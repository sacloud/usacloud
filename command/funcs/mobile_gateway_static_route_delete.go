package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func MobileGatewayStaticRouteDelete(ctx command.Context, params *params.StaticRouteDeleteMobileGatewayParam) error {

	client := ctx.GetAPIClient()
	api := client.GetMobileGatewayAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("MobileGatewayStaticRouteDelete is failed: %s", e)
	}

	if !p.HasStaticRoutes() {
		return fmt.Errorf("MobileGateway[%d] don't have any static-routes", params.Id)
	}

	// validate
	if params.Index <= 0 || params.Index-1 >= len(p.Settings.MobileGateway.StaticRoutes) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	p.Settings.MobileGateway.RemoveStaticRouteAt(params.Index - 1)

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("MobileGatewayStaticRouteDelete is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("MobileGatewayStaticRouteDelete is failed: %s", err)
	}

	return nil
}
