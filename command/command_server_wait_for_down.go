package command

import (
	"fmt"
	"github.com/sacloud/usacloud/command/internal"
)

func ServerWaitForDown(ctx Context, params *WaitForDownServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerWaitForDown is failed: %s", e)
	}

	if p.IsDown() {
		return nil // already downed.
	}

	err := internal.ExecWithProgress(
		fmt.Sprintf("Still waiting for Shutdown[ID:%d]...", params.Id),
		fmt.Sprintf("Shutdown server[ID:%d]", params.Id),
		GlobalOption.Progress,
		func(compChan chan bool, errChan chan error) {
			// call manipurate functions
			err := api.SleepUntilDown(params.Id, client.DefaultTimeoutDuration)
			if err != nil {
				errChan <- err
				return
			}
			compChan <- true
		},
	)
	if err != nil {
		return fmt.Errorf("ServerWaitForDown is failed: %s", err)
	}

	return nil
}
