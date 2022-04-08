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

package note

import (
	cflag2 "github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/examples"
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
	cflag2.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag2.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag2.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag2.NameParameter   `cli:",squash" mapconv:",squash"`
	cflag2.TagsParameter   `cli:",squash" mapconv:",squash"`
	cflag2.IconIDParameter `cli:",squash" mapconv:",squash"`
	Class                  string `cli:",options=note_class" validate:"required,note_class"`
	Content                string `cli:",aliases=contents script scripts" validate:"required" mapconv:",filters=path_or_content"`
}

func newCreateParameter() *createParameter {
	return &createParameter{
		Class: "shell",
	}
}

func init() {
	Resource.AddCommand(createCommand)
}

func (p *createParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &createParameter{
		NameParameter:   examples.Name,
		TagsParameter:   examples.Tags,
		IconIDParameter: examples.IconID,
		Class:           examples.OptionsString("note_class"),
		Content:         examples.ScriptContent,
	}
}
