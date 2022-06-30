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

	cflag.NameParameter `cli:",squash" mapconv:",squash"`
	cflag.DescParameter `cli:",squash" mapconv:",squash"`

	ExpressionsData string                         `cli:"expressions,aliases=rules" mapconv:"-" json:"-"`
	Expressions     []*iaas.PacketFilterExpression `cli:"-" mapconv:"Expression"`
}

func newCreateParameter() *createParameter {
	return &createParameter{}
}

func init() {
	Resource.AddCommand(createCommand)
}

// Customize パラメータ変換処理
func (p *createParameter) Customize(_ cli.Context) error {
	if p.ExpressionsData != "" {
		var expressions []*iaas.PacketFilterExpression
		if err := util.MarshalJSONFromPathOrContent(p.ExpressionsData, &expressions); err != nil {
			return err
		}
		p.Expressions = append(p.Expressions, expressions...)
	}

	return nil
}

func (p *createParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &createParameter{
		ZoneParameter: examples.Zones(ctx.Option().Zones),
		NameParameter: examples.Name,
		DescParameter: examples.Description,
		Expressions: []*iaas.PacketFilterExpression{
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
