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

package test

import (
	"bytes"
	"testing"

	"github.com/sacloud/usacloud/pkg/output"
)

type ColumnDefTestTargets struct {
	ColumnDefs []output.ColumnDef
	Source     interface{}
	Tests      []*ColumnDefTestTarget
}

func (c *ColumnDefTestTargets) ColumnDef(name string) (output.ColumnDef, bool) {
	for _, t := range c.ColumnDefs {
		if t.Name == name {
			return t, true
		}
	}
	return output.ColumnDef{}, false
}

type ColumnDefTestTarget struct {
	ColumnName string
	Expect     string
}

func RunColumnDefTest(t *testing.T, targets ColumnDefTestTargets) {
	tableWriter := output.NewSimpleTableWriter(bytes.NewBufferString(""), targets.ColumnDefs)

	for _, target := range targets.Tests {
		def, ok := targets.ColumnDef(target.ColumnName)
		if !ok {
			t.Fatalf("columnDef %q not exists", target.ColumnName)
		}

		got, err := tableWriter.CellValue(targets.Source, def)
		if err != nil {
			t.Fatalf("got unexpected error: in: %v error: %s", target, err)
		}

		if got != target.Expect {
			t.Fatalf("got unexpected value: in: %v expected: %s got: %s", target, target.Expect, got)
		}
	}
}
