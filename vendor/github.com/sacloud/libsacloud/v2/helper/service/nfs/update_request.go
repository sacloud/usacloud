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

package nfs

import (
	"context"

	"github.com/sacloud/libsacloud/v2/helper/query"

	"github.com/sacloud/libsacloud/v2/helper/service"
	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type UpdateRequest struct {
	Zone string   `validate:"required"`
	ID   types.ID `validate:"required"`

	Name        *string     `request:",omitempty" validate:"omitempty,min=1"`
	Description *string     `request:",omitempty" validate:"omitempty,min=1,max=512"`
	Tags        *types.Tags `request:",omitempty"`
	IconID      *types.ID   `request:",omitempty"`
}

func (req *UpdateRequest) Validate() error {
	return validate.Struct(req)
}

func (req *UpdateRequest) ApplyRequest(ctx context.Context, caller sacloud.APICaller) (*ApplyRequest, error) {
	client := sacloud.NewNFSOp(caller)
	current, err := client.Read(ctx, req.Zone, req.ID)
	if err != nil {
		return nil, err
	}

	// find plan
	plan, err := query.GetNFSPlanInfo(ctx, sacloud.NewNoteOp(caller), current.PlanID)
	if err != nil {
		return nil, err
	}

	applyRequest := &ApplyRequest{
		ID:             current.ID,
		Zone:           req.Zone,
		Name:           current.Name,
		Description:    current.Description,
		Tags:           current.Tags,
		IconID:         current.IconID,
		SwitchID:       current.SwitchID,
		Plan:           plan.DiskPlanID,
		Size:           plan.Size,
		IPAddresses:    current.IPAddresses,
		NetworkMaskLen: current.NetworkMaskLen,
		DefaultRoute:   current.DefaultRoute,
	}
	if err := service.RequestConvertTo(req, applyRequest); err != nil {
		return nil, err
	}
	return applyRequest, nil
}
