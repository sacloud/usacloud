package funcs

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"strings"
)

func VPCRouterDhcpServerInfo(ctx command.Context, params *params.DhcpServerInfoVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterDhcpServerInfo is failed: %s", e)
	}

	if !p.HasDHCPServer() {
		fmt.Fprintf(command.GlobalOption.Err, "VPCRouter[%d] don't have any DHCP servers\n", params.Id)
		return nil
	}

	confList := p.Settings.Router.DHCPServer.Config
	// build parameters to display table
	list := []interface{}{}
	for i := range confList {
		list = append(list, &struct {
			*sacloud.VPCRouterDHCPServerConfig
			DNSServerList string
		}{
			VPCRouterDHCPServerConfig: confList[i],
			DNSServerList:             strings.Join(confList[i].DNSServers, ","),
		})
	}

	return ctx.GetOutput().Print(list...)

}
