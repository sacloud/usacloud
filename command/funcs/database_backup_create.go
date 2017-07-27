package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
)

func DatabaseBackupCreate(ctx command.Context, params *params.BackupCreateDatabaseParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDatabaseAPI()
	info, e := api.Status(params.Id)
	if e != nil {
		return fmt.Errorf("DatabaseBackupCreate is failed: %s", e)
	}

	if !info.IsUp() {
		fmt.Fprintf(command.GlobalOption.Err, "Databaes is not running\n")
		return nil
	}

	err := internal.ExecWithProgress(
		fmt.Sprintf("Still creating backup[ID:%d]...", params.Id),
		fmt.Sprintf("Backup Database[ID:%d]", params.Id),
		command.GlobalOption.Progress,
		func(compChan chan bool, errChan chan error) {
			// call manipurate functions
			_, err := api.Backup(params.Id)
			if err != nil {
				errChan <- err
				return
			}
			compChan <- true
		},
	)
	if err != nil {
		return fmt.Errorf("DatabaseBackupCreate is failed: %s", err)
	}

	// read
	info, e = api.Status(params.Id)
	if e != nil {
		return fmt.Errorf("DatabaseBackupCreate is failed: %s", e)
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
