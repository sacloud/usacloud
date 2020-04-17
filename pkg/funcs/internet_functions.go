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
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
)

type subnet struct {
	*sacloud.Subnet
	IPAddressRangeStart string
	IPAddressRangeEnd   string
}

func getSubnetByID(ctx cli.Context, subnetID sacloud.ID) (*subnet, error) {
	client := ctx.GetAPIClient()
	sn, err := client.GetSubnetAPI().Read(subnetID)
	if err != nil {
		return nil, err
	}

	return &subnet{
		Subnet:              sn,
		IPAddressRangeStart: sn.IPAddresses[0].IPAddress,
		IPAddressRangeEnd:   sn.IPAddresses[len(sn.IPAddresses)-1].IPAddress,
	}, nil
}
