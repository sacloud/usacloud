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

func Update(ctx cli.Context, params *params.UpdateGSLBParam) error {
	client := sacloud.NewGSLBOp(ctx.Client())
	gslb, err := client.Read(ctx, params.Id)
	if err != nil {
		return fmt.Errorf("GSLBUpdate is failed: %s", err)
	}

	// validation
	if (params.Protocol == "http" || params.Protocol == "https") && (params.Path == "" || params.ResponseCode == 0) {
		return fmt.Errorf("path and response-code is required when protocol is http")
	}
	if params.Protocol == "tcp" && params.Port == 0 {
		return fmt.Errorf("port is required when protocol is tcp")
	}

	req := &sacloud.GSLBUpdateRequest{
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
	}

	// set params
	if params.Changed("icon-id") {
		req.IconID = params.IconId
	}
	if params.Changed("name") {
		req.Name = params.Name
	}
	if params.Changed("description") {
		req.Description = params.Description
	}
	if params.Changed("tags") {
		req.Tags = params.Tags
	}
	if params.Changed("protocol") {
		req.HealthCheck.Protocol = types.EGSLBHealthCheckProtocol(params.Protocol)
	}
	if params.Changed("host-header") {
		req.HealthCheck.HostHeader = params.HostHeader
	}
	if params.Changed("path") {
		req.HealthCheck.Path = params.Path
	}
	if params.Changed("response-code") {
		req.HealthCheck.ResponseCode = types.StringNumber(params.ResponseCode)
	}
	if params.Changed("port") {
		req.HealthCheck.Port = types.StringNumber(params.Port)
	}
	if params.Changed("sorry-server") {
		req.SorryServer = params.SorryServer
	}
	if params.Changed("weighted") {
		req.Weighted = types.StringFlag(params.Weighted)
	}
	if params.Changed("delay-loop") {
		req.DelayLoop = params.DelayLoop
	}

	// call Update(id)
	gslb, err = client.Update(ctx, params.Id, req)
	if err != nil {
		return fmt.Errorf("GSLBUpdate is failed: %s", err)
	}

	return ctx.Output().Print(gslb)
}
