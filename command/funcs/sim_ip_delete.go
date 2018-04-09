package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func SIMIpDelete(ctx command.Context, params *params.IpDeleteSIMParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSIMAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("SIMIpDelete is failed: %s", e)
	}

	_, err := api.ClearIP(params.Id)
	if err != nil {
		return fmt.Errorf("SIMIpDelete is failed: %s", err)
	}
	return nil
}
