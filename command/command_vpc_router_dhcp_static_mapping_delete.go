package command

import (
	"fmt"
)

func VPCRouterDhcpStaticMappingDelete(ctx Context, params *DhcpStaticMappingDeleteVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterDhcpStaticMappingDelete is failed: %s", e)
	}

	if !p.HasDHCPStaticMapping() {
		return fmt.Errorf("VPCRouter[%d] don't have any DHCP static mappings", params.Id)
	}

	// validate
	if params.Index <= 0 || params.Index-1 >= len(p.Settings.Router.DHCPStaticMapping.Config) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	p.Settings.Router.RemoveDHCPStaticMappingAt(params.Index - 1)

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterDhcpStaticMappingDelete is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterDhcpStaticMappingDelete is failed: %s", err)
	}

	return nil

}
