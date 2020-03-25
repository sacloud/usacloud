// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-v2-commands'; DO NOT EDIT

package commands

import (
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
	Run: func(cmd *cobra.Command, args []string) {
		// TODO not implements
	},
}

var loadbalancerCreateCmd = &cobra.Command{
	Use: "create",

	Short: "Create Loadbalancer",
	Long:  `Create Loadbalancer`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO not implements
	},
}

var loadbalancerReadCmd = &cobra.Command{
	Use: "read",

	Short: "Read Loadbalancer",
	Long:  `Read Loadbalancer`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO not implements
	},
}

var loadbalancerUpdateCmd = &cobra.Command{
	Use: "update",

	Short: "Update Loadbalancer",
	Long:  `Update Loadbalancer`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO not implements
	},
}

var loadbalancerDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"rm"},
	Short:   "Delete Loadbalancer",
	Long:    `Delete Loadbalancer`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO not implements
	},
}

var loadbalancerBootCmd = &cobra.Command{
	Use:     "boot",
	Aliases: []string{"power-on"},
	Short:   "Boot Loadbalancer",
	Long:    `Boot Loadbalancer`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO not implements
	},
}

var loadbalancerShutdownCmd = &cobra.Command{
	Use:     "shutdown",
	Aliases: []string{"power-off"},
	Short:   "Shutdown Loadbalancer",
	Long:    `Shutdown Loadbalancer`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO not implements
	},
}

var loadbalancerShutdownForceCmd = &cobra.Command{
	Use:     "shutdown-force",
	Aliases: []string{"stop"},
	Short:   "ShutdownForce Loadbalancer",
	Long:    `ShutdownForce Loadbalancer`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO not implements
	},
}

var loadbalancerResetCmd = &cobra.Command{
	Use: "reset",

	Short: "Reset Loadbalancer",
	Long:  `Reset Loadbalancer`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO not implements
	},
}

var loadbalancerWaitForBootCmd = &cobra.Command{
	Use: "wait-for-boot",

	Short: "Wait until boot is completed",
	Long:  `Wait until boot is completed`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO not implements
	},
}

var loadbalancerWaitForDownCmd = &cobra.Command{
	Use: "wait-for-down",

	Short: "Wait until shutdown is completed",
	Long:  `Wait until shutdown is completed`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO not implements
	},
}

var loadbalancerVipInfoCmd = &cobra.Command{
	Use: "vip-info",

	Short: "Show information of VIP(s)",
	Long:  `Show information of VIP(s)`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO not implements
	},
}

var loadbalancerVipAddCmd = &cobra.Command{
	Use: "vip-add",

	Short: "Add VIP to LoadBalancer",
	Long:  `Add VIP to LoadBalancer`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO not implements
	},
}

var loadbalancerVipUpdateCmd = &cobra.Command{
	Use: "vip-update",

	Short: "Update VIP",
	Long:  `Update VIP`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO not implements
	},
}

var loadbalancerVipDeleteCmd = &cobra.Command{
	Use: "vip-delete",

	Short: "Delete VIP from LoadBalancer",
	Long:  `Delete VIP from LoadBalancer`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO not implements
	},
}

var loadbalancerServerInfoCmd = &cobra.Command{
	Use: "server-info",

	Short: "Show servers under VIP(s)",
	Long:  `Show servers under VIP(s)`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO not implements
	},
}

var loadbalancerServerAddCmd = &cobra.Command{
	Use: "server-add",

	Short: "Add server under VIP(s)",
	Long:  `Add server under VIP(s)`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO not implements
	},
}

var loadbalancerServerUpdateCmd = &cobra.Command{
	Use: "server-update",

	Short: "Update server under VIP(s)",
	Long:  `Update server under VIP(s)`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO not implements
	},
}

var loadbalancerServerDeleteCmd = &cobra.Command{
	Use: "server-delete",

	Short: "Delete server under VIP(s)",
	Long:  `Delete server under VIP(s)`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO not implements
	},
}

var loadbalancerMonitorCmd = &cobra.Command{
	Use: "monitor",

	Short: "Monitor Loadbalancer",
	Long:  `Monitor Loadbalancer`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO not implements
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
