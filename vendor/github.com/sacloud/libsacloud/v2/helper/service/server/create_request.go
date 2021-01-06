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
	diskService "github.com/sacloud/libsacloud/v2/helper/service/disk"
	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type CreateRequest struct {
	Zone string `request:"-" validate:"required"`

	Name            string `validate:"required"`
	Description     string `validate:"min=0,max=512"`
	Tags            types.Tags
	IconID          types.ID
	CPU             int
	MemoryGB        int
	Commitment      types.ECommitment
	Generation      types.EPlanGeneration
	InterfaceDriver types.EInterfaceDriver

	BootAfterCreate bool
	CDROMID         types.ID
	PrivateHostID   types.ID

	NetworkInterfaces []*NetworkInterface
	Disks             []*diskService.ApplyRequest
	NoWait            bool
}

func (req *CreateRequest) Validate() error {
	return validate.Struct(req)
}

func (req *CreateRequest) ApplyRequest() *ApplyRequest {
	return &ApplyRequest{
		Zone:              req.Zone,
		Name:              req.Name,
		Description:       req.Description,
		Tags:              req.Tags,
		IconID:            req.IconID,
		CPU:               req.CPU,
		MemoryGB:          req.MemoryGB,
		Commitment:        req.Commitment,
		Generation:        req.Generation,
		InterfaceDriver:   req.InterfaceDriver,
		BootAfterCreate:   req.BootAfterCreate,
		CDROMID:           req.CDROMID,
		PrivateHostID:     req.PrivateHostID,
		NetworkInterfaces: req.NetworkInterfaces,
		Disks:             req.Disks,
		NoWait:            req.NoWait,
	}
}
