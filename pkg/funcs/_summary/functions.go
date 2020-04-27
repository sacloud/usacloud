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

package summary

import "github.com/sacloud/libsacloud/api"

type resourceCounter struct {
	displayName string
	finder      func(client *api.Client) (resCount int, err error)
	global      bool
	paid        bool
}

var resourceCounters = []*resourceCounter{
	{
		displayName: "Server",
		finder:      countServer,
		global:      false,
		paid:        true,
	},
	{
		displayName: "Disk",
		finder:      countDisk,
		global:      false,
		paid:        true,
	},
	{
		displayName: "Archive",
		finder:      countArchive,
		global:      false,
		paid:        true,
	},
	{
		displayName: "ISOImage",
		finder:      countISOImage,
		global:      false,
		paid:        true,
	},
	{
		displayName: "Switch",
		finder:      countSwitch,
		global:      false,
		paid:        true,
	},
	{
		displayName: "Internet",
		finder:      countInternet,
		global:      false,
		paid:        true,
	},
	{
		displayName: "PacketFilter",
		finder:      countPacketFilter,
		global:      false,
		paid:        false,
	},
	{
		displayName: "Bridge",
		finder:      countBridge,
		global:      true,
		paid:        true,
	},
	{
		displayName: "AutoBackup",
		finder:      countAutoBackup,
		global:      false,
		paid:        false,
	},
	{
		displayName: "LoadBalancer",
		finder:      countLoadBalancer,
		global:      false,
		paid:        true,
	},
	{
		displayName: "VPCRouter",
		finder:      countVPCRouter,
		global:      false,
		paid:        true,
	},
	{
		displayName: "Database",
		finder:      countDatabase,
		global:      false,
		paid:        true,
	},
	{
		displayName: "NFS",
		finder:      countNFS,
		global:      false,
		paid:        true,
	},
	{
		displayName: "ProxyLB",
		finder:      countProxyLB,
		global:      true,
		paid:        false,
	},
	{
		displayName: "GSLB",
		finder:      countGSLB,
		global:      true,
		paid:        true,
	},
	{
		displayName: "DNS",
		finder:      countDNS,
		global:      true,
		paid:        true,
	},
	{
		displayName: "SIMpleMonitor",
		finder:      countSIMpleMonitor,
		global:      true,
		paid:        true,
	},
	{
		displayName: "SIM",
		finder:      countSIM,
		global:      true,
		paid:        true,
	},
	{
		displayName: "MobileGateway",
		finder:      countMobileGateway,
		global:      false,
		paid:        true,
	},
	{
		displayName: "License",
		finder:      countLicense,
		global:      true,
		paid:        true,
	},
	{
		displayName: "SSHKey",
		finder:      countSSHKey,
		global:      true,
		paid:        false,
	},
	{
		displayName: "StartupScript",
		finder:      countNote,
		global:      true,
		paid:        false,
	},
	{
		displayName: "Icon",
		finder:      countIcon,
		global:      true,
		paid:        false,
	},
}

func countServer(client *api.Client) (int, error) {
	res, err := client.GetServerAPI().Include("ID").Limit(1).Find()
	if err != nil {
		return 0, err
	}
	return int(res.Total), nil
}

func countDisk(client *api.Client) (int, error) {
	res, err := client.GetDiskAPI().Include("ID").Limit(1).Find()
	if err != nil {
		return 0, err
	}
	return int(res.Total), nil
}

func countArchive(client *api.Client) (int, error) {
	res, err := client.GetArchiveAPI().WithUserScope().Include("ID").Limit(1).Find()
	if err != nil {
		return 0, err
	}
	return int(res.Total), nil
}

func countISOImage(client *api.Client) (int, error) {
	res, err := client.GetCDROMAPI().WithUserScope().Include("ID").Limit(1).Find()
	if err != nil {
		return 0, err
	}
	return int(res.Total), nil
}

func countSwitch(client *api.Client) (int, error) {
	res, err := client.GetSwitchAPI().Include("ID").Limit(1).Find()
	if err != nil {
		return 0, err
	}
	return int(res.Total), nil
}

func countInternet(client *api.Client) (int, error) {
	res, err := client.GetInternetAPI().Include("ID").Limit(1).Find()
	if err != nil {
		return 0, err
	}
	return int(res.Total), nil
}

func countPacketFilter(client *api.Client) (int, error) {
	res, err := client.GetPacketFilterAPI().Include("ID").Limit(1).Find()
	if err != nil {
		return 0, err
	}
	return int(res.Total), nil
}

func countBridge(client *api.Client) (int, error) {
	res, err := client.GetBridgeAPI().Include("ID").Limit(1).Find()
	if err != nil {
		return 0, err
	}
	return int(res.Total), nil
}

func countAutoBackup(client *api.Client) (int, error) {
	res, err := client.GetAutoBackupAPI().Include("ID").Limit(1).Find()
	if err != nil {
		return 0, err
	}
	return int(res.Total), nil
}

func countLoadBalancer(client *api.Client) (int, error) {
	res, err := client.GetLoadBalancerAPI().Include("ID").Limit(1).Find()
	if err != nil {
		return 0, err
	}
	return int(res.Total), nil
}

func countVPCRouter(client *api.Client) (int, error) {
	res, err := client.GetVPCRouterAPI().Include("ID").Limit(1).Find()
	if err != nil {
		return 0, err
	}
	return int(res.Total), nil
}

func countDatabase(client *api.Client) (int, error) {
	res, err := client.GetDatabaseAPI().Include("ID").Limit(1).Find()
	if err != nil {
		return 0, err
	}
	return int(res.Total), nil
}

func countNFS(client *api.Client) (int, error) {
	res, err := client.GetNFSAPI().Include("ID").Limit(1).Find()
	if err != nil {
		return 0, err
	}
	return int(res.Total), nil
}

func countProxyLB(client *api.Client) (int, error) {
	res, err := client.GetProxyLBAPI().Include("ID").Limit(1).Find()
	if err != nil {
		return 0, err
	}
	return int(res.Total), nil
}

func countGSLB(client *api.Client) (int, error) {
	res, err := client.GetGSLBAPI().Include("ID").Limit(1).Find()
	if err != nil {
		return 0, err
	}
	return int(res.Total), nil
}

func countDNS(client *api.Client) (int, error) {
	res, err := client.GetDNSAPI().Include("ID").Limit(1).Find()
	if err != nil {
		return 0, err
	}
	return int(res.Total), nil
}

func countSIMpleMonitor(client *api.Client) (int, error) {
	res, err := client.GetSimpleMonitorAPI().Include("ID").Limit(1).Find()
	if err != nil {
		return 0, err
	}
	return int(res.Total), nil
}

func countSIM(client *api.Client) (int, error) {
	res, err := client.GetSIMAPI().Include("ID").Limit(1).Find()
	if err != nil {
		return 0, err
	}
	return int(res.Total), nil
}

func countMobileGateway(client *api.Client) (int, error) {
	res, err := client.GetMobileGatewayAPI().Include("ID").Limit(1).Find()
	if err != nil {
		return 0, err
	}
	return int(res.Total), nil
}

func countLicense(client *api.Client) (int, error) {
	res, err := client.GetLicenseAPI().Include("ID").Limit(1).Find()
	if err != nil {
		return 0, err
	}
	return int(res.Total), nil
}

func countSSHKey(client *api.Client) (int, error) {
	res, err := client.GetSSHKeyAPI().Include("ID").Limit(1).Find()
	if err != nil {
		return 0, err
	}
	return int(res.Total), nil
}

func countNote(client *api.Client) (int, error) {
	res, err := client.GetNoteAPI().WithUserScope().Include("ID").Limit(1).Find()
	if err != nil {
		return 0, err
	}
	return int(res.Total), nil
}

func countIcon(client *api.Client) (int, error) {
	res, err := client.GetIconAPI().WithUserScope().Include("ID").Limit(1).Find()
	if err != nil {
		return 0, err
	}
	return int(res.Total), nil
}
