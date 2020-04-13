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

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterL2TPServerUpdate(ctx command.Context, params *params.L2TPServerUpdateVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterL2TPServerUpdate is failed: %s", e)
	}

	if !params.Disabled {
		// validate
		targets := []string{"range-start", "range-stop", "pre-shared-secret"}
		for _, t := range targets {
			if !ctx.IsSet(t) {
				return fmt.Errorf("%q: is required when enabled is true", t)
			}
		}

		if !p.HasL2TPIPsecServer() {
			p.InitVPCRouterSetting()
		}
		p.Settings.Router.EnableL2TPIPsecServer(params.PreSharedSecret, params.RangeStart, params.RangeStop)
	} else {
		// validate
		targets := []string{"range-start", "range-stop", "pre-shared-secret"}
		for _, t := range targets {
			if ctx.IsSet(t) {
				return fmt.Errorf("%q: can't set when enabled is false", t)
			}
		}
		if !p.HasL2TPIPsecServer() {
			return nil
		}

		p.Settings.Router.DisableL2TPIPsecServer()
	}

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterL2TPServerUpdate is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterL2TPServerUpdate is failed: %s", err)
	}

	return nil
}
