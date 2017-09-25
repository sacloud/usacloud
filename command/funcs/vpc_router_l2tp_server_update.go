package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterL2tpServerUpdate(ctx command.Context, params *params.L2tpServerUpdateVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterL2tpServerUpdate is failed: %s", e)
	}

	if !params.Disabled {
		// validate
		targets := []string{"range-start", "range-stop", "pre-shared-secret"}
		for _, t := range targets {
			if !ctx.IsSet(t) {
				return fmt.Errorf("%q: is required when enabled is true", t)
			}
		}

		if !p.HasL2TPIPsecServer() {
			p.InitVPCRouterSetting()
		}
		p.Settings.Router.EnableL2TPIPsecServer(params.PreSharedSecret, params.RangeStart, params.RangeStop)
	} else {
		// validate
		targets := []string{"range-start", "range-stop", "pre-shared-secret"}
		for _, t := range targets {
			if ctx.IsSet(t) {
				return fmt.Errorf("%q: can't set when enabled is false", t)
			}
		}
		if !p.HasL2TPIPsecServer() {
			return nil
		}

		p.Settings.Router.DisableL2TPIPsecServer()
	}

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterL2tpServerUpdate is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterL2tpServerUpdate is failed: %s", err)
	}

	return nil
}
