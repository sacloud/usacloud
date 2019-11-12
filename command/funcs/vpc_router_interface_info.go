// Copyright 2017-2019 The Usacloud Authors
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
	"strings"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func VPCRouterInterfaceInfo(ctx command.Context, params *params.InterfaceInfoVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("VPCRouterInterfaceInfo is failed: %s", e)
	}

	// build parameters to display table
	list := []interface{}{}
	for i, nic := range p.GetInterfaces() {
		v := map[string]interface{}{}

		var nicType string
		if i == 0 {
			nicType = "Global"
		} else {
			nicType = "Private"
		}
		v["Index"] = i
		v["Type"] = nicType

		if nic.GetSwitch() != nil {

			sw := ""
			if nic.GetSwitch().Scope == sacloud.ESCopeShared {
				sw = "shared"
			} else {
				sw = nic.GetSwitch().GetStrID()
			}
			v["Switch"] = sw

			// standard & single nic only
			if !p.IsStandardPlan() && p.Settings != nil && p.Settings.Router.Interfaces[i] != nil {
				v["IPAddress1"] = p.Settings.Router.Interfaces[i].IPAddress[0]
				v["IPAddress2"] = p.Settings.Router.Interfaces[i].IPAddress[1]
				v["Alias"] = strings.Join(p.Settings.Router.Interfaces[i].IPAliases, ",")

			}

			if i == 0 {
				v["NetworkMaskLen"] = nic.GetSwitch().Subnet.NetworkMaskLen
				// IP(VIP)
				ip := ""
				if p.IsStandardPlan() {
					// ip
					ip = nic.IPAddress
				} else {
					// VIP
					ip = p.Settings.Router.Interfaces[i].VirtualIPAddress
				}
				v["IPAddress"] = ip
			} else {
				v["NetworkMaskLen"] = p.Settings.Router.Interfaces[i].NetworkMaskLen
				// IP(VIP)
				ip := ""
				if p.IsStandardPlan() {
					// ip
					ip = p.Settings.Router.Interfaces[i].IPAddress[0]
				} else {
					// VIP
					ip = p.Settings.Router.Interfaces[i].VirtualIPAddress
				}
				v["IPAddress"] = ip
			}
		}

		list = append(list, v)
	}

	return ctx.GetOutput().Print(list...)

}
