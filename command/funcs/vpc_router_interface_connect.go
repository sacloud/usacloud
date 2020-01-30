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
	"strconv"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterInterfaceConnect(ctx command.Context, params *params.InterfaceConnectVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterInterfaceConnect is failed: %s", e)
	}
	index, _ := strconv.Atoi(params.Interface)

	// validation
	if !p.IsStandardPlan() {
		// need vip,ip1,ip2
		targets := []string{"ipaddress1", "ipaddress2"}
		for _, t := range targets {
			if !ctx.IsSet(t) {
				return fmt.Errorf("%q: is required when plan is [premium/highspec]", t)
			}
		}
	}
	if p.Interfaces[index].GetSwitch() != nil {
		return fmt.Errorf("Interface[%d] is already connected to switch", index)
	}
	if p.IsUp() {
		return fmt.Errorf("VPCRouter(%d) is still running", params.Id)
	}

	var err error

	_, err = client.Switch.Read(params.SwitchId)
	if err != nil {
		return fmt.Errorf("Switch[%d] is not found", params.SwitchId)
	}

	//connect
	if p.IsStandardPlan() {
		p, err = api.AddStandardInterfaceAt(
			params.Id,
			params.SwitchId,
			params.Ipaddress,
			params.NwMasklen,
			index,
		)
	} else {
		p, err = api.AddPremiumInterfaceAt(
			params.Id,
			params.SwitchId,
			[]string{params.Ipaddress1, params.Ipaddress2},
			params.NwMasklen,
			params.Ipaddress,
			index,
		)
	}
	if err != nil {
		return fmt.Errorf("VPCRouterInterfaceConnect is failed: %s", err)
	}

	if params.WithReboot && p.IsUp() {
		err = internal.ExecWithProgress(
			fmt.Sprintf("Still waiting for reboot[ID:%d]...", params.Id),
			fmt.Sprintf("Connecting interface to switch[ID:%d]", params.Id),
			command.GlobalOption.Progress,
			func(compChan chan bool, errChan chan error) {
				// call manipurate functions
				var err error
				_, err = api.Shutdown(params.Id)
				if err != nil {
					errChan <- err
					return
				}

				err = api.SleepUntilDown(params.Id, client.DefaultTimeoutDuration)
				if err != nil {
					errChan <- err
					return
				}

				_, err = api.Boot(params.Id)
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
			return fmt.Errorf("VPCRouterInterfaceConnect is failed: %s", err)
		}
	}

	return nil

}
