package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
)

func InternetSubnetDelete(ctx command.Context, params *params.SubnetDeleteInternetParam) error {

	client := ctx.GetAPIClient()
	api := client.GetInternetAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("InternetSubnetDelete is failed: %s", e)
	}

	err := internal.ExecWithProgress(
		fmt.Sprintf("Still deleting[ID:%d]...", params.SubnetId),
		fmt.Sprintf("Delete subnet[ID:%d]", params.SubnetId),
		command.GlobalOption.Progress,
		func(compChan chan bool, errChan chan error) {
			_, err := api.DeleteSubnet(params.Id, params.SubnetId)
			if err != nil {
				errChan <- err
				return
			}

			compChan <- true
		},
	)
	if err != nil {
		return fmt.Errorf("InternetSubnetDelete is failed: %s", err)
	}
	return nil
}
