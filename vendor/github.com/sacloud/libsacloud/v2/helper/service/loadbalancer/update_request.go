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

package loadbalancer

import (
	"context"
	"fmt"

	"github.com/sacloud/libsacloud/v2/helper/service"
	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type UpdateRequest struct {
	Zone string   `request:"-" validate:"required"`
	ID   types.ID `request:"-" validate:"required"`

	Name               *string                                 `request:",omitempty" validate:"omitempty,min=1"`
	Description        *string                                 `request:",omitempty" validate:"omitempty,min=1,max=512"`
	Tags               *types.Tags                             `request:",omitempty"`
	IconID             *types.ID                               `request:",omitempty"`
	VirtualIPAddresses *sacloud.LoadBalancerVirtualIPAddresses `request:",omitempty"`

	SettingsHash string
	NoWait       bool
}

func (req *UpdateRequest) Validate() error {
	return validate.Struct(req)
}

func (req *UpdateRequest) ApplyRequest(ctx context.Context, caller sacloud.APICaller) (*ApplyRequest, error) {
	client := sacloud.NewLoadBalancerOp(caller)
	current, err := client.Read(ctx, req.Zone, req.ID)
	if err != nil {
		return nil, err
	}
	if current.Availability != types.Availabilities.Available {
		return nil, fmt.Errorf("target has invalid Availability: Zone=%s ID=%s Availability=%v", req.Zone, req.ID.String(), current.Availability)
	}

	applyRequest := &ApplyRequest{
		ID:                 req.ID,
		Zone:               req.Zone,
		Name:               current.Name,
		Description:        current.Description,
		Tags:               current.Tags,
		IconID:             current.IconID,
		SwitchID:           current.SwitchID,
		PlanID:             current.PlanID,
		VRID:               current.VRID,
		IPAddresses:        current.IPAddresses,
		NetworkMaskLen:     current.NetworkMaskLen,
		DefaultRoute:       current.DefaultRoute,
		VirtualIPAddresses: current.VirtualIPAddresses,
	}
	if err := service.RequestConvertTo(req, applyRequest); err != nil {
		return nil, err
	}
	return applyRequest, nil
}
