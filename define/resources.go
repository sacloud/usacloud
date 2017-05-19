// Package define .
//go:generate ../tools/bin/gen-input-models
//go:generate ../tools/bin/gen-cli-commands
//go:generate ../tools/bin/gen-command-funcs
//go:generate ../tools/bin/gen-command-completion
package define

import "github.com/sacloud/usacloud/schema"

var Resources = map[string]*schema.Resource{
	"AuthStatus":      AuthStatusResource(),
	"Archive":         ArchiveResource(),
	"AutoBackup":      AutoBackupResource(),
	"Bill":            BillResource(),
	"Bridge":          BridgeResource(),
	"Database":        DatabaseResource(),
	"Disk":            DiskResource(),
	"DNS":             DNSResource(),
	"GSLB":            GSLBResource(),
	"Icon":            IconResource(),
	"Interface":       InterfaceResource(),
	"Internet":        InternetResource(),
	"ISOImage":        ISOImageResource(),
	"License":         LicenseResource(),
	"LoadBalancer":    LoadBalancerResource(),
	"ObjectStorage":   ObjectStorageResource(),
	"PacketFilter":    PacketFilterResource(),
	"Price":           PriceResource(),
	"ProductDisk":     ProductDiskResource(),
	"ProductInternet": ProductInternetResource(),
	"ProductLicense":  ProductLicenseResource(),
	"ProductServer":   ProductServerResource(),
	"Region":          RegionResource(),
	"Server":          ServerResource(),
	"SimpleMonitor":   SimpleMonitorResource(),
	"SSHKey":          SSHKeyResource(),
	"StartupScript":   StartupScriptResource(),
	"Switch":          SwitchResource(),
	"WebAccel":        WebAccelResource(),
	"Zone":            ZoneResource(),
}
