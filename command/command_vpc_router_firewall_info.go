package command

import (
	"fmt"
)

func VPCRouterFirewallInfo(ctx Context, params *FirewallInfoVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterFirewallInfo is failed: %s", e)
	}

	if !p.HasFirewall() {
		fmt.Fprintf(GlobalOption.Err, "VPCRouter[%d] don't have any firewall rules\n", params.Id)
		return nil
	}

	ruleList := p.Settings.Router.Firewall.Config[0].Receive
	switch params.Direction {
	case "send":
		ruleList = p.Settings.Router.Firewall.Config[0].Send
	case "receive":
		ruleList = p.Settings.Router.Firewall.Config[0].Receive
	}

	if len(ruleList) == 0 {
		fmt.Fprintf(GlobalOption.Err, "VPCRouter don't have any firewall(%s) rules\n", params.Direction)
		return nil
	}

	// build parameters to display table
	list := []interface{}{}
	for i := range ruleList {
		list = append(list, &ruleList[i])
	}

	return ctx.GetOutput().Print(list...)

}
