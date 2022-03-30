// Copyright 2017-2022 The Usacloud Authors
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

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/usacloud/pkg/cmd/examples"

	"github.com/sacloud/iaas-service-go/disk"

	"github.com/sacloud/iaas-api-go/types"
	"github.com/sacloud/iaas-service-go/server"
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
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag.NameParameter   `cli:",squash" mapconv:",squash"`
	cflag.DescParameter   `cli:",squash" mapconv:",squash"`
	cflag.TagsParameter   `cli:",squash" mapconv:",squash"`
	cflag.IconIDParameter `cli:",squash" mapconv:",squash"`

	CPU        int    `cli:"cpu,aliases=core,category=plan,order=10" validate:"required"`
	Memory     int    `cli:"memory,category=plan,order=20" mapconv:"MemoryGB" validate:"required"`
	GPU        int    `cli:"gpu,category=plan,order=30"`
	Commitment string `cli:",options=server_plan_commitment,category=plan,order=30" mapconv:",filters=server_plan_commitment_to_value" validate:"required,server_plan_commitment"`
	Generation string `cli:",options=server_plan_generation,category=plan,order=40" mapconv:",filters=server_plan_generation_to_value" validate:"required,server_plan_generation"`

	InterfaceDriver string `cli:",options=interface_driver" mapconv:",filters=interface_driver_to_value" validate:"required,interface_driver"`

	BootAfterCreate bool
	CDROMID         types.ID `cli:"cdrom-id,aliases=iso-image-id"`
	PrivateHostID   types.ID

	NetworkInterface     serverNetworkInterface     `cli:",category=network,order=10" mapconv:"-" validate:"omitempty" json:"-"`
	NetworkInterfaceData string                     `cli:"network-interfaces,category=network,order=20" mapconv:"-" json:"-"`
	NetworkInterfaces    []*server.NetworkInterface `cli:"-" mapconv:",omitempty,recursive"`

	Disk      diskApplyParameter    `cli:",category=disk,order=10" mapconv:"-" validate:"omitempty" json:"-"`
	DiskIDs   []types.ID            `cli:"disk-ids,category=disk,order=20" mapconv:"-" json:"-"`
	DisksData string                `cli:"disks,category=disk,order=30" mapconv:"-" json:"-"`
	Disks     []*diskApplyParameter `cli:"-" mapconv:",omitempty,recursive"`

	cflag.NoWaitParameter `cli:",squash" mapconv:",squash"`
}

type serverNetworkInterface struct {
	Upstream       string `cli:",display_options=shared/disconnected/(switch-id),options=shared disconnected"`
	PacketFilterID types.ID
	UserIPAddress  string `validate:"omitempty,ipv4"`
}

type diskApplyParameter struct {
	ID types.ID `json:",omitempty"`

	Name                  string `cli:",category=common" json:",omitempty"` // NOTE: requiredではないためcflag.NameParameterを利用していない
	cflag.DescParameter   `cli:",squash" mapconv:",squash" json:",omitempty"`
	cflag.TagsParameter   `cli:",squash" mapconv:",squash" json:",omitempty"`
	cflag.IconIDParameter `cli:",squash" mapconv:",squash" json:",omitempty"`
	DiskPlan              string     `cli:",options=disk_plan" mapconv:"DiskPlanID,filters=disk_plan_to_value" validate:"omitempty,disk_plan" json:",omitempty"`
	Connection            string     `cli:",options=disk_connection" validate:"omitempty,disk_connection" json:",omitempty"`
	SourceDiskID          types.ID   `json:",omitempty"`
	SourceArchiveID       types.ID   `json:",omitempty"`
	SizeGB                int        `cli:"size,aliases=size-gb" json:",omitempty"`
	DistantFrom           []types.ID `json:",omitempty"`
	OSType                string     `cli:",options=os_type,display_options=os_type_simple" mapconv:",omitempty,filters=os_type_to_value" validate:"omitempty,os_type" json:",omitempty"`

	EditDisk common.EditRequest `cli:"edit,category=edit" mapconv:"EditParameter,omitempty" json:",omitempty"`
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

	if !util.IsEmpty(p.Disk) {
		if err := p.Disk.EditDisk.Validate(); err != nil {
			errs = append(errs, err)
		}
	}
	for _, d := range p.Disks {
		if err := d.EditDisk.Validate(); err != nil {
			errs = append(errs, err)
		}
	}

	if !util.IsEmpty(p.NetworkInterface) {
		nic := &server.NetworkInterface{
			Upstream:       p.NetworkInterface.Upstream,
			PacketFilterID: p.NetworkInterface.PacketFilterID,
			UserIPAddress:  p.NetworkInterface.UserIPAddress,
		}
		if err := nic.Validate(); err != nil {
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
		nic := &server.NetworkInterface{
			Upstream:       p.NetworkInterface.Upstream,
			PacketFilterID: p.NetworkInterface.PacketFilterID,
			UserIPAddress:  p.NetworkInterface.UserIPAddress,
		}
		p.NetworkInterfaces = append(p.NetworkInterfaces, nic)
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

func (p *createParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &createParameter{
		ZoneParameter:   examples.Zones(ctx.Option().Zones),
		NameParameter:   examples.Name,
		DescParameter:   examples.Description,
		TagsParameter:   examples.Tags,
		IconIDParameter: examples.IconID,
		CPU:             1,
		Memory:          2,
		Commitment:      examples.OptionsString("server_plan_commitment"),
		Generation:      examples.OptionsString("server_plan_generation"),
		InterfaceDriver: examples.OptionsString("interface_driver"),
		BootAfterCreate: true,
		CDROMID:         examples.ID,
		PrivateHostID:   examples.ID,
		NetworkInterfaces: []*server.NetworkInterface{
			{
				Upstream:       "shared | disconnected | (switch-id)",
				PacketFilterID: examples.ID,
				UserIPAddress:  examples.IPAddress,
			},
		},
		Disks: []*diskApplyParameter{
			{
				DescParameter: cflag.DescParameter{
					Description: "新規ディスクを作成する例",
				},
				TagsParameter:   examples.Tags,
				IconIDParameter: examples.IconID,
				DiskPlan:        examples.OptionsString("disk_plan"),
				Connection:      examples.OptionsString("disk_connection"),
				SourceDiskID:    examples.ID,
				SourceArchiveID: examples.ID,
				SizeGB:          20,
				DistantFrom:     []types.ID{examples.ID},
				OSType:          examples.OptionsString("os_type"),
				EditDisk: common.EditRequest{
					HostName:            "hostname",
					Password:            "password",
					IPAddress:           examples.IPAddress,
					NetworkMaskLen:      examples.NetworkMaskLen,
					DefaultRoute:        examples.DefaultRoute,
					DisablePWAuth:       true,
					EnableDHCP:          true,
					ChangePartitionUUID: true,
					SSHKeys:             []string{"/path/to/your/public/key", "ssh-rsa ..."},
					SSHKeyIDs:           []types.ID{examples.ID},
					IsSSHKeysEphemeral:  true,
					NoteIDs:             []types.ID{examples.ID},
					IsNotesEphemeral:    true,
					Notes: []*iaas.DiskEditNote{
						{
							ID: examples.ID,
							Variables: map[string]interface{}{
								"variable1": "foo",
								"variable2": "bar",
							},
						},
					},
				},
				NoWait: true,
			},
			{
				ID: examples.ID,
				DescParameter: cflag.DescParameter{
					Description: "既存のディスクを接続する例",
				},
			},
		},
		NoWaitParameter: cflag.NoWaitParameter{
			NoWait: false,
		},
	}
}
