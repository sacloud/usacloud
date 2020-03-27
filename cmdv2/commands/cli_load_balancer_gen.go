// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-v2-commands'; DO NOT EDIT

package commands

import (
	"fmt"

	"github.com/sacloud/usacloud/cmdv2/params"
	"github.com/spf13/cobra"
)

// loadbalancerCmd represents the loadbalancer command
var loadbalancerCmd = &cobra.Command{
	Use:   "load-balancer",
	Short: "A manage commands of LoadBalancer",
	Long:  `A manage commands of LoadBalancer`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

var loadbalancerListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "find", "selector"},
	Short:   "List Loadbalancer",
	Long:    `List Loadbalancer`,
	RunE: func(cmd *cobra.Command, args []string) error {
		listParam, err := params.NewListLoadbalancerParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("list parameter: \n%s\n", debugMarshalIndent(listParam))
		return err
	},
}

var loadbalancerCreateCmd = &cobra.Command{
	Use: "create",

	Short: "Create Loadbalancer",
	Long:  `Create Loadbalancer`,
	RunE: func(cmd *cobra.Command, args []string) error {
		createParam, err := params.NewCreateLoadbalancerParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("create parameter: \n%s\n", debugMarshalIndent(createParam))
		return err
	},
}

var loadbalancerReadCmd = &cobra.Command{
	Use: "read",

	Short: "Read Loadbalancer",
	Long:  `Read Loadbalancer`,
	RunE: func(cmd *cobra.Command, args []string) error {
		readParam, err := params.NewReadLoadbalancerParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("read parameter: \n%s\n", debugMarshalIndent(readParam))
		return err
	},
}

var loadbalancerUpdateCmd = &cobra.Command{
	Use: "update",

	Short: "Update Loadbalancer",
	Long:  `Update Loadbalancer`,
	RunE: func(cmd *cobra.Command, args []string) error {
		updateParam, err := params.NewUpdateLoadbalancerParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("update parameter: \n%s\n", debugMarshalIndent(updateParam))
		return err
	},
}

var loadbalancerDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"rm"},
	Short:   "Delete Loadbalancer",
	Long:    `Delete Loadbalancer`,
	RunE: func(cmd *cobra.Command, args []string) error {
		deleteParam, err := params.NewDeleteLoadbalancerParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("delete parameter: \n%s\n", debugMarshalIndent(deleteParam))
		return err
	},
}

var loadbalancerBootCmd = &cobra.Command{
	Use:     "boot",
	Aliases: []string{"power-on"},
	Short:   "Boot Loadbalancer",
	Long:    `Boot Loadbalancer`,
	RunE: func(cmd *cobra.Command, args []string) error {
		bootParam, err := params.NewBootLoadbalancerParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("boot parameter: \n%s\n", debugMarshalIndent(bootParam))
		return err
	},
}

var loadbalancerShutdownCmd = &cobra.Command{
	Use:     "shutdown",
	Aliases: []string{"power-off"},
	Short:   "Shutdown Loadbalancer",
	Long:    `Shutdown Loadbalancer`,
	RunE: func(cmd *cobra.Command, args []string) error {
		shutdownParam, err := params.NewShutdownLoadbalancerParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("shutdown parameter: \n%s\n", debugMarshalIndent(shutdownParam))
		return err
	},
}

var loadbalancerShutdownForceCmd = &cobra.Command{
	Use:     "shutdown-force",
	Aliases: []string{"stop"},
	Short:   "ShutdownForce Loadbalancer",
	Long:    `ShutdownForce Loadbalancer`,
	RunE: func(cmd *cobra.Command, args []string) error {
		shutdownForceParam, err := params.NewShutdownForceLoadbalancerParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("shutdown-force parameter: \n%s\n", debugMarshalIndent(shutdownForceParam))
		return err
	},
}

var loadbalancerResetCmd = &cobra.Command{
	Use: "reset",

	Short: "Reset Loadbalancer",
	Long:  `Reset Loadbalancer`,
	RunE: func(cmd *cobra.Command, args []string) error {
		resetParam, err := params.NewResetLoadbalancerParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("reset parameter: \n%s\n", debugMarshalIndent(resetParam))
		return err
	},
}

var loadbalancerWaitForBootCmd = &cobra.Command{
	Use: "wait-for-boot",

	Short: "Wait until boot is completed",
	Long:  `Wait until boot is completed`,
	RunE: func(cmd *cobra.Command, args []string) error {
		waitForBootParam, err := params.NewWaitForBootLoadbalancerParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("wait-for-boot parameter: \n%s\n", debugMarshalIndent(waitForBootParam))
		return err
	},
}

var loadbalancerWaitForDownCmd = &cobra.Command{
	Use: "wait-for-down",

	Short: "Wait until shutdown is completed",
	Long:  `Wait until shutdown is completed`,
	RunE: func(cmd *cobra.Command, args []string) error {
		waitForDownParam, err := params.NewWaitForDownLoadbalancerParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("wait-for-down parameter: \n%s\n", debugMarshalIndent(waitForDownParam))
		return err
	},
}

var loadbalancerVipInfoCmd = &cobra.Command{
	Use: "vip-info",

	Short: "Show information of VIP(s)",
	Long:  `Show information of VIP(s)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		vipInfoParam, err := params.NewVipInfoLoadbalancerParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("vip-info parameter: \n%s\n", debugMarshalIndent(vipInfoParam))
		return err
	},
}

var loadbalancerVipAddCmd = &cobra.Command{
	Use: "vip-add",

	Short: "Add VIP to LoadBalancer",
	Long:  `Add VIP to LoadBalancer`,
	RunE: func(cmd *cobra.Command, args []string) error {
		vipAddParam, err := params.NewVipAddLoadbalancerParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("vip-add parameter: \n%s\n", debugMarshalIndent(vipAddParam))
		return err
	},
}

var loadbalancerVipUpdateCmd = &cobra.Command{
	Use: "vip-update",

	Short: "Update VIP",
	Long:  `Update VIP`,
	RunE: func(cmd *cobra.Command, args []string) error {
		vipUpdateParam, err := params.NewVipUpdateLoadbalancerParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("vip-update parameter: \n%s\n", debugMarshalIndent(vipUpdateParam))
		return err
	},
}

var loadbalancerVipDeleteCmd = &cobra.Command{
	Use: "vip-delete",

	Short: "Delete VIP from LoadBalancer",
	Long:  `Delete VIP from LoadBalancer`,
	RunE: func(cmd *cobra.Command, args []string) error {
		vipDeleteParam, err := params.NewVipDeleteLoadbalancerParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("vip-delete parameter: \n%s\n", debugMarshalIndent(vipDeleteParam))
		return err
	},
}

var loadbalancerServerInfoCmd = &cobra.Command{
	Use: "server-info",

	Short: "Show servers under VIP(s)",
	Long:  `Show servers under VIP(s)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		serverInfoParam, err := params.NewServerInfoLoadbalancerParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("server-info parameter: \n%s\n", debugMarshalIndent(serverInfoParam))
		return err
	},
}

var loadbalancerServerAddCmd = &cobra.Command{
	Use: "server-add",

	Short: "Add server under VIP(s)",
	Long:  `Add server under VIP(s)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		serverAddParam, err := params.NewServerAddLoadbalancerParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("server-add parameter: \n%s\n", debugMarshalIndent(serverAddParam))
		return err
	},
}

var loadbalancerServerUpdateCmd = &cobra.Command{
	Use: "server-update",

	Short: "Update server under VIP(s)",
	Long:  `Update server under VIP(s)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		serverUpdateParam, err := params.NewServerUpdateLoadbalancerParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("server-update parameter: \n%s\n", debugMarshalIndent(serverUpdateParam))
		return err
	},
}

var loadbalancerServerDeleteCmd = &cobra.Command{
	Use: "server-delete",

	Short: "Delete server under VIP(s)",
	Long:  `Delete server under VIP(s)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		serverDeleteParam, err := params.NewServerDeleteLoadbalancerParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("server-delete parameter: \n%s\n", debugMarshalIndent(serverDeleteParam))
		return err
	},
}

var loadbalancerMonitorCmd = &cobra.Command{
	Use: "monitor",

	Short: "Monitor Loadbalancer",
	Long:  `Monitor Loadbalancer`,
	RunE: func(cmd *cobra.Command, args []string) error {
		monitorParam, err := params.NewMonitorLoadbalancerParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("monitor parameter: \n%s\n", debugMarshalIndent(monitorParam))
		return err
	},
}

func init() {
	parent := loadbalancerCmd
	parent.AddCommand(loadbalancerListCmd)
	parent.AddCommand(loadbalancerCreateCmd)
	parent.AddCommand(loadbalancerReadCmd)
	parent.AddCommand(loadbalancerUpdateCmd)
	parent.AddCommand(loadbalancerDeleteCmd)
	parent.AddCommand(loadbalancerBootCmd)
	parent.AddCommand(loadbalancerShutdownCmd)
	parent.AddCommand(loadbalancerShutdownForceCmd)
	parent.AddCommand(loadbalancerResetCmd)
	parent.AddCommand(loadbalancerWaitForBootCmd)
	parent.AddCommand(loadbalancerWaitForDownCmd)
	parent.AddCommand(loadbalancerVipInfoCmd)
	parent.AddCommand(loadbalancerVipAddCmd)
	parent.AddCommand(loadbalancerVipUpdateCmd)
	parent.AddCommand(loadbalancerVipDeleteCmd)
	parent.AddCommand(loadbalancerServerInfoCmd)
	parent.AddCommand(loadbalancerServerAddCmd)
	parent.AddCommand(loadbalancerServerUpdateCmd)
	parent.AddCommand(loadbalancerServerDeleteCmd)
	parent.AddCommand(loadbalancerMonitorCmd)
	rootCmd.AddCommand(loadbalancerCmd)
}
