package command

import (
	"fmt"
)

func VPCRouterPortForwardingAdd(ctx Context, params *PortForwardingAddVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterPortForwardingAdd is failed: %s", e)
	}

	if !p.HasSetting() {
		p.InitVPCRouterSetting()
	}

	p.Settings.Router.AddPortForwarding(
		params.Protocol,
		fmt.Sprintf("%d", params.GlobalPort),
		params.PrivateIpaddress,
		fmt.Sprintf("%d", params.PrivatePort),
		params.Description,
	)

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterPortForwardingAdd is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterPortForwardingAdd is failed: %s", err)
	}
	return nil

}
