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
	for i, r := range p.Settings.DNS.ResourceRecordSets {
		// filtering
		if params.Name != "" && params.Name != r.Name {
			continue
		}
		if params.Type != "" && params.Type != r.Type {
			continue
		}

		list = append(list, &dnsRecordValueType{
			DNSRecordSet: &p.Settings.DNS.ResourceRecordSets[i],
			Index:        i + 1, // for display
		})
	}

	return ctx.GetOutput().Print(list...)
}
