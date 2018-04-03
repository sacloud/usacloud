package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func WebAccelRead(ctx command.Context, params *params.ReadWebAccelParam) error {

	client := ctx.GetAPIClient()
	api := client.GetWebAccelAPI()
	p, e := api.Read(fmt.Sprintf("%d", params.Id))
	if e != nil {
		return fmt.Errorf("WebAccelRead is failed: %s", e)
	}

	return ctx.GetOutput().Print(p)
}
