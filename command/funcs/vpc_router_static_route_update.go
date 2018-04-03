package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterStaticRouteUpdate(ctx command.Context, params *params.StaticRouteUpdateVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterStaticRouteUpdate is failed: %s", e)
	}

	if !p.HasStaticRoutes() {
		return fmt.Errorf("VPCRouter[%d] don't have any static-routes", params.Id)
	}

	// validate
	if params.Index <= 0 || params.Index-1 >= len(p.Settings.Router.StaticRoutes.Config) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	cnf := p.Settings.Router.StaticRoutes.Config[params.Index-1]
	if ctx.IsSet("prefix") {
		cnf.Prefix = params.Prefix
	}
	if ctx.IsSet("next-hop") {
		cnf.NextHop = params.NextHop
	}

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterStaticRouteUpdate is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterStaticRouteUpdate is failed: %s", err)
	}

	return nil

}
