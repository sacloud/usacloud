package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterStaticNatInfo(ctx command.Context, params *params.StaticNatInfoVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterStaticNatInfo is failed: %s", e)
	}

	if !p.HasStaticNAT() {
		fmt.Fprintf(command.GlobalOption.Err, "VPCRouter[%d] don't have any static NAT settings\n", params.Id)
		return nil
	}

	confList := p.Settings.Router.StaticNAT.Config
	// build parameters to display table
	list := []interface{}{}
	for i := range confList {
		list = append(list, &confList[i])
	}

	return ctx.GetOutput().Print(list...)
}
