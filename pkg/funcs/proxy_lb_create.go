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

func ProxyLBCreate(ctx cli.Context, params *params.CreateProxyLBParam) error {

	// Validate params
	if ctx.IsSet("sorry-server-ipaddress") || ctx.IsSet("sorry-server-port") {
		if params.SorryServerIpaddress == "" || params.SorryServerPort == 0 {
			return fmt.Errorf("both of sorry-server-ipaddress and sorry-server-port are required")
		}
	}

	client := ctx.GetAPIClient()
	api := client.GetProxyLBAPI()
	p := api.New(params.Name)

	// set params
	p.SetPlan(sacloud.ProxyLBPlan(params.Plan))
	p.SetDescription(params.Description)
	p.SetTags(params.Tags)
	p.SetIconByID(params.IconId)

	protocol := params.Protocol
	switch protocol {
	case "http":
		p.SetHTTPHealthCheck(params.HostHeader, params.Path, params.DelayLoop)
	case "tcp":
		p.SetTCPHealthCheck(params.DelayLoop)
	default:
		return fmt.Errorf("invalid protocol: %s", protocol)
	}

	if params.StickySession {
		p.Settings.ProxyLB.StickySession = sacloud.ProxyLBSessionSetting{
			Enabled: true,
			Method:  sacloud.ProxyLBStickySessionDefaultMethod,
		}
	}

	p.SetSorryServer(params.SorryServerIpaddress, params.SorryServerPort)

	p.Settings.ProxyLB.Timeout = &sacloud.ProxyLBTimeout{
		InactiveSec: params.Timeout,
	}

	// call Create(id)
	res, err := api.Create(p)
	if err != nil {
		return fmt.Errorf("ProxyLBCreate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}
