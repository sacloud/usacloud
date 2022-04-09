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

package gslb

import (
	"net/http"

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

	HealthCheck gslbHealthCheckUpdate `cli:",category=health" mapconv:",omitempty"`

	DelayLoop   *int    `cli:",category=health,order=10" validate:"omitempty,min=10,max=60"`
	Weighted    *bool   `cli:",category=health,order=20"`
	SorryServer *string `validate:"omitempty,ipv4"`

	ServersData        *string           `cli:"servers" mapconv:"-" json:"-"`
	DestinationServers *iaas.GSLBServers `cli:"-"`
}

type gslbHealthCheckUpdate struct {
	Protocol     *string `validate:"omitempty,gslb_protocol"`
	HostHeader   *string
	Path         *string
	ResponseCode *int `cli:"status,aliases=response-code"`
	Port         *int `validate:"omitempty,min=1,max=65535"`
}

func newUpdateParameter() *updateParameter {
	return &updateParameter{}
}

func init() {
	Resource.AddCommand(updateCommand)
}

// Customize パラメータ変換処理
func (p *updateParameter) Customize(_ cli.Context) error {
	if p.ServersData != nil && *p.ServersData != "" {
		var servers iaas.GSLBServers
		if err := util.MarshalJSONFromPathOrContent(*p.ServersData, &servers); err != nil {
			return err
		}
		if p.DestinationServers == nil {
			p.DestinationServers = &iaas.GSLBServers{}
		}
		*p.DestinationServers = append(*p.DestinationServers, servers...)
	}

	return nil
}

func (p *updateParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &updateParameter{
		NameUpdateParameter:   examples.NameUpdate,
		DescUpdateParameter:   examples.DescriptionUpdate,
		TagsUpdateParameter:   examples.TagsUpdate,
		IconIDUpdateParameter: examples.IconIDUpdate,
		HealthCheck: gslbHealthCheckUpdate{
			Protocol:     pointer.NewString(examples.OptionsString("gslb_protocol")),
			HostHeader:   pointer.NewString("www.example.com"),
			Path:         pointer.NewString("/"),
			ResponseCode: pointer.NewInt(http.StatusOK),
			Port:         pointer.NewInt(80),
		},
		DelayLoop:   pointer.NewInt(10),
		Weighted:    pointer.NewBool(true),
		SorryServer: pointer.NewString("192.0.2.1"),
		DestinationServers: &iaas.GSLBServers{
			{
				IPAddress: examples.IPAddress,
				Enabled:   true,
				Weight:    1,
			},
		},
	}
}
