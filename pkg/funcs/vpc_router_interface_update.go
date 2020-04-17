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

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/internal"
	"github.com/sacloud/usacloud/pkg/params"
)

func VPCRouterInterfaceUpdate(ctx cli.Context, params *params.InterfaceUpdateVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterInterfaceUpdate is failed: %s", e)
	}
	index, _ := strconv.Atoi(params.Interface)

	// validation
	if p.Interfaces[index].GetSwitch() == nil {
		return fmt.Errorf("Interface[%d] is already disconnected from switch", index)
	}

	var err error
	if ctx.IsSet("switch-id") {
		_, err = client.Switch.Read(params.SwitchId)
		if err != nil {
			return fmt.Errorf("Switch[%d] is not found", params.SwitchId)
		}

		if p.Interfaces[index].GetSwitch().GetID() == params.SwitchId {
			return fmt.Errorf("Interface[%d] is already connected to switch[%d]", index, params.SwitchId)
		}

		if index == 0 {
			return fmt.Errorf("Can't change switch on interfaces[%d]", index)
		}

		if p.IsUp() {
			return fmt.Errorf("VPCRouter(%d) is still running", params.Id)
		}
	}
	if p.IsStandardPlan() {
		targets := []string{"ipaddress1", "ipaddress2", "alias"}
		for _, t := range targets {
			if ctx.IsSet(t) {
				return fmt.Errorf("%q: can't set when plan is [premium/highspec]", t)
			}
		}
	}
	if ctx.IsSet("alias") && index != 0 {
		return fmt.Errorf("%q: can't set when index is not 0 ", "alias")
	}

	if ctx.IsSet("switch-id") {
		nic := p.Settings.Router.Interfaces[index]
		// disconnect
		_, err = api.DisconnectFromSwitch(params.Id, index)
		if err != nil {
			return fmt.Errorf("VPCRouterInterfaceUpdate is failed: %s", err)
		}

		//connect
		if p.IsStandardPlan() {
			_, err = api.AddStandardInterfaceAt(
				params.Id,
				params.SwitchId,
				nic.IPAddress[0],
				nic.NetworkMaskLen,
				index,
			)
		} else {
			_, err = api.AddPremiumInterfaceAt(
				params.Id,
				params.SwitchId,
				nic.IPAddress,
				nic.NetworkMaskLen,
				nic.VirtualIPAddress,
				index,
			)
		}
		if err != nil {
			return fmt.Errorf("VPCRouterInterfaceUpdate is failed: %s", err)
		}
		p, err = api.Read(params.Id)
		if err != nil {
			return fmt.Errorf("VPCRouterInterfaceUpdate is failed: %s", err)
		}
	}

	// set params
	nic := p.Settings.Router.Interfaces[index]
	if p.IsStandardPlan() {
		if ctx.IsSet("ipaddress") {
			nic.IPAddress[0] = params.Ipaddress
		}
	} else {
		if ctx.IsSet("ipaddress") {
			nic.VirtualIPAddress = params.Ipaddress
		}
		if ctx.IsSet("ipaddress1") {
			nic.IPAddress[0] = params.Ipaddress1
		}
		if ctx.IsSet("ipaddress2") {
			nic.IPAddress[1] = params.Ipaddress2
		}
		if ctx.IsSet("alias") {
			nic.IPAliases = params.Alias
		}
	}
	if ctx.IsSet("nw-masklen") {
		nic.NetworkMaskLen = params.NwMasklen
	}

	p, err = api.UpdateSetting(params.Id, p)
	if err != nil {
		return fmt.Errorf("VPCRouterInterfaceUpdate is failed: %s", err)
	}

	if params.WithReboot && p.IsUp() {
		err = internal.ExecWithProgress(
			fmt.Sprintf("Still waiting for reboot[ID:%d]...", params.Id),
			fmt.Sprintf("Connecting interface to switch[ID:%d]", params.Id),
			ctx.IO().Progress(),
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
			return fmt.Errorf("VPCRouterInterfaceUpdate is failed: %s", err)
		}
	}

	return nil

}
