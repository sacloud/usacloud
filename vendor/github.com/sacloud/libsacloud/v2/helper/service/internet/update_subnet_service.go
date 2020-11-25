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

package internet

import (
	"context"
	"fmt"

	"github.com/sacloud/libsacloud/v2/sacloud"
)

func (s *Service) UpdateSubnet(req *UpdateSubnetRequest) (*sacloud.Subnet, error) {
	return s.UpdateSubnetWithContext(context.Background(), req)
}

func (s *Service) UpdateSubnetWithContext(ctx context.Context, req *UpdateSubnetRequest) (*sacloud.Subnet, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	client := sacloud.NewInternetOp(s.caller)
	current, err := client.Read(ctx, req.Zone, req.ID)
	if err != nil {
		return nil, fmt.Errorf("reading the internet resource[%s] failed: %s", req.ID, err)
	}

	params, err := req.ToRequestParameter(current)
	if err != nil {
		return nil, fmt.Errorf("processing request parameter failed: %s", err)
	}
	result, err := client.UpdateSubnet(ctx, req.Zone, req.ID, req.SubnetID, params)
	if err != nil {
		return nil, err
	}

	subnetOp := sacloud.NewSubnetOp(s.caller)
	return subnetOp.Read(ctx, req.Zone, result.ID)
}
