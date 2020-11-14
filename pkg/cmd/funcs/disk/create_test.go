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
	"errors"
	"strings"
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/usacloud/pkg/validate"

	"github.com/sacloud/usacloud/pkg/cmd/core"

	"github.com/sacloud/libsacloud/v2/sacloud/ostype"

	"github.com/sacloud/usacloud/pkg/cmd/conv"

	"github.com/sacloud/libsacloud/v2/sacloud/types"

	"github.com/sacloud/libsacloud/v2/helper/service/disk"
	"github.com/stretchr/testify/require"
)

func TestCreate_ConvertToServiceRequest(t *testing.T) {
	t.Run("full", func(t *testing.T) {
		in := &createParameter{
			ZoneParameter: core.ZoneParameter{Zone: "is1a"},
			Name:          "name",
			Description:   "desc",
			Tags:          []string{"tag1", "tag2"},
			IconID:        types.ID(1),
			DiskPlan:      "ssd",
			Connection:    types.DiskConnections.VirtIO.String(),
			SizeGB:        20,
			OSType:        "ubuntu",
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

func TestCreateParameter_Validate(t *testing.T) {
	cases := []struct {
		in  *createParameter
		err error
	}{
		// default
		{
			in: newCreateParameter(),
			err: errors.New(strings.Join([]string{
				"validation error:",
				"\t--zone: required",
				"\t--name: required",
			}, "\n")),
		},
		// minimum
		{
			in: &createParameter{
				ZoneParameter: core.ZoneParameter{
					Zone: "is1a",
				},
				Name:       "foobar",
				DiskPlan:   "ssd",
				Connection: "virtio",
			},
			err: nil,
		},
		// invalid tags length
		{
			in: &createParameter{
				ZoneParameter: core.ZoneParameter{
					Zone: "is1a",
				},
				Name:       "foobar",
				DiskPlan:   "ssd",
				Connection: "virtio",
				Tags:       []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "invalid"},
			},
			err: errors.New(strings.Join([]string{
				"validation error:",
				"\t--tags: max=10",
			}, "\n")),
		},
		// invalid tags body
		{
			in: &createParameter{
				ZoneParameter: core.ZoneParameter{
					Zone: "is1a",
				},
				Name:       "foobar",
				DiskPlan:   "ssd",
				Connection: "virtio",
				Tags:       []string{"********10********20********30++x"},
			},
			err: errors.New(strings.Join([]string{
				"validation error:",
				"\t--tags[0]: max=32",
			}, "\n")),
		},
		// custom validation
		{
			in: &createParameter{
				ZoneParameter: core.ZoneParameter{
					Zone: "is1a",
				},
				Name:            "foobar",
				DiskPlan:        "ssd",
				Connection:      "virtio",
				SourceArchiveID: types.ID(1),
				SourceDiskID:    types.ID(1),
			},
			err: errors.New(strings.Join([]string{
				"validation error:",
				"\t--source-archive-id & --source-disk-id: only one of them can be specified",
			}, "\n")),
		},
		// custom validation(with os-type)
		{
			in: &createParameter{
				ZoneParameter: core.ZoneParameter{
					Zone: "is1a",
				},
				Name:            "foobar",
				DiskPlan:        "ssd",
				Connection:      "virtio",
				OSType:          "ubuntu",
				SourceArchiveID: types.ID(1),
			},
			err: errors.New(strings.Join([]string{
				"validation error:",
				"\t--os-type & --source-archive-id & --source-disk-id: only one of them can be specified",
			}, "\n")),
		},
	}

	validate.InitializeValidator(sacloud.SakuraCloudZones)
	for _, tc := range cases {
		err := validateCreateParameter(nil, tc.in)
		require.Equal(t, tc.err, err)
	}
}
