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

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func InterfaceAddDisconnected(ctx cli.Context, params *params.InterfaceAddDisconnectedServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerInterfaceAddDisconnected is failed: %s", e)
	}

	if len(p.GetInterfaces()) >= sacloud.ServerMaxInterfaceLen {
		return fmt.Errorf("Server already connected maximum count of interfaces")
	}

	if p.IsUp() {
		return fmt.Errorf("ServerInterfaceAddDisconnected is failed: %s", "server is running")
	}

	// call manipurate functions
	interfaceAPI := client.GetInterfaceAPI()
	_, err := interfaceAPI.CreateAndConnectToServer(params.Id)
	if err != nil {
		return fmt.Errorf("ServerInterfaceAddDisconnected is failed: %s", err)
	}

	return nil

}