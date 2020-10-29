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

	"github.com/bitly/go-simplejson"
	"github.com/sacloud/go-jmespath"
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

func (o *jsonOutput) Print(target interface{}) error {
	targets := toSlice(target)
	if o.out == nil {
		o.out = os.Stdout
	}
	if o.err == nil {
		o.err = os.Stderr
	}

	if util.IsEmpty(targets) {
		fmt.Fprintln(o.err, "no results")
		return nil
	}

	var values interface{} = targets

	if o.query != "" {
		v, err := o.searchByJMESPath(targets)
		if err != nil {
			return fmt.Errorf("JSONOutput:Query: jmespath.Search is Failed: %s", err)
		}
		values = v
	}

	rawArray, err := json.Marshal(values)
	if err != nil {
		return fmt.Errorf("JSONOutput:Print: json.Marshal is Failed: %s", err)
	}

	j, err := simplejson.NewJson(rawArray)

	if err != nil {
		return fmt.Errorf("JSONOutput:Print: Create SimpleJSON object is failed: %s", err)
	}

	b, err := j.EncodePretty()
	if err != nil {
		return fmt.Errorf("JSONOutput:Print: Print pretty JSON is failed: %s", err)
	}
	o.out.Write(b) // nolint
	fmt.Fprintln(o.out, "")
	return nil
}

func (o *jsonOutput) searchByJMESPath(v interface{}) (result interface{}, err error) {
	defer func() {
		ret := recover()
		if ret != nil {
			fmt.Fprintf(o.err, "jmespath.Search failed: parse error\n")
			err = fmt.Errorf("jmespath.Search failed: %s", ret)
		}
	}()
	result, err = jmespath.Search(o.query, v)
	return
}
