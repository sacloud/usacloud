package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterDhcpStaticMappingInfo(ctx command.Context, params *params.DhcpStaticMappingInfoVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterDhcpStaticMappingInfo is failed: %s", e)
	}

	if !p.HasDHCPStaticMapping() {
		fmt.Fprintf(command.GlobalOption.Err, "VPCRouter[%d] don't have any DHCP static mappings\n", params.Id)
		return nil
	}

	confList := p.Settings.Router.DHCPStaticMapping.Config
	// build parameters to display table
	list := []interface{}{}
	for i := range confList {
		list = append(list, &confList[i])
	}

	return ctx.GetOutput().Print(list...)

}
