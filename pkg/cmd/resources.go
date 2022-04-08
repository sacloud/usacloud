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

//go:generate go run github.com/sacloud/usacloud/tools/gen-commands/
package cmd

import (
	"fmt"
	"os"

	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/archive"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/authstatus"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/autobackup"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/bill"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/bridge"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/cdrom"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/certificateauthority"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/containerregistry"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/coupon"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/database"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/disk"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/diskplan"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/dns"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/enhanceddb"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/esme"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/gslb"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/icon"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/iface"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/internet"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/internetplan"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/ipaddress"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/ipv6addr"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/ipv6net"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/license"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/licenseinfo"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/loadbalancer"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/localrouter"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/mobilegateway"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/nfs"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/note"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/packetfilter"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/privatehost"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/privatehostplan"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/proxylb"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/region"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/server"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/serverplan"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/serviceclass"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/sim"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/simplemonitor"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/sshkey"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/subnet"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/swytch"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/vpcrouter"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/webaccelerator"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iaas/zone"
	"github.com/sacloud/usacloud/pkg/cmd/commands/rest"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/cmd/root"
	_ "github.com/sacloud/usacloud/pkg/cmd/version"
	"github.com/spf13/cobra"
)

var (
	IaaSResources = core.Resources{
		// libsacloud services
		archive.Resource,
		authstatus.Resource,
		autobackup.Resource,
		bill.Resource,
		bridge.Resource,
		cdrom.Resource,
		certificateauthority.Resource,
		containerregistry.Resource,
		coupon.Resource,
		database.Resource,
		disk.Resource,
		diskplan.Resource,
		dns.Resource,
		enhanceddb.Resource,
		esme.Resource,
		gslb.Resource,
		icon.Resource,
		iface.Resource,
		internet.Resource,
		internetplan.Resource,
		ipaddress.Resource,
		ipv6addr.Resource,
		ipv6net.Resource,
		license.Resource,
		licenseinfo.Resource,
		loadbalancer.Resource,
		localrouter.Resource,
		mobilegateway.Resource,
		nfs.Resource,
		note.Resource,
		packetfilter.Resource,
		privatehost.Resource,
		privatehostplan.Resource,
		proxylb.Resource,
		region.Resource,
		server.Resource,
		serverplan.Resource,
		serviceclass.Resource,
		sim.Resource,
		simplemonitor.Resource,
		sshkey.Resource,
		subnet.Resource,
		swytch.Resource,
		vpcrouter.Resource,
		zone.Resource,
		// libsacloud service以外のマニュアル実装分
	}

	MiscResources = core.Resources{
		rest.Resource,
		webaccelerator.Resource,
	}
)

func Resources() core.Resources {
	rs := core.Resources{}
	rs = append(rs, IaaSResources...)
	rs = append(rs, MiscResources...)
	return rs
}

func initCommands() {
	initIaasCommands()
	initMiscCommands()
}

func initIaasCommands() {
	iaasCmd := &cobra.Command{
		Use:   "iaas",
		Short: "SubCommands for IaaS",
		Long:  "SubCommands for IaaS",
	}

	for _, r := range IaaSResources {
		cmd := r.CLICommand()
		if len(cmd.Commands()) > 0 {
			iaasCmd.AddCommand(cmd)
			addHiddenSubCommandToRoot(cmd)
		}
	}
	core.SetSubCommandsUsage(iaasCmd, IaaSResources.CategorizedResources())
	root.Command.AddCommand(iaasCmd)
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

		// コマンドの中にはデフォルトコマンドとして自身のサブコマンドを呼ぶ場合(auth-statusなど)があるため、
		// 末端(childrenがない)コマンドにだけ設定する。(この条件がないと表示が重複する)
		if len(children) == 0 {
			cmd.PersistentPreRun = func(own *cobra.Command, args []string) {
				// この段階ではctx.IO()が参照できないため標準エラーに出力する
				fmt.Fprintln(os.Stderr, "[WARN] This command is deprecated. Please use the command under the `usacloud iaas` subcommand instead.") // nolint
				if cmd.HasParent() {
					parent := cmd.Parent()
					if parent.PersistentPreRun != nil {
						parent.PersistentPreRun(own, args)
					}
				}
			}
		}
	}
	setChildFn(&c)

	c.Hidden = true
	root.Command.AddCommand(&c)
}

func initMiscCommands() {
	for _, r := range MiscResources {
		cmd := r.CLICommand()
		if len(cmd.Commands()) > 0 {
			root.Command.AddCommand(cmd)
		}
	}
}
