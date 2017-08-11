// Package define .
//go:generate go run ../tools/gen-input-models/main.go
//go:generate go run ../tools/gen-cli-commands/main.go
//go:generate go run ../tools/gen-command-funcs/main.go
//go:generate go run ../tools/gen-command-completion/main.go
package define

import "github.com/sacloud/usacloud/schema"

var Resources = map[string]*schema.Resource{
	"AuthStatus":      AuthStatusResource(),
	"Archive":         ArchiveResource(),
	"AutoBackup":      AutoBackupResource(),
	"Bill":            BillResource(),
	"Bridge":          BridgeResource(),
	"Config":          ConfigResource(),
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
	"VPCRouter":       VPCRouterResource(),
	"WebAccel":        WebAccelResource(),
	"Zone":            ZoneResource(),
}
