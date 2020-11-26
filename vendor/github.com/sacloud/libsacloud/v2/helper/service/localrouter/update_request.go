// Copyright 2016-2020 The Libsacloud Authors
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

package localrouter

import (
	"context"

	"github.com/sacloud/libsacloud/v2/helper/service"
	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type UpdateRequest struct {
	ID types.ID `request:"-" validate:"required"`

	Name         *string                            `request:",omitempty" validate:"omitempty,min=1"`
	Description  *string                            `request:",omitempty" validate:"omitempty,min=1,max=512"`
	Tags         *types.Tags                        `request:",omitempty"`
	IconID       *types.ID                          `request:",omitempty"`
	Switch       *sacloud.LocalRouterSwitch         `request:",omitempty"`
	Interface    *sacloud.LocalRouterInterface      `request:",omitempty"`
	Peers        *[]*sacloud.LocalRouterPeer        `request:",omitempty"`
	StaticRoutes *[]*sacloud.LocalRouterStaticRoute `request:",omitempty"`
	SettingsHash string
}

func (req *UpdateRequest) Validate() error {
	return validate.Struct(req)
}

func (req *UpdateRequest) Builder(ctx context.Context, caller sacloud.APICaller) (*Builder, error) {
	builder, err := BuilderFromResource(ctx, caller, req.ID)
	if err != nil {
		return nil, err
	}
	if err := service.RequestConvertTo(req, builder); err != nil {
		return nil, err
	}
	return builder, nil
}
