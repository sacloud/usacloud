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
//go:generate go run ../tools/gen-input-models/main.go
//go:generate go run ../tools/gen-cli-commands/main.go
//go:generate go run ../tools/gen-command-funcs/main.go
//go:generate go run ../tools/gen-input-v2-models/main.go
//go:generate go run ../tools/gen-cli-v2-commands/main.go
//go:generate go run ../tools/gen-cli-v2-resource-finder/main.go
//go:generate go run ../tools/gen-cli-v2-usage/main.go
//go:generate go run ../tools/gen-cli-v2-root-command/main.go
package define

import "github.com/sacloud/usacloud/schema"

var Resources = map[string]*schema.Resource{
	"AuthStatus":      AuthStatusResource(),
	"Archive":         ArchiveResource(),
	"AutoBackup":      AutoBackupResource(),
	"Bill":            BillResource(),
	"Bridge":          BridgeResource(),
	"Config":          ConfigResource(),
	"Coupon":          CouponResource(),
	"Database":        DatabaseResource(),
	"Disk":            DiskResource(),
	"DNS":             DNSResource(),
	"GSLB":            GSLBResource(),
	"Icon":            IconResource(),
	"Interface":       InterfaceResource(),
	"Internet":        InternetResource(),
	"IPv4":            IPv4Resource(),
	"IPv6":            IPv6Resource(),
	"ISOImage":        ISOImageResource(),
	"License":         LicenseResource(),
	"LoadBalancer":    LoadBalancerResource(),
	"MobileGateway":   MobileGatewayResource(),
	"NFS":             NFSResource(),
	"ObjectStorage":   ObjectStorageResource(),
	"PacketFilter":    PacketFilterResource(),
	"ProxyLB":         ProxyLBResource(),
	"Price":           PriceResource(),
	"PrivateHost":     PrivateHostResource(),
	"ProductDisk":     ProductDiskResource(),
	"ProductInternet": ProductInternetResource(),
	"ProductLicense":  ProductLicenseResource(),
	"ProductServer":   ProductServerResource(),
	"Region":          RegionResource(),
	"Self":            SelfResource(),
	"Server":          ServerResource(),
	"SIM":             SIMResource(),
	"SimpleMonitor":   SimpleMonitorResource(),
	"SSHKey":          SSHKeyResource(),
	"StartupScript":   StartupScriptResource(),
	"Switch":          SwitchResource(),
	"VPCRouter":       VPCRouterResource(),
	"WebAccel":        WebAccelResource(),
	"Zone":            ZoneResource(),
	"Summary":         SummaryResource(),
}
