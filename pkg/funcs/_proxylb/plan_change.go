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

package proxylb

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func PlanChange(ctx cli.Context, params *params.PlanChangeProxyLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetProxyLBAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ProxyLBPlanChange is failed: %s", e)
	}

	// call manipurate functions
	res, err := api.ChangePlan(params.Id, sacloud.ProxyLBPlan(params.Plan))
	if err != nil {
		return fmt.Errorf("ProxyLBPlanChange is failed: %s", err)
	}
	return ctx.Output().Print(res)
}
