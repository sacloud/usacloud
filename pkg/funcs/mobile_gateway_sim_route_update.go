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

func MobileGatewaySIMRouteUpdate(ctx cli.Context, params *params.SIMRouteUpdateMobileGatewayParam) error {

	client := ctx.GetAPIClient()
	api := client.GetMobileGatewayAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("MobileGatewaySIMRouteUpdate is failed: %s", e)
	}

	routes, err := api.GetSIMRoutes(params.Id)
	if err != nil {
		return fmt.Errorf("MobileGatewaySIMRouteUpdate is failed: %s", err)
	}

	if len(routes) == 0 {
		fmt.Fprintf(ctx.IO().Err(), "MobileGateway[%d] don't have any SIM routes\n", params.Id)
		return nil
	}

	// validate
	if params.Index <= 0 || params.Index-1 >= len(routes) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	route := routes[params.Index-1]
	if ctx.IsSet("prefix") {
		route.Prefix = params.Prefix
	}
	if ctx.IsSet("sim") {
		route.ResourceID = fmt.Sprintf("%d", params.SIM)
	}

	simRoutes := &sacloud.MobileGatewaySIMRoutes{
		SIMRoutes: routes,
	}

	if _, err := api.SetSIMRoutes(params.Id, simRoutes); err != nil {
		return fmt.Errorf("MobileGatewayStaticRouteUpdate is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("MobileGatewayStaticRouteUpdate is failed: %s", err)
	}

	return nil
}
