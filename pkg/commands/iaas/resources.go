// Copyright 2017-2022 The sacloud/usacloud Authors
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
	"github.com/sacloud/usacloud/pkg/commands/iaas/archive"
	"github.com/sacloud/usacloud/pkg/commands/iaas/authstatus"
	"github.com/sacloud/usacloud/pkg/commands/iaas/autobackup"
	"github.com/sacloud/usacloud/pkg/commands/iaas/autoscale"
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
	"github.com/sacloud/usacloud/pkg/commands/iaas/zone"
	"github.com/sacloud/usacloud/pkg/core"
)

var Resources = core.Resources{
	archive.Resource,
	authstatus.Resource,
	autobackup.Resource,
	autoscale.Resource,
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
}
