package funcs

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func DNSRecordInfo(ctx command.Context, params *params.RecordInfoDNSParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDNSAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DNSRecordList is failed: %s", e)
	}

	if len(p.Settings.DNS.ResourceRecordSets) == 0 {
		fmt.Fprintf(command.GlobalOption.Err, "DNS zone don't have any records\n")
		return nil
	}

	list := []interface{}{}
	for i := range p.Settings.DNS.ResourceRecordSets {
		list = append(list, &p.Settings.DNS.ResourceRecordSets[i])
	}

	return ctx.GetOutput().Print(list...)
}
