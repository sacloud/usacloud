package command

import (
	"fmt"
)

func ServerReset(ctx Context, params *ResetServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerReset is failed: %s", e)
	}

	// call manipurate functions
	_, err := api.RebootForce(params.Id)
	if err != nil {
		return fmt.Errorf("ServerReset is failed: %s", err)
	}

	// wait for boot
	if !params.Async {
		err := api.SleepUntilUp(params.Id, client.DefaultTimeoutDuration)
		if err != nil {
			return fmt.Errorf("ServerReset is failed: %s", err)
		}
	}

	return nil
	// return ctx.GetOutput().Print(p)

}
