package command

import (
	"fmt"
)

func ArchiveWaitForCopy(ctx Context, params *WaitForCopyArchiveParam) error {

	client := ctx.GetAPIClient()
	api := client.GetArchiveAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ArchiveWaitForCopy is failed: %s", e)
	}

	err := api.SleepWhileCopying(p.ID, client.DefaultTimeoutDuration)
	if err != nil {
		return fmt.Errorf("ArchiveWaitForCopy is failed: %s", err)
	}

	return nil

}
