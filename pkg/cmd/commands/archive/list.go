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

package archive

import (
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
)

var listCommand = &core.Command{
	Name:               "list",
	Aliases:            []string{"ls", "find", "select"},
	Category:           "basic",
	Order:              10,
	ServiceFuncAltName: "Find",
	NoProgress:         true,

	ColumnDefs: defaultColumnDefs,

	ParameterInitializer: func() interface{} {
		return newListParameter()
	},
}

type listParameter struct {
	cflag.ZoneParameter        `cli:",squash" mapconv:",squash"`
	cflag.InputParameter       `cli:",squash" mapconv:"-"`
	cflag.LimitOffsetParameter `cli:",squash" mapconv:",squash"`
	cflag.OutputParameter      `cli:",squash" mapconv:"-"`

	cflag.FilterByNamesParameter `cli:",squash" mapconv:",omitempty,squash"`
	cflag.FilterByTagsParameter  `cli:",squash" mapconv:",omitempty,squash"`
	cflag.FilterByScopeParameter `cli:",squash" mapconv:",omitempty,squash"`
	OSType                       string `cli:",category=filter,options=os_type_simple,order=100" mapconv:",omitempty,filters=os_type_to_value" validate:"omitempty,os_type"`
}

func newListParameter() *listParameter {
	return &listParameter{}
}

func init() {
	Resource.AddCommand(listCommand)
}
