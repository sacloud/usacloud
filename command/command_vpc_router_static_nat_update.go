package command

import (
	"fmt"
)

func VPCRouterStaticNatUpdate(ctx Context, params *StaticNatUpdateVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterStaticNatUpdate is failed: %s", e)
	}
	if p.IsStandardPlan() {
		return fmt.Errorf("Static NAT is not supported on standard plan")
	}
	if !p.HasStaticNAT() {
		return fmt.Errorf("VPCRouter[%d] don't have any static NAT settings", params.Id)
	}

	// validate

	// index
	if params.Index <= 0 || params.Index-1 >= len(p.Settings.Router.StaticNAT.Config) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	natConfig := p.Settings.Router.StaticNAT.Config[params.Index-1]
	if ctx.IsSet("global") {
		natConfig.GlobalAddress = params.Global
	}
	if ctx.IsSet("private") {
		natConfig.PrivateAddress = params.Private
	}
	if ctx.IsSet("description") {
		natConfig.Description = params.Description
	}

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterStaticNatUpdate is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterStaticNatUpdate is failed: %s", err)
	}

	return nil
}
