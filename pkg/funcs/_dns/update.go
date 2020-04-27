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
)

func Update(ctx cli.Context, params *params.UpdateDNSParam) error {
	client := sacloud.NewDNSOp(ctx.Client())
	zone, err := client.Read(ctx, params.Id)
	if err != nil {
		return fmt.Errorf("Update is failed: %s", err)
	}

	req := &sacloud.DNSUpdateRequest{
		Description:  zone.Description,
		Tags:         zone.Tags,
		IconID:       zone.IconID,
		Records:      zone.Records,
		SettingsHash: zone.SettingsHash,
	}

	// set params
	if params.Changed("description") {
		req.Description = params.Description
	}
	if params.Changed("tags") {
		req.Tags = params.Tags
	}
	if params.Changed("icon-id") {
		req.IconID = params.IconId
	}

	// call Update(id)
	zone, err = client.Update(ctx, params.Id, req)
	if err != nil {
		return fmt.Errorf("DNSUpdate is failed: %s", err)
	}

	return ctx.Output().Print(zone)
}
