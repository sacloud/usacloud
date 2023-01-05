// Copyright 2017-2023 The sacloud/usacloud Authors
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
	"errors"
	"strings"
	"testing"

	"github.com/sacloud/iaas-api-go/types"
	"github.com/sacloud/iaas-service-go/disk"
	"github.com/sacloud/packages-go/pointer"
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/conv"
	"github.com/sacloud/usacloud/pkg/validate"
	"github.com/stretchr/testify/require"
)

func TestUpdate_ConvertToServiceRequest(t *testing.T) {
	id0 := types.ID(0)
	id2 := types.ID(2)
	t.Run("full", func(t *testing.T) {
		in := &updateParameter{
			ZoneParameter:         cflag.ZoneParameter{Zone: "is1a"},
			IDParameter:           cflag.IDParameter{ID: "1"},
			NameUpdateParameter:   cflag.NameUpdateParameter{Name: pointer.NewString("name")},
			DescUpdateParameter:   cflag.DescUpdateParameter{Description: pointer.NewString("desc")},
			TagsUpdateParameter:   cflag.TagsUpdateParameter{Tags: pointer.NewStringSlice([]string{"tag1", "tag2"})},
			IconIDUpdateParameter: cflag.IconIDUpdateParameter{IconID: &id2},
			Connection:            pointer.NewString(types.DiskConnections.VirtIO.String()),
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
			Tags:        &types.Tags{"tag1", "tag2"},
			IconID:      &id2,
			Connection:  &types.DiskConnections.VirtIO,
		}, out)
	})

	t.Run("nil", func(t *testing.T) {
		in := &updateParameter{
			ZoneParameter:       cflag.ZoneParameter{Zone: "is1a"},
			IDParameter:         cflag.IDParameter{ID: "1"},
			NameUpdateParameter: cflag.NameUpdateParameter{Name: pointer.NewString("name")},
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
		in := &updateParameter{
			ZoneParameter:         cflag.ZoneParameter{Zone: "is1a"},
			IDParameter:           cflag.IDParameter{ID: "1"},
			NameUpdateParameter:   cflag.NameUpdateParameter{Name: pointer.NewString("name")},
			DescUpdateParameter:   cflag.DescUpdateParameter{Description: pointer.NewString("")},
			TagsUpdateParameter:   cflag.TagsUpdateParameter{Tags: pointer.NewStringSlice([]string{})},
			IconIDUpdateParameter: cflag.IconIDUpdateParameter{IconID: &id0},
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
			Tags:        &types.Tags{},
			IconID:      &id0,
			Connection:  nil,
		}, out)
	})
}

func TestUpdateParameter_Validate(t *testing.T) {
	cases := []struct {
		in  *updateParameter
		err error
	}{
		// default
		{
			in: newUpdateParameter(),
			err: errors.New(strings.Join([]string{
				"validation error:",
				"\t--zone: required",
			}, "\n")),
		},
		// valid
		{
			in: &updateParameter{
				ZoneParameter:       cflag.ZoneParameter{Zone: "is1a"},
				IDParameter:         cflag.IDParameter{ID: "1"},
				NameUpdateParameter: cflag.NameUpdateParameter{Name: pointer.NewString("1")},
				DescUpdateParameter: cflag.DescUpdateParameter{Description: pointer.NewString("")},
				TagsUpdateParameter: cflag.TagsUpdateParameter{Tags: pointer.NewStringSlice([]string{})},
				Connection:          pointer.NewString("ide"),
			},
			err: nil,
		},
		// name
		{
			in: &updateParameter{
				ZoneParameter:       cflag.ZoneParameter{Zone: "is1a"},
				IDParameter:         cflag.IDParameter{ID: "1"},
				NameUpdateParameter: cflag.NameUpdateParameter{Name: pointer.NewString("")},
			},
			err: errors.New(strings.Join([]string{
				"validation error:",
				"\t--name: min=1",
			}, "\n")),
		},
	}

	for _, tc := range cases {
		err := validate.Exec(tc.in)
		require.Equal(t, tc.err, err)
	}
}
