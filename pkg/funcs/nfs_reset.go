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

func NFSReset(ctx cli.Context, params *params.ResetNFSParam) error {

	client := ctx.GetAPIClient()
	api := client.GetNFSAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("NFSReset is failed: %s", e)
	}

	err := internal.ExecWithProgress(
		fmt.Sprintf("Still resetting[ID:%d]...", params.Id),
		fmt.Sprintf("Reset nfs[ID:%d]", params.Id),
		ctx.IO().Progress(),
		ctx.Option().NoColor,
		func(compChan chan bool, errChan chan error) {
			// call manipurate functions
			_, err := api.RebootForce(params.Id)
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
		return fmt.Errorf("NFSReset is failed: %s", err)
	}

	return nil

}
