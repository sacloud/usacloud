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

package server

import (
	"context"
	"fmt"

	"github.com/sacloud/libsacloud/v2/helper/service"
	diskService "github.com/sacloud/libsacloud/v2/helper/service/disk"
	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type UpdateRequest struct {
	Zone string   `validate:"required"`
	ID   types.ID `validate:"required"`

	Name            *string                 `request:",omitempty" validate:"omitempty,min=1"`
	Description     *string                 `request:",omitempty" validate:"omitempty,min=0,max=512"`
	Tags            *types.Tags             `request:",omitempty"`
	IconID          *types.ID               `request:",omitempty"`
	CPU             *int                    `request:",omitempty"`
	MemoryGB        *int                    `request:",omitempty"`
	Commitment      *types.ECommitment      `request:",omitempty"`
	Generation      *types.EPlanGeneration  `request:",omitempty"`
	InterfaceDriver *types.EInterfaceDriver `request:",omitempty"`

	CDROMID       *types.ID `request:",omitempty"`
	PrivateHostID *types.ID `request:",omitempty"`

	NetworkInterfaces *[]*NetworkInterface         `request:",omitempty"`
	Disks             *[]*diskService.ApplyRequest `request:",omitempty"`
	NoWait            bool
	ForceShutdown     bool
}

func (req *UpdateRequest) Validate() error {
	return validate.Struct(req)
}

func (req *UpdateRequest) ApplyRequest(ctx context.Context, caller sacloud.APICaller) (*ApplyRequest, error) {
	applyRequest, err := req.applyRequestFromResource(ctx, caller)
	if err != nil {
		return nil, err
	}
	return applyRequest, nil
}

func (req *UpdateRequest) applyRequestFromResource(ctx context.Context, caller sacloud.APICaller) (*ApplyRequest, error) {
	serverOp := sacloud.NewServerOp(caller)
	current, err := serverOp.Read(ctx, req.Zone, req.ID)
	if err != nil {
		return nil, err
	}
	if current.Availability != types.Availabilities.Available {
		return nil, fmt.Errorf("target has invalid Availability: Zone=%s ID=%s Availability=%v", req.Zone, req.ID.String(), current.Availability)
	}

	var nics []*NetworkInterface
	for _, iface := range current.Interfaces {
		var upstream string

		switch {
		case iface.SwitchID.IsEmpty():
			upstream = "disconnected"
		case iface.SwitchScope == types.Scopes.Shared:
			upstream = "shared"
		default:
			upstream = iface.SwitchID.String()
		}

		nics = append(nics, &NetworkInterface{
			Upstream:       upstream,
			PacketFilterID: iface.PacketFilterID,
			UserIPAddress:  iface.UserIPAddress,
		})
	}

	diskOp := sacloud.NewDiskOp(caller)
	var disks []*diskService.ApplyRequest
	for _, d := range current.Disks {
		disk, err := diskOp.Read(ctx, req.Zone, d.ID)
		if err != nil {
			return nil, err
		}
		disks = append(disks, &diskService.ApplyRequest{
			Zone:        req.Zone,
			ID:          disk.ID,
			Name:        disk.Name,
			Description: disk.Description,
			Tags:        disk.Tags,
			IconID:      disk.IconID,
			DiskPlanID:  disk.DiskPlanID,
			Connection:  disk.Connection,
			ServerID:    current.ID,
			SizeGB:      disk.GetSizeGB(),
			NoWait:      req.NoWait,
		})
	}

	applyRequest := &ApplyRequest{
		Zone:              req.Zone,
		ID:                req.ID,
		Name:              current.Name,
		Description:       current.Description,
		Tags:              current.Tags,
		IconID:            current.IconID,
		CPU:               current.CPU,
		MemoryGB:          current.GetMemoryGB(),
		Commitment:        current.ServerPlanCommitment,
		Generation:        current.ServerPlanGeneration,
		InterfaceDriver:   current.InterfaceDriver,
		CDROMID:           current.CDROMID,
		PrivateHostID:     current.PrivateHostID,
		NetworkInterfaces: nics,
		Disks:             disks,
	}

	if err := service.RequestConvertTo(req, applyRequest); err != nil {
		return nil, err
	}
	return applyRequest, nil
}
