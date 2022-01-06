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

package bill

import (
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
)

var listCommand = &core.Command{
	Name:       "list",
	Aliases:    []string{"ls", "find", "select"},
	Category:   "basic",
	Order:      10,
	NoProgress: true,

	ColumnDefs: listColumnDefs,

	ParameterInitializer: func() interface{} {
		return newListParameter()
	},
}

type listParameter struct {
	cflag.CommonParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter `cli:",squash" mapconv:"-"`

	Year  int `cli:",desc=(*required when --month is specified)" validate:"required_with=Month"`
	Month int `cli:",desc=(*required when --year is specified)" validate:"required_with=Year,min=0,max=12"`
}

func newListParameter() *listParameter {
	return &listParameter{}
}

func init() {
	Resource.AddCommand(listCommand)
}
