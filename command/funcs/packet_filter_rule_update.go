package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func PacketFilterRuleUpdate(ctx command.Context, params *params.RuleUpdatePacketFilterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetPacketFilterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("PacketFilterRuleUpdate is failed: %s", e)
	}

	if len(p.Expression) == 0 {
		fmt.Fprintf(command.GlobalOption.Err, "PacketFilter1 don't have any rules\n")
		return nil
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

	index := params.Index - 1

	exp := p.Expression[index]

	if ctx.IsSet("protocol") {
		exp.Protocol = params.Protocol
	}
	if ctx.IsSet("source-network") {
		exp.SourceNetwork = params.SourceNetwork
	}
	if ctx.IsSet("source-port") {
		exp.SourcePort = params.SourcePort
	}
	if ctx.IsSet("destination-port") {
		exp.DestinationPort = params.DestinationPort
	}
	if ctx.IsSet("description") {
		exp.Description = params.Description
	}
	if ctx.IsSet("action") {
		exp.Action = params.Action
	}

	// call manipurate functions
	p, err := api.Update(params.Id, p)
	if err != nil {
		return fmt.Errorf("PacketFilterRuleUpdate is failed: %s", err)
	}

	list := []interface{}{}
	for i := range p.Expression {
		list = append(list, p.Expression[i])
	}

	return ctx.GetOutput().Print(list...)

}
