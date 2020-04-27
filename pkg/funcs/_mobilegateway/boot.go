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
	"github.com/sacloud/usacloud/pkg/progress"
)

func Boot(ctx cli.Context, params *params.BootMobileGatewayParam) error {

	client := ctx.GetAPIClient()
	api := client.GetMobileGatewayAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("MobileGatewayBoot is failed: %s", e)
	}

	if p.IsUp() {
		return nil // already booted.
	}

	err := progress.ExecWithProgress(
		fmt.Sprintf("Still booting[ID:%d]...", params.Id),
		fmt.Sprintf("Boot mobile-gateway[ID:%d]", params.Id),
		ctx.IO().Progress(),
		ctx.Option().NoColor,
		func(compChan chan bool, errChan chan error) {
			// call manipurate functions
			_, err := api.Boot(params.Id)
			if err != nil {
				errChan <- err
				return
			}
			err = api.SleepUntilUp(params.Id, client.DefaultTimeoutDuration)
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

	return nil

}