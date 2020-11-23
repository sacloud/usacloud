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
	"errors"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Builder コンテナレジストリのビルダー
type Builder struct {
	ID types.ID

	Name           string
	Description    string
	Tags           types.Tags
	IconID         types.ID
	AccessLevel    types.EContainerRegistryAccessLevel
	VirtualDomain  string
	SubDomainLabel string
	Users          []*User
	SettingsHash   string
	Client         sacloud.ContainerRegistryAPI
}

// User represents API parameter/response structure
type User struct {
	UserName   string
	Password   string
	Permission types.EContainerRegistryPermission
}

// BuilderFromResource 既存のリソースからビルダーを作成
func BuilderFromResource(ctx context.Context, caller sacloud.APICaller, id types.ID) (*Builder, error) {
	client := sacloud.NewContainerRegistryOp(caller)
	current, err := client.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	users, err := client.ListUsers(ctx, id) // NOTE: ユーザーが登録されていなくても200が返る
	if err != nil {
		return nil, err
	}

	builder := &Builder{
		ID:             id,
		Name:           current.Name,
		Description:    current.Description,
		Tags:           current.Tags,
		IconID:         current.IconID,
		AccessLevel:    current.AccessLevel,
		VirtualDomain:  current.VirtualDomain,
		SubDomainLabel: current.SubDomainLabel,
		SettingsHash:   current.SettingsHash,
		Client:         client,
	}
	if users != nil {
		for _, user := range users.Users {
			builder.Users = append(builder.Users, &User{
				UserName:   user.UserName,
				Password:   "", // パスワードは参照できないため常に空
				Permission: user.Permission,
			})
		}
	}
	return builder, nil
}

func (b *Builder) Build(ctx context.Context) (*sacloud.ContainerRegistry, error) {
	if b.ID.IsEmpty() {
		return b.create(ctx)
	}
	return b.update(ctx)
}

func (b *Builder) create(ctx context.Context) (*sacloud.ContainerRegistry, error) {
	created, err := b.Client.Create(ctx, &sacloud.ContainerRegistryCreateRequest{
		Name:           b.Name,
		Description:    b.Description,
		Tags:           b.Tags,
		IconID:         b.IconID,
		AccessLevel:    b.AccessLevel,
		VirtualDomain:  b.VirtualDomain,
		SubDomainLabel: b.SubDomainLabel,
	})
	if err != nil {
		return nil, err
	}

	if len(b.Users) == 0 {
		return created, nil
	}
	return created, b.reconcileUsers(ctx, created.ID)
}

func (b *Builder) update(ctx context.Context) (*sacloud.ContainerRegistry, error) {
	current, err := b.Client.Read(ctx, b.ID)
	if err != nil {
		return nil, err
	}
	if current.SubDomainLabel != b.SubDomainLabel {
		return nil, errors.New("SubDomainLabel cannot be changed")
	}

	updated, err := b.Client.Update(ctx, b.ID, &sacloud.ContainerRegistryUpdateRequest{
		Name:          b.Name,
		Description:   b.Description,
		Tags:          b.Tags,
		IconID:        b.IconID,
		AccessLevel:   b.AccessLevel,
		VirtualDomain: b.VirtualDomain,
		SettingsHash:  b.SettingsHash,
	})
	if err != nil {
		return nil, err
	}
	return updated, b.reconcileUsers(ctx, updated.ID)
}

func (b *Builder) reconcileUsers(ctx context.Context, id types.ID) error {
	currentUsers, err := b.Client.ListUsers(ctx, id)
	if err != nil {
		return err
	}
	if currentUsers != nil {
		// delete
		for _, username := range b.deletedUsers(currentUsers.Users) {
			if err := b.Client.DeleteUser(ctx, id, username); err != nil {
				return err
			}
		}
		// update
		for _, user := range b.updatedUsers(currentUsers.Users) {
			if err := b.Client.UpdateUser(ctx, id, user.UserName, &sacloud.ContainerRegistryUserUpdateRequest{
				Password:   user.Password,
				Permission: user.Permission,
			}); err != nil {
				return err
			}
		}
		// create
		for _, user := range b.createdUsers(currentUsers.Users) {
			if err := b.Client.AddUser(ctx, id, &sacloud.ContainerRegistryUserCreateRequest{
				UserName:   user.UserName,
				Password:   user.Password,
				Permission: user.Permission,
			}); err != nil {
				return err
			}
		}
	}
	return nil
}

func (b *Builder) deletedUsers(currentUsers []*sacloud.ContainerRegistryUser) []string {
	var results []string
	for _, current := range currentUsers {
		exists := false
		for _, desired := range b.Users {
			if current.UserName == desired.UserName {
				exists = true
				break
			}
		}
		if !exists {
			results = append(results, current.UserName)
		}
	}
	return results
}

func (b *Builder) updatedUsers(currentUsers []*sacloud.ContainerRegistryUser) []*User {
	var results []*User
	for _, current := range currentUsers {
		for _, desired := range b.Users {
			if current.UserName == desired.UserName {
				results = append(results, desired)
				break
			}
		}
	}
	return results
}

func (b *Builder) createdUsers(currentUsers []*sacloud.ContainerRegistryUser) []*User {
	var results []*User
	for _, created := range b.Users {
		exists := false
		for _, current := range currentUsers {
			if created.UserName == current.UserName {
				exists = true
				break
			}
		}
		if !exists {
			results = append(results, created)
		}
	}
	return results
}
