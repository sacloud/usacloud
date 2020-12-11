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

package gslb

import (
	"github.com/sacloud/libsacloud/v2/sacloud"
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
	cflag.InputParameter   `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag.NameParameter   `cli:",squash" mapconv:",squash"`
	cflag.DescParameter   `cli:",squash" mapconv:",squash"`
	cflag.TagsParameter   `cli:",squash" mapconv:",squash"`
	cflag.IconIDParameter `cli:",squash" mapconv:",squash"`

	HealthCheck struct {
		Protocol     string `validate:"required,gslb_protocol"`
		HostHeader   string
		Path         string
		ResponseCode int `cli:"status,aliases=response-code"`
		Port         int `validate:"omitempty,min=1,max=65535"`
	} `cli:",category=health"`

	DelayLoop int  `cli:",category=health,order=10" validate:"required,min=10,max=60"`
	Weighted  bool `cli:",category=health,order=20"`

	SorryServer string `validate:"omitempty,ipv4"`

	ServersData        string              `cli:"servers" mapconv:"-"`
	DestinationServers sacloud.GSLBServers `cli:"-"`
}

func newCreateParameter() *createParameter {
	return &createParameter{
		DelayLoop: 10,
	}
}

func init() {
	Resource.AddCommand(createCommand)
}

// Customize パラメータ変換処理
func (p *createParameter) Customize(_ cli.Context) error {
	if p.ServersData != "" {
		var servers sacloud.GSLBServers
		if err := util.MarshalJSONFromPathOrContent(p.ServersData, &servers); err != nil {
			return err
		}
		p.DestinationServers = append(p.DestinationServers, servers...)
	}

	return nil
}
