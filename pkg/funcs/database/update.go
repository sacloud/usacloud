// Copyright 2017-2020 The Usacloud Authors
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
	"strings"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
	"github.com/sacloud/usacloud/pkg/weekday"
)

func Update(ctx cli.Context, params *params.UpdateDatabaseParam) error {
	client := sacloud.NewDatabaseOp(ctx.Client())
	db, err := client.Read(ctx, ctx.Zone(), params.Id)
	if err != nil {
		return fmt.Errorf("DatabaseUpdate is failed: %s", err)
	}

	// is need apply config?
	configTargets := []string{
		"password", "port", "source-networks",
		"enable-web-ui", "enable-backup", "backup-time", "backup-weekdays",
		"replica-user-password", "enable-replication",
	}
	isNeedApplyConfig := false
	for _, target := range configTargets {
		if params.Changed(target) {
			isNeedApplyConfig = true
			break
		}
	}

	// is slave?
	if isNeedApplyConfig && db.ReplicationSetting != nil && db.ReplicationSetting.Model == types.DatabaseReplicationModels.MasterSlave {
		return fmt.Errorf("slave's database setting couldn't update")
	}

	// is need shutdown?
	needShutdownTargets := []string{"replica-user-password", "enable-replication"}
	isNeedShutdown := false
	for _, target := range needShutdownTargets {
		if params.Changed(target) {
			isNeedShutdown = true
			break
		}
	}
	if isNeedShutdown && db.InstanceStatus.IsUp() {
		return fmt.Errorf("need shutdown to update replication parameters[%s]", strings.Join(needShutdownTargets, "/"))
	}

	if params.Changed("description") {
		db.Description = params.Description
	}
	if params.Changed("icon-id") {
		db.IconID = params.IconId
	}
	if params.Changed("name") {
		db.Name = params.Name
	}
	if params.Changed("tags") {
		db.Tags = params.Tags
	}
	if params.Changed("password") {
		db.CommonSetting.UserPassword = params.Password
	}
	if params.Changed("replica-user-password") {
		db.CommonSetting.ReplicaPassword = params.ReplicaUserPassword
	}
	if params.Changed("enable-replication") {
		if params.EnableReplication {
			db.ReplicationSetting = &sacloud.DatabaseReplicationSetting{
				Model: types.DatabaseReplicationModels.MasterSlave,
			}
		} else {
			db.ReplicationSetting = nil
		}
	}

	if params.Changed("port") {
		db.CommonSetting.ServicePort = params.Port
	}
	if params.Changed("source-networks") {
		db.CommonSetting.SourceNetwork = params.SourceNetworks
	}
	if params.Changed("enable-web-ui") {
		db.CommonSetting.WebUI = types.WebUI(fmt.Sprintf("%t", params.EnableWebUi))
	}

	if params.Changed("backup-weekdays") {
		if db.BackupSetting == nil {
			db.BackupSetting = &sacloud.DatabaseSettingBackup{}
		}
		db.BackupSetting.DayOfWeek = weekday.FromStrings(params.BackupWeekdays)
	}
	if params.Changed("backup-time") {
		if db.BackupSetting == nil {
			db.BackupSetting = &sacloud.DatabaseSettingBackup{}
		}
		db.BackupSetting.Time = params.BackupTime
	}
	if params.Changed("enable-backup") {
		if !params.EnableBackup {
			db.BackupSetting = nil
		}
	}

	// update
	_, err = client.Update(ctx, ctx.Zone(), params.Id, &sacloud.DatabaseUpdateRequest{
		Name:        db.Name,
		Description: db.Description,
		Tags:        db.Tags,
		IconID:      db.IconID,
		CommonSetting: &sacloud.DatabaseSettingCommon{
			WebUI:           db.CommonSetting.WebUI,
			ServicePort:     db.CommonSetting.ServicePort,
			SourceNetwork:   db.CommonSetting.SourceNetwork,
			DefaultUser:     db.CommonSetting.DefaultUser,
			UserPassword:    db.CommonSetting.UserPassword,
			ReplicaUser:     db.CommonSetting.ReplicaUser,
			ReplicaPassword: db.CommonSetting.ReplicaPassword,
		},
		BackupSetting:      db.BackupSetting,
		ReplicationSetting: db.ReplicationSetting,
		SettingsHash:       db.SettingsHash,
	})
	if err != nil {
		return fmt.Errorf("DatabaseUpdate is failed: %s", err)
	}

	if err := client.Config(ctx, ctx.Zone(), params.Id); err != nil {
		return fmt.Errorf("DatabaseUpdate is failed: %s", err)
	}

	// read again
	db, err = client.Read(ctx, ctx.Zone(), params.Id)
	if err != nil {
		return fmt.Errorf("DatabaseUpdate is failed: %s", err)
	}
	return ctx.Output().Print(db)

}
