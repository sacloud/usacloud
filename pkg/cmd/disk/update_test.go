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

	"github.com/sacloud/usacloud/pkg/cmd/conv"

	"github.com/sacloud/libsacloud/v2/sacloud/pointer"
	"github.com/sacloud/libsacloud/v2/sacloud/types"

	"github.com/sacloud/libsacloud/v2/helper/service/disk"
	"github.com/stretchr/testify/require"
)

func TestUpdate_ConvertToServiceRequest(t *testing.T) {
	t.Run("full", func(t *testing.T) {
		in := &UpdateParameter{
			OutputParameter: nil,
			Zone:            "is1a",
			ID:              types.ID(1),
			Name:            pointer.NewString("name"),
			Description:     pointer.NewString("desc"),
			Tags:            pointer.NewStringSlice([]string{"tag1", "tag2"}),
			IconID:          pointer.NewID(types.ID(2)),
			Connection:      pointer.NewString(types.DiskConnections.VirtIO.String()),
		}

		out := &disk.UpdateRequest{}
		if err := conv.ConvertTo(in, out); err != nil {
			t.Fatal(err)
		}

		require.EqualValues(t, &disk.UpdateRequest{
			Zone:        "is1a",
			ID:          types.ID(1),
			Name:        pointer.NewString("name"),
			Description: pointer.NewString("desc"),
			Tags:        pointer.NewTags(types.Tags{"tag1", "tag2"}),
			IconID:      pointer.NewID(types.ID(2)),
			Connection:  &types.DiskConnections.VirtIO,
		}, out)
	})

	t.Run("nil", func(t *testing.T) {
		in := &UpdateParameter{
			OutputParameter: nil,
			Zone:            "is1a",
			ID:              types.ID(1),
			Name:            pointer.NewString("name"),
		}

		out := &disk.UpdateRequest{}
		if err := conv.ConvertTo(in, out); err != nil {
			t.Fatal(err)
		}

		require.EqualValues(t, &disk.UpdateRequest{
			Zone:        "is1a",
			ID:          types.ID(1),
			Name:        pointer.NewString("name"),
			Description: nil,
			Tags:        nil,
			IconID:      nil,
			Connection:  nil,
		}, out)
	})
	t.Run("empty", func(t *testing.T) {
		in := &UpdateParameter{
			OutputParameter: nil,
			Zone:            "is1a",
			ID:              types.ID(1),
			Name:            pointer.NewString("name"),
			Description:     pointer.NewString(""),
			Tags:            pointer.NewStringSlice([]string{}),
			IconID:          pointer.NewID(types.ID(0)),
		}

		out := &disk.UpdateRequest{}
		if err := conv.ConvertTo(in, out); err != nil {
			t.Fatal(err)
		}

		require.EqualValues(t, &disk.UpdateRequest{
			Zone:        "is1a",
			ID:          types.ID(1),
			Name:        pointer.NewString("name"),
			Description: pointer.NewString(""),
			Tags:        pointer.NewTags(types.Tags{}),
			IconID:      pointer.NewID(0),
			Connection:  nil,
		}, out)
	})
}
