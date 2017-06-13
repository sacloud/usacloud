package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterPortForwardingInfo(ctx command.Context, params *params.PortForwardingInfoVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterPortForwardingInfo is failed: %s", e)
	}

	if !p.HasPortForwarding() {
		fmt.Fprintf(command.GlobalOption.Err, "VPCRouter[%d] don't have any port-forwarding settings\n", params.Id)
		return nil
	}

	confList := p.Settings.Router.PortForwarding.Config
	// build parameters to display table
	list := []interface{}{}
	for i := range confList {
		list = append(list, &confList[i])
	}

	return ctx.GetOutput().Print(list...)

}
