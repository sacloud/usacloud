package command

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command/internal"
	"os"
)

func ServerDelete(ctx Context, params *DeleteServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()

	p, err := api.Read(params.Id)
	if err != nil {
		return fmt.Errorf("ServerDelete is failed: %s", err)
	}

	if p.IsUp() {
		if params.Force {

			compChan := make(chan bool)
			errChan := make(chan error)
			spinner := internal.NewProgress(
				fmt.Sprintf("Still waiting for Shutdown[ID:%d]...", params.Id),
				fmt.Sprintf("Shutdown server[ID:%d]", params.Id),
				GlobalOption.Progress)

			go func() {
				spinner.Start()
				// call manipurate functions
				var err error
				_, err = api.Stop(params.Id)
				if err != nil {
					errChan <- err
					return
				}

				err = api.SleepUntilDown(params.Id, client.DefaultTimeoutDuration)
				if err != nil {
					errChan <- err
					return
				}
				compChan <- true
			}()

		down:
			for {
				select {
				case <-compChan:
					spinner.Stop()
					break down
				case err := <-errChan:
					return fmt.Errorf("ServerDelete is failed: %s", err)
				}
			}

		} else {
			return fmt.Errorf("Server(%d) is still running", params.Id)
		}
	}

	// call Delete(id)
	var res *sacloud.Server
	if params.WithDisk && len(p.Disks) > 0 {
		res, err = api.DeleteWithDisk(params.Id, p.GetDiskIDs())
		if err != nil {
			return fmt.Errorf("ServerDelete is failed: %s", err)
		}
	} else {
		res, err = api.Delete(params.Id)
		if err != nil {
			return fmt.Errorf("ServerDelete is failed: %s", err)
		}
	}

	// Delete generated ssh-key on default location
	keyFile, err := getSSHPrivateKeyStorePath(res.ID)
	if err != nil {
		return fmt.Errorf("ServerDelete is failed: %s", err)
	}
	if _, e := os.Stat(keyFile); e == nil {
		err = os.Remove(keyFile)
		if err != nil {
			return fmt.Errorf("ServerDelete is failed: %s", err)
		}
	}

	return ctx.GetOutput().Print(res)

}
