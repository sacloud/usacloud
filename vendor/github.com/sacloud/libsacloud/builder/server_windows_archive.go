package builder

import "github.com/sacloud/libsacloud/sacloud"

// PublicArchiveWindowsServerBuilder Windows系パブリックアーカイブを利用して構築を行うサーバービルダー
//
// 基本的なディスク設定やディスクの追加に対応していますが、ディスクの修正機能には非対応です。
type PublicArchiveWindowsServerBuilder struct {
	*serverBuilder
}

/*---------------------------------------------------------
  common properties
---------------------------------------------------------*/

// GetServerName サーバー名 取得
func (b *PublicArchiveWindowsServerBuilder) GetServerName() string {
	return b.serverName
}

// SetServerName サーバー名 設定
func (b *PublicArchiveWindowsServerBuilder) SetServerName(serverName string) {
	b.serverName = serverName
}

// WithServerName サーバー名 設定
func (b *PublicArchiveWindowsServerBuilder) WithServerName(serverName string) *PublicArchiveWindowsServerBuilder {
	b.SetServerName(serverName)
	return b
}

// GetCore CPUコア数 取得
func (b *PublicArchiveWindowsServerBuilder) GetCore() int {
	return b.core
}

// SetCore CPUコア数 設定
func (b *PublicArchiveWindowsServerBuilder) SetCore(core int) {
	b.core = core
}

// WithCore CPUコア数 設定
func (b *PublicArchiveWindowsServerBuilder) WithCore(core int) *PublicArchiveWindowsServerBuilder {
	b.SetCore(core)
	return b
}

// GetMemory メモリサイズ(GB単位) 取得
func (b *PublicArchiveWindowsServerBuilder) GetMemory() int {
	return b.memory
}

// SetMemory メモリサイズ(GB単位) 設定
func (b *PublicArchiveWindowsServerBuilder) SetMemory(memory int) {
	b.memory = memory
}

// WithMemory メモリサイズ(GB単位) 設定
func (b *PublicArchiveWindowsServerBuilder) WithMemory(memory int) *PublicArchiveWindowsServerBuilder {
	b.SetMemory(memory)
	return b
}

// GetInterfaceDriver インターフェースドライバ 取得
func (b *PublicArchiveWindowsServerBuilder) GetInterfaceDriver() sacloud.EInterfaceDriver {
	return b.interfaceDriver
}

// SetInterfaceDriver インターフェースドライバ 設定
func (b *PublicArchiveWindowsServerBuilder) SetInterfaceDriver(interfaceDriver sacloud.EInterfaceDriver) {
	b.interfaceDriver = interfaceDriver
}

// WithInterfaceDriver インターフェースドライバ 設定
func (b *PublicArchiveWindowsServerBuilder) WithInterfaceDriver(interfaceDriver sacloud.EInterfaceDriver) *PublicArchiveWindowsServerBuilder {
	b.SetInterfaceDriver(interfaceDriver)
	return b
}

// GetDescription 説明 取得
func (b *PublicArchiveWindowsServerBuilder) GetDescription() string {
	return b.description
}

// SetDescription 説明 設定
func (b *PublicArchiveWindowsServerBuilder) SetDescription(description string) {
	b.description = description
}

// WithDescription 説明 設定
func (b *PublicArchiveWindowsServerBuilder) WithDescription(description string) *PublicArchiveWindowsServerBuilder {
	b.SetDescription(description)
	return b
}

// GetIconID アイコンID 取得
func (b *PublicArchiveWindowsServerBuilder) GetIconID() int64 {
	return b.iconID
}

// SetIconID アイコンID 設定
func (b *PublicArchiveWindowsServerBuilder) SetIconID(iconID int64) {
	b.iconID = iconID
}

// WithIconID アイコンID 設定
func (b *PublicArchiveWindowsServerBuilder) WithIconID(iconID int64) *PublicArchiveWindowsServerBuilder {
	b.SetIconID(iconID)
	return b
}

// GetPrivateHostID 専有ID 取得
func (b *PublicArchiveWindowsServerBuilder) GetPrivateHostID() int64 {
	return b.privateHostID
}

// SetPrivateHostID 専有ホストID 設定
func (b *PublicArchiveWindowsServerBuilder) SetPrivateHostID(privateHostID int64) {
	b.privateHostID = privateHostID
}

// WithPrivateHostID 専有ホストID 設定
func (b *PublicArchiveWindowsServerBuilder) WithPrivateHostID(privateHostID int64) *PublicArchiveWindowsServerBuilder {
	b.privateHostID = privateHostID
	return b
}

// IsBootAfterCreate サーバー作成後すぐに起動フラグ 取得
func (b *PublicArchiveWindowsServerBuilder) IsBootAfterCreate() bool {
	return b.bootAfterCreate
}

// SetBootAfterCreate サーバー作成後すぐに起動フラグ 設定
func (b *PublicArchiveWindowsServerBuilder) SetBootAfterCreate(bootAfterCreate bool) {
	b.bootAfterCreate = bootAfterCreate
}

// WithBootAfterCreate サーバー作成後すぐに起動フラグ 設定
func (b *PublicArchiveWindowsServerBuilder) WithBootAfterCreate(bootAfterCreate bool) *PublicArchiveWindowsServerBuilder {
	b.SetBootAfterCreate(bootAfterCreate)
	return b
}

// GetTags タグ 取得
func (b *PublicArchiveWindowsServerBuilder) GetTags() []string {
	return b.tags
}

// SetTags タグ 設定
func (b *PublicArchiveWindowsServerBuilder) SetTags(tags []string) {
	b.tags = tags
}

// WithTags タグ 設定
func (b *PublicArchiveWindowsServerBuilder) WithTags(tags []string) *PublicArchiveWindowsServerBuilder {
	b.SetTags(tags)
	return b
}

/*---------------------------------------------------------
  for nic functioms
---------------------------------------------------------*/

// ClearNICConnections NIC接続設定 クリア
func (b *PublicArchiveWindowsServerBuilder) ClearNICConnections() {
	b.nicConnections = nil
}

// WithEmptyNICConnections NIC接続設定 クリア
func (b *PublicArchiveWindowsServerBuilder) WithEmptyNICConnections() *PublicArchiveWindowsServerBuilder {
	b.nicConnections = nil
	return b
}

// AddPublicNWConnectedNIC 共有セグメントへの接続追加(注:共有セグメントはeth0のみ接続可能)
func (b *PublicArchiveWindowsServerBuilder) AddPublicNWConnectedNIC() {
	b.nicConnections = append(b.nicConnections, "shared")
}

// WithAddPublicNWConnectedNIC 共有セグメントへの接続追加(注:共有セグメントはeth0のみ接続可能)
func (b *PublicArchiveWindowsServerBuilder) WithAddPublicNWConnectedNIC() *PublicArchiveWindowsServerBuilder {
	b.AddPublicNWConnectedNIC()
	return b
}

// AddExistsSwitchConnectedNIC スイッチ or ルーター+スイッチへの接続追加(注:ルーター+スイッチはeth0のみ接続可能)
func (b *PublicArchiveWindowsServerBuilder) AddExistsSwitchConnectedNIC(switchID string, ipaddress string, networkMaskLen int, defaultRoute string) {
	b.nicConnections = append(b.nicConnections, switchID)
	b.disk.SetIPAddress(ipaddress)
	b.disk.SetNetworkMaskLen(networkMaskLen)
	b.disk.SetDefaultRoute(defaultRoute)
}

// WithAddExistsSwitchConnectedNIC スイッチ or ルーター+スイッチへの接続追加(注:ルーター+スイッチはeth0のみ接続可能)
func (b *PublicArchiveWindowsServerBuilder) WithAddExistsSwitchConnectedNIC(switchID string, ipaddress string, networkMaskLen int, defaultRoute string) *PublicArchiveWindowsServerBuilder {
	b.AddExistsSwitchConnectedNIC(switchID, ipaddress, networkMaskLen, defaultRoute)
	return b
}

// AddDisconnectedNIC 切断されたNIC追加
func (b *PublicArchiveWindowsServerBuilder) AddDisconnectedNIC() {
	b.nicConnections = append(b.nicConnections, "")
}

// WithAddDisconnectedNIC 切断されたNIC追加
func (b *PublicArchiveWindowsServerBuilder) WithAddDisconnectedNIC() *PublicArchiveWindowsServerBuilder {
	b.nicConnections = append(b.nicConnections, "")
	return b
}

// GetPacketFilterIDs パケットフィルタID 取得
func (b *PublicArchiveWindowsServerBuilder) GetPacketFilterIDs() []int64 {
	return b.packetFilterIDs
}

// SetPacketFilterIDs パケットフィルタID 設定
func (b *PublicArchiveWindowsServerBuilder) SetPacketFilterIDs(ids []int64) {
	b.packetFilterIDs = ids
}

// WithPacketFilterIDs パケットフィルタID 設定
func (b *PublicArchiveWindowsServerBuilder) WithPacketFilterIDs(ids []int64) *PublicArchiveWindowsServerBuilder {
	b.packetFilterIDs = ids
	return b
}

// GetISOImageID ISOイメージ(CDROM)ID 取得
func (b *PublicArchiveWindowsServerBuilder) GetISOImageID() int64 {
	return b.isoImageID
}

// SetISOImageID ISOイメージ(CDROM)ID 設定
func (b *PublicArchiveWindowsServerBuilder) SetISOImageID(id int64) {
	b.isoImageID = id
}

// WithISOImageID ISOイメージ(CDROM)ID 設定
func (b *PublicArchiveWindowsServerBuilder) WithISOImageID(id int64) *PublicArchiveWindowsServerBuilder {
	b.SetISOImageID(id)
	return b
}

/*---------------------------------------------------------
  for disk properties
---------------------------------------------------------*/

// GetDiskSize ディスクサイズ(GB単位) 取得
func (b *PublicArchiveWindowsServerBuilder) GetDiskSize() int {
	return b.disk.GetSize()
}

// SetDiskSize ディスクサイズ(GB単位) 設定
func (b *PublicArchiveWindowsServerBuilder) SetDiskSize(diskSize int) {
	b.disk.SetSize(diskSize)
}

// WithDiskSize ディスクサイズ(GB単位) 設定
func (b *PublicArchiveWindowsServerBuilder) WithDiskSize(diskSize int) *PublicArchiveWindowsServerBuilder {
	b.SetDiskSize(diskSize)
	return b
}

// GetDistantFrom ストレージ隔離対象ディスク 取得
func (b *PublicArchiveWindowsServerBuilder) GetDistantFrom() []int64 {
	return b.disk.GetDistantFrom()
}

// SetDistantFrom ストレージ隔離対象ディスク 設定
func (b *PublicArchiveWindowsServerBuilder) SetDistantFrom(distantFrom []int64) {
	b.disk.SetDistantFrom(distantFrom)
}

// WithDistantFrom ストレージ隔離対象ディスク 設定
func (b *PublicArchiveWindowsServerBuilder) WithDistantFrom(distantFrom []int64) *PublicArchiveWindowsServerBuilder {
	b.SetDistantFrom(distantFrom)
	return b
}

// AddDistantFrom ストレージ隔離対象ディスク 追加
func (b *PublicArchiveWindowsServerBuilder) AddDistantFrom(diskID int64) {
	b.disk.AddDistantFrom(diskID)
}

// WithAddDistantFrom ストレージ隔離対象ディスク 追加
func (b *PublicArchiveWindowsServerBuilder) WithAddDistantFrom(diskID int64) *PublicArchiveWindowsServerBuilder {
	b.AddDistantFrom(diskID)
	return b
}

// ClearDistantFrom ストレージ隔離対象ディスク クリア
func (b *PublicArchiveWindowsServerBuilder) ClearDistantFrom() {
	b.disk.ClearDistantFrom()
}

// WithEmptyDistantFrom ストレージ隔離対象ディスク クリア
func (b *PublicArchiveWindowsServerBuilder) WithEmptyDistantFrom() *PublicArchiveWindowsServerBuilder {
	b.ClearDistantFrom()
	return b
}

// GetDiskPlanID ディスクプラン(SSD/HDD) 取得
func (b *PublicArchiveWindowsServerBuilder) GetDiskPlanID() sacloud.DiskPlanID {
	return b.disk.GetPlanID()
}

// SetDiskPlanID ディスクプラン(SSD/HDD) 設定
func (b *PublicArchiveWindowsServerBuilder) SetDiskPlanID(diskPlanID sacloud.DiskPlanID) {
	b.disk.SetPlanID(diskPlanID)
}

// WithDiskPlanID ディスクプラン(SSD/HDD) 設定
func (b *PublicArchiveWindowsServerBuilder) WithDiskPlanID(diskPlanID sacloud.DiskPlanID) *PublicArchiveWindowsServerBuilder {
	b.SetDiskPlanID(diskPlanID)
	return b
}

// SetDiskPlan ディスクプラン(ssd/hdd) 設定(文字列から)
func (b *PublicArchiveWindowsServerBuilder) SetDiskPlan(plan string) {
	b.disk.SetPlan(plan)
}

// WithDiskPlan ディスクプラン(ssd/hdd) 設定(文字列から)
func (b *PublicArchiveWindowsServerBuilder) WithDiskPlan(plan string) *PublicArchiveWindowsServerBuilder {
	b.SetDiskPlan(plan)
	return b
}

// GetDiskConnection ディスク接続方法(VirtIO/IDE) 取得
func (b *PublicArchiveWindowsServerBuilder) GetDiskConnection() sacloud.EDiskConnection {
	return b.disk.GetConnection()
}

// SetDiskConnection ディスク接続方法(VirtIO/IDE) 設定
func (b *PublicArchiveWindowsServerBuilder) SetDiskConnection(diskConnection sacloud.EDiskConnection) {
	b.disk.SetConnection(diskConnection)
}

// WithDiskConnection ディスク接続方法(VirtIO/IDE) 設定
func (b *PublicArchiveWindowsServerBuilder) WithDiskConnection(diskConnection sacloud.EDiskConnection) *PublicArchiveWindowsServerBuilder {
	b.SetDiskConnection(diskConnection)
	return b
}

/*---------------------------------------------------------
  for disk edit properties
---------------------------------------------------------*/

// GetSourceArchiveID ソースアーカイブID 取得
func (b *PublicArchiveWindowsServerBuilder) GetSourceArchiveID() int64 {
	return b.disk.GetSourceArchiveID()
}

// GetSourceDiskID ソースディスクID 設定
func (b *PublicArchiveWindowsServerBuilder) GetSourceDiskID() int64 {
	return b.disk.GetSourceDiskID()
}

/*---------------------------------------------------------
  for event handler
---------------------------------------------------------*/

// SetEventHandler イベントハンドラ 設定
func (b *PublicArchiveWindowsServerBuilder) SetEventHandler(event ServerBuildEvents, handler ServerBuildEventHandler) {
	b.buildEventHandlers[event] = handler
}

// WithEventHandler イベントハンドラ 設定
func (b *PublicArchiveWindowsServerBuilder) WithEventHandler(event ServerBuildEvents, handler ServerBuildEventHandler) *PublicArchiveWindowsServerBuilder {
	b.SetEventHandler(event, handler)
	return b
}

// ClearEventHandler イベントハンドラ クリア
func (b *PublicArchiveWindowsServerBuilder) ClearEventHandler(event ServerBuildEvents) {
	delete(b.buildEventHandlers, event)
}

// WithEmptyEventHandler イベントハンドラ クリア
func (b *PublicArchiveWindowsServerBuilder) WithEmptyEventHandler(event ServerBuildEvents) *PublicArchiveWindowsServerBuilder {
	b.ClearEventHandler(event)
	return b
}

// GetEventHandler イベントハンドラ 取得
func (b *PublicArchiveWindowsServerBuilder) GetEventHandler(event ServerBuildEvents) *ServerBuildEventHandler {
	if handler, ok := b.buildEventHandlers[event]; ok {
		return &handler
	}
	return nil
}

// SetDiskEventHandler ディスクイベントハンドラ 設定
func (b *PublicArchiveWindowsServerBuilder) SetDiskEventHandler(event DiskBuildEvents, handler DiskBuildEventHandler) {
	b.disk.SetEventHandler(event, handler)
}

// WithDiskEventHandler ディスクイベントハンドラ 設定
func (b *PublicArchiveWindowsServerBuilder) WithDiskEventHandler(event DiskBuildEvents, handler DiskBuildEventHandler) *PublicArchiveWindowsServerBuilder {
	b.SetDiskEventHandler(event, handler)
	return b
}

// ClearDiskEventHandler ディスクイベントハンドラ クリア
func (b *PublicArchiveWindowsServerBuilder) ClearDiskEventHandler(event DiskBuildEvents) *PublicArchiveWindowsServerBuilder {
	b.disk.ClearEventHandler(event)
	return b
}

// WithEmptyDiskEventHandler ディスクイベントハンドラ クリア
func (b *PublicArchiveWindowsServerBuilder) WithEmptyDiskEventHandler(event DiskBuildEvents) *PublicArchiveWindowsServerBuilder {
	b.ClearDiskEventHandler(event)
	return b
}

// GetDiskEventHandler ディスクイベントハンドラ 取得
func (b *PublicArchiveWindowsServerBuilder) GetDiskEventHandler(event DiskBuildEvents) *DiskBuildEventHandler {
	return b.disk.GetEventHandler(event)
}

/*---------------------------------------------------------
  for additional disks
---------------------------------------------------------*/

// AddAdditionalDisk 追加ディスク 追加
func (b *PublicArchiveWindowsServerBuilder) AddAdditionalDisk(diskBuilder *DiskBuilder) {
	b.additionalDisks = append(b.additionalDisks, diskBuilder)
}

// WithAddAdditionalDisk 追加ディスク 追加
func (b *PublicArchiveWindowsServerBuilder) WithAddAdditionalDisk(diskBuilder *DiskBuilder) *PublicArchiveWindowsServerBuilder {
	b.additionalDisks = append(b.additionalDisks, diskBuilder)
	return b
}

// ClearAdditionalDisks 追加ディスク クリア
func (b *PublicArchiveWindowsServerBuilder) ClearAdditionalDisks() {
	b.additionalDisks = []*DiskBuilder{}
}

// WithEmptyAdditionalDisks 追加ディスク クリア
func (b *PublicArchiveWindowsServerBuilder) WithEmptyAdditionalDisks() *PublicArchiveWindowsServerBuilder {
	b.ClearAdditionalDisks()
	return b
}

// GetAdditionalDisks 追加ディスク 取得
func (b *PublicArchiveWindowsServerBuilder) GetAdditionalDisks() []*DiskBuilder {
	return b.additionalDisks
}
