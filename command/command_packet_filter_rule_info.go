package command

import (
	"fmt"
)

func PacketFilterRuleInfo(ctx Context, params *RuleInfoPacketFilterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetPacketFilterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("PacketFilterRuleList is failed: %s", e)
	}

	if len(p.Expression) == 0 {
		fmt.Fprintf(GlobalOption.Err, "PacketFilter don't have any rules\n")
		return nil
	}

	list := []interface{}{}
	for i := range p.Expression {
		list = append(list, p.Expression[i])
	}

	return ctx.GetOutput().Print(list...)
}
