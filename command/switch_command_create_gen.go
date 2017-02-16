package command

import (
	"fmt"
)

func SwitchCreate(ctx Context, params *CreateSwitchParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSwitchAPI()
	p := api.New()

	// set params

	p.SetName(params.Name)

	p.SetDescription(params.Description)

	p.SetTags(params.Tags)

	p.SetIconByID(params.Icon)

	// call Create(id)
	res, err := api.Create(p)
	if err != nil {
		return fmt.Errorf("SwitchCreate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}
