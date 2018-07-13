package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ServerUpdate(ctx command.Context, params *params.UpdateServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerUpdate is failed: %s", e)
	}

	// validate
	if ctx.IsSet("interface-driver") && !p.IsDown() {
		return fmt.Errorf("ServerUpdate is failed: %s", "server is running")
	}

	// set params
	if ctx.IsSet("name") {
		p.SetName(params.Name)
	}
	if ctx.IsSet("description") {
		p.SetDescription(params.Description)
	}
	if ctx.IsSet("tags") {
		p.SetTags(params.Tags)
	}
	if ctx.IsSet("icon-id") {
		p.SetIconByID(params.IconId)
	}
	if ctx.IsSet("interface-driver") {
		p.SetInterfaceDriverByString(params.InterfaceDriver)
	}

	// call Update(id)
	res, err := api.Update(params.Id, p)
	if err != nil {
		return fmt.Errorf("ServerUpdate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}
