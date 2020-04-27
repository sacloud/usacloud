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

package nfs

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
	"github.com/sacloud/usacloud/pkg/progress"
)

func Create(ctx cli.Context, params *params.CreateNFSParam) error {

	client := ctx.GetAPIClient()
	api := client.GetNFSAPI()

	p := &sacloud.CreateNFSValue{
		SwitchID:     params.SwitchId,
		IPAddress:    params.Ipaddress,
		MaskLen:      params.NwMaskLen,
		DefaultRoute: params.DefaultRoute,
		Name:         params.Name,
		Description:  params.Description,
		Tags:         params.Tags,
		Icon:         sacloud.NewResource(params.IconId),
	}

	var plan sacloud.NFSPlan
	switch params.Plan {
	case "ssd":
		plan = sacloud.NFSPlanSSD
	case "hdd":
		plan = sacloud.NFSPlanHDD
	default:
		return fmt.Errorf("NFSCreate is failed: invalid plan %s", params.Plan)
	}

	// call Create(id)
	res, err := api.CreateWithPlan(p, plan, sacloud.NFSSize(params.Size))
	if err != nil {
		return fmt.Errorf("NFSCreate is failed: %s", err)
	}

	// wait for boot
	err = progress.ExecWithProgress(
		fmt.Sprintf("Still creating[ID:%d]...", res.ID),
		fmt.Sprintf("Create nfs[ID:%d]", res.ID),
		ctx.IO().Progress(),
		ctx.Option().NoColor,
		func(compChan chan bool, errChan chan error) {
			// call manipurate functions
			err := api.SleepWhileCopying(res.ID, client.DefaultTimeoutDuration, 20)
			if err != nil {
				errChan <- err
				return
			}
			err = api.SleepUntilUp(res.ID, client.DefaultTimeoutDuration)
			if err != nil {
				errChan <- err
				return
			}
			compChan <- true
		},
	)
	if err != nil {
		return fmt.Errorf("NFSCreate is failed: %s", err)
	}

	return ctx.Output().Print(res)

}
