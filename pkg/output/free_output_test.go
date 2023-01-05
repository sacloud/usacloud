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

package output

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type dummyOption struct{}

func (o dummyOption) OutputTypeFlagValue() string  { return "" }
func (o dummyOption) GetColumn() []string          { return []string{} }
func (o dummyOption) FormatFlagValue() string      { return "test ID:{{.ID}}" }
func (o dummyOption) FormatFileFlagValue() string  { return "" }
func (o dummyOption) QuietFlagValue() bool         { return false }
func (o dummyOption) QueryFlagValue() string       { return "" }
func (o dummyOption) QueryDriverFlagValue() string { return "" }

func TestFreeOutput_Print(t *testing.T) {
	buf := bytes.NewBufferString("")
	o := NewFreeOutput(buf, os.Stderr, dummyOption{})

	type dummy struct {
		ID int64
	}

	values := Contents{
		{Value: &dummy{ID: 1}},
		{Value: &dummy{ID: 2}},
	}

	err := o.Print(values)

	assert.NoError(t, err)
	assert.Equal(t, testFreeOutputText, buf.String())
}

var testFreeOutputText = `test ID:1
test ID:2
`
