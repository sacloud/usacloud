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

package disk

import (
	"fmt"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func ServerConnect(ctx cli.Context, params *params.ServerConnectDiskParam) error {
	client := sacloud.NewDiskOp(ctx.Client())
	serverClient := sacloud.NewServerOp(ctx.Client())

	disk, err := client.Read(ctx, ctx.Zone(), params.Id)
	if err != nil {
		return err
	}

	// disk is disconnected from server?
	if !disk.ServerID.IsEmpty() {
		return fmt.Errorf("disk is already connected to server[%d]", disk.ServerID)
	}

	// server is exists?
	server, err := serverClient.Read(ctx, ctx.Zone(), params.ServerId)
	if err != nil {
		return fmt.Errorf("reading server[%d] returns error: %s", disk.ServerID, err)
	}

	// server is stopped?
	if !server.InstanceStatus.IsDown() {
		return fmt.Errorf("server[%d] is running", disk.ServerID)
	}

	return client.ConnectToServer(ctx, ctx.Zone(), params.Id, params.ServerId)
}
