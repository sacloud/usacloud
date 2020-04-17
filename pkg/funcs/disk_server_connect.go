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

func DiskServerConnect(ctx cli.Context, params *params.ServerConnectDiskParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDiskAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DiskServerConnect is failed: %s", e)
	}

	// disk is disconnected from server?
	if p.Server != nil {
		return fmt.Errorf("DiskServerConnect is failed: %s", "Disk is already connected to server")
	}

	// server is exists?
	server, e := client.GetServerAPI().Read(params.ServerId)
	if e != nil || server == nil {
		return fmt.Errorf("DiskServerConnect is failed: Server(ID:%d) isnot exists: %s", params.Id, e)
	}

	// server is stopped?
	if !server.IsDown() {
		return fmt.Errorf("DiskServerConnect is failed: %s", "Server needs to be stopped")
	}

	// call manipurate functions
	_, err := api.ConnectToServer(params.Id, params.ServerId)
	if err != nil {
		return fmt.Errorf("DiskServerConnect is failed: %s", err)
	}

	return nil
}
