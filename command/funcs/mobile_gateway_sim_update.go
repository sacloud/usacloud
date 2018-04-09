package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func MobileGatewaySimUpdate(ctx command.Context, params *params.SimUpdateMobileGatewayParam) error {

	client := ctx.GetAPIClient()
	api := client.GetMobileGatewayAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("MobileGatewaySimUpdate is failed: %s", e)
	}

	// set IPAddress
	simAPI := client.GetSIMAPI()
	_, err := simAPI.AssignIP(params.SimId, params.Ipaddress)
	if err != nil {
		return fmt.Errorf("MobileGatewaySimUpdate is failed: %s", err)
	}

	return nil

}
