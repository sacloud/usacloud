package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
)

func InternetSubnetUpdate(ctx command.Context, params *params.SubnetUpdateInternetParam) error {

	client := ctx.GetAPIClient()
	api := client.GetInternetAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("InternetSubnetUpdate is failed: %s", e)
	}

	var sn *subnet
	err := internal.ExecWithProgress(
		fmt.Sprintf("Still updating[ID:%d]...", params.SubnetId),
		fmt.Sprintf("Update subnet[ID:%d]", params.SubnetId),
		command.GlobalOption.Progress,
		func(compChan chan bool, errChan chan error) {
			subnet, err := api.UpdateSubnet(params.Id, params.SubnetId, params.NextHop)
			if err != nil {
				errChan <- err
				return
			}
			sn, err = getSubnetByID(ctx, subnet.ID)
			if err != nil {
				errChan <- err
				return
			}
			compChan <- true
		},
	)
	if err != nil {
		return fmt.Errorf("InternetSubnetUpdate is failed: %s", err)
	}
	return ctx.GetOutput().Print(sn)
}
