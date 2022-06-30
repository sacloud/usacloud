// Copyright 2017-2022 The sacloud/usacloud Authors
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
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

type simple struct {
	Field1 string `cli:"example-field1"`
}

type unexported struct {
	unexported string // nolint
}

type noTag struct {
	NoTag string
}

type ignore struct {
	Ignore string `cli:"-"`
}

type recursive struct {
	Nested1 *nested1 `cli:",category=parent"`
}

type nested1 struct {
	Field1 string
	Field2 *nested2
}

type nested2 struct {
	Field1 string
}

type squash struct {
	Nested1 *nested1 `cli:",squash,category=parent"`
	Nested2 *nestedSquash
}

type nestedSquash struct {
	Nested1 *nested1 `cli:",squash"`
}

func TestParser_Parse(t *testing.T) {
	cases := []struct {
		in        interface{}
		expect    []StructField
		errString string
	}{
		{
			in:        nil,
			errString: "value required",
		},
		{
			in:        "",
			errString: `unsupported value: ""`,
		},
		{
			in:        1,
			errString: `unsupported value: 1`,
		},
		{
			in:        []string{},
			errString: `unsupported value: []string{}`,
		},
		{
			in: simple{},
			expect: []StructField{
				{
					StructField: reflect.TypeOf(simple{}).Field(0),
					Tag: Tag{
						FieldName: "Field1",
						FlagName:  "example-field1",
					},
				},
			},
		},
		{
			in: &simple{}, // with pointer
			expect: []StructField{
				{
					StructField: reflect.TypeOf(simple{}).Field(0),
					Tag: Tag{
						FieldName: "Field1",
						FlagName:  "example-field1",
					},
				},
			},
		},
		{
			in:     &unexported{},
			expect: nil,
		},
		{
			in: &noTag{},
			expect: []StructField{
				{
					StructField: reflect.TypeOf(noTag{}).Field(0),
					Tag: Tag{
						FieldName: "NoTag",
						FlagName:  "no-tag",
					},
				},
			},
		},
		{
			in:     &ignore{},
			expect: nil,
		},
		{
			in: &recursive{},
			expect: []StructField{
				{
					StructField: reflect.TypeOf(nested1{}).Field(0),
					Tag: Tag{
						FieldName: "Nested1.Field1",
						FlagName:  "nested1-field1",
						Category:  "parent",
					},
				},
				{
					StructField: reflect.TypeOf(nested2{}).Field(0),
					Tag: Tag{
						FieldName: "Nested1.Field2.Field1",
						FlagName:  "nested1-field2-field1",
						Category:  "parent",
					},
				},
			},
		},
		{
			in: &squash{},
			expect: []StructField{
				{
					StructField: reflect.TypeOf(nested1{}).Field(0),
					Tag: Tag{
						FieldName: "Nested1.Field1",
						FlagName:  "field1",
						Category:  "parent",
					},
				},
				{
					StructField: reflect.TypeOf(nested2{}).Field(0),
					Tag: Tag{
						FieldName: "Nested1.Field2.Field1",
						FlagName:  "field2-field1",
						Category:  "parent",
					},
				},
				{
					StructField: reflect.TypeOf(nested1{}).Field(0),
					Tag: Tag{
						FieldName: "Nested2.Nested1.Field1",
						FlagName:  "nested2-field1",
					},
				},
				{
					StructField: reflect.TypeOf(nested2{}).Field(0),
					Tag: Tag{
						FieldName: "Nested2.Nested1.Field2.Field1",
						FlagName:  "nested2-field2-field1",
					},
				},
			},
		},
	}

	for _, tc := range cases {
		fields, err := Parse(tc.in)
		if tc.errString != "" {
			require.EqualError(t, err, tc.errString)
			continue
		}
		require.EqualValues(t, tc.expect, fields)
	}
}
