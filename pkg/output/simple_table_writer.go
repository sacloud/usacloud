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
	"fmt"
	"io"

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

func newSimpleTableWriter(out io.Writer, columnDefs []ColumnDef) tableWriter {
	w := &simpleTableWriter{
		table:      tablewriter.NewWriter(out),
		columnDefs: columnDefs,
	}

	var headers []string
	for _, def := range columnDefs {
		name := def.Name
		if name == "__ORDER__" {
			name = "#"
		}
		headers = append(headers, name)
	}

	w.table.SetHeader(headers)
	w.table.SetAlignment(tablewriter.ALIGN_LEFT)
	w.table.SetAutoWrapText(false)
	w.table.SetAutoFormatHeaders(false)
	return w
}

func (w *simpleTableWriter) append(values map[string]string) {
	rowValeus := []string{}
	for _, def := range w.columnDefs {
		collected := ""

		if def.FormatFunc == nil {
			exists := false
			var sources []interface{}
			// collect source values
			for i, source := range def.GetSources() {
				var s string
				if v, ok := values[source]; ok {
					s = v

					if i < len(def.ValueMapping) {
						mapping := def.ValueMapping[i]
						if mapped, ok := mapping[s]; ok {
							s = mapped
						}
					}
				}
				if s != "" {
					exists = true
				}
				sources = append(sources, s)
			}
			//format
			if exists {
				format := def.GetFormat()
				collected = fmt.Sprintf(format, sources...)
			}
		} else {
			collected = def.FormatFunc(values)
		}

		if collected == "" {
			collected = "-"
		}
		rowValeus = append(rowValeus, collected)
	}

	w.table.Append(rowValeus)
}

func (w *simpleTableWriter) render() {
	w.table.Render()
}
