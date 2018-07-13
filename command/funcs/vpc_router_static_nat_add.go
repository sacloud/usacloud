package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterStaticNatAdd(ctx command.Context, params *params.StaticNatAddVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterStaticNatAdd is failed: %s", e)
	}

	if p.IsStandardPlan() {
		return fmt.Errorf("Static NAT is not supported on standard plan")
	}

	if !p.HasSetting() {
		p.InitVPCRouterSetting()
	}

	// validate
	p.Settings.Router.AddStaticNAT(params.Global, params.Private, params.Description)

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterStaticNatAdd is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterStaticNatAdd is failed: %s", err)
	}
	return nil
}
