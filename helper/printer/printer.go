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

package printer

import (
	"fmt"
	"io"

	"github.com/fatih/color"
	"github.com/sacloud/usacloud/command"
)

// Fprint is delegates to *color.Color or fmt depending on command.GlobalOption.NoColor flag
func Fprint(w io.Writer, c *color.Color, a ...interface{}) {
	if command.GlobalOption.NoColor {
		fmt.Fprint(w, a...)
	} else {
		c.Fprint(w, a...)
	}
}

// Fprintf is delegates to *color.Color or fmt depending on command.GlobalOption.NoColor flag
func Fprintf(w io.Writer, c *color.Color, format string, a ...interface{}) {
	if command.GlobalOption.NoColor {
		fmt.Fprintf(w, format, a...)
	} else {
		c.Fprintf(w, format, a...)
	}
}
