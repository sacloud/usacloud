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

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func Create(ctx cli.Context, params *params.CreateGSLBParam) error {
	client := sacloud.NewGSLBOp(ctx.Client())
	gslb, err := client.Create(ctx, &sacloud.GSLBCreateRequest{
		HealthCheck: &sacloud.GSLBHealthCheck{
			Protocol:     types.EGSLBHealthCheckProtocol(params.Protocol),
			HostHeader:   params.HostHeader,
			Path:         params.Path,
			ResponseCode: types.StringNumber(params.ResponseCode),
			Port:         types.StringNumber(params.Port),
		},
		DelayLoop:   params.DelayLoop,
		Weighted:    types.StringFlag(params.Weighted),
		SorryServer: params.SorryServer,
		Name:        params.Name,
		Description: params.Description,
		Tags:        params.Tags,
		IconID:      params.IconId,
	})
	if err != nil {
		return fmt.Errorf("GSLBCreate is failed: %s", err)
	}

	return ctx.Output().Print(gslb)
}
