package output

import (
	"io"
	"sort"
	"strings"

	"github.com/olekukonko/tablewriter"
)

// detailTableWriter is write table that have two columns( attr_name , value )
type detailTableWriter struct {
	table    *tablewriter.Table
	includes []string
	excludes []string
}

func newDetailTableWriter(out io.Writer, includes []string, excludes []string) tableWriter {
	w := &detailTableWriter{
		table:    tablewriter.NewWriter(out),
		includes: includes,
		excludes: excludes,
	}

	w.table.SetHeader([]string{"Name", "Value"})
	w.table.SetAlignment(tablewriter.ALIGN_LEFT)
	w.table.SetAutoWrapText(false)
	w.table.SetAutoFormatHeaders(false)
	return w
}

func (w *detailTableWriter) append(values map[string]string) {

	var keys []string
	for k := range values {
		keys = append(keys, k)
	}
	sortedKeys := sortedKeys(keys)
	sort.Sort(sortedKeys)

	includes := w.includes
	if len(includes) == 0 {
		includes = []string{""}
	}

	// include
	for _, k := range sortedKeys {
		if w.needInclude(k, includes) && !w.needExclude(k, w.excludes) && k != "__ORDER__" {

			sv := values[k]
			if sv == "" {
				sv = "-"
			}

			w.table.Append([]string{k, sv})
		}
	}

}

func (w *detailTableWriter) needInclude(targetName string, includes []string) bool {
	for _, include := range includes {
		if include == "" {
			return true
		}

		res := strings.HasPrefix(targetName, include)
		if res {
			return res
		}
	}
	return false
}

func (w *detailTableWriter) needExclude(targetName string, excludes []string) bool {
	for _, exclude := range excludes {
		if exclude != "" {
			res := strings.HasPrefix(targetName, exclude)
			if res {
				return res
			}
		}
	}
	return false
}

func (w *detailTableWriter) render() {
	w.table.Render()
}

type sortedKeys []string

func (l sortedKeys) Len() int {
	return len(l)
}

func (l sortedKeys) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l sortedKeys) Less(i, j int) bool {

	c1 := strings.Count(l[i], ".")
	c2 := strings.Count(l[j], ".")
	if c1 != c2 {
		return c1 < c2
	}
	return l[i] < l[j]
}
