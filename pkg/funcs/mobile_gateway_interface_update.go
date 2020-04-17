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
	"github.com/sacloud/usacloud/pkg/params"
)

func MobileGatewayInterfaceUpdate(ctx cli.Context, params *params.InterfaceUpdateMobileGatewayParam) error {

	client := ctx.GetAPIClient()
	api := client.GetMobileGatewayAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("MobileGatewayInterfaceUpdate is failed: %s", e)
	}

	if len(p.Interfaces) < 2 || p.Interfaces[1].GetSwitch() == nil {
		return fmt.Errorf("Interface[%s] is already disconnected from switch", "eth1")
	}
	if p.IsUp() {
		return fmt.Errorf("MobileGateway(%d) is still running", params.Id)
	}

	var err error

	p.SetPrivateInterface(params.Ipaddress, params.NwMasklen)
	_, err = api.Update(p.ID, p)
	if err != nil {
		return fmt.Errorf("MobileGatewayInterfaceConnect is failed: %s", err)
	}

	_, err = api.Config(p.ID)
	if err != nil {
		return fmt.Errorf("MobileGatewayInterfaceConnect is failed: %s", err)
	}

	return nil

}
