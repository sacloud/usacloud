package funcs

import (
	"fmt"
	"time"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
)

const maxErrorCount = 10

func ServerBoot(ctx command.Context, params *params.BootServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerBoot is failed: %s", e)
	}

	if p.IsUp() {
		return nil // already booted.
	}

	err := internal.ExecWithProgress(
		fmt.Sprintf("Still booting[ID:%d]...", params.Id),
		fmt.Sprintf("Boot server[ID:%d]", params.Id),
		command.GlobalOption.Progress,
		func(compChan chan bool, errChan chan error) {
			// call manipurate functions
			errCount := 0
			var err error

			for errCount < maxErrorCount {
				_, err = api.Boot(params.Id)
				if err == nil {
					break
				}
				errCount++
				time.Sleep(1 * time.Second)
			}
			if err != nil {
				errChan <- err
				return
			}

			err = api.SleepUntilUp(params.Id, client.DefaultTimeoutDuration)
			if err != nil {
				errChan <- err
				return
			}
			compChan <- true
		},
	)
	if err != nil {
		return fmt.Errorf("ServerBoot is failed: %s", err)
	}

	return nil
}
