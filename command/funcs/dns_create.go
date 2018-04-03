package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func DNSCreate(ctx command.Context, params *params.CreateDNSParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDNSAPI()
	p := api.New(params.Name)

	// set params
	p.SetTags(params.Tags)
	p.SetIconByID(params.IconId)
	p.SetName(params.Name)
	p.SetDescription(params.Description)

	// call Create(id)
	res, err := api.Create(p)
	if err != nil {
		return fmt.Errorf("DNSCreate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}
