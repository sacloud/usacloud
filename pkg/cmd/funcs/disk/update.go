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
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cmd/core"
)

var updateCommand = &core.Command{
	Name:         "update",
	Category:     "basics",
	Order:        40,
	SelectorType: core.SelectorTypeRequireMulti,

	ParameterInitializer: func() interface{} {
		return newUpdateParameter()
	},
}

type updateParameter struct {
	core.ZoneParameter `cli:",squash" mapconv:",squash"`
	core.IDParameter   `cli:",squash" mapconv:",squash"`

	Name        *string   `cli:",category=disk" validate:"omitempty,min=1"`
	Description *string   `cli:",category=disk" validate:"omitempty,description"`
	Tags        *[]string `cli:",category=disk" validate:"omitempty,tags"`
	IconID      *types.ID `cli:",category=disk"`
	Connection  *string   `cli:",category=disk,options=disk_connection" validate:"omitempty,disk_connection"`

	core.ConfirmParameter `cli:",squash" mapconv:"-"`
	core.OutputParameter  `cli:",squash" mapconv:"-"`
}

func newUpdateParameter() *updateParameter {
	return &updateParameter{
		// TODO デフォルト値はここで設定する
	}
}

func init() {
	Resource.AddCommand(updateCommand)
}
