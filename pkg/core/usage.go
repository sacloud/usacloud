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
	"strings"

	"github.com/sacloud/usacloud/pkg/version"
	"github.com/spf13/cobra"
)

const originalCommandsUsage = `Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}`

const commandUsageTemplate = ` === %s ===
%s
`

const commandUsageWrapperTemplate = `Available Commands:
%s`

func buildSubCommandsUsage(rootCmd *cobra.Command, resources []*Resource, appendManualImplCommands bool) string {
	line := "    %s %s"
	var usages []string
	for _, r := range resources {
		cmd := lookupCmd(rootCmd, r.Name)

		if cmd.IsAvailableCommand() {
			t := fmt.Sprintf("%%-%ds", cmd.NamePadding())
			name := fmt.Sprintf(t, cmd.Name())
			usages = append(usages, fmt.Sprintf(line, name, cmd.Short))
		}
	}
	if appendManualImplCommands {
		// completionを追加
		if cmd := lookupCmd(rootCmd, "completion"); cmd != nil {
			t := fmt.Sprintf("%%-%ds", cmd.NamePadding())
			name := fmt.Sprintf(t, cmd.Name())
			usages = append(usages, fmt.Sprintf(line, name, cmd.Short))
		}
		// update-selfを追加
		if cmd := lookupCmd(rootCmd, "update-self"); cmd != nil {
			t := fmt.Sprintf("%%-%ds", cmd.NamePadding())
			name := fmt.Sprintf(t, cmd.Name())
			usages = append(usages, fmt.Sprintf(line, name, cmd.Short))
		}
	}

	return strings.TrimRight(strings.Join(usages, "\n"), "\n")
}

func SetSubCommandsUsage(cmd *cobra.Command, commands []*CategorizedResources) {
	cmd.SetUsageTemplate("")
	var usages []string
	for _, c := range commands {
		usages = append(usages, fmt.Sprintf(commandUsageTemplate, c.Category.DisplayName, buildSubCommandsUsage(cmd, c.Resources, c.Category.Key == "other")))
	}
	usage := fmt.Sprintf(commandUsageWrapperTemplate, strings.TrimRight(strings.Join(usages, "\n"), "\n"))

	usageTemplate := strings.Replace(cmd.UsageTemplate(), originalCommandsUsage, usage, 1) +
		fmt.Sprintf("\nCopyright %s The Usacloud Authors\n", version.CopyrightYear)

	cmd.SetUsageTemplate(usageTemplate)
}

func buildCommandUsages(rootCmd *cobra.Command, commands []*Command) string {
	line := "    %s %s"
	var usages []string
	for _, c := range commands {
		cmd := lookupCmd(rootCmd, c.Name)

		if cmd.IsAvailableCommand() {
			t := fmt.Sprintf("%%-%ds", cmd.NamePadding())
			name := fmt.Sprintf(t, cmd.Name())
			usages = append(usages, fmt.Sprintf(line, name, cmd.Short))
		}
	}
	return strings.TrimRight(strings.Join(usages, "\n"), "\n")
}

func buildCommandsUsage(cmd *cobra.Command, commands []*CategorizedCommands) {
	cmd.SetUsageTemplate("")
	var usages []string
	for _, c := range commands {
		usages = append(usages, fmt.Sprintf(commandUsageTemplate, c.Category.DisplayName, buildCommandUsages(cmd, c.Commands)))
	}
	usage := fmt.Sprintf(commandUsageWrapperTemplate, strings.TrimRight(strings.Join(usages, "\n"), "\n"))
	cmd.SetUsageTemplate(strings.Replace(cmd.UsageTemplate(), originalCommandsUsage, usage, 1))
}

const originalFlagsUsage = `Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}`

const flagsUsageTemplate = `  === %s ===

%s`

const flagsUsageWrapperTemplate = `Flags:

%s`

func BuildFlagsUsage(cmd *cobra.Command, sets []*FlagSet) {
	var usages []string
	for _, fs := range sets {
		usages = append(usages, fmt.Sprintf(flagsUsageTemplate, fs.Title, fs.Flags.FlagUsages()))
	}
	usage := fmt.Sprintf(flagsUsageWrapperTemplate, strings.TrimRight(strings.Join(usages, "\n"), "\n"))
	cmd.SetUsageTemplate(strings.Replace(cmd.UsageTemplate(), originalFlagsUsage, usage, 1))
}
