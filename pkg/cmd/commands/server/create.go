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
	Aliases:  []string{"build"}, // v0との互換用
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
	cflag.InputParameter   `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag.NameParameter   `cli:",squash" mapconv:",squash"`
	cflag.DescParameter   `cli:",squash" mapconv:",squash"`
	cflag.TagsParameter   `cli:",squash" mapconv:",squash"`
	cflag.IconIDParameter `cli:",squash" mapconv:",squash"`

	CPU        int    `cli:"cpu,aliases=core,category=plan,order=10" validate:"required"`
	Memory     int    `cli:"memory,category=plan,order=20" mapconv:"MemoryGB" validate:"required"`
	Commitment string `cli:",options=server_plan_commitment,category=plan,order=30" mapconv:",filters=server_plan_commitment_to_value" validate:"required,server_plan_commitment"`
	Generation string `cli:",options=server_plan_generation,category=plan,order=40" mapconv:",filters=server_plan_generation_to_value" validate:"required,server_plan_generation"`

	InterfaceDriver string `cli:",options=interface_dirver" mapconv:",filters=interface_driver_to_value" validate:"required,interface_driver"`

	BootAfterCreate bool
	CDROMID         types.ID `cli:"cdrom-id,aliases=iso-image-id"`
	PrivateHostID   types.ID

	NetworkInterface     server.NetworkInterface    `cli:",category=network,order=10" mapconv:"-" validate:"omitempty"`
	NetworkInterfaceData string                     `cli:"network-interfaces,category=network,order=20" mapconv:"-"`
	NetworkInterfaces    []*server.NetworkInterface `cli:"-" mapconv:",omitempty,recursive"`

	Disk      diskApplyParameter    `cli:",category=disk,order=10" mapconv:"-" validate:"omitempty"`
	DiskIDs   []types.ID            `cli:"disk-ids,category=disk,order=20" mapconv:"-"`
	DisksData string                `cli:"disks,category=disk,order=30" mapconv:"-"`
	Disks     []*diskApplyParameter `cli:"-" mapconv:",omitempty,recursive"`

	cflag.NoWaitParameter `cli:",squash" mapconv:",squash"`
}

type diskApplyParameter struct {
	ID types.ID

	Name                  string `cli:",category=common"` // NOTE: requiredではないためcflag.NameParameterを利用していない
	cflag.DescParameter   `cli:",squash" mapconv:",squash"`
	cflag.TagsParameter   `cli:",squash" mapconv:",squash"`
	cflag.IconIDParameter `cli:",squash" mapconv:",squash"`
	DiskPlan              string `cli:",options=disk_plan" mapconv:"DiskPlanID,filters=disk_plan_to_value" validate:"omitempty,disk_plan"`
	Connection            string `cli:",options=disk_connection" validate:"omitempty,disk_connection"`
	SourceDiskID          types.ID
	SourceArchiveID       types.ID
	ServerID              types.ID
	SizeGB                int `cli:"size,aliases=size-gb"`
	DistantFrom           []types.ID
	OSType                string `cli:",options=os_type_simple" mapconv:",omitempty,filters=os_type_to_value" validate:"omitempty,os_type"`

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
				ID:              diskID,
				Name:            disk.Name,
				DescParameter:   cflag.DescParameter{Description: disk.Description},
				TagsParameter:   cflag.TagsParameter{Tags: disk.Tags},
				IconIDParameter: cflag.IconIDParameter{IconID: disk.IconID},
				Connection:      disk.Connection.String(),
				NoWait:          p.NoWait,
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
