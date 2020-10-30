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

package ftag

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
	Nested1 *nested1
}

type nested1 struct {
	Field1 string
	Field2 *nested2
}

type nested2 struct {
	Field1 string
}

type squash struct {
	Nested1 *nested1 `cli:",squash"`
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
						Name: "example-field1",
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
						Name: "example-field1",
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
						Name: "no-tag",
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
						Name: "nested-1-field-1",
					},
				},
				{
					StructField: reflect.TypeOf(nested2{}).Field(0),
					Tag: Tag{
						Name: "nested-1-field-2-field-1",
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
						Name: "field-1",
					},
				},
				{
					StructField: reflect.TypeOf(nested2{}).Field(0),
					Tag: Tag{
						Name: "field-2-field-1",
					},
				},
				{
					StructField: reflect.TypeOf(nested1{}).Field(0),
					Tag: Tag{
						Name: "nested-2-field-1",
					},
				},
				{
					StructField: reflect.TypeOf(nested2{}).Field(0),
					Tag: Tag{
						Name: "nested-2-field-2-field-1",
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
