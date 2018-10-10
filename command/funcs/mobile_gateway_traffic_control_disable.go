package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func MobileGatewayTrafficControlDisable(ctx command.Context, params *params.TrafficControlDisableMobileGatewayParam) error {

	client := ctx.GetAPIClient()
	api := client.GetMobileGatewayAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("MobileGatewayTrafficControlDisable is failed: %s", e)
	}

	// set params
	_, err := api.DisableTrafficMonitoringConfig(params.Id)
	if err != nil {
		return fmt.Errorf("MobileGatewayTrafficControlDisable is failed: %s", err)
	}
	return nil
}
