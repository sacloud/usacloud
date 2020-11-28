// Copyright 2016-2020 The Libsacloud Authors
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

package vpcrouter

import (
	"context"

	"github.com/sacloud/libsacloud/v2/helper/power"
	"github.com/sacloud/libsacloud/v2/sacloud"
)

func (s *Service) Boot(req *BootRequest) error {
	return s.BootWithContext(context.Background(), req)
}

func (s *Service) BootWithContext(ctx context.Context, req *BootRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	client := sacloud.NewVPCRouterOp(s.caller)
	if req.NoWait {
		return client.Boot(ctx, req.Zone, req.ID)
	}
	return power.BootVPCRouter(ctx, client, req.Zone, req.ID)
}
