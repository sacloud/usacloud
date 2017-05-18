package command

import (
	"fmt"
)

func LoadBalancerVipDelete(ctx Context, params *VipDeleteLoadBalancerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetLoadBalancerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("LoadBalancerVipDelete is failed: %s", e)
	}

	// index
	if params.Index <= 0 || params.Index-1 >= len(p.Settings.LoadBalancer) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	// delete VIP
	vip := p.Settings.LoadBalancer[params.Index-1]
	p.DeleteLoadBalancerSetting(vip.VirtualIPAddress, vip.Port)

	p, err := client.LoadBalancer.Update(params.Id, p)
	if err != nil {
		return fmt.Errorf("LoadBalancerVipDelete is failed: %s", err)
	}
	_, err = client.LoadBalancer.Config(params.Id)
	if err != nil {
		return fmt.Errorf("LoadBalancerVipDelete is failed: %s", err)
	}

	return nil

}
