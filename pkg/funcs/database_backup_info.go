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
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func DatabaseBackupInfo(ctx cli.Context, params *params.BackupInfoDatabaseParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDatabaseAPI()
	info, e := api.Status(params.Id)
	if e != nil {
		return fmt.Errorf("DatabaseBackupInfo is failed: %s", e)
	}

	if !info.IsUp() {
		fmt.Fprintf(ctx.IO().Err(), "Databaes is not running\n")
		return nil
	}
	if !hasDatabaseBackup(info) {
		fmt.Fprintf(ctx.IO().Err(), "There is no backup in the database\n")
		return nil
	}

	list := []interface{}{}
	for _, history := range info.DBConf.Backup.History {
		size := int64(0)
		if history.Size > 0 {
			size = history.Size / 1024 / 1024
		}
		list = append(list, &backupHistory{
			DatabaseBackupHistory: history,
			SizeMB:                size,
		})
	}

	return ctx.GetOutput().Print(list...)
}

type backupHistory struct {
	*sacloud.DatabaseBackupHistory
	SizeMB int64
}

func hasDatabaseBackup(status *sacloud.DatabaseStatus) bool {
	return status.DBConf != nil &&
		status.DBConf.Backup != nil &&
		len(status.DBConf.Backup.History) > 0
}
