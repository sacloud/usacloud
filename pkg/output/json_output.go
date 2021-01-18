// Copyright 2017-2021 The Usacloud Authors
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

	"github.com/sacloud/usacloud/pkg/query"

	"github.com/fatih/structs"
	"github.com/sacloud/usacloud/pkg/util"
)

type jsonOutput struct {
	out    io.Writer
	err    io.Writer
	query  string
	driver string
}

func NewJSONOutput(out io.Writer, err io.Writer, query string, driver string) Output {
	return &jsonOutput{
		out:    out,
		err:    err,
		query:  query,
		driver: driver,
	}
}

func (o *jsonOutput) Print(contents Contents) error {
	targets := contents.Values()
	if o.out == nil {
		o.out = os.Stdout
	}
	if o.err == nil {
		o.err = os.Stderr
	}

	if util.IsEmpty(targets) {
		return nil
	}

	// queryが指定されている場合はメタデータなしでJSON出力処理を行う
	if o.query != "" {
		return o.printWithQuery(targets)
	}

	// queryが指定されていない場合はメタデータありのJSON出力処理を行う
	return o.printWithMetaData(contents)
}

func (o *jsonOutput) printWithMetaData(contents Contents) error {
	var results []interface{}
	for i, v := range contents.Values() {
		if !structs.IsStruct(v) {
			results = append(results, v)
			continue
		}

		mapValue := structs.Map(v)
		// zone
		if contents[i].Zone != "" {
			if _, ok := mapValue["Zone"]; !ok {
				mapValue["Zone"] = contents[i].Zone
			}
		}
		// ID
		if !contents[i].ID.IsEmpty() {
			if _, ok := mapValue["ID"]; !ok {
				mapValue["ID"] = contents[i].ID
			}
		}
		results = append(results, mapValue)
	}

	return o.printOutput(results)
}

func (o *jsonOutput) printWithQuery(values []interface{}) error {
	q, err := util.StringFromPathOrContent(o.query)
	if err != nil {
		return fmt.Errorf("JSONOutput:Query: loading query from %q Failed: %s", o.query, err)
	}

	results, err := query.Executor(o.driver)(values, q)
	if err != nil {
		return err
	}
	return o.printOutput(results)
}

func (o *jsonOutput) printOutput(v interface{}) error {
	data, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return fmt.Errorf("JSONOutput:Print: MarshalIndent failed: %s", err)
	}

	if _, err := o.out.Write(data); err != nil {
		return err
	}
	if _, err := fmt.Fprintln(o.out, ""); err != nil {
		return err
	}
	return nil
}
