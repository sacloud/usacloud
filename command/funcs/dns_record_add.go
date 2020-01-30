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

package funcs

import (
	"fmt"
	"strings"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func DNSRecordAdd(ctx command.Context, params *params.RecordAddDNSParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDNSAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DNSRecordAdd is failed: %s", e)
	}

	// validate maxlen
	if len(p.Settings.DNS.ResourceRecordSets) == 1000 {
		return fmt.Errorf("DNS already have max(1000) records")
	}

	// validate per types
	var record *sacloud.DNSRecordSet
	t := strings.ToUpper(params.Type)
	switch t {
	case "MX":
		record = p.CreateNewMXRecord(params.Name, params.Value, params.Ttl, params.MxPriority)
	case "SRV":
		record = p.CreateNewSRVRecord(params.Name, params.SrvTarget, params.Ttl, params.SrvPriority, params.SrvWeight, params.SrvPort)
	default:
		record = p.CreateNewRecord(params.Name, t, params.Value, params.Ttl)
	}

	// save current index
	index := len(p.Settings.DNS.ResourceRecordSets)

	p.AddRecord(record)

	// update
	p, e = api.Update(params.Id, p)
	if e != nil {
		return fmt.Errorf("DNSRecordAdd is failed: %s", e)
	}

	list := []interface{}{}
	list = append(list, &dnsRecordValueType{
		DNSRecordSet: &p.Settings.DNS.ResourceRecordSets[index],
		Index:        index + 1, // for display
	})
	return ctx.GetOutput().Print(list...)

}
