package command

import (
	"fmt"
)

func DiskReinstallToBlank(ctx Context, params *ReinstallToBlankDiskParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDiskAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DiskReinstallToBlank is failed: %s", e)
	}

	// call manipurate functions
	_, err := api.ReinstallFromBlank(params.Id, p.SizeMB)
	if err != nil {
		return fmt.Errorf("DiskReinstallToBlank is failed: %s", err)
	}

	if !params.Async {
		err = api.SleepWhileCopying(params.Id, client.DefaultTimeoutDuration)
		if err != nil {
			return fmt.Errorf("DiskReinstallToBlank is failed: %s", err)
		}
		p, err = api.Read(params.Id)
		if err != nil {
			return fmt.Errorf("DiskReinstallToBlank is failed: %s", err)
		}
	}

	return ctx.GetOutput().Print(p)

}
