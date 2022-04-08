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
	"github.com/sacloud/usacloud/pkg/ccol"
	cflag2 "github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/output"
)

var listParametersCommand = &core.Command{
	Name:               "list-parameters",
	Aliases:            []string{"list-parameter"},
	Category:           "basic",
	Order:              500,
	ServiceFuncAltName: "ListParameter",
	NoProgress:         true,
	SelectorType:       core.SelectorTypeRequireMulti,

	ColumnDefs: []output.ColumnDef{
		ccol.Zone,
		ccol.ID,
		{
			Name: "Key",
		},
		{
			Name:     "CurrentValue",
			Template: "{{ .Value }}",
		},
		{
			Name:     "Type",
			Template: "{{ .Meta.Type }}",
		},
		{
			Name:     "About",
			Template: "{{ .Meta.Text | to_single_line | ellipsis 30 }}",
		},
		{
			Name:     "Example",
			Template: "{{ .Meta.Example }}",
		},
		{
			Name:     "Reboot",
			Template: "{{ .Meta.Reboot}}",
		},
	},

	ParameterInitializer: func() interface{} {
		return newListParametersParameter()
	},
}

type listParametersParameter struct {
	cflag2.ZoneParameter   `cli:",squash" mapconv:",squash"`
	cflag2.IDParameter     `cli:",squash" mapconv:",squash"`
	cflag2.CommonParameter `cli:",squash" mapconv:"-"`
	cflag2.OutputParameter `cli:",squash" mapconv:"-"`
}

func newListParametersParameter() *listParametersParameter {
	return &listParametersParameter{}
}

func init() {
	Resource.AddCommand(listParametersCommand)
}
