package command

import (
	"fmt"
	"github.com/sacloud/usacloud/command/internal"
)

func LoadBalancerDelete(ctx Context, params *DeleteLoadBalancerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetLoadBalancerAPI()

	p, err := api.Read(params.Id)
	if err != nil {
		return fmt.Errorf("ServerDelete is failed: %s", err)
	}

	if p.IsUp() {
		if params.Force {

			compChan := make(chan bool)
			errChan := make(chan error)
			spinner := internal.NewProgress(
				fmt.Sprintf("Still waiting for delete[ID:%d]...", params.Id),
				fmt.Sprintf("Delete server[ID:%d]", params.Id),
				GlobalOption.Progress)

			go func() {
				spinner.Start()
				// call manipurate functions
				var err error
				_, err = api.Stop(params.Id)
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
					return fmt.Errorf("LoadBalancerDelete is failed: %s", err)
				}
			}

		} else {
			return fmt.Errorf("LoadBalancer(%d) is still running", params.Id)
		}
	}

	// call Delete(id)
	res, err := api.Delete(params.Id)
	if err != nil {
		return fmt.Errorf("LoadBalancerDelete is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)
}
