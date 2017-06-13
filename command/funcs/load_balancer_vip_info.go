package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func LoadBalancerVipInfo(ctx command.Context, params *params.VipInfoLoadBalancerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetLoadBalancerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("LoadBalancerVipInfo is failed: %s", e)
	}

	vips := p.Settings.LoadBalancer
	if len(vips) == 0 {
		fmt.Fprintf(command.GlobalOption.Err, "LoadBalancer don't have any VIPs\n")
		return nil
	}

	// build parameters to display table
	list := []interface{}{}
	for i := range vips {
		list = append(list, &vips[i])
	}

	return ctx.GetOutput().Print(list...)

}
