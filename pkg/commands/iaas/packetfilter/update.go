// Copyright 2017-2022 The sacloud/usacloud Authors
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

package packetfilter

import (
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

	cflag.NameUpdateParameter `cli:",squash" mapconv:",omitempty,squash"`
	cflag.DescUpdateParameter `cli:",squash" mapconv:",omitempty,squash"`
	ExpressionsData           *string                         `cli:"expressions,aliases=rules" mapconv:"-" json:"-"`
	Expressions               *[]*iaas.PacketFilterExpression `cli:"-" mapconv:"Expression"`
}

func newUpdateParameter() *updateParameter {
	return &updateParameter{}
}

func init() {
	Resource.AddCommand(updateCommand)
}

// Customize パラメータ変換処理
func (p *updateParameter) Customize(_ cli.Context) error {
	if p.ExpressionsData != nil && *p.ExpressionsData != "" {
		var expressions []*iaas.PacketFilterExpression
		if err := util.MarshalJSONFromPathOrContent(*p.ExpressionsData, &expressions); err != nil {
			return err
		}
		if p.Expressions == nil {
			p.Expressions = &[]*iaas.PacketFilterExpression{}
		}
		*p.Expressions = append(*p.Expressions, expressions...)
	}

	return nil
}

func (p *updateParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &updateParameter{
		ZoneParameter:       examples.Zones(ctx.Option().Zones),
		NameUpdateParameter: examples.NameUpdate,
		DescUpdateParameter: examples.DescriptionUpdate,
		Expressions: &[]*iaas.PacketFilterExpression{
			{
				Protocol:        types.Protocol(examples.OptionsString("packetfilter_protocol")),
				SourceNetwork:   "192.0.2.1 | 192.0.2.0/24",
				SourcePort:      "1024 | 1024-2048",
				DestinationPort: "1024 | 1024-2048",
				Action:          types.Action(examples.OptionsString("packetfilter_action")),
				Description:     "description",
			},
			{
				Protocol:        types.Protocols.TCP,
				DestinationPort: "22",
				Action:          types.Actions.Allow,
				Description:     "allow ssh",
			},
			{
				Protocol: types.Protocols.ICMP,
				Action:   types.Actions.Allow,
			},
			{
				Protocol:        types.Protocols.TCP,
				DestinationPort: "32768-61000",
				Action:          types.Actions.Allow,
			},
			{
				Protocol:        types.Protocols.UDP,
				DestinationPort: "32768-61000",
				Action:          types.Actions.Allow,
			},
			{
				Protocol: types.Protocols.Fragment,
				Action:   types.Actions.Allow,
			},
			{
				Protocol: types.Protocols.IP,
				Action:   types.Actions.Deny,
			},
		},
	}
}
