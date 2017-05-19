package command

import (
	"fmt"
)

func VPCRouterPortForwardingDelete(ctx Context, params *PortForwardingDeleteVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterPortForwardingDelete is failed: %s", e)
	}

	if !p.HasPortForwarding() {
		return fmt.Errorf("VPCRouter[%d] don't have any port-forwarding settings", params.Id)
	}

	// validate
	if params.Index <= 0 || params.Index-1 >= len(p.Settings.Router.PortForwarding.Config) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	pfConfig := p.Settings.Router.PortForwarding.Config[params.Index-1]
	p.Settings.Router.RemovePortForwarding(
		pfConfig.Protocol,
		pfConfig.GlobalPort,
		pfConfig.PrivateAddress,
		pfConfig.PrivatePort,
	)

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterPortForwardingDelete is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterPortForwardingDelete is failed: %s", err)
	}

	return nil

}
