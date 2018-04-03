package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func PacketFilterRuleAdd(ctx command.Context, params *params.RuleAddPacketFilterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetPacketFilterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("PacketFilterRuleAdd is failed: %s", e)
	}

	// validate
	switch params.Protocol {
	case "icmp", "fragment", "ip":
		if params.SourcePort != "" || params.DestinationPort != "" {
			return fmt.Errorf("When protocol is %s , can't set Source/Destination port", params.Protocol)
		}
	}

	// index
	if params.Index <= 0 || params.Index-1 >= len(p.Expression) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	action := true
	switch params.Action {
	case "allow":
		action = true
	case "deny":
		action = false
	}

	index := params.Index - 1

	switch params.Protocol {
	case "tcp":
		p.AddTCPRuleAt(params.SourceNetwork, params.SourcePort, params.DestinationPort, params.Description, action, index)
	case "udp":
		p.AddUDPRuleAt(params.SourceNetwork, params.SourcePort, params.DestinationPort, params.Description, action, index)
	case "icmp":
		p.AddICMPRuleAt(params.SourceNetwork, params.Description, action, index)
	case "fragment":
		p.AddFragmentRuleAt(params.SourceNetwork, params.Description, action, index)
	case "ip":
		p.AddIPRuleAt(params.SourceNetwork, params.Description, action, index)
	default:
		panic(fmt.Errorf("Unknown protocol %s", params.Protocol))
	}

	// call manipurate functions
	p, err := api.Update(params.Id, p)
	if err != nil {
		return fmt.Errorf("PacketFilterRuleAdd is failed: %s", err)
	}

	list := []interface{}{}
	for i := range p.Expression {
		list = append(list, p.Expression[i])
	}

	return ctx.GetOutput().Print(list...)

}
