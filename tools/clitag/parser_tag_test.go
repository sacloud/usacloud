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

package clitag

import (
	"testing"

	require "github.com/stretchr/testify/require"
)

func TestParser_parseTag(t *testing.T) {
	cases := []struct {
		in        string
		expect    Tag
		errString string
	}{
		{
			in:     ``,
			expect: Tag{},
		},
		{
			in: `example`,
			expect: Tag{
				FlagName: "example",
			},
		},
		{
			in:     `-`,
			expect: Tag{Ignore: true},
		},
		{
			in: `,aliases=foo bar`,
			expect: Tag{
				Aliases: []string{"foo", "bar"},
			},
		},
		{
			in: `,aliases= foo bar`,
			expect: Tag{
				Aliases: []string{"foo", "bar"},
			},
		},
		{
			in: `,short=e,desc=desc,squash,category=foo,order=10`,
			expect: Tag{
				Shorthand:   "e",
				Description: "desc",
				Squash:      true,
				Category:    "foo",
				Order:       10,
			},
		},
		{
			in:        `,short==,desc=desc`,
			errString: `got invalid tag value: "short=="`,
		},
		{
			in:        `,short=ee`,
			errString: `got invalid tag value: key 'short' must have only 1 character: "short=ee"`,
		},
		{
			in:        `,foo=bar`,
			errString: `got invalid tag key: "foo=bar"`,
		},
	}

	for _, tc := range cases {
		parser := &Parser{}
		tag, err := parser.parseTag(tc.in)
		if tc.errString != "" {
			require.EqualError(t, err, tc.errString, tc.in)
			continue
		}
		require.EqualValues(t, tc.expect, tag, tc.in)
	}
}
