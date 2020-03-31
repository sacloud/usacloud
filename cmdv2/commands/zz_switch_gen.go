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
	switchListParam             = params.NewListSwitchParam()
	switchCreateParam           = params.NewCreateSwitchParam()
	switchReadParam             = params.NewReadSwitchParam()
	switchUpdateParam           = params.NewUpdateSwitchParam()
	switchDeleteParam           = params.NewDeleteSwitchParam()
	switchBridgeConnectParam    = params.NewBridgeConnectSwitchParam()
	switchBridgeDisconnectParam = params.NewBridgeDisconnectSwitchParam()
)

// switchCmd represents the command to manage SAKURA Cloud Switch
var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "A manage commands of Switch",
	Long:  `A manage commands of Switch`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

var switchListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "find", "selector"},
	Short:   "List Switch",
	Long:    `List Switch`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := switchListParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("list parameter: \n%s\n", debugMarshalIndent(switchListParam))
		return err
	},
}

func switchListCmdInit() {
	fs := switchListCmd.Flags()
	fs.StringSliceVarP(&switchListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &switchListParam.Id), "id", "", "set filter by id(s)")
	fs.IntVarP(&switchListParam.From, "from", "", 0, "set offset")
	fs.StringSliceVarP(&switchListParam.Tags, "tags", "", []string{}, "set filter by tags(AND)")
	fs.IntVarP(&switchListParam.Max, "max", "", 0, "set limit")
	fs.StringSliceVarP(&switchListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
}

var switchCreateCmd = &cobra.Command{
	Use: "create",

	Short: "Create Switch",
	Long:  `Create Switch`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := switchCreateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("create parameter: \n%s\n", debugMarshalIndent(switchCreateParam))
		return err
	},
}

func switchCreateCmdInit() {
	fs := switchCreateCmd.Flags()
	fs.StringVarP(&switchCreateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&switchCreateParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&switchCreateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &switchCreateParam.IconId), "icon-id", "", "set Icon ID")
}

var switchReadCmd = &cobra.Command{
	Use: "read",

	Short: "Read Switch",
	Long:  `Read Switch`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := switchReadParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("read parameter: \n%s\n", debugMarshalIndent(switchReadParam))
		return err
	},
}

func switchReadCmdInit() {
}

var switchUpdateCmd = &cobra.Command{
	Use: "update",

	Short: "Update Switch",
	Long:  `Update Switch`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := switchUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("update parameter: \n%s\n", debugMarshalIndent(switchUpdateParam))
		return err
	},
}

func switchUpdateCmdInit() {
	fs := switchUpdateCmd.Flags()
	fs.StringVarP(&switchUpdateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&switchUpdateParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&switchUpdateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &switchUpdateParam.IconId), "icon-id", "", "set Icon ID")
}

var switchDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"rm"},
	Short:   "Delete Switch",
	Long:    `Delete Switch`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := switchDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("delete parameter: \n%s\n", debugMarshalIndent(switchDeleteParam))
		return err
	},
}

func switchDeleteCmdInit() {
}

var switchBridgeConnectCmd = &cobra.Command{
	Use: "bridge-connect",

	Short: "BridgeConnect Switch",
	Long:  `BridgeConnect Switch`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := switchBridgeConnectParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("bridge-connect parameter: \n%s\n", debugMarshalIndent(switchBridgeConnectParam))
		return err
	},
}

func switchBridgeConnectCmdInit() {
	fs := switchBridgeConnectCmd.Flags()
	fs.VarP(newIDValue(0, &switchBridgeConnectParam.BridgeId), "bridge-id", "", "set Bridge ID")
}

var switchBridgeDisconnectCmd = &cobra.Command{
	Use: "bridge-disconnect",

	Short: "BridgeDisconnect Switch",
	Long:  `BridgeDisconnect Switch`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := switchBridgeDisconnectParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("bridge-disconnect parameter: \n%s\n", debugMarshalIndent(switchBridgeDisconnectParam))
		return err
	},
}

func switchBridgeDisconnectCmdInit() {
}

func init() {
	parent := switchCmd

	switchListCmdInit()
	parent.AddCommand(switchListCmd)

	switchCreateCmdInit()
	parent.AddCommand(switchCreateCmd)

	switchReadCmdInit()
	parent.AddCommand(switchReadCmd)

	switchUpdateCmdInit()
	parent.AddCommand(switchUpdateCmd)

	switchDeleteCmdInit()
	parent.AddCommand(switchDeleteCmd)

	switchBridgeConnectCmdInit()
	parent.AddCommand(switchBridgeConnectCmd)

	switchBridgeDisconnectCmdInit()
	parent.AddCommand(switchBridgeDisconnectCmd)

	rootCmd.AddCommand(parent)
}
