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

package base

import (
	"log"
	"reflect"
	"sort"

	"github.com/spf13/cobra"
)

type Resource struct {
	Name    string
	Aliases []string
	Usage   string

	CommandCategories   []Category
	categorizedCommands []*CategorizedCommands
	DefaultCommandName  string

	Category            Category
	SkipApplyConfigFile bool

	Warning string

	IsGlobalResource bool

	ServiceType reflect.Type // リソースに対応するlibsacloud serviceの型情報、コード生成用
}

func (r *Resource) CLICommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     r.Name,
		Aliases: r.Aliases,
		Short:   r.Usage,
		Long:    r.Usage,
		RunE: func(cmd *cobra.Command, args []string) error {
			if r.DefaultCommandName == "" {
				cmd.HelpFunc()(cmd, args)
				return nil
			}
			return r.runDefaultCmd(cmd, args, r.DefaultCommandName)
		},
	}
	for _, c := range r.Commands() {
		subCmd := c.CLICommand()
		cmd.AddCommand(subCmd)

		// フラグの引き継ぎ
		if c.Name == r.DefaultCommandName {
			parameter := c.ParameterInitializer().(FlagInitializer)
			parameter.SetupCobraCommandFlags(cmd)
		}
	}
	return cmd
}

func (r *Resource) runDefaultCmd(cmd *cobra.Command, args []string, commandName string) error {
	defaultCmd := lookupCmd(cmd, commandName)
	if defaultCmd == nil {
		cmd.HelpFunc()(cmd, args)
		return nil
	}
	return defaultCmd.RunE(defaultCmd, args)
}

func (r *Resource) Commands() []*Command {
	var commands []*Command
	for _, cat := range r.categorizedCommands {
		commands = append(commands, cat.Commands...)
	}
	return commands
}

func (r *Resource) AddCommand(command *Command) {
	command.resource = r

	categoryKey := command.Category
	category := r.commandCategory(categoryKey)
	if category == nil {
		log.Fatalf("resource %q does not have category %q", r.Name, categoryKey)
	}

	found := false
	for _, cat := range r.categorizedCommands {
		if cat.Category.Equals(category) {
			// 同じNameのコマンドが同一カテゴリーに存在したらエラーとする
			for _, c := range cat.Commands {
				if c.Name == command.Name {
					log.Fatalf("resource %q already has same command %q", r.Name, command.Name)
				}
			}
			cat.Commands = append(cat.Commands, command)
		}
		sort.Slice(cat.Commands, func(i, j int) bool {
			return cat.Commands[i].Order < cat.Commands[j].Order
		})
		found = true
	}

	if !found {
		r.categorizedCommands = append(r.categorizedCommands, &CategorizedCommands{
			Category: *category,
			Commands: []*Command{command},
		})
	}
}

func (r *Resource) commandCategory(key string) *Category {
	for _, c := range r.CommandCategories {
		if c.Key == key {
			return &c
		}
	}
	return nil
}
