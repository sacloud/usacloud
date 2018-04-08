package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func SIMImeiLock(ctx command.Context, params *params.ImeiLockSIMParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSIMAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("SIMImeiLock is failed: %s", e)
	}

	_, err := api.IMEILock(params.Id, params.Imei)
	if err != nil {
		return fmt.Errorf("SIMImeiLock is failed: %s", err)
	}
	return nil
}
