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

package loadbalancer

import (
	"context"
	"errors"

	"github.com/sacloud/libsacloud/v2/helper/wait"
	"github.com/sacloud/libsacloud/v2/pkg/util"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type Builder struct {
	ID   types.ID
	Zone string

	Name               string
	Description        string
	Tags               types.Tags
	IconID             types.ID
	SwitchID           types.ID
	PlanID             types.ID
	VRID               int
	IPAddresses        []string
	NetworkMaskLen     int
	DefaultRoute       string
	VirtualIPAddresses sacloud.LoadBalancerVirtualIPAddresses

	NoWait       bool
	SettingsHash string // for update
	Client       sacloud.LoadBalancerAPI
}

func (b *Builder) Build(ctx context.Context) (*sacloud.LoadBalancer, error) {
	if b.ID.IsEmpty() {
		return b.create(ctx)
	}
	return b.update(ctx)
}

func (b *Builder) create(ctx context.Context) (*sacloud.LoadBalancer, error) {
	created, err := b.Client.Create(ctx, b.Zone, &sacloud.LoadBalancerCreateRequest{
		SwitchID:           b.SwitchID,
		PlanID:             b.PlanID,
		VRID:               b.VRID,
		IPAddresses:        b.IPAddresses,
		NetworkMaskLen:     b.NetworkMaskLen,
		DefaultRoute:       b.DefaultRoute,
		Name:               b.Name,
		Description:        b.Description,
		Tags:               b.Tags,
		IconID:             b.IconID,
		VirtualIPAddresses: b.VirtualIPAddresses,
	})
	if err != nil {
		return nil, err
	}
	if b.NoWait {
		return created, nil
	}
	return wait.UntilLoadBalancerIsUp(ctx, b.Client, b.Zone, created.ID)
}

func (b *Builder) validateForUpdate(current *sacloud.LoadBalancer) error {
	if current.SwitchID != b.SwitchID {
		return errors.New("SwitchID cannot be changed")
	}
	if current.PlanID != b.PlanID {
		return errors.New("PlanID cannot be changed")
	}
	if current.VRID != b.VRID {
		return errors.New("VRID cannot be changed")
	}
	if !util.DeepEqual(current.IPAddresses, b.IPAddresses) {
		return errors.New("IPAddresses cannot be changed")
	}
	if current.NetworkMaskLen != b.NetworkMaskLen {
		return errors.New("NetworkMaskLen cannot be changed")
	}
	if current.DefaultRoute != b.DefaultRoute {
		return errors.New("DefaultRoute cannot be changed")
	}
	return nil
}

func (b *Builder) update(ctx context.Context) (*sacloud.LoadBalancer, error) {
	current, err := b.Client.Read(ctx, b.Zone, b.ID)
	if err != nil {
		return nil, err
	}
	if err := b.validateForUpdate(current); err != nil {
		return nil, err
	}

	updated, err := b.Client.Update(ctx, b.Zone, b.ID, &sacloud.LoadBalancerUpdateRequest{
		Name:               b.Name,
		Description:        b.Description,
		Tags:               b.Tags,
		IconID:             b.IconID,
		VirtualIPAddresses: b.VirtualIPAddresses,
		SettingsHash:       b.SettingsHash,
	})
	if err != nil {
		return nil, err
	}
	return updated, b.Client.Config(ctx, b.Zone, b.ID)
}
