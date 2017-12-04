package funcs

import "github.com/sacloud/libsacloud/sacloud"

func initLoadBalancerSettings(lb *sacloud.LoadBalancer) {
	if lb.Settings == nil {
		lb.Settings = &sacloud.LoadBalancerSettings{}
	}
	if lb.Settings.LoadBalancer == nil {
		lb.Settings.LoadBalancer = []*sacloud.LoadBalancerSetting{}
	}
}
