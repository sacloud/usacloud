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

func VPCRouterFirewallDelete(ctx command.Context, params *params.FirewallDeleteVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterFirewallDelete is failed: %s", e)
	}

	if !p.HasFirewall() {
		return fmt.Errorf("VPCRouter[%d] don't have any firewall rules[%s]", params.Id, params.Direction)
	}

	// validate
	cnt := len(p.Settings.Router.Firewall.Config[params.Interface].Receive)
	if params.Direction == "send" {
		cnt = len(p.Settings.Router.Firewall.Config[params.Interface].Send)
	}

	if params.Index <= 0 || params.Index-1 >= cnt {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	switch params.Direction {
	case "send":
		p.Settings.Router.RemoveFirewallRuleSendAt(params.Interface, params.Index-1)
	case "receive":
		p.Settings.Router.RemoveFirewallRuleReceiveAt(params.Interface, params.Index-1)
	}

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterFirewallDelete is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterFirewallDelete is failed: %s", err)
	}

	return nil

}
