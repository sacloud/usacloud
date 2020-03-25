package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// stubCmd represents the stub command
var stubCmd = &cobra.Command{
	Use:   "stub",
	Short: "A stub command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		dumpFlagsAndArgs(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(stubCmd)
}

func dumpFlagsAndArgs(cmd *cobra.Command, args []string) {
	fmt.Println("Arguments:")
	for _, arg := range args {
		fmt.Println("\t" + arg)
	}
	fmt.Println("Flags:")
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		fmt.Println("\t" + f.Name + "\t: " + f.Value.String())
	})
}
