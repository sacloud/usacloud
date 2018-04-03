package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterDhcpServerUpdate(ctx command.Context, params *params.DhcpServerUpdateVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterDhcpServerUpdate is failed: %s", e)
	}

	if !p.HasDHCPServer() {
		return fmt.Errorf("VPCRouter[%d] don't have any DHCP servers", params.Id)
	}

	_, cnf := p.Settings.Router.FindDHCPServerAt(params.Interface)
	if cnf == nil {
		return fmt.Errorf("DHCP server is not found on eth%d", params.Interface)
	}

	if ctx.IsSet("range-start") {
		cnf.RangeStart = params.RangeStart
	}
	if ctx.IsSet("range-stop") {
		cnf.RangeStop = params.RangeStop
	}
	if ctx.IsSet("dns_servers") {
		cnf.DNSServers = params.DnsServers
	}

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterDhcpServerUpdate is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterDhcpServerUpdate is failed: %s", err)
	}

	return nil

}
