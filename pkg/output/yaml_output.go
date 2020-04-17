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
	"fmt"
	"io"
	"os"

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

func (o *yamlOutput) Print(targets ...interface{}) error {
	if o.out == nil {
		o.out = os.Stdout
	}
	if o.err == nil {
		o.err = os.Stderr
	}

	if len(targets) == 0 {
		fmt.Fprintf(o.err, "Result is empty\n")
		return nil
	}

	b, err := yaml.Marshal(targets)
	if err != nil {
		return fmt.Errorf("YAMLOutput:Print: yaml.Marshal is Failed: %s", err)
	}
	o.out.Write(b)
	fmt.Fprintln(o.out, "")
	return nil

}
