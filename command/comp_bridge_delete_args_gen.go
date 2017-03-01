// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
)

func BridgeDeleteCompleteArgs(ctx Context, params *DeleteBridgeParam) {

	if !GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetBridgeAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}
	for i := range res.Bridges {
		fmt.Println(res.Bridges[i].ID)
	}

}
