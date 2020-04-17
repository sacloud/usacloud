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

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func DNSRecordDelete(ctx cli.Context, params *params.RecordDeleteDNSParam) error {

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
