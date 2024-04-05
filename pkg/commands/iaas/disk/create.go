// Copyright 2017-2023 The sacloud/usacloud Authors
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

package disk

import (
	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/commands/iaas/common"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/examples"
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

	cflag.NameParameter   `cli:",squash" mapconv:",squash"`
	cflag.DescParameter   `cli:",squash" mapconv:",squash"`
	cflag.TagsParameter   `cli:",squash" mapconv:",squash"`
	cflag.IconIDParameter `cli:",squash" mapconv:",squash"`
	DiskPlan              string `cli:",options=disk_plan,category=plan,order=10" mapconv:"DiskPlanID,filters=disk_plan_to_value" validate:"required,disk_plan"`
	SizeGB                int    `cli:"size,category=plan,order=20"`
	Connection            string `cli:"connector,aliases=connection,options=disk_connection,category=plan,order=30" validate:"required,disk_connection"`
	EncryptionAlgorithm   string `cli:"encryption-algorithm,options=disk_encryption_algorithm,category=plan,order=40" validate:"omitempty,disk_encryption_algorithm"`

	OSType          string   `cli:",options=os_type,display_options=os_type_simple,category=source,order=10" mapconv:",omitempty,filters=os_type_to_value" validate:"omitempty,os_type"`
	SourceDiskID    types.ID `cli:",category=source,order=20"`
	SourceArchiveID types.ID `cli:",category=source,order=30"`

	ServerID    types.ID
	DistantFrom []types.ID

	EditDisk              common.EditRequest `cli:",category=edit" mapconv:"EditParameter,omitempty"`
	cflag.NoWaitParameter `cli:",squash" mapconv:",squash"`
}

func validateCreateParameter(ctx cli.Context, parameter interface{}) error {
	if err := validate.Exec(parameter); err != nil {
		return err
	}
	p := parameter.(*createParameter)

	var errs []error
	targets := []*validate.Target{
		{FlagName: "--os-type", Value: p.OSType},
		{FlagName: "--source-archive-id", Value: p.SourceArchiveID},
		{FlagName: "--source-disk-id", Value: p.SourceDiskID},
	}
	if err := validate.ConflictWith(targets...); err != nil {
		errs = append(errs, err)
	}

	return validate.NewValidationError(errs...)
}

func newCreateParameter() *createParameter {
	return &createParameter{
		DiskPlan:   "ssd",
		Connection: "virtio",
		SizeGB:     20,
	}
}

func init() {
	Resource.AddCommand(createCommand)
}

// Customize パラメータ変換処理
func (p *createParameter) Customize(ctx cli.Context) error {
	return p.EditDisk.Customize(ctx)
}

func (p *createParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &createParameter{
		ZoneParameter:       examples.Zones(ctx.Option().Zones),
		NameParameter:       examples.Name,
		DescParameter:       examples.Description,
		TagsParameter:       examples.Tags,
		IconIDParameter:     examples.IconID,
		DiskPlan:            examples.OptionsString("disk_plan"),
		SizeGB:              20,
		Connection:          examples.OptionsString("disk_connection"),
		EncryptionAlgorithm: examples.OptionsString("disk_encryption_algorithm"),
		OSType:              examples.OptionsString("os_type"),
		SourceDiskID:        examples.ID,
		SourceArchiveID:     examples.ID,
		ServerID:            examples.ID,
		DistantFrom:         []types.ID{examples.ID},
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
		NoWaitParameter: cflag.NoWaitParameter{
			NoWait: false,
		},
	}
}
