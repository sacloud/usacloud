// Copyright 2016-2021 The Libsacloud Authors
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
	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type CreateRequest struct {
	Name         string `validate:"required"`
	Description  string `validate:"min=0,max=512"`
	Tags         types.Tags
	IconID       types.ID
	Switch       *sacloud.LocalRouterSwitch    `request:",omitempty" validate:"required_with=Interface"`
	Interface    *sacloud.LocalRouterInterface `request:",omitempty" validate:"required_with=Switch"`
	Peers        []*sacloud.LocalRouterPeer
	StaticRoutes []*sacloud.LocalRouterStaticRoute
}

func (req *CreateRequest) Validate() error {
	return validate.Struct(req)
}

func (req *CreateRequest) Builder(caller sacloud.APICaller) *Builder {
	return &Builder{
		Name:         req.Name,
		Description:  req.Description,
		Tags:         req.Tags,
		IconID:       req.IconID,
		Switch:       req.Switch,
		Interface:    req.Interface,
		Peers:        req.Peers,
		StaticRoutes: req.StaticRoutes,
		SettingsHash: "",
		Caller:       caller,
	}
}
