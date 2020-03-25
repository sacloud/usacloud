package commands

import (
	"log"
	"os"

	"github.com/sacloud/usacloud/pkg/utils"

	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

// rootCmd represents the base command when called without any sub-commands
var rootCmd = &cobra.Command{
	Use:   "usacloud [global options] <command> <sub-command> [options] [arguments]",
	Short: "Usacloud is CLI for manage to resources on the SAKURA Cloud",
	Long:  `CLI to manage to resources on the SAKURA Cloud`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		cliOption.loadGlobalFlags(cmd.PersistentFlags())
		return utils.FlattenErrors(cliOption.Validate(true))
	},
}

func init() {
	rootCmd.Flags().SortFlags = false
	rootCmd.PersistentFlags().SortFlags = false

	initGlobalFlags(rootCmd.PersistentFlags())
}
