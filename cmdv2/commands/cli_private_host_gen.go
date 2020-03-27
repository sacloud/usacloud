// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-v2-commands'; DO NOT EDIT

package commands

import (
	"fmt"

	"github.com/sacloud/usacloud/cmdv2/params"
	"github.com/spf13/cobra"
)

// privatehostCmd represents the privatehost command
var privatehostCmd = &cobra.Command{
	Use:   "private-host",
	Short: "A manage commands of PrivateHost",
	Long:  `A manage commands of PrivateHost`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

var privatehostListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "find", "selector"},
	Short:   "List Privatehost",
	Long:    `List Privatehost`,
	RunE: func(cmd *cobra.Command, args []string) error {
		listParam, err := params.NewListPrivatehostParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("list parameter: \n%s\n", debugMarshalIndent(listParam))
		return err
	},
}

var privatehostCreateCmd = &cobra.Command{
	Use: "create",

	Short: "Create Privatehost",
	Long:  `Create Privatehost`,
	RunE: func(cmd *cobra.Command, args []string) error {
		createParam, err := params.NewCreatePrivatehostParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("create parameter: \n%s\n", debugMarshalIndent(createParam))
		return err
	},
}

var privatehostReadCmd = &cobra.Command{
	Use: "read",

	Short: "Read Privatehost",
	Long:  `Read Privatehost`,
	RunE: func(cmd *cobra.Command, args []string) error {
		readParam, err := params.NewReadPrivatehostParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("read parameter: \n%s\n", debugMarshalIndent(readParam))
		return err
	},
}

var privatehostUpdateCmd = &cobra.Command{
	Use: "update",

	Short: "Update Privatehost",
	Long:  `Update Privatehost`,
	RunE: func(cmd *cobra.Command, args []string) error {
		updateParam, err := params.NewUpdatePrivatehostParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("update parameter: \n%s\n", debugMarshalIndent(updateParam))
		return err
	},
}

var privatehostDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"rm"},
	Short:   "Delete Privatehost",
	Long:    `Delete Privatehost`,
	RunE: func(cmd *cobra.Command, args []string) error {
		deleteParam, err := params.NewDeletePrivatehostParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("delete parameter: \n%s\n", debugMarshalIndent(deleteParam))
		return err
	},
}

var privatehostServerInfoCmd = &cobra.Command{
	Use:     "server-info",
	Aliases: []string{"server-list"},
	Short:   "ServerInfo Privatehost",
	Long:    `ServerInfo Privatehost`,
	RunE: func(cmd *cobra.Command, args []string) error {
		serverInfoParam, err := params.NewServerInfoPrivatehostParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("server-info parameter: \n%s\n", debugMarshalIndent(serverInfoParam))
		return err
	},
}

var privatehostServerAddCmd = &cobra.Command{
	Use: "server-add",

	Short: "ServerAdd Privatehost",
	Long:  `ServerAdd Privatehost`,
	RunE: func(cmd *cobra.Command, args []string) error {
		serverAddParam, err := params.NewServerAddPrivatehostParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("server-add parameter: \n%s\n", debugMarshalIndent(serverAddParam))
		return err
	},
}

var privatehostServerDeleteCmd = &cobra.Command{
	Use: "server-delete",

	Short: "ServerDelete Privatehost",
	Long:  `ServerDelete Privatehost`,
	RunE: func(cmd *cobra.Command, args []string) error {
		serverDeleteParam, err := params.NewServerDeletePrivatehostParam(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("server-delete parameter: \n%s\n", debugMarshalIndent(serverDeleteParam))
		return err
	},
}

func init() {
	parent := privatehostCmd
	parent.AddCommand(privatehostListCmd)
	parent.AddCommand(privatehostCreateCmd)
	parent.AddCommand(privatehostReadCmd)
	parent.AddCommand(privatehostUpdateCmd)
	parent.AddCommand(privatehostDeleteCmd)
	parent.AddCommand(privatehostServerInfoCmd)
	parent.AddCommand(privatehostServerAddCmd)
	parent.AddCommand(privatehostServerDeleteCmd)
	rootCmd.AddCommand(privatehostCmd)
}
