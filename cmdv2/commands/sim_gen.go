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
	simListParam          = params.NewListSimParam()
	simCreateParam        = params.NewCreateSimParam()
	simReadParam          = params.NewReadSimParam()
	simUpdateParam        = params.NewUpdateSimParam()
	simDeleteParam        = params.NewDeleteSimParam()
	simCarrierInfoParam   = params.NewCarrierInfoSimParam()
	simCarrierUpdateParam = params.NewCarrierUpdateSimParam()
	simActivateParam      = params.NewActivateSimParam()
	simDeactivateParam    = params.NewDeactivateSimParam()
	simImeiLockParam      = params.NewImeiLockSimParam()
	simIpAddParam         = params.NewIpAddSimParam()
	simImeiUnlockParam    = params.NewImeiUnlockSimParam()
	simIpDeleteParam      = params.NewIpDeleteSimParam()
	simLogsParam          = params.NewLogsSimParam()
	simMonitorParam       = params.NewMonitorSimParam()
)

// simCmd represents the command to manage SAKURA Cloud SIM
var simCmd = &cobra.Command{
	Use:   "sim",
	Short: "A manage commands of SIM",
	Long:  `A manage commands of SIM`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

var simListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "find", "selector"},
	Short:   "List Sim",
	Long:    `List Sim`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := simListParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("list parameter: \n%s\n", debugMarshalIndent(simListParam))
		return err
	},
}

func simListCmdInit() {
	fs := simListCmd.Flags()
	fs.StringSliceVarP(&simListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringSliceVarP(&simListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &simListParam.Id), "id", "", "set filter by id(s)")
	fs.StringSliceVarP(&simListParam.Tags, "tags", "", []string{}, "set filter by tags(AND)")
	fs.IntVarP(&simListParam.From, "from", "", 0, "set offset")
	fs.IntVarP(&simListParam.Max, "max", "", 0, "set limit")
}

var simCreateCmd = &cobra.Command{
	Use: "create",

	Short: "Create Sim",
	Long:  `Create Sim`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := simCreateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("create parameter: \n%s\n", debugMarshalIndent(simCreateParam))
		return err
	},
}

func simCreateCmdInit() {
	fs := simCreateCmd.Flags()
	fs.BoolVarP(&simCreateParam.Disabled, "disabled", "", false, "")
	fs.StringSliceVarP(&simCreateParam.Carrier, "carrier", "", []string{}, "")
	fs.StringVarP(&simCreateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&simCreateParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&simCreateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.StringVarP(&simCreateParam.Iccid, "iccid", "", "", "")
	fs.StringVarP(&simCreateParam.Passcode, "passcode", "", "", "")
	fs.StringVarP(&simCreateParam.Imei, "imei", "", "", "")
	fs.VarP(newIDValue(0, &simCreateParam.IconId), "icon-id", "", "set Icon ID")
}

var simReadCmd = &cobra.Command{
	Use: "read",

	Short: "Read Sim",
	Long:  `Read Sim`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := simReadParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("read parameter: \n%s\n", debugMarshalIndent(simReadParam))
		return err
	},
}

func simReadCmdInit() {
}

var simUpdateCmd = &cobra.Command{
	Use: "update",

	Short: "Update Sim",
	Long:  `Update Sim`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := simUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("update parameter: \n%s\n", debugMarshalIndent(simUpdateParam))
		return err
	},
}

func simUpdateCmdInit() {
	fs := simUpdateCmd.Flags()
	fs.StringVarP(&simUpdateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&simUpdateParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&simUpdateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &simUpdateParam.IconId), "icon-id", "", "set Icon ID")
}

var simDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"rm"},
	Short:   "Delete Sim",
	Long:    `Delete Sim`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := simDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("delete parameter: \n%s\n", debugMarshalIndent(simDeleteParam))
		return err
	},
}

func simDeleteCmdInit() {
	fs := simDeleteCmd.Flags()
	fs.BoolVarP(&simDeleteParam.Force, "force", "f", false, "forced-delete flag if SIM is still activating")
}

var simCarrierInfoCmd = &cobra.Command{
	Use:     "carrier-info",
	Aliases: []string{"carrier-list"},
	Short:   "CarrierInfo Sim",
	Long:    `CarrierInfo Sim`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := simCarrierInfoParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("carrier-info parameter: \n%s\n", debugMarshalIndent(simCarrierInfoParam))
		return err
	},
}

func simCarrierInfoCmdInit() {
}

var simCarrierUpdateCmd = &cobra.Command{
	Use: "carrier-update",

	Short: "CarrierUpdate Sim",
	Long:  `CarrierUpdate Sim`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := simCarrierUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("carrier-update parameter: \n%s\n", debugMarshalIndent(simCarrierUpdateParam))
		return err
	},
}

func simCarrierUpdateCmdInit() {
	fs := simCarrierUpdateCmd.Flags()
	fs.StringSliceVarP(&simCarrierUpdateParam.Carrier, "carrier", "", []string{}, "")
}

var simActivateCmd = &cobra.Command{
	Use: "activate",

	Short: "Activate Sim",
	Long:  `Activate Sim`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := simActivateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("activate parameter: \n%s\n", debugMarshalIndent(simActivateParam))
		return err
	},
}

func simActivateCmdInit() {
}

var simDeactivateCmd = &cobra.Command{
	Use: "deactivate",

	Short: "Deactivate Sim",
	Long:  `Deactivate Sim`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := simDeactivateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("deactivate parameter: \n%s\n", debugMarshalIndent(simDeactivateParam))
		return err
	},
}

func simDeactivateCmdInit() {
}

var simImeiLockCmd = &cobra.Command{
	Use: "imei-lock",

	Short: "ImeiLock Sim",
	Long:  `ImeiLock Sim`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := simImeiLockParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("imei-lock parameter: \n%s\n", debugMarshalIndent(simImeiLockParam))
		return err
	},
}

func simImeiLockCmdInit() {
	fs := simImeiLockCmd.Flags()
	fs.StringVarP(&simImeiLockParam.Imei, "imei", "", "", "")
}

var simIpAddCmd = &cobra.Command{
	Use: "ip-add",

	Short: "IpAdd Sim",
	Long:  `IpAdd Sim`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := simIpAddParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("ip-add parameter: \n%s\n", debugMarshalIndent(simIpAddParam))
		return err
	},
}

func simIpAddCmdInit() {
	fs := simIpAddCmd.Flags()
	fs.StringVarP(&simIpAddParam.Ip, "ip", "", "", "")
}

var simImeiUnlockCmd = &cobra.Command{
	Use: "imei-unlock",

	Short: "ImeiUnlock Sim",
	Long:  `ImeiUnlock Sim`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := simImeiUnlockParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("imei-unlock parameter: \n%s\n", debugMarshalIndent(simImeiUnlockParam))
		return err
	},
}

func simImeiUnlockCmdInit() {
}

var simIpDeleteCmd = &cobra.Command{
	Use:     "ip-delete",
	Aliases: []string{"ip-del"},
	Short:   "IpDelete Sim",
	Long:    `IpDelete Sim`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := simIpDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("ip-delete parameter: \n%s\n", debugMarshalIndent(simIpDeleteParam))
		return err
	},
}

func simIpDeleteCmdInit() {
}

var simLogsCmd = &cobra.Command{
	Use: "logs",

	Short: "Logs Sim",
	Long:  `Logs Sim`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := simLogsParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("logs parameter: \n%s\n", debugMarshalIndent(simLogsParam))
		return err
	},
}

func simLogsCmdInit() {
	fs := simLogsCmd.Flags()
	fs.BoolVarP(&simLogsParam.Follow, "follow", "f", false, "follow log output")
	fs.Int64VarP(&simLogsParam.RefreshInterval, "refresh-interval", "", 3, "log refresh interval second")
}

var simMonitorCmd = &cobra.Command{
	Use: "monitor",

	Short: "Monitor Sim",
	Long:  `Monitor Sim`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := simMonitorParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("monitor parameter: \n%s\n", debugMarshalIndent(simMonitorParam))
		return err
	},
}

func simMonitorCmdInit() {
	fs := simMonitorCmd.Flags()
	fs.StringVarP(&simMonitorParam.Start, "start", "", "", "set start-time")
	fs.StringVarP(&simMonitorParam.End, "end", "", "", "set end-time")
	fs.StringVarP(&simMonitorParam.KeyFormat, "key-format", "", "sakuracloud.sim.{{.ID}}", "set monitoring value key-format")
}

func init() {
	parent := simCmd

	simListCmdInit()
	parent.AddCommand(simListCmd)

	simCreateCmdInit()
	parent.AddCommand(simCreateCmd)

	simReadCmdInit()
	parent.AddCommand(simReadCmd)

	simUpdateCmdInit()
	parent.AddCommand(simUpdateCmd)

	simDeleteCmdInit()
	parent.AddCommand(simDeleteCmd)

	simCarrierInfoCmdInit()
	parent.AddCommand(simCarrierInfoCmd)

	simCarrierUpdateCmdInit()
	parent.AddCommand(simCarrierUpdateCmd)

	simActivateCmdInit()
	parent.AddCommand(simActivateCmd)

	simDeactivateCmdInit()
	parent.AddCommand(simDeactivateCmd)

	simImeiLockCmdInit()
	parent.AddCommand(simImeiLockCmd)

	simIpAddCmdInit()
	parent.AddCommand(simIpAddCmd)

	simImeiUnlockCmdInit()
	parent.AddCommand(simImeiUnlockCmd)

	simIpDeleteCmdInit()
	parent.AddCommand(simIpDeleteCmd)

	simLogsCmdInit()
	parent.AddCommand(simLogsCmd)

	simMonitorCmdInit()
	parent.AddCommand(simMonitorCmd)

	rootCmd.AddCommand(parent)
}
