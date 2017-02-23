package command

import (
	"fmt"
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

	// wait for down
	err := api.SleepUntilDown(params.Id, client.DefaultTimeoutDuration)
	if err != nil {
		return fmt.Errorf("ServerWaitForDown is failed: %s", err)
	}

	return nil

}
