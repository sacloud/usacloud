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
	"github.com/sacloud/libsacloud/v2/helper/service/server"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/util"
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
	cflag.InputParameter   `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag.NameUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.DescUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.TagsUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.IconIDUpdateParameter `cli:",squash" mapconv:",omitempty,squash"`

	CPU        *int    `cli:"cpu,aliases=core,category=plan,order=10"`
	Memory     *int    `cli:"memory,category=plan,order=20" mapconv:"MemoryGB"`
	Commitment *string `cli:",options=server_plan_commitment,category=plan,order=30" mapconv:",omitempty,filters=server_plan_commitment_to_value" validate:"omitempty,server_plan_commitment"`
	Generation *string `cli:",options=server_plan_generation,category=plan,order=40" mapconv:",omitempty,filters=server_plan_generation_to_value" validate:"omitempty,server_plan_generation"`

	InterfaceDriver *string `cli:",options=interface_dirver" mapconv:",omitempty,filters=interface_driver_to_value" validate:"omitempty,interface_driver"`

	CDROMID       *types.ID `cli:"cdrom-id,aliases=iso-image-id"`
	PrivateHostID *types.ID

	NetworkInterfaceData string                      `cli:"network-interfaces" mapconv:"-"`
	NetworkInterfaces    *[]*server.NetworkInterface `cli:"-" mapconv:",omitempty,recursive"`

	DisksData string                 `cli:"disks" mapconv:"-"`
	Disks     *[]*diskApplyParameter `cli:"-" mapconv:",omitempty,recursive"`

	cflag.NoWaitParameter `cli:",squash" mapconv:",squash"`
	ForceShutdown         bool // DeleteのForceと区別するために-fは定義しない
}

func newUpdateParameter() *updateParameter {
	return &updateParameter{}
}

func init() {
	Resource.AddCommand(updateCommand)
}

// Customize パラメータ変換処理
func (p *updateParameter) Customize(_ cli.Context) error {
	if p.NetworkInterfaceData != "" {
		var nics []*server.NetworkInterface
		if err := util.MarshalJSONFromPathOrContent(p.NetworkInterfaceData, &nics); err != nil {
			return err
		}
		if p.NetworkInterfaces == nil {
			p.NetworkInterfaces = &[]*server.NetworkInterface{}
		}
		*p.NetworkInterfaces = append(*p.NetworkInterfaces, nics...)
	}

	if p.DisksData != "" {
		var disks []*diskApplyParameter
		if err := util.MarshalJSONFromPathOrContent(p.DisksData, &disks); err != nil {
			return err
		}
		if p.Disks == nil {
			p.Disks = &[]*diskApplyParameter{}
		}
		*p.Disks = append(*p.Disks, disks...)
	}
	return nil
}
