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
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
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

	Target      string   `validate:"required"`
	Description string   `validate:"description"`
	Tags        []string `validate:"tags"`
	IconID      types.ID

	DelayLoop   int `validate:"min=60,max=3600"`
	Enabled     bool
	HealthCheck createParameterHealthCheck

	NotifyEmailEnabled bool
	NotifyEmailHTML    bool
	NotifySlackEnabled string
	SlackWebhooksURL   string
	NotifyInterval     int `validate:"min=3600,max=259200"`
}

type createParameterHealthCheck struct {
	Protocol          string `cli:",options=simple_monitor_protocol" mapconv:",filters=simple_monitor_protocol_to_value" validate:"required,simple_monitor_protocol"`
	Port              int
	Path              string
	Status            int
	SNI               bool
	Host              string
	BasicAuthUsername string
	BasicAuthPassword string
	QName             string
	ExpectedData      string
	Community         string
	SNMPVersion       string
	OID               string
	RemainingDays     int
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
