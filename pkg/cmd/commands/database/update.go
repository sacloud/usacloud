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
	cflag.ZoneParameter    `cli:",squash" mapconv:",squash"`
	cflag.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	Name        *string   `validate:"omitempty,min=1"`
	Description *string   `validate:"omitempty,description"`
	Tags        *[]string `validate:"omitempty,tags"`
	IconID      *types.ID

	SourceNetwork         *[]string ` validate:"omitempty,dive,cidrv4"`
	EnableReplication     *bool
	ReplicaUserPassword   *string ` validate:"omitempty,required_with=EnableReplication"`
	EnableWebUI           *bool
	EnableBackup          *bool
	BackupWeekdays        *[]string `cli:",options=weekdays" mapconv:",omitempty,filters=weekdays" validate:"omitempty,required_with=EnableBackup,max=7,weekdays"`
	BackupStartTimeHour   *int      `request:",omitempty" validate:"omitempty,min=0,max=23"`
	BackupStartTimeMinute *int      `request:",omitempty" validate:"omitempty,oneof=0 15 30 45"`

	SettingsHash string
	NoWait       bool
}

func newUpdateParameter() *updateParameter {
	return &updateParameter{}
}

func init() {
	Resource.AddCommand(updateCommand)
}
