package command

import (
	"fmt"
)

func VPCRouterFirewallDelete(ctx Context, params *FirewallDeleteVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterFirewallDelete is failed: %s", e)
	}

	if !p.HasFirewall() {
		return fmt.Errorf("VPCRouter[%d] don't have any firewall rules[%s]", params.Id, params.Direction)
	}

	// validate
	if params.Index <= 0 {
		cnt := len(p.Settings.Router.Firewall.Config[0].Receive)
		if params.Direction == "send" {
			cnt = len(p.Settings.Router.Firewall.Config[0].Send)
		}

		if params.Index-1 >= cnt {
			return fmt.Errorf("index(%d) is out of range", params.Index)
		}
	}

	switch params.Direction {
	case "send":
		p.Settings.Router.RemoveFirewallRuleSendAt(params.Index - 1)
	case "receive":
		p.Settings.Router.RemoveFirewallRuleReceiveAt(params.Index - 1)
	}

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterFirewallDelete is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterFirewallDelete is failed: %s", err)
	}

	return nil

}
