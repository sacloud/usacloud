package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterStaticRouteInfo(ctx command.Context, params *params.StaticRouteInfoVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterStaticRouteInfo is failed: %s", e)
	}

	if !p.HasStaticRoutes() {
		fmt.Fprintf(command.GlobalOption.Err, "VPCRouter[%d] don't have any static-routes\n", params.Id)
		return nil
	}

	confList := p.Settings.Router.StaticRoutes.Config
	// build parameters to display table
	list := []interface{}{}
	for i := range confList {
		list = append(list, &confList[i])
	}

	return ctx.GetOutput().Print(list...)
}
