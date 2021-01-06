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
	"fmt"
	"io"
	"os"

	"github.com/fatih/structs"
	"github.com/ghodss/yaml"
	"github.com/sacloud/usacloud/pkg/util"
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

	var results []interface{}
	for i, v := range targets {
		if !structs.IsStruct(v) {
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

	b, err := yaml.Marshal(results)
	if err != nil {
		return fmt.Errorf("YAMLOutput:Print: yaml.Marshal is Failed: %s", err)
	}
	o.out.Write(b)          // nolint
	fmt.Fprintln(o.out, "") // nolint
	return nil
}
