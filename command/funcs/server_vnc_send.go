package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/helper/vnc"
	"io/ioutil"
	"strings"
)

func ServerVncSend(ctx command.Context, params *params.VncSendServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerVncSend is failed: %s", e)
	}

	if !p.IsUp() && params.WaitForBoot {

		err := internal.ExecWithProgress(
			fmt.Sprintf("Still booting[ID:%d]...", params.Id),
			fmt.Sprintf("Connect to server[ID:%d]", params.Id),
			command.GlobalOption.Progress,
			func(compChan chan bool, errChan chan error) {
				// call manipurate functions
				err := api.SleepUntilUp(params.Id, client.DefaultTimeoutDuration)
				if err != nil {
					errChan <- err
					return
				}
				compChan <- true
			},
		)
		if err != nil {
			return fmt.Errorf("ServerVncSend is failed: %s", err)
		}
	}

	cmd := ""
	if params.CommandFile != "" {
		b, err := ioutil.ReadFile(params.CommandFile)
		if err != nil {
			return fmt.Errorf("ServerVncSend is failed: %s", err)
		}
		cmd = string(b)
	}
	if params.Command != "" {
		cmd = params.Command
	}
	if cmd == "" {
		return fmt.Errorf("Command or CommandFile is required")
	}
	// remove newline from command
	cmd = strings.Replace(cmd, "\r", "", -1)
	cmd = strings.Replace(cmd, "\n", "", -1)

	// create option
	option := vnc.NewSendCommandOption()
	option.UseUSKeyboard = params.UseUsKeyboard
	option.Debug = params.Debug
	option.ProgressWriter = command.GlobalOption.Progress

	// VNCProxy(call sacloud API)
	vncProxyInfo, e := api.GetVNCProxy(params.Id)
	if e != nil {
		return fmt.Errorf("ServerVncSend is failed: %s", e)
	}

	return vnc.SendCommand(vncProxyInfo, cmd, option)
}
