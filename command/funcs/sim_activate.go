package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func SIMActivate(ctx command.Context, params *params.ActivateSIMParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSIMAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("SIMActivate is failed: %s", e)
	}

	_, err := api.Activate(params.Id)
	if err != nil {
		return fmt.Errorf("SIMActivate is failed: %s", err)
	}
	return nil
}
