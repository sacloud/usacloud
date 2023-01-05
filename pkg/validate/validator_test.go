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

package validate

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidator_Exec(t *testing.T) {
	type Foo struct {
		Required string `validate:"required"`
		Min2     string `validate:"min=2"`
		Max5     string `validate:"max=5"`
	}

	v := &Foo{
		Required: "",
		Min2:     "1",
		Max5:     "123456",
	}

	err := Exec(v)
	require.EqualError(t, err, strings.Join([]string{
		"validation error:",
		"\t--required: required",
		"\t--min2: min=2",
		"\t--max5: max=5",
	}, "\n"))
}
