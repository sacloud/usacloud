package command

import (
	"fmt"
)

func ISOImageUpdate(ctx Context, params *UpdateISOImageParam) error {

	client := ctx.GetAPIClient()
	api := client.GetCDROMAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ISOImageUpdate is failed: %s", e)
	}

	// set params

	if ctx.IsSet("icon") {
		p.SetIconByID(params.Icon)
	}
	if ctx.IsSet("name") {
		p.SetName(params.Name)
	}
	if ctx.IsSet("description") {
		p.SetDescription(params.Description)
	}
	if ctx.IsSet("tags") {
		p.SetTags(params.Tags)
	}

	// call Update(id)
	res, err := api.Update(params.Id, p)
	if err != nil {
		return fmt.Errorf("ISOImageUpdate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}
