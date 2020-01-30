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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func DNSRecordBulkUpdate(ctx command.Context, params *params.RecordBulkUpdateDNSParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDNSAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DNSRecordBulkUpdate is failed: %s", e)
	}

	// validate
	buf, err := ioutil.ReadFile(params.File)
	if err != nil {
		return fmt.Errorf("DNSRecordBulkUpdate is failed: %s", err)
	}

	var records dnsRecordsType
	if err := json.Unmarshal(buf, &records); err != nil {
		return fmt.Errorf("DNSRecordBulkUpdate is failed: %s", err)
	}

	if len(records) == 0 {
		return fmt.Errorf("file %q don't have any records", params.File)
	}

	// set params
	switch params.Mode {
	case "upsert-only":
		indexedRecords := map[int]*sacloud.DNSRecordSet{}
		if p.HasDNSRecord() {
			for i := range p.Settings.DNS.ResourceRecordSets {
				indexedRecords[i+1] = &p.Settings.DNS.ResourceRecordSets[i]
			}
		}

		for _, record := range records {
			if record.Index == 0 {
				continue
			}
			if dest, ok := indexedRecords[record.Index]; ok {
				dest.Name = record.Name
				dest.Type = record.Type
				dest.RData = record.RData
				dest.TTL = record.TTL
			} else {
				indexedRecords[record.Index] = &sacloud.DNSRecordSet{
					Name:  record.Name,
					Type:  record.Type,
					RData: record.RData,
					TTL:   record.TTL,
				}
			}
		}

		var keys []int
		for k := range indexedRecords {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		v := make([]sacloud.DNSRecordSet, 0, len(indexedRecords))
		for _, key := range keys {
			v = append(v, *indexedRecords[key])
		}
		p.Settings.DNS.ResourceRecordSets = v

	case "sync":
		p.ClearRecords()
		sort.Slice(records, func(i, j int) bool { return records[i].Index < records[j].Index })
		for _, record := range records {
			p.AddRecord(&sacloud.DNSRecordSet{
				Name:  record.Name,
				Type:  record.Type,
				RData: record.RData,
				TTL:   record.TTL,
			})
		}
	default:
		return fmt.Errorf("Invalid Mode: %q", params.Mode)
	}

	// update records
	p, e = api.Update(params.Id, p)
	if e != nil {
		return fmt.Errorf("DNSRecordBulkUpdate is failed: %s", e)
	}

	return nil
}
