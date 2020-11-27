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

package proxylb

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
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	Name        string   `validate:"required"`
	Description string   `validate:"description"`
	Tags        []string `validate:"tags"`
	IconID      types.ID

	Plan           int `validate:"required,proxylb_plan"`
	HealthCheck    createParameterHealthCheck
	SorryServer    createParameterSorryServer `mapconv:",omitempty"`
	LetsEncrypt    createParameterLetsEncrypt
	StickySession  createParameterStickySession `mapconv:",omitempty"`
	Timeout        createParameterTimeout       `cli:",squash"`
	UseVIPFailover bool                         `cli:"vip-fail-over"`
	Region         string                       `cli:",options=proxylb_region" validate:"required,proxylb_region"`

	BindPortsData string                     `cli:"bind-ports" mapconv:"-"`
	BindPorts     []*sacloud.ProxyLBBindPort `cli:"-"`

	ServersData string                   `cli:"servers" mapconv:"-"`
	Servers     []*sacloud.ProxyLBServer `cli:"-"`

	RulesData string                 `cli:"rules" mapconv:"-"`
	Rules     []*sacloud.ProxyLBRule `cli:"-"`
}

type createParameterHealthCheck struct {
	Protocol  string `validate:"required,proxylb_protocol"`
	Path      string
	Host      string
	DelayLoop int `validate:"required,min=10,max=60"`
}

type createParameterSorryServer struct {
	IPAddress string `cli:",aliases=ipaddress" validate:"omitempty,ipv4"`
	Port      int    `validate:"omitempty,min=0,max=65535"`
}

type createParameterLetsEncrypt struct {
	CommonName string `validate:"omitempty,fqdn"`
	Enabled    bool
}

type createParameterStickySession struct {
	Method  string
	Enabled bool
}

type createParameterTimeout struct {
	InactiveSec int `validate:"omitempty,min=10,max=600"`
}

func newCreateParameter() *createParameter {
	return &createParameter{
		Plan:   types.ProxyLBPlans.CPS100.Int(),
		Region: types.ProxyLBRegions.IS1.String(),
		HealthCheck: createParameterHealthCheck{
			Protocol:  types.ProxyLBProtocols.HTTP.String(),
			Path:      "/",
			Host:      "",
			DelayLoop: 10,
		},
		Timeout: createParameterTimeout{InactiveSec: 10},
	}
}

func init() {
	Resource.AddCommand(createCommand)
}

// Customize パラメータ変換処理
func (p *createParameter) Customize(_ cli.Context) error {
	if p.BindPortsData != "" {
		var bindPorts []*sacloud.ProxyLBBindPort
		if err := util.MarshalJSONFromPathOrContent(p.BindPortsData, &bindPorts); err != nil {
			return err
		}
		p.BindPorts = append(p.BindPorts, bindPorts...)
	}

	if p.ServersData != "" {
		var servers []*sacloud.ProxyLBServer
		if err := util.MarshalJSONFromPathOrContent(p.ServersData, &servers); err != nil {
			return err
		}
		p.Servers = append(p.Servers, servers...)
	}

	if p.RulesData != "" {
		var rules []*sacloud.ProxyLBRule
		if err := util.MarshalJSONFromPathOrContent(p.RulesData, &rules); err != nil {
			return err
		}
		p.Rules = append(p.Rules, rules...)
	}
	return nil
}
