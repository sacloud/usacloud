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

package disk

import (
	"github.com/sacloud/usacloud/pkg/cmd/core"
)

var deleteCommand = &core.Command{
	Name:         "delete",
	Aliases:      []string{"rm"},
	Category:     "basics",
	Order:        50,
	SelectorType: core.SelectorTypeRequireMulti,

	ParameterInitializer: func() interface{} {
		return newDeleteParameter()
	},
}

type deleteParameter struct {
	core.ZoneParameter `cli:",squash" mapconv:",squash"`
	core.IDParameter   `cli:",squash" mapconv:",squash"`

	FailIfNotFound bool `cli:",category=disk"`

	core.ConfirmParameter `cli:",squash" mapconv:"-"`
	core.OutputParameter  `cli:",squash" mapconv:"-"`
}

func newDeleteParameter() *deleteParameter {
	return &deleteParameter{}
}

func init() {
	Resource.AddCommand(deleteCommand)
}
