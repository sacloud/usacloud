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
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/astaxie/flatmap"
	"github.com/bitly/go-simplejson"
	"github.com/sacloud/usacloud/pkg/util"
)

type rowOutput struct {
	Out       io.Writer
	Err       io.Writer
	Separator rune
	Columns   []string
}

func NewRowOutput(out io.Writer, err io.Writer, separator rune, option Option) Output {
	return &rowOutput{
		Out:       out,
		Err:       err,
		Separator: separator,
		Columns:   option.GetColumn(),
	}
}

func (o *rowOutput) Print(target interface{}) error {
	targets := toSlice(target)
	if o.Out == nil {
		o.Out = os.Stdout
	}
	if o.Err == nil {
		o.Err = os.Stderr
	}

	if util.IsEmpty(targets) {
		fmt.Fprintln(o.Err, "no results")
		return nil
	}

	w := csv.NewWriter(o.Out)
	w.UseCRLF = true
	w.Comma = o.Separator

	header := header(o.Columns)
	headerExists := map[string]bool{}

	// targets -> byte[] -> []interface{}
	rawArray, err := json.Marshal(targets)
	if err != nil {
		return fmt.Errorf("RowOutput:Print: json.Marshal is failed: %s", err)
	}

	j, err := simplejson.NewJson(rawArray)
	if err != nil {
		return fmt.Errorf("RowOutput:Print: create simplejson is failed: %s", err)
	}

	// first, collect header
	if len(header) == 0 {
		header = append(header, "RowNumber")
		for i := 0; i < sliceLen(targets); i++ {
			// interface{} -> map[string]interface{}
			v := j.GetIndex(i)
			mapValue, err := v.Map()
			if err != nil {
				return fmt.Errorf("RowOutput:Print: json format is invalid: %v", err)
			}

			// to flatmap( map[string]string )
			flatMap, err := flatmap.Flatten(mapValue)
			if err != nil {
				return fmt.Errorf("TableOutput:Print: create flatmap is failed: %v", err)
			}

			for k := range flatMap {
				if _, ok := headerExists[k]; !ok {
					header = append(header, k)
					headerExists[k] = true
				}
			}
		}
	}

	sort.Sort(header)
	// write header
	w.Write(header) // nolint

	// next, collect values
	for rowIndex := 0; rowIndex < sliceLen(targets); rowIndex++ {
		// interface{} -> map[string]interface{}
		v := j.GetIndex(rowIndex)
		mapValue, err := v.Map()
		if err != nil {
			return fmt.Errorf("RowOutput:Print: json format is invalid: %v", err)
		}

		// to flatmap( map[string]string )
		flatMap, err := flatmap.Flatten(mapValue)
		if err != nil {
			return fmt.Errorf("TableOutput:Print: create flatmap is failed: %v", err)
		}

		row := []string{}
		for i := range header {
			k := header[i]
			value := ""
			if k == "RowNumber" {
				value = fmt.Sprintf("%d", rowIndex+1)
			} else {
				if v, ok := flatMap[k]; ok {
					value = v
				}
			}
			row = append(row, value)
		}

		w.Write(row) // nolint
	}

	w.Flush()
	return nil
}

type header []string

func (s header) Len() int {
	return len(s)
}

func (s header) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s header) Less(i, j int) bool {
	if s[j] == "RowNumber" {
		return false
	}
	if s[i] == "RowNumber" {
		return true
	}
	if s[j] == "ID" {
		return false
	}
	if s[i] == "ID" {
		return true
	}
	return s[i] < s[j]
}
