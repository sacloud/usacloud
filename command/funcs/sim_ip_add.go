package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func SIMIpAdd(ctx command.Context, params *params.IpAddSIMParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSIMAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("SIMIpAdd is failed: %s", e)
	}

	_, err := api.AssignIP(params.Id, params.Ip)
	if err != nil {
		return fmt.Errorf("SIMIpAdd is failed: %s", err)
	}
	return nil
}
