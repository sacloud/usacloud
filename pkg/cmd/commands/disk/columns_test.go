// Copyright 2017-2021 The Usacloud Authors
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

package disk

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/pkg/size"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/test"
)

func TestDiskDefaultColumnDefs(t *testing.T) {
	cases := []struct {
		in    *sacloud.Disk
		tests []*test.ColumnDefTestTarget
	}{
		{
			in: &sacloud.Disk{
				ID:           1,
				Name:         "disk",
				Description:  "desc",
				Tags:         types.Tags{"tag1", "tag2"},
				Availability: types.Availabilities.Available,
				Connection:   types.DiskConnections.VirtIO,
				SizeMB:       20 * size.GiB,
				DiskPlanID:   types.DiskPlans.SSD,
				DiskPlanName: "SSDプラン",
				Storage:      &sacloud.Storage{Name: "dummy"},
				ServerID:     2,
				ServerName:   "server",
			},
			tests: []*test.ColumnDefTestTarget{
				{
					ColumnName: "ID",
					Expect:     "1",
				},
				{
					ColumnName: "Plan",
					Expect:     "ssd",
				},

				{
					ColumnName: "Size",
					Expect:     "20GB",
				},
				{
					ColumnName: "Server",
					Expect:     "2(server)",
				},
			},
		},
	}

	for _, tc := range cases {
		test.RunColumnDefTest(t, test.ColumnDefTestTargets{
			ColumnDefs: defaultColumnDefs,
			Source:     tc.in,
			Tests:      tc.tests,
		})
	}
}
