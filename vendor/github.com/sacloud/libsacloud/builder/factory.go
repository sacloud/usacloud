package builder

import (
	"fmt"

	"github.com/sacloud/libsacloud/api"
	"github.com/sacloud/libsacloud/sacloud/ostype"
)

// ServerDiskless ディスクレスサーバービルダー
func ServerDiskless(client *api.Client, name string) DisklessServerBuilder {
	b := newServerBuilder(client, name)
	b.hasCommonProperty = true
	b.hasNetworkInterfaceProperty = true
	b.hasServerEventProperty = true
	return b
}

// ServerPublicArchiveUnix ディスクの編集が可能なLinux(Unix)系パブリックアーカイブを利用するビルダー
func ServerPublicArchiveUnix(client *api.Client, os ostype.ArchiveOSTypes, name string, password string) PublicArchiveUnixServerBuilder {
	b := newServerBuilder(client, name)
	b.hasCommonProperty = true
	b.hasNetworkInterfaceProperty = true
	b.hasDiskProperty = true
	b.hasDiskSourceProperty = true
	b.hasDiskEditProperty = true
	b.hasServerEventProperty = true
	b.hasDiskEventProperty = true
	b.hasAdditionalDiskProperty = true

	b.serverPublicArchiveUnix(os, password)
	return b
}

// ServerPublicArchiveFixedUnix ディスクの編集が不可なLinux(Unix)系パブリックアーカイブを利用するビルダー
func ServerPublicArchiveFixedUnix(client *api.Client, os ostype.ArchiveOSTypes, name string) FixedUnixArchiveServerBuilder {
	b := newServerBuilder(client, name)
	b.hasCommonProperty = true
	b.hasNetworkInterfaceProperty = true
	b.hasDiskProperty = true
	b.hasDiskSourceProperty = true
	b.hasServerEventProperty = true
	b.hasDiskEventProperty = true
	b.hasAdditionalDiskProperty = true

	b.serverPublicArchiveFixedUnix(os)
	return b
}

// ServerPublicArchiveWindows Windows系パブリックアーカイブを利用するビルダー
func ServerPublicArchiveWindows(client *api.Client, os ostype.ArchiveOSTypes, name string) PublicArchiveWindowsServerBuilder {
	b := newServerBuilder(client, name)
	b.hasCommonProperty = true
	b.hasNetworkInterfaceProperty = true
	b.hasDiskProperty = true
	b.hasDiskSourceProperty = true
	b.hasServerEventProperty = true
	b.hasDiskEventProperty = true
	b.hasAdditionalDiskProperty = true

	b.serverPublicArchiveWindows(os)
	return b
}

//ServerBlankDisk 空のディスクを利用するビルダー
func ServerBlankDisk(client *api.Client, name string) BlankDiskServerBuilder {
	b := newServerBuilder(client, name)
	b.hasCommonProperty = true
	b.hasNetworkInterfaceProperty = true
	b.hasDiskProperty = true
	b.hasServerEventProperty = true
	b.hasDiskEventProperty = true
	b.hasAdditionalDiskProperty = true

	b.serverFromBlank()
	return b
}

// ServerFromExistsDisk 既存ディスクを接続するビルダー
func ServerFromExistsDisk(client *api.Client, name string, sourceDiskID int64) ConnectDiskServerBuilder {
	b := newServerBuilder(client, name)
	b.hasCommonProperty = true
	b.hasNetworkInterfaceProperty = true
	b.hasAdditionalDiskProperty = true

	b.connectDiskIDs = []int64{sourceDiskID}
	return b
}

// ServerFromDisk 既存ディスクをコピーして新たなディスクを作成するビルダー
func ServerFromDisk(client *api.Client, name string, sourceDiskID int64) CommonServerBuilder {
	b := newServerBuilder(client, name)
	b.hasCommonProperty = true
	b.hasNetworkInterfaceProperty = true
	b.hasDiskProperty = true
	b.hasDiskSourceProperty = true
	b.hasDiskEditProperty = true
	b.hasServerEventProperty = true
	b.hasDiskEventProperty = true
	b.hasAdditionalDiskProperty = true

	b.serverFromDisk(sourceDiskID)
	return b
}

// ServerFromArchive 既存アーカイブをコピーして新たなディスクを作成するビルダー
func ServerFromArchive(client *api.Client, name string, sourceArchiveID int64) CommonServerBuilder {
	b := newServerBuilder(client, name)
	b.hasCommonProperty = true
	b.hasNetworkInterfaceProperty = true
	b.hasDiskProperty = true
	b.hasDiskSourceProperty = true
	b.hasDiskEditProperty = true
	b.hasServerEventProperty = true
	b.hasDiskEventProperty = true
	b.hasAdditionalDiskProperty = true

	b.serverFromArchive(sourceArchiveID)
	return b
}

func (b *serverBuilder) serverPublicArchiveUnix(os ostype.ArchiveOSTypes, password string) {
	if !os.IsSupportDiskEdit() {
		b.errors = append(b.errors, fmt.Errorf("%q is not support EditDisk", os))
	}

	archive, err := b.client.Archive.FindByOSType(os)
	if err != nil {
		b.errors = append(b.errors, err)
	}

	b.disk = Disk(b.client, b.serverName)
	b.disk.sourceArchiveID = archive.ID
	b.disk.password = password

}

func (b *serverBuilder) serverPublicArchiveFixedUnix(os ostype.ArchiveOSTypes) {
	archive, err := b.client.Archive.FindByOSType(os)
	if err != nil {
		b.errors = append(b.errors, err)
	}

	b.disk = Disk(b.client, b.serverName)
	b.disk.sourceArchiveID = archive.ID
}

func (b *serverBuilder) serverPublicArchiveWindows(os ostype.ArchiveOSTypes) {
	if !os.IsWindows() {
		b.errors = append(b.errors, fmt.Errorf("%q is not windows", os))
	}

	archive, err := b.client.Archive.FindByOSType(os)
	if err != nil {
		b.errors = append(b.errors, err)
	}

	b.disk = Disk(b.client, b.serverName)
	b.disk.sourceArchiveID = archive.ID
	b.disk.sourceDiskID = 0
	b.disk.forceEditDisk = true
}

func (b *serverBuilder) serverFromDisk(sourceDiskID int64) {
	b.disk = Disk(b.client, b.serverName)
	b.disk.sourceArchiveID = 0
	b.disk.sourceDiskID = sourceDiskID
}

func (b *serverBuilder) serverFromArchive(sourceArchiveID int64) {

	b.disk = Disk(b.client, b.serverName)
	b.disk.sourceArchiveID = sourceArchiveID
	b.disk.sourceDiskID = 0
}

func (b *serverBuilder) serverFromBlank() {
	b.disk = Disk(b.client, b.serverName)
	b.disk.sourceArchiveID = 0
	b.disk.sourceDiskID = 0
}
