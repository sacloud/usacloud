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

package server

import (
	"context"
	"fmt"

	"github.com/sacloud/libsacloud/v2/sacloud/types"

	"github.com/sacloud/libsacloud/v2/sacloud"
)

func (s *Service) ChangePlan(req *ChangePlanRequest) (*sacloud.Server, error) {
	return s.ChangePlanWithContext(context.Background(), req)
}

func (s *Service) ChangePlanWithContext(ctx context.Context, req *ChangePlanRequest) (*sacloud.Server, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	client := sacloud.NewServerOp(s.caller)
	current, err := client.Read(ctx, req.Zone, req.ID)
	if err != nil {
		return nil, err
	}
	if !current.InstanceStatus.IsDown() {
		return nil, fmt.Errorf("server[%s] is still running", req.ID)
	}

	changeReq := &sacloud.ServerChangePlanRequest{
		CPU:                  current.CPU,
		MemoryMB:             current.MemoryMB,
		ServerPlanGeneration: current.ServerPlanGeneration,
		ServerPlanCommitment: current.ServerPlanCommitment,
	}
	if req.CPU > 0 {
		changeReq.CPU = req.CPU
	}
	if req.MemoryMB > 0 {
		changeReq.MemoryMB = req.MemoryMB
	}
	if req.ServerPlanGeneration != types.PlanGenerations.Default {
		changeReq.ServerPlanGeneration = req.ServerPlanGeneration
	}
	if req.ServerPlanCommitment != types.Commitments.Unknown {
		changeReq.ServerPlanCommitment = req.ServerPlanCommitment
	}

	return client.ChangePlan(ctx, req.Zone, req.ID, changeReq)
}
