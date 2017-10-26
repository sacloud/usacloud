package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
)

func InternetSubnetAdd(ctx command.Context, params *params.SubnetAddInternetParam) error {

	client := ctx.GetAPIClient()
	api := client.GetInternetAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("InternetSubnetCreate is failed: %s", e)
	}

	var sn *subnet
	err := internal.ExecWithProgress(
		"Still creating...",
		"Add subnet",
		command.GlobalOption.Progress,
		func(compChan chan bool, errChan chan error) {
			subnet, err := api.AddSubnet(params.Id, params.NwMasklen, params.NextHop)
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
		return fmt.Errorf("InternetSubnetCreate is failed: %s", err)
	}

	return ctx.GetOutput().Print(sn)

}
