// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-funcs'; DO NOT EDIT

package command

import (
	"fmt"
)

func ProductDiskRead(ctx Context, params *ReadProductDiskParam) error {

	client := ctx.GetAPIClient()
	api := client.GetProductDiskAPI()

	// set params

	// call Read(id)
	res, err := api.Read(params.Id)
	if err != nil {
		return fmt.Errorf("ProductDiskRead is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}
