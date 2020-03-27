// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-v2-commands'; DO NOT EDIT

package commands

import (
	"fmt"

	"github.com/sacloud/usacloud/cmdv2/params"
	"github.com/spf13/cobra"
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Show summary of resource usage",
	Long:  `Show summary of resource usage`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO not implements: call show func as default
	},
}

var summaryShowCmd = &cobra.Command{
	Use: "show",

	Short: "Show Summary (default)",
	Long:  `Show Summary (default)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		showParam, err := params.NewShowSummaryParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("show parameter: \n%s\n", debugMarshalIndent(showParam))
		return err
	},
}

func init() {
	parent := summaryCmd
	parent.AddCommand(summaryShowCmd)
	rootCmd.AddCommand(summaryCmd)
}
