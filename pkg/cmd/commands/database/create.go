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
	cflag.InputParameter   `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag.NameParameter   `cli:",squash" mapconv:",squash"`
	cflag.DescParameter   `cli:",squash" mapconv:",squash"`
	cflag.TagsParameter   `cli:",squash" mapconv:",squash"`
	cflag.IconIDParameter `cli:",squash" mapconv:",squash"`

	DatabaseType string `cli:",options=database_type,category=plan,order=10" mapconv:",filters=database_type_to_value" validate:"required,database_type"`
	PlanID       string `cli:"plan,options=database_plan,category=plan,order=20" mapconv:",filters=database_plan_to_value" validate:"required,database_plan"`

	SwitchID       types.ID `cli:",category=network,order=10" validate:"required"`
	IPAddresses    []string `cli:"ip-address,aliases=ipaddress,category=network,order=20" validate:"required,min=1,max=2,dive,ipv4"`
	NetworkMaskLen int      `cli:"netmask,aliases=network-mask-len,category=network,order=30" validate:"required,min=1,max=32"`
	DefaultRoute   string   `cli:"gateway,aliases=default-route,category=network,order=40" validate:"omitempty,ipv4"`
	Port           int      `cli:",category=network,order=50" validate:"omitempty,min=1,max=65535"`
	SourceNetwork  []string `cli:"source-range,aliases=source-network,category=network,order=60" validate:"omitempty,dive,cidrv4"`

	Username string `cli:",category=user,order=10" validate:"required"`
	Password string `cli:",category=user,order=20" validate:"required"`

	EnableReplication   bool     `cli:",category=replication,order=10"`
	ReplicaUserPassword string   `cli:",category=replication,order=20" validate:"required_with=EnableReplication"`
	EnableWebUI         bool     `cli:",category=WebUI"`
	EnableBackup        bool     `cli:",category=backup,order=10"`
	BackupWeekdays      []string `cli:",options=weekdays,category=backup,order=20" mapconv:",omitempty,filters=weekdays" validate:"required_with=EnableBackup,max=7,weekdays"`

	BackupStartTimeHour   int `cli:",category=backup,order=30" validate:"omitempty,min=0,max=23"`
	BackupStartTimeMinute int `cli:",options=backup_start_minute,category=backup,order=40" validate:"omitempty,backup_start_minute"`

	cflag.NoWaitParameter `cli:",squash" mapconv:",squash"`
}

func newCreateParameter() *createParameter {
	return &createParameter{}
}

func init() {
	Resource.AddCommand(createCommand)
}
