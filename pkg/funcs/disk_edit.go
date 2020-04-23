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
	"github.com/sacloud/usacloud/pkg/progress"
	"github.com/sacloud/usacloud/pkg/utils"
)

func DiskEdit(ctx cli.Context, params *params.EditDiskParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDiskAPI()
	p := buildDiskEditValue(ctx, params)

	// wait for copy with progress
	err := progress.ExecWithProgress(
		fmt.Sprintf("Still editing[ID:%d]...", params.Id),
		fmt.Sprintf("Edit disk[ID:%d]", params.Id),
		ctx.IO().Progress(),
		ctx.Option().NoColor,
		func(compChan chan bool, errChan chan error) {
			// call manipurate functions
			_, err := api.Config(params.Id, p)
			if err != nil {
				errChan <- err
				return
			}
			if err := api.SleepWhileCopying(params.Id, client.DefaultTimeoutDuration); err != nil {
				errChan <- err
				return
			}

			compChan <- true
		},
	)

	if err != nil {
		return fmt.Errorf("DiskEdit is failed: %s", err)
	}

	// read
	res, err := api.Read(params.Id)
	if err != nil {
		return fmt.Errorf("DiskEdit is failed: %s", err)
	}
	return ctx.GetOutput().Print(res)
}

func buildDiskEditValue(ctx cli.Context, params *params.EditDiskParam) *sacloud.DiskEditValue {
	p := ctx.GetAPIClient().GetDiskAPI().NewCondig()
	p.SetBackground(true)

	// set params
	if ctx.IsSet("hostname") {
		p.SetHostName(params.Hostname)
	}
	if ctx.IsSet("password") {
		p.SetPassword(params.Password)
	}
	if ctx.IsSet("ssh-key-ids") {
		p.SetSSHKeys(utils.StringIDs(params.SSHKeyIds))
	}
	if ctx.IsSet("disable-password-auth") {
		p.SetDisablePWAuth(params.DisablePasswordAuth)
	}
	if ctx.IsSet("startup-script-ids") {
		p.SetNotes(utils.StringIDs(params.StartupScriptIds))
	}
	if ctx.IsSet("ipaddress") {
		p.SetUserIPAddress(params.Ipaddress)
	}
	if ctx.IsSet("default-route") {
		p.SetDefaultRoute(params.DefaultRoute)
	}
	if ctx.IsSet("nw-masklen") {
		p.SetNetworkMaskLen(fmt.Sprintf("%d", params.NwMasklen))
	}

	return p
}
