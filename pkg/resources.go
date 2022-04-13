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
package pkg

import (
	"fmt"
	"os"

	"github.com/sacloud/usacloud/pkg/commands/completion"
	"github.com/sacloud/usacloud/pkg/commands/iaas"
	"github.com/sacloud/usacloud/pkg/commands/iaas/archive"
	"github.com/sacloud/usacloud/pkg/commands/iaas/authstatus"
	"github.com/sacloud/usacloud/pkg/commands/iaas/autobackup"
	"github.com/sacloud/usacloud/pkg/commands/iaas/bill"
	"github.com/sacloud/usacloud/pkg/commands/iaas/bridge"
	"github.com/sacloud/usacloud/pkg/commands/iaas/cdrom"
	"github.com/sacloud/usacloud/pkg/commands/iaas/certificateauthority"
	"github.com/sacloud/usacloud/pkg/commands/iaas/containerregistry"
	"github.com/sacloud/usacloud/pkg/commands/iaas/coupon"
	"github.com/sacloud/usacloud/pkg/commands/iaas/database"
	"github.com/sacloud/usacloud/pkg/commands/iaas/disk"
	"github.com/sacloud/usacloud/pkg/commands/iaas/diskplan"
	"github.com/sacloud/usacloud/pkg/commands/iaas/dns"
	"github.com/sacloud/usacloud/pkg/commands/iaas/enhanceddb"
	"github.com/sacloud/usacloud/pkg/commands/iaas/esme"
	"github.com/sacloud/usacloud/pkg/commands/iaas/gslb"
	"github.com/sacloud/usacloud/pkg/commands/iaas/icon"
	"github.com/sacloud/usacloud/pkg/commands/iaas/iface"
	"github.com/sacloud/usacloud/pkg/commands/iaas/internet"
	"github.com/sacloud/usacloud/pkg/commands/iaas/internetplan"
	"github.com/sacloud/usacloud/pkg/commands/iaas/ipaddress"
	"github.com/sacloud/usacloud/pkg/commands/iaas/ipv6addr"
	"github.com/sacloud/usacloud/pkg/commands/iaas/ipv6net"
	"github.com/sacloud/usacloud/pkg/commands/iaas/license"
	"github.com/sacloud/usacloud/pkg/commands/iaas/licenseinfo"
	"github.com/sacloud/usacloud/pkg/commands/iaas/loadbalancer"
	"github.com/sacloud/usacloud/pkg/commands/iaas/localrouter"
	"github.com/sacloud/usacloud/pkg/commands/iaas/mobilegateway"
	"github.com/sacloud/usacloud/pkg/commands/iaas/nfs"
	"github.com/sacloud/usacloud/pkg/commands/iaas/note"
	"github.com/sacloud/usacloud/pkg/commands/iaas/packetfilter"
	"github.com/sacloud/usacloud/pkg/commands/iaas/privatehost"
	"github.com/sacloud/usacloud/pkg/commands/iaas/privatehostplan"
	"github.com/sacloud/usacloud/pkg/commands/iaas/proxylb"
	"github.com/sacloud/usacloud/pkg/commands/iaas/region"
	"github.com/sacloud/usacloud/pkg/commands/iaas/server"
	"github.com/sacloud/usacloud/pkg/commands/iaas/serverplan"
	"github.com/sacloud/usacloud/pkg/commands/iaas/serviceclass"
	"github.com/sacloud/usacloud/pkg/commands/iaas/sim"
	"github.com/sacloud/usacloud/pkg/commands/iaas/simplemonitor"
	"github.com/sacloud/usacloud/pkg/commands/iaas/sshkey"
	"github.com/sacloud/usacloud/pkg/commands/iaas/subnet"
	"github.com/sacloud/usacloud/pkg/commands/iaas/swytch"
	"github.com/sacloud/usacloud/pkg/commands/iaas/vpcrouter"
	"github.com/sacloud/usacloud/pkg/commands/iaas/webaccelerator"
	"github.com/sacloud/usacloud/pkg/commands/iaas/zone"
	"github.com/sacloud/usacloud/pkg/commands/phy"
	phyServer "github.com/sacloud/usacloud/pkg/commands/phy/server"
	phyService "github.com/sacloud/usacloud/pkg/commands/phy/service"
	"github.com/sacloud/usacloud/pkg/commands/rest"
	"github.com/sacloud/usacloud/pkg/commands/root"
	"github.com/sacloud/usacloud/pkg/commands/version"
	"github.com/sacloud/usacloud/pkg/core"
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

	PhyResources = core.Resources{
		phyService.Resource,
		phyServer.Resource,
	}

	MiscResources = core.Resources{
		rest.Resource,
		webaccelerator.Resource,
	}

	RootCommands = []*cobra.Command{
		completion.Command,
		version.Command,
	}
)

func Resources() core.Resources {
	rs := core.Resources{}
	rs = append(rs, IaaSResources...)
	rs = append(rs, PhyResources...)
	rs = append(rs, MiscResources...)
	return rs
}

func initCommands() {
	initIaasCommands()
	initPhyCommands()
	initMiscCommands()
	initRootCommands()
}

func initIaasCommands() {
	for _, r := range IaaSResources {
		cmd := r.CLICommand()
		if len(cmd.Commands()) > 0 {
			iaas.Command.AddCommand(cmd)
			addHiddenSubCommandToRoot(cmd)
		}
	}
	core.SetSubCommandsUsage(iaas.Command, IaaSResources.CategorizedResources(iaas.ResourceCategories))
	root.Command.AddCommand(iaas.Command)
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

func initPhyCommands() {
	for _, r := range PhyResources {
		cmd := r.CLICommand()
		if len(cmd.Commands()) > 0 {
			phy.Command.AddCommand(cmd)
		}
	}
	core.SetSubCommandsUsage(phy.Command, PhyResources.CategorizedResources(phy.ResourceCategories))
	root.Command.AddCommand(phy.Command)
}

func initMiscCommands() {
	for _, r := range MiscResources {
		cmd := r.CLICommand()
		if len(cmd.Commands()) > 0 {
			root.Command.AddCommand(cmd)
		}
	}
}

func initRootCommands() {
	root.Command.AddCommand(RootCommands...)
}
