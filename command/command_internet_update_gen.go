// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-funcs'; DO NOT EDIT

package command

import (
	"fmt"
)

func InternetUpdate(ctx Context, params *UpdateInternetParam) error {

	client := ctx.GetAPIClient()
	api := client.GetInternetAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("InternetUpdate is failed: %s", e)
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
	if ctx.IsSet("band-width") {
		p.SetBandWidthMbps(params.BandWidth)
	}
	if ctx.IsSet("name") {
		p.SetName(params.Name)
	}

	// call Update(id)
	res, err := api.Update(params.Id, p)
	if err != nil {
		return fmt.Errorf("InternetUpdate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}
