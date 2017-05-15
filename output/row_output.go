package output

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/astaxie/flatmap"
	"github.com/bitly/go-simplejson"
	"io"
	"os"
	"sort"
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

func (o *rowOutput) Print(targets ...interface{}) error {
	if o.Out == nil {
		o.Out = os.Stdout
	}
	if o.Err == nil {
		o.Err = os.Stderr
	}

	if len(targets) == 0 {
		fmt.Fprintf(o.Err, "Result is empty\n")
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
		for i := range targets {
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
	w.Write(header)

	// next, collect values
	for rowIndex := range targets {

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

		w.Write(row)
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
