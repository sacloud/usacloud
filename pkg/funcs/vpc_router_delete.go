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
	"github.com/sacloud/usacloud/pkg/internal"
	"github.com/sacloud/usacloud/pkg/params"
)

func VPCRouterDelete(ctx cli.Context, params *params.DeleteVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()

	p, err := api.Read(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterDelete is failed: %s", err)
	}

	if p.IsUp() {
		if params.Force {

			err = internal.ExecWithProgress(
				fmt.Sprintf("Still waiting for delete[ID:%d]...", params.Id),
				fmt.Sprintf("Delete vpc-router[ID:%d]", params.Id),
				ctx.IO().Progress(),
				func(compChan chan bool, errChan chan error) {
					// call manipurate functions
					var err error
					_, err = api.Stop(params.Id)
					if err != nil {
						errChan <- err
						return
					}

					err = api.SleepUntilDown(params.Id, client.DefaultTimeoutDuration)
					if err != nil {
						errChan <- err
						return
					}
					compChan <- true
				},
			)
			if err != nil {
				return fmt.Errorf("VPCRouterDelete is failed: %s", err)
			}

		} else {
			return fmt.Errorf("VPCRouter(%d) is still running", params.Id)
		}
	}

	// call Delete(id)
	res, err := api.Delete(params.Id)
	if err != nil {
		return fmt.Errorf("VPCRouterDelete is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}
