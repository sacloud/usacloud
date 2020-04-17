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

func LoadBalancerServerAdd(ctx cli.Context, params *params.ServerAddLoadBalancerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetLoadBalancerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("LoadBalancerServerAdd is failed: %s", e)
	}

	// validation keys
	if (!ctx.IsSet("vip") || !ctx.IsSet("port")) && !ctx.IsSet("vip-index") {
		return fmt.Errorf("%q or %q and %q: is required", "vip-index", "vip", "port")
	}

	// validate health check
	if params.Protocol == "http" || params.Protocol == "https" {
		if params.Path == "" || params.ResponseCode <= 0 {
			return fmt.Errorf("%q and %q: is required when protocol is http/https", "path", "response-code")
		}
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

	// validate IP duplicate
	for _, s := range vip.Servers {
		if s.IPAddress == params.Ipaddress {
			return fmt.Errorf("IPAddress(%q) is already used", params.Ipaddress)
		}
	}

	// validate server count(max = 40)
	if len(vip.Servers) >= 40 {
		return fmt.Errorf("VIP(%s:%s) already has maximum count of servers", vip.VirtualIPAddress, vip.Port)
	}

	var server = &sacloud.LoadBalancerServer{}
	server.IPAddress = params.Ipaddress
	server.Port = vip.Port
	server.Enabled = "True"
	if params.Disabled {
		server.Enabled = "False"
	}
	server.HealthCheck = &sacloud.LoadBalancerHealthCheck{}

	server.HealthCheck.Protocol = params.Protocol

	switch server.HealthCheck.Protocol {
	case "http", "https":
		server.HealthCheck.Path = params.Path
		server.HealthCheck.Status = fmt.Sprintf("%d", params.ResponseCode)
	}
	vip.AddServer(server)

	p, err := client.LoadBalancer.Update(params.Id, p)
	if err != nil {
		return fmt.Errorf("LoadBalancerServerAdd is failed: %s", err)
	}
	_, err = client.LoadBalancer.Config(params.Id)
	if err != nil {
		return fmt.Errorf("LoadBalancerServerAdd is failed: %s", err)
	}

	return nil
}
