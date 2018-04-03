package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
)

func LoadBalancerShutdown(ctx command.Context, params *params.ShutdownLoadBalancerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetLoadBalancerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("LoadBalancerShutdown is failed: %s", e)
	}

	if p.IsDown() {
		return nil // already downed.
	}

	err := internal.ExecWithProgress(
		fmt.Sprintf("Still waiting for Shutdown[ID:%d]...", params.Id),
		fmt.Sprintf("Shutdown load-balancer[ID:%d]", params.Id),
		command.GlobalOption.Progress,
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
		return fmt.Errorf("LoadBalancerShutdown is failed: %s", err)
	}

	return nil
}
