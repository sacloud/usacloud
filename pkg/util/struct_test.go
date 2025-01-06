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

package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func setup() (*os.File, func(), error) {
	tempFile, err := os.CreateTemp("", "*_test.json")
	if err != nil {
		return nil, nil, err
	}
	return tempFile, func() { os.Remove(tempFile.Name()) }, nil
}

func TestBytesFromPathOrContent(t *testing.T) {
	in := `{"Foo":"Bar"}`

	// setup
	tempFile, cleaner, err := setup()
	if err != nil {
		t.Fatal(err)
	}
	defer cleaner()

	if _, err := tempFile.WriteString(in); err != nil {
		t.Fatal(err)
	}

	// from string
	fromString, err := BytesFromPathOrContent(in)
	if err != nil {
		t.Fatal(err)
	}
	if string(fromString) != in {
		t.Fatalf("got unexpected value: expected:%s got:%s", in, string(fromString))
	}

	// from File
	fromFile, err := BytesFromPathOrContent(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	if string(fromFile) != in {
		t.Fatalf("got unexpected value: expected:%s got:%s", in, string(fromFile))
	}
}

func TestMarshalJSONFromPathOrContent(t *testing.T) {
	in := `{"Foo":"Bar"}`

	// setup
	tempFile, cleaner, err := setup()
	if err != nil {
		t.Fatal(err)
	}
	defer cleaner()

	if _, err := tempFile.WriteString(in); err != nil {
		t.Fatal(err)
	}

	type target struct {
		Foo string
	}

	cases := []struct {
		in        string
		expect    *target
		errString string
	}{
		{
			in:     tempFile.Name(),
			expect: &target{Foo: "Bar"},
		},
		{
			in:     in,
			expect: &target{Foo: "Bar"},
		},
		{
			in:        "",
			errString: "pathOrContent required",
		},
		{
			in:        `{"Foo": 1}`,
			errString: "json: cannot unmarshal number into Go struct field target.Foo of type string",
		},
		{
			in:     `{"Bar": 1}`, // 存在しないキーはエラーとしない(json.Unmarshalと同じ)
			expect: &target{},
		},
	}

	for _, tc := range cases {
		dest := &target{}
		err := MarshalJSONFromPathOrContent(tc.in, dest)
		require.Equal(t, tc.errString == "", err == nil)
		if err != nil {
			require.EqualError(t, err, tc.errString)
			continue
		}
		require.EqualValues(t, tc.expect, dest)
	}
}
