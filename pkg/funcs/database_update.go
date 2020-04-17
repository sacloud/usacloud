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

package funcs

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func DatabaseUpdate(ctx cli.Context, params *params.UpdateDatabaseParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDatabaseAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DatabaseUpdate is failed: %s", e)
	}

	// is need apply config?
	configTargets := []string{
		"password", "port", "source-networks",
		"enable-web-ui", "enable-backup", "backup-time", "backup-weekdays",
		"replica-user-password", "enable-replication",
	}
	isNeedApplyConfig := cli.IsSetOr(ctx, configTargets...)

	// is slave?
	if isNeedApplyConfig && p.Settings.DBConf.Replication != nil && p.Settings.DBConf.Replication.Model == sacloud.DatabaseReplicationModelAsyncReplica {
		return fmt.Errorf("Replication slave can't update database settings")
	}

	// is need shutdown?
	needShutdownTargets := []string{"replica-user-password", "enable-replication"}
	isNeedShutdown := cli.IsSetOr(ctx, needShutdownTargets...)
	if isNeedShutdown && p.IsUp() {
		return fmt.Errorf("Need shutdown to change replication parameters[%s]", strings.Join(needShutdownTargets, "/"))
	}

	if ctx.IsSet("description") {
		p.SetDescription(params.Description)
	}
	if ctx.IsSet("icon-id") {
		p.SetIconByID(params.IconId)
	}
	if ctx.IsSet("name") {
		p.SetName(params.Name)
	}
	if ctx.IsSet("tags") {
		p.SetTags(params.Tags)
	}

	// update
	p, err := api.Update(params.Id, p)
	if err != nil {
		return fmt.Errorf("DatabaseUpdate is failed: %s", err)
	}

	// set params
	if ctx.IsSet("password") {
		p.Settings.DBConf.Common.UserPassword = params.Password
	}
	if ctx.IsSet("replica-user-password") {
		p.Settings.DBConf.Common.UserPassword = params.ReplicaUserPassword
	}
	if ctx.IsSet("enable-replication") {
		if params.EnableReplication {
			p.Settings.DBConf.Replication = &sacloud.DatabaseReplicationSetting{
				Model: sacloud.DatabaseReplicationModelMasterSlave,
			}
		} else {
			p.Settings.DBConf.Replication = nil
		}
	}

	if ctx.IsSet("port") {
		p.Settings.DBConf.Common.ServicePort = json.Number(fmt.Sprintf("%d", params.Port))
	}
	if ctx.IsSet("source-networks") {
		p.Settings.DBConf.Common.SourceNetwork = params.SourceNetworks
	}
	if ctx.IsSet("enable-web-ui") {
		p.Settings.DBConf.Common.WebUI = params.EnableWebUi
	}

	if ctx.IsSet("backup-weekdays") {

		if p.Settings.DBConf.Backup == nil {
			p.Settings.DBConf.Backup = &sacloud.DatabaseBackupSetting{}
		}

		exists := false
		for _, v := range params.BackupWeekdays {
			if v == "all" {
				exists = true
				break
			}
		}
		if exists {
			p.Settings.DBConf.Backup.DayOfWeek = sacloud.AllowDatabaseBackupWeekdays()
		} else {
			p.Settings.DBConf.Backup.DayOfWeek = params.BackupWeekdays
		}
	}
	if ctx.IsSet("backup-time") {
		if p.Settings.DBConf.Backup == nil {
			p.Settings.DBConf.Backup = &sacloud.DatabaseBackupSetting{}
		}

		p.Settings.DBConf.Backup.Time = params.BackupTime
	}

	if ctx.IsSet("enable-backup") {
		if !params.EnableBackup {
			p.Settings.DBConf.Backup = nil
		}
	}

	if isNeedApplyConfig {
		_, err = api.UpdateSetting(params.Id, p)
		if err != nil {
			return fmt.Errorf("DatabaseUpdate is failed: %s", err)
		}
		_, err = api.Config(params.Id)
		if err != nil {
			return fmt.Errorf("DatabaseUpdate is failed: %s", err)
		}
	}

	// read again
	p, e = api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DatabaseUpdate is failed: %s", e)
	}
	return ctx.GetOutput().Print(p)

}
