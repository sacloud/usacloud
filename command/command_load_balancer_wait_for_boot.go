package command

import (
	"fmt"
	"github.com/sacloud/usacloud/command/internal"
)

func LoadBalancerWaitForBoot(ctx Context, params *WaitForBootLoadBalancerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetLoadBalancerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("LoadBalancerWaitForBoot is failed: %s", e)
	}

	if p.IsUp() {
		return nil // already booted.
	}

	compChan := make(chan bool)
	errChan := make(chan error)
	spinner := internal.NewProgress(
		fmt.Sprintf("Still booting[ID:%d]...", params.Id),
		fmt.Sprintf("Boot LoadBalancer[ID:%d]", params.Id),
		GlobalOption.Progress)

	go func() { // wait for copy with progress

		spinner.Start()
		// call manipurate functions
		err := api.SleepUntilUp(params.Id, client.DefaultTimeoutDuration)
		if err != nil {
			errChan <- err
			return
		}
		compChan <- true
	}()

boot:
	for {
		select {
		case <-compChan:
			spinner.Stop()
			break boot
		case err := <-errChan:
			return fmt.Errorf("LoadBalancerWaitForBoot is failed: %s", err)
		}
	}

	return nil

}
