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

func ProxyLBUpdate(ctx command.Context, params *params.UpdateProxyLBParam) error {

	// Validate params
	if ctx.IsSet("sorry-server-ipaddress") || ctx.IsSet("sorry-server-port") {
		if params.SorryServerIpaddress == "" || params.SorryServerPort == 0 {
			return fmt.Errorf("both of sorry-server-ipaddress and sorry-server-port are required")
		}
	}

	client := ctx.GetAPIClient()
	api := client.GetProxyLBAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ProxyLBUpdate is failed: %s", e)
	}

	// set params

	if ctx.IsSet("name") {
		p.SetName(params.Name)
	}
	if ctx.IsSet("description") {
		p.SetDescription(params.Description)
	}
	if ctx.IsSet("tags") {
		p.SetTags(params.Tags)
	}
	if ctx.IsSet("icon-id") {
		p.SetIconByID(params.IconId)
	}

	protocol := p.Settings.ProxyLB.HealthCheck.Protocol
	if ctx.IsSet("protocol") {
		protocol = params.Protocol
	}

	delayLoop := p.Settings.ProxyLB.HealthCheck.DelayLoop
	switch protocol {
	case "http":
		hostHeader := p.Settings.ProxyLB.HealthCheck.Host
		if ctx.IsSet("host-header") {
			hostHeader = params.HostHeader
		}

		path := p.Settings.ProxyLB.HealthCheck.Path
		if ctx.IsSet("path") {
			path = params.Path
		}

		if ctx.IsSet("delay-loop") {
			delayLoop = params.DelayLoop
		}

		p.SetHTTPHealthCheck(hostHeader, path, delayLoop)
	case "tcp":

		if ctx.IsSet("delay-loop") {
			delayLoop = params.DelayLoop
		}

		p.SetTCPHealthCheck(params.DelayLoop)
	default:
		return fmt.Errorf("invalid protocol: %s", protocol)
	}

	if ctx.IsSet("sticky-session") {
		if params.StickySession {
			p.Settings.ProxyLB.StickySession = sacloud.ProxyLBSessionSetting{
				Enabled: true,
				Method:  sacloud.ProxyLBStickySessionDefaultMethod,
			}
		} else {
			p.Settings.ProxyLB.StickySession = sacloud.ProxyLBSessionSetting{
				Enabled: false,
			}
		}
	}

	if ctx.IsSet("sorry-server-ipaddress") || ctx.IsSet("sorry-server-port") {
		p.SetSorryServer(params.SorryServerIpaddress, params.SorryServerPort)
	}

	if ctx.IsSet("timeout") {
		p.Settings.ProxyLB.Timeout = &sacloud.ProxyLBTimeout{
			InactiveSec: params.Timeout,
		}
	}

	// call Update(id)
	res, err := api.Update(params.Id, p)
	if err != nil {
		return fmt.Errorf("ProxyLBUpdate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}
