package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterStaticRouteAdd(ctx command.Context, params *params.StaticRouteAddVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterStaticRouteAdd is failed: %s", e)
	}

	if !p.HasSetting() {
		p.InitVPCRouterSetting()
	}

	p.Settings.Router.AddStaticRoute(params.Prefix, params.NextHop)

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterStaticRouteAdd is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterStaticRouteAdd is failed: %s", err)
	}
	return nil

}
