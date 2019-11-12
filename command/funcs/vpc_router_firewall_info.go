// Copyright 2017-2019 The Usacloud Authors
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
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterFirewallInfo(ctx command.Context, params *params.FirewallInfoVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterFirewallInfo is failed: %s", e)
	}

	if !p.HasFirewall() {
		fmt.Fprintf(command.GlobalOption.Err, "VPCRouter[%d] don't have any firewall rules\n", params.Id)
		return nil
	}

	ruleList := p.Settings.Router.Firewall.Config[params.Interface].Receive
	switch params.Direction {
	case "send":
		ruleList = p.Settings.Router.Firewall.Config[params.Interface].Send
	case "receive":
		ruleList = p.Settings.Router.Firewall.Config[params.Interface].Receive
	}

	if len(ruleList) == 0 {
		fmt.Fprintf(command.GlobalOption.Err, "VPCRouter don't have any firewall(%s) rules\n", params.Direction)
		return nil
	}

	// build parameters to display table
	list := []interface{}{}
	for i := range ruleList {
		list = append(list, &struct {
			*sacloud.VPCRouterFirewallRule
			Interface int
		}{
			VPCRouterFirewallRule: ruleList[i],
			Interface:             params.Interface,
		})
	}

	return ctx.GetOutput().Print(list...)

}
