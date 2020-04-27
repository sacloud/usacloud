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

package gslb

import (
	"fmt"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/util"

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func ServerUpdate(ctx cli.Context, params *params.ServerUpdateGSLBParam) error {
	client := sacloud.NewGSLBOp(ctx.Client())
	gslb, err := client.Read(ctx, params.Id)
	if err != nil {
		return fmt.Errorf("GSLBServerUpdate is failed: %s", err)
	}

	if len(gslb.DestinationServers) == 0 {
		return fmt.Errorf("GSLB don't have any servers")
	}

	// validate index
	if err := util.ValidIndex(params.Index, len(gslb.DestinationServers)); err != nil {
		return err
	}

	// validate duplicate
	if params.Changed("ipaddress") {
		for i, s := range gslb.DestinationServers {
			if i != params.Index-1 && s.IPAddress == params.Ipaddress {
				return fmt.Errorf("GSLB already have server(%s)", params.Ipaddress)
			}
		}
	}

	server := gslb.DestinationServers[params.Index-1]

	if params.Changed("ipaddress") {
		server.IPAddress = params.Ipaddress
	}

	if params.Changed("disalbed") {
		server.Enabled = types.StringFlag(!params.Disabled)
	}

	if params.Changed("weight") {
		server.Weight = types.StringNumber(params.Weight)
	}

	gslb, err = client.UpdateSettings(ctx, params.Id, &sacloud.GSLBUpdateSettingsRequest{
		HealthCheck:        gslb.HealthCheck,
		DelayLoop:          gslb.DelayLoop,
		Weighted:           gslb.Weighted,
		SorryServer:        gslb.SorryServer,
		DestinationServers: gslb.DestinationServers,
		SettingsHash:       gslb.SettingsHash,
	})
	if err != nil {
		return fmt.Errorf("GSLBServerUpdate is failed: %s", err)
	}

	var list []interface{}
	for i := range gslb.DestinationServers {
		list = append(list, gslb.DestinationServers[i])
	}

	return ctx.Output().Print(list...)
}
