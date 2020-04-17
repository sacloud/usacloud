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

func MobileGatewaySIMRouteAdd(ctx cli.Context, params *params.SIMRouteAddMobileGatewayParam) error {

	client := ctx.GetAPIClient()
	api := client.GetMobileGatewayAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("MobileGatewaySIMRouteAdd is failed: %s", e)
	}

	routeList, err := api.GetSIMRoutes(params.Id)
	if err != nil {
		return fmt.Errorf("MobileGatewaySIMRouteAdd is failed: %s", err)
	}

	routes := &sacloud.MobileGatewaySIMRoutes{
		SIMRoutes: routeList,
	}

	if _, exists := routes.FindSIMRoute(params.SIM, params.Prefix); exists != nil {
		fmt.Fprintf(ctx.IO().Out(), "SIM Route[%s -> %d] already exists", params.Prefix, params.SIM)
		return nil
	}

	if _, err := client.GetSIMAPI().Read(params.SIM); err != nil {
		return fmt.Errorf("SIM[%d] is not found: %s", params.SIM, err)
	}

	if _, err := api.AddSIMRoute(params.Id, params.SIM, params.Prefix); err != nil {
		return fmt.Errorf("MobileGatewaySIMRouteAdd is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("MobileGatewaySIMRouteAdd is failed: %s", err)
	}
	return nil
}
