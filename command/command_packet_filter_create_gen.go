// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-funcs'; DO NOT EDIT

package command

import (
	"fmt"
)

func PacketFilterCreate(ctx Context, params *CreatePacketFilterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetPacketFilterAPI()
	p := api.New()

	// set params

	p.SetDescription(params.Description)

	p.SetName(params.Name)

	// call Create(id)
	res, err := api.Create(p)
	if err != nil {
		return fmt.Errorf("PacketFilterCreate is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}
