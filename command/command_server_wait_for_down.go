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

	compChan := make(chan bool)
	errChan := make(chan error)
	spinner := internal.NewProgress(
		fmt.Sprintf("Still waiting for Shutdown[ID:%d]...", params.Id),
		fmt.Sprintf("Shutdown server[ID:%d]", params.Id),
		GlobalOption.Progress)

	go func() {
		spinner.Start()
		// call manipurate functions
		err := api.SleepUntilDown(params.Id, client.DefaultTimeoutDuration)
		if err != nil {
			errChan <- err
			return
		}
		compChan <- true
	}()

down:
	for {
		select {
		case <-compChan:
			spinner.Stop()
			break down
		case err := <-errChan:
			return fmt.Errorf("ServerWaitForDown is failed: %s", err)
		}
	}

	return nil
}
