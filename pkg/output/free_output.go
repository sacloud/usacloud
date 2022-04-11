// Copyright 2017-2022 The Usacloud Authors
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
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/fatih/structs"
	"github.com/sacloud/usacloud/pkg/util"
)

type freeOutput struct {
	Out    io.Writer
	Err    io.Writer
	Format string
}

func NewFreeOutput(out io.Writer, err io.Writer, option Option) Output {
	return &freeOutput{
		Out:    out,
		Err:    err,
		Format: option.FormatFlagValue(),
	}
}

func (o *freeOutput) Print(contents Contents) error {
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

	format, err := util.StringFromPathOrContent(o.Format)
	if err != nil {
		return fmt.Errorf("FreeOutput:Print: load format from %q failed failed: %s", o.Format, err)
	}

	t, err := newTemplate().Parse(format)
	if err != nil {
		return fmt.Errorf("invalid output format %q: %s", format, err)
	}

	for i, v := range targets {
		if !structs.IsStruct(v) {
			continue
		}
		mapValue := structs.Map(v)

		// original value
		if _, ok := mapValue["OriginalValue"]; !ok {
			mapValue["OriginalValue"] = v
		}
		mapValue["RowNumber"] = fmt.Sprintf("%d", i+1)
		mapValue["__ORDER__"] = fmt.Sprintf("%d", i+1)
		if contents[i].Zone != "" {
			if _, ok := mapValue["Zone"]; !ok {
				mapValue["Zone"] = contents[i].Zone
			}
		}
		if contents[i].ID != "" {
			if _, ok := mapValue["ID"]; !ok {
				mapValue["ID"] = contents[i].ID
			}
		}

		buf := bytes.NewBufferString("")
		if err := t.Execute(buf, mapValue); err != nil {
			return err
		}

		o.Out.Write(buf.Bytes()) // nolint
		fmt.Fprintln(o.Out, "")
	}

	return nil
}
