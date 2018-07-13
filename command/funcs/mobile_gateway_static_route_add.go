package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func MobileGatewayStaticRouteAdd(ctx command.Context, params *params.StaticRouteAddMobileGatewayParam) error {

	client := ctx.GetAPIClient()
	api := client.GetMobileGatewayAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("MobileGatewayStaticRouteAdd is failed: %s", e)
	}

	if _, exists := p.Settings.MobileGateway.FindStaticRoute(params.Prefix, params.NextHop); exists != nil {
		fmt.Fprintf(command.GlobalOption.Out, "StaticRoute[%s -> %s] already exists", params.Prefix, params.NextHop)
		return nil
	}

	p.Settings.MobileGateway.AddStaticRoute(params.Prefix, params.NextHop)
	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("MobileGatewayStaticRouteAdd is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("MobileGatewayStaticRouteAdd is failed: %s", err)
	}
	return nil
}
