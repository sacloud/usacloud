package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterFirewallAdd(ctx command.Context, params *params.FirewallAddVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterFirewallAdd is failed: %s", e)
	}

	if !p.HasSetting() {
		p.InitVPCRouterSetting()
	}

	// validate
	if params.Protocol == "icmp" || params.Protocol == "ip" {
		if params.SourcePort > 0 {
			return fmt.Errorf("%q: can't set when protocol is [icmp/ip]", "source-port")
		}
		if params.DestinationPort > 0 {
			return fmt.Errorf("%q: can't set when protocol is [icmp/ip]", "destination-port")
		}
	}

	// build parameters
	isAllow := false
	if params.Action == "allow" {
		isAllow = true
	}
	sourcePort := ""
	destPort := ""
	if params.Protocol == "tcp" || params.Protocol == "udp" {
		if params.SourcePort > 0 {
			sourcePort = fmt.Sprintf("%d", params.SourcePort)
		}
		if params.DestinationPort > 0 {
			destPort = fmt.Sprintf("%d", params.DestinationPort)
		}
	}

	var f func(isAllow bool, protocol string, sourceNetwork string, sourcePort string, destNetwork string, destPort string, logging bool, description string)
	switch params.Direction {
	case "send":
		f = p.Settings.Router.AddFirewallRuleSend
	case "receive":
		f = p.Settings.Router.AddFirewallRuleReceive
	}
	// add
	f(
		isAllow,
		params.Protocol,
		params.SourceNetwork,
		sourcePort,
		params.DestinationNetwork,
		destPort,
		params.EnableLogging,
		params.Description,
	)

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterFirewallAdd is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterFirewallAdd is failed: %s", err)
	}
	return nil

}
