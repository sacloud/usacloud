package command

import (
	"fmt"
)

func InternetCreate(ctx Context, params *CreateInternetParam) error {

	client := ctx.GetAPIClient()
	api := client.GetInternetAPI()
	p := api.New()

	// set params

	p.SetTags(params.Tags)

	p.SetIconByID(params.Icon)

	p.SetNetworkMaskLen(params.NwMasklen)

	p.SetName(params.Name)

	p.SetDescription(params.Description)

	// call Create(id)
	res, err := api.Create(p)
	if err != nil {
		return fmt.Errorf("InternetCreate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}
