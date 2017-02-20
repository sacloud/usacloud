// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-funcs'; DO NOT EDIT

package command

import (
	"fmt"
)

func PacketFilterDelete(ctx Context, params *DeletePacketFilterParam) error {

	client := ctx.GetAPIClient()
	api := client.GetPacketFilterAPI()

	// set params

	// call Delete(id)
	res, err := api.Delete(params.Id)
	if err != nil {
		return fmt.Errorf("PacketFilterDelete is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}
