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

package cdrom

import (
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
	SizeGB                int `cli:"size,desc=(*required when --source-file is specified)" validate:"required_with=SourceFile,cdrom_sizes"`

	SourceFile string `mapconv:"SourceReader,filters=path_to_reader" validate:"required,file"`
}

func newCreateParameter() *createParameter {
	return &createParameter{
		SizeGB: 5,
	}
}

func init() {
	Resource.AddCommand(createCommand)
}

func (p *createParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &createParameter{
		ZoneParameter:   examples.Zones(ctx.Option().Zones),
		NameParameter:   examples.Name,
		DescParameter:   examples.Description,
		TagsParameter:   examples.Tags,
		IconIDParameter: examples.IconID,
		SizeGB:          5,
		SourceFile:      "/path/to/iso/file",
	}
}
