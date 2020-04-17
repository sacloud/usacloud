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

package funcs

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func LoadBalancerVipAdd(ctx cli.Context, params *params.VipAddLoadBalancerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetLoadBalancerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("LoadBalancerVipAdd is failed: %s", e)
	}

	initLoadBalancerSettings(p)

	// validate
	if len(p.Settings.LoadBalancer) >= 4 {
		return fmt.Errorf("LoadBalancer already has maximum count of VIP")
	}

	for _, v := range p.Settings.LoadBalancer {
		if v.VirtualIPAddress == params.Vip && v.Port == fmt.Sprintf("%d", params.Port) {
			return fmt.Errorf("VIP(%s:%d) is already used", params.Vip, params.Port)
		}
	}

	// set params
	var vip = &sacloud.LoadBalancerSetting{
		VirtualIPAddress: params.Vip,
		Port:             fmt.Sprintf("%d", params.Port),
		DelayLoop:        fmt.Sprintf("%d", params.DelayLoop),
		SorryServer:      params.SorryServer,
		Description:      params.Description,
	}

	p.AddLoadBalancerSetting(vip)
	p, err := client.LoadBalancer.Update(params.Id, p)
	if err != nil {
		return fmt.Errorf("LoadBalancerVipAdd is failed: %s", err)
	}
	_, err = client.LoadBalancer.Config(params.Id)
	if err != nil {
		return fmt.Errorf("LoadBalancerVipAdd is failed: %s", err)
	}

	return nil
}
