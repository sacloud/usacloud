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
	"fmt"

	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type CreateRequest struct {
	Zone string `request:"-" validate:"required"`

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
	SourceNetwork         []string `validate:"dive,cidrv4"`
	DatabaseType          string   `validate:"required,oneof=mariadb postgresql"`
	Username              string   `validate:"required"`
	Password              string   `validate:"required"`
	EnableReplication     bool
	ReplicaUserPassword   string `validate:"required_with=EnableReplication"`
	EnableWebUI           bool
	EnableBackup          bool
	BackupWeekdays        []types.EBackupSpanWeekday `validate:"required_with=EnableBackup,max=7"`
	BackupStartTimeHour   int                        `validate:"omitempty,min=0,max=23"`
	BackupStartTimeMinute int                        `validate:"omitempty,oneof=0 15 30 45"`
}

func (req *CreateRequest) Validate() error {
	return validate.Struct(req)
}

func (req *CreateRequest) ToRequestParameter() (*sacloud.DatabaseCreateRequest, error) {
	r := &sacloud.DatabaseCreateRequest{
		PlanID:         req.PlanID,
		SwitchID:       req.SwitchID,
		IPAddresses:    req.IPAddresses,
		NetworkMaskLen: req.NetworkMaskLen,
		DefaultRoute:   req.DefaultRoute,
		Conf: &sacloud.DatabaseRemarkDBConfCommon{
			DatabaseName:     types.RDBMSVersions[types.RDBMSTypesPostgreSQL].Name,
			DatabaseVersion:  types.RDBMSVersions[types.RDBMSTypesPostgreSQL].Version,
			DatabaseRevision: types.RDBMSVersions[types.RDBMSTypesPostgreSQL].Revision,
		},
		CommonSetting: &sacloud.DatabaseSettingCommon{
			WebUI:           types.ToWebUI(req.EnableWebUI),
			ServicePort:     req.Port,
			SourceNetwork:   req.SourceNetwork,
			DefaultUser:     req.Username,
			UserPassword:    req.Password,
			ReplicaUser:     "",
			ReplicaPassword: req.ReplicaUserPassword,
		},
		Name:        req.Name,
		Description: req.Description,
		Tags:        req.Tags,
		IconID:      req.IconID,
	}

	if req.EnableBackup {
		r.BackupSetting = &sacloud.DatabaseSettingBackup{
			Time:      fmt.Sprintf("%02d:%02d", req.BackupStartTimeHour, req.BackupStartTimeMinute),
			DayOfWeek: req.BackupWeekdays,
		}
	}
	if req.EnableReplication {
		r.ReplicationSetting = &sacloud.DatabaseReplicationSetting{
			Model: types.DatabaseReplicationModels.MasterSlave,
		}
	}
	return r, nil
}
