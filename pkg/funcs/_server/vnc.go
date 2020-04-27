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

package server

import (
	"fmt"

	"github.com/sacloud/libsacloud/utils/server"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
	"github.com/sacloud/usacloud/pkg/progress"
)

func Vnc(ctx cli.Context, params *params.VncServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerVnc is failed: %s", e)
	}

	if !p.IsUp() && params.WaitForBoot {

		err := progress.ExecWithProgress(
			fmt.Sprintf("Still booting[ID:%d]...", params.Id),
			fmt.Sprintf("Connect to server[ID:%d]", params.Id),
			ctx.IO().Progress(),
			ctx.Option().NoColor,
			func(compChan chan bool, errChan chan error) {
				// call manipurate functions
				err := api.SleepUntilUp(params.Id, client.DefaultTimeoutDuration)
				if err != nil {
					errChan <- err
					return
				}
				compChan <- true
			},
		)
		if err != nil {
			return fmt.Errorf("ServerVnc is failed: %s", err)
		}
	}
	vncProxyInfo, e := api.GetVNCProxy(params.Id)
	if e != nil {
		return fmt.Errorf("ServerVnc is failed: %s", e)
	}

	return server.StartDefaultVNCClient(vncProxyInfo)
}
