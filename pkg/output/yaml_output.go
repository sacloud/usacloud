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

	"github.com/sacloud/usacloud/pkg/util"

	"github.com/ghodss/yaml"
)

type yamlOutput struct {
	out io.Writer
	err io.Writer
}

func NewYAMLOutput(out io.Writer, err io.Writer) Output {
	return &yamlOutput{
		out: out,
		err: err,
	}
}

func (o *yamlOutput) Print(contents Contents) error {
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

	// HACK: ゾーンの値を追加するためにsimplejsonにして操作する
	// targets -> byte[] -> []interface{}
	rawArray, err := json.Marshal(targets)
	if err != nil {
		return fmt.Errorf("YAMLOutput:Print: json.Marshal is failed: %s", err)
	}

	j, err := simplejson.NewJson(rawArray)
	if err != nil {
		return fmt.Errorf("YAMLOutput:Print: create simplejson is failed: %s", err)
	}

	for i := 0; i < len(targets); i++ {
		if contents[i].Zone != "" {
			row := j.GetIndex(i)
			if _, ok := row.CheckGet("Zone"); !ok {
				row.Set("Zone", contents[i].Zone)
			}
		}
		if !contents[i].ID.IsEmpty() {
			row := j.GetIndex(i)
			if _, ok := row.CheckGet("ID"); !ok {
				row.Set("ID", contents[i].ID)
			}
		}
	}

	b, err := yaml.Marshal(j)
	if err != nil {
		return fmt.Errorf("YAMLOutput:Print: yaml.Marshal is Failed: %s", err)
	}
	o.out.Write(b)          // nolint
	fmt.Fprintln(o.out, "") // nolint
	return nil
}
