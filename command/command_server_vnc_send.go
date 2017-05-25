package command

import (
	"fmt"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/vnc"
	"io/ioutil"
	"strings"
)

func ServerVncSend(ctx Context, params *VncSendServerParam) error {

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
			GlobalOption.Progress,
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

	command := ""
	if params.CommandFile != "" {
		b, err := ioutil.ReadFile(params.CommandFile)
		if err != nil {
			return fmt.Errorf("ServerVncSend is failed: %s", err)
		}
		command = string(b)
	}
	if params.Command != "" {
		command = params.Command
	}
	if command == "" {
		return fmt.Errorf("Command or CommandFile is required")
	}
	// remove newline from command
	command = strings.Replace(command, "\r", "", -1)
	command = strings.Replace(command, "\n", "", -1)

	// create option
	option := vnc.NewSendCommandOption()
	option.UseUSKeyboard = params.UseUsKeyboard
	option.Debug = params.Debug
	option.ProgressWriter = GlobalOption.Progress

	// VNCProxy(call sacloud API)
	vncProxyInfo, e := api.GetVNCProxy(params.Id)
	if e != nil {
		return fmt.Errorf("ServerVncSend is failed: %s", e)
	}

	return vnc.SendCommand(vncProxyInfo, command, option)
}
