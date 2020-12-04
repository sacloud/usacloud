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

package naming

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHoge(t *testing.T) {
	t.Log(ToKebabCase("DNS1"))
}

func TestNames(t *testing.T) {
	cases := []struct {
		f   func(string) string
		in  string
		out string
	}{
		{
			f:   ToCamelCase,
			in:  "AuthStatus",
			out: "AuthStatus",
		},
		{
			f:   ToSnakeCase,
			in:  "IPv4",
			out: "ipv4",
		},
		{
			f:   ToCamelCase,
			in:  "IPv4",
			out: "IPv4",
		},
		{
			f:   ToCamelCaseWithFirstLower,
			in:  "IPv4",
			out: "ipv4",
		},
		{
			f:   ToSnakeCase,
			in:  "DNS",
			out: "dns",
		},
		{
			f:   ToCamelCase,
			in:  "DNS",
			out: "DNS",
		},
		{
			f:   ToCamelCaseWithFirstLower,
			in:  "DNS",
			out: "dns",
		},
		{
			f:   ToCamelCase,
			in:  "ipv6-enable",
			out: "IPv6Enable",
		},
		{
			f:   ToCamelCase,
			in:  "sim-info",
			out: "SIMInfo",
		},
		{
			f:   ToCamelCase,
			in:  "simple-monitor",
			out: "SimpleMonitor",
		},
	}

	for _, tc := range cases {
		got := tc.f(tc.in)
		require.Equal(t, tc.out, got)
	}
}
