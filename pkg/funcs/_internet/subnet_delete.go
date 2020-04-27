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

package internet

import (
	"fmt"

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
	"github.com/sacloud/usacloud/pkg/progress"
)

func SubnetDelete(ctx cli.Context, params *params.SubnetDeleteInternetParam) error {

	client := ctx.GetAPIClient()
	api := client.GetInternetAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("InternetSubnetDelete is failed: %s", e)
	}

	err := progress.ExecWithProgress(
		fmt.Sprintf("Still deleting[ID:%d]...", params.SubnetId),
		fmt.Sprintf("Delete subnet[ID:%d]", params.SubnetId),
		ctx.IO().Progress(),
		ctx.Option().NoColor,
		func(compChan chan bool, errChan chan error) {
			_, err := api.DeleteSubnet(params.Id, params.SubnetId)
			if err != nil {
				errChan <- err
				return
			}

			compChan <- true
		},
	)
	if err != nil {
		return fmt.Errorf("InternetSubnetDelete is failed: %s", err)
	}
	return nil
}