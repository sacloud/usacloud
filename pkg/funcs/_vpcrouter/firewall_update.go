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

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func FirewallUpdate(ctx cli.Context, params *params.FirewallUpdateVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterFirewallUpdate is failed: %s", e)
	}

	if !p.HasFirewall() {
		return fmt.Errorf("VPCRouter[%d] don't have any firewall rules[%s]", params.Id, params.Direction)
	}

	// validate
	if params.Protocol == "icmp" || params.Protocol == "ip" {
		if params.SourcePort > 0 {
			return fmt.Errorf("%q: can't set when protocol is [icmp/ip]", "source-port")
		}
		if params.DestinationPort > 0 {
			return fmt.Errorf("%q: can't set when protocol is [icmp/ip]", "destination-port")
		}
	}

	// validate
	cnt := len(p.Settings.Router.Firewall.Config[params.Interface].Receive)
	if params.Direction == "send" {
		cnt = len(p.Settings.Router.Firewall.Config[params.Interface].Send)
	}

	if params.Index <= 0 || params.Index-1 >= cnt {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	var rule *sacloud.VPCRouterFirewallRule
	switch params.Direction {
	case "send":
		rule = p.Settings.Router.Firewall.Config[params.Interface].Send[params.Index-1]
	case "receive":
		rule = p.Settings.Router.Firewall.Config[params.Interface].Receive[params.Index-1]
	}

	if ctx.IsSet("protocol") {
		rule.Protocol = params.Protocol
	}
	if rule.Protocol == "icmp" || rule.Protocol == "ip" {
		rule.DestinationPort = ""
		rule.SourcePort = ""
	} else {
		if ctx.IsSet("source-port") {
			rule.SourcePort = fmt.Sprintf("%d", params.SourcePort)
		}
		if ctx.IsSet("destination-port") {
			rule.DestinationPort = fmt.Sprintf("%d", params.DestinationPort)
		}
	}

	if ctx.IsSet("action") {
		rule.Action = params.Action
	}
	if ctx.IsSet("source-network") {
		rule.SourceNetwork = params.SourceNetwork
	}
	if ctx.IsSet("destination-network") {
		rule.DestinationNetwork = params.DestinationNetwork
	}

	if ctx.IsSet("enable-logging") {
		v := "False"
		if params.EnableLogging {
			v = "True"
		}
		rule.Logging = v
	}

	if ctx.IsSet("description") {
		rule.Description = params.Description
	}

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterFirewallUpdate is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterFirewallUpdate is failed: %s", err)
	}

	return nil

}
