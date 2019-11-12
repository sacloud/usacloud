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
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/astaxie/flatmap"
	"github.com/bitly/go-simplejson"
)

type tableOutput struct {
	Out           io.Writer
	Err           io.Writer
	IncludeFields []string
	ExcludeFields []string
	ColumnDefs    []ColumnDef
	TableType     TableType
}

func NewTableOutput(out io.Writer, err io.Writer, formater Formatter) Output {
	return &tableOutput{
		Out:           out,
		Err:           err,
		IncludeFields: formater.GetIncludeFields(),
		ExcludeFields: formater.GetExcludeFields(),
		ColumnDefs:    formater.GetColumnDefs(),
		TableType:     formater.GetTableType(),
	}
}

func (o *tableOutput) Print(targets ...interface{}) error {
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

	table, err := o.getTableWriter()
	if err != nil {
		return fmt.Errorf("TableOutput:Print: %s", err)
	}

	// targets -> byte[] -> []interface{}
	rawArray, err := json.Marshal(targets)
	if err != nil {
		return fmt.Errorf("TableOutput:Print: json.Marshal is failed: %s", err)
	}

	j, err := simplejson.NewJson(rawArray)
	if err != nil {
		return fmt.Errorf("TableOutput:Print: create simplejson is failed: %s", err)
	}

	for i := range targets {

		// interface{} -> map[string]interface{}
		v := j.GetIndex(i)
		mapValue, err := v.Map()
		if err != nil {
			return fmt.Errorf("TableOutput:Print: json format is invalid: %v", err)
		}

		// to flatmap( map[string]string )
		flatMap, err := flatmap.Flatten(mapValue)
		if err != nil {
			return fmt.Errorf("TableOutput:Print: create flatmap is failed: %v", err)
		}

		flatMap["__ORDER__"] = fmt.Sprintf("%d", i+1)
		table.append(flatMap)
	}

	table.render()
	return nil
}

func (o *tableOutput) getTableWriter() (tableWriter, error) {
	switch o.TableType {
	case TableSimple:
		return newSimpleTableWriter(o.Out, o.ColumnDefs), nil
	case TableDetail:
		return newDetailTableWriter(o.Out, o.IncludeFields, o.ExcludeFields), nil
	default:
		return nil, fmt.Errorf("unknown OutputTableType(%v)", o.TableType)
	}
}
