package command

import (
	"fmt"
	"github.com/sacloud/usacloud/command/internal"
)

func VPCRouterDelete(ctx Context, params *DeleteVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()

	p, err := api.Read(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterDelete is failed: %s", err)
	}

	if p.IsUp() {
		if params.Force {

			err = internal.ExecWithProgress(
				fmt.Sprintf("Still waiting for delete[ID:%d]...", params.Id),
				fmt.Sprintf("Delete vpc-router[ID:%d]", params.Id),
				GlobalOption.Progress,
				func(compChan chan bool, errChan chan error) {
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
				},
			)
			if err != nil {
				return fmt.Errorf("VPCRouterDelete is failed: %s", err)
			}

		} else {
			return fmt.Errorf("VPCRouter(%d) is still running", params.Id)
		}
	}

	// call Delete(id)
	res, err := api.Delete(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterDelete is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}
