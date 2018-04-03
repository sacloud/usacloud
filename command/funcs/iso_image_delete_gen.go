// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-funcs'; DO NOT EDIT

package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ISOImageDelete(ctx command.Context, params *params.DeleteISOImageParam) error {

	client := ctx.GetAPIClient()
	api := client.GetCDROMAPI()

	// set params

	// call Delete(id)
	res, err := api.Delete(params.Id)
	if err != nil {
		return fmt.Errorf("ISOImageDelete is failed: %s", err)
	}

	return ctx.GetOutput().Print(res)

}
