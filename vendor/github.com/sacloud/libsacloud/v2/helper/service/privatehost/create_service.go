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

package privatehost

import (
	"context"
	"fmt"

	"github.com/sacloud/libsacloud/v2/sacloud"
)

func (s *Service) Create(req *CreateRequest) (*sacloud.PrivateHost, error) {
	return s.CreateWithContext(context.Background(), req)
}

func (s *Service) CreateWithContext(ctx context.Context, req *CreateRequest) (*sacloud.PrivateHost, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	if req.PlanID.IsEmpty() {
		planOp := sacloud.NewPrivateHostPlanOp(s.caller)
		plans, err := planOp.Find(ctx, req.Zone, nil)
		if err != nil {
			return nil, err
		}
		for _, p := range plans.PrivateHostPlans {
			if p.Class == req.Class {
				req.PlanID = p.ID
				break
			}
		}
		if req.PlanID.IsEmpty() {
			return nil, fmt.Errorf("PrivateHostPlan with Class=%q not found", req.Class)
		}
	}

	params, err := req.ToRequestParameter()
	if err != nil {
		return nil, err
	}

	privateHostOp := sacloud.NewPrivateHostOp(s.caller)
	return privateHostOp.Create(ctx, req.Zone, params)
}
