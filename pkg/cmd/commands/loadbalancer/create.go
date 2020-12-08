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

	SwitchID       types.ID `validate:"required"`
	PlanID         string   `cli:"plan,options=loadbalancer_plan" mapconv:",filters=loadbalancer_plan_to_value" validate:"required,loadbalancer_plan"`
	VRID           int      `cli:"vrid"`
	IPAddresses    []string `cli:"ip-address,aliases=ipaddress" validate:"required,min=1,max=2,dive,ipv4"`
	NetworkMaskLen int      `validate:"required,min=1,max=32"`
	DefaultRoute   string   `validate:"omitempty,ipv4"`
	Port           int      `validate:"omitempty,min=1,max=65535"`

	VirtualIPAddressesData string                                 `cli:"virtual-ip-addresses" mapconv:"-"`
	VirtualIPAddresses     sacloud.LoadBalancerVirtualIPAddresses `cli:"-"`

	NoWait bool
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
