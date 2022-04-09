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

package simplemonitor

import (
	"net/http"

	"github.com/sacloud/packages-go/pointer"
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/examples"
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

	cflag.DescUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.TagsUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.IconIDUpdateParameter `cli:",squash" mapconv:",omitempty,squash"`

	DelayLoop        *int `validate:"omitempty,min=60,max=3600"`
	MaxCheckAttempts *int `validate:"omitempty,min=1,max=10"`
	RetryInterval    *int `validate:"omitempty,min=10,max=3600"`
	Timeout          *int `validate:"omitempty,min=1,max=30"`
	Enabled          *bool
	HealthCheck      updateParameterHealthCheck `mapconv:",omitempty"`

	NotifyEmailEnabled *bool
	NotifyEmailHTML    *bool
	NotifySlackEnabled *bool
	SlackWebhooksURL   *string
	NotifyInterval     *int `validate:"omitempty,min=3600,max=259200"`
}

type updateParameterHealthCheck struct {
	Protocol          *string `cli:",options=simple_monitor_protocol" mapconv:",omitempty,filters=simple_monitor_protocol_to_value" validate:"omitempty,simple_monitor_protocol" json:",omitempty"`
	Port              *int    `json:",omitempty"`
	Path              *string `json:",omitempty"`
	Status            *int    `json:",omitempty"`
	ContainsString    *string `json:",omitempty"`
	SNI               *bool   `json:",omitempty"`
	Host              *string `json:",omitempty"`
	BasicAuthUsername *string `json:",omitempty"`
	BasicAuthPassword *string `json:",omitempty"`
	QName             *string `json:",omitempty"`
	ExpectedData      *string `json:",omitempty"`
	Community         *string `json:",omitempty"`
	SNMPVersion       *string `json:",omitempty"`
	OID               *string `json:",omitempty"`
	RemainingDays     *int    `json:",omitempty"`
	HTTP2             *bool   `cli:"http2" json:",omitempty"`
	FTPS              *string `cli:",options=simple_monitor_ftps" mapconv:",omitempty,filters=simple_monitor_ftps_to_value" validate:"omitempty,simple_monitor_ftps" json:",omitempty"`
	VerifySNI         *bool   `cli:"verify-sni" json:",omitempty"`
}

func newUpdateParameter() *updateParameter {
	return &updateParameter{}
}

func init() {
	Resource.AddCommand(updateCommand)
}

func (p *updateParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &updateParameter{
		DescUpdateParameter:   examples.DescriptionUpdate,
		TagsUpdateParameter:   examples.TagsUpdate,
		IconIDUpdateParameter: examples.IconIDUpdate,
		DelayLoop:             pointer.NewInt(60),
		MaxCheckAttempts:      pointer.NewInt(3),
		RetryInterval:         pointer.NewInt(10),
		Timeout:               pointer.NewInt(10),
		Enabled:               pointer.NewBool(true),
		HealthCheck: updateParameterHealthCheck{
			Protocol:          pointer.NewString(examples.OptionsString("simple_monitor_protocol")),
			Port:              pointer.NewInt(80),
			Path:              pointer.NewString("/healthz"),
			Status:            pointer.NewInt(http.StatusOK),
			ContainsString:    pointer.NewString("ok"),
			SNI:               pointer.NewBool(true),
			Host:              pointer.NewString("www2.example.com"),
			BasicAuthUsername: pointer.NewString("username"),
			BasicAuthPassword: pointer.NewString("password"),
			HTTP2:             pointer.NewBool(true),
			FTPS:              pointer.NewString(examples.OptionsString("simple_monitor_ftps")),
			VerifySNI:         pointer.NewBool(true),
		},
		NotifyEmailEnabled: pointer.NewBool(true),
		NotifyEmailHTML:    pointer.NewBool(true),
		NotifySlackEnabled: pointer.NewBool(true),
		SlackWebhooksURL:   &examples.SlackNotifyWebhooksURL,
		NotifyInterval:     pointer.NewInt(60 * 60 * 2),
	}
}
