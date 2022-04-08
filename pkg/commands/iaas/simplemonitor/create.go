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

	cflag2 "github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/examples"
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
	cflag2.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag2.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag2.OutputParameter  `cli:",squash" mapconv:"-"`

	Target                 string `validate:"required"`
	cflag2.DescParameter   `cli:",squash" mapconv:",squash"`
	cflag2.TagsParameter   `cli:",squash" mapconv:",squash"`
	cflag2.IconIDParameter `cli:",squash" mapconv:",squash"`

	DelayLoop        int `validate:"min=60,max=3600"`
	MaxCheckAttempts int `validate:"min=1,max=10"`
	RetryInterval    int `validate:"min=10,max=3600"`

	Timeout     int `validate:"omitempty,min=1,max=30"`
	Enabled     bool
	HealthCheck createParameterHealthCheck

	NotifyEmailEnabled bool
	NotifyEmailHTML    bool
	NotifySlackEnabled bool
	SlackWebhooksURL   string
	NotifyInterval     int `validate:"min=3600,max=259200"`
}

type createParameterHealthCheck struct {
	Protocol          string `cli:",options=simple_monitor_protocol" mapconv:",filters=simple_monitor_protocol_to_value" validate:"required,simple_monitor_protocol" json:",omitempty"`
	Port              int    `json:",omitempty"`
	Path              string `json:",omitempty"`
	Status            int    `json:",omitempty"`
	ContainsString    string `json:",omitempty"`
	SNI               bool   `json:",omitempty"`
	Host              string `json:",omitempty"`
	BasicAuthUsername string `json:",omitempty"`
	BasicAuthPassword string `json:",omitempty"`
	QName             string `json:",omitempty"`
	ExpectedData      string `json:",omitempty"`
	Community         string `json:",omitempty"`
	SNMPVersion       string `json:",omitempty"`
	OID               string `json:",omitempty"`
	RemainingDays     int    `json:",omitempty"`
	HTTP2             bool   `cli:"http2" json:",omitempty"`
	FTPS              string `cli:",options=simple_monitor_ftps" mapconv:",filters=simple_monitor_ftps_to_value" validate:"omitempty,simple_monitor_ftps" json:",omitempty"`
	VerifySNI         bool   `cli:"verify-sni" json:",omitempty"`
}

func newCreateParameter() *createParameter {
	return &createParameter{
		Enabled:          true,
		MaxCheckAttempts: 3,
		RetryInterval:    10,
		DelayLoop:        60,
		NotifyInterval:   60 * 60 * 2, // 2時間
	}
}

func init() {
	Resource.AddCommand(createCommand)
}

func (p *createParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &createParameter{
		Target:           "www.example.com",
		DescParameter:    examples.Description,
		TagsParameter:    examples.Tags,
		IconIDParameter:  examples.IconID,
		DelayLoop:        60,
		MaxCheckAttempts: 3,
		RetryInterval:    10,
		Timeout:          10,
		Enabled:          true,
		HealthCheck: createParameterHealthCheck{
			Protocol:          examples.OptionsString("simple_monitor_protocol"),
			Port:              80,
			Path:              "/healthz",
			Status:            http.StatusOK,
			ContainsString:    "ok",
			SNI:               true,
			Host:              "www2.example.com",
			BasicAuthUsername: "username",
			BasicAuthPassword: "password",
			HTTP2:             true,
			FTPS:              examples.OptionsString("simple_monitor_ftps"),
			VerifySNI:         true,
		},
		NotifyEmailEnabled: true,
		NotifyEmailHTML:    true,
		NotifySlackEnabled: true,
		SlackWebhooksURL:   examples.SlackNotifyWebhooksURL,
		NotifyInterval:     60 * 60 * 2,
	}
}
