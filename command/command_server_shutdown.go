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

	compChan := make(chan bool)
	errChan := make(chan error)
	spinner := internal.NewSpinner(
		"Waiting for Shutdown...",
		"Shutdown is complete.\n",
		internal.CharSetProgress,
		GlobalOption.Progress)

	go func() {
		spinner.Start()
		// call manipurate functions
		var err error
		if params.Force {
			_, err = api.Stop(params.Id)
		} else {
			_, err = api.Shutdown(params.Id)
		}
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
	}()

down:
	for {
		select {
		case <-compChan:
			spinner.Stop()
			break down
		case err := <-errChan:
			return fmt.Errorf("ServerShutdown is failed: %s", err)
		}
	}

	return nil

}
