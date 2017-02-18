package command

import (
	"fmt"
	"github.com/sacloud/libsacloud/sacloud"
)

func ServerIsoEject(ctx Context, params *IsoEjectServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerIsoEject is failed: %s", e)
	}

	if p.Instance.CDROM == nil || p.Instance.CDROM.ID == sacloud.EmptyID {
		fmt.Fprintf(GlobalOption.Err, "ISOImage isnot inserted to server\n")
		return nil
	}

	// call manipurate functions
	_, err := api.EjectCDROM(params.Id, p.Instance.CDROM.ID)
	if err != nil {
		return fmt.Errorf("ServerIsoEject is failed: %s", err)
	}

	return nil
	// return ctx.GetOutput().Print(res)
}
