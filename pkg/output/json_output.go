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

	"github.com/fatih/structs"
	"github.com/jmespath/go-jmespath"
	"github.com/sacloud/usacloud/pkg/util"
)

type jsonOutput struct {
	out   io.Writer
	err   io.Writer
	query string
}

func NewJSONOutput(out io.Writer, err io.Writer, query string) Output {
	return &jsonOutput{
		out:   out,
		err:   err,
		query: query,
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

	if o.query != "" {
		query, err := util.StringFromPathOrContent(o.query)
		if err != nil {
			return fmt.Errorf("JSONOutput:Query: loading query from %q Failed: %s", o.query, err)
		}

		v, err := o.searchByJMESPath(targets, query)
		if err != nil {
			return fmt.Errorf("JSONOutput:Query: jmespath.Search is Failed: %s", err)
		}

		switch v := v.(type) {
		case []interface{}:
			targets = v
		default:
			targets = []interface{}{v}
		}
	}

	var results []interface{}
	for i, v := range targets {
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

	b, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return fmt.Errorf("JSONOutput:Print: MarshalIndent failed: %s", err)
	}

	o.out.Write(b) // nolint
	fmt.Fprintln(o.out, "")
	return nil
}

func (o *jsonOutput) searchByJMESPath(v interface{}, query string) (result interface{}, err error) {
	defer func() {
		ret := recover()
		if ret != nil {
			fmt.Fprintf(o.err, "jmespath.Search failed: parse error\n")
			err = fmt.Errorf("jmespath.Search failed: %s", ret)
		}
	}()
	result, err = jmespath.Search(query, v)
	return
}
