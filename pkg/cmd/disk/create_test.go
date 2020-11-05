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

package disk

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud/ostype"

	"github.com/sacloud/usacloud/pkg/cmd/conv"

	"github.com/sacloud/libsacloud/v2/sacloud/types"

	"github.com/sacloud/libsacloud/v2/helper/service/disk"
	"github.com/stretchr/testify/require"
)

func TestCreate_ConvertToServiceRequest(t *testing.T) {
	t.Run("full", func(t *testing.T) {
		in := &CreateParameter{
			Zone:        "is1a",
			Name:        "name",
			Description: "desc",
			Tags:        []string{"tag1", "tag2"},
			IconID:      types.ID(1),
			DiskPlan:    "ssd",
			Connection:  types.DiskConnections.VirtIO.String(),
			SizeGB:      20,
			OSType:      "ubuntu",
		}

		out := &disk.CreateRequest{}
		if err := conv.ConvertTo(in, out); err != nil {
			t.Fatal(err)
		}

		require.EqualValues(t, &disk.CreateRequest{
			Zone:        "is1a",
			Name:        "name",
			Description: "desc",
			Tags:        types.Tags{"tag1", "tag2"},
			IconID:      types.ID(1),
			DiskPlanID:  types.DiskPlans.SSD,
			Connection:  types.DiskConnections.VirtIO,
			SizeGB:      20,
			OSType:      ostype.Ubuntu,
		}, out)
	})
}
