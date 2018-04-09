package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func SIMDeactivate(ctx command.Context, params *params.DeactivateSIMParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSIMAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("SIMDeactivate is failed: %s", e)
	}

	_, err := api.Deactivate(params.Id)
	if err != nil {
		return fmt.Errorf("SIMDeactivate is failed: %s", err)
	}
	return nil
}
