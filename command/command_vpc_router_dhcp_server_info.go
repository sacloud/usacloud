package command

import (
	"fmt"
)

func VPCRouterDhcpServerInfo(ctx Context, params *DhcpServerInfoVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterDhcpServerInfo is failed: %s", e)
	}

	if !p.HasDHCPServer() {
		fmt.Fprintf(GlobalOption.Err, "VPCRouter[%d] don't have any DHCP servers\n", params.Id)
		return nil
	}

	confList := p.Settings.Router.DHCPServer.Config
	// build parameters to display table
	list := []interface{}{}
	for i := range confList {
		list = append(list, &confList[i])
	}

	return ctx.GetOutput().Print(list...)

}
