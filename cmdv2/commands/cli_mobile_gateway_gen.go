// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-v2-commands'; DO NOT EDIT

package commands

import (
	"fmt"

	"github.com/sacloud/usacloud/cmdv2/params"
	"github.com/spf13/cobra"
)

// mobilegatewayCmd represents the mobilegateway command
var mobilegatewayCmd = &cobra.Command{
	Use:   "mobile-gateway",
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
		listParam, err := params.NewListMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("list parameter: \n%s\n", debugMarshalIndent(listParam))
		return err
	},
}

var mobilegatewayCreateCmd = &cobra.Command{
	Use: "create",

	Short: "Create Mobilegateway",
	Long:  `Create Mobilegateway`,
	RunE: func(cmd *cobra.Command, args []string) error {
		createParam, err := params.NewCreateMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("create parameter: \n%s\n", debugMarshalIndent(createParam))
		return err
	},
}

var mobilegatewayReadCmd = &cobra.Command{
	Use: "read",

	Short: "Read Mobilegateway",
	Long:  `Read Mobilegateway`,
	RunE: func(cmd *cobra.Command, args []string) error {
		readParam, err := params.NewReadMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("read parameter: \n%s\n", debugMarshalIndent(readParam))
		return err
	},
}

var mobilegatewayUpdateCmd = &cobra.Command{
	Use: "update",

	Short: "Update Mobilegateway",
	Long:  `Update Mobilegateway`,
	RunE: func(cmd *cobra.Command, args []string) error {
		updateParam, err := params.NewUpdateMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("update parameter: \n%s\n", debugMarshalIndent(updateParam))
		return err
	},
}

var mobilegatewayDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"rm"},
	Short:   "Delete Mobilegateway",
	Long:    `Delete Mobilegateway`,
	RunE: func(cmd *cobra.Command, args []string) error {
		deleteParam, err := params.NewDeleteMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("delete parameter: \n%s\n", debugMarshalIndent(deleteParam))
		return err
	},
}

var mobilegatewayBootCmd = &cobra.Command{
	Use:     "boot",
	Aliases: []string{"power-on"},
	Short:   "Boot Mobilegateway",
	Long:    `Boot Mobilegateway`,
	RunE: func(cmd *cobra.Command, args []string) error {
		bootParam, err := params.NewBootMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("boot parameter: \n%s\n", debugMarshalIndent(bootParam))
		return err
	},
}

var mobilegatewayShutdownCmd = &cobra.Command{
	Use:     "shutdown",
	Aliases: []string{"power-off"},
	Short:   "Shutdown Mobilegateway",
	Long:    `Shutdown Mobilegateway`,
	RunE: func(cmd *cobra.Command, args []string) error {
		shutdownParam, err := params.NewShutdownMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("shutdown parameter: \n%s\n", debugMarshalIndent(shutdownParam))
		return err
	},
}

var mobilegatewayShutdownForceCmd = &cobra.Command{
	Use:     "shutdown-force",
	Aliases: []string{"stop"},
	Short:   "ShutdownForce Mobilegateway",
	Long:    `ShutdownForce Mobilegateway`,
	RunE: func(cmd *cobra.Command, args []string) error {
		shutdownForceParam, err := params.NewShutdownForceMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("shutdown-force parameter: \n%s\n", debugMarshalIndent(shutdownForceParam))
		return err
	},
}

var mobilegatewayResetCmd = &cobra.Command{
	Use: "reset",

	Short: "Reset Mobilegateway",
	Long:  `Reset Mobilegateway`,
	RunE: func(cmd *cobra.Command, args []string) error {
		resetParam, err := params.NewResetMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("reset parameter: \n%s\n", debugMarshalIndent(resetParam))
		return err
	},
}

var mobilegatewayWaitForBootCmd = &cobra.Command{
	Use: "wait-for-boot",

	Short: "Wait until boot is completed",
	Long:  `Wait until boot is completed`,
	RunE: func(cmd *cobra.Command, args []string) error {
		waitForBootParam, err := params.NewWaitForBootMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("wait-for-boot parameter: \n%s\n", debugMarshalIndent(waitForBootParam))
		return err
	},
}

var mobilegatewayWaitForDownCmd = &cobra.Command{
	Use: "wait-for-down",

	Short: "Wait until shutdown is completed",
	Long:  `Wait until shutdown is completed`,
	RunE: func(cmd *cobra.Command, args []string) error {
		waitForDownParam, err := params.NewWaitForDownMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("wait-for-down parameter: \n%s\n", debugMarshalIndent(waitForDownParam))
		return err
	},
}

var mobilegatewayInterfaceInfoCmd = &cobra.Command{
	Use:     "interface-info",
	Aliases: []string{"interface-list"},
	Short:   "Show information of NIC(s) connected to mobile-gateway",
	Long:    `Show information of NIC(s) connected to mobile-gateway`,
	RunE: func(cmd *cobra.Command, args []string) error {
		interfaceInfoParam, err := params.NewInterfaceInfoMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("interface-info parameter: \n%s\n", debugMarshalIndent(interfaceInfoParam))
		return err
	},
}

var mobilegatewayInterfaceConnectCmd = &cobra.Command{
	Use: "interface-connect",

	Short: "Connected to switch",
	Long:  `Connected to switch`,
	RunE: func(cmd *cobra.Command, args []string) error {
		interfaceConnectParam, err := params.NewInterfaceConnectMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("interface-connect parameter: \n%s\n", debugMarshalIndent(interfaceConnectParam))
		return err
	},
}

var mobilegatewayInterfaceUpdateCmd = &cobra.Command{
	Use: "interface-update",

	Short: "Update interface",
	Long:  `Update interface`,
	RunE: func(cmd *cobra.Command, args []string) error {
		interfaceUpdateParam, err := params.NewInterfaceUpdateMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("interface-update parameter: \n%s\n", debugMarshalIndent(interfaceUpdateParam))
		return err
	},
}

var mobilegatewayInterfaceDisconnectCmd = &cobra.Command{
	Use: "interface-disconnect",

	Short: "Disconnected to switch",
	Long:  `Disconnected to switch`,
	RunE: func(cmd *cobra.Command, args []string) error {
		interfaceDisconnectParam, err := params.NewInterfaceDisconnectMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("interface-disconnect parameter: \n%s\n", debugMarshalIndent(interfaceDisconnectParam))
		return err
	},
}

var mobilegatewayTrafficControlInfoCmd = &cobra.Command{
	Use: "traffic-control-info",

	Short: "Show information of traffic-control",
	Long:  `Show information of traffic-control`,
	RunE: func(cmd *cobra.Command, args []string) error {
		trafficControlInfoParam, err := params.NewTrafficControlInfoMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("traffic-control-info parameter: \n%s\n", debugMarshalIndent(trafficControlInfoParam))
		return err
	},
}

var mobilegatewayTrafficControlEnableCmd = &cobra.Command{
	Use: "traffic-control-enable",

	Short: "Enable traffic-control",
	Long:  `Enable traffic-control`,
	RunE: func(cmd *cobra.Command, args []string) error {
		trafficControlEnableParam, err := params.NewTrafficControlEnableMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("traffic-control-enable parameter: \n%s\n", debugMarshalIndent(trafficControlEnableParam))
		return err
	},
}

var mobilegatewayTrafficControlUpdateCmd = &cobra.Command{
	Use: "traffic-control-update",

	Short: "Update traffic-control config",
	Long:  `Update traffic-control config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		trafficControlUpdateParam, err := params.NewTrafficControlUpdateMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("traffic-control-update parameter: \n%s\n", debugMarshalIndent(trafficControlUpdateParam))
		return err
	},
}

var mobilegatewayTrafficControlDisableCmd = &cobra.Command{
	Use:     "traffic-control-disable",
	Aliases: []string{"traffic-control-delete"},
	Short:   "Disable traffic-control config",
	Long:    `Disable traffic-control config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		trafficControlDisableParam, err := params.NewTrafficControlDisableMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("traffic-control-disable parameter: \n%s\n", debugMarshalIndent(trafficControlDisableParam))
		return err
	},
}

var mobilegatewayStaticRouteInfoCmd = &cobra.Command{
	Use:     "static-route-info",
	Aliases: []string{"static-route-list"},
	Short:   "Show information of static-routes",
	Long:    `Show information of static-routes`,
	RunE: func(cmd *cobra.Command, args []string) error {
		staticRouteInfoParam, err := params.NewStaticRouteInfoMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("static-route-info parameter: \n%s\n", debugMarshalIndent(staticRouteInfoParam))
		return err
	},
}

var mobilegatewayStaticRouteAddCmd = &cobra.Command{
	Use: "static-route-add",

	Short: "Add static-route",
	Long:  `Add static-route`,
	RunE: func(cmd *cobra.Command, args []string) error {
		staticRouteAddParam, err := params.NewStaticRouteAddMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("static-route-add parameter: \n%s\n", debugMarshalIndent(staticRouteAddParam))
		return err
	},
}

var mobilegatewayStaticRouteUpdateCmd = &cobra.Command{
	Use: "static-route-update",

	Short: "Update static-route",
	Long:  `Update static-route`,
	RunE: func(cmd *cobra.Command, args []string) error {
		staticRouteUpdateParam, err := params.NewStaticRouteUpdateMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("static-route-update parameter: \n%s\n", debugMarshalIndent(staticRouteUpdateParam))
		return err
	},
}

var mobilegatewayStaticRouteDeleteCmd = &cobra.Command{
	Use: "static-route-delete",

	Short: "Delete static-route",
	Long:  `Delete static-route`,
	RunE: func(cmd *cobra.Command, args []string) error {
		staticRouteDeleteParam, err := params.NewStaticRouteDeleteMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("static-route-delete parameter: \n%s\n", debugMarshalIndent(staticRouteDeleteParam))
		return err
	},
}

var mobilegatewaySimInfoCmd = &cobra.Command{
	Use:     "sim-info",
	Aliases: []string{"interface-list"},
	Short:   "Show information of NIC(s) connected to mobile-gateway",
	Long:    `Show information of NIC(s) connected to mobile-gateway`,
	RunE: func(cmd *cobra.Command, args []string) error {
		simInfoParam, err := params.NewSimInfoMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("sim-info parameter: \n%s\n", debugMarshalIndent(simInfoParam))
		return err
	},
}

var mobilegatewaySimAddCmd = &cobra.Command{
	Use: "sim-add",

	Short: "Connected to switch",
	Long:  `Connected to switch`,
	RunE: func(cmd *cobra.Command, args []string) error {
		simAddParam, err := params.NewSimAddMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("sim-add parameter: \n%s\n", debugMarshalIndent(simAddParam))
		return err
	},
}

var mobilegatewaySimUpdateCmd = &cobra.Command{
	Use: "sim-update",

	Short: "Connected to switch",
	Long:  `Connected to switch`,
	RunE: func(cmd *cobra.Command, args []string) error {
		simUpdateParam, err := params.NewSimUpdateMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("sim-update parameter: \n%s\n", debugMarshalIndent(simUpdateParam))
		return err
	},
}

var mobilegatewaySimDeleteCmd = &cobra.Command{
	Use: "sim-delete",

	Short: "Disconnected to switch",
	Long:  `Disconnected to switch`,
	RunE: func(cmd *cobra.Command, args []string) error {
		simDeleteParam, err := params.NewSimDeleteMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("sim-delete parameter: \n%s\n", debugMarshalIndent(simDeleteParam))
		return err
	},
}

var mobilegatewaySimRouteInfoCmd = &cobra.Command{
	Use:     "sim-route-info",
	Aliases: []string{"sim-route-list"},
	Short:   "Show information of sim-routes",
	Long:    `Show information of sim-routes`,
	RunE: func(cmd *cobra.Command, args []string) error {
		simRouteInfoParam, err := params.NewSimRouteInfoMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("sim-route-info parameter: \n%s\n", debugMarshalIndent(simRouteInfoParam))
		return err
	},
}

var mobilegatewaySimRouteAddCmd = &cobra.Command{
	Use: "sim-route-add",

	Short: "Add sim-route",
	Long:  `Add sim-route`,
	RunE: func(cmd *cobra.Command, args []string) error {
		simRouteAddParam, err := params.NewSimRouteAddMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("sim-route-add parameter: \n%s\n", debugMarshalIndent(simRouteAddParam))
		return err
	},
}

var mobilegatewaySimRouteUpdateCmd = &cobra.Command{
	Use: "sim-route-update",

	Short: "Update sim-route",
	Long:  `Update sim-route`,
	RunE: func(cmd *cobra.Command, args []string) error {
		simRouteUpdateParam, err := params.NewSimRouteUpdateMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("sim-route-update parameter: \n%s\n", debugMarshalIndent(simRouteUpdateParam))
		return err
	},
}

var mobilegatewaySimRouteDeleteCmd = &cobra.Command{
	Use: "sim-route-delete",

	Short: "Delete sim-route",
	Long:  `Delete sim-route`,
	RunE: func(cmd *cobra.Command, args []string) error {
		simRouteDeleteParam, err := params.NewSimRouteDeleteMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("sim-route-delete parameter: \n%s\n", debugMarshalIndent(simRouteDeleteParam))
		return err
	},
}

var mobilegatewayDnsUpdateCmd = &cobra.Command{
	Use: "dns-update",

	Short: "Update interface",
	Long:  `Update interface`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dnsUpdateParam, err := params.NewDnsUpdateMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("dns-update parameter: \n%s\n", debugMarshalIndent(dnsUpdateParam))
		return err
	},
}

var mobilegatewayLogsCmd = &cobra.Command{
	Use: "logs",

	Short: "Logs Mobilegateway",
	Long:  `Logs Mobilegateway`,
	RunE: func(cmd *cobra.Command, args []string) error {
		logsParam, err := params.NewLogsMobilegatewayParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("logs parameter: \n%s\n", debugMarshalIndent(logsParam))
		return err
	},
}

func init() {
	parent := mobilegatewayCmd
	parent.AddCommand(mobilegatewayListCmd)
	parent.AddCommand(mobilegatewayCreateCmd)
	parent.AddCommand(mobilegatewayReadCmd)
	parent.AddCommand(mobilegatewayUpdateCmd)
	parent.AddCommand(mobilegatewayDeleteCmd)
	parent.AddCommand(mobilegatewayBootCmd)
	parent.AddCommand(mobilegatewayShutdownCmd)
	parent.AddCommand(mobilegatewayShutdownForceCmd)
	parent.AddCommand(mobilegatewayResetCmd)
	parent.AddCommand(mobilegatewayWaitForBootCmd)
	parent.AddCommand(mobilegatewayWaitForDownCmd)
	parent.AddCommand(mobilegatewayInterfaceInfoCmd)
	parent.AddCommand(mobilegatewayInterfaceConnectCmd)
	parent.AddCommand(mobilegatewayInterfaceUpdateCmd)
	parent.AddCommand(mobilegatewayInterfaceDisconnectCmd)
	parent.AddCommand(mobilegatewayTrafficControlInfoCmd)
	parent.AddCommand(mobilegatewayTrafficControlEnableCmd)
	parent.AddCommand(mobilegatewayTrafficControlUpdateCmd)
	parent.AddCommand(mobilegatewayTrafficControlDisableCmd)
	parent.AddCommand(mobilegatewayStaticRouteInfoCmd)
	parent.AddCommand(mobilegatewayStaticRouteAddCmd)
	parent.AddCommand(mobilegatewayStaticRouteUpdateCmd)
	parent.AddCommand(mobilegatewayStaticRouteDeleteCmd)
	parent.AddCommand(mobilegatewaySimInfoCmd)
	parent.AddCommand(mobilegatewaySimAddCmd)
	parent.AddCommand(mobilegatewaySimUpdateCmd)
	parent.AddCommand(mobilegatewaySimDeleteCmd)
	parent.AddCommand(mobilegatewaySimRouteInfoCmd)
	parent.AddCommand(mobilegatewaySimRouteAddCmd)
	parent.AddCommand(mobilegatewaySimRouteUpdateCmd)
	parent.AddCommand(mobilegatewaySimRouteDeleteCmd)
	parent.AddCommand(mobilegatewayDnsUpdateCmd)
	parent.AddCommand(mobilegatewayLogsCmd)
	rootCmd.AddCommand(mobilegatewayCmd)
}
