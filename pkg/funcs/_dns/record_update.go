// Copyright 2017-2020 The Usacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dns

import (
	"fmt"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
	"github.com/sacloud/usacloud/pkg/util"
)

func RecordUpdate(ctx cli.Context, params *params.RecordUpdateDNSParam) error {
	client := sacloud.NewDNSOp(ctx.Client())
	zone, err := client.Read(ctx, params.Id)
	if err != nil {
		return fmt.Errorf("DNSRecordAdd is failed: %s", err)
	}
	if len(zone.Records) == 0 {
		return fmt.Errorf("DNS zone don't have any records")
	}

	// validate index
	if err := util.ValidIndex(params.Index, len(zone.Records)); err != nil {
		return err
	}

	// TODO libsacloud v2実装まで保留
	return nil
	//// set params
	//record := &p.Settings.DNS.ResourceRecordSets[params.Index-1]
	//t := record.Type
	//
	//if ctx.IsSet("type") {
	//	t = params.Type
	//}
	//
	//t = strings.ToUpper(t)
	//switch t {
	//case "MX":
	//	rdata := strings.Split(record.RData, " ")
	//	name := record.Name
	//	mxPriority, _ := strconv.Atoi(rdata[0])
	//	value := ""
	//	if len(rdata) > 1 {
	//		value = rdata[1]
	//	}
	//
	//	ttl := record.TTL
	//
	//	if ctx.IsSet("name") {
	//		name = params.Name
	//	}
	//	if ctx.IsSet("mx-priority") {
	//		mxPriority = params.MxPriority
	//	}
	//	if ctx.IsSet("value") {
	//		value = params.Value
	//	}
	//	if ctx.IsSet("ttl") {
	//		ttl = params.Ttl
	//	}
	//
	//	record = p.CreateNewMXRecord(name, value, ttl, mxPriority)
	//case "SRV":
	//
	//	rdata := strings.Split(record.RData, " ")
	//	priority := 0
	//	weight := 0
	//	port := 0
	//	target := ""
	//
	//	priority, _ = strconv.Atoi(rdata[0])
	//	if len(rdata) > 1 {
	//		weight, _ = strconv.Atoi(rdata[1])
	//	}
	//	if len(rdata) > 2 {
	//		port, _ = strconv.Atoi(rdata[2])
	//	}
	//	if len(rdata) > 3 {
	//		target = rdata[3]
	//	}
	//
	//	name := record.Name
	//	ttl := record.TTL
	//
	//	if ctx.IsSet("srv-priority") {
	//		priority = params.SrvPriority
	//	}
	//	if ctx.IsSet("srv-weight") {
	//		weight = params.SrvWeight
	//	}
	//	if ctx.IsSet("srv-port") {
	//		port = params.SrvPort
	//	}
	//	if ctx.IsSet("srv-target") {
	//		target = params.SrvTarget
	//	}
	//	if ctx.IsSet("name") {
	//		name = params.Name
	//	}
	//	if ctx.IsSet("ttl") {
	//		ttl = params.Ttl
	//	}
	//
	//	record = p.CreateNewSRVRecord(name, target, ttl, priority, weight, port)
	//default:
	//	name := record.Name
	//	value := record.RData
	//	ttl := record.TTL
	//
	//	if ctx.IsSet("name") {
	//		name = params.Name
	//	}
	//	if ctx.IsSet("value") {
	//		value = params.Value
	//	}
	//	if ctx.IsSet("ttl") {
	//		ttl = params.Ttl
	//	}
	//
	//	record = p.CreateNewRecord(name, t, value, ttl)
	//}
	//
	//p.Settings.DNS.ResourceRecordSets[params.Index-1] = *record
	//
	//// update
	//p, e = api.Update(params.Id, p)
	//if e != nil {
	//	return fmt.Errorf("DNSRecordUpdate is failed: %s", e)
	//}
	//
	//list := []interface{}{}
	//list = append(list, &dnsRecordValueType{
	//	DNSRecordSet: record,
	//	Index:        params.Index, // for display
	//})
	//return ctx.Output().Print(list...)
}
