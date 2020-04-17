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

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func PrivateHostCreate(ctx cli.Context, params *params.CreatePrivateHostParam) error {

	client := ctx.GetAPIClient()
	api := client.GetPrivateHostAPI()
	p := api.New()

	// set params

	p.SetName(params.Name)
	p.SetDescription(params.Description)
	p.SetTags(params.Tags)
	p.SetIconByID(params.IconId)

	// set plan(There have only one plan now)
	plans, err := client.Product.GetProductPrivateHostAPI().FilterBy("Class", "dynamic").Find()
	if err != nil || len(plans.PrivateHostPlans) == 0 {
		return fmt.Errorf("PrivateHostCreate is failed: can't find any private-host plan %s", err)
	}
	p.SetPrivateHostPlanByID(plans.PrivateHostPlans[0].ID)

	// call Create(id)
	res, err := api.Create(p)
	if err != nil {
		return fmt.Errorf("PrivateHostCreate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}
