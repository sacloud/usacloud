package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func MobileGatewaySimDelete(ctx command.Context, params *params.SimDeleteMobileGatewayParam) error {

	client := ctx.GetAPIClient()
	api := client.GetMobileGatewayAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("MobileGatewaySimDelete is failed: %s", e)
	}

	// clear IPAddress
	simAPI := client.GetSIMAPI()
	_, err := simAPI.ClearIP(params.SimId)
	if err != nil {
		return fmt.Errorf("MobileGatewaySimDelete is failed: %s", err)
	}

	_, err = api.DeleteSIM(p.ID, params.SimId)
	if err != nil {
		return fmt.Errorf("MobileGatewaySimDelete is failed: %s", err)
	}

	return nil
}
