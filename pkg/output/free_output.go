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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/sacloud/usacloud/pkg/util"

	"github.com/bitly/go-simplejson"
)

type freeOutput struct {
	Out        io.Writer
	Err        io.Writer
	Format     string
	FormatFile string
}

func NewFreeOutput(out io.Writer, err io.Writer, option Option) Output {
	return &freeOutput{
		Out:        out,
		Err:        err,
		Format:     option.FormatFlagValue(),
		FormatFile: option.FormatFileFlagValue(),
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

	if o.FormatFile != "" {
		format, err := ioutil.ReadFile(o.FormatFile)
		if err != nil {
			return fmt.Errorf("FreeOutput:Print: read format-file is failed: %s", err)
		}
		o.Format = string(format)
	}

	// targets -> byte[] -> []interface{}
	rawArray, err := json.Marshal(targets)
	if err != nil {
		return fmt.Errorf("FreeOutput:Print: json.Marshal is failed: %s", err)
	}

	j, err := simplejson.NewJson(rawArray)
	if err != nil {
		return fmt.Errorf("FreeOutput:Print: create simplejson is failed: %s", err)
	}

	t := template.New("t")
	_, err = t.Parse(o.Format)
	if err != nil {
		return fmt.Errorf("Output format is invalid: %s", err)
	}

	for i := 0; i < len(targets); i++ {
		// interface{} -> map[string]interface{}
		v := j.GetIndex(i)
		mapValue, err := v.Map()
		if err != nil {
			return fmt.Errorf("FreeOutput:Print: json format is invalid: %v", err)
		}
		mapValue["RowNumber"] = fmt.Sprintf("%d", i+1)
		if _, ok := mapValue["Zone"]; !ok {
			mapValue["Zone"] = contents[i].Zone
		}

		buf := bytes.NewBufferString("")
		err = t.Execute(buf, mapValue)
		if err != nil {
			return err
		}

		o.Out.Write(buf.Bytes()) // nolint
		fmt.Fprintln(o.Out, "")
	}

	return nil
}
