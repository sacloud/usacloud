package command

import (
	"fmt"
)

func DNSRecordInfo(ctx Context, params *RecordInfoDNSParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDNSAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DNSRecordList is failed: %s", e)
	}

	if len(p.Settings.DNS.ResourceRecordSets) == 0 {
		fmt.Fprintf(GlobalOption.Err, "DNS zone don't have any records\n")
		return nil
	}

	list := []interface{}{}
	for i := range p.Settings.DNS.ResourceRecordSets {
		list = append(list, &p.Settings.DNS.ResourceRecordSets[i])
	}

	return ctx.GetOutput().Print(list...)
}
