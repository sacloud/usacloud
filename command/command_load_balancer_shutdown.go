package command

import (
	"fmt"
	"github.com/sacloud/usacloud/command/internal"
)

func LoadBalancerShutdown(ctx Context, params *ShutdownLoadBalancerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetLoadBalancerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("LoadBalancerShutdown is failed: %s", e)
	}

	if p.IsDown() {
		return nil // already downed.
	}

	compChan := make(chan bool)
	errChan := make(chan error)
	spinner := internal.NewProgress(
		fmt.Sprintf("Still waiting for Shutdown[ID:%d]...", params.Id),
		fmt.Sprintf("Shutdown LoadBalancer[ID:%d]", params.Id),
		GlobalOption.Progress)

	go func() {
		spinner.Start()
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
	}()

down:
	for {
		select {
		case <-compChan:
			spinner.Stop()
			break down
		case err := <-errChan:
			return fmt.Errorf("LoadBalancerShutdown is failed: %s", err)
		}
	}

	return nil

}
