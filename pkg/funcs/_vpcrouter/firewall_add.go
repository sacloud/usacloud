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

func FirewallAdd(ctx cli.Context, params *params.FirewallAddVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterFirewallAdd is failed: %s", e)
	}

	if !p.HasSetting() {
		p.InitVPCRouterSetting()
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

	// build parameters
	isAllow := false
	if params.Action == "allow" {
		isAllow = true
	}
	sourcePort := ""
	destPort := ""
	if params.Protocol == "tcp" || params.Protocol == "udp" {
		if params.SourcePort > 0 {
			sourcePort = fmt.Sprintf("%d", params.SourcePort)
		}
		if params.DestinationPort > 0 {
			destPort = fmt.Sprintf("%d", params.DestinationPort)
		}
	}

	var f func(ifIndex int, isAllow bool, protocol string, sourceNetwork string, sourcePort string,
		destNetwork string, destPort string, logging bool, description string) (int, *sacloud.VPCRouterFirewallRule)
	switch params.Direction {
	case "send":
		f = p.Settings.Router.AddFirewallRuleSend
	case "receive":
		f = p.Settings.Router.AddFirewallRuleReceive
	}
	// add
	f(
		params.Interface,
		isAllow,
		params.Protocol,
		params.SourceNetwork,
		sourcePort,
		params.DestinationNetwork,
		destPort,
		params.EnableLogging,
		params.Description,
	)

	_, err := api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterFirewallAdd is failed: %s", err)
	}
	_, err = api.Config(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterFirewallAdd is failed: %s", err)
	}
	return nil

}
