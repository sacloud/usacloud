package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func PrivateHostCreate(ctx command.Context, params *params.CreatePrivateHostParam) error {

	client := ctx.GetAPIClient()
	api := client.GetPrivateHostAPI()
	p := api.New()

	// set params

	p.SetName(params.Name)
	p.SetDescription(params.Description)
	p.SetTags(params.Tags)
	p.SetIconByID(params.IconId)

	// set plan(There have only one plan now)
	plans, err := client.Product.GetProductPrivateHostAPI().Find()
	if err != nil || len(plans.PrivateHostPlans) == 0 {
		return fmt.Errorf("PrivateHostCreate is failed: can't find any private-host plan %s", err)
	}
	p.SetPrivateHostPlanByID(plans.PrivateHostPlans[0].ID)

	// call Create(id)
	res, err := api.Create(p)
	if err != nil {
		return fmt.Errorf("PrivateHostCreate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}
