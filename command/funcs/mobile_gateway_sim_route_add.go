package funcs

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func MobileGatewaySimRouteAdd(ctx command.Context, params *params.SimRouteAddMobileGatewayParam) error {

	client := ctx.GetAPIClient()
	api := client.GetMobileGatewayAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("MobileGatewaySimRouteAdd is failed: %s", e)
	}

	routeList, err := api.GetSIMRoutes(params.Id)
	if err != nil {
		return fmt.Errorf("MobileGatewaySimRouteAdd is failed: %s", err)
	}

	routes := &sacloud.MobileGatewaySIMRoutes{
		SIMRoutes: routeList,
	}

	if _, exists := routes.FindSIMRoute(params.Sim, params.Prefix); exists != nil {
		fmt.Fprintf(command.GlobalOption.Out, "SIM Route[%s -> %d] already exists", params.Prefix, params.Sim)
		return nil
	}

	if _, err := client.GetSIMAPI().Read(params.Sim); err != nil {
		return fmt.Errorf("SIM[%d] is not found: %s", params.Sim, err)
	}

	if _, err := api.AddSIMRoute(params.Id, params.Sim, params.Prefix); err != nil {
		return fmt.Errorf("MobileGatewaySimRouteAdd is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("MobileGatewaySimRouteAdd is failed: %s", err)
	}
	return nil
}
