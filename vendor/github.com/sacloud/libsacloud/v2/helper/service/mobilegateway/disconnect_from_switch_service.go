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

	"github.com/sacloud/libsacloud/v2/sacloud"
)

func (s *Service) DisconnectFromSwitch(req *DisconnectFromSwitchRequest) error {
	return s.DisconnectFromSwitchWithContext(context.Background(), req)
}

func (s *Service) DisconnectFromSwitchWithContext(ctx context.Context, req *DisconnectFromSwitchRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	client := sacloud.NewMobileGatewayOp(s.caller)
	return client.DisconnectFromSwitch(ctx, req.Zone, req.ID)
}
