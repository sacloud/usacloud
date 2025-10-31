// Copyright 2017-2025 The sacloud/usacloud Authors
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

package root

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"sync"

	"github.com/fatih/color"
	"github.com/sacloud/usacloud/pkg/config"
	"github.com/sacloud/usacloud/pkg/printer"
	"github.com/sacloud/usacloud/pkg/version"
	"github.com/spf13/cobra"
)

// Command represents the base command when called without any sub-commands
var Command = &cobra.Command{
	Use:   "usacloud [global options] <command> <sub-command> [options] [arguments]",
	Short: "Usacloud is CLI for manage to resources on the SAKURA Cloud",
	Long:  `CLI to manage to resources on the SAKURA Cloud`,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if cmd.Name() == "update-self" { // update-selfだけ例外扱い。Note: 例外扱いするコマンドが増えるようであれば実装を修正する
			return
		}
		once.Do(func() {
			noColor, _ := cmd.PersistentFlags().GetBool("no-color") // ignore error
			alertNewVersionReleased(noColor)
		})
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		v, err := cmd.Flags().GetBool("version")
		if err != nil {
			return err
		}
		if v {
			cmd.Root().SetArgs([]string{"version"})
			return cmd.Root().Execute()
		}
		return cmd.Help()
	},
}

func init() {
	Command.Flags().SortFlags = false
	Command.PersistentFlags().SortFlags = false

	config.TheClient.SetEnviron(slices.Clone(os.Environ()))
	config.InitConfig(Command.PersistentFlags())

	// This AddGoFlagSet() silently ignores duplicated flags;
	// They need extra touches.  Done in config.LoadConfigValue().
	Command.PersistentFlags().AddGoFlagSet(config.TheClient.FlagSet(flag.ContinueOnError))
}

const newVersionAlertTemplate = `
[NOTICE]
  Current version of Usacloud is out of date.

    - Current version: v%s
    - Latest version:  %s

  You can update Usacloud in the following ways.

    - Use a package manager (e.g. ` + "`brew upgrade sacloud/usacloud/usacloud`" + `)
    - Use ` + "`usacloud update-self`" + ` command
    - Download directly from %s

`

var once sync.Once

func alertNewVersionReleased(noColor bool) {
	releaseInfo, err := version.NewVersionReleased()
	if err != nil {
		handleGatheringReleaseInfoError(err)
	}
	if releaseInfo != nil {
		newVersionReleased, err := releaseInfo.Release.GreaterThanCurrent()
		if err != nil {
			handleGatheringReleaseInfoError(err)
		}
		if newVersionReleased {
			p := &printer.Printer{NoColor: noColor}
			p.Fprintf(os.Stderr, color.New(color.FgYellow), newVersionAlertTemplate, version.Version, releaseInfo.Release.TagName, releaseInfo.Release.URL)
		}
	}
}

func handleGatheringReleaseInfoError(err error) {
	if err == nil {
		return
	}
	if os.Getenv("USACLOUD_TRACE") != "" {
		fmt.Fprintln(os.Stderr, err)
	}
}
