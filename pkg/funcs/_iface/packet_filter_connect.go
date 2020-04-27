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

package iface

import (
	"fmt"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func PacketFilterConnect(ctx cli.Context, params *params.PacketFilterConnectInterfaceParam) error {
	ifClient := sacloud.NewInterfaceOp(ctx.Client())
	pfClient := sacloud.NewPacketFilterOp(ctx.Client())

	iface, err := ifClient.Read(ctx, ctx.Zone(), params.Id)
	if err != nil {
		return fmt.Errorf("InterfacePacketFilterConnect is failed: %s", err)
	}

	// read packet filter
	if _, err := pfClient.Read(ctx, ctx.Zone(), params.PacketFilterId); err != nil {
		return fmt.Errorf("InterfacePacketFilterConnect is failed: %s", err)
	}

	if !iface.PacketFilterID.IsEmpty() {
		return fmt.Errorf("interface is already connected packet filter(%d)", iface.PacketFilterID)
	}

	if err := ifClient.ConnectToPacketFilter(ctx, ctx.Zone(), params.Id, params.PacketFilterId); err != nil {
		return fmt.Errorf("InterfacePacketFilterConnect is failed: %s", err)
	}
	return nil
}
