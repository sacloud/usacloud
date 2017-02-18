package command

import (
	"fmt"
	"github.com/sacloud/libsacloud/builder"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/libsacloud/sacloud/ostype"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ServerBuild(ctx Context, params *BuildServerParam) error {

	client := ctx.GetAPIClient()

	// TODO validation

	// select builder
	var sb interface{}

	switch params.DiskMode {
	case "create":
		if params.SourceDiskId > 0 {
			sb = builder.ServerFromDisk(client, params.Name, params.SourceDiskId)
		} else if params.SourceArchiveId > 0 {
			sb = builder.ServerFromArchive(client, params.Name, params.SourceArchiveId)
		} else {
			// Windows?
			if isWindows(params.OsType) {
				sb = builder.ServerPublicArchiveWindows(client, strToOSType(params.OsType), params.Name)
			} else {
				sb = builder.ServerPublicArchiveUnix(client, strToOSType(params.OsType), params.Name, params.Password)
			}
		}
	case "connect":
		sb = builder.ServerFromExistsDisk(client, params.Name, params.DiskId)
	case "diskless":
		sb = builder.ServerDiskless(client, params.Name)
	}

	// set network params
	if sb, ok := sb.(serverNetworkParams); ok {
		switch params.NetworkMode {
		case "shared":
			sb.AddPublicNWConnectedNIC()
		case "switch":
			switch sb := sb.(type) {
			case serverConnectSwitchParam:
				sb.AddExistsSwitchConnectedNIC(fmt.Sprintf("%d", params.SwitchId))
			case serverConnectSwitchParamWithEditableDisk:
				sb.AddExistsSwitchConnectedNIC(
					fmt.Sprintf("%d", params.SwitchId),
					params.Ipaddress,
					params.NwMasklen,
					params.DefaultRoute,
				)
			default:
				panic(fmt.Errorf("This server builder Can't connect to switch : %#v", sb))
			}

		case "disconnect":
			sb.AddDisconnectedNIC()
		case "none":
			// noop
		default:
			panic(fmt.Errorf("Unknown NetworkMode : %s", params.NetworkMode))
		}

		sb.SetUseVirtIONetPCI(params.UseNicVirtio)
		sb.SetPacketFilterIDs([]int64{params.PacketFilterId})
	}

	// set disk edit params
	if sb, ok := sb.(serverEditDiskParam); ok {
		sb.SetHostName(params.Hostname)
		sb.SetPassword(params.Password)
		sb.SetDisablePWAuth(params.DisablePasswordAuth)

		for _, v := range params.StartupScriptIds {
			sb.AddNoteID(v)
		}
		for _, v := range params.StartupScripts {
			sb.AddNote(v)
		}
		sb.SetNotesEphemeral(params.StartupScriptsEphemeral)

		// SSH Key generate params
		switch params.SshKeyMode {
		case "id":
			for _, v := range params.SshKeyIds {
				sb.AddSSHKeyID(v)
			}
		case "generate":
			sb.SetGenerateSSHKeyName(params.SshKeyName)
			sb.SetGenerateSSHKeyPassPhrase(params.SshKeyPassPhrase)
			sb.SetGenerateSSHKeyDescription(params.SshKeyDescription)
		case "upload":
			// pubkey(text)
			for _, v := range params.SshKeyPublicKeys {
				sb.AddSSHKey(v)
			}
			// pubkey(from file)
			for _, v := range params.SshKeyPublicKeyFiles {
				sb.AddSSHKey(v)

			}
			sb.SetSSHKeysEphemeral(params.SshKeyEphemeral)
		}

	}

	// set disk params
	if sb, ok := sb.(serverDiskParams); ok {
		sb.SetDiskPlan(params.DiskPlan)
		sb.SetDiskConnection(sacloud.EDiskConnection(params.DiskConnection))
		sb.SetDiskSize(params.DiskSize)
		sb.SetDistantFrom(params.DistantFrom)
	} else {
		panic(fmt.Errorf("ServerCreate is failed: %s", "ServerBuilder not implements disk property."))
	}

	// set common params
	var b serverBuilder
	b, ok := sb.(serverBuilder)
	if !ok {
		panic(fmt.Errorf("ServerCreate is failed: %s", "ServerBuilder not implements common property."))
	}

	tags := params.GetTags()

	b.SetCore(params.GetCore())
	b.SetMemory(params.GetMemory())
	b.SetServerName(params.GetName())
	b.SetDescription(params.GetDescription())
	if params.UsKeyboard {
		tags = append(tags, sacloud.TagKeyboardUS)
	}
	b.SetTags(tags)
	b.SetIconID(params.IconId)
	b.SetBootAfterCreate(!params.DisableBootAfterCreate)
	b.SetISOImageID(params.GetIsoImageId())

	// call Create(id)
	res, err := b.Build()
	if err != nil {
		return fmt.Errorf("ServerCreate is failed: %s", err)
	}

	if len(res.Disks) > 0 && res.Disks[0].GeneratedSSHKey != nil {
		path := params.SshKeyPrivateKeyOutput
		if path == "" {
			p, err := getSSHPrivateKeyStorePath(res.Server.ID)
			if err != nil {
				return fmt.Errorf("ServerCreate is failed: getting HomeDir is failed:%s", err)
			}
			path = p
		}
		pKey := res.Disks[0].GeneratedSSHKey.PrivateKey
		dir := filepath.Dir(path)
		if err := os.MkdirAll(dir, 0600); err != nil {
			return fmt.Errorf("ServerCreate is failed: creating directory(%s) is failed:%s", dir, err)
		}

		err = ioutil.WriteFile(path, []byte(pKey), os.FileMode(0600))
		if err != nil {
			return fmt.Errorf("ServerCreate is failed: Writing private key to %s is failed:%s", params.SshKeyPrivateKeyOutput, err)
		}
	}

	return ctx.GetOutput().Print(res)
}

func isWindows(osType string) bool {
	windowsTypes := []string{
		"windows2008", "windows2008-rds", "windows2008-rds-office",
		"windows2012", "windows2012-rds", "windows2012-rds-office",
		"windows2016",
	}
	for _, v := range windowsTypes {
		if v == osType {
			return true
		}
	}
	return false
}

func strToOSType(osType string) ostype.ArchiveOSTypes {
	switch osType {
	case "centos":
		return ostype.CentOS
	case "ubuntu":
		return ostype.Ubuntu
	case "debian":
		return ostype.Debian
	case "vyos":
		return ostype.VyOS
	case "coreos":
		return ostype.CoreOS
	case "kusanagi":
		return ostype.Kusanagi
	case "site-guard":
		return ostype.SiteGuard
	case "freebsd":
		return ostype.FreeBSD
	case "windows2008":
		return ostype.Windows2008
	case "windows2008-rds":
		return ostype.Windows2008RDS
	case "windows2008-rds-office":
		return ostype.Windows2008RDSOffice
	case "windows2012":
		return ostype.Windows2012
	case "windows2012-rds":
		return ostype.Windows2012RDS
	case "windows2012-rds-office":
		return ostype.Windows2012RDSOffice
	case "windows2016":
		return ostype.Windows2016
	default:
		return ostype.Custom
	}
}

type serverBuilder interface {
	SetCore(int)
	SetMemory(int)
	SetServerName(string)
	SetDescription(string)
	SetTags([]string)
	SetIconID(int64)
	SetBootAfterCreate(bool)
	SetISOImageID(int64)

	Build() (*builder.ServerBuildResult, error)
}

type serverDiskParams interface {
	SetDiskPlan(string)
	SetDiskConnection(sacloud.EDiskConnection)
	SetDiskSize(int)
	SetDistantFrom([]int64)
}

type serverNetworkParams interface {
	SetUseVirtIONetPCI(bool)
	SetPacketFilterIDs([]int64)
	AddPublicNWConnectedNIC()
	AddDisconnectedNIC()
}

type serverConnectSwitchParamWithEditableDisk interface {
	AddExistsSwitchConnectedNIC(id string, ipaddress string, maskLen int, defRoute string)
}

type serverConnectSwitchParam interface {
	AddExistsSwitchConnectedNIC(id string)
}

type serverEditDiskParam interface {
	SetHostName(string)
	SetPassword(string)
	SetDisablePWAuth(bool)
	AddNote(string)
	AddNoteID(int64)
	SetNotesEphemeral(bool)
	AddSSHKey(string)
	AddSSHKeyID(int64)
	SetSSHKeysEphemeral(bool)
	SetGenerateSSHKeyName(string)
	SetGenerateSSHKeyPassPhrase(string)
	SetGenerateSSHKeyDescription(string)
}
