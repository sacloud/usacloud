package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
)

func DatabaseWaitForBoot(ctx command.Context, params *params.WaitForBootDatabaseParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDatabaseAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DatabaseWaitForBoot is failed: %s", e)
	}

	if p.IsUp() {
		return nil // already booted.
	}

	err := internal.ExecWithProgress(
		fmt.Sprintf("Still booting[ID:%d]...", params.Id),
		fmt.Sprintf("Boot database[ID:%d]", params.Id),
		command.GlobalOption.Progress,
		func(compChan chan bool, errChan chan error) {
			// call manipurate functions
			err := api.SleepUntilUp(params.Id, client.DefaultTimeoutDuration)
			if err != nil {
				errChan <- err
				return
			}
			compChan <- true
		},
	)
	if err != nil {
		return fmt.Errorf("DatabaseWaitForBoot is failed: %s", err)
	}

	return nil

}
