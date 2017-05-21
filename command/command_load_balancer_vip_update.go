package command

import (
	"fmt"
)

func LoadBalancerVipUpdate(ctx Context, params *VipUpdateLoadBalancerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetLoadBalancerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("LoadBalancerVipUpdate is failed: %s", e)
	}

	// validate

	// index
	if params.Index <= 0 || params.Index-1 >= len(p.Settings.LoadBalancer) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	for i, v := range p.Settings.LoadBalancer {
		if i == params.Index-1 {
			continue
		}

		if v.VirtualIPAddress == params.Vip {
			return fmt.Errorf("VIP(%q) is already used", params.Vip)
		}
	}

	vip := p.Settings.LoadBalancer[params.Index-1]

	if ctx.IsSet("vip") {
		vip.VirtualIPAddress = params.Vip
	}
	if ctx.IsSet("port") {
		vip.Port = fmt.Sprintf("%d", params.Port)
	}
	if ctx.IsSet("delay-loop") {
		vip.DelayLoop = fmt.Sprintf("%d", params.DelayLoop)
	}
	if ctx.IsSet("sorry-server") {
		vip.SorryServer = params.SorryServer
	}

	p, err := client.LoadBalancer.Update(params.Id, p)
	if err != nil {
		return fmt.Errorf("LoadBalancerVipUpdate is failed: %s", err)
	}
	_, err = client.LoadBalancer.Config(params.Id)
	if err != nil {
		return fmt.Errorf("LoadBalancerVipUpdate is failed: %s", err)
	}

	return nil
}
