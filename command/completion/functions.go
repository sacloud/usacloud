// Copyright 2017-2019 The Usacloud Authors
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

package completion

import (
	"fmt"
	"strconv"

	"gopkg.in/urfave/cli.v2"
)

func FlagNames(c *cli.Context, commandName string) {
	comm := c.App.Command(commandName)
	if comm == nil {
		return
	}
	for _, f := range comm.VisibleFlags() {
		for _, n := range f.Names() {
			format := "--%s\n"
			if len(n) == 1 {
				format = "-%s\n"
			}
			fmt.Printf(format, n)
		}
	}
}

func isSakuraID(id string) bool {
	_, err := strconv.ParseInt(id, 10, 64)
	return err == nil
}
