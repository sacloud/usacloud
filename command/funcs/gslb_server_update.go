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

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func GSLBServerUpdate(ctx command.Context, params *params.ServerUpdateGSLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetGSLBAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("GSLBServerUpdate is failed: %s", e)
	}

	if len(p.Settings.GSLB.Servers) == 0 {
		return fmt.Errorf("GSLB don't have any servers")
	}

	// validate index
	if params.Index <= 0 || params.Index-1 >= len(p.Settings.GSLB.Servers) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	// validate duplicate
	if ctx.IsSet("ipaddress") {
		for i, s := range p.Settings.GSLB.Servers {
			if i != params.Index && s.IPAddress == params.Ipaddress {
				return fmt.Errorf("GSLB already have server(%s)", params.Ipaddress)
			}
		}
	}

	server := &p.Settings.GSLB.Servers[params.Index-1]

	if ctx.IsSet("ipaddress") {
		server.IPAddress = params.Ipaddress
	}

	if ctx.IsSet("disalbed") {
		// update
		enabled := "True"
		if params.Disabled {
			enabled = "False"
		}
		server.Enabled = enabled

	}

	if ctx.IsSet("weight") {
		if params.Weight != 0 {
			server.Weight = fmt.Sprintf("%d", params.Weight)
		}
	}

	p, e = api.Update(params.Id, p)
	if e != nil {
		return fmt.Errorf("GSLBServerUpdate is failed: %s", e)
	}

	list := []interface{}{}
	for i := range p.Settings.GSLB.Servers {
		list = append(list, p.Settings.GSLB.Servers[i])
	}

	return ctx.GetOutput().Print(list...)

}
