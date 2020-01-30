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
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func InterfacePacketFilterConnect(ctx command.Context, params *params.PacketFilterConnectInterfaceParam) error {

	client := ctx.GetAPIClient()
	api := client.GetInterfaceAPI()
	packetFilterAPI := client.GetPacketFilterAPI()

	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("InterfacePacketFilterConnect is failed: %s", e)
	}

	// read packet filter
	_, e = packetFilterAPI.Read(params.PacketFilterId)
	if e != nil {
		return fmt.Errorf("InterfacePacketFilterConnect is failed: %s", e)
	}

	if p.PacketFilter != nil && p.PacketFilter.ID > sacloud.EmptyID {
		return fmt.Errorf("interface is already connected packet filter(%d)", p.PacketFilter.ID)
	}

	// call manipurate functions
	_, err := api.ConnectToPacketFilter(params.Id, params.PacketFilterId)
	if err != nil {
		return fmt.Errorf("InterfacePacketFilterConnect is failed: %s", err)
	}
	return nil
}
