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

package iaas

import (
	"github.com/sacloud/usacloud/pkg/commands/iaas/category"
	"github.com/sacloud/usacloud/pkg/commands/root"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "iaas",
	Short: "SubCommands for IaaS",
	Long:  "SubCommands for IaaS",
}

func init() {
	for _, r := range Resources {
		cmd := r.CLICommand()
		if len(cmd.Commands()) > 0 {
			Command.AddCommand(cmd)
			addHiddenSubCommandToRoot(cmd)
		}
	}
	core.SetSubCommandsUsage(Command, Resources.CategorizedResources(category.ResourceCategories))
}

// addHiddenSubCommandToRoot 互換性維持のためにroot直下にHidden=trueの状態でコマンドを追加する
func addHiddenSubCommandToRoot(cmd *cobra.Command) {
	c := *cmd
	var setChildFn func(cmd *cobra.Command)
	setChildFn = func(cmd *cobra.Command) {
		children := cmd.Commands()
		cmd.ResetCommands()
		for _, child := range children {
			c := *child
			setChildFn(&c)
			cmd.AddCommand(&c)
		}

		// Note:当面は警告を表示しないようにする。https://github.com/sacloud/usacloud/issues/911
		//
		// コマンドの中にはデフォルトコマンドとして自身のサブコマンドを呼ぶ場合(auth-statusなど)があるため、
		// 末端(childrenがない)コマンドにだけ設定する。(この条件がないと表示が重複する)
		//if len(children) == 0 {
		//	cmd.PersistentPreRun = func(own *cobra.Command, args []string) {
		//		// この段階ではctx.IO()が参照できないため標準エラーに出力する
		//		fmt.Fprintln(os.Stderr, "[WARN] This command is deprecated. Please use the command under the `usacloud iaas` subcommand instead.") // nolint
		//		if cmd.HasParent() {
		//			parent := cmd.Parent()
		//			if parent.PersistentPreRun != nil {
		//				parent.PersistentPreRun(own, args)
		//			}
		//		}
		//	}
		//}
	}
	setChildFn(&c)

	c.Hidden = true
	root.Command.AddCommand(&c)
}
