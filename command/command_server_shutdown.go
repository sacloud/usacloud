package command

import (
	"fmt"
	"github.com/sacloud/usacloud/command/internal"
)

func ServerShutdown(ctx Context, params *ShutdownServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerShutdown is failed: %s", e)
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
			compChan <- true
		},
	)
	if err != nil {
		return fmt.Errorf("ServerShutdown is failed: %s", err)
	}

	return nil
}
