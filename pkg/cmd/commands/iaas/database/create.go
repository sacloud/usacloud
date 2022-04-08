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

package database

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/sacloud/iaas-api-go/types"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/cmd/examples"
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
	cflag.ZoneParameter    `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag.NameParameter   `cli:",squash" mapconv:",squash"`
	cflag.DescParameter   `cli:",squash" mapconv:",squash"`
	cflag.TagsParameter   `cli:",squash" mapconv:",squash"`
	cflag.IconIDParameter `cli:",squash" mapconv:",squash"`

	DatabaseType string `cli:",options=database_type,category=plan,order=10" mapconv:",filters=database_type_to_value" validate:"required,database_type"`
	PlanID       string `cli:"plan,options=database_plan,category=plan,order=20" mapconv:",filters=database_plan_to_value" validate:"required,database_plan"`

	SwitchID       types.ID `cli:",category=network,order=10" validate:"required"`
	IPAddresses    []string `cli:"ip-address,aliases=ipaddress,category=network,order=20" validate:"required,min=1,max=1,dive,ipv4"`
	NetworkMaskLen int      `cli:"netmask,aliases=network-mask-len,category=network,order=30" validate:"required,min=1,max=32"`
	DefaultRoute   string   `cli:"gateway,aliases=default-route,category=network,order=40" validate:"omitempty,ipv4"`
	Port           int      `cli:",category=network,order=50" validate:"omitempty,min=1,max=65535"`
	SourceNetwork  []string `cli:"source-range,aliases=source-network,category=network,order=60" validate:"omitempty,dive,cidrv4"`

	Username string `cli:",category=user,order=10" validate:"required"`
	Password string `cli:",category=user,order=20" validate:"required"`

	EnableReplication   bool     `cli:",category=replication,order=10"`
	ReplicaUserPassword string   `cli:",category=replication,order=20,desc=(*required when --enable-replication is specified)" validate:"required_with=EnableReplication"`
	EnableWebUI         bool     `cli:",category=WebUI"`
	EnableBackup        bool     `cli:",category=backup,order=10"`
	BackupWeekdays      []string `cli:",options=weekdays,category=backup,order=20,desc=(*required when --enable-backup is specified)" mapconv:",omitempty,filters=weekdays" validate:"required_with=EnableBackup,max=7,weekdays"`

	BackupStartTimeHour   int `cli:",category=backup,order=30" validate:"omitempty,min=0,max=23"`
	BackupStartTimeMinute int `cli:",options=backup_start_minute,category=backup,order=40" validate:"omitempty,backup_start_minute"`

	DatabaseParametersData []string               `cli:"database-parameters" json:"-" mapconv:"-"`
	DatabaseParameters     map[string]interface{} `cli:"-" mapconv:"Parameters"`

	cflag.NoWaitParameter `cli:",squash" mapconv:",squash"`
}

func newCreateParameter() *createParameter {
	return &createParameter{}
}

func init() {
	Resource.AddCommand(createCommand)
}

func validateCreateParameter(ctx cli.Context, parameter interface{}) error {
	if err := validate.Exec(parameter); err != nil {
		return err
	}
	p := parameter.(*createParameter)

	var errs []error
	pattern := regexp.MustCompile(`^\w+=\w*$`)
	for _, param := range p.DatabaseParametersData {
		if !pattern.MatchString(param) {
			errs = append(errs, validate.NewFlagError("--database-parameters",
				fmt.Sprintf("must be in a format like 'key=value': %s", param),
			))
		}
	}
	return validate.NewValidationError(errs...)
}

func (p *createParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &createParameter{
		ZoneParameter:         examples.Zones(ctx.Option().Zones),
		NameParameter:         examples.Name,
		DescParameter:         examples.Description,
		TagsParameter:         examples.Tags,
		IconIDParameter:       examples.IconID,
		DatabaseType:          examples.OptionsString("database_type"),
		PlanID:                examples.OptionsString("database_plan"),
		SwitchID:              examples.ID,
		IPAddresses:           []string{examples.IPAddress},
		NetworkMaskLen:        examples.NetworkMaskLen,
		DefaultRoute:          examples.DefaultRoute,
		Port:                  5432,
		SourceNetwork:         []string{"192.0.2.0/24"},
		Username:              "username",
		Password:              "password",
		EnableReplication:     true,
		ReplicaUserPassword:   "password",
		EnableWebUI:           true,
		EnableBackup:          true,
		BackupWeekdays:        []string{examples.OptionsString("weekdays")},
		BackupStartTimeHour:   1,
		BackupStartTimeMinute: 30,
		DatabaseParameters: map[string]interface{}{
			"max_connections": "150",
		},
		NoWaitParameter: cflag.NoWaitParameter{
			NoWait: false,
		},
	}
}

// Customize パラメータ変換処理
func (p *createParameter) Customize(_ cli.Context) error {
	parameters := make(map[string]interface{})
	for _, p := range p.DatabaseParametersData {
		keyValues := strings.Split(p, "=")
		if len(keyValues) == 0 {
			continue // バリデーションで弾いているため到達しないはず
		}
		key := keyValues[0]
		var value interface{}
		if len(keyValues) > 1 {
			value = keyValues[1]
		}
		parameters[key] = value
	}
	p.DatabaseParameters = parameters
	return nil
}
