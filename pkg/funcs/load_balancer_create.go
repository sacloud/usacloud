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
	"github.com/sacloud/usacloud/pkg/internal"
	"github.com/sacloud/usacloud/pkg/params"
)

func LoadBalancerCreate(ctx cli.Context, params *params.CreateLoadBalancerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetLoadBalancerAPI()

	// validate
	if params.HighAvailability && params.Ipaddress2 == "" {
		return fmt.Errorf("%q: is required when using high-availability flag", "ipaddress2")
	}

	p := &sacloud.CreateLoadBalancerValue{
		SwitchID:     params.SwitchId,
		VRID:         params.Vrid,
		IPAddress1:   params.Ipaddress1,
		MaskLen:      params.NwMaskLen,
		DefaultRoute: params.DefaultRoute,
		Name:         params.Name,
		Description:  params.Description,
		Tags:         params.Tags,
		Icon:         sacloud.NewResource(params.IconId),
	}

	switch params.Plan {
	case "standard":
		p.Plan = sacloud.LoadBalancerPlanStandard
	case "highspec":
		p.Plan = sacloud.LoadBalancerPlanPremium
	}

	var lb *sacloud.LoadBalancer
	var err error
	if params.HighAvailability {
		//冗長構成
		lb, err = sacloud.CreateNewLoadBalancerDouble(&sacloud.CreateDoubleLoadBalancerValue{
			CreateLoadBalancerValue: p,
			IPAddress2:              params.Ipaddress2,
		}, nil)

		if err != nil {
			return fmt.Errorf("LoadBalancerCreate is failed: %s", err)
		}

	} else {
		lb, err = sacloud.CreateNewLoadBalancerSingle(p, nil)
		if err != nil {
			return fmt.Errorf("LoadBalancerCreate is failed: %s", err)
		}
	}

	// call Create(id)
	res, err := api.Create(lb)
	if err != nil {
		return fmt.Errorf("LoadBalancerCreate is failed: %s", err)
	}

	// wait for boot
	err = internal.ExecWithProgress(
		fmt.Sprintf("Still creating[ID:%d]...", res.ID),
		fmt.Sprintf("Create load-balancer[ID:%d]", res.ID),
		ctx.IO().Progress(),
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
		return fmt.Errorf("LoadBalancerCreate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)
}
