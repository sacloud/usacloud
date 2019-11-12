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
	"strings"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterDhcpServerInfo(ctx command.Context, params *params.DhcpServerInfoVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterDhcpServerInfo is failed: %s", e)
	}

	if !p.HasDHCPServer() {
		fmt.Fprintf(command.GlobalOption.Err, "VPCRouter[%d] don't have any DHCP servers\n", params.Id)
		return nil
	}

	confList := p.Settings.Router.DHCPServer.Config
	// build parameters to display table
	list := []interface{}{}
	for i := range confList {
		list = append(list, &struct {
			*sacloud.VPCRouterDHCPServerConfig
			DNSServerList string
		}{
			VPCRouterDHCPServerConfig: confList[i],
			DNSServerList:             strings.Join(confList[i].DNSServers, ","),
		})
	}

	return ctx.GetOutput().Print(list...)

}
