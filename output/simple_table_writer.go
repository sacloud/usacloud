package output

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io"
)

// simpleTableWriter is write table that output one line for each value
type simpleTableWriter struct {
	table      *tablewriter.Table
	columnDefs []ColumnDef
}

func newSimpleTableWriter(out io.Writer, columnDefs []ColumnDef) tableWriter {
	w := &simpleTableWriter{
		table:      tablewriter.NewWriter(out),
		columnDefs: columnDefs,
	}

	var headers []string
	for _, def := range columnDefs {
		headers = append(headers, def.Name)
	}

	w.table.SetHeader(headers)
	w.table.SetAlignment(tablewriter.ALIGN_LEFT)
	w.table.SetAutoWrapText(false)
	return w
}

func (w *simpleTableWriter) append(values map[string]string) {

	rowValeus := []string{}
	for _, def := range w.columnDefs {
		collected := ""

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
