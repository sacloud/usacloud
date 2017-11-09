package builder

import (
	"fmt"
	"github.com/sacloud/libsacloud/api"
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/libsacloud/sacloud/ostype"
)

/**********************************************************
  Type : ServerBuildEvents
**********************************************************/

// ServerBuildEvents サーバー構築時イベント種別
type ServerBuildEvents int

const (
	// ServerBuildOnStart サーバー構築 開始
	ServerBuildOnStart ServerBuildEvents = iota

	// ServerBuildOnSetPlanBefore サーバープラン設定 開始時
	ServerBuildOnSetPlanBefore

	// ServerBuildOnSetPlanAfter サーバープラン設定 終了時
	ServerBuildOnSetPlanAfter

	// ServerBuildOnCreateServerBefore サーバー作成 開始時
	ServerBuildOnCreateServerBefore

	// ServerBuildOnCreateServerAfter サーバー作成 終了時
	ServerBuildOnCreateServerAfter

	// ServerBuildOnInsertCDROMBefore ISOイメージ挿入 開始時
	ServerBuildOnInsertCDROMBefore

	// ServerBuildOnInsertCDROMAfter ISOイメージ挿入 終了時
	ServerBuildOnInsertCDROMAfter

	// ServerBuildOnBootBefore サーバー起動 開始時
	ServerBuildOnBootBefore

	// ServerBuildOnBootAfter サーバー起動 終了時
	ServerBuildOnBootAfter

	// ServerBuildOnComplete サーバー構築 完了
	ServerBuildOnComplete
)

// ServerBuildEventHandler サーバー構築時イベントハンドラ
type ServerBuildEventHandler func(value *ServerBuildValue, result *ServerBuildResult)

/**********************************************************
  Type : ServerBuilder
**********************************************************/

//serverBuilder サーバービルダー基底
type serverBuilder struct {
	*baseBuilder
	buildEventHandlers map[ServerBuildEvents]ServerBuildEventHandler
	// for server
	serverName      string
	core            int
	memory          int
	interfaceDriver sacloud.EInterfaceDriver
	description     string
	iconID          int64
	tags            []string
	bootAfterCreate bool

	// CDROM
	isoImageID int64

	// privateHost
	privateHostID int64

	// for nic
	nicConnections []string

	// for PacketFilter
	packetFilterIDs []int64

	// for disks
	disk            *DiskBuilder
	additionalDisks []*DiskBuilder

	connectDiskIDs []int64

	currentBuildValue  *ServerBuildValue
	currentBuildResult *ServerBuildResult
}

const (
	// DefaultCore コア数(デフォルト値)
	DefaultCore = 1
	// DefaultMemory メモリサイズ(デフォルト値)
	DefaultMemory = 1
	// DefaultDescription 説明 (デフォルト値)
	DefaultDescription = ""
	// DefaultIconID アイコンID(デフォルト値)
	DefaultIconID = int64(0)
	// DefaultBootAfterCreate サーバー作成後すぐに起動フラグ(デフォルト値)
	DefaultBootAfterCreate = true
)

var (
	// DefaultInterfaceDriver インターフェースドライバ(デフォルト値)
	DefaultInterfaceDriver = sacloud.InterfaceDriverVirtIO
)

func newServerBuilder(client *api.Client, serverName string) *serverBuilder {
	return &serverBuilder{
		baseBuilder: &baseBuilder{
			client: client,
			errors: []error{},
		},
		buildEventHandlers: map[ServerBuildEvents]ServerBuildEventHandler{},
		serverName:         serverName,
		core:               DefaultCore,
		memory:             DefaultMemory,
		interfaceDriver:    DefaultInterfaceDriver,
		description:        DefaultDescription,
		iconID:             DefaultIconID,
		bootAfterCreate:    DefaultBootAfterCreate,
	}

}

/*---------------------------------------------------------
  for connect disk functions
---------------------------------------------------------*/

// ServerDiskless ディスクレスサーバービルダー
func ServerDiskless(client *api.Client, name string) *DisklessServerBuilder {
	b := newServerBuilder(client, name)
	return &DisklessServerBuilder{
		serverBuilder: b,
	}
}

// ServerPublicArchiveUnix ディスクの編集が可能なLinux(Unix)系パブリックアーカイブを利用するビルダー
func ServerPublicArchiveUnix(client *api.Client, os ostype.ArchiveOSTypes, name string, password string) *PublicArchiveUnixServerBuilder {

	b := newServerBuilder(client, name)
	b.ServerPublicArchiveUnix(os, password)
	return &PublicArchiveUnixServerBuilder{
		serverBuilder: b,
	}

}

// ServerPublicArchiveWindows Windows系パブリックアーカイブを利用するビルダー
func ServerPublicArchiveWindows(client *api.Client, os ostype.ArchiveOSTypes, name string) *PublicArchiveWindowsServerBuilder {

	b := newServerBuilder(client, name)
	b.ServerPublicArchiveWindows(os)
	return &PublicArchiveWindowsServerBuilder{
		serverBuilder: b,
	}

}

//ServerBlankDisk 空のディスクを利用するビルダー
func ServerBlankDisk(client *api.Client, name string) *BlankDiskServerBuilder {

	b := newServerBuilder(client, name)
	b.ServerFromBlank()
	return &BlankDiskServerBuilder{
		serverBuilder: b,
	}

}

// ServerFromExistsDisk 既存ディスクを接続するビルダー
func ServerFromExistsDisk(client *api.Client, name string, sourceDiskID int64) *ConnectDiskServerBuilder {
	b := newServerBuilder(client, name)
	b.connectDiskIDs = []int64{sourceDiskID}
	return &ConnectDiskServerBuilder{
		serverBuilder: b,
	}
}

// ServerFromDisk 既存ディスクをコピーして新たなディスクを作成するビルダー
func ServerFromDisk(client *api.Client, name string, sourceDiskID int64) *CommonServerBuilder {
	b := newServerBuilder(client, name)

	b.ServerFromDisk(sourceDiskID)
	return &CommonServerBuilder{
		serverBuilder: b,
	}

}

// ServerFromArchive 既存アーカイブをコピーして新たなディスクを作成するビルダー
func ServerFromArchive(client *api.Client, name string, sourceArchiveID int64) *CommonServerBuilder {
	b := newServerBuilder(client, name)

	b.ServerFromArchive(sourceArchiveID)
	return &CommonServerBuilder{
		serverBuilder: b,
	}

}

/*---------------------------------------------------------
  Inner functions
---------------------------------------------------------*/

func (b *serverBuilder) ServerPublicArchiveUnix(os ostype.ArchiveOSTypes, password string) {
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

func (b *serverBuilder) ServerPublicArchiveWindows(os ostype.ArchiveOSTypes) {
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

func (b *serverBuilder) ServerFromDisk(sourceDiskID int64) {
	b.disk = Disk(b.client, b.serverName)
	b.disk.sourceArchiveID = 0
	b.disk.sourceDiskID = sourceDiskID
}

func (b *serverBuilder) ServerFromArchive(sourceArchiveID int64) {

	b.disk = Disk(b.client, b.serverName)
	b.disk.sourceArchiveID = sourceArchiveID
	b.disk.sourceDiskID = 0
}

func (b *serverBuilder) ServerFromBlank() {
	b.disk = Disk(b.client, b.serverName)
	b.disk.sourceArchiveID = 0
	b.disk.sourceDiskID = 0
}

// Build サーバーの構築
func (b *serverBuilder) Build() (*ServerBuildResult, error) {

	// start
	b.callEventHandlerIfExists(ServerBuildOnStart)
	b.currentBuildValue = &ServerBuildValue{}
	b.currentBuildResult = &ServerBuildResult{}

	if len(b.errors) > 0 {
		return b.currentBuildResult, b.getFlattenErrors()
	}

	// build parameter
	if err := b.buildParams(); err != nil {
		return b.currentBuildResult, err
	}

	// create server
	b.callEventHandlerIfExists(ServerBuildOnCreateServerBefore)
	if err := b.createServer(); err != nil {
		return b.currentBuildResult, err
	}
	b.callEventHandlerIfExists(ServerBuildOnCreateServerAfter)

	// create disks
	if err := b.createDisks(); err != nil {
		return b.currentBuildResult, err
	}

	// connect exists disks
	if err := b.connectDisks(); err != nil {
		return b.currentBuildResult, err
	}

	// insert cdrom
	if b.isoImageID > 0 {
		b.callEventHandlerIfExists(ServerBuildOnInsertCDROMBefore)
		if err := b.insertCDROM(); err != nil {
			return b.currentBuildResult, err
		}
		b.callEventHandlerIfExists(ServerBuildOnInsertCDROMAfter)
	}

	// connect packet filters
	if err := b.connectPacketFilters(); err != nil {
		return b.currentBuildResult, err
	}

	// boot server
	if b.bootAfterCreate {
		b.callEventHandlerIfExists(ServerBuildOnBootBefore)
		if err := b.bootServer(); err != nil {
			return b.currentBuildResult, err
		}
		b.callEventHandlerIfExists(ServerBuildOnBootAfter)
	}

	// complete
	b.callEventHandlerIfExists(ServerBuildOnComplete)
	return b.currentBuildResult, nil
}

func (b *serverBuilder) buildParams() error {

	v := b.currentBuildValue
	v.Server = b.client.Server.New()
	return b.buildServerParams()
}

func (b *serverBuilder) buildServerParams() error {

	v := b.currentBuildValue
	b.callEventHandlerIfExists(ServerBuildOnSetPlanBefore)

	// plan
	plan, err := b.client.Product.Server.GetBySpec(b.core, b.memory)
	if err != nil {
		err = fmt.Errorf("Error building server parameters : setting plan / [%s]", err)
		return err
	}

	b.callEventHandlerIfExists(ServerBuildOnSetPlanAfter)

	s := v.Server
	s.Name = b.serverName
	s.SetServerPlanByID(plan.GetStrID())
	s.Description = b.description
	s.InterfaceDriver = b.interfaceDriver
	// tags
	for _, tag := range b.tags {
		if !s.HasTag(tag) {
			s.AppendTag(tag)
		}
	}
	if b.iconID > 0 {
		s.SetIconByID(b.iconID)
	}

	if b.privateHostID > 0 {
		s.SetPrivateHostByID(b.privateHostID)
	}

	// NIC
	for _, nic := range b.nicConnections {
		switch nic {
		case "shared":
			s.AddPublicNWConnectedParam()
			break
		case "":
			s.AddEmptyConnectedParam()
			break
		default:
			s.AddExistsSwitchConnectedParam(nic)
		}
	}

	return nil
}

func (b *serverBuilder) createDisks() error {
	if b.disk != nil {
		// build disk
		if b.currentBuildResult.Server.ID > 0 {
			b.disk.SetServerID(b.currentBuildResult.Server.ID)
		}
		diskBuildResult, err := b.disk.Build()
		if err != nil {
			return err
		}
		b.currentBuildResult.addDisk(diskBuildResult)
	}
	// build additional disks
	if len(b.additionalDisks) > 0 {
		for _, diskBuilder := range b.additionalDisks {
			if b.currentBuildResult.Server.ID > 0 {
				diskBuilder.SetServerID(b.currentBuildResult.Server.ID)
			}
			res, err := diskBuilder.Build()
			if err != nil {
				return err
			}
			b.currentBuildResult.addDisk(res)
		}
	}
	return nil
}

func (b *serverBuilder) createServer() error {
	server, err := b.client.Server.Create(b.currentBuildValue.Server)
	if err != nil {
		return err
	}
	b.currentBuildResult.Server = server
	return nil
}

func (b *serverBuilder) connectDisks() error {
	server := b.currentBuildResult.Server
	for _, diskID := range b.connectDiskIDs {
		_, err := b.client.Disk.ConnectToServer(diskID, server.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *serverBuilder) connectPacketFilters() error {
	server := b.currentBuildResult.Server
	for i, pfID := range b.packetFilterIDs {
		if len(server.Interfaces) <= i {
			return fmt.Errorf("Number of packet filter and NIC are different")
		}
		if pfID > 0 {
			_, err := b.client.Interface.ConnectToPacketFilter(server.Interfaces[i].ID, pfID)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (b *serverBuilder) insertCDROM() error {
	server := b.currentBuildResult.Server
	_, err := b.client.Server.InsertCDROM(server.ID, b.isoImageID)
	if err != nil {
		return err
	}
	return nil
}

func (b *serverBuilder) bootServer() error {
	server := b.currentBuildResult.Server
	_, err := b.client.Server.Boot(server.ID)
	if err != nil {
		return err
	}

	if err := b.client.Server.SleepUntilUp(server.ID, b.client.DefaultTimeoutDuration); err != nil {
		return err
	}

	// refresh CurrentBildResult.Server
	s, err := b.client.Server.Read(server.ID)
	if err != nil {
		return err
	}
	b.currentBuildResult.Server = s

	return nil
}

func (b *serverBuilder) callEventHandlerIfExists(event ServerBuildEvents) {
	if handler, ok := b.buildEventHandlers[event]; ok {
		handler(b.currentBuildValue, b.currentBuildResult)
	}
}

/**********************************************************
  Type : ServerBuildValue
**********************************************************/

// ServerBuildValue サーバー構築用パラメータ
type ServerBuildValue struct {
	// Server サーバー作成用パラメータ
	Server *sacloud.Server
}

/**********************************************************
  Type : ServerBuildResult
**********************************************************/

// ServerBuildResult サーバー構築結果
type ServerBuildResult struct {
	// Server サーバー
	Server *sacloud.Server
	// Disks ディスク構築結果
	Disks []*DiskBuildResult
}

func (s *ServerBuildResult) addDisk(disk *DiskBuildResult) {
	s.Disks = append(s.Disks, disk)
}
