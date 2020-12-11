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

	VirtualIPAddressesData *string                                 `cli:"virtual-ip-addresses,category=network" mapconv:"-"`
	VirtualIPAddresses     *sacloud.LoadBalancerVirtualIPAddresses `cli:"-"`
	cflag.NoWaitParameter  `cli:",squash" mapconv:",squash"`
}

func newUpdateParameter() *updateParameter {
	return &updateParameter{}
}

func init() {
	Resource.AddCommand(updateCommand)
}

// Customize パラメータ変換処理
func (p *updateParameter) Customize(_ cli.Context) error {
	if p.VirtualIPAddressesData != nil && *p.VirtualIPAddressesData != "" {
		var vips sacloud.LoadBalancerVirtualIPAddresses
		if err := util.MarshalJSONFromPathOrContent(*p.VirtualIPAddressesData, &vips); err != nil {
			return err
		}
		if p.VirtualIPAddresses == nil {
			p.VirtualIPAddresses = &sacloud.LoadBalancerVirtualIPAddresses{}
		}
		*p.VirtualIPAddresses = append(*p.VirtualIPAddresses, vips...)
	}

	return nil
}
