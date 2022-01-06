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

package disk

import (
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/pointer"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/commands/common"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/cmd/examples"
)

var updateCommand = &core.Command{
	Name:         "update",
	Category:     "basic",
	Order:        40,
	SelectorType: core.SelectorTypeRequireMulti,

	ColumnDefs: defaultColumnDefs,

	ParameterInitializer: func() interface{} {
		return newUpdateParameter()
	},
}

type updateParameter struct {
	cflag.ZoneParameter    `cli:",squash" mapconv:",squash"`
	cflag.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag.NameUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.DescUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.TagsUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.IconIDUpdateParameter `cli:",squash" mapconv:",omitempty,squash"`

	Connection            *string            `cli:"connector,aliases=connection,options=disk_connection,category=plan" validate:"omitempty,disk_connection"`
	EditDisk              common.EditRequest `cli:",category=edit" mapconv:"EditParameter,omitempty"`
	cflag.NoWaitParameter `cli:",squash" mapconv:",squash"`
}

func newUpdateParameter() *updateParameter {
	return &updateParameter{}
}

func init() {
	Resource.AddCommand(updateCommand)
}

// Customize パラメータ変換処理
func (p *updateParameter) Customize(ctx cli.Context) error {
	return p.EditDisk.Customize(ctx)
}

func (p *updateParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &updateParameter{
		ZoneParameter:         examples.Zones(ctx.Option().Zones),
		NameUpdateParameter:   examples.NameUpdate,
		DescUpdateParameter:   examples.DescriptionUpdate,
		TagsUpdateParameter:   examples.TagsUpdate,
		IconIDUpdateParameter: examples.IconIDUpdate,
		Connection:            pointer.NewString(examples.OptionsString("disk_connection")),
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
			Notes: []*sacloud.DiskEditNote{
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
