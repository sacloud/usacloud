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
	"github.com/sacloud/usacloud/pkg/util/progress"
)

func MobileGatewayCreate(ctx cli.Context, params *params.CreateMobileGatewayParam) error {

	client := ctx.GetAPIClient()
	api := client.GetMobileGatewayAPI()

	createMGWValues := &sacloud.CreateMobileGatewayValue{
		Name:        params.Name,
		Description: params.Description,
		Tags:        params.Tags,
	}
	mgwSetting := &sacloud.MobileGatewaySetting{
		InternetConnection: &sacloud.MGWInternetConnection{
			Enabled: "False",
		},
		Interfaces: []*sacloud.MGWInterface{
			nil,
		},
	}
	if params.InternetConnection {
		mgwSetting.InternetConnection.Enabled = "True"
	}

	p, err := sacloud.CreateNewMobileGateway(createMGWValues, mgwSetting)
	if err != nil {
		return fmt.Errorf("MobileGatewayCreate is failed: %s", err)
	}

	p.SetIconByID(params.IconId)

	var res *sacloud.MobileGateway
	err = progress.ExecWithProgress(
		fmt.Sprintf("Still creating..."),
		fmt.Sprintf("Create mobile-gateway"),
		ctx.IO().Progress(),
		ctx.Option().NoColor,
		func(compChan chan bool, errChan chan error) {
			// call manipurate functions
			// call Create(id)
			res, err = api.Create(p)
			if err != nil {
				errChan <- fmt.Errorf("MobileGatewayCreate is failed: %s", err)
				return
			}
			err = api.SleepWhileCopying(res.ID, client.DefaultTimeoutDuration, 20)
			if err != nil {
				errChan <- err
				return
			}

			compChan <- true
		},
	)
	if err != nil {
		return fmt.Errorf("MobileGatewayBoot is failed: %s", err)
	}

	res, err = api.Read(res.ID)
	if err != nil {
		return fmt.Errorf("MobileGatewayBoot is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)
}
