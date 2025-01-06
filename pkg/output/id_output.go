// Copyright 2017-2025 The sacloud/usacloud Authors
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
	"github.com/sacloud/usacloud/pkg/util"
)

type idOutput struct {
	Out io.Writer
	Err io.Writer
}

var idOutputTargetColumns = []string{"ID", "Current.ID", "Key", "AccountID", "BillID", "Index", "RowNumber"}

func NewIDOutput(out io.Writer, err io.Writer) Output {
	return &idOutput{
		Out: out,
		Err: err,
	}
}

func (o *idOutput) Print(contents Contents) error {
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
	for i, v := range targets {
		if !structs.IsStruct(v) {
			continue
		}
		mapValue := structs.Map(v)

		var value interface{}
		for _, key := range idOutputTargetColumns {
			if key == "RowNumber" {
				value = fmt.Sprintf("%d", i+1)
				break
			} else {
				if id, ok := mapValue[key]; ok {
					value = id
					break
				}
			}
		}

		fmt.Fprintf(o.Out, "%v\n", value)
	}

	return nil
}
