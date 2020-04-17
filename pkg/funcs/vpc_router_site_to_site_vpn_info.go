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
	"strings"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func VPCRouterSiteToSiteVPNInfo(ctx cli.Context, params *params.SiteToSiteVPNInfoVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterSiteToSiteVPNInfo is failed: %s", e)
	}

	if !p.HasSiteToSiteIPsecVPN() {
		fmt.Fprintf(ctx.IO().Err(), "VPCRouter[%d] don't have any site-to-site IPSec VPN settings\n", params.Id)
		return nil
	}

	confList := p.Settings.Router.SiteToSiteIPsecVPN.Config

	type s2sSetting struct {
		*sacloud.VPCRouterSiteToSiteIPsecVPNConfig
		RoutesJoined      string
		LocalPrefixJoined string
	}

	// build parameters to display table
	list := []interface{}{}
	for i := range confList {
		s := &s2sSetting{
			VPCRouterSiteToSiteIPsecVPNConfig: confList[i],
			RoutesJoined:                      strings.Join(confList[i].Routes, ","),
			LocalPrefixJoined:                 strings.Join(confList[i].LocalPrefix, ","),
		}
		list = append(list, s)
	}

	return ctx.GetOutput().Print(list...)

}
