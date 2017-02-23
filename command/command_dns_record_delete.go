package command

import (
	"fmt"
)

func DNSRecordDelete(ctx Context, params *RecordDeleteDNSParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDNSAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DNSRecordDelete is failed: %s", e)
	}

	if len(p.Settings.DNS.ResourceRecordSets) == 0 {
		return fmt.Errorf("DNS zone don't have any records")
	}

	// validate index
	if params.Index <= 0 || params.Index-1 >= len(p.Settings.DNS.ResourceRecordSets) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	// delete
	recordSet := p.Settings.DNS.ResourceRecordSets
	p.ClearRecords()
	for i, r := range recordSet {
		if i != params.Index-1 {
			p.Settings.DNS.ResourceRecordSets = append(p.Settings.DNS.ResourceRecordSets, r)
		}
	}

	// update
	p, e = api.Update(params.Id, p)
	if e != nil {
		return fmt.Errorf("DNSRecordDelete is failed: %s", e)
	}

	list := []interface{}{}
	for i := range p.Settings.DNS.ResourceRecordSets {
		list = append(list, &p.Settings.DNS.ResourceRecordSets[i])
	}

	return ctx.GetOutput().Print(list...)

}
