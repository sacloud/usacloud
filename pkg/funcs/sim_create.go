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
	"github.com/sacloud/usacloud/pkg/define"
	"github.com/sacloud/usacloud/pkg/params"
)

func SIMCreate(ctx cli.Context, params *params.CreateSIMParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSIMAPI()
	p := api.New(params.Name, params.Iccid, params.Passcode)

	// set params
	p.SetDescription(params.Description)
	p.SetTags(params.Tags)
	p.SetIconByID(params.IconId)

	// call Create(id)
	res, err := api.Create(p)
	if err != nil {
		return fmt.Errorf("SIMCreate is failed: %s", err)
	}

	var carriers []*sacloud.SIMNetworkOperatorConfig
	for _, carrier := range params.Carrier {
		carriers = append(carriers, &sacloud.SIMNetworkOperatorConfig{
			Allow: true,
			Name:  define.SIMCarrier[carrier],
		})
	}
	if _, err := api.SetNetworkOperator(res.ID, carriers...); err != nil {
		return fmt.Errorf("SIMCreate is failed: %s", err)
	}

	if !params.Disabled {
		// activate sim
		if _, err := api.Activate(res.ID); err != nil {
			return fmt.Errorf("SIMCreate is failed: %s", err)
		}
	}

	if params.Imei != "" {
		// set imei lock
		if _, err := api.IMEILock(res.ID, params.Imei); err != nil {
			return fmt.Errorf("SIMCreate is failed: %s", err)
		}
	}

	return ctx.GetOutput().Print(res)

}
