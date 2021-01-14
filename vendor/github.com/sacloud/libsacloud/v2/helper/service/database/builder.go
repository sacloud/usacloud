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

package database

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	databaseBuilder "github.com/sacloud/libsacloud/v2/helper/builder/database"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// Builder helper/builder/databaseの移行までの暫定実装
//
// 移行するまではhelper/builder/databaseを呼び出す処理のみ行う
type Builder struct {
	ID   types.ID `request:"-"`
	Zone string   `request:"-"`

	Name                  string `validate:"required"`
	Description           string `validate:"min=0,max=512"`
	Tags                  types.Tags
	IconID                types.ID
	PlanID                types.ID `validate:"required"`
	SwitchID              types.ID `validate:"required"`
	IPAddresses           []string `validate:"required,min=1,max=2,dive,ipv4"`
	NetworkMaskLen        int      `validate:"required,min=1,max=32"`
	DefaultRoute          string   `validate:"omitempty,ipv4"`
	Port                  int      `validate:"omitempty,min=1,max=65535"`
	SourceNetwork         []string `validate:"omitempty,dive,cidrv4"`
	DatabaseType          string   `validate:"required,oneof=mariadb postgres"`
	Username              string   `validate:"required"`
	Password              string   `validate:"required"`
	EnableReplication     bool
	ReplicaUserPassword   string `validate:"required_with=EnableReplication"`
	EnableWebUI           bool
	EnableBackup          bool
	BackupWeekdays        []types.EBackupSpanWeekday `validate:"required_with=EnableBackup,max=7"`
	BackupStartTimeHour   int                        `validate:"omitempty,min=0,max=23"`
	BackupStartTimeMinute int                        `validate:"omitempty,oneof=0 15 30 45"`
	Parameters            map[string]interface{}

	NoWait bool

	Caller sacloud.APICaller `request:"-"`
}

func BuilderFromResource(ctx context.Context, caller sacloud.APICaller, zone string, id types.ID) (*Builder, error) {
	client := sacloud.NewDatabaseOp(caller)
	current, err := client.Read(ctx, zone, id)
	if err != nil {
		return nil, err
	}
	var bkHour, bkMinute int
	var bkWeekdays []types.EBackupSpanWeekday
	if current.BackupSetting != nil {
		bkWeekdays = current.BackupSetting.DayOfWeek
		if current.BackupSetting.Time != "" {
			timeStrings := strings.Split(current.BackupSetting.Time, ":")
			if len(timeStrings) == 2 {
				hour, err := strconv.ParseInt(timeStrings[0], 10, 64)
				if err != nil {
					return nil, err
				}
				bkHour = int(hour)

				minute, err := strconv.ParseInt(timeStrings[1], 10, 64)
				if err != nil {
					return nil, err
				}
				bkMinute = int(minute)
			}
		}
	}
	parameters, err := client.GetParameter(ctx, zone, id)
	if err != nil {
		return nil, err
	}

	return &Builder{
		ID:                    current.ID,
		Zone:                  zone,
		Name:                  current.Name,
		Description:           current.Description,
		Tags:                  current.Tags,
		IconID:                current.IconID,
		PlanID:                current.PlanID,
		SwitchID:              current.SwitchID,
		IPAddresses:           current.IPAddresses,
		NetworkMaskLen:        current.NetworkMaskLen,
		DefaultRoute:          current.DefaultRoute,
		Port:                  current.CommonSetting.ServicePort,
		SourceNetwork:         current.CommonSetting.SourceNetwork,
		DatabaseType:          current.Conf.DatabaseName,
		Username:              current.CommonSetting.DefaultUser,
		Password:              current.CommonSetting.UserPassword,
		EnableReplication:     current.ReplicationSetting != nil,
		ReplicaUserPassword:   current.CommonSetting.ReplicaPassword,
		EnableWebUI:           current.CommonSetting.WebUI.Bool(),
		EnableBackup:          current.BackupSetting != nil,
		BackupWeekdays:        bkWeekdays,
		BackupStartTimeHour:   bkHour,
		BackupStartTimeMinute: bkMinute,
		Parameters:            parameters.Settings,
		Caller:                caller,
	}, nil
}

func (b *Builder) actualBuilder() *databaseBuilder.Builder {
	replicaUser := ""
	replicaPassword := ""
	if b.EnableReplication {
		replicaUser = "replica"
		replicaPassword = b.ReplicaUserPassword
	}
	builder := &databaseBuilder.Builder{
		PlanID:         b.PlanID,
		SwitchID:       b.SwitchID,
		IPAddresses:    b.IPAddresses,
		NetworkMaskLen: b.NetworkMaskLen,
		DefaultRoute:   b.DefaultRoute,
		Conf: &sacloud.DatabaseRemarkDBConfCommon{
			DatabaseName:     types.RDBMSVersions[types.RDBMSTypeFromString(b.DatabaseType)].Name,
			DatabaseVersion:  types.RDBMSVersions[types.RDBMSTypeFromString(b.DatabaseType)].Version,
			DatabaseRevision: types.RDBMSVersions[types.RDBMSTypeFromString(b.DatabaseType)].Revision,
		},
		CommonSetting: &sacloud.DatabaseSettingCommon{
			WebUI:           types.ToWebUI(b.EnableWebUI),
			ServicePort:     b.Port,
			SourceNetwork:   b.SourceNetwork,
			DefaultUser:     b.Username,
			UserPassword:    b.Password,
			ReplicaUser:     replicaUser,
			ReplicaPassword: replicaPassword,
		},
		Name:        b.Name,
		Description: b.Description,
		Tags:        b.Tags,
		IconID:      b.IconID,
		Parameters:  b.Parameters,
		NoWait:      b.NoWait,
		Client:      databaseBuilder.NewAPIClient(b.Caller),
	}
	if b.EnableBackup {
		builder.BackupSetting = &sacloud.DatabaseSettingBackup{
			Time:      fmt.Sprintf("%02d:%02d", b.BackupStartTimeHour, b.BackupStartTimeMinute),
			DayOfWeek: b.BackupWeekdays,
		}
	}
	if b.EnableReplication {
		builder.ReplicationSetting = &sacloud.DatabaseReplicationSetting{
			Model: types.DatabaseReplicationModels.MasterSlave,
		}
	}
	return builder
}

func (b *Builder) Build(ctx context.Context) (*sacloud.Database, error) {
	if b.ID.IsEmpty() {
		return b.create(ctx)
	}
	return b.update(ctx)
}

func (b *Builder) create(ctx context.Context) (*sacloud.Database, error) {
	return b.actualBuilder().Build(ctx, b.Zone)
}

func (b *Builder) update(ctx context.Context) (*sacloud.Database, error) {
	return b.actualBuilder().Update(ctx, b.Zone, b.ID)
}
