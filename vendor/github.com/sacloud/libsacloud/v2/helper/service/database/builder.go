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

package database

import (
	"context"

	databaseBuilder "github.com/sacloud/libsacloud/v2/helper/builder/database"
	"github.com/sacloud/libsacloud/v2/helper/service"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Builder helper/builder/databaseの移行までの暫定実装
//
// 移行するまではhelper/builder/databaseを呼び出す処理のみ行う
type Builder struct {
	ID   types.ID `request:"-"`
	Zone string   `request:"-"`

	Name        string
	Description string
	Tags        types.Tags
	IconID      types.ID

	PlanID             types.ID
	SwitchID           types.ID
	IPAddresses        []string
	NetworkMaskLen     int
	DefaultRoute       string
	Conf               *sacloud.DatabaseRemarkDBConfCommon
	SourceID           types.ID
	CommonSetting      *sacloud.DatabaseSettingCommon
	BackupSetting      *sacloud.DatabaseSettingBackup
	ReplicationSetting *sacloud.DatabaseReplicationSetting

	SettingsHash string

	Caller sacloud.APICaller `request:"-"`
}

func BuilderFromResource(ctx context.Context, caller sacloud.APICaller, zone string, id types.ID) (*Builder, error) {
	client := sacloud.NewDatabaseOp(caller)
	current, err := client.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	return &Builder{
		ID:                 current.ID,
		Zone:               zone,
		Name:               current.Name,
		Description:        current.Description,
		Tags:               current.Tags,
		IconID:             current.IconID,
		PlanID:             current.PlanID,
		SwitchID:           current.SwitchID,
		IPAddresses:        current.IPAddresses,
		NetworkMaskLen:     current.NetworkMaskLen,
		DefaultRoute:       current.DefaultRoute,
		Conf:               current.Conf,
		CommonSetting:      current.CommonSetting,
		BackupSetting:      current.BackupSetting,
		ReplicationSetting: current.ReplicationSetting,
		SettingsHash:       current.SettingsHash,
		Caller:             caller,
	}, nil
}

func (b *Builder) Build(ctx context.Context) (*sacloud.Database, error) {
	if b.ID.IsEmpty() {
		return b.create(ctx)
	}
	return b.update(ctx)
}

func (b *Builder) create(ctx context.Context) (*sacloud.Database, error) {
	builder := &databaseBuilder.Builder{
		Client: databaseBuilder.NewAPIClient(b.Caller),
	}
	if err := service.RequestConvertTo(b, builder); err != nil {
		return nil, err
	}
	return builder.Build(ctx, b.Zone)
}

func (b *Builder) update(ctx context.Context) (*sacloud.Database, error) {
	builder := &databaseBuilder.Builder{
		Client: databaseBuilder.NewAPIClient(b.Caller),
	}
	if err := service.RequestConvertTo(b, builder); err != nil {
		return nil, err
	}
	return builder.Update(ctx, b.Zone, b.ID)
}
