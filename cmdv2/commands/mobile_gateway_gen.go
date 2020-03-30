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
	mobilegatewayListParam                  = params.NewListMobilegatewayParam()
	mobilegatewayCreateParam                = params.NewCreateMobilegatewayParam()
	mobilegatewayReadParam                  = params.NewReadMobilegatewayParam()
	mobilegatewayUpdateParam                = params.NewUpdateMobilegatewayParam()
	mobilegatewayDeleteParam                = params.NewDeleteMobilegatewayParam()
	mobilegatewayBootParam                  = params.NewBootMobilegatewayParam()
	mobilegatewayShutdownParam              = params.NewShutdownMobilegatewayParam()
	mobilegatewayShutdownForceParam         = params.NewShutdownForceMobilegatewayParam()
	mobilegatewayResetParam                 = params.NewResetMobilegatewayParam()
	mobilegatewayWaitForBootParam           = params.NewWaitForBootMobilegatewayParam()
	mobilegatewayWaitForDownParam           = params.NewWaitForDownMobilegatewayParam()
	mobilegatewayInterfaceInfoParam         = params.NewInterfaceInfoMobilegatewayParam()
	mobilegatewayInterfaceConnectParam      = params.NewInterfaceConnectMobilegatewayParam()
	mobilegatewayInterfaceUpdateParam       = params.NewInterfaceUpdateMobilegatewayParam()
	mobilegatewayInterfaceDisconnectParam   = params.NewInterfaceDisconnectMobilegatewayParam()
	mobilegatewayTrafficControlInfoParam    = params.NewTrafficControlInfoMobilegatewayParam()
	mobilegatewayTrafficControlEnableParam  = params.NewTrafficControlEnableMobilegatewayParam()
	mobilegatewayTrafficControlUpdateParam  = params.NewTrafficControlUpdateMobilegatewayParam()
	mobilegatewayTrafficControlDisableParam = params.NewTrafficControlDisableMobilegatewayParam()
	mobilegatewayStaticRouteInfoParam       = params.NewStaticRouteInfoMobilegatewayParam()
	mobilegatewayStaticRouteAddParam        = params.NewStaticRouteAddMobilegatewayParam()
	mobilegatewayStaticRouteUpdateParam     = params.NewStaticRouteUpdateMobilegatewayParam()
	mobilegatewayStaticRouteDeleteParam     = params.NewStaticRouteDeleteMobilegatewayParam()
	mobilegatewaySimInfoParam               = params.NewSimInfoMobilegatewayParam()
	mobilegatewaySimAddParam                = params.NewSimAddMobilegatewayParam()
	mobilegatewaySimUpdateParam             = params.NewSimUpdateMobilegatewayParam()
	mobilegatewaySimDeleteParam             = params.NewSimDeleteMobilegatewayParam()
	mobilegatewaySimRouteInfoParam          = params.NewSimRouteInfoMobilegatewayParam()
	mobilegatewaySimRouteAddParam           = params.NewSimRouteAddMobilegatewayParam()
	mobilegatewaySimRouteUpdateParam        = params.NewSimRouteUpdateMobilegatewayParam()
	mobilegatewaySimRouteDeleteParam        = params.NewSimRouteDeleteMobilegatewayParam()
	mobilegatewayDnsUpdateParam             = params.NewDnsUpdateMobilegatewayParam()
	mobilegatewayLogsParam                  = params.NewLogsMobilegatewayParam()
)

// mobilegatewayCmd represents the command to manage SAKURA Cloud MobileGateway
var mobilegatewayCmd = &cobra.Command{
	Use:   "mobilegateway",
	Short: "A manage commands of MobileGateway",
	Long:  `A manage commands of MobileGateway`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

var mobilegatewayListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "find", "selector"},
	Short:   "List Mobilegateway",
	Long:    `List Mobilegateway`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewayListParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("list parameter: \n%s\n", debugMarshalIndent(mobilegatewayListParam))
		return err
	},
}

func mobilegatewayListCmdInit() {
	fs := mobilegatewayListCmd.Flags()
	fs.StringSliceVarP(&mobilegatewayListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &mobilegatewayListParam.Id), "id", "", "set filter by id(s)")
	fs.IntVarP(&mobilegatewayListParam.From, "from", "", 0, "set offset")
	fs.IntVarP(&mobilegatewayListParam.Max, "max", "", 0, "set limit")
	fs.StringSliceVarP(&mobilegatewayListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringSliceVarP(&mobilegatewayListParam.Tags, "tags", "", []string{}, "set filter by tags(AND)")
}

var mobilegatewayCreateCmd = &cobra.Command{
	Use: "create",

	Short: "Create Mobilegateway",
	Long:  `Create Mobilegateway`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewayCreateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("create parameter: \n%s\n", debugMarshalIndent(mobilegatewayCreateParam))
		return err
	},
}

func mobilegatewayCreateCmdInit() {
	fs := mobilegatewayCreateCmd.Flags()
	fs.VarP(newIDValue(0, &mobilegatewayCreateParam.IconId), "icon-id", "", "set Icon ID")
	fs.BoolVarP(&mobilegatewayCreateParam.InternetConnection, "internet-connection", "", false, "connect to internet")
	fs.StringVarP(&mobilegatewayCreateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&mobilegatewayCreateParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&mobilegatewayCreateParam.Tags, "tags", "", []string{}, "set resource tags")
}

var mobilegatewayReadCmd = &cobra.Command{
	Use: "read",

	Short: "Read Mobilegateway",
	Long:  `Read Mobilegateway`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewayReadParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("read parameter: \n%s\n", debugMarshalIndent(mobilegatewayReadParam))
		return err
	},
}

func mobilegatewayReadCmdInit() {
}

var mobilegatewayUpdateCmd = &cobra.Command{
	Use: "update",

	Short: "Update Mobilegateway",
	Long:  `Update Mobilegateway`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewayUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("update parameter: \n%s\n", debugMarshalIndent(mobilegatewayUpdateParam))
		return err
	},
}

func mobilegatewayUpdateCmdInit() {
	fs := mobilegatewayUpdateCmd.Flags()
	fs.StringVarP(&mobilegatewayUpdateParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&mobilegatewayUpdateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &mobilegatewayUpdateParam.IconId), "icon-id", "", "set Icon ID")
	fs.BoolVarP(&mobilegatewayUpdateParam.InternetConnection, "internet-connection", "", false, "connect to internet")
	fs.StringVarP(&mobilegatewayUpdateParam.Name, "name", "", "", "set resource display name")
}

var mobilegatewayDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"rm"},
	Short:   "Delete Mobilegateway",
	Long:    `Delete Mobilegateway`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewayDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("delete parameter: \n%s\n", debugMarshalIndent(mobilegatewayDeleteParam))
		return err
	},
}

func mobilegatewayDeleteCmdInit() {
	fs := mobilegatewayDeleteCmd.Flags()
	fs.BoolVarP(&mobilegatewayDeleteParam.Force, "force", "f", false, "forced-shutdown flag if mobile-gateway is running")
}

var mobilegatewayBootCmd = &cobra.Command{
	Use:     "boot",
	Aliases: []string{"power-on"},
	Short:   "Boot Mobilegateway",
	Long:    `Boot Mobilegateway`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewayBootParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("boot parameter: \n%s\n", debugMarshalIndent(mobilegatewayBootParam))
		return err
	},
}

func mobilegatewayBootCmdInit() {
}

var mobilegatewayShutdownCmd = &cobra.Command{
	Use:     "shutdown",
	Aliases: []string{"power-off"},
	Short:   "Shutdown Mobilegateway",
	Long:    `Shutdown Mobilegateway`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewayShutdownParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("shutdown parameter: \n%s\n", debugMarshalIndent(mobilegatewayShutdownParam))
		return err
	},
}

func mobilegatewayShutdownCmdInit() {
}

var mobilegatewayShutdownForceCmd = &cobra.Command{
	Use:     "shutdown-force",
	Aliases: []string{"stop"},
	Short:   "ShutdownForce Mobilegateway",
	Long:    `ShutdownForce Mobilegateway`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewayShutdownForceParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("shutdown-force parameter: \n%s\n", debugMarshalIndent(mobilegatewayShutdownForceParam))
		return err
	},
}

func mobilegatewayShutdownForceCmdInit() {
}

var mobilegatewayResetCmd = &cobra.Command{
	Use: "reset",

	Short: "Reset Mobilegateway",
	Long:  `Reset Mobilegateway`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewayResetParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("reset parameter: \n%s\n", debugMarshalIndent(mobilegatewayResetParam))
		return err
	},
}

func mobilegatewayResetCmdInit() {
}

var mobilegatewayWaitForBootCmd = &cobra.Command{
	Use: "wait-for-boot",

	Short: "Wait until boot is completed",
	Long:  `Wait until boot is completed`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewayWaitForBootParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("wait-for-boot parameter: \n%s\n", debugMarshalIndent(mobilegatewayWaitForBootParam))
		return err
	},
}

func mobilegatewayWaitForBootCmdInit() {
}

var mobilegatewayWaitForDownCmd = &cobra.Command{
	Use: "wait-for-down",

	Short: "Wait until shutdown is completed",
	Long:  `Wait until shutdown is completed`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewayWaitForDownParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("wait-for-down parameter: \n%s\n", debugMarshalIndent(mobilegatewayWaitForDownParam))
		return err
	},
}

func mobilegatewayWaitForDownCmdInit() {
}

var mobilegatewayInterfaceInfoCmd = &cobra.Command{
	Use:     "interface-info",
	Aliases: []string{"interface-list"},
	Short:   "Show information of NIC(s) connected to mobile-gateway",
	Long:    `Show information of NIC(s) connected to mobile-gateway`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewayInterfaceInfoParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("interface-info parameter: \n%s\n", debugMarshalIndent(mobilegatewayInterfaceInfoParam))
		return err
	},
}

func mobilegatewayInterfaceInfoCmdInit() {
}

var mobilegatewayInterfaceConnectCmd = &cobra.Command{
	Use: "interface-connect",

	Short: "Connected to switch",
	Long:  `Connected to switch`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewayInterfaceConnectParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("interface-connect parameter: \n%s\n", debugMarshalIndent(mobilegatewayInterfaceConnectParam))
		return err
	},
}

func mobilegatewayInterfaceConnectCmdInit() {
	fs := mobilegatewayInterfaceConnectCmd.Flags()
	fs.VarP(newIDValue(0, &mobilegatewayInterfaceConnectParam.SwitchId), "switch-id", "", "set connect switch ID")
	fs.StringVarP(&mobilegatewayInterfaceConnectParam.Ipaddress, "ipaddress", "", "", "set ipaddress")
	fs.IntVarP(&mobilegatewayInterfaceConnectParam.NwMasklen, "nw-masklen", "", 24, "set ipaddress prefix")
}

var mobilegatewayInterfaceUpdateCmd = &cobra.Command{
	Use: "interface-update",

	Short: "Update interface",
	Long:  `Update interface`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewayInterfaceUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("interface-update parameter: \n%s\n", debugMarshalIndent(mobilegatewayInterfaceUpdateParam))
		return err
	},
}

func mobilegatewayInterfaceUpdateCmdInit() {
	fs := mobilegatewayInterfaceUpdateCmd.Flags()
	fs.StringVarP(&mobilegatewayInterfaceUpdateParam.Ipaddress, "ipaddress", "", "", "set ipaddress")
	fs.IntVarP(&mobilegatewayInterfaceUpdateParam.NwMasklen, "nw-masklen", "", 24, "set ipaddress prefix")
}

var mobilegatewayInterfaceDisconnectCmd = &cobra.Command{
	Use: "interface-disconnect",

	Short: "Disconnected to switch",
	Long:  `Disconnected to switch`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewayInterfaceDisconnectParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("interface-disconnect parameter: \n%s\n", debugMarshalIndent(mobilegatewayInterfaceDisconnectParam))
		return err
	},
}

func mobilegatewayInterfaceDisconnectCmdInit() {
}

var mobilegatewayTrafficControlInfoCmd = &cobra.Command{
	Use: "traffic-control-info",

	Short: "Show information of traffic-control",
	Long:  `Show information of traffic-control`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewayTrafficControlInfoParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("traffic-control-info parameter: \n%s\n", debugMarshalIndent(mobilegatewayTrafficControlInfoParam))
		return err
	},
}

func mobilegatewayTrafficControlInfoCmdInit() {
}

var mobilegatewayTrafficControlEnableCmd = &cobra.Command{
	Use: "traffic-control-enable",

	Short: "Enable traffic-control",
	Long:  `Enable traffic-control`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewayTrafficControlEnableParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("traffic-control-enable parameter: \n%s\n", debugMarshalIndent(mobilegatewayTrafficControlEnableParam))
		return err
	},
}

func mobilegatewayTrafficControlEnableCmdInit() {
	fs := mobilegatewayTrafficControlEnableCmd.Flags()
	fs.StringVarP(&mobilegatewayTrafficControlEnableParam.SlackWebhookUrl, "slack-webhook-url", "", "", "")
	fs.BoolVarP(&mobilegatewayTrafficControlEnableParam.AutoTrafficShaping, "auto-traffic-shaping", "", false, "")
	fs.IntVarP(&mobilegatewayTrafficControlEnableParam.Quota, "quota", "", 512, "")
	fs.IntVarP(&mobilegatewayTrafficControlEnableParam.BandWidthLimit, "band-width-limit", "", 0, "")
	fs.BoolVarP(&mobilegatewayTrafficControlEnableParam.EnableEmail, "enable-email", "", false, "")
}

var mobilegatewayTrafficControlUpdateCmd = &cobra.Command{
	Use: "traffic-control-update",

	Short: "Update traffic-control config",
	Long:  `Update traffic-control config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewayTrafficControlUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("traffic-control-update parameter: \n%s\n", debugMarshalIndent(mobilegatewayTrafficControlUpdateParam))
		return err
	},
}

func mobilegatewayTrafficControlUpdateCmdInit() {
	fs := mobilegatewayTrafficControlUpdateCmd.Flags()
	fs.IntVarP(&mobilegatewayTrafficControlUpdateParam.BandWidthLimit, "band-width-limit", "", 0, "")
	fs.BoolVarP(&mobilegatewayTrafficControlUpdateParam.EnableEmail, "enable-email", "", false, "")
	fs.StringVarP(&mobilegatewayTrafficControlUpdateParam.SlackWebhookUrl, "slack-webhook-url", "", "", "")
	fs.BoolVarP(&mobilegatewayTrafficControlUpdateParam.AutoTrafficShaping, "auto-traffic-shaping", "", false, "")
	fs.IntVarP(&mobilegatewayTrafficControlUpdateParam.Quota, "quota", "", 0, "")
}

var mobilegatewayTrafficControlDisableCmd = &cobra.Command{
	Use:     "traffic-control-disable",
	Aliases: []string{"traffic-control-delete"},
	Short:   "Disable traffic-control config",
	Long:    `Disable traffic-control config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewayTrafficControlDisableParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("traffic-control-disable parameter: \n%s\n", debugMarshalIndent(mobilegatewayTrafficControlDisableParam))
		return err
	},
}

func mobilegatewayTrafficControlDisableCmdInit() {
}

var mobilegatewayStaticRouteInfoCmd = &cobra.Command{
	Use:     "static-route-info",
	Aliases: []string{"static-route-list"},
	Short:   "Show information of static-routes",
	Long:    `Show information of static-routes`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewayStaticRouteInfoParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("static-route-info parameter: \n%s\n", debugMarshalIndent(mobilegatewayStaticRouteInfoParam))
		return err
	},
}

func mobilegatewayStaticRouteInfoCmdInit() {
}

var mobilegatewayStaticRouteAddCmd = &cobra.Command{
	Use: "static-route-add",

	Short: "Add static-route",
	Long:  `Add static-route`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewayStaticRouteAddParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("static-route-add parameter: \n%s\n", debugMarshalIndent(mobilegatewayStaticRouteAddParam))
		return err
	},
}

func mobilegatewayStaticRouteAddCmdInit() {
	fs := mobilegatewayStaticRouteAddCmd.Flags()
	fs.StringVarP(&mobilegatewayStaticRouteAddParam.Prefix, "prefix", "", "", "set prefix")
	fs.StringVarP(&mobilegatewayStaticRouteAddParam.NextHop, "next-hop", "", "", "set next-hop")
}

var mobilegatewayStaticRouteUpdateCmd = &cobra.Command{
	Use: "static-route-update",

	Short: "Update static-route",
	Long:  `Update static-route`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewayStaticRouteUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("static-route-update parameter: \n%s\n", debugMarshalIndent(mobilegatewayStaticRouteUpdateParam))
		return err
	},
}

func mobilegatewayStaticRouteUpdateCmdInit() {
	fs := mobilegatewayStaticRouteUpdateCmd.Flags()
	fs.IntVarP(&mobilegatewayStaticRouteUpdateParam.Index, "index", "", 0, "index of target static-route")
	fs.StringVarP(&mobilegatewayStaticRouteUpdateParam.Prefix, "prefix", "", "", "set prefix")
	fs.StringVarP(&mobilegatewayStaticRouteUpdateParam.NextHop, "next-hop", "", "", "set next-hop")
}

var mobilegatewayStaticRouteDeleteCmd = &cobra.Command{
	Use: "static-route-delete",

	Short: "Delete static-route",
	Long:  `Delete static-route`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewayStaticRouteDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("static-route-delete parameter: \n%s\n", debugMarshalIndent(mobilegatewayStaticRouteDeleteParam))
		return err
	},
}

func mobilegatewayStaticRouteDeleteCmdInit() {
	fs := mobilegatewayStaticRouteDeleteCmd.Flags()
	fs.IntVarP(&mobilegatewayStaticRouteDeleteParam.Index, "index", "", 0, "index of target static-route")
}

var mobilegatewaySimInfoCmd = &cobra.Command{
	Use:     "sim-info",
	Aliases: []string{"interface-list"},
	Short:   "Show information of NIC(s) connected to mobile-gateway",
	Long:    `Show information of NIC(s) connected to mobile-gateway`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewaySimInfoParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("sim-info parameter: \n%s\n", debugMarshalIndent(mobilegatewaySimInfoParam))
		return err
	},
}

func mobilegatewaySimInfoCmdInit() {
}

var mobilegatewaySimAddCmd = &cobra.Command{
	Use: "sim-add",

	Short: "Connected to switch",
	Long:  `Connected to switch`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewaySimAddParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("sim-add parameter: \n%s\n", debugMarshalIndent(mobilegatewaySimAddParam))
		return err
	},
}

func mobilegatewaySimAddCmdInit() {
	fs := mobilegatewaySimAddCmd.Flags()
	fs.StringVarP(&mobilegatewaySimAddParam.Ipaddress, "ipaddress", "", "", "set ipaddress")
	fs.VarP(newIDValue(0, &mobilegatewaySimAddParam.SimId), "sim-id", "", "")
}

var mobilegatewaySimUpdateCmd = &cobra.Command{
	Use: "sim-update",

	Short: "Connected to switch",
	Long:  `Connected to switch`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewaySimUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("sim-update parameter: \n%s\n", debugMarshalIndent(mobilegatewaySimUpdateParam))
		return err
	},
}

func mobilegatewaySimUpdateCmdInit() {
	fs := mobilegatewaySimUpdateCmd.Flags()
	fs.VarP(newIDValue(0, &mobilegatewaySimUpdateParam.SimId), "sim-id", "", "")
	fs.StringVarP(&mobilegatewaySimUpdateParam.Ipaddress, "ipaddress", "", "", "set ipaddress")
}

var mobilegatewaySimDeleteCmd = &cobra.Command{
	Use: "sim-delete",

	Short: "Disconnected to switch",
	Long:  `Disconnected to switch`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewaySimDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("sim-delete parameter: \n%s\n", debugMarshalIndent(mobilegatewaySimDeleteParam))
		return err
	},
}

func mobilegatewaySimDeleteCmdInit() {
	fs := mobilegatewaySimDeleteCmd.Flags()
	fs.VarP(newIDValue(0, &mobilegatewaySimDeleteParam.SimId), "sim-id", "", "")
}

var mobilegatewaySimRouteInfoCmd = &cobra.Command{
	Use:     "sim-route-info",
	Aliases: []string{"sim-route-list"},
	Short:   "Show information of sim-routes",
	Long:    `Show information of sim-routes`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewaySimRouteInfoParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("sim-route-info parameter: \n%s\n", debugMarshalIndent(mobilegatewaySimRouteInfoParam))
		return err
	},
}

func mobilegatewaySimRouteInfoCmdInit() {
}

var mobilegatewaySimRouteAddCmd = &cobra.Command{
	Use: "sim-route-add",

	Short: "Add sim-route",
	Long:  `Add sim-route`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewaySimRouteAddParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("sim-route-add parameter: \n%s\n", debugMarshalIndent(mobilegatewaySimRouteAddParam))
		return err
	},
}

func mobilegatewaySimRouteAddCmdInit() {
	fs := mobilegatewaySimRouteAddCmd.Flags()
	fs.StringVarP(&mobilegatewaySimRouteAddParam.Prefix, "prefix", "", "", "set prefix")
	fs.VarP(newIDValue(0, &mobilegatewaySimRouteAddParam.Sim), "sim", "", "set sim")
}

var mobilegatewaySimRouteUpdateCmd = &cobra.Command{
	Use: "sim-route-update",

	Short: "Update sim-route",
	Long:  `Update sim-route`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewaySimRouteUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("sim-route-update parameter: \n%s\n", debugMarshalIndent(mobilegatewaySimRouteUpdateParam))
		return err
	},
}

func mobilegatewaySimRouteUpdateCmdInit() {
	fs := mobilegatewaySimRouteUpdateCmd.Flags()
	fs.IntVarP(&mobilegatewaySimRouteUpdateParam.Index, "index", "", 0, "index of target sim-route")
	fs.StringVarP(&mobilegatewaySimRouteUpdateParam.Prefix, "prefix", "", "", "set prefix")
	fs.VarP(newIDValue(0, &mobilegatewaySimRouteUpdateParam.Sim), "sim", "", "set sim")
}

var mobilegatewaySimRouteDeleteCmd = &cobra.Command{
	Use: "sim-route-delete",

	Short: "Delete sim-route",
	Long:  `Delete sim-route`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewaySimRouteDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("sim-route-delete parameter: \n%s\n", debugMarshalIndent(mobilegatewaySimRouteDeleteParam))
		return err
	},
}

func mobilegatewaySimRouteDeleteCmdInit() {
	fs := mobilegatewaySimRouteDeleteCmd.Flags()
	fs.IntVarP(&mobilegatewaySimRouteDeleteParam.Index, "index", "", 0, "index of target sim-route")
}

var mobilegatewayDnsUpdateCmd = &cobra.Command{
	Use: "dns-update",

	Short: "Update interface",
	Long:  `Update interface`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewayDnsUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("dns-update parameter: \n%s\n", debugMarshalIndent(mobilegatewayDnsUpdateParam))
		return err
	},
}

func mobilegatewayDnsUpdateCmdInit() {
	fs := mobilegatewayDnsUpdateCmd.Flags()
	fs.StringVarP(&mobilegatewayDnsUpdateParam.Dns1, "dns-1", "", "", "set DNS server address")
	fs.StringVarP(&mobilegatewayDnsUpdateParam.Dns2, "dns-2", "", "", "set DNS server address")
}

var mobilegatewayLogsCmd = &cobra.Command{
	Use: "logs",

	Short: "Logs Mobilegateway",
	Long:  `Logs Mobilegateway`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mobilegatewayLogsParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("logs parameter: \n%s\n", debugMarshalIndent(mobilegatewayLogsParam))
		return err
	},
}

func mobilegatewayLogsCmdInit() {
	fs := mobilegatewayLogsCmd.Flags()
	fs.Int64VarP(&mobilegatewayLogsParam.RefreshInterval, "refresh-interval", "", 3, "log refresh interval second")
	fs.BoolVarP(&mobilegatewayLogsParam.Follow, "follow", "f", false, "follow log output")
}

func init() {
	parent := mobilegatewayCmd

	mobilegatewayListCmdInit()
	parent.AddCommand(mobilegatewayListCmd)

	mobilegatewayCreateCmdInit()
	parent.AddCommand(mobilegatewayCreateCmd)

	mobilegatewayReadCmdInit()
	parent.AddCommand(mobilegatewayReadCmd)

	mobilegatewayUpdateCmdInit()
	parent.AddCommand(mobilegatewayUpdateCmd)

	mobilegatewayDeleteCmdInit()
	parent.AddCommand(mobilegatewayDeleteCmd)

	mobilegatewayBootCmdInit()
	parent.AddCommand(mobilegatewayBootCmd)

	mobilegatewayShutdownCmdInit()
	parent.AddCommand(mobilegatewayShutdownCmd)

	mobilegatewayShutdownForceCmdInit()
	parent.AddCommand(mobilegatewayShutdownForceCmd)

	mobilegatewayResetCmdInit()
	parent.AddCommand(mobilegatewayResetCmd)

	mobilegatewayWaitForBootCmdInit()
	parent.AddCommand(mobilegatewayWaitForBootCmd)

	mobilegatewayWaitForDownCmdInit()
	parent.AddCommand(mobilegatewayWaitForDownCmd)

	mobilegatewayInterfaceInfoCmdInit()
	parent.AddCommand(mobilegatewayInterfaceInfoCmd)

	mobilegatewayInterfaceConnectCmdInit()
	parent.AddCommand(mobilegatewayInterfaceConnectCmd)

	mobilegatewayInterfaceUpdateCmdInit()
	parent.AddCommand(mobilegatewayInterfaceUpdateCmd)

	mobilegatewayInterfaceDisconnectCmdInit()
	parent.AddCommand(mobilegatewayInterfaceDisconnectCmd)

	mobilegatewayTrafficControlInfoCmdInit()
	parent.AddCommand(mobilegatewayTrafficControlInfoCmd)

	mobilegatewayTrafficControlEnableCmdInit()
	parent.AddCommand(mobilegatewayTrafficControlEnableCmd)

	mobilegatewayTrafficControlUpdateCmdInit()
	parent.AddCommand(mobilegatewayTrafficControlUpdateCmd)

	mobilegatewayTrafficControlDisableCmdInit()
	parent.AddCommand(mobilegatewayTrafficControlDisableCmd)

	mobilegatewayStaticRouteInfoCmdInit()
	parent.AddCommand(mobilegatewayStaticRouteInfoCmd)

	mobilegatewayStaticRouteAddCmdInit()
	parent.AddCommand(mobilegatewayStaticRouteAddCmd)

	mobilegatewayStaticRouteUpdateCmdInit()
	parent.AddCommand(mobilegatewayStaticRouteUpdateCmd)

	mobilegatewayStaticRouteDeleteCmdInit()
	parent.AddCommand(mobilegatewayStaticRouteDeleteCmd)

	mobilegatewaySimInfoCmdInit()
	parent.AddCommand(mobilegatewaySimInfoCmd)

	mobilegatewaySimAddCmdInit()
	parent.AddCommand(mobilegatewaySimAddCmd)

	mobilegatewaySimUpdateCmdInit()
	parent.AddCommand(mobilegatewaySimUpdateCmd)

	mobilegatewaySimDeleteCmdInit()
	parent.AddCommand(mobilegatewaySimDeleteCmd)

	mobilegatewaySimRouteInfoCmdInit()
	parent.AddCommand(mobilegatewaySimRouteInfoCmd)

	mobilegatewaySimRouteAddCmdInit()
	parent.AddCommand(mobilegatewaySimRouteAddCmd)

	mobilegatewaySimRouteUpdateCmdInit()
	parent.AddCommand(mobilegatewaySimRouteUpdateCmd)

	mobilegatewaySimRouteDeleteCmdInit()
	parent.AddCommand(mobilegatewaySimRouteDeleteCmd)

	mobilegatewayDnsUpdateCmdInit()
	parent.AddCommand(mobilegatewayDnsUpdateCmd)

	mobilegatewayLogsCmdInit()
	parent.AddCommand(mobilegatewayLogsCmd)

	rootCmd.AddCommand(parent)
}
