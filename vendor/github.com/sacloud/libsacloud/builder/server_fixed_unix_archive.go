package builder

import "github.com/sacloud/libsacloud/sacloud"

// FixedUnixArchiveServerBuilder ディスクの修正不可なUNIX系パブリックアーカイブを利用して構築を行うサーバービルダー
//
// 基本的なディスク設定やディスクの追加に対応していますが、ディスクの修正機能には非対応です。
type FixedUnixArchiveServerBuilder struct {
	*serverBuilder
}

/*---------------------------------------------------------
  common properties
---------------------------------------------------------*/

// GetServerName サーバー名 取得
func (b *FixedUnixArchiveServerBuilder) GetServerName() string {
	return b.serverName
}

// SetServerName サーバー名 設定
func (b *FixedUnixArchiveServerBuilder) SetServerName(serverName string) {
	b.serverName = serverName
}

// WithServerName サーバー名 設定
func (b *FixedUnixArchiveServerBuilder) WithServerName(serverName string) *FixedUnixArchiveServerBuilder {
	b.SetServerName(serverName)
	return b
}

// GetCore CPUコア数 取得
func (b *FixedUnixArchiveServerBuilder) GetCore() int {
	return b.core
}

// SetCore CPUコア数 設定
func (b *FixedUnixArchiveServerBuilder) SetCore(core int) {
	b.core = core
}

// WithCore CPUコア数 設定
func (b *FixedUnixArchiveServerBuilder) WithCore(core int) *FixedUnixArchiveServerBuilder {
	b.SetCore(core)
	return b
}

// GetMemory メモリサイズ(GB単位) 取得
func (b *FixedUnixArchiveServerBuilder) GetMemory() int {
	return b.memory
}

// SetMemory メモリサイズ(GB単位) 設定
func (b *FixedUnixArchiveServerBuilder) SetMemory(memory int) {
	b.memory = memory
}

// WithMemory メモリサイズ(GB単位) 設定
func (b *FixedUnixArchiveServerBuilder) WithMemory(memory int) *FixedUnixArchiveServerBuilder {
	b.SetMemory(memory)
	return b
}

// GetInterfaceDriver インターフェースドライバ 取得
func (b *FixedUnixArchiveServerBuilder) GetInterfaceDriver() sacloud.EInterfaceDriver {
	return b.interfaceDriver
}

// SetInterfaceDriver インターフェースドライバ 設定
func (b *FixedUnixArchiveServerBuilder) SetInterfaceDriver(interfaceDriver sacloud.EInterfaceDriver) {
	b.interfaceDriver = interfaceDriver
}

// WithInterfaceDriver インターフェースドライバ 設定
func (b *FixedUnixArchiveServerBuilder) WithInterfaceDriver(interfaceDriver sacloud.EInterfaceDriver) *FixedUnixArchiveServerBuilder {
	b.SetInterfaceDriver(interfaceDriver)
	return b
}

// GetDescription 説明 取得
func (b *FixedUnixArchiveServerBuilder) GetDescription() string {
	return b.description
}

// SetDescription 説明 設定
func (b *FixedUnixArchiveServerBuilder) SetDescription(description string) {
	b.description = description
}

// WithDescription 説明 設定
func (b *FixedUnixArchiveServerBuilder) WithDescription(description string) *FixedUnixArchiveServerBuilder {
	b.SetDescription(description)
	return b
}

// GetIconID アイコンID 取得
func (b *FixedUnixArchiveServerBuilder) GetIconID() int64 {
	return b.iconID
}

// SetIconID アイコンID 設定
func (b *FixedUnixArchiveServerBuilder) SetIconID(iconID int64) {
	b.iconID = iconID
}

// WithIconID アイコンID 設定
func (b *FixedUnixArchiveServerBuilder) WithIconID(iconID int64) *FixedUnixArchiveServerBuilder {
	b.SetIconID(iconID)
	return b
}

// GetPrivateHostID 専有ID 取得
func (b *FixedUnixArchiveServerBuilder) GetPrivateHostID() int64 {
	return b.privateHostID
}

// SetPrivateHostID 専有ホストID 設定
func (b *FixedUnixArchiveServerBuilder) SetPrivateHostID(privateHostID int64) {
	b.privateHostID = privateHostID
}

// WithPrivateHostID 専有ホストID 設定
func (b *FixedUnixArchiveServerBuilder) WithPrivateHostID(privateHostID int64) *FixedUnixArchiveServerBuilder {
	b.privateHostID = privateHostID
	return b
}

// IsBootAfterCreate サーバー作成後すぐに起動フラグ 取得
func (b *FixedUnixArchiveServerBuilder) IsBootAfterCreate() bool {
	return b.bootAfterCreate
}

// SetBootAfterCreate サーバー作成後すぐに起動フラグ 設定
func (b *FixedUnixArchiveServerBuilder) SetBootAfterCreate(bootAfterCreate bool) {
	b.bootAfterCreate = bootAfterCreate
}

// WithBootAfterCreate サーバー作成後すぐに起動フラグ 設定
func (b *FixedUnixArchiveServerBuilder) WithBootAfterCreate(bootAfterCreate bool) *FixedUnixArchiveServerBuilder {
	b.SetBootAfterCreate(bootAfterCreate)
	return b
}

// GetTags タグ 取得
func (b *FixedUnixArchiveServerBuilder) GetTags() []string {
	return b.tags
}

// SetTags タグ 設定
func (b *FixedUnixArchiveServerBuilder) SetTags(tags []string) {
	b.tags = tags
}

// WithTags タグ 設定
func (b *FixedUnixArchiveServerBuilder) WithTags(tags []string) *FixedUnixArchiveServerBuilder {
	b.SetTags(tags)
	return b
}

/*---------------------------------------------------------
  for nic functioms
---------------------------------------------------------*/

// ClearNICConnections NIC接続設定 クリア
func (b *FixedUnixArchiveServerBuilder) ClearNICConnections() {
	b.nicConnections = nil
}

// WithEmptyNICConnections NIC接続設定 クリア
func (b *FixedUnixArchiveServerBuilder) WithEmptyNICConnections() *FixedUnixArchiveServerBuilder {
	b.nicConnections = nil
	return b
}

// AddPublicNWConnectedNIC 共有セグメントへの接続追加(注:共有セグメントはeth0のみ接続可能)
func (b *FixedUnixArchiveServerBuilder) AddPublicNWConnectedNIC() {
	b.nicConnections = append(b.nicConnections, "shared")
}

// WithAddPublicNWConnectedNIC 共有セグメントへの接続追加(注:共有セグメントはeth0のみ接続可能)
func (b *FixedUnixArchiveServerBuilder) WithAddPublicNWConnectedNIC() *FixedUnixArchiveServerBuilder {
	b.AddPublicNWConnectedNIC()
	return b
}

// AddExistsSwitchConnectedNIC スイッチ or ルーター+スイッチへの接続追加(注:ルーター+スイッチはeth0のみ接続可能)
func (b *FixedUnixArchiveServerBuilder) AddExistsSwitchConnectedNIC(switchID string, ipaddress string, networkMaskLen int, defaultRoute string) {
	b.nicConnections = append(b.nicConnections, switchID)
	b.disk.SetIPAddress(ipaddress)
	b.disk.SetNetworkMaskLen(networkMaskLen)
	b.disk.SetDefaultRoute(defaultRoute)
}

// WithAddExistsSwitchConnectedNIC スイッチ or ルーター+スイッチへの接続追加(注:ルーター+スイッチはeth0のみ接続可能)
func (b *FixedUnixArchiveServerBuilder) WithAddExistsSwitchConnectedNIC(switchID string, ipaddress string, networkMaskLen int, defaultRoute string) *FixedUnixArchiveServerBuilder {
	b.AddExistsSwitchConnectedNIC(switchID, ipaddress, networkMaskLen, defaultRoute)
	return b
}

// AddDisconnectedNIC 切断されたNIC追加
func (b *FixedUnixArchiveServerBuilder) AddDisconnectedNIC() {
	b.nicConnections = append(b.nicConnections, "")
}

// WithAddDisconnectedNIC 切断されたNIC追加
func (b *FixedUnixArchiveServerBuilder) WithAddDisconnectedNIC() *FixedUnixArchiveServerBuilder {
	b.nicConnections = append(b.nicConnections, "")
	return b
}

// GetPacketFilterIDs パケットフィルタID 取得
func (b *FixedUnixArchiveServerBuilder) GetPacketFilterIDs() []int64 {
	return b.packetFilterIDs
}

// SetPacketFilterIDs パケットフィルタID 設定
func (b *FixedUnixArchiveServerBuilder) SetPacketFilterIDs(ids []int64) {
	b.packetFilterIDs = ids
}

// WithPacketFilterIDs パケットフィルタID 設定
func (b *FixedUnixArchiveServerBuilder) WithPacketFilterIDs(ids []int64) *FixedUnixArchiveServerBuilder {
	b.packetFilterIDs = ids
	return b
}

// GetISOImageID ISOイメージ(CDROM)ID 取得
func (b *FixedUnixArchiveServerBuilder) GetISOImageID() int64 {
	return b.isoImageID
}

// SetISOImageID ISOイメージ(CDROM)ID 設定
func (b *FixedUnixArchiveServerBuilder) SetISOImageID(id int64) {
	b.isoImageID = id
}

// WithISOImageID ISOイメージ(CDROM)ID 設定
func (b *FixedUnixArchiveServerBuilder) WithISOImageID(id int64) *FixedUnixArchiveServerBuilder {
	b.SetISOImageID(id)
	return b
}

/*---------------------------------------------------------
  for disk properties
---------------------------------------------------------*/

// GetDiskSize ディスクサイズ(GB単位) 取得
func (b *FixedUnixArchiveServerBuilder) GetDiskSize() int {
	return b.disk.GetSize()
}

// SetDiskSize ディスクサイズ(GB単位) 設定
func (b *FixedUnixArchiveServerBuilder) SetDiskSize(diskSize int) {
	b.disk.SetSize(diskSize)
}

// WithDiskSize ディスクサイズ(GB単位) 設定
func (b *FixedUnixArchiveServerBuilder) WithDiskSize(diskSize int) *FixedUnixArchiveServerBuilder {
	b.SetDiskSize(diskSize)
	return b
}

// GetDistantFrom ストレージ隔離対象ディスク 取得
func (b *FixedUnixArchiveServerBuilder) GetDistantFrom() []int64 {
	return b.disk.GetDistantFrom()
}

// SetDistantFrom ストレージ隔離対象ディスク 設定
func (b *FixedUnixArchiveServerBuilder) SetDistantFrom(distantFrom []int64) {
	b.disk.SetDistantFrom(distantFrom)
}

// WithDistantFrom ストレージ隔離対象ディスク 設定
func (b *FixedUnixArchiveServerBuilder) WithDistantFrom(distantFrom []int64) *FixedUnixArchiveServerBuilder {
	b.SetDistantFrom(distantFrom)
	return b
}

// AddDistantFrom ストレージ隔離対象ディスク 追加
func (b *FixedUnixArchiveServerBuilder) AddDistantFrom(diskID int64) {
	b.disk.AddDistantFrom(diskID)
}

// WithAddDistantFrom ストレージ隔離対象ディスク 追加
func (b *FixedUnixArchiveServerBuilder) WithAddDistantFrom(diskID int64) *FixedUnixArchiveServerBuilder {
	b.AddDistantFrom(diskID)
	return b
}

// ClearDistantFrom ストレージ隔離対象ディスク クリア
func (b *FixedUnixArchiveServerBuilder) ClearDistantFrom() {
	b.disk.ClearDistantFrom()
}

// WithEmptyDistantFrom ストレージ隔離対象ディスク クリア
func (b *FixedUnixArchiveServerBuilder) WithEmptyDistantFrom() *FixedUnixArchiveServerBuilder {
	b.ClearDistantFrom()
	return b
}

// GetDiskPlanID ディスクプラン(SSD/HDD) 取得
func (b *FixedUnixArchiveServerBuilder) GetDiskPlanID() sacloud.DiskPlanID {
	return b.disk.GetPlanID()
}

// SetDiskPlanID ディスクプラン(SSD/HDD) 設定
func (b *FixedUnixArchiveServerBuilder) SetDiskPlanID(diskPlanID sacloud.DiskPlanID) {
	b.disk.SetPlanID(diskPlanID)
}

// WithDiskPlanID ディスクプラン(SSD/HDD) 設定
func (b *FixedUnixArchiveServerBuilder) WithDiskPlanID(diskPlanID sacloud.DiskPlanID) *FixedUnixArchiveServerBuilder {
	b.SetDiskPlanID(diskPlanID)
	return b
}

// SetDiskPlan ディスクプラン(ssd/hdd) 設定(文字列から)
func (b *FixedUnixArchiveServerBuilder) SetDiskPlan(plan string) {
	b.disk.SetPlan(plan)
}

// WithDiskPlan ディスクプラン(ssd/hdd) 設定(文字列から)
func (b *FixedUnixArchiveServerBuilder) WithDiskPlan(plan string) *FixedUnixArchiveServerBuilder {
	b.SetDiskPlan(plan)
	return b
}

// GetDiskConnection ディスク接続方法(VirtIO/IDE) 取得
func (b *FixedUnixArchiveServerBuilder) GetDiskConnection() sacloud.EDiskConnection {
	return b.disk.GetConnection()
}

// SetDiskConnection ディスク接続方法(VirtIO/IDE) 設定
func (b *FixedUnixArchiveServerBuilder) SetDiskConnection(diskConnection sacloud.EDiskConnection) {
	b.disk.SetConnection(diskConnection)
}

// WithDiskConnection ディスク接続方法(VirtIO/IDE) 設定
func (b *FixedUnixArchiveServerBuilder) WithDiskConnection(diskConnection sacloud.EDiskConnection) *FixedUnixArchiveServerBuilder {
	b.SetDiskConnection(diskConnection)
	return b
}

/*---------------------------------------------------------
  for disk edit properties
---------------------------------------------------------*/

// GetSourceArchiveID ソースアーカイブID 取得
func (b *FixedUnixArchiveServerBuilder) GetSourceArchiveID() int64 {
	return b.disk.GetSourceArchiveID()
}

// GetSourceDiskID ソースディスクID 設定
func (b *FixedUnixArchiveServerBuilder) GetSourceDiskID() int64 {
	return b.disk.GetSourceDiskID()
}

/*---------------------------------------------------------
  for event handler
---------------------------------------------------------*/

// SetEventHandler イベントハンドラ 設定
func (b *FixedUnixArchiveServerBuilder) SetEventHandler(event ServerBuildEvents, handler ServerBuildEventHandler) {
	b.buildEventHandlers[event] = handler
}

// WithEventHandler イベントハンドラ 設定
func (b *FixedUnixArchiveServerBuilder) WithEventHandler(event ServerBuildEvents, handler ServerBuildEventHandler) *FixedUnixArchiveServerBuilder {
	b.SetEventHandler(event, handler)
	return b
}

// ClearEventHandler イベントハンドラ クリア
func (b *FixedUnixArchiveServerBuilder) ClearEventHandler(event ServerBuildEvents) {
	delete(b.buildEventHandlers, event)
}

// WithEmptyEventHandler イベントハンドラ クリア
func (b *FixedUnixArchiveServerBuilder) WithEmptyEventHandler(event ServerBuildEvents) *FixedUnixArchiveServerBuilder {
	b.ClearEventHandler(event)
	return b
}

// GetEventHandler イベントハンドラ 取得
func (b *FixedUnixArchiveServerBuilder) GetEventHandler(event ServerBuildEvents) *ServerBuildEventHandler {
	if handler, ok := b.buildEventHandlers[event]; ok {
		return &handler
	}
	return nil
}

// SetDiskEventHandler ディスクイベントハンドラ 設定
func (b *FixedUnixArchiveServerBuilder) SetDiskEventHandler(event DiskBuildEvents, handler DiskBuildEventHandler) {
	b.disk.SetEventHandler(event, handler)
}

// WithDiskEventHandler ディスクイベントハンドラ 設定
func (b *FixedUnixArchiveServerBuilder) WithDiskEventHandler(event DiskBuildEvents, handler DiskBuildEventHandler) *FixedUnixArchiveServerBuilder {
	b.SetDiskEventHandler(event, handler)
	return b
}

// ClearDiskEventHandler ディスクイベントハンドラ クリア
func (b *FixedUnixArchiveServerBuilder) ClearDiskEventHandler(event DiskBuildEvents) *FixedUnixArchiveServerBuilder {
	b.disk.ClearEventHandler(event)
	return b
}

// WithEmptyDiskEventHandler ディスクイベントハンドラ クリア
func (b *FixedUnixArchiveServerBuilder) WithEmptyDiskEventHandler(event DiskBuildEvents) *FixedUnixArchiveServerBuilder {
	b.ClearDiskEventHandler(event)
	return b
}

// GetDiskEventHandler ディスクイベントハンドラ 取得
func (b *FixedUnixArchiveServerBuilder) GetDiskEventHandler(event DiskBuildEvents) *DiskBuildEventHandler {
	return b.disk.GetEventHandler(event)
}

/*---------------------------------------------------------
  for additional disks
---------------------------------------------------------*/

// AddAdditionalDisk 追加ディスク 追加
func (b *FixedUnixArchiveServerBuilder) AddAdditionalDisk(diskBuilder *DiskBuilder) {
	b.additionalDisks = append(b.additionalDisks, diskBuilder)
}

// WithAddAdditionalDisk 追加ディスク 追加
func (b *FixedUnixArchiveServerBuilder) WithAddAdditionalDisk(diskBuilder *DiskBuilder) *FixedUnixArchiveServerBuilder {
	b.additionalDisks = append(b.additionalDisks, diskBuilder)
	return b
}

// ClearAdditionalDisks 追加ディスク クリア
func (b *FixedUnixArchiveServerBuilder) ClearAdditionalDisks() {
	b.additionalDisks = []*DiskBuilder{}
}

// WithEmptyAdditionalDisks 追加ディスク クリア
func (b *FixedUnixArchiveServerBuilder) WithEmptyAdditionalDisks() *FixedUnixArchiveServerBuilder {
	b.ClearAdditionalDisks()
	return b
}

// GetAdditionalDisks 追加ディスク 取得
func (b *FixedUnixArchiveServerBuilder) GetAdditionalDisks() []*DiskBuilder {
	return b.additionalDisks
}
