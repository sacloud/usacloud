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

package localrouter

import (
	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/packages-go/pointer"
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
	cflag.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag.NameUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.DescUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.TagsUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.IconIDUpdateParameter `cli:",squash" mapconv:",omitempty,squash"`

	Switch switchParameterUpdate

	Interface interfaceParameterUpdate `cli:",squash"`

	PeersData *string                  `cli:"peers" mapconv:"-" json:"-"`
	Peers     *[]*iaas.LocalRouterPeer `cli:"-"`

	StaticRoutesData *string                         `cli:"static-routes" mapconv:"-" json:"-"`
	StaticRoutes     *[]*iaas.LocalRouterStaticRoute `cli:"-"`
}

type switchParameterUpdate struct {
	Code     *string
	Category *string
	ZoneID   *string
}

type interfaceParameterUpdate struct {
	VirtualIPAddress *string   `validate:"omitempty,ipv4"`
	IPAddress        *[]string `cli:"ip-addresses" validate:"omitempty,min=2,max=2,dive,ipv4"`
	NetworkMaskLen   *int      `cli:"netmask,aliases=network-mask-len" validate:"omitempty,min=8,max=28"`
	VRID             *int      `cli:"vrid" validate:"omitempty"`
}

func newUpdateParameter() *updateParameter {
	return &updateParameter{}
}

func init() {
	Resource.AddCommand(updateCommand)
}

// Customize パラメータ変換処理
func (p *updateParameter) Customize(_ cli.Context) error {
	if p.PeersData != nil && *p.PeersData != "" {
		var peers []*iaas.LocalRouterPeer
		if err := util.MarshalJSONFromPathOrContent(*p.PeersData, &peers); err != nil {
			return err
		}
		if p.Peers == nil {
			p.Peers = &[]*iaas.LocalRouterPeer{}
		}
		*p.Peers = append(*p.Peers, peers...)
	}
	if p.StaticRoutesData != nil && *p.StaticRoutesData != "" {
		var staticRoutes []*iaas.LocalRouterStaticRoute
		if err := util.MarshalJSONFromPathOrContent(*p.StaticRoutesData, &staticRoutes); err != nil {
			return err
		}
		if p.StaticRoutes == nil {
			p.StaticRoutes = &[]*iaas.LocalRouterStaticRoute{}
		}
		*p.StaticRoutes = append(*p.StaticRoutes, staticRoutes...)
	}
	return nil
}

func (p *updateParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &updateParameter{
		NameUpdateParameter:   examples.NameUpdate,
		DescUpdateParameter:   examples.DescriptionUpdate,
		TagsUpdateParameter:   examples.TagsUpdate,
		IconIDUpdateParameter: examples.IconIDUpdate,
		Switch: switchParameterUpdate{
			Category: pointer.NewString("cloud"),
			Code:     pointer.NewString(examples.ID.String()),
			ZoneID:   pointer.NewString(examples.ZonesString(ctx.Option().Zones)),
		},
		Interface: interfaceParameterUpdate{
			VirtualIPAddress: pointer.NewString(examples.VirtualIPAddress),
			IPAddress:        pointer.NewStringSlice(examples.IPAddresses),
			NetworkMaskLen:   pointer.NewInt(examples.NetworkMaskLen),
			VRID:             pointer.NewInt(1),
		},
		Peers: &[]*iaas.LocalRouterPeer{
			{
				ID:          examples.ID,
				SecretKey:   "*****",
				Enabled:     true,
				Description: "example-peer",
			},
		},
		StaticRoutes: &[]*iaas.LocalRouterStaticRoute{
			{
				Prefix:  "192.0.2.0/24",
				NextHop: "192.0.2.1",
			},
		},
	}
}
