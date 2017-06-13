package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
	"strconv"
)

func VPCRouterInterfaceDisconnect(ctx command.Context, params *params.InterfaceDisconnectVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterInterfaceDisconnect is failed: %s", e)
	}
	index, _ := strconv.Atoi(params.Index)

	// validation
	if p.Interfaces[index].GetSwitch() == nil {
		return fmt.Errorf("Interface[%d] is already disconnected from switch", index)
	}
	if p.IsUp() {
		return fmt.Errorf("VPCRouter(%d) is still running", params.Id)
	}

	// disconnect
	_, err := api.DeleteInterfaceAt(params.Id, index)
	if err != nil {
		return fmt.Errorf("VPCRouterInterfaceDisconnect is failed: %s", err)
	}

	if params.WithReboot && p.IsUp() {
		err := internal.ExecWithProgress(
			fmt.Sprintf("Still waiting for reboot[ID:%d]...", params.Id),
			fmt.Sprintf("Disconnecting interface to switch[ID:%d]", params.Id),
			command.GlobalOption.Progress,
			func(compChan chan bool, errChan chan error) {
				// call manipurate functions
				var err error
				_, err = api.Shutdown(params.Id)
				if err != nil {
					errChan <- err
					return
				}

				err = api.SleepUntilDown(params.Id, client.DefaultTimeoutDuration)
				if err != nil {
					errChan <- err
					return
				}

				_, err = api.Boot(params.Id)
				if err != nil {
					errChan <- err
					return
				}
				err = api.SleepUntilUp(params.Id, client.DefaultTimeoutDuration)
				if err != nil {
					errChan <- err
					return
				}

				compChan <- true
			},
		)
		if err != nil {
			return fmt.Errorf("VPCRouterInterfaceDisconnect is failed: %s", err)
		}
	}

	return nil

}
