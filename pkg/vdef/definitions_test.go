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

package vdef

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefinition_initialize(t *testing.T) {
	// definitionsのinitで各所にfunc登録されるか
	definitions["test"] = []*definition{}
	registerFunctions()

	testNames := []string{
		"test_to_key", "test_to_value",
	}
	for _, name := range testNames {
		_, ok := ConverterFilters[name]
		require.True(t, ok, "%s is not registered to ConverterFilters", name)

		_, ok = TemplateFuncMap[name]
		require.True(t, ok, "%s is not registered to TemplateFuncMap", name)
	}

	_, ok := validatorAliases["test"]
	require.True(t, ok, "test is not registered to validatorAliases")

	_, ok = FlagOptionsMap["test"]
	require.True(t, ok, "test is not registered to FlagOptionsMap")
}

func TestConverterFuncGenerator(t *testing.T) {
	def := []*definition{
		{key: "string", value: "stringValue"},
		{key: 1, value: "intValue"},
		{key: "intKey", value: 1},
	}

	t.Run("convertFuncToValue", func(t *testing.T) {
		f := convertFuncToValue("test", def)
		require.NotNil(t, f)

		cases := []struct {
			in     interface{}
			expect interface{}
			err    string
		}{
			{
				in:     "string",
				expect: "stringValue",
			},
			{
				in:     1,
				expect: "intValue",
			},
			{
				in:  int64(1),
				err: "key 1 not found in test",
			},
			{
				in:     []interface{}{"string", 1},
				expect: []interface{}{"stringValue", "intValue"},
			},
			{
				in:     "foobar",
				expect: nil,
				err:    `key foobar not found in test`,
			},
		}

		for _, tc := range cases {
			res, err := f(tc.in)
			if err != nil {
				require.EqualError(t, err, tc.err)
				continue
			}
			require.EqualValues(t, tc.expect, res)
		}
	})
	t.Run("convertFuncToKey", func(t *testing.T) {
		f := convertFuncToKey("test", def)
		require.NotNil(t, f)

		cases := []struct {
			in     interface{}
			expect interface{}
			err    string
		}{
			{
				in:     "stringValue",
				expect: "string",
			},
			{
				in:     "intValue",
				expect: 1,
			},
			{
				in:     1,
				expect: "intKey",
			},
			{
				in:     int64(1),
				expect: "intKey",
				err:    `value 1 not found in test`,
			},
			{
				in:     []interface{}{"stringValue", "intValue"},
				expect: []interface{}{"string", 1},
			},
			{
				in:  "foobar",
				err: `value foobar not found in test`,
			},
		}

		for _, tc := range cases {
			res, err := f(tc.in)
			if err != nil {
				require.EqualError(t, err, tc.err)
				continue
			}
			require.EqualValues(t, tc.expect, res)
		}
	})
}
