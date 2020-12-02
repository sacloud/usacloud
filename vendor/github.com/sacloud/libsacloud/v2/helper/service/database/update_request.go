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
	"fmt"
	"strconv"
	"strings"

	"github.com/sacloud/libsacloud/v2/helper/service"

	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type UpdateRequest struct {
	Zone string   `validate:"required"`
	ID   types.ID `validate:"required"`

	Name        *string     `request:",omitempty" validate:"omitempty,min=1"`
	Description *string     `request:",omitempty" validate:"omitempty,min=1,max=512"`
	Tags        *types.Tags `request:",omitempty"`
	IconID      *types.ID   `request:",omitempty"`

	SourceNetwork         *[]string                   `request:",omitempty" validate:"omitempty,dive,cidrv4"`
	EnableReplication     *bool                       `request:",omitempty"`
	ReplicaUserPassword   *string                     `request:",omitempty" validate:"required_with=EnableReplication"`
	EnableWebUI           *bool                       `request:",omitempty"`
	EnableBackup          *bool                       `request:",omitempty"`
	BackupWeekdays        *[]types.EBackupSpanWeekday `request:",omitempty" validate:"required_with=EnableBackup,max=7"`
	BackupStartTimeHour   *int                        `request:",omitempty" validate:"omitempty,min=0,max=23"`
	BackupStartTimeMinute *int                        `request:",omitempty" validate:"omitempty,oneof=0 15 30 45"`

	SettingsHash string
	NoWait       bool
}

func (req *UpdateRequest) Validate() error {
	return validate.Struct(req)
}

func (req *UpdateRequest) ApplyRequest(ctx context.Context, caller sacloud.APICaller) (*ApplyRequest, error) {
	dbOp := sacloud.NewDatabaseOp(caller)
	current, err := dbOp.Read(ctx, req.Zone, req.ID)
	if err != nil {
		return nil, err
	}

	if current.Availability != types.Availabilities.Available {
		return nil, fmt.Errorf("target has invalid Availability: Zone=%s ID=%s Availability=%v", req.Zone, req.ID.String(), current.Availability)
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

	applyRequest := &ApplyRequest{
		Zone:                  req.Zone,
		ID:                    req.ID,
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
		NoWait:                false,
	}

	if err := service.RequestConvertTo(req, applyRequest); err != nil {
		return nil, err
	}
	return applyRequest, nil
}
