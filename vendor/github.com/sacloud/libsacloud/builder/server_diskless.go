package builder

import "github.com/sacloud/libsacloud/sacloud"

// DisklessServerBuilder ディスクレス サーバービルダー
//
// ディスクレスのサーバーを構築します。 ディスク関連の設定に非対応です。
type DisklessServerBuilder struct {
	*serverBuilder
}

/*---------------------------------------------------------
  common properties
---------------------------------------------------------*/

// GetServerName サーバー名 取得
func (b *DisklessServerBuilder) GetServerName() string {
	return b.serverName
}

// SetServerName サーバー名 設定
func (b *DisklessServerBuilder) SetServerName(serverName string) {
	b.serverName = serverName
}

// WithServerName サーバー名 設定
func (b *DisklessServerBuilder) WithServerName(serverName string) *DisklessServerBuilder {
	b.SetServerName(serverName)
	return b
}

// GetCore CPUコア数 取得
func (b *DisklessServerBuilder) GetCore() int {
	return b.core
}

// SetCore CPUコア数 設定
func (b *DisklessServerBuilder) SetCore(core int) {
	b.core = core
}

// WithCore CPUコア数 設定
func (b *DisklessServerBuilder) WithCore(core int) *DisklessServerBuilder {
	b.SetCore(core)
	return b
}

// GetMemory メモリサイズ(GB単位) 取得
func (b *DisklessServerBuilder) GetMemory() int {
	return b.memory
}

// SetMemory メモリサイズ(GB単位) 設定
func (b *DisklessServerBuilder) SetMemory(memory int) {
	b.memory = memory
}

// WithMemory メモリサイズ(GB単位) 設定
func (b *DisklessServerBuilder) WithMemory(memory int) *DisklessServerBuilder {
	b.SetMemory(memory)
	return b
}

// GetInterfaceDriver インターフェースドライバ 取得
func (b *DisklessServerBuilder) GetInterfaceDriver() sacloud.EInterfaceDriver {
	return b.interfaceDriver
}

// SetInterfaceDriver インターフェースドライバ 設定
func (b *DisklessServerBuilder) SetInterfaceDriver(interfaceDriver sacloud.EInterfaceDriver) {
	b.interfaceDriver = interfaceDriver
}

// WithInterfaceDriver インターフェースドライバ 設定
func (b *DisklessServerBuilder) WithInterfaceDriver(interfaceDriver sacloud.EInterfaceDriver) *DisklessServerBuilder {
	b.SetInterfaceDriver(interfaceDriver)
	return b
}

// GetDescription 説明 取得
func (b *DisklessServerBuilder) GetDescription() string {
	return b.description
}

// SetDescription 説明 設定
func (b *DisklessServerBuilder) SetDescription(description string) {
	b.description = description
}

// WithDescription 説明 設定
func (b *DisklessServerBuilder) WithDescription(description string) *DisklessServerBuilder {
	b.SetDescription(description)
	return b
}

// GetIconID アイコンID 取得
func (b *DisklessServerBuilder) GetIconID() int64 {
	return b.iconID
}

// SetIconID アイコンID 設定
func (b *DisklessServerBuilder) SetIconID(iconID int64) {
	b.iconID = iconID
}

// WithIconID アイコンID 設定
func (b *DisklessServerBuilder) WithIconID(iconID int64) *DisklessServerBuilder {
	b.SetIconID(iconID)
	return b
}

// GetPrivateHostID 専有ホストID 取得
func (b *DisklessServerBuilder) GetPrivateHostID() int64 {
	return b.privateHostID
}

// SetPrivateHostID 専有ホストID 設定
func (b *DisklessServerBuilder) SetPrivateHostID(privateHostID int64) {
	b.privateHostID = privateHostID
}

// WithPrivateHostID 専有ホストID 設定
func (b *DisklessServerBuilder) WithPrivateHostID(privateHostID int64) *DisklessServerBuilder {
	b.privateHostID = privateHostID
	return b
}

// IsBootAfterCreate サーバー作成後すぐに起動フラグ 取得
func (b *DisklessServerBuilder) IsBootAfterCreate() bool {
	return b.bootAfterCreate
}

// SetBootAfterCreate サーバー作成後すぐに起動フラグ 設定
func (b *DisklessServerBuilder) SetBootAfterCreate(bootAfterCreate bool) {
	b.bootAfterCreate = bootAfterCreate
}

// WithBootAfterCreate サーバー作成後すぐに起動フラグ 設定
func (b *DisklessServerBuilder) WithBootAfterCreate(bootAfterCreate bool) *DisklessServerBuilder {
	b.SetBootAfterCreate(bootAfterCreate)
	return b
}

// GetTags タグ 取得
func (b *DisklessServerBuilder) GetTags() []string {
	return b.tags
}

// SetTags タグ 設定
func (b *DisklessServerBuilder) SetTags(tags []string) {
	b.tags = tags
}

// WithTags タグ 設定
func (b *DisklessServerBuilder) WithTags(tags []string) *DisklessServerBuilder {
	b.SetTags(tags)
	return b
}

/*---------------------------------------------------------
  for nic functioms
---------------------------------------------------------*/

// ClearNICConnections NIC接続設定 クリア
func (b *DisklessServerBuilder) ClearNICConnections() {
	b.nicConnections = nil
}

// WithEmptyNICConnections NIC接続設定 クリア
func (b *DisklessServerBuilder) WithEmptyNICConnections() *DisklessServerBuilder {
	b.ClearNICConnections()
	return b
}

// AddPublicNWConnectedNIC 共有セグメントへの接続追加(注:共有セグメントはeth0のみ接続可能)
func (b *DisklessServerBuilder) AddPublicNWConnectedNIC() {
	b.nicConnections = append(b.nicConnections, "shared")
}

// WithAddPublicNWConnectedNIC 共有セグメントへの接続追加(注:共有セグメントはeth0のみ接続可能)
func (b *DisklessServerBuilder) WithAddPublicNWConnectedNIC() *DisklessServerBuilder {
	b.AddPublicNWConnectedNIC()
	return b
}

// AddExistsSwitchConnectedNIC スイッチ or ルーター+スイッチへの接続追加(注:ルーター+スイッチはeth0のみ接続可能)
func (b *DisklessServerBuilder) AddExistsSwitchConnectedNIC(switchID string) {
	b.nicConnections = append(b.nicConnections, switchID)
}

// WithAddExistsSwitchConnectedNIC スイッチ or ルーター+スイッチへの接続追加(注:ルーター+スイッチはeth0のみ接続可能)
func (b *DisklessServerBuilder) WithAddExistsSwitchConnectedNIC(switchID string) *DisklessServerBuilder {
	b.AddExistsSwitchConnectedNIC(switchID)
	return b
}

// AddDisconnectedNIC 切断されたNIC追加
func (b *DisklessServerBuilder) AddDisconnectedNIC() {
	b.nicConnections = append(b.nicConnections, "")
}

// WithAddDisconnectedNIC 切断されたNIC追加
func (b *DisklessServerBuilder) WithAddDisconnectedNIC() *DisklessServerBuilder {
	b.AddDisconnectedNIC()
	return b
}

// GetPacketFilterIDs パケットフィルタID 取得
func (b *DisklessServerBuilder) GetPacketFilterIDs() []int64 {
	return b.packetFilterIDs
}

// SetPacketFilterIDs パケットフィルタID 設定
func (b *DisklessServerBuilder) SetPacketFilterIDs(ids []int64) {
	b.packetFilterIDs = ids
}

// WithPacketFilterIDs パケットフィルタID 設定
func (b *DisklessServerBuilder) WithPacketFilterIDs(ids []int64) *DisklessServerBuilder {
	b.packetFilterIDs = ids
	return b
}

// GetISOImageID ISOイメージ(CDROM)ID 取得
func (b *DisklessServerBuilder) GetISOImageID() int64 {
	return b.isoImageID
}

// SetISOImageID ISOイメージ(CDROM)ID 設定
func (b *DisklessServerBuilder) SetISOImageID(id int64) {
	b.isoImageID = id
}

// WithISOImageID ISOイメージ(CDROM)ID 設定
func (b *DisklessServerBuilder) WithISOImageID(id int64) *DisklessServerBuilder {
	b.isoImageID = id
	return b
}

/*---------------------------------------------------------
  for event handler
---------------------------------------------------------*/

// SetEventHandler イベントハンドラ 設定
func (b *DisklessServerBuilder) SetEventHandler(event ServerBuildEvents, handler ServerBuildEventHandler) {
	b.buildEventHandlers[event] = handler
}

// WithEventHandler イベントハンドラ 設定
func (b *DisklessServerBuilder) WithEventHandler(event ServerBuildEvents, handler ServerBuildEventHandler) *DisklessServerBuilder {
	b.SetEventHandler(event, handler)
	return b
}

// ClearEventHandler イベントハンドラ クリア
func (b *DisklessServerBuilder) ClearEventHandler(event ServerBuildEvents) {
	delete(b.buildEventHandlers, event)
}

// WithEmptyEventHandler イベントハンドラ クリア
func (b *DisklessServerBuilder) WithEmptyEventHandler(event ServerBuildEvents) *DisklessServerBuilder {
	b.ClearEventHandler(event)
	return b
}

// GetEventHandler イベントハンドラ 取得
func (b *DisklessServerBuilder) GetEventHandler(event ServerBuildEvents) *ServerBuildEventHandler {
	if handler, ok := b.buildEventHandlers[event]; ok {
		return &handler
	}
	return nil
}
