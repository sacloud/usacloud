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
	packetfilterListParam                = params.NewListPacketfilterParam()
	packetfilterCreateParam              = params.NewCreatePacketfilterParam()
	packetfilterReadParam                = params.NewReadPacketfilterParam()
	packetfilterUpdateParam              = params.NewUpdatePacketfilterParam()
	packetfilterDeleteParam              = params.NewDeletePacketfilterParam()
	packetfilterRuleInfoParam            = params.NewRuleInfoPacketfilterParam()
	packetfilterRuleAddParam             = params.NewRuleAddPacketfilterParam()
	packetfilterRuleUpdateParam          = params.NewRuleUpdatePacketfilterParam()
	packetfilterRuleDeleteParam          = params.NewRuleDeletePacketfilterParam()
	packetfilterInterfaceConnectParam    = params.NewInterfaceConnectPacketfilterParam()
	packetfilterInterfaceDisconnectParam = params.NewInterfaceDisconnectPacketfilterParam()
)

// packetfilterCmd represents the command to manage SAKURA Cloud PacketFilter
var packetfilterCmd = &cobra.Command{
	Use:   "packetfilter",
	Short: "A manage commands of PacketFilter",
	Long:  `A manage commands of PacketFilter`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

var packetfilterListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "find"},
	Short:   "List Packetfilter",
	Long:    `List Packetfilter`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := packetfilterListParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("list parameter: \n%s\n", debugMarshalIndent(packetfilterListParam))
		return err
	},
}

func packetfilterListCmdInit() {
	fs := packetfilterListCmd.Flags()
	fs.IntVarP(&packetfilterListParam.Max, "max", "", 0, "set limit")
	fs.StringSliceVarP(&packetfilterListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringSliceVarP(&packetfilterListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &packetfilterListParam.Id), "id", "", "set filter by id(s)")
	fs.IntVarP(&packetfilterListParam.From, "from", "", 0, "set offset")
}

var packetfilterCreateCmd = &cobra.Command{
	Use: "create",

	Short: "Create Packetfilter",
	Long:  `Create Packetfilter`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := packetfilterCreateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("create parameter: \n%s\n", debugMarshalIndent(packetfilterCreateParam))
		return err
	},
}

func packetfilterCreateCmdInit() {
	fs := packetfilterCreateCmd.Flags()
	fs.StringVarP(&packetfilterCreateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&packetfilterCreateParam.Description, "description", "", "", "set resource description")
}

var packetfilterReadCmd = &cobra.Command{
	Use: "read",

	Short: "Read Packetfilter",
	Long:  `Read Packetfilter`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := packetfilterReadParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("read parameter: \n%s\n", debugMarshalIndent(packetfilterReadParam))
		return err
	},
}

func packetfilterReadCmdInit() {
}

var packetfilterUpdateCmd = &cobra.Command{
	Use: "update",

	Short: "Update Packetfilter",
	Long:  `Update Packetfilter`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := packetfilterUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("update parameter: \n%s\n", debugMarshalIndent(packetfilterUpdateParam))
		return err
	},
}

func packetfilterUpdateCmdInit() {
	fs := packetfilterUpdateCmd.Flags()
	fs.StringVarP(&packetfilterUpdateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&packetfilterUpdateParam.Description, "description", "", "", "set resource description")
}

var packetfilterDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"rm"},
	Short:   "Delete Packetfilter",
	Long:    `Delete Packetfilter`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := packetfilterDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("delete parameter: \n%s\n", debugMarshalIndent(packetfilterDeleteParam))
		return err
	},
}

func packetfilterDeleteCmdInit() {
}

var packetfilterRuleInfoCmd = &cobra.Command{
	Use:     "rule-info",
	Aliases: []string{"rules", "rule-list"},
	Short:   "RuleInfo Packetfilter",
	Long:    `RuleInfo Packetfilter`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := packetfilterRuleInfoParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("rule-info parameter: \n%s\n", debugMarshalIndent(packetfilterRuleInfoParam))
		return err
	},
}

func packetfilterRuleInfoCmdInit() {
}

var packetfilterRuleAddCmd = &cobra.Command{
	Use: "rule-add",

	Short: "RuleAdd Packetfilter",
	Long:  `RuleAdd Packetfilter`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := packetfilterRuleAddParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("rule-add parameter: \n%s\n", debugMarshalIndent(packetfilterRuleAddParam))
		return err
	},
}

func packetfilterRuleAddCmdInit() {
	fs := packetfilterRuleAddCmd.Flags()
	fs.StringVarP(&packetfilterRuleAddParam.SourceNetwork, "source-network", "", "", "set source network[A.A.A.A] or [A.A.A.A/N (N=1..31)] or [A.A.A.A/M.M.M.M]")
	fs.StringVarP(&packetfilterRuleAddParam.SourcePort, "source-port", "", "", "set source port[N (N=0..65535)] or [N-N (N=0..65535)] or [0xPPPP/0xMMMM]")
	fs.StringVarP(&packetfilterRuleAddParam.DestinationPort, "destination-port", "", "", "set destination port[N (N=0..65535)] or [N-N (N=0..65535)] or [0xPPPP/0xMMMM]")
	fs.StringVarP(&packetfilterRuleAddParam.Action, "action", "", "", "set action[allow/deny]")
	fs.StringVarP(&packetfilterRuleAddParam.Description, "description", "", "", "set resource description")
	fs.IntVarP(&packetfilterRuleAddParam.Index, "index", "", 1, "index to insert rule into")
	fs.StringVarP(&packetfilterRuleAddParam.Protocol, "protocol", "", "", "set target protocol[tcp/udp/icmp/fragment/ip]")
}

var packetfilterRuleUpdateCmd = &cobra.Command{
	Use: "rule-update",

	Short: "RuleUpdate Packetfilter",
	Long:  `RuleUpdate Packetfilter`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := packetfilterRuleUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("rule-update parameter: \n%s\n", debugMarshalIndent(packetfilterRuleUpdateParam))
		return err
	},
}

func packetfilterRuleUpdateCmdInit() {
	fs := packetfilterRuleUpdateCmd.Flags()
	fs.StringVarP(&packetfilterRuleUpdateParam.Action, "action", "", "", "set action[allow/deny]")
	fs.StringVarP(&packetfilterRuleUpdateParam.Description, "description", "", "", "set resource description")
	fs.IntVarP(&packetfilterRuleUpdateParam.Index, "index", "", 0, "index of target rule")
	fs.StringVarP(&packetfilterRuleUpdateParam.Protocol, "protocol", "", "", "set target protocol[tcp/udp/icmp/fragment/ip]")
	fs.StringVarP(&packetfilterRuleUpdateParam.SourceNetwork, "source-network", "", "", "set source network[A.A.A.A] or [A.A.A.A/N (N=1..31)] or [A.A.A.A/M.M.M.M]")
	fs.StringVarP(&packetfilterRuleUpdateParam.SourcePort, "source-port", "", "", "set source port[N (N=0..65535)] or [N-N (N=0..65535)] or [0xPPPP/0xMMMM]")
	fs.StringVarP(&packetfilterRuleUpdateParam.DestinationPort, "destination-port", "", "", "set destination port[N (N=0..65535)] or [N-N (N=0..65535)] or [0xPPPP/0xMMMM]")
}

var packetfilterRuleDeleteCmd = &cobra.Command{
	Use: "rule-delete",

	Short: "RuleDelete Packetfilter",
	Long:  `RuleDelete Packetfilter`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := packetfilterRuleDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("rule-delete parameter: \n%s\n", debugMarshalIndent(packetfilterRuleDeleteParam))
		return err
	},
}

func packetfilterRuleDeleteCmdInit() {
	fs := packetfilterRuleDeleteCmd.Flags()
	fs.IntVarP(&packetfilterRuleDeleteParam.Index, "index", "", 0, "index of target rule")
}

var packetfilterInterfaceConnectCmd = &cobra.Command{
	Use: "interface-connect",

	Short: "InterfaceConnect Packetfilter",
	Long:  `InterfaceConnect Packetfilter`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := packetfilterInterfaceConnectParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("interface-connect parameter: \n%s\n", debugMarshalIndent(packetfilterInterfaceConnectParam))
		return err
	},
}

func packetfilterInterfaceConnectCmdInit() {
	fs := packetfilterInterfaceConnectCmd.Flags()
	fs.VarP(newIDValue(0, &packetfilterInterfaceConnectParam.InterfaceId), "interface-id", "", "set interface ID")
}

var packetfilterInterfaceDisconnectCmd = &cobra.Command{
	Use: "interface-disconnect",

	Short: "InterfaceDisconnect Packetfilter",
	Long:  `InterfaceDisconnect Packetfilter`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := packetfilterInterfaceDisconnectParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("interface-disconnect parameter: \n%s\n", debugMarshalIndent(packetfilterInterfaceDisconnectParam))
		return err
	},
}

func packetfilterInterfaceDisconnectCmdInit() {
	fs := packetfilterInterfaceDisconnectCmd.Flags()
	fs.VarP(newIDValue(0, &packetfilterInterfaceDisconnectParam.InterfaceId), "interface-id", "", "set interface ID")
}

func init() {
	parent := packetfilterCmd

	packetfilterListCmdInit()
	parent.AddCommand(packetfilterListCmd)

	packetfilterCreateCmdInit()
	parent.AddCommand(packetfilterCreateCmd)

	packetfilterReadCmdInit()
	parent.AddCommand(packetfilterReadCmd)

	packetfilterUpdateCmdInit()
	parent.AddCommand(packetfilterUpdateCmd)

	packetfilterDeleteCmdInit()
	parent.AddCommand(packetfilterDeleteCmd)

	packetfilterRuleInfoCmdInit()
	parent.AddCommand(packetfilterRuleInfoCmd)

	packetfilterRuleAddCmdInit()
	parent.AddCommand(packetfilterRuleAddCmd)

	packetfilterRuleUpdateCmdInit()
	parent.AddCommand(packetfilterRuleUpdateCmd)

	packetfilterRuleDeleteCmdInit()
	parent.AddCommand(packetfilterRuleDeleteCmd)

	packetfilterInterfaceConnectCmdInit()
	parent.AddCommand(packetfilterInterfaceConnectCmd)

	packetfilterInterfaceDisconnectCmdInit()
	parent.AddCommand(packetfilterInterfaceDisconnectCmd)

	rootCmd.AddCommand(parent)
}
