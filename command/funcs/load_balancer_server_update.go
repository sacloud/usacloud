// Copyright 2017-2019 The Usacloud Authors
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

func LoadBalancerServerUpdate(ctx command.Context, params *params.ServerUpdateLoadBalancerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetLoadBalancerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("LoadBalancerServerUpdate is failed: %s", e)
	}

	// validation keys
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

	// find target server
	var server *sacloud.LoadBalancerServer
	for _, s := range vip.Servers {
		if s.IPAddress == params.Ipaddress {
			server = s
			break
		}
	}
	if server == nil {
		return fmt.Errorf("server(%s) is not found", params.Ipaddress)
	}

	if ctx.IsSet("disabled") {
		server.Enabled = "True"
		if params.Disabled {
			server.Enabled = "False"
		}
	}

	if ctx.IsSet("protocol") {

		// validate when http/https
		if params.Protocol == "http" || params.Protocol == "https" {
			if server.HealthCheck.Protocol != "http" && server.HealthCheck.Protocol != "https" {
				if !ctx.IsSet("path") {
					return fmt.Errorf("%q: is required when protocol is http/https", "path")
				}
				if !ctx.IsSet("response-code") {
					return fmt.Errorf("%q: is required when protocol is http/https", "response-code")
				}
			}
		}

		server.HealthCheck.Protocol = params.Protocol
	}

	switch server.HealthCheck.Protocol {
	case "http", "https":
		if ctx.IsSet("path") {
			server.HealthCheck.Path = params.Path
		}
		if ctx.IsSet("response-code") {
			server.HealthCheck.Status = fmt.Sprintf("%d", params.ResponseCode)
		}
	default:
		server.HealthCheck.Path = ""
		server.HealthCheck.Status = ""
	}

	p, err := client.LoadBalancer.Update(params.Id, p)
	if err != nil {
		return fmt.Errorf("LoadBalancerServerUpdate is failed: %s", err)
	}
	_, err = client.LoadBalancer.Config(params.Id)
	if err != nil {
		return fmt.Errorf("LoadBalancerServerUpdate is failed: %s", err)
	}

	return nil
}
