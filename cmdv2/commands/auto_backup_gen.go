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

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/cmdv2/params"
	"github.com/spf13/cobra"
)

var (
	autobackupListParam   = params.NewListAutobackupParam()
	autobackupCreateParam = params.NewCreateAutobackupParam()
	autobackupReadParam   = params.NewReadAutobackupParam()
	autobackupUpdateParam = params.NewUpdateAutobackupParam()
	autobackupDeleteParam = params.NewDeleteAutobackupParam()
)

// autobackupCmd represents the command to manage SAKURA Cloud AutoBackup
var autobackupCmd = &cobra.Command{
	Use:   "autobackup",
	Short: "A manage commands of AutoBackup",
	Long:  `A manage commands of AutoBackup`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

var autobackupListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "find", "selector"},
	Short:   "List Autobackup",
	Long:    `List Autobackup`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := autobackupListParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("list parameter: \n%s\n", debugMarshalIndent(autobackupListParam))
		return err
	},
}

func autobackupListCmdInit() {
	fs := autobackupListCmd.Flags()
	fs.IntVarP(&autobackupListParam.Max, "max", "", 0, "set limit")
	fs.StringSliceVarP(&autobackupListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringSliceVarP(&autobackupListParam.Tags, "tags", "", []string{}, "set filter by tags(AND)")
	fs.StringSliceVarP(&autobackupListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &autobackupListParam.Id), "id", "", "set filter by id(s)")
	fs.IntVarP(&autobackupListParam.From, "from", "", 0, "set offset")
}

var autobackupCreateCmd = &cobra.Command{
	Use: "create",

	Short: "Create Autobackup",
	Long:  `Create Autobackup`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := autobackupCreateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("create parameter: \n%s\n", debugMarshalIndent(autobackupCreateParam))
		return err
	},
}

func autobackupCreateCmdInit() {
	fs := autobackupCreateCmd.Flags()
	fs.IntVarP(&autobackupCreateParam.Generation, "generation", "", 1, "set backup generation[1-10]")
	fs.StringVarP(&autobackupCreateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&autobackupCreateParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&autobackupCreateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &autobackupCreateParam.IconId), "icon-id", "", "set Icon ID")
	fs.VarP(newIDValue(0, &autobackupCreateParam.DiskId), "disk-id", "", "set target diskID ")
	fs.StringSliceVarP(&autobackupCreateParam.Weekdays, "weekdays", "", []string{"all"}, "set backup target weekdays[all or mon/tue/wed/thu/fri/sat/sun]")
}

var autobackupReadCmd = &cobra.Command{
	Use: "read",

	Short: "Read Autobackup",
	Long:  `Read Autobackup`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := autobackupReadParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("read parameter: \n%s\n", debugMarshalIndent(autobackupReadParam))
		return err
	},
}

func autobackupReadCmdInit() {
}

var autobackupUpdateCmd = &cobra.Command{
	Use: "update",

	Short: "Update Autobackup",
	Long:  `Update Autobackup`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := autobackupUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("update parameter: \n%s\n", debugMarshalIndent(autobackupUpdateParam))
		return err
	},
}

func autobackupUpdateCmdInit() {
	fs := autobackupUpdateCmd.Flags()
	fs.IntVarP(&autobackupUpdateParam.Generation, "generation", "", 0, "set backup generation[1-10]")
	fs.StringVarP(&autobackupUpdateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&autobackupUpdateParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&autobackupUpdateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &autobackupUpdateParam.IconId), "icon-id", "", "set Icon ID")
	fs.StringSliceVarP(&autobackupUpdateParam.Weekdays, "weekdays", "", []string{}, "set backup target weekdays[all or mon/tue/wed/thu/fri/sat/sun]")
}

var autobackupDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"rm"},
	Short:   "Delete Autobackup",
	Long:    `Delete Autobackup`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := autobackupDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("delete parameter: \n%s\n", debugMarshalIndent(autobackupDeleteParam))
		return err
	},
}

func autobackupDeleteCmdInit() {
}

func init() {
	parent := autobackupCmd

	autobackupListCmdInit()
	parent.AddCommand(autobackupListCmd)

	autobackupCreateCmdInit()
	parent.AddCommand(autobackupCreateCmd)

	autobackupReadCmdInit()
	parent.AddCommand(autobackupReadCmd)

	autobackupUpdateCmdInit()
	parent.AddCommand(autobackupUpdateCmd)

	autobackupDeleteCmdInit()
	parent.AddCommand(autobackupDeleteCmd)

	rootCmd.AddCommand(parent)
}
