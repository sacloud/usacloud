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
	cflag.InputParameter   `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag.NameUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.DescUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.TagsUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.IconIDUpdateParameter `cli:",squash" mapconv:",omitempty,squash"`

	SourceNetwork *[]string `cli:",category=network" validate:"omitempty,dive,cidrv4"`

	EnableReplication     *bool     `cli:",category=replication,order=10"`
	ReplicaUserPassword   *string   `cli:",category=replication,order=20" validate:"omitempty,required_with=EnableReplication"`
	EnableWebUI           *bool     `cli:",category=WebUI"`
	EnableBackup          *bool     `cli:",category=backup,order=10"`
	BackupWeekdays        *[]string `cli:",options=weekdays,category=backup,order=20" mapconv:",omitempty,filters=weekdays" validate:"omitempty,required_with=EnableBackup,max=7,weekdays"`
	BackupStartTimeHour   *int      `cli:",category=backup,order=30" mapconv:",omitempty" validate:"omitempty,min=0,max=23"`
	BackupStartTimeMinute *int      `cli:",category=backup,order=40" mapconv:",omitempty" validate:"omitempty,oneof=0 15 30 45"`

	cflag.NoWaitParameter `cli:",squash" mapconv:",squash"`
}

func newUpdateParameter() *updateParameter {
	return &updateParameter{}
}

func init() {
	Resource.AddCommand(updateCommand)
}
