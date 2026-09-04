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

package common

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecodeJSON(t *testing.T) {
	type input struct {
		Name string `json:"Name"`
	}

	t.Run("valid", func(t *testing.T) {
		var value input
		require.NoError(t, DecodeJSON(strings.NewReader(`{"Name":"example"}`), &value))
		require.Equal(t, "example", value.Name)
	})

	t.Run("unknown field", func(t *testing.T) {
		var value input
		err := DecodeJSON(strings.NewReader(`{"Unknown":"value"}`), &value)
		require.ErrorContains(t, err, `unknown field "Unknown"`)
	})

	t.Run("multiple values", func(t *testing.T) {
		var value input
		err := DecodeJSON(strings.NewReader(`{"Name":"first"} {"Name":"second"}`), &value)
		require.EqualError(t, err, "decode JSON: multiple values are not allowed")
	})
}
