// Copyright 2017-2025 The sacloud/usacloud Authors
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

package gslb

import (
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/core"
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
	cflag.CommonParameter      `cli:",squash" mapconv:"-"`
	cflag.OutputParameter      `cli:",squash" mapconv:"-"`
	cflag.LimitOffsetParameter `cli:",squash" mapconv:",squash"`

	cflag.FilterByNamesParameter `cli:",squash" mapconv:",omitempty,squash"`
	cflag.FilterByTagsParameter  `cli:",squash" mapconv:",omitempty,squash"`
}

func newListParameter() *listParameter {
	return &listParameter{}
}

func init() {
	Resource.AddCommand(listCommand)
}
