// Copyright 2016-2021 The Libsacloud Authors
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

package mobilegateway

import (
	"context"
	"fmt"

	"github.com/sacloud/libsacloud/v2/sacloud"
)

func (s *Service) UpdateSIM(req *UpdateSIMRequest) error {
	return s.UpdateSIMWithContext(context.Background(), req)
}

func (s *Service) UpdateSIMWithContext(ctx context.Context, req *UpdateSIMRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	mgwOp := sacloud.NewMobileGatewayOp(s.caller)
	simOp := sacloud.NewSIMOp(s.caller)

	sims, err := mgwOp.ListSIM(ctx, req.Zone, req.ID)
	if err != nil {
		return err
	}
	if !sims.Exists(req.SIMID) {
		return fmt.Errorf("SIM[%s] not found in MobileGatewaySIMs", req.SIMID.String())
	}

	return simOp.AssignIP(ctx, req.SIMID, &sacloud.SIMAssignIPRequest{IP: req.IPAddress})
}
