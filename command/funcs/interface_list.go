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
	"github.com/sacloud/usacloud/pkg/utils"
)

func InterfaceList(ctx command.Context, params *params.ListInterfaceParam) error {

	client := ctx.GetAPIClient()
	finder := client.GetInterfaceAPI()

	finder.SetEmpty()

	if !utils.IsEmpty(params.Name) {
		for _, v := range params.Name {
			finder.SetFilterBy("Name", v)
		}
	}
	if !utils.IsEmpty(params.Id) {
		for _, v := range params.Id {
			finder.SetFilterMultiBy("ID", v)
		}
	}
	if !utils.IsEmpty(params.From) {
		finder.SetOffset(params.From)
	}
	if !utils.IsEmpty(params.Max) {
		finder.SetLimit(params.Max)
	}
	if !utils.IsEmpty(params.Sort) {
		for _, v := range params.Sort {
			setSortBy(finder, v)
		}
	}

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return fmt.Errorf("InterfaceList is failed: %s", err)
	}

	list := []interface{}{}
	ignoreTags := []string{"@appliance-database", "@appliance-loadbalancer", "@appliance-vpcrouter"}

Outer:
	for i, nic := range res.Interfaces {
		// customize: ignore appliance interface
		for _, t := range ignoreTags {
			if nic.Server.HasTag(t) {
				continue Outer
			}
		}

		list = append(list, &res.Interfaces[i])
	}

	return ctx.GetOutput().Print(list...)

}
