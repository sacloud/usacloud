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

	Name        *string   `validate:"omitempty,min=1"`
	Description *string   `validate:"omitempty,description"`
	Tags        *[]string `validate:"omitempty,tags"`
	IconID      *types.ID

	Plan *int `validate:"omitempty,proxylb_plan"`

	HealthCheck   updateParameterHealthCheck   `mapconv:",omitempty"`
	SorryServer   updateParameterSorryServer   `mapconv:",omitempty"`
	LetsEncrypt   updateParameterLetsEncrypt   `mapconv:",omitempty"`
	StickySession updateParameterStickySession `mapconv:",omitempty"`
	Timeout       updateParameterTimeout       `cli:",squash"`

	BindPortsData *string                     `cli:"bind-ports" mapconv:"-"`
	BindPorts     *[]*sacloud.ProxyLBBindPort `cli:"-"`

	ServersData *string                   `cli:"servers" mapconv:"-"`
	Servers     *[]*sacloud.ProxyLBServer `cli:"-"`

	RulesData    *string                 `cli:"rules" mapconv:"-"`
	Rules        *[]*sacloud.ProxyLBRule `cli:"-"`
	SettingsHash *string
}

func newUpdateParameter() *updateParameter {
	return &updateParameter{
		// TODO デフォルト値はここで設定する
	}
}

type updateParameterHealthCheck struct {
	Protocol  *string `validate:"omitempty,proxylb_protocol"`
	Path      *string
	Host      *string
	DelayLoop *int `validate:"omitempty,min=10,max=60"`
}

type updateParameterSorryServer struct {
	IPAddress *string `cli:",aliases=ipaddress" validate:"omitempty,ipv4"`
	Port      *int    `validate:"omitempty,min=0,max=65535"`
}

type updateParameterLetsEncrypt struct {
	CommonName *string `validate:"omitempty,fqdn"`
	Enabled    *bool
}

type updateParameterStickySession struct {
	Method  *string
	Enabled *bool
}

type updateParameterTimeout struct {
	InactiveSec *int `validate:"omitempty,min=10,max=600"`
}

func init() {
	Resource.AddCommand(updateCommand)
}

// Customize パラメータ変換処理
func (p *updateParameter) Customize(_ cli.Context) error {
	if p.BindPortsData != nil && *p.BindPortsData != "" {
		var bindPorts []*sacloud.ProxyLBBindPort
		if err := util.MarshalJSONFromPathOrContent(*p.BindPortsData, &bindPorts); err != nil {
			return err
		}
		if p.BindPorts == nil {
			p.BindPorts = &[]*sacloud.ProxyLBBindPort{}
		}
		*p.BindPorts = append(*p.BindPorts, bindPorts...)
	}
	if p.ServersData != nil && *p.ServersData != "" {
		var servers []*sacloud.ProxyLBServer
		if err := util.MarshalJSONFromPathOrContent(*p.ServersData, &servers); err != nil {
			return err
		}
		if p.Servers == nil {
			p.Servers = &[]*sacloud.ProxyLBServer{}
		}
		*p.Servers = append(*p.Servers, servers...)
	}
	if p.RulesData != nil && *p.RulesData != "" {
		var rules []*sacloud.ProxyLBRule
		if err := util.MarshalJSONFromPathOrContent(*p.RulesData, &rules); err != nil {
			return err
		}
		if p.Rules == nil {
			p.Rules = &[]*sacloud.ProxyLBRule{}
		}
		*p.Rules = append(*p.Rules, rules...)
	}
	return nil
}
