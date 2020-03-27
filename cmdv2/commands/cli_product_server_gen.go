// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-v2-commands'; DO NOT EDIT

package commands

import (
	"fmt"

	"github.com/sacloud/usacloud/cmdv2/params"
	"github.com/spf13/cobra"
)

// productserverCmd represents the productserver command
var productserverCmd = &cobra.Command{
	Use:   "product-server",
	Short: "A manage commands of ProductServer",
	Long:  `A manage commands of ProductServer`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO not implements: call list func as default
	},
}

var productserverListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "find"},
	Short:   "List Productserver (default)",
	Long:    `List Productserver (default)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		listParam, err := params.NewListProductserverParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("list parameter: \n%s\n", debugMarshalIndent(listParam))
		return err
	},
}

var productserverReadCmd = &cobra.Command{
	Use: "read",

	Short: "Read Productserver",
	Long:  `Read Productserver`,
	RunE: func(cmd *cobra.Command, args []string) error {
		readParam, err := params.NewReadProductserverParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("read parameter: \n%s\n", debugMarshalIndent(readParam))
		return err
	},
}

func init() {
	parent := productserverCmd
	parent.AddCommand(productserverListCmd)
	parent.AddCommand(productserverReadCmd)
	rootCmd.AddCommand(productserverCmd)
}
