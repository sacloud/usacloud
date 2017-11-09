package builder

import "github.com/sacloud/libsacloud/sacloud"

// ConnectDiskServerBuilder ブランクディスクを利用して構築を行うサーバービルダー
//
// すでに存在するディスクを持ちます。ディスクの追加に対応していますが、ディスクの修正機能には非対応です。
type ConnectDiskServerBuilder struct {
	*serverBuilder
}

/*---------------------------------------------------------
  common properties
---------------------------------------------------------*/

// GetServerName サーバー名 取得
func (b *ConnectDiskServerBuilder) GetServerName() string {
	return b.serverName
}

// SetServerName サーバー名 設定
func (b *ConnectDiskServerBuilder) SetServerName(serverName string) {
	b.serverName = serverName
}

// WithServerName サーバー名 設定
func (b *ConnectDiskServerBuilder) WithServerName(serverName string) *ConnectDiskServerBuilder {
	b.SetServerName(serverName)
	return b
}

// GetCore CPUコア数 取得
func (b *ConnectDiskServerBuilder) GetCore() int {
	return b.core
}

// SetCore CPUコア数 設定
func (b *ConnectDiskServerBuilder) SetCore(core int) {
	b.core = core
}

// WithCore CPUコア数 設定
func (b *ConnectDiskServerBuilder) WithCore(core int) *ConnectDiskServerBuilder {
	b.SetCore(core)
	return b
}

// GetMemory メモリサイズ(GB単位) 取得
func (b *ConnectDiskServerBuilder) GetMemory() int {
	return b.memory
}

// SetMemory メモリサイズ(GB単位) 設定
func (b *ConnectDiskServerBuilder) SetMemory(memory int) {
	b.memory = memory
}

// WithMemory メモリサイズ(GB単位) 設定
func (b *ConnectDiskServerBuilder) WithMemory(memory int) *ConnectDiskServerBuilder {
	b.SetMemory(memory)
	return b
}

// GetInterfaceDriver インターフェースドライバ 取得
func (b *ConnectDiskServerBuilder) GetInterfaceDriver() sacloud.EInterfaceDriver {
	return b.interfaceDriver
}

// SetInterfaceDriver インターフェースドライバ 設定
func (b *ConnectDiskServerBuilder) SetInterfaceDriver(interfaceDriver sacloud.EInterfaceDriver) {
	b.interfaceDriver = interfaceDriver
}

// WithInterfaceDriver インターフェースドライバ 設定
func (b *ConnectDiskServerBuilder) WithInterfaceDriver(interfaceDriver sacloud.EInterfaceDriver) *ConnectDiskServerBuilder {
	b.SetInterfaceDriver(interfaceDriver)
	return b
}

// GetDescription 説明 取得
func (b *ConnectDiskServerBuilder) GetDescription() string {
	return b.description
}

// SetDescription 説明 設定
func (b *ConnectDiskServerBuilder) SetDescription(description string) {
	b.description = description
}

// WithDescription 説明 設定
func (b *ConnectDiskServerBuilder) WithDescription(description string) *ConnectDiskServerBuilder {
	b.SetDescription(description)
	return b
}

// GetIconID アイコンID 取得
func (b *ConnectDiskServerBuilder) GetIconID() int64 {
	return b.iconID
}

// SetIconID アイコンID 設定
func (b *ConnectDiskServerBuilder) SetIconID(iconID int64) {
	b.iconID = iconID
}

// WithIconID アイコンID 設定
func (b *ConnectDiskServerBuilder) WithIconID(iconID int64) *ConnectDiskServerBuilder {
	b.iconID = iconID
	return b
}

// GetPrivateHostID 専有ホストID 取得
func (b *ConnectDiskServerBuilder) GetPrivateHostID() int64 {
	return b.privateHostID
}

// SetPrivateHostID 専有ホストID 設定
func (b *ConnectDiskServerBuilder) SetPrivateHostID(privateHostID int64) {
	b.privateHostID = privateHostID
}

// WithPrivateHostID 専有ホストID 設定
func (b *ConnectDiskServerBuilder) WithPrivateHostID(privateHostID int64) *ConnectDiskServerBuilder {
	b.privateHostID = privateHostID
	return b
}

// IsBootAfterCreate サーバー作成後すぐに起動フラグ 取得
func (b *ConnectDiskServerBuilder) IsBootAfterCreate() bool {
	return b.bootAfterCreate
}

// SetBootAfterCreate サーバー作成後すぐに起動フラグ 設定
func (b *ConnectDiskServerBuilder) SetBootAfterCreate(bootAfterCreate bool) {
	b.bootAfterCreate = bootAfterCreate
}

// WithBootAfterCreate サーバー作成後すぐに起動フラグ 設定
func (b *ConnectDiskServerBuilder) WithBootAfterCreate(bootAfterCreate bool) *ConnectDiskServerBuilder {
	b.SetBootAfterCreate(bootAfterCreate)
	return b
}

// GetTags タグ 取得
func (b *ConnectDiskServerBuilder) GetTags() []string {
	return b.tags
}

// SetTags タグ 設定
func (b *ConnectDiskServerBuilder) SetTags(tags []string) {
	b.tags = tags
}

// WithTags タグ 設定
func (b *ConnectDiskServerBuilder) WithTags(tags []string) *ConnectDiskServerBuilder {
	b.SetTags(tags)
	return b
}

/*---------------------------------------------------------
  for nic functioms
---------------------------------------------------------*/

// ClearNICConnections NIC接続設定 クリア
func (b *ConnectDiskServerBuilder) ClearNICConnections() *ConnectDiskServerBuilder {
	b.nicConnections = nil
	return b
}

// AddPublicNWConnectedNIC 共有セグメントへの接続追加(注:共有セグメントはeth0のみ接続可能)
func (b *ConnectDiskServerBuilder) AddPublicNWConnectedNIC() {
	b.nicConnections = append(b.nicConnections, "shared")
}

// WithAddPublicNWConnectedNIC 共有セグメントへの接続追加(注:共有セグメントはeth0のみ接続可能)
func (b *ConnectDiskServerBuilder) WithAddPublicNWConnectedNIC() *ConnectDiskServerBuilder {
	b.AddPublicNWConnectedNIC()
	return b
}

// AddExistsSwitchConnectedNIC スイッチ or ルーター+スイッチへの接続追加(注:ルーター+スイッチはeth0のみ接続可能)
func (b *ConnectDiskServerBuilder) AddExistsSwitchConnectedNIC(switchID string) {
	b.nicConnections = append(b.nicConnections, switchID)
}

// WithAddExistsSwitchConnectedNIC スイッチ or ルーター+スイッチへの接続追加(注:ルーター+スイッチはeth0のみ接続可能)
func (b *ConnectDiskServerBuilder) WithAddExistsSwitchConnectedNIC(switchID string) *ConnectDiskServerBuilder {
	b.AddExistsSwitchConnectedNIC(switchID)
	return b
}

// AddDisconnectedNIC 切断されたNIC追加
func (b *ConnectDiskServerBuilder) AddDisconnectedNIC() {
	b.nicConnections = append(b.nicConnections, "")
}

// WithAddDisconnectedNIC 切断されたNIC追加
func (b *ConnectDiskServerBuilder) WithAddDisconnectedNIC() *ConnectDiskServerBuilder {
	b.AddDisconnectedNIC()
	return b
}

// GetPacketFilterIDs パケットフィルタID 取得
func (b *ConnectDiskServerBuilder) GetPacketFilterIDs() []int64 {
	return b.packetFilterIDs
}

// SetPacketFilterIDs パケットフィルタID 設定
func (b *ConnectDiskServerBuilder) SetPacketFilterIDs(ids []int64) {
	b.packetFilterIDs = ids
}

// WithPacketFilterIDs パケットフィルタID 設定
func (b *ConnectDiskServerBuilder) WithPacketFilterIDs(ids []int64) *ConnectDiskServerBuilder {
	b.packetFilterIDs = ids
	return b
}

// GetISOImageID ISOイメージ(CDROM)ID 取得
func (b *ConnectDiskServerBuilder) GetISOImageID() int64 {
	return b.isoImageID
}

// SetISOImageID ISOイメージ(CDROM)ID 設定
func (b *ConnectDiskServerBuilder) SetISOImageID(id int64) {
	b.isoImageID = id
}

// WithISOImageID ISOイメージ(CDROM)ID 設定
func (b *ConnectDiskServerBuilder) WithISOImageID(id int64) *ConnectDiskServerBuilder {
	b.SetISOImageID(id)
	return b
}

/*---------------------------------------------------------
  for additional disks
---------------------------------------------------------*/

// AddAdditionalDisk 追加ディスク 追加
func (b *ConnectDiskServerBuilder) AddAdditionalDisk(diskBuilder *DiskBuilder) {
	b.additionalDisks = append(b.additionalDisks, diskBuilder)
}

// WithAddAdditionalDisk 追加ディスク 追加
func (b *ConnectDiskServerBuilder) WithAddAdditionalDisk(diskBuilder *DiskBuilder) *ConnectDiskServerBuilder {
	b.AddAdditionalDisk(diskBuilder)
	return b
}

// ClearAdditionalDisks 追加ディスク クリア
func (b *ConnectDiskServerBuilder) ClearAdditionalDisks() {
	b.additionalDisks = []*DiskBuilder{}
}

// WithEmptyAdditionalDisks 追加ディスク クリア
func (b *ConnectDiskServerBuilder) WithEmptyAdditionalDisks() *ConnectDiskServerBuilder {
	b.ClearAdditionalDisks()
	return b
}

// GetAdditionalDisks 追加ディスク 取得
func (b *ConnectDiskServerBuilder) GetAdditionalDisks() []*DiskBuilder {
	return b.additionalDisks
}
