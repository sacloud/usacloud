package command

import (
	"fmt"
)

func VPCRouterPptpServerUpdate(ctx Context, params *PptpServerUpdateVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterPptpServerUpdate is failed: %s", e)
	}

	if params.Enabled == "true" {
		// validate
		targets := []string{"range-start", "range-stop"}
		for _, t := range targets {
			if !ctx.IsSet(t) {
				return fmt.Errorf("%q: is required when enabled is true", t)
			}
		}

		if !p.HasPPTPServer() {
			p.InitVPCRouterSetting()
		}
		p.Settings.Router.EnablePPTPServer(params.RangeStart, params.RangeStop)
	} else {
		// validate
		targets := []string{"range-start", "range-stop"}
		for _, t := range targets {
			if ctx.IsSet(t) {
				return fmt.Errorf("%q: can't set when enabled is false", t)
			}
		}
		if !p.HasPPTPServer() {
			return nil
		}

		p.Settings.Router.DisablePPTPServer()
	}

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterPptpServerUpdate is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterPptpServerUpdate is failed: %s", err)
	}

	return nil
}
