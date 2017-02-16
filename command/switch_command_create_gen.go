// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-funcs'; DO NOT EDIT

package command

import (
	"fmt"
)

func SwitchCreate(ctx Context, params *CreateSwitchParam) error {

	client := ctx.GetAPIClient()
	api := client.GetSwitchAPI()
	p := api.New()

	// set params

	p.SetDescription(params.Description)

	p.SetTags(params.Tags)

	p.SetIconByID(params.Icon)

	p.SetName(params.Name)

	// call Create(id)
	res, err := api.Create(p)
	if err != nil {
		return fmt.Errorf("SwitchCreate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}
