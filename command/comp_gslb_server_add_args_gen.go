// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
)

func GSLBServerAddCompleteArgs(ctx Context, params *ServerAddGSLBParam) {

	if !GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetGSLBAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}
	for i := range res.CommonServiceGSLBItems {
		fmt.Println(res.CommonServiceGSLBItems[i].ID)
	}

}
