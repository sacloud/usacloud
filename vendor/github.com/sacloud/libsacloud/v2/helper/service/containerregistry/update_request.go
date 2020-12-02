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

package containerregistry

import (
	"context"

	"github.com/sacloud/libsacloud/v2/helper/service"
	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type UpdateRequest struct {
	ID types.ID `request:"-" validate:"required"`

	Name          *string                              `request:",omitempty" validate:"omitempty,min=1"`
	Description   *string                              `request:",omitempty" validate:"omitempty,min=1,max=512"`
	Tags          *types.Tags                          `request:",omitempty"`
	IconID        *types.ID                            `request:",omitempty"`
	AccessLevel   *types.EContainerRegistryAccessLevel `request:",omitempty"`
	VirtualDomain *string                              `request:",omitempty"`
	Users         *[]*User                             `request:",omitempty"`
	SettingsHash  string
}

func (req *UpdateRequest) Validate() error {
	return validate.Struct(req)
}

func (req *UpdateRequest) ApplyRequest(ctx context.Context, caller sacloud.APICaller) (*ApplyRequest, error) {
	client := sacloud.NewContainerRegistryOp(caller)
	current, err := client.Read(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	users, err := client.ListUsers(ctx, req.ID) // NOTE: ユーザーが登録されていなくても200が返る
	if err != nil {
		return nil, err
	}

	applyRequest := &ApplyRequest{
		ID:             req.ID,
		Name:           current.Name,
		Description:    current.Description,
		Tags:           current.Tags,
		IconID:         current.IconID,
		AccessLevel:    current.AccessLevel,
		VirtualDomain:  current.VirtualDomain,
		SubDomainLabel: current.SubDomainLabel,
		SettingsHash:   current.SettingsHash,
	}
	if users != nil {
		for _, user := range users.Users {
			applyRequest.Users = append(applyRequest.Users, &User{
				UserName:   user.UserName,
				Password:   "", // パスワードは参照できないため常に空
				Permission: user.Permission,
			})
		}
	}

	if err := service.RequestConvertTo(req, applyRequest); err != nil {
		return nil, err
	}
	return applyRequest, nil
}
