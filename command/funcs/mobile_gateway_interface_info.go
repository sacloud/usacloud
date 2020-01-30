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

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func MobileGatewayInterfaceInfo(ctx command.Context, params *params.InterfaceInfoMobileGatewayParam) error {

	client := ctx.GetAPIClient()
	api := client.GetMobileGatewayAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("MobileGatewayInterfaceInfo is failed: %s", e)
	}

	res := []interface{}{
		map[string]interface{}{
			"Type":           "Public",
			"Switch":         "shared",
			"IPAddress":      p.Interfaces[0].IPAddress,
			"NetworkMaskLen": p.Interfaces[0].Switch.Subnet.NetworkMaskLen,
		},
	}

	if len(p.Settings.MobileGateway.Interfaces) > 1 && p.Settings.MobileGateway.Interfaces[1] != nil {
		res = append(res, map[string]interface{}{
			"Type":           "Private",
			"Switch":         p.Interfaces[1].Switch.GetStrID(),
			"IPAddress":      p.Settings.MobileGateway.Interfaces[1].IPAddress[0],
			"NetworkMaskLen": p.Settings.MobileGateway.Interfaces[1].NetworkMaskLen,
		})
	}

	return ctx.GetOutput().Print(res...)
}
