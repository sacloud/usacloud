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

package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const originalCommandsUsage = `Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}`

const commandUsageTemplate = `  <%s>
%s
`

const commandUsageWrapperTemplate = `Available Commands:
%s`

type commandSet struct {
	title    string
	commands []*cobra.Command
}

func (c *commandSet) CommandUsages() string {
	line := "    %s %s"
	var usages []string
	for _, co := range c.commands {
		if co.IsAvailableCommand() {
			t := fmt.Sprintf("%%-%ds", co.NamePadding())
			name := fmt.Sprintf(t, co.Name())
			usages = append(usages, fmt.Sprintf(line, name, co.Short))
		}
	}
	return strings.TrimRight(strings.Join(usages, "\n"), "\n")
}

func buildCommandsUsage(cmd *cobra.Command, commands []*commandSet) {
	cmd.SetUsageTemplate("")
	var usages []string
	for _, c := range commands {
		usages = append(usages, fmt.Sprintf(commandUsageTemplate, c.title, c.CommandUsages()))
	}
	usage := fmt.Sprintf(commandUsageWrapperTemplate, strings.TrimRight(strings.Join(usages, "\n"), "\n"))
	cmd.SetUsageTemplate(strings.Replace(cmd.UsageTemplate(), originalCommandsUsage, usage, 1))
}

func lookupCmd(cmd *cobra.Command, name string) *cobra.Command {
	for _, c := range cmd.Commands() {
		if c.Name() == name {
			return c
		}
	}
	return nil
}

func runDefaultCmd(cmd *cobra.Command, args []string, commandName string) error {
	defaultCmd := lookupCmd(cmd, commandName)
	if defaultCmd == nil {
		cmd.HelpFunc()(cmd, args)
		return nil
	}
	return defaultCmd.RunE(defaultCmd, args)
}

const originalFlagsUsage = `Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}`

const flagsUsageTemplate = `  <%s>
%s`

const flagsUsageWrapperTemplate = `Flags:
%s`

type flagSet struct {
	title string
	flags *pflag.FlagSet
}

func buildFlagsUsage(cmd *cobra.Command, sets []*flagSet) {
	var usages []string
	for _, fs := range sets {
		usages = append(usages, fmt.Sprintf(flagsUsageTemplate, fs.title, fs.flags.FlagUsages()))
	}
	usage := fmt.Sprintf(flagsUsageWrapperTemplate, strings.TrimRight(strings.Join(usages, "\n"), "\n"))
	cmd.SetUsageTemplate(strings.Replace(cmd.UsageTemplate(), originalFlagsUsage, usage, 1))
}
