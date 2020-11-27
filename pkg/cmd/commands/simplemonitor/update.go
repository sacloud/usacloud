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

	Description *string   `validate:"omitempty,description"`
	Tags        *[]string `validate:"omitempty,tags"`
	IconID      *types.ID

	DelayLoop   *int `validate:"omitempty,min=60,max=3600"`
	Enabled     *bool
	HealthCheck updateParameterHealthCheck `mapconv:",omitempty"`

	NotifyEmailEnabled *bool
	NotifyEmailHTML    *bool
	NotifySlackEnabled *string
	SlackWebhooksURL   *string
	NotifyInterval     *int `validate:"omitempty,min=3600,max=259200"`

	SettingsHash string
}

type updateParameterHealthCheck struct {
	Protocol          *string `cli:",options=simple_monitor_protocol" mapconv:",omitempty,filters=simple_monitor_protocol_to_value" validate:"omitempty,simple_monitor_protocol"`
	Port              *int
	Path              *string
	Status            *int
	SNI               *bool
	Host              *string
	BasicAuthUsername *string
	BasicAuthPassword *string
	QName             *string
	ExpectedData      *string
	Community         *string
	SNMPVersion       *string
	OID               *string
	RemainingDays     *int
}

func newUpdateParameter() *updateParameter {
	return &updateParameter{
		// TODO デフォルト値はここで設定する
	}
}

func init() {
	Resource.AddCommand(updateCommand)
}
