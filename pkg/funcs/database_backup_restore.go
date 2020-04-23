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

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
	"github.com/sacloud/usacloud/pkg/progress"
)

func DatabaseBackupRestore(ctx cli.Context, params *params.BackupRestoreDatabaseParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDatabaseAPI()
	info, e := api.Status(params.Id)
	if e != nil {
		return fmt.Errorf("DatabaseBackupRestore is failed: %s", e)
	}

	if !info.IsUp() {
		fmt.Fprintf(ctx.IO().Err(), "Databaes is not running\n")
		return nil
	}
	if !hasDatabaseBackup(info) {
		fmt.Fprintf(ctx.IO().Err(), "There is no backup in the database\n")
		return nil
	}
	// index
	if params.Index <= 0 || params.Index-1 >= len(info.DBConf.Backup.History) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	backupID := info.DBConf.Backup.History[params.Index-1].ID()
	err := progress.ExecWithProgress(
		fmt.Sprintf("Still restoring from backup[ID:%d:%s]...", params.Id, backupID),
		fmt.Sprintf("Restore Database[ID:%d]", params.Id),
		ctx.IO().Progress(),
		ctx.Option().NoColor,
		func(compChan chan bool, errChan chan error) {
			// call manipurate functions
			_, err := api.Restore(params.Id, backupID)
			if err != nil {
				errChan <- err
				return
			}
			compChan <- true
		},
	)
	if err != nil {
		return fmt.Errorf("DatabaseBackupRestore is failed: %s", err)
	}

	// read
	info, e = api.Status(params.Id)
	if e != nil {
		return fmt.Errorf("DatabaseBackupRestore is failed: %s", e)
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
