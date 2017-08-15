package funcs

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
		displayName: "SimpleMonitor",
		finder:      countSimpleMonitor,
		global:      true,
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
	res, err := client.GetServerAPI().Include("ID").Find()
	if err != nil {
		return 0, err
	}
	return len(res.Servers), nil
}

func countDisk(client *api.Client) (int, error) {
	res, err := client.GetDiskAPI().Include("ID").Find()
	if err != nil {
		return 0, err
	}
	return len(res.Disks), nil
}

func countArchive(client *api.Client) (int, error) {
	res, err := client.GetArchiveAPI().WithUserScope().Include("ID").Find()
	if err != nil {
		return 0, err
	}
	return len(res.Archives), nil
}

func countISOImage(client *api.Client) (int, error) {
	res, err := client.GetCDROMAPI().WithUserScope().Include("ID").Find()
	if err != nil {
		return 0, err
	}
	return len(res.CDROMs), nil
}

func countSwitch(client *api.Client) (int, error) {
	res, err := client.GetSwitchAPI().Include("ID").Find()
	if err != nil {
		return 0, err
	}
	return len(res.Switches), nil
}

func countInternet(client *api.Client) (int, error) {
	res, err := client.GetInternetAPI().Include("ID").Find()
	if err != nil {
		return 0, err
	}
	return len(res.Internet), nil
}

func countPacketFilter(client *api.Client) (int, error) {
	res, err := client.GetPacketFilterAPI().Include("ID").Find()
	if err != nil {
		return 0, err
	}
	return len(res.PacketFilters), nil
}

func countBridge(client *api.Client) (int, error) {
	res, err := client.GetBridgeAPI().Include("ID").Find()
	if err != nil {
		return 0, err
	}
	return len(res.Bridges), nil
}

func countAutoBackup(client *api.Client) (int, error) {
	res, err := client.GetAutoBackupAPI().Include("ID").Find()
	if err != nil {
		return 0, err
	}
	return len(res.CommonServiceAutoBackupItems), nil
}

func countLoadBalancer(client *api.Client) (int, error) {
	res, err := client.GetLoadBalancerAPI().Include("ID").Find()
	if err != nil {
		return 0, err
	}
	return len(res.LoadBalancers), nil
}

func countVPCRouter(client *api.Client) (int, error) {
	res, err := client.GetVPCRouterAPI().Include("ID").Find()
	if err != nil {
		return 0, err
	}
	return len(res.VPCRouters), nil
}

func countDatabase(client *api.Client) (int, error) {
	res, err := client.GetDatabaseAPI().Include("ID").Find()
	if err != nil {
		return 0, err
	}
	return len(res.Databases), nil
}

func countGSLB(client *api.Client) (int, error) {
	res, err := client.GetGSLBAPI().Include("ID").Find()
	if err != nil {
		return 0, err
	}
	return len(res.CommonServiceGSLBItems), nil
}

func countDNS(client *api.Client) (int, error) {
	res, err := client.GetDNSAPI().Include("ID").Find()
	if err != nil {
		return 0, err
	}
	return len(res.CommonServiceDNSItems), nil
}

func countSimpleMonitor(client *api.Client) (int, error) {
	res, err := client.GetSimpleMonitorAPI().Include("ID").Find()
	if err != nil {
		return 0, err
	}
	return len(res.SimpleMonitors), nil
}

func countLicense(client *api.Client) (int, error) {
	res, err := client.GetLicenseAPI().Include("ID").Find()
	if err != nil {
		return 0, err
	}
	return len(res.Licenses), nil
}

func countSSHKey(client *api.Client) (int, error) {
	res, err := client.GetSSHKeyAPI().Include("ID").Find()
	if err != nil {
		return 0, err
	}
	return len(res.SSHKeys), nil
}

func countNote(client *api.Client) (int, error) {
	res, err := client.GetNoteAPI().WithUserScope().Include("ID").Find()
	if err != nil {
		return 0, err
	}
	return len(res.Notes), nil
}

func countIcon(client *api.Client) (int, error) {
	res, err := client.GetIconAPI().WithUserScope().Include("ID").Find()
	if err != nil {
		return 0, err
	}
	return len(res.Icons), nil
}
