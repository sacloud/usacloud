package builder

import (
	"github.com/sacloud/libsacloud/sacloud"
)

// BlankDiskServerBuilder ブランクディスクを利用して構築を行うサーバービルダー
//
// 空のディスクを持ちます。基本的なディスク設定やディスクの追加に対応していますが、ディスクの修正機能には非対応です。
type BlankDiskServerBuilder struct {
	*serverBuilder
}

/*---------------------------------------------------------
  common properties
---------------------------------------------------------*/

// GetServerName サーバー名 取得
func (b *BlankDiskServerBuilder) GetServerName() string {
	return b.serverName
}

// SetServerName サーバー名 設定
func (b *BlankDiskServerBuilder) SetServerName(serverName string) {
	b.serverName = serverName
}

// WithServerName サーバー名 設定
func (b *BlankDiskServerBuilder) WithServerName(serverName string) *BlankDiskServerBuilder {
	b.SetServerName(serverName)
	return b
}

// GetCore CPUコア数 取得
func (b *BlankDiskServerBuilder) GetCore() int {
	return b.core
}

// SetCore CPUコア数 設定
func (b *BlankDiskServerBuilder) SetCore(core int) {
	b.core = core
}

// WithCore CPUコア数 設定
func (b *BlankDiskServerBuilder) WithCore(core int) *BlankDiskServerBuilder {
	b.SetCore(core)
	return b
}

// GetMemory メモリサイズ(GB単位) 取得
func (b *BlankDiskServerBuilder) GetMemory() int {
	return b.memory
}

// SetMemory メモリサイズ(GB単位) 設定
func (b *BlankDiskServerBuilder) SetMemory(memory int) {
	b.memory = memory
}

// WithMemory メモリサイズ(GB単位) 設定
func (b *BlankDiskServerBuilder) WithMemory(memory int) *BlankDiskServerBuilder {
	b.SetMemory(memory)
	return b
}

// GetInterfaceDriver インターフェースドライバ 取得
func (b *BlankDiskServerBuilder) GetInterfaceDriver() sacloud.EInterfaceDriver {
	return b.interfaceDriver
}

// SetInterfaceDriver インターフェースドライバ 設定
func (b *BlankDiskServerBuilder) SetInterfaceDriver(interfaceDriver sacloud.EInterfaceDriver) {
	b.interfaceDriver = interfaceDriver
}

// WithInterfaceDriver インターフェースドライバ 設定
func (b *BlankDiskServerBuilder) WithInterfaceDriver(interfaceDriver sacloud.EInterfaceDriver) *BlankDiskServerBuilder {
	b.SetInterfaceDriver(interfaceDriver)
	return b
}

// GetDescription 説明 取得
func (b *BlankDiskServerBuilder) GetDescription() string {
	return b.description
}

// SetDescription 説明 設定
func (b *BlankDiskServerBuilder) SetDescription(description string) {
	b.description = description
}

// WithDescription 説明 設定
func (b *BlankDiskServerBuilder) WithDescription(description string) *BlankDiskServerBuilder {
	b.SetDescription(description)
	return b
}

// GetIconID アイコンID 取得
func (b *BlankDiskServerBuilder) GetIconID() int64 {
	return b.iconID
}

// SetIconID アイコンID 設定
func (b *BlankDiskServerBuilder) SetIconID(iconID int64) {
	b.iconID = iconID
}

// WithIconID アイコンID 設定
func (b *BlankDiskServerBuilder) WithIconID(iconID int64) *BlankDiskServerBuilder {
	b.iconID = iconID
	return b
}

// GetPrivateHostID アイコンID 取得
func (b *BlankDiskServerBuilder) GetPrivateHostID() int64 {
	return b.privateHostID
}

// SetPrivateHostID アイコンID 設定
func (b *BlankDiskServerBuilder) SetPrivateHostID(privateHostID int64) {
	b.privateHostID = privateHostID
}

// WithPrivateHostID アイコンID 設定
func (b *BlankDiskServerBuilder) WithPrivateHostID(privateHostID int64) *BlankDiskServerBuilder {
	b.privateHostID = privateHostID
	return b
}

// IsBootAfterCreate サーバー作成後すぐに起動フラグ 取得
func (b *BlankDiskServerBuilder) IsBootAfterCreate() bool {
	return b.bootAfterCreate
}

// SetBootAfterCreate サーバー作成後すぐに起動フラグ 設定
func (b *BlankDiskServerBuilder) SetBootAfterCreate(bootAfterCreate bool) {
	b.bootAfterCreate = bootAfterCreate
}

// WithBootAfterCreate サーバー作成後すぐに起動フラグ 設定
func (b *BlankDiskServerBuilder) WithBootAfterCreate(bootAfterCreate bool) *BlankDiskServerBuilder {
	b.SetBootAfterCreate(bootAfterCreate)
	return b
}

// GetTags タグ 取得
func (b *BlankDiskServerBuilder) GetTags() []string {
	return b.tags
}

// SetTags タグ 設定
func (b *BlankDiskServerBuilder) SetTags(tags []string) {
	b.tags = tags
}

// WithTags タグ 設定
func (b *BlankDiskServerBuilder) WithTags(tags []string) *BlankDiskServerBuilder {
	b.SetTags(tags)
	return b
}

/*---------------------------------------------------------
  for nic functioms
---------------------------------------------------------*/

// ClearNICConnections NIC接続設定 クリア
func (b *BlankDiskServerBuilder) ClearNICConnections() *BlankDiskServerBuilder {
	b.nicConnections = nil
	return b
}

// AddPublicNWConnectedNIC 共有セグメントへの接続追加(注:共有セグメントはeth0のみ接続可能)
func (b *BlankDiskServerBuilder) AddPublicNWConnectedNIC() {
	b.nicConnections = append(b.nicConnections, "shared")
}

// WithAddPublicNWConnectedNIC 共有セグメントへの接続追加(注:共有セグメントはeth0のみ接続可能)
func (b *BlankDiskServerBuilder) WithAddPublicNWConnectedNIC() *BlankDiskServerBuilder {
	b.AddPublicNWConnectedNIC()
	return b
}

// AddExistsSwitchConnectedNIC スイッチ or ルーター+スイッチへの接続追加(注:ルーター+スイッチはeth0のみ接続可能)
func (b *BlankDiskServerBuilder) AddExistsSwitchConnectedNIC(switchID string) {
	b.nicConnections = append(b.nicConnections, switchID)
}

// WithAddExistsSwitchConnectedNIC スイッチ or ルーター+スイッチへの接続追加(注:ルーター+スイッチはeth0のみ接続可能)
func (b *BlankDiskServerBuilder) WithAddExistsSwitchConnectedNIC(switchID string) *BlankDiskServerBuilder {
	b.AddExistsSwitchConnectedNIC(switchID)
	return b
}

// AddDisconnectedNIC 切断されたNIC追加
func (b *BlankDiskServerBuilder) AddDisconnectedNIC() {
	b.nicConnections = append(b.nicConnections, "")
}

// WithAddDisconnectedNIC 切断されたNIC追加
func (b *BlankDiskServerBuilder) WithAddDisconnectedNIC() *BlankDiskServerBuilder {
	b.AddDisconnectedNIC()
	return b
}

// GetPacketFilterIDs パケットフィルタID 取得
func (b *BlankDiskServerBuilder) GetPacketFilterIDs() []int64 {
	return b.packetFilterIDs
}

// SetPacketFilterIDs パケットフィルタID 設定
func (b *BlankDiskServerBuilder) SetPacketFilterIDs(ids []int64) {
	b.packetFilterIDs = ids
}

// WithPacketFilterIDs パケットフィルタID 設定
func (b *BlankDiskServerBuilder) WithPacketFilterIDs(ids []int64) *BlankDiskServerBuilder {
	b.packetFilterIDs = ids
	return b
}

// GetISOImageID ISOイメージ(CDROM)ID 取得
func (b *BlankDiskServerBuilder) GetISOImageID() int64 {
	return b.isoImageID
}

// SetISOImageID ISOイメージ(CDROM)ID 設定
func (b *BlankDiskServerBuilder) SetISOImageID(id int64) {
	b.isoImageID = id
}

// WithISOImageID ISOイメージ(CDROM)ID 設定
func (b *BlankDiskServerBuilder) WithISOImageID(id int64) *BlankDiskServerBuilder {
	b.SetISOImageID(id)
	return b
}

/*---------------------------------------------------------
  for disk properties
---------------------------------------------------------*/

// GetDiskSize ディスクサイズ(GB単位) 取得
func (b *BlankDiskServerBuilder) GetDiskSize() int {
	return b.disk.GetSize()
}

// SetDiskSize ディスクサイズ(GB単位) 設定
func (b *BlankDiskServerBuilder) SetDiskSize(diskSize int) {
	b.disk.SetSize(diskSize)
}

// WithDiskSize ディスクサイズ(GB単位) 設定
func (b *BlankDiskServerBuilder) WithDiskSize(diskSize int) *BlankDiskServerBuilder {
	b.SetDiskSize(diskSize)
	return b
}

// GetDistantFrom ストレージ隔離対象ディスク 取得
func (b *BlankDiskServerBuilder) GetDistantFrom() []int64 {
	return b.disk.GetDistantFrom()
}

// SetDistantFrom ストレージ隔離対象ディスク 設定
func (b *BlankDiskServerBuilder) SetDistantFrom(distantFrom []int64) {
	b.disk.SetDistantFrom(distantFrom)
}

// WithDistantFrom ストレージ隔離対象ディスク 設定
func (b *BlankDiskServerBuilder) WithDistantFrom(distantFrom []int64) *BlankDiskServerBuilder {
	b.SetDistantFrom(distantFrom)
	return b
}

// AddDistantFrom ストレージ隔離対象ディスク 追加
func (b *BlankDiskServerBuilder) AddDistantFrom(diskID int64) {
	b.disk.AddDistantFrom(diskID)
}

// WithAddDistantFrom ストレージ隔離対象ディスク 追加
func (b *BlankDiskServerBuilder) WithAddDistantFrom(diskID int64) *BlankDiskServerBuilder {
	b.AddDistantFrom(diskID)
	return b
}

// ClearDistantFrom ストレージ隔離対象ディスク クリア
func (b *BlankDiskServerBuilder) ClearDistantFrom() {
	b.disk.ClearDistantFrom()
}

// WithEmptyDistantFrom ストレージ隔離対象ディスク クリア
func (b *BlankDiskServerBuilder) WithEmptyDistantFrom() *BlankDiskServerBuilder {
	b.ClearDistantFrom()
	return b
}

// GetDiskPlanID ディスクプラン(SSD/HDD) 取得
func (b *BlankDiskServerBuilder) GetDiskPlanID() sacloud.DiskPlanID {
	return b.disk.GetPlanID()
}

// SetDiskPlanID ディスクプラン(SSD/HDD) 設定
func (b *BlankDiskServerBuilder) SetDiskPlanID(diskPlanID sacloud.DiskPlanID) {
	b.disk.SetPlanID(diskPlanID)
}

// WithDiskPlanID ディスクプラン(SSD/HDD) 設定
func (b *BlankDiskServerBuilder) WithDiskPlanID(diskPlanID sacloud.DiskPlanID) *BlankDiskServerBuilder {
	b.SetDiskPlanID(diskPlanID)
	return b
}

// SetDiskPlan ディスクプラン(ssd/hdd) 設定(文字列から)
func (b *BlankDiskServerBuilder) SetDiskPlan(plan string) {
	b.disk.SetPlan(plan)
}

// WithDiskPlan ディスクプラン(ssd/hdd) 設定(文字列から)
func (b *BlankDiskServerBuilder) WithDiskPlan(plan string) *BlankDiskServerBuilder {
	b.SetDiskPlan(plan)
	return b
}

// GetDiskConnection ディスク接続方法(VirtIO/IDE) 取得
func (b *BlankDiskServerBuilder) GetDiskConnection() sacloud.EDiskConnection {
	return b.disk.GetConnection()
}

// SetDiskConnection ディスク接続方法(VirtIO/IDE) 設定
func (b *BlankDiskServerBuilder) SetDiskConnection(diskConnection sacloud.EDiskConnection) {
	b.disk.SetConnection(diskConnection)
}

// WithDiskConnection ディスク接続方法(VirtIO/IDE) 設定
func (b *BlankDiskServerBuilder) WithDiskConnection(diskConnection sacloud.EDiskConnection) *BlankDiskServerBuilder {
	b.SetDiskConnection(diskConnection)
	return b
}

/*---------------------------------------------------------
  for event handler
---------------------------------------------------------*/

// SetEventHandler イベントハンドラ 設定
func (b *BlankDiskServerBuilder) SetEventHandler(event ServerBuildEvents, handler ServerBuildEventHandler) {
	b.buildEventHandlers[event] = handler
}

// WithEventHandler イベントハンドラ 設定
func (b *BlankDiskServerBuilder) WithEventHandler(event ServerBuildEvents, handler ServerBuildEventHandler) *BlankDiskServerBuilder {
	b.SetEventHandler(event, handler)
	return b
}

// ClearEventHandler イベントハンドラ クリア
func (b *BlankDiskServerBuilder) ClearEventHandler(event ServerBuildEvents) {
	delete(b.buildEventHandlers, event)
}

// WithEmptyEventHandler イベントハンドラ クリア
func (b *BlankDiskServerBuilder) WithEmptyEventHandler(event ServerBuildEvents) *BlankDiskServerBuilder {
	delete(b.buildEventHandlers, event)
	return b
}

// GetEventHandler イベントハンドラ 取得
func (b *BlankDiskServerBuilder) GetEventHandler(event ServerBuildEvents) *ServerBuildEventHandler {
	if handler, ok := b.buildEventHandlers[event]; ok {
		return &handler
	}
	return nil
}

// SetDiskEventHandler ディスクイベントハンドラ 設定
func (b *BlankDiskServerBuilder) SetDiskEventHandler(event DiskBuildEvents, handler DiskBuildEventHandler) {
	b.disk.SetEventHandler(event, handler)
}

// WithDiskEventHandler ディスクイベントハンドラ 設定
func (b *BlankDiskServerBuilder) WithDiskEventHandler(event DiskBuildEvents, handler DiskBuildEventHandler) *BlankDiskServerBuilder {
	b.SetDiskEventHandler(event, handler)
	return b
}

// ClearDiskEventHandler ディスクイベントハンドラ クリア
func (b *BlankDiskServerBuilder) ClearDiskEventHandler(event DiskBuildEvents) {
	b.disk.ClearEventHandler(event)
}

// WithEmptyDiskEventHandler ディスクイベントハンドラ クリア
func (b *BlankDiskServerBuilder) WithEmptyDiskEventHandler(event DiskBuildEvents) *BlankDiskServerBuilder {
	b.ClearDiskEventHandler(event)
	return b
}

// GetDiskEventHandler ディスクイベントハンドラ 取得
func (b *BlankDiskServerBuilder) GetDiskEventHandler(event DiskBuildEvents) *DiskBuildEventHandler {
	return b.disk.GetEventHandler(event)
}

/*---------------------------------------------------------
  for additional disks
---------------------------------------------------------*/

// AddAdditionalDisk 追加ディスク 追加
func (b *BlankDiskServerBuilder) AddAdditionalDisk(diskBuilder *DiskBuilder) {
	b.additionalDisks = append(b.additionalDisks, diskBuilder)
}

// WithAddAdditionalDisk 追加ディスク 追加
func (b *BlankDiskServerBuilder) WithAddAdditionalDisk(diskBuilder *DiskBuilder) *BlankDiskServerBuilder {
	b.AddAdditionalDisk(diskBuilder)
	return b
}

// ClearAdditionalDisks 追加ディスク クリア
func (b *BlankDiskServerBuilder) ClearAdditionalDisks() {
	b.additionalDisks = []*DiskBuilder{}
}

// WithEmptyAdditionalDisks 追加ディスク クリア
func (b *BlankDiskServerBuilder) WithEmptyAdditionalDisks() *BlankDiskServerBuilder {
	b.ClearAdditionalDisks()
	return b
}

// GetAdditionalDisks 追加ディスク 取得
func (b *BlankDiskServerBuilder) GetAdditionalDisks() []*DiskBuilder {
	return b.additionalDisks
}
