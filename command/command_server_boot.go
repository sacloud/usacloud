package command

import (
	"fmt"
	"github.com/sacloud/usacloud/command/internal"
)

func ServerBoot(ctx Context, params *BootServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerBoot is failed: %s", e)
	}

	if p.IsUp() {
		return nil // already booted.
	}

	compChan := make(chan bool)
	errChan := make(chan error)
	spinner := internal.NewSpinner(
		"Booting...",
		"Boot server is complete.\n",
		internal.CharSetProgress,
		GlobalOption.Progress)

	go func() {
		spinner.Start()
		// call manipurate functions
		_, err := api.Boot(params.Id)
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
	}()

boot:
	for {
		select {
		case <-compChan:
			spinner.Stop()
			break boot
		case err := <-errChan:
			return fmt.Errorf("ServerBoot is failed: %s", err)
		}
	}

	return nil

}
