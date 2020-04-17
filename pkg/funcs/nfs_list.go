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

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
	"github.com/sacloud/usacloud/pkg/utils"
)

func NFSList(ctx cli.Context, params *params.ListNFSParam) error {

	client := ctx.GetAPIClient()
	finder := client.GetNFSAPI()

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
		return fmt.Errorf("NFSList is failed: %s", err)
	}

	// get plans
	plans, err := client.NFS.GetNFSPlans()
	if err != nil {
		return fmt.Errorf("NFSList is failed: %s", err)
	}

	list := []interface{}{}
	type nfsValues struct {
		*sacloud.NFS
		PlanName string
		Size     int
	}
	for i := range res.NFS {

		if !params.GetCommandDef().Params["tags"].FilterFunc(list, &res.NFS[i], params.Tags) {
			continue
		}

		var planName string
		var size int

		plan, planDetail := plans.FindByPlanID(res.NFS[i].GetPlanID())
		if planDetail != nil {
			planName = plan.String()
			size = planDetail.Size
		}

		list = append(list, &nfsValues{
			NFS:      &res.NFS[i],
			PlanName: planName,
			Size:     size,
		})
	}
	return ctx.GetOutput().Print(list...)

}
