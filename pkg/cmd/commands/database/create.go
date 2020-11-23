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

package database

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
	cflag.ZoneParameter    `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	Name           string   `validate:"required"`
	Description    string   `validate:"description"`
	Tags           []string `validate:"tags"`
	IconID         types.ID
	PlanID         string   `cli:"plan,options=database_plan" mapconv:",filters=database_plan_to_value" validate:"required,database_plan"`
	SwitchID       types.ID `validate:"required"`
	IPAddresses    []string `validate:"required,min=1,max=2,dive,ipv4"`
	NetworkMaskLen int      `validate:"required,min=1,max=32"`
	DefaultRoute   string   `validate:"omitempty,ipv4"`
	Port           int      `validate:"omitempty,min=1,max=65535"`
	SourceNetwork  []string `validate:"omitempty,dive,cidrv4"`
	DatabaseType   string   `cli:",options=database_type" mapconv:",filters=database_type_to_value" validate:"required,database_type"`

	Username string `validate:"required"`
	Password string `validate:"required"`

	EnableReplication   bool
	ReplicaUserPassword string `validate:"required_with=EnableReplication"`
	EnableWebUI         bool
	EnableBackup        bool
	BackupWeekdays      []string `cli:",options=weekdays" mapconv:",omitempty,filters=weekdays" validate:"required_with=EnableBackup,max=7,weekdays"`

	BackupStartTimeHour   int `validate:"omitempty,min=0,max=23"`
	BackupStartTimeMinute int `cli:",options=backup_start_minute" validate:"omitempty,backup_start_minute"`
}

func newCreateParameter() *createParameter {
	return &createParameter{}
}

func init() {
	Resource.AddCommand(createCommand)
}
