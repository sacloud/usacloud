package command

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
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
			_, err = api.Stop(params.Id)
			if err != nil {
				return fmt.Errorf("ServerDelete is failed: %s", err)
			}

			err = api.SleepUntilDown(params.Id, client.DefaultTimeoutDuration)
			if err != nil {
				return fmt.Errorf("ServerDelete is failed: %s", err)
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
