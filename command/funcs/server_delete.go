package funcs

import (
	"fmt"
	"os"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/internal"
	"github.com/sacloud/usacloud/command/params"
)

func ServerDelete(ctx command.Context, params *params.DeleteServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()

	p, err := api.Read(params.Id)
	if err != nil {
		return fmt.Errorf("ServerDelete is failed: %s", err)
	}

	if p.IsUp() {
		if params.Force {

			err := internal.ExecWithProgress(
				fmt.Sprintf("Still waiting for Delete[ID:%d]...", params.Id),
				fmt.Sprintf("Delete server[ID:%d]", params.Id),
				command.GlobalOption.Progress,
				func(compChan chan bool, errChan chan error) {
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
				},
			)
			if err != nil {
				return fmt.Errorf("ServerDelete is failed: %s", err)
			}
		} else {
			return fmt.Errorf("Server(%d) is still running", params.Id)
		}
	}

	// call Delete(id)
	var res *sacloud.Server
	if !params.WithoutDisk && len(p.Disks) > 0 {
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
