package command

import (
	"fmt"
)

func VPCRouterStaticRouteDelete(ctx Context, params *StaticRouteDeleteVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterStaticRouteDelete is failed: %s", e)
	}

	if !p.HasStaticRoutes() {
		return fmt.Errorf("VPCRouter[%d] don't have any static-routes", params.Id)
	}

	// validate
	if params.Index <= 0 || params.Index-1 >= len(p.Settings.Router.StaticRoutes.Config) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	p.Settings.Router.RemoveStaticRouteAt(params.Index - 1)

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterStaticRouteDelete is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterStaticRouteDelete is failed: %s", err)
	}

	return nil
}
