package command

import (
	"fmt"
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

	// call manipurate functions
	var err error
	if params.Force {
		_, err = api.Stop(params.Id)
	} else {
		_, err = api.Shutdown(params.Id)
	}
	if err != nil {
		return fmt.Errorf("ServerShutdown is failed: %s", err)
	}

	// wait for down
	if !params.Async {
		err := api.SleepUntilDown(params.Id, client.DefaultTimeoutDuration)
		if err != nil {
			return fmt.Errorf("ServerShutdown is failed: %s", err)
		}
	}

	return nil

}
