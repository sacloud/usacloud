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

package internet

import (
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/core"
)

var listSubnetCommand = &core.Command{
	Name:       "list-subnet",
	Aliases:    []string{"subnet-list"},
	Category:   "subnet",
	Order:      10,
	NoProgress: true,

	ColumnDefs:   subnetColumnDefs,
	SelectorType: core.SelectorTypeRequireMulti,

	ParameterInitializer: func() interface{} {
		return newListSubnetParameter()
	},
}

type listSubnetParameter struct {
	cflag.ZoneParameter   `cli:",squash" mapconv:",squash"`
	cflag.IDParameter     `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter `cli:",squash" mapconv:"-"`
}

func newListSubnetParameter() *listSubnetParameter {
	return &listSubnetParameter{}
}

func init() {
	Resource.AddCommand(listSubnetCommand)
}
