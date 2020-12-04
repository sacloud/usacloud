// Copyright 2017-2020 The Usacloud Authors
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
	"fmt"

	"github.com/sacloud/libsacloud/v2/helper/service/disk"

	"github.com/sacloud/libsacloud/v2/helper/service/server"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/commands/common"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/util"
	"github.com/sacloud/usacloud/pkg/validate"
)

var createCommand = &core.Command{
	Name:     "create",
	Category: "basic",
	Order:    20,

	ColumnDefs: defaultColumnDefs,

	ParameterInitializer: func() interface{} {
		return newCreateParameter()
	},

	ValidateFunc: validateCreateParameter,
}

type createParameter struct {
	cflag.ZoneParameter    `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	Name        string   `validate:"required"`
	Description string   `validate:"description"`
	Tags        []string `validate:"tags"`
	IconID      types.ID

	CPU             int    `cli:"cpu,aliases=core" validate:"required"`
	Memory          int    `cli:"memory" mapconv:"MemoryGB" validate:"required"`
	Commitment      string `cli:",options=server_plan_commitment" mapconv:",filters=server_plan_commitment_to_value" validate:"required,server_plan_commitment"`
	Generation      string `cli:",options=server_plan_generation" mapconv:",filters=server_plan_generation_to_value" validate:"required,server_plan_generation"`
	InterfaceDriver string `cli:",options=interface_dirver" mapconv:",filters=interface_driver_to_value" validate:"required,interface_driver"`

	BootAfterCreate bool
	CDROMID         types.ID `cli:"cdrom-id,aliases=iso-image-id"`
	PrivateHostID   types.ID

	NetworkInterface     server.NetworkInterface    `mapconv:"-" validate:"omitempty"`
	NetworkInterfaceData string                     `cli:"network-interfaces" mapconv:"-"`
	NetworkInterfaces    []*server.NetworkInterface `cli:"-" mapconv:",omitempty,recursive"`

	Disk      diskApplyParameter    `mapconv:"-" validate:"omitempty"`
	DisksData string                `cli:"disks" mapconv:"-"`
	Disks     []*diskApplyParameter `cli:"-" mapconv:",omitempty,recursive"`
	DiskIDs   []types.ID            `cli:"disk-ids" mapconv:"-"`

	NoWait bool
}

type diskApplyParameter struct {
	ID types.ID

	Name            string   // 省略時はサーバ名が利用される
	Description     string   `validate:"description"`
	Tags            []string `validate:"tags"`
	IconID          types.ID
	DiskPlan        string `cli:",options=disk_plan" mapconv:"DiskPlanID,filters=disk_plan_to_value" validate:"omitempty,disk_plan"`
	Connection      string `cli:",options=disk_connection" validate:"omitempty,disk_connection"`
	SourceDiskID    types.ID
	SourceArchiveID types.ID
	ServerID        types.ID
	SizeGB          int `cli:"size,aliases=size-gb"`
	DistantFrom     []types.ID
	OSType          string `cli:",options=os_type" mapconv:",omitempty,filters=os_type_to_value" validate:"omitempty,os_type"`

	EditDisk common.EditRequest `cli:"edit,category=edit" mapconv:"EditParameter,omitempty"`
	NoWait   bool
}

func newCreateParameter() *createParameter {
	return &createParameter{
		CPU:             1,
		Memory:          1,
		Commitment:      types.Commitments.Standard.String(),
		Generation:      "default",
		InterfaceDriver: types.InterfaceDrivers.VirtIO.String(),
	}
}

func init() {
	Resource.AddCommand(createCommand)
}

func validateCreateParameter(_ cli.Context, parameter interface{}) error {
	if err := validate.Exec(parameter); err != nil {
		return err
	}

	p, ok := parameter.(*createParameter)
	if !ok {
		return fmt.Errorf("invalid parameter: %v", parameter)
	}

	var errs []error
	// conflict with
	targets := []*validate.Target{
		{FlagName: "--network-interface", Value: p.NetworkInterface},
		{FlagName: "--network-interfaces", Value: p.NetworkInterfaceData},
	}
	if err := validate.ConflictWith(targets...); err != nil {
		errs = append(errs, err)
	}

	targets = []*validate.Target{
		{FlagName: "--disk", Value: p.Disk},
		{FlagName: "--disks", Value: p.Disks},
		{FlagName: "--disk-ids", Value: p.DiskIDs},
	}
	if err := validate.ConflictWith(targets...); err != nil {
		errs = append(errs, err)
	}

	if !util.IsEmpty(p.NetworkInterface) {
		if err := p.NetworkInterface.Validate(); err != nil {
			errs = append(errs, err)
		}
	}
	for _, nic := range p.NetworkInterfaces {
		if err := nic.Validate(); err != nil {
			errs = append(errs, err)
		}
	}
	return validate.NewValidationError(errs...)
}

// Customize パラメータ変換処理
func (p *createParameter) Customize(ctx cli.Context) error {
	// network interfaces
	if !util.IsEmpty(p.NetworkInterface) {
		p.NetworkInterfaces = append(p.NetworkInterfaces, &p.NetworkInterface)
	}
	if p.NetworkInterfaceData != "" {
		var nics []*server.NetworkInterface
		if err := util.MarshalJSONFromPathOrContent(p.NetworkInterfaceData, &nics); err != nil {
			return err
		}
		p.NetworkInterfaces = append(p.NetworkInterfaces, nics...)
	}

	// disk
	if !util.IsEmpty(p.Disk) {
		p.Disks = append(p.Disks, &p.Disk)
	}
	if p.DisksData != "" {
		var disks []*diskApplyParameter
		if err := util.MarshalJSONFromPathOrContent(p.DisksData, &disks); err != nil {
			return err
		}
		p.Disks = append(p.Disks, disks...)
	}
	if len(p.DiskIDs) > 0 {
		diskService := disk.New(ctx.Client())
		for _, diskID := range p.DiskIDs {
			disk, err := diskService.Read(&disk.ReadRequest{
				Zone: p.Zone,
				ID:   diskID,
			})
			if err != nil {
				return err
			}
			p.Disks = append(p.Disks, &diskApplyParameter{
				ID:          diskID,
				Name:        disk.Name,
				Description: disk.Description,
				Tags:        disk.Tags,
				IconID:      disk.IconID,
				Connection:  disk.Connection.String(),
				NoWait:      p.NoWait,
			})
		}
	}

	// set default value to disk
	for _, disk := range p.Disks {
		if disk.Name == "" {
			disk.Name = p.Name
		}
		if disk.DiskPlan == "" {
			disk.DiskPlan = "ssd"
		}
		if disk.Connection == "" {
			disk.Connection = types.DiskConnections.VirtIO.String()
		}

		if err := disk.EditDisk.Customize(ctx); err != nil {
			return err
		}
	}

	return nil
}
