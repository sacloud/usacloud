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

package simplemonitor

import (
	"net/http"

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/cmd/examples"
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

	Target                string `validate:"required"`
	cflag.DescParameter   `cli:",squash" mapconv:",squash"`
	cflag.TagsParameter   `cli:",squash" mapconv:",squash"`
	cflag.IconIDParameter `cli:",squash" mapconv:",squash"`

	DelayLoop   int `validate:"min=60,max=3600"`
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
}

func newCreateParameter() *createParameter {
	return &createParameter{
		Enabled:        true,
		DelayLoop:      60,
		NotifyInterval: 60 * 60 * 2, // 2時間
	}
}

func init() {
	Resource.AddCommand(createCommand)
}

func (p *createParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &createParameter{
		Target:          "www.example.com",
		DescParameter:   examples.Description,
		TagsParameter:   examples.Tags,
		IconIDParameter: examples.IconID,
		DelayLoop:       60,
		Enabled:         true,
		HealthCheck: createParameterHealthCheck{
			Protocol:          examples.OptionsString("simple_monitor_protocol"),
			Port:              80,
			Path:              "/healthz",
			Status:            http.StatusOK,
			SNI:               true,
			Host:              "www2.example.com",
			BasicAuthUsername: "username",
			BasicAuthPassword: "password",
		},
		NotifyEmailEnabled: true,
		NotifyEmailHTML:    true,
		NotifySlackEnabled: true,
		SlackWebhooksURL:   examples.SlackNotifyWebhooksURL,
		NotifyInterval:     60 * 60 * 2,
	}
}
