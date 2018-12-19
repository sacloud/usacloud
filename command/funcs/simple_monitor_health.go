package funcs

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func SimpleMonitorHealth(ctx command.Context, params *params.HealthSimpleMonitorParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSimpleMonitorAPI()

	res, err := api.Health(params.Id)
	if err != nil {
		return fmt.Errorf("SimpleMonitorHealth is failed: %s", err)
	}

	output := struct {
		ID int64
		*sacloud.SimpleMonitorHealthCheckStatus
	}{
		ID:                             params.Id,
		SimpleMonitorHealthCheckStatus: res,
	}

	return ctx.GetOutput().Print(output)

}
