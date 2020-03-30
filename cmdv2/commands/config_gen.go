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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-v2-commands'; DO NOT EDIT

package commands

import (
	"fmt"

	"github.com/sacloud/usacloud/cmdv2/params"
	"github.com/spf13/cobra"
)

var (
	configCurrentParam = params.NewCurrentConfigParam()
	configDeleteParam  = params.NewDeleteConfigParam()
	configEditParam    = params.NewEditConfigParam()
	configListParam    = params.NewListConfigParam()
	configMigrateParam = params.NewMigrateConfigParam()
	configShowParam    = params.NewShowConfigParam()
	configUseParam     = params.NewUseConfigParam()
)

// configCmd represents the command to manage SAKURA Cloud Config
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A manage command of APIKey settings",
	Long:  `A manage command of APIKey settings`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO not implements: call edit func as default
	},
}

var configCurrentCmd = &cobra.Command{
	Use: "current",

	Short: "Current Config",
	Long:  `Current Config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := configCurrentParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("current parameter: \n%s\n", debugMarshalIndent(configCurrentParam))
		return err
	},
}

func configCurrentCmdInit() {
}

var configDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"rm"},
	Short:   "Delete Config",
	Long:    `Delete Config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := configDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("delete parameter: \n%s\n", debugMarshalIndent(configDeleteParam))
		return err
	},
}

func configDeleteCmdInit() {
}

var configEditCmd = &cobra.Command{
	Use: "edit",

	Short: "Edit Config (default)",
	Long:  `Edit Config (default)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := configEditParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("edit parameter: \n%s\n", debugMarshalIndent(configEditParam))
		return err
	},
}

func configEditCmdInit() {
	fs := configEditCmd.Flags()
	fs.StringVarP(&configEditParam.DefaultOutputType, "default-output-type", "", "", "Default output format type")
	fs.StringVarP(&configEditParam.Token, "token", "", "", "API Token of SakuraCloud")
	fs.StringVarP(&configEditParam.Secret, "secret", "", "", "API Secret of SakuraCloud")
	fs.StringVarP(&configEditParam.Zone, "zone", "", "", "Target zone of SakuraCloud")
}

var configListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List Config",
	Long:    `List Config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := configListParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("list parameter: \n%s\n", debugMarshalIndent(configListParam))
		return err
	},
}

func configListCmdInit() {
}

var configMigrateCmd = &cobra.Command{
	Use: "migrate",

	Short: "Migrate Config",
	Long:  `Migrate Config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := configMigrateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("migrate parameter: \n%s\n", debugMarshalIndent(configMigrateParam))
		return err
	},
}

func configMigrateCmdInit() {
}

var configShowCmd = &cobra.Command{
	Use: "show",

	Short: "Show Config",
	Long:  `Show Config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := configShowParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("show parameter: \n%s\n", debugMarshalIndent(configShowParam))
		return err
	},
}

func configShowCmdInit() {
}

var configUseCmd = &cobra.Command{
	Use: "use",

	Short: "Use Config",
	Long:  `Use Config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := configUseParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("use parameter: \n%s\n", debugMarshalIndent(configUseParam))
		return err
	},
}

func configUseCmdInit() {
}

func init() {
	parent := configCmd

	configCurrentCmdInit()
	parent.AddCommand(configCurrentCmd)

	configDeleteCmdInit()
	parent.AddCommand(configDeleteCmd)

	configEditCmdInit()
	parent.AddCommand(configEditCmd)

	configListCmdInit()
	parent.AddCommand(configListCmd)

	configMigrateCmdInit()
	parent.AddCommand(configMigrateCmd)

	configShowCmdInit()
	parent.AddCommand(configShowCmd)

	configUseCmdInit()
	parent.AddCommand(configUseCmd)

	rootCmd.AddCommand(parent)
}
