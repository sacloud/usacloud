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

package proxylb

import (
	"net/http"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/iaas-api-go/types"
	"github.com/sacloud/packages-go/pointer"
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/examples"
	"github.com/sacloud/usacloud/pkg/util"
	"github.com/sacloud/usacloud/pkg/validate"
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
	ValidateFunc: validateUpdateParameter,
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

	Plan *string `cli:",options=proxylb_plan" mapconv:",omitempty,filters=proxylb_plan_to_value" validate:"omitempty,proxylb_plan"`

	HealthCheck   updateParameterHealthCheck   `mapconv:",omitempty"`
	SorryServer   updateParameterSorryServer   `mapconv:",omitempty"`
	LetsEncrypt   updateParameterLetsEncrypt   `mapconv:",omitempty"`
	StickySession updateParameterStickySession `mapconv:",omitempty"`
	Gzip          updateParameterGzip          `mapconv:",omitempty"`
	ProxyProtocol updateParameterProxyProtocol `mapconv:",omitempty"`
	Syslog        updateParameterSyslog        `mapconv:",omitempty"`
	Timeout       updateParameterTimeout       `cli:",squash"`

	BindPortsData *string                  `cli:"bind-ports" mapconv:"-"`
	BindPorts     *[]*iaas.ProxyLBBindPort `cli:"-"`

	ServersData *string                `cli:"servers" mapconv:"-"`
	Servers     *[]*iaas.ProxyLBServer `cli:"-"`

	RulesData *string              `cli:"rules" mapconv:"-"`
	Rules     *[]*iaas.ProxyLBRule `cli:"-"`
}

func newUpdateParameter() *updateParameter {
	return &updateParameter{}
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
	CommonName      *string   `validate:"omitempty,fqdn"`
	SubjectAltNames *[]string `validate:"omitempty,dive,fqdn"`
	Enabled         *bool
	AcceptTOS       bool `cli:"accept-tos,desc=The flag to accept the current Let's Encrypt terms of service(see: https://letsencrypt.org/repository/)" mapconv:"-"`
}

type updateParameterStickySession struct {
	Method  *string
	Enabled *bool
}

type updateParameterGzip struct {
	Enabled *bool
}

type updateParameterProxyProtocol struct {
	Enabled *bool
}

type updateParameterSyslog struct {
	Server *string `validate:"omitempty,ipv4"`
	Port   *int    `validate:"omitempty,min=0,max=65535"`
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
		var bindPorts []*iaas.ProxyLBBindPort
		if err := util.MarshalJSONFromPathOrContent(*p.BindPortsData, &bindPorts); err != nil {
			return err
		}
		if p.BindPorts == nil {
			p.BindPorts = &[]*iaas.ProxyLBBindPort{}
		}
		*p.BindPorts = append(*p.BindPorts, bindPorts...)
	}
	if p.ServersData != nil && *p.ServersData != "" {
		var servers []*iaas.ProxyLBServer
		if err := util.MarshalJSONFromPathOrContent(*p.ServersData, &servers); err != nil {
			return err
		}
		if p.Servers == nil {
			p.Servers = &[]*iaas.ProxyLBServer{}
		}
		*p.Servers = append(*p.Servers, servers...)
	}
	if p.RulesData != nil && *p.RulesData != "" {
		var rules []*iaas.ProxyLBRule
		if err := util.MarshalJSONFromPathOrContent(*p.RulesData, &rules); err != nil {
			return err
		}
		if p.Rules == nil {
			p.Rules = &[]*iaas.ProxyLBRule{}
		}
		*p.Rules = append(*p.Rules, rules...)
	}
	return nil
}

func (p *updateParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &updateParameter{
		NameUpdateParameter:   examples.NameUpdate,
		DescUpdateParameter:   examples.DescriptionUpdate,
		TagsUpdateParameter:   examples.TagsUpdate,
		IconIDUpdateParameter: examples.IconIDUpdate,
		Plan:                  pointer.NewString(examples.OptionsString("proxylb_plan")),
		HealthCheck: updateParameterHealthCheck{
			Protocol:  pointer.NewString(examples.OptionsString("proxylb_protocol")),
			Path:      pointer.NewString("/healthz"),
			Host:      pointer.NewString("www.example.com"),
			DelayLoop: pointer.NewInt(10),
		},
		SorryServer: updateParameterSorryServer{
			IPAddress: &examples.IPAddress,
			Port:      pointer.NewInt(80),
		},
		LetsEncrypt: updateParameterLetsEncrypt{
			CommonName: pointer.NewString("www.example.com"),
			Enabled:    pointer.NewBool(true),
		},
		StickySession: updateParameterStickySession{
			Method:  pointer.NewString("cookie"),
			Enabled: pointer.NewBool(true),
		},
		Gzip: updateParameterGzip{
			Enabled: pointer.NewBool(true),
		},
		ProxyProtocol: updateParameterProxyProtocol{
			Enabled: pointer.NewBool(true),
		},
		Syslog: updateParameterSyslog{
			Server: &examples.IPAddress,
			Port:   pointer.NewInt(514),
		},
		Timeout: updateParameterTimeout{
			InactiveSec: pointer.NewInt(10),
		},
		BindPorts: &[]*iaas.ProxyLBBindPort{
			{
				ProxyMode:       types.EProxyLBProxyMode(examples.OptionsString("proxylb_proxy_mode")),
				Port:            80,
				RedirectToHTTPS: true,
				SupportHTTP2:    true,
				AddResponseHeader: []*iaas.ProxyLBResponseHeader{
					{
						Header: "Cache-Control",
						Value:  "public, max-age=900",
					},
				},
				SSLPolicy: examples.OptionsString("proxylb_ssl_policy"),
			},
		},
		Servers: &[]*iaas.ProxyLBServer{
			{
				IPAddress:   examples.IPAddress,
				Port:        80,
				ServerGroup: "group1",
				Enabled:     true,
			},
		},
		Rules: &[]*iaas.ProxyLBRule{
			{
				Action:      types.ProxyLBRuleActions.Forward,
				Host:        "www2.example.com",
				Path:        "/foo1",
				ServerGroup: "group1",
			},
			{
				Action:             types.ProxyLBRuleActions.Redirect,
				Host:               "www2.example.com",
				Path:               "/foo2",
				ServerGroup:        "group1",
				RedirectLocation:   "/redirect",
				RedirectStatusCode: http.StatusMovedPermanently,
			},
			{
				Action:           types.ProxyLBRuleActions.Fixed,
				Host:             "www2.example.com",
				Path:             "/foo3",
				ServerGroup:      "group1",
				FixedStatusCode:  http.StatusOK,
				FixedContentType: types.ProxyLBFixedContentTypes.Plain,
				FixedMessageBody: "your-content",
			},
		},
	}
}

func validateUpdateParameter(_ cli.Context, parameter interface{}) error {
	if err := validate.Exec(parameter); err != nil {
		return err
	}
	p := parameter.(*updateParameter)
	if p.LetsEncrypt.Enabled != nil && *p.LetsEncrypt.Enabled && !p.LetsEncrypt.AcceptTOS {
		return validate.NewValidationError(
			validate.NewFlagError("--lets-encrypt-accept-tos", "required when --lets-encrypt-enabled=true"),
		)
	}
	return nil
}
