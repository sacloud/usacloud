// Copyright 2017-2025 The sacloud/usacloud Authors
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
	"net/http"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/examples"
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
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag.NameUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.DescUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.TagsUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.IconIDUpdateParameter `cli:",squash" mapconv:",omitempty,squash"`

	VirtualIPAddressesData *string                              `cli:"virtual-ip-addresses,aliases=vips,category=network" mapconv:"-" json:"-"`
	VirtualIPAddresses     *iaas.LoadBalancerVirtualIPAddresses `cli:"-"`
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
		var vips iaas.LoadBalancerVirtualIPAddresses
		if err := util.MarshalJSONFromPathOrContent(*p.VirtualIPAddressesData, &vips); err != nil {
			return err
		}
		if p.VirtualIPAddresses == nil {
			p.VirtualIPAddresses = &iaas.LoadBalancerVirtualIPAddresses{}
		}
		*p.VirtualIPAddresses = append(*p.VirtualIPAddresses, vips...)
	}

	return nil
}

func (p *updateParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &updateParameter{
		ZoneParameter:         examples.Zones(ctx.Option().Zones),
		NameUpdateParameter:   examples.NameUpdate,
		DescUpdateParameter:   examples.DescriptionUpdate,
		TagsUpdateParameter:   examples.TagsUpdate,
		IconIDUpdateParameter: examples.IconIDUpdate,
		VirtualIPAddresses: &iaas.LoadBalancerVirtualIPAddresses{
			{
				VirtualIPAddress: examples.VirtualIPAddress,
				Port:             80,
				DelayLoop:        10,
				SorryServer:      "192.0.2.1",
				Description:      "example",
				Servers: iaas.LoadBalancerServers{
					{
						IPAddress: "192.0.2.101",
						Port:      80,
						Enabled:   true,
						HealthCheck: &iaas.LoadBalancerServerHealthCheck{
							Protocol:     types.ELoadBalancerHealthCheckProtocol(examples.OptionsString("loadbalancer_server_protocol")),
							Path:         "/",
							ResponseCode: http.StatusOK,
						},
					},
				},
			},
		},
	}
}
