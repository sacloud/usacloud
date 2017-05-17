package output

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io"
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
