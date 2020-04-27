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

func SiteToSiteVPNPeers(ctx cli.Context, params *params.SiteToSiteVPNPeersVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterSiteToSiteVPNPeers is failed: %s", e)
	}

	if !p.HasSiteToSiteIPsecVPN() {
		fmt.Fprintf(ctx.IO().Err(), "VPCRouter[%d] don't have any site-to-site IPSec VPN settings\n", params.Id)
		return nil
	}

	status, err := api.Status(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterSiteToSiteVPNPeers is failed: %s", err)
	}

	// build parameters to display table
	list := []interface{}{}
	for i := range status.SiteToSiteIPsecVPNPeers {
		list = append(list, &status.SiteToSiteIPsecVPNPeers[i])
	}

	return ctx.Output().Print(list...)
}
