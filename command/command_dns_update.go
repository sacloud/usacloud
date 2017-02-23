package command

import (
	"fmt"
)

func DNSUpdate(ctx Context, params *UpdateDNSParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDNSAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DNSUpdate is failed: %s", e)
	}

	// set params

	if ctx.IsSet("description") {
		p.SetDescription(params.Description)
	}
	if ctx.IsSet("tags") {
		p.SetTags(params.Tags)
	}
	if ctx.IsSet("icon-id") {
		p.SetIconByID(params.IconId)
	}

	// call Update(id)
	res, err := api.Update(params.Id, p)
	if err != nil {
		return fmt.Errorf("DNSUpdate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}
