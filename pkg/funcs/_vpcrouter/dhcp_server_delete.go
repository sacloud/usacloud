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

package vpcrouter

import (
	"fmt"

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func DhcpServerDelete(ctx cli.Context, params *params.DhcpServerDeleteVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterDhcpServerDelete is failed: %s", e)
	}

	if !p.HasDHCPServer() {
		return fmt.Errorf("VPCRouter[%d] don't have any DHCP servers", params.Id)
	}

	_, cnf := p.Settings.Router.FindDHCPServerAt(params.Interface)
	if cnf == nil {
		return fmt.Errorf("DHCP server is not found on eth%d", params.Interface)
	}

	p.Settings.Router.RemoveDHCPServerAt(params.Interface)

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterDhcpServerDelete is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterDhcpServerDelete is failed: %s", err)
	}

	return nil
}
