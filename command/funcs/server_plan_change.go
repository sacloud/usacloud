package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"os"
)

func ServerPlanChange(ctx command.Context, params *params.PlanChangeServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerPlanChange is failed: %s", e)
	}

	if !p.IsDown() {
		return fmt.Errorf("ServerPlanChange is failed: %s", "server is running")
	}

	plan, err := client.GetProductServerAPI().GetBySpec(params.Core, params.Memory)
	if err != nil {
		return fmt.Errorf("ServerPlanChange is failed: plan is invalid: %s", err)
	}

	// call manipurate functions
	res, err := api.ChangePlan(params.Id, fmt.Sprintf("%d", plan.ID))
	if err != nil {
		return fmt.Errorf("ServerPlanChange is failed: %s", err)
	}

	// if exists ssh private-key file on default location, rename it.
	beforeKeyPath, err := getSSHPrivateKeyStorePath(p.ID)
	if err != nil {
		return fmt.Errorf("ServerPlanChange is failed: %s", err)
	}
	if _, err := os.Stat(beforeKeyPath); err == nil {
		afterKeyPath, err := getSSHPrivateKeyStorePath(res.ID)
		if err != nil {
			return fmt.Errorf("ServerPlanChange is failed: %s", err)
		}
		err = os.Rename(beforeKeyPath, afterKeyPath)
		if err != nil {
			return fmt.Errorf("ServerPlanChange is failed: %s", err)
		}
	}

	return ctx.GetOutput().Print(res)

}
