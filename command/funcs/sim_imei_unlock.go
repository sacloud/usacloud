package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func SIMImeiUnlock(ctx command.Context, params *params.ImeiUnlockSIMParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSIMAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("SIMImeiUnlock is failed: %s", e)
	}

	_, err := api.IMEIUnlock(params.Id)
	if err != nil {
		return fmt.Errorf("SIMImeiUnlock is failed: %s", err)
	}
	return nil
}
