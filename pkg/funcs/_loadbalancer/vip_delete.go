// Copyright 2017-2020 The Usacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package loadbalancer

import (
	"fmt"

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func VipDelete(ctx cli.Context, params *params.VipDeleteLoadBalancerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetLoadBalancerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("LoadBalancerVipDelete is failed: %s", e)
	}

	initLoadBalancerSettings(p)

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