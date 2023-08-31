// Copyright 2017-2023 The sacloud/usacloud Authors
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

	"github.com/sacloud/packages-go/pointer"
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/examples"
	"github.com/sacloud/usacloud/pkg/validate"
)

var updateCommand = &core.Command{
	Name:         "update",
	Category:     "basic",
	Order:        40,
	SelectorType: core.SelectorTypeRequireMulti,

	ColumnDefs: defaultColumnDefs,

	ValidateFunc: validateUpdateParameter,
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

	cflag.NameUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.DescUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.TagsUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.IconIDUpdateParameter `cli:",squash" mapconv:",omitempty,squash"`

	SourceNetwork *[]string `cli:"source-range,aliases=source-network,category=network" validate:"omitempty,dive,cidrv4"`

	EnableReplication     *bool     `cli:",category=replication,order=10"`
	ReplicaUserPassword   *string   `cli:",category=replication,order=20,desc=(*required when --enable-replication is specified)" validate:"omitempty,required_with=EnableReplication"`
	EnableWebUI           *bool     `cli:",category=WebUI"`
	EnableBackup          *bool     `cli:",category=backup,order=10"`
	BackupWeekdays        *[]string `cli:",options=weekdays,category=backup,order=20" mapconv:",omitempty,filters=dereference weekdays" validate:"omitempty,max=7,weekdays"`
	BackupStartTimeHour   *int      `cli:",category=backup,order=30" mapconv:",omitempty" validate:"omitempty,min=0,max=23"`
	BackupStartTimeMinute *int      `cli:",category=backup,order=40" mapconv:",omitempty" validate:"omitempty,oneof=0 15 30 45"`

	DatabaseParametersData *[]string               `cli:"database-parameters" json:"-" mapconv:"-"`
	DatabaseParameters     *map[string]interface{} `cli:"-" mapconv:"Parameters"`

	cflag.NoWaitParameter `cli:",squash" mapconv:",squash"`
}

func newUpdateParameter() *updateParameter {
	return &updateParameter{}
}

func init() {
	Resource.AddCommand(updateCommand)
}

func validateUpdateParameter(ctx cli.Context, parameter interface{}) error {
	if err := validate.Exec(parameter); err != nil {
		return err
	}
	p := parameter.(*updateParameter)

	var errs []error
	pattern := regexp.MustCompile(`^\w+=\w*$`)
	if p.DatabaseParametersData != nil {
		for _, param := range *p.DatabaseParametersData {
			if !pattern.MatchString(param) {
				errs = append(errs, validate.NewFlagError("--database-parameters",
					fmt.Sprintf("must be in a format like 'key=value': %s", param),
				))
			}
		}
	}
	return validate.NewValidationError(errs...)
}

func (p *updateParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &updateParameter{
		ZoneParameter:         examples.Zones(ctx.Option().Zones),
		NameUpdateParameter:   examples.NameUpdate,
		DescUpdateParameter:   examples.DescriptionUpdate,
		TagsUpdateParameter:   examples.TagsUpdate,
		IconIDUpdateParameter: examples.IconIDUpdate,
		SourceNetwork:         pointer.NewStringSlice([]string{"192.0.2.0/24"}),
		EnableReplication:     pointer.NewBool(true),
		ReplicaUserPassword:   pointer.NewString("password"),
		EnableWebUI:           pointer.NewBool(true),
		EnableBackup:          pointer.NewBool(true),
		BackupWeekdays:        pointer.NewStringSlice([]string{examples.OptionsString("weekdays")}),
		BackupStartTimeHour:   pointer.NewInt(1),
		BackupStartTimeMinute: pointer.NewInt(30),
		NoWaitParameter: cflag.NoWaitParameter{
			NoWait: false,
		},
	}
}

// Customize パラメータ変換処理
func (p *updateParameter) Customize(_ cli.Context) error {
	if p.DatabaseParametersData != nil {
		parameters := make(map[string]interface{})
		for _, p := range *p.DatabaseParametersData {
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
		p.DatabaseParameters = &parameters
	}
	return nil
}
