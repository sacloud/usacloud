package command

import (
	"fmt"
)

func ServerWaitForBoot(ctx Context, params *WaitForBootServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerWaitForBoot is failed: %s", e)
	}

	if p.IsUp() {
		return nil // already booted.
	}

	// wait for boot
	err := api.SleepUntilUp(params.Id, client.DefaultTimeoutDuration)
	if err != nil {
		return fmt.Errorf("ServerWaitForBoot is failed: %s", err)
	}

	return nil
	// return ctx.GetOutput().Print(p)
}
