package builder

import "github.com/sacloud/libsacloud/sacloud"

// PublicArchiveUnixServerBuilder Linux(Unix)系パブリックアーカイブを利用して構築を行うサーバービルダー
//
// 基本的なディスク設定やディスクの追加、ディスクの修正機能に対応しています。
type PublicArchiveUnixServerBuilder struct {
	*serverBuilder
}

/*---------------------------------------------------------
  common properties
---------------------------------------------------------*/

// GetServerName サーバー名 取得
func (b *PublicArchiveUnixServerBuilder) GetServerName() string {
	return b.serverName
}

// SetServerName サーバー名 設定
func (b *PublicArchiveUnixServerBuilder) SetServerName(serverName string) {
	b.serverName = serverName
}

// WithServerName サーバー名 設定
func (b *PublicArchiveUnixServerBuilder) WithServerName(serverName string) *PublicArchiveUnixServerBuilder {
	b.SetServerName(serverName)
	return b
}

// GetCore CPUコア数 取得
func (b *PublicArchiveUnixServerBuilder) GetCore() int {
	return b.core
}

// SetCore CPUコア数 設定
func (b *PublicArchiveUnixServerBuilder) SetCore(core int) {
	b.core = core
}

// WithCore CPUコア数 設定
func (b *PublicArchiveUnixServerBuilder) WithCore(core int) *PublicArchiveUnixServerBuilder {
	b.SetCore(core)
	return b
}

// GetMemory メモリサイズ(GB単位) 取得
func (b *PublicArchiveUnixServerBuilder) GetMemory() int {
	return b.memory
}

// SetMemory メモリサイズ(GB単位) 設定
func (b *PublicArchiveUnixServerBuilder) SetMemory(memory int) {
	b.memory = memory
}

// WithMemory メモリサイズ(GB単位) 設定
func (b *PublicArchiveUnixServerBuilder) WithMemory(memory int) *PublicArchiveUnixServerBuilder {
	b.SetMemory(memory)
	return b
}

// GetInterfaceDriver インターフェースドライバ 取得
func (b *PublicArchiveUnixServerBuilder) GetInterfaceDriver() sacloud.EInterfaceDriver {
	return b.interfaceDriver
}

// SetInterfaceDriver インターフェースドライバ 設定
func (b *PublicArchiveUnixServerBuilder) SetInterfaceDriver(interfaceDriver sacloud.EInterfaceDriver) {
	b.interfaceDriver = interfaceDriver
}

// WithInterfaceDriver インターフェースドライバ 設定
func (b *PublicArchiveUnixServerBuilder) WithInterfaceDriver(interfaceDriver sacloud.EInterfaceDriver) *PublicArchiveUnixServerBuilder {
	b.SetInterfaceDriver(interfaceDriver)
	return b
}

// GetDescription 説明 取得
func (b *PublicArchiveUnixServerBuilder) GetDescription() string {
	return b.description
}

// SetDescription 説明 設定
func (b *PublicArchiveUnixServerBuilder) SetDescription(description string) {
	b.description = description
}

// WithDescription 説明 設定
func (b *PublicArchiveUnixServerBuilder) WithDescription(description string) *PublicArchiveUnixServerBuilder {
	b.SetDescription(description)
	return b
}

// GetIconID アイコンID 取得
func (b *PublicArchiveUnixServerBuilder) GetIconID() int64 {
	return b.iconID
}

// SetIconID アイコンID 設定
func (b *PublicArchiveUnixServerBuilder) SetIconID(iconID int64) {
	b.iconID = iconID
}

// WithIconID アイコンID 設定
func (b *PublicArchiveUnixServerBuilder) WithIconID(iconID int64) *PublicArchiveUnixServerBuilder {
	b.SetIconID(iconID)
	return b
}

// GetPrivateHostID 専有ホストID 取得
func (b *PublicArchiveUnixServerBuilder) GetPrivateHostID() int64 {
	return b.privateHostID
}

// SetPrivateHostID 専有ホストID 設定
func (b *PublicArchiveUnixServerBuilder) SetPrivateHostID(privateHostID int64) {
	b.privateHostID = privateHostID
}

// WithPrivateHostID 専有ホストID 設定
func (b *PublicArchiveUnixServerBuilder) WithPrivateHostID(privateHostID int64) *PublicArchiveUnixServerBuilder {
	b.privateHostID = privateHostID
	return b
}

// IsBootAfterCreate サーバー作成後すぐに起動フラグ 取得
func (b *PublicArchiveUnixServerBuilder) IsBootAfterCreate() bool {
	return b.bootAfterCreate
}

// SetBootAfterCreate サーバー作成後すぐに起動フラグ 設定
func (b *PublicArchiveUnixServerBuilder) SetBootAfterCreate(bootAfterCreate bool) {
	b.bootAfterCreate = bootAfterCreate
}

// WithBootAfterCreate サーバー作成後すぐに起動フラグ 設定
func (b *PublicArchiveUnixServerBuilder) WithBootAfterCreate(bootAfterCreate bool) *PublicArchiveUnixServerBuilder {
	b.SetBootAfterCreate(bootAfterCreate)
	return b
}

// GetTags タグ 取得
func (b *PublicArchiveUnixServerBuilder) GetTags() []string {
	return b.tags
}

// SetTags タグ 設定
func (b *PublicArchiveUnixServerBuilder) SetTags(tags []string) {
	b.tags = tags
}

// WithTags タグ 設定
func (b *PublicArchiveUnixServerBuilder) WithTags(tags []string) *PublicArchiveUnixServerBuilder {
	b.SetTags(tags)
	return b
}

/*---------------------------------------------------------
  for nic functioms
---------------------------------------------------------*/

// ClearNICConnections NIC接続設定 クリア
func (b *PublicArchiveUnixServerBuilder) ClearNICConnections() {
	b.nicConnections = nil
	b.disk.SetIPAddress("")
	b.disk.SetNetworkMaskLen(0)
	b.disk.SetDefaultRoute("")
}

// WithEmptyNICConnections NIC接続設定 クリア
func (b *PublicArchiveUnixServerBuilder) WithEmptyNICConnections() *PublicArchiveUnixServerBuilder {
	b.ClearNICConnections()
	return b
}

// AddPublicNWConnectedNIC 共有セグメントへの接続追加(注:共有セグメントはeth0のみ接続可能)
func (b *PublicArchiveUnixServerBuilder) AddPublicNWConnectedNIC() {
	b.nicConnections = append(b.nicConnections, "shared")
}

// WithAddPublicNWConnectedNIC 共有セグメントへの接続追加(注:共有セグメントはeth0のみ接続可能)
func (b *PublicArchiveUnixServerBuilder) WithAddPublicNWConnectedNIC() *PublicArchiveUnixServerBuilder {
	b.AddPublicNWConnectedNIC()
	return b
}

// AddExistsSwitchConnectedNIC スイッチ or ルーター+スイッチへの接続追加(注:ルーター+スイッチはeth0のみ接続可能)
func (b *PublicArchiveUnixServerBuilder) AddExistsSwitchConnectedNIC(switchID string, ipaddress string, networkMaskLen int, defaultRoute string) {
	b.nicConnections = append(b.nicConnections, switchID)
	b.disk.SetIPAddress(ipaddress)
	b.disk.SetNetworkMaskLen(networkMaskLen)
	b.disk.SetDefaultRoute(defaultRoute)
}

// WithAddExistsSwitchConnectedNIC スイッチ or ルーター+スイッチへの接続追加(注:ルーター+スイッチはeth0のみ接続可能)
func (b *PublicArchiveUnixServerBuilder) WithAddExistsSwitchConnectedNIC(switchID string, ipaddress string, networkMaskLen int, defaultRoute string) *PublicArchiveUnixServerBuilder {
	b.AddExistsSwitchConnectedNIC(switchID, ipaddress, networkMaskLen, defaultRoute)
	return b
}

// AddDisconnectedNIC 切断されたNIC追加
func (b *PublicArchiveUnixServerBuilder) AddDisconnectedNIC() {
	b.nicConnections = append(b.nicConnections, "")
}

// WithAddDisconnectedNIC 切断されたNIC追加
func (b *PublicArchiveUnixServerBuilder) WithAddDisconnectedNIC() *PublicArchiveUnixServerBuilder {
	b.nicConnections = append(b.nicConnections, "")
	return b
}

// GetPacketFilterIDs パケットフィルタID 取得
func (b *PublicArchiveUnixServerBuilder) GetPacketFilterIDs() []int64 {
	return b.packetFilterIDs
}

// SetPacketFilterIDs パケットフィルタID 設定
func (b *PublicArchiveUnixServerBuilder) SetPacketFilterIDs(ids []int64) {
	b.packetFilterIDs = ids
}

// WithPacketFilterIDs パケットフィルタID 設定
func (b *PublicArchiveUnixServerBuilder) WithPacketFilterIDs(ids []int64) *PublicArchiveUnixServerBuilder {
	b.packetFilterIDs = ids
	return b
}

// GetISOImageID ISOイメージ(CDROM)ID 取得
func (b *PublicArchiveUnixServerBuilder) GetISOImageID() int64 {
	return b.isoImageID
}

// SetISOImageID ISOイメージ(CDROM)ID 設定
func (b *PublicArchiveUnixServerBuilder) SetISOImageID(id int64) {
	b.isoImageID = id
}

// WithISOImageID ISOイメージ(CDROM)ID 設定
func (b *PublicArchiveUnixServerBuilder) WithISOImageID(id int64) *PublicArchiveUnixServerBuilder {
	b.SetISOImageID(id)
	return b
}

/*---------------------------------------------------------
  for disk properties
---------------------------------------------------------*/

// GetDiskSize ディスクサイズ(GB単位) 取得
func (b *PublicArchiveUnixServerBuilder) GetDiskSize() int {
	return b.disk.GetSize()
}

// SetDiskSize ディスクサイズ(GB単位) 設定
func (b *PublicArchiveUnixServerBuilder) SetDiskSize(diskSize int) {
	b.disk.SetSize(diskSize)
}

// WithDiskSize ディスクサイズ(GB単位) 設定
func (b *PublicArchiveUnixServerBuilder) WithDiskSize(diskSize int) *PublicArchiveUnixServerBuilder {
	b.SetDiskSize(diskSize)
	return b
}

// GetDistantFrom ストレージ隔離対象ディスク 取得
func (b *PublicArchiveUnixServerBuilder) GetDistantFrom() []int64 {
	return b.disk.GetDistantFrom()
}

// SetDistantFrom ストレージ隔離対象ディスク 設定
func (b *PublicArchiveUnixServerBuilder) SetDistantFrom(distantFrom []int64) {
	b.disk.SetDistantFrom(distantFrom)
}

// WithDistantFrom ストレージ隔離対象ディスク 設定
func (b *PublicArchiveUnixServerBuilder) WithDistantFrom(distantFrom []int64) *PublicArchiveUnixServerBuilder {
	b.SetDistantFrom(distantFrom)
	return b
}

// AddDistantFrom ストレージ隔離対象ディスク 追加
func (b *PublicArchiveUnixServerBuilder) AddDistantFrom(diskID int64) {
	b.disk.AddDistantFrom(diskID)
}

// WithAddDistantFrom ストレージ隔離対象ディスク 追加
func (b *PublicArchiveUnixServerBuilder) WithAddDistantFrom(diskID int64) *PublicArchiveUnixServerBuilder {
	b.AddDistantFrom(diskID)
	return b
}

// ClearDistantFrom ストレージ隔離対象ディスク クリア
func (b *PublicArchiveUnixServerBuilder) ClearDistantFrom() {
	b.disk.ClearDistantFrom()
}

// WithEmptyDistantFrom ストレージ隔離対象ディスク クリア
func (b *PublicArchiveUnixServerBuilder) WithEmptyDistantFrom() *PublicArchiveUnixServerBuilder {
	b.ClearDistantFrom()
	return b
}

// GetDiskPlanID ディスクプラン(SSD/HDD) 取得
func (b *PublicArchiveUnixServerBuilder) GetDiskPlanID() sacloud.DiskPlanID {
	return b.disk.GetPlanID()
}

// SetDiskPlanID ディスクプラン(SSD/HDD) 設定
func (b *PublicArchiveUnixServerBuilder) SetDiskPlanID(diskPlanID sacloud.DiskPlanID) {
	b.disk.SetPlanID(diskPlanID)
}

// WithDiskPlanID ディスクプラン(SSD/HDD) 設定
func (b *PublicArchiveUnixServerBuilder) WithDiskPlanID(diskPlanID sacloud.DiskPlanID) *PublicArchiveUnixServerBuilder {
	b.SetDiskPlanID(diskPlanID)
	return b
}

// SetDiskPlan ディスクプラン(ssd/hdd) 設定(文字列から)
func (b *PublicArchiveUnixServerBuilder) SetDiskPlan(plan string) {
	b.disk.SetPlan(plan)
}

// WithDiskPlan ディスクプラン(ssd/hdd) 設定(文字列から)
func (b *PublicArchiveUnixServerBuilder) WithDiskPlan(plan string) *PublicArchiveUnixServerBuilder {
	b.SetDiskPlan(plan)
	return b
}

// GetDiskConnection ディスク接続方法(VirtIO/IDE) 取得
func (b *PublicArchiveUnixServerBuilder) GetDiskConnection() sacloud.EDiskConnection {
	return b.disk.GetConnection()
}

// SetDiskConnection ディスク接続方法(VirtIO/IDE) 設定
func (b *PublicArchiveUnixServerBuilder) SetDiskConnection(diskConnection sacloud.EDiskConnection) {
	b.disk.SetConnection(diskConnection)
}

// WithDiskConnection ディスク接続方法(VirtIO/IDE) 設定
func (b *PublicArchiveUnixServerBuilder) WithDiskConnection(diskConnection sacloud.EDiskConnection) *PublicArchiveUnixServerBuilder {
	b.SetDiskConnection(diskConnection)
	return b
}

/*---------------------------------------------------------
  for disk edit properties
---------------------------------------------------------*/

// GetSourceArchiveID ソースアーカイブID 取得
func (b *PublicArchiveUnixServerBuilder) GetSourceArchiveID() int64 {
	return b.disk.GetSourceArchiveID()
}

// GetSourceDiskID ソースディスクID 設定
func (b *PublicArchiveUnixServerBuilder) GetSourceDiskID() int64 {
	return b.disk.GetSourceDiskID()
}

// GetPassword パスワード 取得
func (b *PublicArchiveUnixServerBuilder) GetPassword() string {
	return b.disk.GetPassword()
}

// SetPassword パスワード 設定
func (b *PublicArchiveUnixServerBuilder) SetPassword(password string) {
	b.disk.SetPassword(password)
}

// WithPassword パスワード 設定
func (b *PublicArchiveUnixServerBuilder) WithPassword(password string) *PublicArchiveUnixServerBuilder {
	b.SetPassword(password)
	return b
}

// GetHostName ホスト名 取得
func (b *PublicArchiveUnixServerBuilder) GetHostName() string {
	return b.disk.GetHostName()
}

// SetHostName ホスト名 設定
func (b *PublicArchiveUnixServerBuilder) SetHostName(hostName string) {
	b.disk.SetHostName(hostName)
}

// WithHostName ホスト名 設定
func (b *PublicArchiveUnixServerBuilder) WithHostName(hostName string) *PublicArchiveUnixServerBuilder {
	b.SetHostName(hostName)
	return b
}

// IsDisablePWAuth パスワード認証無効化フラグ 取得
func (b *PublicArchiveUnixServerBuilder) IsDisablePWAuth() bool {
	return b.disk.IsDisablePWAuth()
}

// SetDisablePWAuth パスワード認証無効化フラグ 設定
func (b *PublicArchiveUnixServerBuilder) SetDisablePWAuth(disable bool) {
	b.disk.SetDisablePWAuth(disable)
}

// WithDisablePWAuth パスワード認証無効化フラグ 設定
func (b *PublicArchiveUnixServerBuilder) WithDisablePWAuth(disable bool) *PublicArchiveUnixServerBuilder {
	b.SetDisablePWAuth(disable)
	return b
}

// AddSSHKey 公開鍵 追加
func (b *PublicArchiveUnixServerBuilder) AddSSHKey(sshKey string) {
	b.disk.AddSSHKey(sshKey)
}

// WithAddSSHKey 公開鍵 追加
func (b *PublicArchiveUnixServerBuilder) WithAddSSHKey(sshKey string) *PublicArchiveUnixServerBuilder {
	b.disk.AddSSHKey(sshKey)
	return b
}

// ClearSSHKey 公開鍵 クリア
func (b *PublicArchiveUnixServerBuilder) ClearSSHKey() {
	b.disk.ClearSSHKey()
}

// WithEmptySSHKey 公開鍵 クリア
func (b *PublicArchiveUnixServerBuilder) WithEmptySSHKey() *PublicArchiveUnixServerBuilder {
	b.ClearSSHKey()
	return b
}

// GetSSHKeys 公開鍵 取得
func (b *PublicArchiveUnixServerBuilder) GetSSHKeys() []string {
	return b.disk.GetSSHKeys()
}

// AddSSHKeyID 公開鍵ID 追加
func (b *PublicArchiveUnixServerBuilder) AddSSHKeyID(sshKeyID int64) {
	b.disk.AddSSHKeyID(sshKeyID)
}

// WithAddSSHKeyID 公開鍵ID 追加
func (b *PublicArchiveUnixServerBuilder) WithAddSSHKeyID(sshKeyID int64) *PublicArchiveUnixServerBuilder {
	b.AddSSHKeyID(sshKeyID)
	return b
}

// ClearSSHKeyIDs 公開鍵ID クリア
func (b *PublicArchiveUnixServerBuilder) ClearSSHKeyIDs() {
	b.disk.ClearSSHKeyIDs()
}

// WithEmptySSHKeyIDs 公開鍵ID クリア
func (b *PublicArchiveUnixServerBuilder) WithEmptySSHKeyIDs() *PublicArchiveUnixServerBuilder {
	b.ClearSSHKeyIDs()
	return b
}

// GetSSHKeyIds 公開鍵ID 取得
func (b *PublicArchiveUnixServerBuilder) GetSSHKeyIds() []int64 {
	return b.disk.GetSSHKeyIds()
}

// AddNote スタートアップスクリプト 追加
func (b *PublicArchiveUnixServerBuilder) AddNote(note string) {
	b.disk.AddNote(note)
}

// WithAddNote スタートアップスクリプト 追加
func (b *PublicArchiveUnixServerBuilder) WithAddNote(note string) *PublicArchiveUnixServerBuilder {
	b.AddNote(note)
	return b
}

// ClearNotes スタートアップスクリプト クリア
func (b *PublicArchiveUnixServerBuilder) ClearNotes() {
	b.disk.ClearNotes()
}

// WithEmptyNotes スタートアップスクリプト クリア
func (b *PublicArchiveUnixServerBuilder) WithEmptyNotes() *PublicArchiveUnixServerBuilder {
	b.ClearNotes()
	return b
}

// GetNotes スタートアップスクリプト 取得
func (b *PublicArchiveUnixServerBuilder) GetNotes() []string {
	return b.disk.GetNotes()
}

// AddNoteID スタートアップスクリプト 追加
func (b *PublicArchiveUnixServerBuilder) AddNoteID(noteID int64) {
	b.disk.AddNoteID(noteID)
}

// WithAddNoteID スタートアップスクリプト 追加
func (b *PublicArchiveUnixServerBuilder) WithAddNoteID(noteID int64) *PublicArchiveUnixServerBuilder {
	b.AddNoteID(noteID)
	return b
}

// ClearNoteIDs スタートアップスクリプト クリア
func (b *PublicArchiveUnixServerBuilder) ClearNoteIDs() {
	b.disk.ClearNoteIDs()
}

// WithEmptyNoteIDs スタートアップスクリプト クリア
func (b *PublicArchiveUnixServerBuilder) WithEmptyNoteIDs() *PublicArchiveUnixServerBuilder {
	b.ClearNoteIDs()
	return b
}

// GetNoteIDs スタートアップスクリプトID 取得
func (b *PublicArchiveUnixServerBuilder) GetNoteIDs() []int64 {
	return b.disk.GetNoteIDs()
}

// IsSSHKeysEphemeral ディスク作成後の公開鍵削除フラグ 取得
func (b *PublicArchiveUnixServerBuilder) IsSSHKeysEphemeral() bool {
	return b.disk.IsSSHKeysEphemeral()
}

// SetSSHKeysEphemeral ディスク作成後の公開鍵削除フラグ 設定
func (b *PublicArchiveUnixServerBuilder) SetSSHKeysEphemeral(isEphemeral bool) {
	b.disk.SetSSHKeysEphemeral(isEphemeral)
}

// WithSSHKeysEphemeral ディスク作成後の公開鍵削除フラグ 設定
func (b *PublicArchiveUnixServerBuilder) WithSSHKeysEphemeral(isEphemeral bool) *PublicArchiveUnixServerBuilder {
	b.SetSSHKeysEphemeral(isEphemeral)
	return b
}

// IsNotesEphemeral ディスク作成後のスタートアップスクリプト削除フラグ 取得
func (b *PublicArchiveUnixServerBuilder) IsNotesEphemeral() bool {
	return b.disk.IsNotesEphemeral()
}

// SetNotesEphemeral ディスク作成後のスタートアップスクリプト削除フラグ 設定
func (b *PublicArchiveUnixServerBuilder) SetNotesEphemeral(isEphemeral bool) {
	b.disk.SetNotesEphemeral(isEphemeral)
}

// WithNotesEphemeral ディスク作成後のスタートアップスクリプト削除フラグ 設定
func (b *PublicArchiveUnixServerBuilder) WithNotesEphemeral(isEphemeral bool) *PublicArchiveUnixServerBuilder {
	b.SetNotesEphemeral(isEphemeral)
	return b
}

// GetGenerateSSHKeyName SSHキー生成 名称 取得
func (b *PublicArchiveUnixServerBuilder) GetGenerateSSHKeyName() string {
	return b.disk.GetGenerateSSHKeyName()
}

// SetGenerateSSHKeyName SSHキー生成 名称 設定
func (b *PublicArchiveUnixServerBuilder) SetGenerateSSHKeyName(name string) {
	b.disk.SetGenerateSSHKeyName(name)
}

// WithGenerateSSHKeyName SSHキー生成 名称 設定
func (b *PublicArchiveUnixServerBuilder) WithGenerateSSHKeyName(name string) *PublicArchiveUnixServerBuilder {
	b.disk.SetGenerateSSHKeyName(name)
	return b
}

// GetGenerateSSHKeyPassPhrase SSHキー生成 パスフレーズ 取得
func (b *PublicArchiveUnixServerBuilder) GetGenerateSSHKeyPassPhrase() string {
	return b.disk.GetGenerateSSHKeyPassPhrase()
}

// SetGenerateSSHKeyPassPhrase SSHキー生成 パスフレーズ 設定
func (b *PublicArchiveUnixServerBuilder) SetGenerateSSHKeyPassPhrase(pass string) {
	b.disk.SetGenerateSSHKeyPassPhrase(pass)
}

// WithGenerateSSHKeyPassPhrase SSHキー生成 パスフレーズ 設定
func (b *PublicArchiveUnixServerBuilder) WithGenerateSSHKeyPassPhrase(pass string) *PublicArchiveUnixServerBuilder {
	b.disk.SetGenerateSSHKeyPassPhrase(pass)
	return b
}

// GetGenerateSSHKeyDescription SSHキー生成 説明 取得
func (b *PublicArchiveUnixServerBuilder) GetGenerateSSHKeyDescription() string {
	return b.disk.GetGenerateSSHKeyDescription()
}

// SetGenerateSSHKeyDescription SSHキー生成 説明 設定
func (b *PublicArchiveUnixServerBuilder) SetGenerateSSHKeyDescription(desc string) {
	b.disk.SetGenerateSSHKeyDescription(desc)
}

// WithGenerateSSHKeyDescription SSHキー生成 説明 設定
func (b *PublicArchiveUnixServerBuilder) WithGenerateSSHKeyDescription(desc string) *PublicArchiveUnixServerBuilder {
	b.disk.SetGenerateSSHKeyDescription(desc)
	return b
}

/*---------------------------------------------------------
  for event handler
---------------------------------------------------------*/

// SetEventHandler イベントハンドラ 設定
func (b *PublicArchiveUnixServerBuilder) SetEventHandler(event ServerBuildEvents, handler ServerBuildEventHandler) {
	b.buildEventHandlers[event] = handler
}

// WithEventHandler イベントハンドラ 設定
func (b *PublicArchiveUnixServerBuilder) WithEventHandler(event ServerBuildEvents, handler ServerBuildEventHandler) *PublicArchiveUnixServerBuilder {
	b.SetEventHandler(event, handler)
	return b
}

// ClearEventHandler イベントハンドラ クリア
func (b *PublicArchiveUnixServerBuilder) ClearEventHandler(event ServerBuildEvents) {
	delete(b.buildEventHandlers, event)
}

// WithEmptyEventHandler イベントハンドラ クリア
func (b *PublicArchiveUnixServerBuilder) WithEmptyEventHandler(event ServerBuildEvents) *PublicArchiveUnixServerBuilder {
	delete(b.buildEventHandlers, event)
	return b
}

// GetEventHandler イベントハンドラ 取得
func (b *PublicArchiveUnixServerBuilder) GetEventHandler(event ServerBuildEvents) *ServerBuildEventHandler {
	if handler, ok := b.buildEventHandlers[event]; ok {
		return &handler
	}
	return nil
}

// SetDiskEventHandler ディスクイベントハンドラ 設定
func (b *PublicArchiveUnixServerBuilder) SetDiskEventHandler(event DiskBuildEvents, handler DiskBuildEventHandler) {
	b.disk.SetEventHandler(event, handler)
}

// WithDiskEventHandler ディスクイベントハンドラ 設定
func (b *PublicArchiveUnixServerBuilder) WithDiskEventHandler(event DiskBuildEvents, handler DiskBuildEventHandler) *PublicArchiveUnixServerBuilder {
	b.SetDiskEventHandler(event, handler)
	return b
}

// ClearDiskEventHandler ディスクイベントハンドラ クリア
func (b *PublicArchiveUnixServerBuilder) ClearDiskEventHandler(event DiskBuildEvents) {
	b.disk.ClearEventHandler(event)
}

// WithEmptyDiskEventHandler ディスクイベントハンドラ クリア
func (b *PublicArchiveUnixServerBuilder) WithEmptyDiskEventHandler(event DiskBuildEvents) *PublicArchiveUnixServerBuilder {
	b.ClearDiskEventHandler(event)
	return b
}

// GetDiskEventHandler ディスクイベントハンドラ 取得
func (b *PublicArchiveUnixServerBuilder) GetDiskEventHandler(event DiskBuildEvents) *DiskBuildEventHandler {
	return b.disk.GetEventHandler(event)
}

/*---------------------------------------------------------
  for additional disks
---------------------------------------------------------*/

// AddAdditionalDisk 追加ディスク 追加
func (b *PublicArchiveUnixServerBuilder) AddAdditionalDisk(diskBuilder *DiskBuilder) {
	b.additionalDisks = append(b.additionalDisks, diskBuilder)
}

// WithAddAdditionalDisk 追加ディスク 追加
func (b *PublicArchiveUnixServerBuilder) WithAddAdditionalDisk(diskBuilder *DiskBuilder) *PublicArchiveUnixServerBuilder {
	b.AddAdditionalDisk(diskBuilder)
	return b
}

// ClearAdditionalDisks 追加ディスク クリア
func (b *PublicArchiveUnixServerBuilder) ClearAdditionalDisks() {
	b.additionalDisks = []*DiskBuilder{}
}

// WithEmptyAdditionalDisks 追加ディスク クリア
func (b *PublicArchiveUnixServerBuilder) WithEmptyAdditionalDisks() *PublicArchiveUnixServerBuilder {
	b.ClearAdditionalDisks()
	return b
}

// GetAdditionalDisks 追加ディスク 取得
func (b *PublicArchiveUnixServerBuilder) GetAdditionalDisks() []*DiskBuilder {
	return b.additionalDisks
}
