package funcs

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func DNSRecordDelete(ctx command.Context, params *params.RecordDeleteDNSParam) error {

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

	var targetRecord *sacloud.DNSRecordSet
	for i, r := range recordSet {
		if i == params.Index-1 {
			targetRecord = &r
		} else {
			p.Settings.DNS.ResourceRecordSets = append(p.Settings.DNS.ResourceRecordSets, r)
		}
	}

	// update
	p, e = api.Update(params.Id, p)
	if e != nil {
		return fmt.Errorf("DNSRecordDelete is failed: %s", e)
	}

	list := []interface{}{}
	list = append(list, &dnsRecordValueType{
		DNSRecordSet: targetRecord,
		Index:        params.Index, // for display
	})

	return ctx.GetOutput().Print(list...)

}
