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
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/cmd/examples"
	"github.com/sacloud/usacloud/pkg/util"
	"github.com/sacloud/usacloud/pkg/validate"
)

var createCommand = &core.Command{
	Name:     "create",
	Category: "basic",
	Order:    20,

	ColumnDefs: defaultColumnDefs,

	ParameterInitializer: func() interface{} {
		return newCreateParameter()
	},
	ValidateFunc: validateCreateParameter,
}

type createParameter struct {
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag.NameParameter   `cli:",squash" mapconv:",squash"`
	cflag.DescParameter   `cli:",squash" mapconv:",squash"`
	cflag.TagsParameter   `cli:",squash" mapconv:",squash"`
	cflag.IconIDParameter `cli:",squash" mapconv:",squash"`

	Plan           string `cli:",options=proxylb_plan" mapconv:",filters=proxylb_plan_to_value" validate:"required,proxylb_plan"`
	HealthCheck    createParameterHealthCheck
	SorryServer    createParameterSorryServer `mapconv:",omitempty"`
	LetsEncrypt    createParameterLetsEncrypt
	StickySession  createParameterStickySession `mapconv:",omitempty"`
	Gzip           createParameterGzip          `mapconv:",omitempty"`
	ProxyProtocol  createParameterProxyProtocol `mapconv:",omitempty"`
	Syslog         createParameterSyslog        `mapconv:",omitempty"`
	Timeout        createParameterTimeout       `cli:",squash"`
	UseVIPFailover bool                         `cli:"vip-fail-over"`
	Region         string                       `cli:",options=proxylb_region" validate:"required,proxylb_region"`

	BindPortsData string                  `cli:"bind-ports" mapconv:"-" json:"-"`
	BindPorts     []*iaas.ProxyLBBindPort `cli:"-"`

	ServersData string                `cli:"servers" mapconv:"-" json:"-"`
	Servers     []*iaas.ProxyLBServer `cli:"-"`

	RulesData string              `cli:"rules" mapconv:"-" json:"-"`
	Rules     []*iaas.ProxyLBRule `cli:"-"`
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

type createParameterSyslog struct {
	Server string `validate:"omitempty,ipv4"`
	Port   int    `validate:"omitempty,min=0,max=65535"`
}

type createParameterLetsEncrypt struct {
	CommonName      string `validate:"omitempty,fqdn"`
	Enabled         bool
	SubjectAltNames []string `validate:"omitempty,dive,fqdn"`
	AcceptTOS       bool     `cli:"accept-tos,desc=The flag to accept the current Let's Encrypt terms of service(see: https://letsencrypt.org/repository/)" mapconv:"-"`
}

type createParameterStickySession struct {
	Method  string
	Enabled bool
}

type createParameterGzip struct {
	Enabled bool
}

type createParameterProxyProtocol struct {
	Enabled bool
}

type createParameterTimeout struct {
	InactiveSec int `validate:"omitempty,min=10,max=600"`
}

func newCreateParameter() *createParameter {
	return &createParameter{
		Plan:   types.ProxyLBPlans.CPS100.String(),
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
		var bindPorts []*iaas.ProxyLBBindPort
		if err := util.MarshalJSONFromPathOrContent(p.BindPortsData, &bindPorts); err != nil {
			return err
		}
		p.BindPorts = append(p.BindPorts, bindPorts...)
	}

	if p.ServersData != "" {
		var servers []*iaas.ProxyLBServer
		if err := util.MarshalJSONFromPathOrContent(p.ServersData, &servers); err != nil {
			return err
		}
		p.Servers = append(p.Servers, servers...)
	}

	if p.RulesData != "" {
		var rules []*iaas.ProxyLBRule
		if err := util.MarshalJSONFromPathOrContent(p.RulesData, &rules); err != nil {
			return err
		}
		p.Rules = append(p.Rules, rules...)
	}
	return nil
}

func (p *createParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &createParameter{
		NameParameter:   examples.Name,
		DescParameter:   examples.Description,
		TagsParameter:   examples.Tags,
		IconIDParameter: examples.IconID,
		Plan:            examples.OptionsString("proxylb_plan"),
		HealthCheck: createParameterHealthCheck{
			Protocol:  examples.OptionsString("proxylb_protocol"),
			Path:      "/healthz",
			Host:      "www.example.com",
			DelayLoop: 10,
		},
		SorryServer: createParameterSorryServer{
			IPAddress: examples.IPAddress,
			Port:      80,
		},
		LetsEncrypt: createParameterLetsEncrypt{
			CommonName:      "www.example.com",
			SubjectAltNames: []string{"www1.example.com"},
			Enabled:         true,
		},
		StickySession: createParameterStickySession{
			Method:  "cookie",
			Enabled: true,
		},
		Gzip: createParameterGzip{
			Enabled: true,
		},
		ProxyProtocol: createParameterProxyProtocol{
			Enabled: true,
		},
		Syslog: createParameterSyslog{
			Server: examples.IPAddress,
			Port:   514,
		},
		Timeout: createParameterTimeout{
			InactiveSec: 10,
		},
		UseVIPFailover: true,
		Region:         examples.OptionsString("proxylb_region"),
		BindPorts: []*iaas.ProxyLBBindPort{
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
		Servers: []*iaas.ProxyLBServer{
			{
				IPAddress:   examples.IPAddress,
				Port:        80,
				ServerGroup: "group1",
				Enabled:     true,
			},
		},
		Rules: []*iaas.ProxyLBRule{
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

func validateCreateParameter(_ cli.Context, parameter interface{}) error {
	if err := validate.Exec(parameter); err != nil {
		return err
	}
	p := parameter.(*createParameter)
	if p.LetsEncrypt.Enabled && !p.LetsEncrypt.AcceptTOS {
		return validate.NewValidationError(
			validate.NewFlagError("--lets-encrypt-accept-tos", "required when --lets-encrypt-enabled=true"),
		)
	}
	return nil
}
