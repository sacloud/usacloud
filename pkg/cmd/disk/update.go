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
	"github.com/sacloud/usacloud/pkg/cmd/base"
)

var updateCommand = &base.Command{
	Name:     "update",
	Category: "basics",
	Order:    40,

	ParameterInitializer: func() interface{} {
		return newUpdateParameter()
	},
}

type updateParameter struct {
	base.ZoneParameter `cli:",squash" mapconv:",squash"`
	base.IDParameter   `cli:",squash" mapconv:",squash"`

	Name        *string   `cli:",category=disk" validate:"omitempty,min=1"`
	Description *string   `cli:",category=disk" validate:"omitempty,description"`
	Tags        *[]string `cli:",category=disk" validate:"omitempty,tags"`
	IconID      *types.ID `cli:",category=disk"`
	Connection  *string   `cli:",category=disk,options=disk_connection" validate:"omitempty,disk_connection"`

	base.ConfirmParameter `cli:",squash" mapconv:"-"`
	base.OutputParameter  `cli:",squash" mapconv:"-"`
}

func newUpdateParameter() *updateParameter {
	return &updateParameter{
		// TODO デフォルト値はここで設定する
	}
}

func init() {
	Resource.AddCommand(updateCommand)
}
