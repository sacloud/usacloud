package command

import (
	"fmt"
	"github.com/sacloud/usacloud/vnc"
)

func ServerVnc(ctx Context, params *VncServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()

	vncProxyInfo, e := api.GetVNCProxy(params.Id)
	if e != nil {
		return fmt.Errorf("ServerVnc is failed: %s", e)
	}

	return vnc.OpenVNCClient(vncProxyInfo)
}
