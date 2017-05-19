package command

import (
	"fmt"
)

func VPCRouterDhcpStaticMappingAdd(ctx Context, params *DhcpStaticMappingAddVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterDhcpStaticMappingAdd is failed: %s", e)
	}

	if !p.HasSetting() {
		p.InitVPCRouterSetting()
	}

	p.Settings.Router.AddDHCPStaticMapping(
		params.Ipaddress,
		params.Macaddress,
	)

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterDhcpStaticMappingAdd is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterDhcpStaticMappingAdd is failed: %s", err)
	}
	return nil

}
