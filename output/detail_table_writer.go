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
