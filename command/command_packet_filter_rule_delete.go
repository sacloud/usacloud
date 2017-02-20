package command

import (
	"fmt"
)

func PacketFilterRuleDelete(ctx Context, params *RuleDeletePacketFilterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetPacketFilterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("PacketFilterRuleDelete is failed: %s", e)
	}

	if len(p.Expression) == 0 {
		fmt.Fprintf(GlobalOption.Err, "PacketFilter1 don't have any rules\n")
		return nil
	}

	// index
	if params.Index <= 0 || params.Index-1 >= len(p.Expression) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	p.RemoveRuleAt(params.Index - 1)

	// call manipurate functions
	p, err := api.Update(params.Id, p)
	if err != nil {
		return fmt.Errorf("PacketFilterRuleDelete is failed: %s", err)
	}

	list := []interface{}{}
	for i := range p.Expression {
		list = append(list, p.Expression[i])
	}

	return ctx.GetOutput().Print(list...)

}
