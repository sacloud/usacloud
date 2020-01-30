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

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ProxyLBServerAdd(ctx command.Context, params *params.ServerAddProxyLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetProxyLBAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ProxyLBServerAdd is failed: %s", e)
	}

	// check duplicate
	for _, server := range p.Settings.ProxyLB.Servers {
		if server.IPAddress == params.Ipaddress && server.Port == params.Port {
			return fmt.Errorf("server %s/%d is already exists", server.IPAddress, server.Port)
		}
	}

	p.AddServer(params.Ipaddress, params.Port, !params.Disabled)

	p, e = api.UpdateSetting(params.Id, p)
	if e != nil {
		return fmt.Errorf("ProxyLBServerAdd is failed: %s", e)
	}

	var list []interface{}
	for i := range p.Settings.ProxyLB.Servers {
		list = append(list, &p.Settings.ProxyLB.Servers[i])
	}
	return ctx.GetOutput().Print(list...)
}
