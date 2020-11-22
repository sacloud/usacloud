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
	"sync"
	"text/template"

	"github.com/olekukonko/tablewriter"
)

// SimpleTableWriter is write table that output one line for each value
type SimpleTableWriter struct {
	table      *tablewriter.Table
	columnDefs []ColumnDef
	template   *template.Template
	initOnce   sync.Once
}

func NewSimpleTableWriter(out io.Writer, columnDefs []ColumnDef) *SimpleTableWriter {
	if len(columnDefs) == 0 {
		columnDefs = []ColumnDef{
			{Name: "{{.__ORDER__}}"},
			{Name: "{{.ID}}"},
			{Name: "{{.Name}}"},
		}
	}

	w := &SimpleTableWriter{
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

// CellValue 指定のColumnDefを用いてvから出力を組み立てる
//
// NOTE: pkg/cmd/commands/*でColumnDefのテストをするためにエクスポートしている
func (w *SimpleTableWriter) CellValue(v interface{}, def ColumnDef) (string, error) {
	w.initOnce.Do(func() {
		w.template = newTemplate()
	})

	if _, err := w.template.Parse(def.Template); err != nil {
		return "", fmt.Errorf("invalid output format %q: %s", def.Template, err)
	}

	buf := bytes.NewBufferString("")
	if err := w.template.Execute(buf, v); err != nil {
		return "", err
	}
	s := buf.String()
	if s == "" || s == "<no value>" { // HACK: map[string]interface{}の場合、missingkey=zeroオプションありでもキーがない場合は<no value>となる
		s = "-"
	}
	return s, nil
}

func (w *SimpleTableWriter) append(values interface{}) error {
	var rowValeus []string
	for _, def := range w.columnDefs {
		s, err := w.CellValue(values, def)
		if err != nil {
			return err
		}
		rowValeus = append(rowValeus, s)
	}

	w.table.Append(rowValeus)
	return nil
}

func (w *SimpleTableWriter) render() {
	w.table.Render()
}
