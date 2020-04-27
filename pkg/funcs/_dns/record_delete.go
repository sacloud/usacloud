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

	"github.com/sacloud/usacloud/pkg/util"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func RecordDelete(ctx cli.Context, params *params.RecordDeleteDNSParam) error {
	client := sacloud.NewDNSOp(ctx.Client())
	zone, err := client.Read(ctx, params.Id)
	if err != nil {
		return fmt.Errorf("DNSRecordDelete is failed: %s", err)
	}

	if len(zone.Records) == 0 {
		return fmt.Errorf("DNS zone doesn't have any records")
	}

	// validate index
	if err := util.ValidIndex(params.Index, len(zone.Records)); err != nil {
		return nil
	}

	// TODO libsacloud v2で実装されるまで保留
	return nil

	//// delete
	//recordSet := p.Settings.DNS.ResourceRecordSets
	//p.ClearRecords()
	//
	//var targetRecord *sacloud.DNSRecordSet
	//for i, r := range recordSet {
	//	if i == params.Index-1 {
	//		targetRecord = &r
	//	} else {
	//		p.Settings.DNS.ResourceRecordSets = append(p.Settings.DNS.ResourceRecordSets, r)
	//	}
	//}
	//
	//// update
	//p, e = api.Update(params.Id, p)
	//if e != nil {
	//	return fmt.Errorf("DNSRecordDelete is failed: %s", e)
	//}
	//
	//list := []interface{}{}
	//list = append(list, &dnsRecordValueType{
	//	DNSRecordSet: targetRecord,
	//	Index:        params.Index, // for display
	//})
	//
	//return ctx.Output().Print(list...)

}
