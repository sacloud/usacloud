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
	"strings"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/params"
)

func RecordInfo(ctx cli.Context, params *params.RecordInfoDNSParam) error {
	client := sacloud.NewDNSOp(ctx.Client())
	zone, err := client.Read(ctx, params.Id)
	if err != nil {
		return fmt.Errorf("DNSRecordAdd is failed: %s", err)
	}

	if len(zone.Records) == 0 {
		fmt.Fprintf(ctx.IO().Err(), "DNS zone don't have any records\n")
		return nil
	}

	var list []interface{}
	for i, r := range zone.Records {
		// filtering
		if params.Name != "" && params.Name != r.Name {
			continue
		}
		if params.Type != "" && strings.ToUpper(params.Type) != r.Type.String() {
			continue
		}

		list = append(list, &dnsRecordValueType{
			DNSRecord: zone.Records[i],
			Index:     i + 1, // for display
		})
	}

	return ctx.Output().Print(list...)
}
