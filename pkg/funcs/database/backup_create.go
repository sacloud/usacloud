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
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func BackupCreate(ctx cli.Context, params *params.BackupCreateDatabaseParam) error {
	// TODO libsacloud v2実装まで保留
	return nil
	//client := sacloud.NewDatabaseOp(ctx.Client())
	//client := ctx.GetAPIClient()
	//api := client.GetDatabaseAPI()
	//info, e := api.Status(params.Id)
	//if e != nil {
	//	return fmt.Errorf("DatabaseBackupCreate is failed: %s", e)
	//}
	//
	//if !info.IsUp() {
	//	fmt.Fprintf(ctx.IO().Err(), "Databaes is not running\n")
	//	return nil
	//}
	//
	//err := progress.ExecWithProgress(
	//	fmt.Sprintf("Still creating backup[ID:%d]...", params.Id),
	//	fmt.Sprintf("Backup Database[ID:%d]", params.Id),
	//	ctx.IO().Progress(),
	//	ctx.Option().NoColor,
	//	func(compChan chan bool, errChan chan error) {
	//		// call manipurate functions
	//		_, err := api.Backup(params.Id)
	//		if err != nil {
	//			errChan <- err
	//			return
	//		}
	//		compChan <- true
	//	},
	//)
	//if err != nil {
	//	return fmt.Errorf("DatabaseBackupCreate is failed: %s", err)
	//}
	//
	//// read
	//info, e = api.Status(params.Id)
	//if e != nil {
	//	return fmt.Errorf("DatabaseBackupCreate is failed: %s", e)
	//}
	//list := []interface{}{}
	//for _, history := range info.DBConf.Backup.History {
	//	size := int64(0)
	//	if history.Size > 0 {
	//		size = history.Size / 1024 / 1024
	//	}
	//	list = append(list, &backupHistory{
	//		DatabaseBackupHistory: history,
	//		SizeMB:                size,
	//	})
	//}
	//
	//return ctx.Output().Print(list...)

}
