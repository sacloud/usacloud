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

package loadbalancer

import (
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/util"
)

var createCommand = &core.Command{
	Name:     "create",
	Category: "basic",
	Order:    20,

	ColumnDefs: defaultColumnDefs,

	ParameterInitializer: func() interface{} {
		return newCreateParameter()
	},

	// TODO プランごとにIPアドレス数のバリデーションが必要
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

	PlanID string `cli:"plan,options=loadbalancer_plan,category=plan,order=10" mapconv:",filters=loadbalancer_plan_to_value" validate:"required,loadbalancer_plan"`

	VRID           int      `cli:"vrid,category=network,order=10"`
	SwitchID       types.ID `cli:",category=network,order=20" validate:"required"`
	IPAddresses    []string `cli:"ip-address,aliases=ipaddress,category=network,order=30" validate:"required,min=1,max=2,dive,ipv4"`
	NetworkMaskLen int      `cli:",category=network,order=40" validate:"required,min=1,max=32"`
	DefaultRoute   string   `cli:",category=network,order=50" validate:"omitempty,ipv4"`
	Port           int      `cli:",category=network,order=60" validate:"omitempty,min=1,max=65535"`

	VirtualIPAddressesData string                                 `cli:"virtual-ip-addresses,category=network,order=70" mapconv:"-"`
	VirtualIPAddresses     sacloud.LoadBalancerVirtualIPAddresses `cli:"-"`

	cflag.NoWaitParameter `cli:",squash" mapconv:",squash"`
}

func newCreateParameter() *createParameter {
	return &createParameter{
		PlanID: "standard",
	}
}

func init() {
	Resource.AddCommand(createCommand)
}

// Customize パラメータ変換処理
func (p *createParameter) Customize(_ cli.Context) error {
	if p.VirtualIPAddressesData != "" {
		var vips sacloud.LoadBalancerVirtualIPAddresses
		if err := util.MarshalJSONFromPathOrContent(p.VirtualIPAddressesData, &vips); err != nil {
			return err
		}
		p.VirtualIPAddresses = append(p.VirtualIPAddresses, vips...)
	}

	return nil
}
