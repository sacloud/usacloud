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

func PacketFilterRuleDelete(ctx command.Context, params *params.RuleDeletePacketFilterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetPacketFilterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("PacketFilterRuleDelete is failed: %s", e)
	}

	if len(p.Expression) == 0 {
		fmt.Fprintf(command.GlobalOption.Err, "PacketFilter don't have any rules\n")
		return nil
	}

	// index
	if params.Index <= 0 || params.Index-1 >= len(p.Expression) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	p.RemoveRuleAt(params.Index - 1)

	// call manipurate functions
	p, err := api.Update(params.Id, p)
	if err != nil {
		return fmt.Errorf("PacketFilterRuleDelete is failed: %s", err)
	}

	list := []interface{}{}
	for i := range p.Expression {
		list = append(list, p.Expression[i])
	}

	return ctx.GetOutput().Print(list...)

}
