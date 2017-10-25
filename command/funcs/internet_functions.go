package funcs

import (
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
)

type subnet struct {
	*sacloud.Subnet
	IPAddressRangeStart string
	IPAddressRangeEnd   string
}

func getSubnetByID(ctx command.Context, subnetID int64) (*subnet, error) {
	client := ctx.GetAPIClient()
	sn, err := client.GetSubnetAPI().Read(subnetID)
	if err != nil {
		return nil, err
	}

	return &subnet{
		Subnet:              sn,
		IPAddressRangeStart: sn.IPAddresses[0].IPAddress,
		IPAddressRangeEnd:   sn.IPAddresses[len(sn.IPAddresses)-1].IPAddress,
	}, nil
}
