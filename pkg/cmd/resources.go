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

	"github.com/sacloud/usacloud/pkg/cmd/commands/archive"
	"github.com/sacloud/usacloud/pkg/cmd/commands/authstatus"
	"github.com/sacloud/usacloud/pkg/cmd/commands/autobackup"
	"github.com/sacloud/usacloud/pkg/cmd/commands/bill"
	"github.com/sacloud/usacloud/pkg/cmd/commands/bridge"
	"github.com/sacloud/usacloud/pkg/cmd/commands/cdrom"
	"github.com/sacloud/usacloud/pkg/cmd/commands/certificateauthority"
	"github.com/sacloud/usacloud/pkg/cmd/commands/containerregistry"
	"github.com/sacloud/usacloud/pkg/cmd/commands/coupon"
	"github.com/sacloud/usacloud/pkg/cmd/commands/database"
	"github.com/sacloud/usacloud/pkg/cmd/commands/disk"
	"github.com/sacloud/usacloud/pkg/cmd/commands/diskplan"
	"github.com/sacloud/usacloud/pkg/cmd/commands/dns"
	"github.com/sacloud/usacloud/pkg/cmd/commands/enhanceddb"
	"github.com/sacloud/usacloud/pkg/cmd/commands/esme"
	"github.com/sacloud/usacloud/pkg/cmd/commands/gslb"
	"github.com/sacloud/usacloud/pkg/cmd/commands/icon"
	"github.com/sacloud/usacloud/pkg/cmd/commands/iface"
	"github.com/sacloud/usacloud/pkg/cmd/commands/internet"
	"github.com/sacloud/usacloud/pkg/cmd/commands/internetplan"
	"github.com/sacloud/usacloud/pkg/cmd/commands/ipaddress"
	"github.com/sacloud/usacloud/pkg/cmd/commands/ipv6addr"
	"github.com/sacloud/usacloud/pkg/cmd/commands/ipv6net"
	"github.com/sacloud/usacloud/pkg/cmd/commands/license"
	"github.com/sacloud/usacloud/pkg/cmd/commands/licenseinfo"
	"github.com/sacloud/usacloud/pkg/cmd/commands/loadbalancer"
	"github.com/sacloud/usacloud/pkg/cmd/commands/localrouter"
	"github.com/sacloud/usacloud/pkg/cmd/commands/mobilegateway"
	"github.com/sacloud/usacloud/pkg/cmd/commands/nfs"
	"github.com/sacloud/usacloud/pkg/cmd/commands/note"
	"github.com/sacloud/usacloud/pkg/cmd/commands/packetfilter"
	"github.com/sacloud/usacloud/pkg/cmd/commands/privatehost"
	"github.com/sacloud/usacloud/pkg/cmd/commands/privatehostplan"
	"github.com/sacloud/usacloud/pkg/cmd/commands/proxylb"
	"github.com/sacloud/usacloud/pkg/cmd/commands/region"
	"github.com/sacloud/usacloud/pkg/cmd/commands/rest"
	"github.com/sacloud/usacloud/pkg/cmd/commands/server"
	"github.com/sacloud/usacloud/pkg/cmd/commands/serverplan"
	"github.com/sacloud/usacloud/pkg/cmd/commands/serviceclass"
	"github.com/sacloud/usacloud/pkg/cmd/commands/sim"
	"github.com/sacloud/usacloud/pkg/cmd/commands/simplemonitor"
	"github.com/sacloud/usacloud/pkg/cmd/commands/sshkey"
	"github.com/sacloud/usacloud/pkg/cmd/commands/subnet"
	"github.com/sacloud/usacloud/pkg/cmd/commands/swytch"
	"github.com/sacloud/usacloud/pkg/cmd/commands/vpcrouter"
	"github.com/sacloud/usacloud/pkg/cmd/commands/webaccelerator"
	"github.com/sacloud/usacloud/pkg/cmd/commands/zone"
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
