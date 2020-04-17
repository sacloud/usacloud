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

	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func DNSRecordInfo(ctx cli.Context, params *params.RecordInfoDNSParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDNSAPI()
	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("DNSRecordList is failed: %s", e)
	}

	if len(p.Settings.DNS.ResourceRecordSets) == 0 {
		fmt.Fprintf(ctx.IO().Err(), "DNS zone don't have any records\n")
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
