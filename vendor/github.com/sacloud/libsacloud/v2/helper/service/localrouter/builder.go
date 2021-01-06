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
	"context"

	localrouterBuilder "github.com/sacloud/libsacloud/v2/helper/builder/localrouter"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Builder ローカルルータの構築を行う
type Builder struct {
	ID          types.ID
	Name        string
	Description string
	Tags        types.Tags
	IconID      types.ID

	Switch       *sacloud.LocalRouterSwitch
	Interface    *sacloud.LocalRouterInterface
	Peers        []*sacloud.LocalRouterPeer
	StaticRoutes []*sacloud.LocalRouterStaticRoute

	SettingsHash string

	Caller sacloud.APICaller
}

func BuilderFromResource(ctx context.Context, caller sacloud.APICaller, id types.ID) (*Builder, error) {
	client := sacloud.NewLocalRouterOp(caller)
	current, err := client.Read(ctx, id)
	if err != nil {
		return nil, err
	}

	return &Builder{
		Name:         current.Name,
		Description:  current.Description,
		Tags:         current.Tags,
		IconID:       current.IconID,
		Switch:       current.Switch,
		Interface:    current.Interface,
		Peers:        current.Peers,
		StaticRoutes: current.StaticRoutes,
		Caller:       caller,
	}, nil
}

func (b *Builder) Build(ctx context.Context) (*sacloud.LocalRouter, error) {
	builder := &localrouterBuilder.Builder{
		Name:         b.Name,
		Description:  b.Description,
		Tags:         b.Tags,
		IconID:       b.IconID,
		Switch:       b.Switch,
		Interface:    b.Interface,
		Peers:        b.Peers,
		StaticRoutes: b.StaticRoutes,
		SettingsHash: b.SettingsHash,
		Client:       localrouterBuilder.NewAPIClient(b.Caller),
	}

	if b.ID.IsEmpty() {
		return builder.Build(ctx)
	}
	return builder.Update(ctx, b.ID)
}
