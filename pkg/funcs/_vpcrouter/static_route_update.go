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

func StaticRouteUpdate(ctx cli.Context, params *params.StaticRouteUpdateVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterStaticRouteUpdate is failed: %s", e)
	}

	if !p.HasStaticRoutes() {
		return fmt.Errorf("VPCRouter[%d] don't have any static-routes", params.Id)
	}

	// validate
	if params.Index <= 0 || params.Index-1 >= len(p.Settings.Router.StaticRoutes.Config) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	cnf := p.Settings.Router.StaticRoutes.Config[params.Index-1]
	if ctx.IsSet("prefix") {
		cnf.Prefix = params.Prefix
	}
	if ctx.IsSet("next-hop") {
		cnf.NextHop = params.NextHop
	}

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterStaticRouteUpdate is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterStaticRouteUpdate is failed: %s", err)
	}

	return nil

}