package funcs

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func DatabaseBackupInfo(ctx command.Context, params *params.BackupInfoDatabaseParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDatabaseAPI()
	info, e := api.Status(params.Id)
	if e != nil {
		return fmt.Errorf("DatabaseBackupInfo is failed: %s", e)
	}

	if !info.IsUp() {
		fmt.Fprintf(command.GlobalOption.Err, "Databaes is not running\n")
		return nil
	}
	if !hasDatabaseBackup(info) {
		fmt.Fprintf(command.GlobalOption.Err, "There is no backup in the database\n")
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
