package funcs

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ProxyLBPlanChange(ctx command.Context, params *params.PlanChangeProxyLBParam) error {

	client := ctx.GetAPIClient()
	api := client.GetProxyLBAPI()
	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ProxyLBPlanChange is failed: %s", e)
	}

	// call manipurate functions
	res, err := api.ChangePlan(params.Id, sacloud.ProxyLBPlan(params.Plan))
	if err != nil {
		return fmt.Errorf("ProxyLBPlanChange is failed: %s", err)
	}
	return ctx.GetOutput().Print(res)
}
