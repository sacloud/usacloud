package command

import (
	"fmt"
)

func LoadBalancerVipInfo(ctx Context, params *VipInfoLoadBalancerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetLoadBalancerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("LoadBalancerVipInfo is failed: %s", e)
	}

	vips := p.Settings.LoadBalancer
	if len(vips) == 0 {
		fmt.Fprintf(GlobalOption.Err, "LoadBalancer don't have any VIPs\n")
		return nil
	}

	// build parameters to display table
	list := []interface{}{}
	for i := range vips {
		list = append(list, &vips[i])
	}

	return ctx.GetOutput().Print(list...)

}
