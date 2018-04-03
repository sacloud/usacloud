package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterDhcpStaticMappingUpdate(ctx command.Context, params *params.DhcpStaticMappingUpdateVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterDhcpStaticMappingUpdate is failed: %s", e)
	}

	if !p.HasDHCPStaticMapping() {
		return fmt.Errorf("VPCRouter[%d] don't have any DHCP static mappings", params.Id)
	}

	// validate
	if params.Index <= 0 || params.Index-1 >= len(p.Settings.Router.DHCPStaticMapping.Config) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	cnf := p.Settings.Router.DHCPStaticMapping.Config[params.Index-1]
	if ctx.IsSet("macaddress") {
		cnf.MACAddress = params.Macaddress
	}
	if ctx.IsSet("ipaddress") {
		cnf.IPAddress = params.Ipaddress
	}

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterDhcpStaticMappingUpdate is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterDhcpStaticMappingUpdate is failed: %s", err)
	}

	return nil

}
