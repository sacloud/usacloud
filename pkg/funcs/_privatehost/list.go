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

package privatehost

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
	"github.com/sacloud/usacloud/pkg/util"
)

func List(ctx cli.Context, params *params.ListPrivateHostParam) error {

	client := ctx.GetAPIClient()
	finder := client.GetPrivateHostAPI()

	finder.SetEmpty()

	if !util.IsEmpty(params.Name) {
		for _, v := range params.Name {
			finder.SetFilterBy("Name", v)
		}
	}
	if !util.IsEmpty(params.Id) {
		for _, v := range params.Id {
			finder.SetFilterMultiBy("ID", v)
		}
	}
	if !util.IsEmpty(params.From) {
		finder.SetOffset(params.From)
	}
	if !util.IsEmpty(params.Max) {
		finder.SetLimit(params.Max)
	}
	if !util.IsEmpty(params.Sort) {
		for _, v := range params.Sort {
			setSortBy(finder, v)
		}
	}

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return fmt.Errorf("PrivateHostList is failed: %s", err)
	}

	list := []interface{}{}

	for i, h := range res.PrivateHosts {

		if !params.GetCommandDef().Params["tags"].FilterFunc(list, &res.PrivateHosts[i], params.Tags) {
			continue
		}

		list = append(list, &struct {
			*sacloud.PrivateHost
			AssignedCore   int
			TotalCore      int
			AssignedMemory int
			TotalMemory    int
		}{
			PrivateHost:    &res.PrivateHosts[i],
			AssignedCore:   h.AssignedCPU,
			TotalCore:      h.Plan.CPU,
			AssignedMemory: h.GetAssignedMemoryGB(),
			TotalMemory:    h.Plan.GetMemoryGB(),
		})
	}
	return ctx.Output().Print(list...)

}