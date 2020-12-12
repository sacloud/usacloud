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

package autobackup

import (
	"github.com/sacloud/libsacloud/v2/sacloud/types"
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
	cflag.ZoneParameter    `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag.NameParameter   `cli:",squash" mapconv:",squash"`
	cflag.DescParameter   `cli:",squash" mapconv:",squash"`
	cflag.TagsParameter   `cli:",squash" mapconv:",squash"`
	cflag.IconIDParameter `cli:",squash" mapconv:",squash"`

	DiskID           types.ID `validate:"required"`
	Weekdays         []string `cli:",options=weekdays" mapconv:"BackupSpanWeekdays,omitempty,filters=weekdays" validate:"required,weekdays"`
	MaxNumOfArchives int      `cli:"max-backup-num" mapconv:"MaximumNumberOfArchives" validate:"required,min=1,max=10"`
}

func newCreateParameter() *createParameter {
	return &createParameter{}
}

func init() {
	Resource.AddCommand(createCommand)
}

func (p *createParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &createParameter{
		ZoneParameter:    examples.Zones(ctx.Option().Zones),
		NameParameter:    examples.Name,
		DescParameter:    examples.Description,
		TagsParameter:    examples.Tags,
		IconIDParameter:  examples.IconID,
		DiskID:           examples.ID,
		Weekdays:         []string{examples.OptionsString("weekdays")},
		MaxNumOfArchives: 5,
	}
}
