package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func DatabaseBackupUnlock(ctx command.Context, params *params.BackupUnlockDatabaseParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDatabaseAPI()
	info, e := api.Status(params.Id)
	if e != nil {
		return fmt.Errorf("DatabaseBackupUnlock is failed: %s", e)
	}

	if !info.IsUp() {
		fmt.Fprintf(command.GlobalOption.Err, "Databaes is not running\n")
		return nil
	}
	if !hasDatabaseBackup(info) {
		fmt.Fprintf(command.GlobalOption.Err, "There is no backup in the database\n")
		return nil
	}
	// index
	if params.Index <= 0 || params.Index-1 >= len(info.DBConf.Backup.History) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	backupID := info.DBConf.Backup.History[params.Index-1].Id()
	_, err := api.HistoryUnlock(params.Id, backupID)
	if err != nil {
		return fmt.Errorf("DatabaseBackupUnlock is failed: %s", e)
	}

	// read
	info, e = api.Status(params.Id)
	if e != nil {
		return fmt.Errorf("DatabaseBackupUnlock is failed: %s", e)
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
