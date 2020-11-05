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

// Package define .
//go:generate go run github.com/sacloud/usacloud/tools/gen-command-params/
//go:generate go run github.com/sacloud/usacloud/tools/gen-command-funcs/
//go:generate go run github.com/sacloud/usacloud/tools/gen-cli-commands/
//go:generate go run github.com/sacloud/usacloud/tools/gen-cli-resource-finder/
//go:generate go run github.com/sacloud/usacloud/tools/gen-cli-usage/
//go:generate go run github.com/sacloud/usacloud/tools/gen-cli-normalize-flag-name/
//go:generate go run github.com/sacloud/usacloud/tools/gen-cli-root-command/
//go:generate go run github.com/sacloud/usacloud/tools/gen-cli-params-v1/ # TODO 実験的実装
package define

import "github.com/sacloud/usacloud/pkg/schema"

var Resources = map[string]*schema.Resource{
	"AuthStatus": AuthStatusResource(),
	"Archive":    ArchiveResource(),
	"AutoBackup": AutoBackupResource(),
	"Bill":       BillResource(),
	"Bridge":     BridgeResource(),
	"Config":     ConfigResource(),
	"Coupon":     CouponResource(),
	"Database":   DatabaseResource(),
	"Disk":       DiskResource(),
	//"DNS":           DNSResource(),
	//"GSLB":          GSLBResource(),
	//"Icon":          IconResource(),
	//"Interface":     InterfaceResource(),
	//"Internet":      InternetResource(),
	//"IPv4":          IPv4Resource(),
	//"IPv6":          IPv6Resource(),
	//"ISOImage":      ISOImageResource(),
	//"License":       LicenseResource(),
	//"LoadBalancer":  LoadBalancerResource(),
	//"MobileGateway": MobileGatewayResource(),
	//"NFS":           NFSResource(),
	//"ObjectStorage": ObjectStorageResource(),
	//"PacketFilter":  PacketFilterResource(),
	//"ProxyLB":       ProxyLBResource(),
	//// TODO libsacloud v2でPrice API未実装
	////"Price":           PriceResource(),
	//"PrivateHost":     PrivateHostResource(),
	//"ProductDisk":     ProductDiskResource(),
	//"ProductInternet": ProductInternetResource(),
	//"ProductLicense":  ProductLicenseResource(),
	//"ProductServer":   ProductServerResource(),
	//"Region":          RegionResource(),
	//"Self":            SelfResource(),
	//"Server":          ServerResource(),
	//"SIM":             SIMResource(),
	//"SimpleMonitor":   SimpleMonitorResource(),
	//"SSHKey":          SSHKeyResource(),
	//"StartupScript":   StartupScriptResource(),
	//"Switch":          SwitchResource(),
	//"VPCRouter":       VPCRouterResource(),
	//"WebAccel":        WebAccelResource(),
	//"Zone":            ZoneResource(),
	//"Summary":         SummaryResource(),
}
