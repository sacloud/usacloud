// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
)

func InterfacePacketFilterConnectCompleteArgs(ctx Context, params *PacketFilterConnectInterfaceParam) {

	if !GlobalOption.Valid {
		return
	}

	client := ctx.GetAPIClient()
	finder := client.GetInterfaceAPI()
	finder.SetEmpty()

	// call Find()
	res, err := finder.Find()
	if err != nil {
		return
	}
	for i := range res.Interfaces {
		fmt.Println(res.Interfaces[i].ID)
	}

}
