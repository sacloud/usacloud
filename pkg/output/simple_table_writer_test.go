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

package output

import (
	"io"
	"io/ioutil"
	"testing"

	"github.com/olekukonko/tablewriter"
	"github.com/stretchr/testify/assert"
)

type testTable struct {
	*tablewriter.Table
	values [][]string
}

func (t *testTable) Append(row []string) {
	t.values = append(t.values, row)
	t.Table.Append(row)
}

type testSimpleTableWriter struct {
	*simpleTableWriter
}

func (w *testSimpleTableWriter) getValues() [][]string {
	t := w.table.(*testTable)
	return t.values
}

func newTestSimpleTableWriter(out io.Writer, columnDefs []ColumnDef) *testSimpleTableWriter {
	w := &testSimpleTableWriter{
		simpleTableWriter: newSimpleTableWriter(out, columnDefs),
	}

	t := &testTable{
		Table: tablewriter.NewWriter(out),
	}
	w.table = t
	return w
}

func simpleTableTestValue() map[string]string {
	return map[string]string{
		"ID":                         "999999999999",
		"Name":                       "TestValue",
		"Dummy":                      "1",
		"Interfaces.0.IPAddress":     "192.2.0.1",
		"Interfaces.0.UserIPAddress": "192.2.0.2",
		"Interfaces.0.Switch.UserSubnet.NetworkMaskLen": "24",
		"Interfaces.0.Switch.Scope":                     "shared",
	}
}

func TestSimpleTableWriter_Basic(t *testing.T) {
	defs := []ColumnDef{
		{
			Name: "ID",
		},
	}

	writer := newTestSimpleTableWriter(ioutil.Discard, defs)
	value := simpleTableTestValue()
	if err := writer.append(value); err != nil {
		t.Fatal(err)
	}

	values := writer.getValues()

	assert.Len(t, values, 1)
	assert.Equal(t, values[0][0], value["ID"])
}

func TestSimpleTableWriter_EmptyColumn(t *testing.T) {
	defs := []ColumnDef{
		{
			Name: "EmptyCol",
		},
	}

	writer := newTestSimpleTableWriter(ioutil.Discard, defs)
	value := simpleTableTestValue()
	if err := writer.append(value); err != nil {
		t.Fatal(err)
	}

	values := writer.getValues()

	assert.Len(t, values, 1)
	assert.Equal(t, values[0][0], "-")
}
