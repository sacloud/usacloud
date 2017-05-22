package command

import (
	"fmt"
	"github.com/sacloud/usacloud/command/internal"
)

func VPCRouterCreate(ctx Context, params *CreateVPCRouterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetVPCRouterAPI()
	p := api.New()

	// validate
	isStandard := params.Plan == "standard"
	if !isStandard {
		targets := []string{"switch-id", "vip", "ipaddress1", "ipaddress2"}
		for _, t := range targets {
			if !ctx.IsSet(t) {
				return fmt.Errorf("%q: is required when plan is [premium/highspec]", t)
			}
		}
	}

	// set params
	switch params.Plan {
	case "standard":
		p.SetStandardPlan()
	case "premium":
		p.SetPremiumPlan(
			fmt.Sprintf("%d", params.SwitchId),
			params.Vip,
			params.Ipaddress1,
			params.Ipaddress2,
			params.Vrid,
			[]string{},
		)
	case "highspec":
		p.SetHighSpecPlan(
			fmt.Sprintf("%d", params.SwitchId),
			params.Vip,
			params.Ipaddress1,
			params.Ipaddress2,
			params.Vrid,
			[]string{},
		)
	}

	p.SetName(params.Name)
	p.SetDescription(params.Description)
	p.SetTags(params.Tags)
	p.SetIconByID(params.IconId)

	// call Create(id)
	res, err := api.Create(p)
	if err != nil {
		return fmt.Errorf("VPCRouterCreate is failed: %s", err)
	}
	// wait for boot
	err = internal.ExecWithProgress(
		fmt.Sprintf("Still creating[ID:%d]...", res.ID),
		fmt.Sprintf("Create vpc-router[ID:%d]", res.ID),
		GlobalOption.Progress,
		func(compChan chan bool, errChan chan error) {
			// call manipurate functions
			err := api.SleepWhileCopying(res.ID, client.DefaultTimeoutDuration, 20)
			if err != nil {
				errChan <- err
				return
			}
			if params.BootAfterCreate {
				_, err := api.Boot(res.ID)
				if err != nil {
					errChan <- err
					return
				}
				err = api.SleepUntilUp(res.ID, client.DefaultTimeoutDuration)
				if err != nil {
					errChan <- err
					return
				}
			}

			compChan <- true
		},
	)
	if err != nil {
		return fmt.Errorf("VPCRouterCreate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)
}
