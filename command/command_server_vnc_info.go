package command

import (
	"fmt"
)

func ServerVncInfo(ctx Context, params *VncInfoServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()

	vncProxyInfo, e := api.GetVNCProxy(params.Id)
	if e != nil {
		return fmt.Errorf("ServerVncInfo is failed: %s", e)
	}

	return ctx.GetOutput().Print(vncProxyInfo)

}
