package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func MobileGatewaySimRouteInfo(ctx command.Context, params *params.SimRouteInfoMobileGatewayParam) error {

	client := ctx.GetAPIClient()
	api := client.GetMobileGatewayAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("MobileGatewaySimRouteInfo is failed: %s", e)
	}

	routes, err := api.GetSIMRoutes(params.Id)
	if err != nil {
		return fmt.Errorf("MobileGatewaySimRouteInfo is failed: %s", err)
	}

	if len(routes) == 0 {
		fmt.Fprintf(command.GlobalOption.Err, "MobileGateway[%d] don't have any SIM routes\n", params.Id)
		return nil
	}

	// build parameters to display table
	list := []interface{}{}
	for i := range routes {
		list = append(list, &routes[i])
	}

	return ctx.GetOutput().Print(list...)

}
