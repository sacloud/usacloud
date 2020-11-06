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

	"github.com/sacloud/libsacloud/v2/helper/service/disk"
	"github.com/sacloud/libsacloud/v2/pkg/mapconv"
	"github.com/sacloud/usacloud/pkg/cmd/base"
	"github.com/stretchr/testify/require"
)

func TestList_ConvertToServiceRequest(t *testing.T) {
	in := &ListParameter{
		OutputParameter: nil,
		Zone:            "is1a",
		Names:           []string{"name1", "name2"},
		Tags:            []string{"tag1", "tag2"},
		FindParameter: &base.FindParameter{
			Count: 1,
			From:  2,
		},
	}

	out := &disk.FindRequest{}
	if err := mapconv.ConvertTo(in, out); err != nil {
		t.Fatal(err)
	}

	require.EqualValues(t, &disk.FindRequest{
		Zone:  "is1a",
		Names: []string{"name1", "name2"},
		Tags:  []string{"tag1", "tag2"},
		Count: 1,
		From:  2,
	}, out)
}
