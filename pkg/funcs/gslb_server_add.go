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

func GSLBServerAdd(ctx cli.Context, params *params.ServerAddGSLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetGSLBAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("GSLBServerAdd is failed: %s", e)
	}

	// validate maxlen
	if len(p.Settings.GSLB.Servers) == 6 {
		return fmt.Errorf("GSLB already have max(6) servers")
	}

	// validate duplicate
	for _, s := range p.Settings.GSLB.Servers {
		if s.IPAddress == params.Ipaddress {
			return fmt.Errorf("GSLB already have server(%s)", params.Ipaddress)
		}
	}

	// add
	enabled := "True"
	if params.Disabled {
		enabled = "False"
	}

	server := &sacloud.GSLBServer{
		IPAddress: params.Ipaddress,
		Enabled:   enabled,
	}

	if params.Weight != 0 {
		server.Weight = fmt.Sprintf("%d", params.Weight)
	}
	p.AddGSLBServer(server)

	p, e = api.Update(params.Id, p)
	if e != nil {
		return fmt.Errorf("GSLBServerAdd is failed: %s", e)
	}

	list := []interface{}{}
	for i := range p.Settings.GSLB.Servers {
		list = append(list, p.Settings.GSLB.Servers[i])
	}

	return ctx.GetOutput().Print(list...)

}
