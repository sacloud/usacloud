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
	"github.com/sacloud/usacloud/pkg/progress"
)

func Create(ctx cli.Context, params *params.CreateVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p := api.New()

	// validate
	isStandard := params.Plan == "standard"
	if !isStandard {
		targets := []string{"switch-id", "vip", "ipaddress1", "ipaddress2"}
		for _, t := range targets {
			if !ctx.IsSet(t) {
				return fmt.Errorf("%q: is required when plan is [premium/highspec/highspec1600/highspec4000]", t)
			}
		}
	}

	// set params
	switch params.Plan {
	case "standard":
		p.SetStandardPlan()
	case "premium":
		p.SetPremiumPlan(
			fmt.Sprintf("%d", params.SwitchId),
			params.Vip,
			params.Ipaddress1,
			params.Ipaddress2,
			params.Vrid,
			[]string{},
		)
	case "highspec", "highspec1600":
		p.SetHighSpecPlan(
			fmt.Sprintf("%d", params.SwitchId),
			params.Vip,
			params.Ipaddress1,
			params.Ipaddress2,
			params.Vrid,
			[]string{},
		)
	case "highspec4000":
		p.SetHighSpec4000MbpsPlan(
			fmt.Sprintf("%d", params.SwitchId),
			params.Vip,
			params.Ipaddress1,
			params.Ipaddress2,
			params.Vrid,
			[]string{},
		)
	}

	p.SetName(params.Name)
	p.SetDescription(params.Description)
	p.SetTags(params.Tags)
	p.SetIconByID(params.IconId)

	p.InitVPCRouterSetting()
	p.Settings.Router.InternetConnection = &sacloud.VPCRouterInternetConnection{
		Enabled: "True",
	}
	if params.DisableInternetConnection {
		p.Settings.Router.InternetConnection.Enabled = "False"
	}

	// call Create(id)
	res, err := api.Create(p)
	if err != nil {
		return fmt.Errorf("VPCRouterCreate is failed: %s", err)
	}
	// wait for boot
	err = progress.ExecWithProgress(
		fmt.Sprintf("Still creating[ID:%d]...", res.ID),
		fmt.Sprintf("Create vpc-router[ID:%d]", res.ID),
		ctx.IO().Progress(),
		ctx.Option().NoColor,
		func(compChan chan bool, errChan chan error) {
			// call manipurate functions
			err := api.SleepWhileCopying(res.ID, client.DefaultTimeoutDuration, 20)
			if err != nil {
				errChan <- err
				return
			}
			if params.BootAfterCreate {
				_, err := api.Boot(res.ID)
				if err != nil {
					errChan <- err
					return
				}
				err = api.SleepUntilUp(res.ID, client.DefaultTimeoutDuration)
				if err != nil {
					errChan <- err
					return
				}
			}

			compChan <- true
		},
	)
	if err != nil {
		return fmt.Errorf("VPCRouterCreate is failed: %s", err)
	}

	return ctx.Output().Print(res)
}
