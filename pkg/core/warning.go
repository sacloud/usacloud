// Copyright 2017-2023 The sacloud/usacloud Authors
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

package core

import (
	"fmt"
	"io"

	"github.com/fatih/color"
)

func (c *Command) printWarning(writer io.Writer, noColor bool, warn string) {
	if warn == "" {
		return
	}
	if noColor {
		fmt.Fprintf(writer, "[WARN] %s\n", warn)
	} else {
		out := color.New(color.FgYellow)
		out.Fprintf(writer, "[WARN] %s\n", warn)
	}
}
