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

package mobilegateway

import (
	"fmt"

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func SIMInfo(ctx cli.Context, params *params.SIMInfoMobileGatewayParam) error {

	client := ctx.GetAPIClient()
	api := client.GetMobileGatewayAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("MobileGatewaySIMInfo is failed: %s", e)
	}

	sims, err := api.ListSIM(p.ID, nil)
	if err != nil {
		return fmt.Errorf("MobileGatewaySIMInfo is failed: %s", e)
	}

	list := []interface{}{}
	for i := range sims {
		list = append(list, &sims[i])
	}
	return ctx.Output().Print(list...)
}