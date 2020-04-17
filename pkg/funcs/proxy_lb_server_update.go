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

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func ProxyLBServerUpdate(ctx cli.Context, params *params.ServerUpdateProxyLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetProxyLBAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ProxyLBServerUpdate is failed: %s", e)
	}

	if len(p.Settings.ProxyLB.Servers) == 0 {
		return fmt.Errorf("ProxyLB don't have any servers")
	}

	// validate index
	if params.Index <= 0 || params.Index-1 >= len(p.Settings.ProxyLB.Servers) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	server := &p.Settings.ProxyLB.Servers[params.Index-1]

	if ctx.IsSet("ipaddress") {
		server.IPAddress = params.Ipaddress
	}
	if ctx.IsSet("port") {
		server.Port = params.Port
	}
	if ctx.IsSet("disabled") {
		server.Enabled = !params.Disabled
	}

	p, e = api.UpdateSetting(params.Id, p)
	if e != nil {
		return fmt.Errorf("ProxyLBServerUpdate is failed: %s", e)
	}

	var list []interface{}
	for i := range p.Settings.ProxyLB.Servers {
		list = append(list, &p.Settings.ProxyLB.Servers[i])
	}
	return ctx.GetOutput().Print(list...)

}
