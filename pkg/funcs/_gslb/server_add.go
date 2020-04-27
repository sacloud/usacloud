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

	"github.com/sacloud/libsacloud/v2/sacloud/types"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func ServerAdd(ctx cli.Context, params *params.ServerAddGSLBParam) error {
	client := sacloud.NewGSLBOp(ctx.Client())
	gslb, err := client.Read(ctx, params.Id)
	if err != nil {
		return fmt.Errorf("GSLBServerAdd is failed: %s", err)

	}

	// validate maxlen
	if len(gslb.DestinationServers) == 6 {
		return fmt.Errorf("GSLB already have max(6) servers")
	}

	// validate duplicate
	for _, s := range gslb.DestinationServers {
		if s.IPAddress == params.Ipaddress {
			return fmt.Errorf("GSLB already have server(%s)", params.Ipaddress)
		}
	}

	gslb.DestinationServers = append(gslb.DestinationServers, &sacloud.GSLBServer{
		IPAddress: params.Ipaddress,
		Enabled:   true,
		Weight:    types.StringNumber(params.Weight),
	})

	gslb, err = client.Update(ctx, params.Id, &sacloud.GSLBUpdateRequest{
		Name:               gslb.Name,
		Description:        gslb.Description,
		Tags:               gslb.Tags,
		IconID:             gslb.IconID,
		HealthCheck:        gslb.HealthCheck,
		DelayLoop:          gslb.DelayLoop,
		Weighted:           gslb.Weighted,
		SorryServer:        gslb.SorryServer,
		DestinationServers: gslb.DestinationServers,
		SettingsHash:       gslb.SettingsHash,
	})
	if err != nil {
		return fmt.Errorf("GSLBServerAdd is failed: %s", err)
	}

	var list []interface{}
	for i := range gslb.DestinationServers {
		list = append(list, gslb.DestinationServers[i])
	}
	return ctx.Output().Print(list...)
}
