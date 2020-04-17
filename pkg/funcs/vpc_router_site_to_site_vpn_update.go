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

func VPCRouterSiteToSiteVPNUpdate(ctx cli.Context, params *params.SiteToSiteVPNUpdateVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterSiteToSiteVPNUpdate is failed: %s", e)
	}

	if !p.HasSiteToSiteIPsecVPN() {
		return fmt.Errorf("VPCRouter[%d] don't have any site-to-site IPSec VPN settings", params.Id)
	}

	// validate
	if params.Index <= 0 || params.Index-1 >= len(p.Settings.Router.SiteToSiteIPsecVPN.Config) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	cnf := p.Settings.Router.SiteToSiteIPsecVPN.Config[params.Index-1]
	if ctx.IsSet("peer") {
		cnf.Peer = params.Peer
	}
	if ctx.IsSet("remote-id") {
		cnf.RemoteID = params.RemoteId
	}
	if ctx.IsSet("pre-shared-secret") {
		cnf.PreSharedSecret = params.PreSharedSecret
	}
	if ctx.IsSet("routes") {
		cnf.Routes = params.Routes
	}
	if ctx.IsSet("local-prefix") {
		cnf.LocalPrefix = params.LocalPrefix
	}

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterSiteToSiteVPNUpdate is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterSiteToSiteVPNUpdate is failed: %s", err)
	}

	return nil

}
