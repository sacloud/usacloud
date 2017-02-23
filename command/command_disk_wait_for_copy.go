package command

import (
	"fmt"
)

func DiskWaitForCopy(ctx Context, params *WaitForCopyDiskParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDiskAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DiskWaitForCopy is failed: %s", e)
	}

	err := api.SleepWhileCopying(p.ID, client.DefaultTimeoutDuration)
	if err != nil {
		return fmt.Errorf("DiskWaitForCopy is failed: %s", e)
	}

	// read again
	p, err = api.Read(p.ID)
	if err != nil {
		return fmt.Errorf("DiskWaitForCopy is failed: %s", e)
	}

	return ctx.GetOutput().Print(p)

}
