package command

import (
	"fmt"
)

func ServerBoot(ctx Context, params *BootServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerBoot is failed: %s", e)
	}

	if p.IsUp() {
		return nil // already booted.
	}

	// call manipurate functions
	_, err := api.Boot(params.Id)
	if err != nil {
		return fmt.Errorf("ServerBoot is failed: %s", err)
	}

	// wait for boot
	if !params.Async {
		err := api.SleepUntilUp(params.Id, client.DefaultTimeoutDuration)
		if err != nil {
			return fmt.Errorf("ServerBoot is failed: %s", err)
		}
	}

	return nil
	// return ctx.GetOutput().Print(p)

}
