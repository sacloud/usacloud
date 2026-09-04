// Copyright 2017-2025 The sacloud/usacloud Authors
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

package output

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

type unionID struct {
	Type   string
	String string
	Int    int
}

type unionIDResource struct {
	ID unionID
}

func TestIDOutputPrintsUnionIDValue(t *testing.T) {
	tests := []struct {
		name     string
		resource unionIDResource
		expected string
	}{
		{
			name:     "integer ID",
			resource: unionIDResource{ID: unionID{Type: "int", Int: 113802049593}},
			expected: "113802049593\n",
		},
		{
			name:     "string ID",
			resource: unionIDResource{ID: unionID{Type: "string", String: "queue-id"}},
			expected: "queue-id\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out bytes.Buffer
			contents := Contents{{Value: tt.resource}}

			require.NoError(t, NewIDOutput(&out, nil).Print(contents))
			require.Equal(t, tt.expected, out.String())
		})
	}
}
