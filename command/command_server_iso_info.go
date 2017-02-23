package command

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
)

func ServerIsoInfo(ctx Context, params *IsoInfoServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerIsoInfo is failed: %s", e)
	}

	if p.Instance.CDROM == nil || p.Instance.CDROM.ID == sacloud.EmptyID {
		fmt.Fprintf(GlobalOption.Err, "ISOImage isnot inserted to server\n")
		return nil
	}

	return ctx.GetOutput().Print(p.Instance.CDROM)

}
