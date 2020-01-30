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
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func LoadBalancerServerInfo(ctx command.Context, params *params.ServerInfoLoadBalancerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetLoadBalancerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("LoadBalancerServerInfo is failed: %s", e)
	}

	// validation
	if (!ctx.IsSet("vip") || !ctx.IsSet("port")) && !ctx.IsSet("vip-index") {
		return fmt.Errorf("%q or %q and %q: is required", "vip-index", "vip", "port")
	}

	var vip *sacloud.LoadBalancerSetting
	if params.VipIndex > 0 {
		// use VIP index
		if params.VipIndex-1 >= len(p.Settings.LoadBalancer) {
			return fmt.Errorf("vip-index(%d) is out of range", params.VipIndex)
		}
		vip = p.Settings.LoadBalancer[params.VipIndex-1]
	} else {
		// use VIP and port
		for _, v := range p.Settings.LoadBalancer {
			if v.VirtualIPAddress == params.Vip && v.Port == fmt.Sprintf("%d", params.Port) {
				vip = v
				break
			}
		}
		if vip == nil {
			return fmt.Errorf("VIP(%s:%d) is not found", params.Vip, params.Port)
		}
	}

	// build parameters to display table
	list := []interface{}{}
	for i := range vip.Servers {
		list = append(list, &vip.Servers[i])
	}

	return ctx.GetOutput().Print(list...)

}
