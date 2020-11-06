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

type UpdateParameter struct {
	Zone string   `cli:"-" validate:"required"`
	ID   types.ID `cli:"-" validate:"required"`

	Name        *string   `cli:",category=disk" validate:"omitempty,min=1"`
	Description *string   `cli:",category=disk" validate:"omitempty,min=1,max=512"`
	Tags        *[]string `cli:",category=diks"`
	IconID      *types.ID `cli:",category=disk"`
	Connection  *string   `cli:",category=disk,options=disk_connection"`

	*base.ConfirmParameter `cli:",squash" mapconv:"-"`
	*base.OutputParameter  `cli:",squash" mapconv:"-"`
}

func NewUpdateParameter() *UpdateParameter {
	return &UpdateParameter{
		// TODO デフォルト値はここで設定する
	}
}
