package funcs

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func MobileGatewaySimRouteUpdate(ctx command.Context, params *params.SimRouteUpdateMobileGatewayParam) error {

	client := ctx.GetAPIClient()
	api := client.GetMobileGatewayAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("MobileGatewaySimRouteUpdate is failed: %s", e)
	}

	routes, err := api.GetSIMRoutes(params.Id)
	if err != nil {
		return fmt.Errorf("MobileGatewaySimRouteUpdate is failed: %s", err)
	}

	if len(routes) == 0 {
		fmt.Fprintf(command.GlobalOption.Err, "MobileGateway[%d] don't have any SIM routes\n", params.Id)
		return nil
	}

	// validate
	if params.Index <= 0 || params.Index-1 >= len(routes) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	route := routes[params.Index-1]
	if ctx.IsSet("prefix") {
		route.Prefix = params.Prefix
	}
	if ctx.IsSet("sim") {
		route.ResourceID = fmt.Sprintf("%d", params.Sim)
	}

	simRoutes := &sacloud.MobileGatewaySIMRoutes{
		SIMRoutes: routes,
	}

	if _, err := api.SetSIMRoutes(params.Id, simRoutes); err != nil {
		return fmt.Errorf("MobileGatewayStaticRouteUpdate is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("MobileGatewayStaticRouteUpdate is failed: %s", err)
	}

	return nil
}
