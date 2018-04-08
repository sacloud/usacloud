package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func MobileGatewaySimAdd(ctx command.Context, params *params.SimAddMobileGatewayParam) error {

	client := ctx.GetAPIClient()
	api := client.GetMobileGatewayAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("MobileGatewaySimAdd is failed: %s", e)
	}

	_, err := api.AddSIM(p.ID, params.SimId)
	if err != nil {
		return fmt.Errorf("MobileGatewaySimAdd is failed: %s", err)
	}

	// set IPAddress
	simAPI := client.GetSIMAPI()
	_, err = simAPI.AssignIP(params.SimId, params.Ipaddress)
	if err != nil {
		return fmt.Errorf("MobileGatewaySimAdd is failed: %s", err)
	}

	return nil
}
