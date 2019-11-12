// Copyright 2017-2019 The Usacloud Authors
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
	"fmt"
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
		simpleTableWriter: newSimpleTableWriter(out, columnDefs).(*simpleTableWriter),
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
	writer.append(value)

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
	writer.append(value)

	values := writer.getValues()

	assert.Len(t, values, 1)
	assert.Equal(t, values[0][0], "-")
}

func TestSimpleTableWriter_Format(t *testing.T) {
	format := "%s(%s)"
	defs := []ColumnDef{
		{
			Name:    "Formated",
			Format:  format,
			Sources: []string{"Name", "ID"},
		},
	}
	writer := newTestSimpleTableWriter(ioutil.Discard, defs)
	value := simpleTableTestValue()
	writer.append(value)

	values := writer.getValues()
	assert.Equal(t, values[0][0], fmt.Sprintf(format, value["Name"], value["ID"]))

	// reverse sources order
	defs = []ColumnDef{
		{
			Name:    "Formated",
			Format:  format,
			Sources: []string{"ID", "Name"},
		},
	}
	writer = newTestSimpleTableWriter(ioutil.Discard, defs)
	writer.append(value)

	values = writer.getValues()
	assert.NotEqual(t, values[0][0], fmt.Sprintf(format, value["Name"], value["ID"]))

}

func TestSimpleTableWriter_ValueMapping(t *testing.T) {
	defs := []ColumnDef{
		{
			Name: "Dummy",
			ValueMapping: []map[string]string{
				{
					"1": "test",
				},
			},
		},
	}

	writer := newTestSimpleTableWriter(ioutil.Discard, defs)
	value := simpleTableTestValue()
	writer.append(value)

	values := writer.getValues()
	assert.Equal(t, values[0][0], "test")

	// mapping not exists
	value["Dummy"] = "2"

	writer = newTestSimpleTableWriter(ioutil.Discard, defs)
	writer.append(value)

	values = writer.getValues()
	assert.Equal(t, values[0][0], "2")
}

func TestSimpleTableWriter_ValueMappingMulti(t *testing.T) {
	format := "%s:%s"
	defs := []ColumnDef{
		{
			Name:    "Dummy",
			Sources: []string{"ID", "Dummy"},
			ValueMapping: []map[string]string{
				{ // for ID
				},
				{ // for Dummy
					"1": "test",
				},
			},
			Format: format,
		},
	}

	writer := newTestSimpleTableWriter(ioutil.Discard, defs)
	value := simpleTableTestValue()
	writer.append(value)

	values := writer.getValues()
	assert.Equal(t, values[0][0], fmt.Sprintf(format, value["ID"], "test"))

	// mapping not exists
	value["Dummy"] = "2"

	writer = newTestSimpleTableWriter(ioutil.Discard, defs)
	writer.append(value)

	values = writer.getValues()
	assert.Equal(t, values[0][0], fmt.Sprintf(format, value["ID"], "2"))
}

func TestSimpleTableWriter_CustomFormat(t *testing.T) {

	formatFunc := func(values map[string]string) string {
		if scope, ok := values["Interfaces.0.Switch.Scope"]; ok {
			format := "%s/%s"
			switch scope {
			case "shared":
				return fmt.Sprintf(format,
					values["Interfaces.0.IPAddress"],
					values["Interfaces.0.Switch.UserSubnet.NetworkMaskLen"],
				)
			case "user":
				return fmt.Sprintf(format,
					values["Interfaces.0.UserIPAddress"],
					values["Interfaces.0.Switch.UserSubnet.NetworkMaskLen"],
				)

			}

		}

		return ""
	}

	defs := []ColumnDef{
		{
			Name:       "Dummy",
			FormatFunc: formatFunc,
		},
	}

	writer := newTestSimpleTableWriter(ioutil.Discard, defs)
	value := simpleTableTestValue()
	writer.append(value)

	values := writer.getValues()
	assert.Equal(t, values[0][0],
		fmt.Sprintf("%s/%s",
			value["Interfaces.0.IPAddress"],
			value["Interfaces.0.Switch.UserSubnet.NetworkMaskLen"],
		),
	)

	// mapping not exists
	value["Interfaces.0.Switch.Scope"] = "user"

	writer = newTestSimpleTableWriter(ioutil.Discard, defs)
	writer.append(value)

	values = writer.getValues()
	assert.Equal(t, values[0][0],
		fmt.Sprintf("%s/%s",
			value["Interfaces.0.UserIPAddress"],
			value["Interfaces.0.Switch.UserSubnet.NetworkMaskLen"],
		),
	)
}
