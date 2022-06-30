// Copyright 2017-2022 The Usacloud Authors
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

package completion

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long: `To load completions:

Bash:

$ source <(usacloud completion bash)

# To load completions for each session, execute once:
Linux:
  $ usacloud completion bash > /etc/bash_completion.d/usacloud
MacOS:
  $ usacloud completion bash > /usr/local/etc/bash_completion.d/usacloud

Zsh:

# If shell completion is not already enabled in your environment you will need
# to enable it.  You can execute the following once:

$ echo "autoload -U compinit; compinit" >> ~/.zshrc

# To load completions for each session, execute once:
$ usacloud completion zsh > "${fpath[1]}/_usacloud"

# You will need to start a new shell for this setup to take effect.

Fish:

$ usacloud completion fish | source

# To load completions for each session, execute once:
$ usacloud completion fish > ~/.config/fish/completions/usacloud.fish
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// iaasサブコマンド配下のコマンドは互換性維持のためHidden=trueでRoot直下にも配置されている。
		// これらのコマンドも補完するためにこのタイミングでRoot直下のコマンドのHiddenをfalseにする。
		// Note: もしRoot直下にHidden=trueなコマンドが欲しくなった場合は修正が必要
		for _, c := range cmd.Root().Commands() {
			c.Hidden = false
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		switch args[0] {
		case "bash":
			err = cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			err = cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			err = cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			err = cmd.Root().GenPowerShellCompletion(os.Stdout)
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}
