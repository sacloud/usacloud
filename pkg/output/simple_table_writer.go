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
	"bytes"
	"fmt"
	"io"
	"text/template"

	"github.com/olekukonko/tablewriter"
)

// simpleTableWriter is write table that output one line for each value
type simpleTableWriter struct {
	table      tableHandler
	columnDefs []ColumnDef
}

type tableHandler interface {
	SetHeader(keys []string)
	SetAlignment(align int)
	SetAutoWrapText(auto bool)
	SetAutoFormatHeaders(auto bool)
	Append(row []string)
	Render()
}

func newSimpleTableWriter(out io.Writer, columnDefs []ColumnDef) *simpleTableWriter {
	if len(columnDefs) == 0 {
		columnDefs = []ColumnDef{
			{Name: "{{.__ORDER__}}"},
			{Name: "{{.ID}}"},
			{Name: "{{.Name}}"},
		}
	}

	w := &simpleTableWriter{
		table:      tablewriter.NewWriter(out),
		columnDefs: columnDefs,
	}

	var headers []string
	for i, def := range columnDefs {
		name := def.Name
		if name == "__ORDER__" {
			name = "#"
		}
		headers = append(headers, name)

		if def.Template == "" {
			columnDefs[i].Template = fmt.Sprintf("{{.%s}}", def.Name)
		}
	}

	w.table.SetHeader(headers)
	w.table.SetAlignment(tablewriter.ALIGN_LEFT)
	w.table.SetAutoWrapText(false)
	w.table.SetAutoFormatHeaders(false)
	return w
}

func (w *simpleTableWriter) append(values interface{}) error {
	var rowValeus []string
	for _, def := range w.columnDefs {
		t := template.Must(template.New("output").Option("missingkey=zero").Parse(def.Template))

		buf := bytes.NewBufferString("")
		err := t.Execute(buf, values)
		if err != nil {
			return err
		}
		s := buf.String()
		if s == "" || s == "<no value>" { // HACK: map[string]interface{}の場合、missingkey=zeroオプションありでもキーがない場合は<no value>となる
			s = "-"
		}
		rowValeus = append(rowValeus, s)
	}

	w.table.Append(rowValeus)
	return nil
}

func (w *simpleTableWriter) render() {
	w.table.Render()
}
