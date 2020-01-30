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

func VPCRouterPortForwardingUpdate(ctx command.Context, params *params.PortForwardingUpdateVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterPortForwardingUpdate is failed: %s", e)
	}

	if !p.HasPortForwarding() {
		return fmt.Errorf("VPCRouter[%d] don't have any port-forwarding settings", params.Id)
	}

	// validate
	if params.Index <= 0 || params.Index-1 >= len(p.Settings.Router.PortForwarding.Config) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	pfConfig := p.Settings.Router.PortForwarding.Config[params.Index-1]
	if ctx.IsSet("protocol") {
		pfConfig.Protocol = params.Protocol
	}
	if ctx.IsSet("global-port") {
		pfConfig.GlobalPort = fmt.Sprintf("%d", params.GlobalPort)
	}
	if ctx.IsSet("private-ipaddress") {
		pfConfig.PrivateAddress = params.PrivateIpaddress
	}
	if ctx.IsSet("private-port") {
		pfConfig.PrivatePort = fmt.Sprintf("%d", params.PrivatePort)
	}
	if ctx.IsSet("description") {
		pfConfig.Description = params.Description
	}

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterPortForwardingUpdate is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterPortForwardingUpdate is failed: %s", err)
	}

	return nil
}
