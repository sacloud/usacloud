package funcs

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func DNSRecordUpdate(ctx command.Context, params *params.RecordUpdateDNSParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDNSAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DNSRecordUpdate is failed: %s", e)
	}
	if len(p.Settings.DNS.ResourceRecordSets) == 0 {
		return fmt.Errorf("DNS zone don't have any records")
	}

	// validate index
	if params.Index <= 0 || params.Index-1 >= len(p.Settings.DNS.ResourceRecordSets) {
		return fmt.Errorf("index(%d) is out of range", params.Index)
	}

	// set params
	record := &p.Settings.DNS.ResourceRecordSets[params.Index-1]
	t := record.Type

	if ctx.IsSet("type") {
		t = params.Type
	}

	t = strings.ToUpper(t)
	switch t {
	case "MX":
		rdata := strings.Split(record.RData, " ")
		name := record.Name
		mxPriority, _ := strconv.Atoi(rdata[0])
		value := ""
		if len(rdata) > 1 {
			value = rdata[1]
		}

		ttl := record.TTL

		if ctx.IsSet("name") {
			name = params.Name
		}
		if ctx.IsSet("mx-priority") {
			mxPriority = params.MxPriority
		}
		if ctx.IsSet("value") {
			value = params.Value
		}
		if ctx.IsSet("ttl") {
			ttl = params.Ttl
		}

		record = p.CreateNewMXRecord(name, value, ttl, mxPriority)
	case "SRV":

		rdata := strings.Split(record.RData, " ")
		priority := 0
		weight := 0
		port := 0
		target := ""

		priority, _ = strconv.Atoi(rdata[0])
		if len(rdata) > 1 {
			weight, _ = strconv.Atoi(rdata[1])
		}
		if len(rdata) > 2 {
			port, _ = strconv.Atoi(rdata[2])
		}
		if len(rdata) > 3 {
			target = rdata[3]
		}

		name := record.Name
		ttl := record.TTL

		if ctx.IsSet("srv-priority") {
			priority = params.SrvPriority
		}
		if ctx.IsSet("srv-weight") {
			weight = params.SrvWeight
		}
		if ctx.IsSet("srv-port") {
			port = params.SrvPort
		}
		if ctx.IsSet("srv-target") {
			target = params.SrvTarget
		}
		if ctx.IsSet("name") {
			name = params.Name
		}
		if ctx.IsSet("ttl") {
			ttl = params.Ttl
		}

		record = p.CreateNewSRVRecord(name, target, ttl, priority, weight, port)
	default:
		name := record.Name
		value := record.RData
		ttl := record.TTL

		if ctx.IsSet("name") {
			name = params.Name
		}
		if ctx.IsSet("value") {
			value = params.Value
		}
		if ctx.IsSet("ttl") {
			ttl = params.Ttl
		}

		record = p.CreateNewRecord(name, t, value, ttl)
	}

	p.Settings.DNS.ResourceRecordSets[params.Index-1] = *record

	// update
	p, e = api.Update(params.Id, p)
	if e != nil {
		return fmt.Errorf("DNSRecordUpdate is failed: %s", e)
	}

	list := []interface{}{}
	list = append(list, &dnsRecordValueType{
		DNSRecordSet: record,
		Index:        params.Index, // for display
	})
	return ctx.GetOutput().Print(list...)
}
