package command

import (
	"fmt"
)

func IconUpdate(ctx Context, params *UpdateIconParam) error {

	client := ctx.GetAPIClient()
	api := client.GetIconAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("IconUpdate is failed: %s", e)
	}

	// set params

	if ctx.IsSet("tags") {
		p.SetTags(params.Tags)
	}
	if ctx.IsSet("name") {
		p.SetName(params.Name)
	}

	// call Update(id)
	res, err := api.Update(params.Id, p)
	if err != nil {
		return fmt.Errorf("IconUpdate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}
