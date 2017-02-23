// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-funcs'; DO NOT EDIT

package command

import (
	"fmt"
)

func RegionRead(ctx Context, params *ReadRegionParam) error {

	client := ctx.GetAPIClient()
	api := client.GetRegionAPI()

	// set params

	// call Read(id)
	res, err := api.Read(params.Id)
	if err != nil {
		return fmt.Errorf("RegionRead is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}
