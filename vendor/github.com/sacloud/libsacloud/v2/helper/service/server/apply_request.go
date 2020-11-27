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
	diskBuilder "github.com/sacloud/libsacloud/v2/helper/builder/disk"
	serverBuilder "github.com/sacloud/libsacloud/v2/helper/builder/server"
	diskService "github.com/sacloud/libsacloud/v2/helper/service/disk"

	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/pkg/size"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type ApplyRequest struct {
	Zone string   `request:"-" validate:"required"`
	ID   types.ID `request:"-"`

	Name                 string `validate:"required"`
	Description          string `validate:"min=0,max=512"`
	Tags                 types.Tags
	IconID               types.ID
	CPU                  int
	MemoryMB             int
	ServerPlanCommitment types.ECommitment
	ServerPlanGeneration types.EPlanGeneration
	ConnectedSwitches    []*sacloud.ConnectedSwitch
	InterfaceDriver      types.EInterfaceDriver

	BootAfterCreate bool
	CDROMID         types.ID
	PrivateHostID   types.ID

	// TODO builderパッケージを移植するまではserverBuilderへ依存させておく

	NIC            serverBuilder.NICSettingHolder
	AdditionalNICs []serverBuilder.AdditionalNICSettingHolder
	Disks          []*diskService.ApplyRequest

	ForceShutdown bool
}

func (req *ApplyRequest) Validate() error {
	return validate.Struct(req)
}

func (req *ApplyRequest) Builder(caller sacloud.APICaller) (*serverBuilder.Builder, error) {
	var diskBuilders []diskBuilder.Builder
	for _, d := range req.Disks {
		b, err := d.Builder(caller)
		if err != nil {
			return nil, err
		}
		diskBuilders = append(diskBuilders, b)
	}

	return &serverBuilder.Builder{
		Name:            req.Name,
		CPU:             req.CPU,
		MemoryGB:        req.MemoryMB * size.GiB,
		Commitment:      req.ServerPlanCommitment,
		Generation:      req.ServerPlanGeneration,
		InterfaceDriver: req.InterfaceDriver,
		Description:     req.Description,
		IconID:          req.IconID,
		Tags:            req.Tags,
		BootAfterCreate: req.BootAfterCreate,
		CDROMID:         req.CDROMID,
		PrivateHostID:   req.PrivateHostID,
		NIC:             req.NIC,
		AdditionalNICs:  req.AdditionalNICs,
		DiskBuilders:    diskBuilders,
		Client:          serverBuilder.NewBuildersAPIClient(caller),
		ServerID:        req.ID,
		ForceShutdown:   req.ForceShutdown,
	}, nil
}
