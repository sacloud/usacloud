package command

import (
	"fmt"
)

func VPCRouterStaticNatInfo(ctx Context, params *StaticNatInfoVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterStaticNatInfo is failed: %s", e)
	}

	if !p.HasStaticNAT() {
		fmt.Fprintf(GlobalOption.Err, "VPCRouter[%d] don't have any static NAT settings\n", params.Id)
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
