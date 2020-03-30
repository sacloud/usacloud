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

package commands

import (
	"encoding/json"
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

func debugMarshalIndent(in interface{}) string {
	data, err := json.MarshalIndent(in, "", "    ")
	if err != nil {
		return ""
	}
	return string(data)
}
