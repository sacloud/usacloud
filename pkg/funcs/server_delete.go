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
	"os"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
	"github.com/sacloud/usacloud/pkg/progress"
)

func ServerDelete(ctx cli.Context, params *params.DeleteServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()

	p, err := api.Read(params.Id)
	if err != nil {
		return fmt.Errorf("ServerDelete is failed: %s", err)
	}

	if p.IsUp() {
		if params.Force {

			err := progress.ExecWithProgress(
				fmt.Sprintf("Still waiting for Delete[ID:%d]...", params.Id),
				fmt.Sprintf("Delete server[ID:%d]", params.Id),
				ctx.IO().Progress(),
				ctx.Option().NoColor,
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
				return fmt.Errorf("ServerDelete is failed: %s", err)
			}
		} else {
			return fmt.Errorf("Server(%d) is still running", params.Id)
		}
	}

	// call Delete(id)
	var res *sacloud.Server
	if !params.WithoutDisk && len(p.Disks) > 0 {
		res, err = api.DeleteWithDisk(params.Id, p.GetDiskIDs())
		if err != nil {
			return fmt.Errorf("ServerDelete is failed: %s", err)
		}
	} else {
		res, err = api.Delete(params.Id)
		if err != nil {
			return fmt.Errorf("ServerDelete is failed: %s", err)
		}
	}

	// Delete generated ssh-key on default location
	keyFile, err := getSSHPrivateKeyStorePath(res.ID)
	if err != nil {
		return fmt.Errorf("ServerDelete is failed: %s", err)
	}
	if _, e := os.Stat(keyFile); e == nil {
		err = os.Remove(keyFile)
		if err != nil {
			return fmt.Errorf("ServerDelete is failed: %s", err)
		}
	}

	return ctx.GetOutput().Print(res)
}
