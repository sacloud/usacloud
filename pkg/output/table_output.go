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
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/sacloud/usacloud/pkg/util"

	"github.com/bitly/go-simplejson"
)

type tableOutput struct {
	Out        io.Writer
	Err        io.Writer
	ColumnDefs []ColumnDef
}

func NewTableOutput(out io.Writer, err io.Writer, columnDefs []ColumnDef) Output {
	return &tableOutput{
		Out:        out,
		Err:        err,
		ColumnDefs: columnDefs,
	}
}

func (o *tableOutput) Print(contents Contents) error {
	targets := contents.Values()
	if o.Out == nil {
		o.Out = os.Stdout
	}
	if o.Err == nil {
		o.Err = os.Stderr
	}

	if util.IsEmpty(targets) {
		return nil
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

	table := newSimpleTableWriter(o.Out, o.ColumnDefs)
	for i := 0; i < len(targets); i++ {
		// interface{} -> map[string]interface{}
		v := j.GetIndex(i)
		mapValue, err := v.Map()
		if err != nil {
			return fmt.Errorf("TableOutput:Print: json format is invalid: %v", err)
		}

		mapValue["__ORDER__"] = fmt.Sprintf("%d", i+1)
		if contents[i].Zone != "" {
			if _, ok := mapValue["Zone"]; !ok {
				mapValue["Zone"] = contents[i].Zone
			}
		}
		if !contents[i].ID.IsEmpty() {
			if _, ok := mapValue["ID"]; !ok {
				mapValue["ID"] = contents[i].ID
			}
		}
		if err := table.append(mapValue); err != nil {
			return fmt.Errorf("TableOutput:Print: processing template failed: %v", err)
		}
	}

	table.render()
	return nil
}
