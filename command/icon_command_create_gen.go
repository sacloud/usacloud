package command

import (
	"fmt"
)

func IconCreate(ctx Context, params *CreateIconParam) error {

	client := ctx.GetAPIClient()
	api := client.GetIconAPI()
	p := api.New()

	// set params

	p.SetName(params.Name)

	p.SetTags(params.Tags)

	params.getCommandDef().Params["image"].CustomHandler("Image", params, p)

	// call Create(id)
	res, err := api.Create(p)
	if err != nil {
		return fmt.Errorf("IconCreate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}
