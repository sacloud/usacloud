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

func PacketFilterRuleAdd(ctx command.Context, params *params.RuleAddPacketFilterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetPacketFilterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("PacketFilterRuleAdd is failed: %s", e)
	}

	// validate
	switch params.Protocol {
	case "icmp", "fragment", "ip":
		if params.SourcePort != "" || params.DestinationPort != "" {
			return fmt.Errorf("When protocol is %s , can't set Source/Destination port", params.Protocol)
		}
	}

	// index
	if params.Index <= 0 || params.Index-1 >= len(p.Expression) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	action := true
	switch params.Action {
	case "allow":
		action = true
	case "deny":
		action = false
	}

	index := params.Index - 1

	switch params.Protocol {
	case "tcp":
		p.AddTCPRuleAt(params.SourceNetwork, params.SourcePort, params.DestinationPort, params.Description, action, index)
	case "udp":
		p.AddUDPRuleAt(params.SourceNetwork, params.SourcePort, params.DestinationPort, params.Description, action, index)
	case "icmp":
		p.AddICMPRuleAt(params.SourceNetwork, params.Description, action, index)
	case "fragment":
		p.AddFragmentRuleAt(params.SourceNetwork, params.Description, action, index)
	case "ip":
		p.AddIPRuleAt(params.SourceNetwork, params.Description, action, index)
	default:
		panic(fmt.Errorf("Unknown protocol %s", params.Protocol))
	}

	// call manipurate functions
	p, err := api.Update(params.Id, p)
	if err != nil {
		return fmt.Errorf("PacketFilterRuleAdd is failed: %s", err)
	}

	list := []interface{}{}
	for i := range p.Expression {
		list = append(list, p.Expression[i])
	}

	return ctx.GetOutput().Print(list...)

}
