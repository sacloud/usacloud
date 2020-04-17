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

func LoadBalancerVipInfo(ctx cli.Context, params *params.VipInfoLoadBalancerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetLoadBalancerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("LoadBalancerVipInfo is failed: %s", e)
	}
	initLoadBalancerSettings(p)

	vips := p.Settings.LoadBalancer
	if len(vips) == 0 {
		fmt.Fprintf(ctx.IO().Err(), "LoadBalancer don't have any VIPs\n")
		return nil
	}

	// build parameters to display table
	list := []interface{}{}
	for i := range vips {
		list = append(list, &vips[i])
	}

	return ctx.GetOutput().Print(list...)

}
