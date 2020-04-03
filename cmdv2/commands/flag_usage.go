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

package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const originalFlagsUsage = `Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}`

func setFlagsUsage(cmd *cobra.Command, usage string) {
	cmd.SetUsageTemplate(strings.Replace(cmd.UsageTemplate(), originalFlagsUsage, usage, 1))
}

const flagsUsageTemplate = `%s:
%s`

func buildFlagsUsage(sets []*flagSet) string {
	var usage []string
	for _, fs := range sets {
		usage = append(usage, fmt.Sprintf(flagsUsageTemplate, fs.title, fs.flags.FlagUsages()))
	}
	return strings.TrimRight(strings.Join(usage, "\n"), "\n")
}

type flagSet struct {
	title string
	flags *pflag.FlagSet
}
