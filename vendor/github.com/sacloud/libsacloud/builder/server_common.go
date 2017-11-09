package builder

import "github.com/sacloud/libsacloud/sacloud"

// CommonServerBuilder 既存のアーカイブ or ディスクを利用して構築を行うサーバービルダー
//
// 基本的なディスク設定やディスクの追加、ディスクの修正機能に対応しています。
// ただし、ディスクの修正機能はソースアーカイブ or ソースディスクが対応していない場合は
// さくらのクラウドAPIコール時にエラーとなるため、適切にハンドリングするように実装する必要があります。
type CommonServerBuilder struct {
	*serverBuilder
}

/*---------------------------------------------------------
  common properties
---------------------------------------------------------*/

// GetServerName サーバー名 取得
func (b *CommonServerBuilder) GetServerName() string {
	return b.serverName
}

// SetServerName サーバー名 設定
func (b *CommonServerBuilder) SetServerName(serverName string) {
	b.serverName = serverName
}

// WithServerName サーバー名 設定
func (b *CommonServerBuilder) WithServerName(serverName string) *CommonServerBuilder {
	b.SetServerName(serverName)
	return b
}

// GetCore CPUコア数 取得
func (b *CommonServerBuilder) GetCore() int {
	return b.core
}

// SetCore CPUコア数 設定
func (b *CommonServerBuilder) SetCore(core int) {
	b.core = core
}

// WithCore CPUコア数 設定
func (b *CommonServerBuilder) WithCore(core int) *CommonServerBuilder {
	b.SetCore(core)
	return b
}

// GetMemory メモリサイズ(GB単位) 取得
func (b *CommonServerBuilder) GetMemory() int {
	return b.memory
}

// SetMemory メモリサイズ(GB単位) 設定
func (b *CommonServerBuilder) SetMemory(memory int) {
	b.memory = memory
}

// WithMemory メモリサイズ(GB単位) 設定
func (b *CommonServerBuilder) WithMemory(memory int) *CommonServerBuilder {
	b.SetMemory(memory)
	return b
}

// GetInterfaceDriver インターフェースドライバ 取得
func (b *CommonServerBuilder) GetInterfaceDriver() sacloud.EInterfaceDriver {
	return b.interfaceDriver
}

// SetInterfaceDriver インターフェースドライバ 設定
func (b *CommonServerBuilder) SetInterfaceDriver(interfaceDriver sacloud.EInterfaceDriver) {
	b.interfaceDriver = interfaceDriver
}

// WithInterfaceDriver インターフェースドライバ 設定
func (b *CommonServerBuilder) WithInterfaceDriver(interfaceDriver sacloud.EInterfaceDriver) *CommonServerBuilder {
	b.SetInterfaceDriver(interfaceDriver)
	return b
}

// GetDescription 説明 取得
func (b *CommonServerBuilder) GetDescription() string {
	return b.description
}

// SetDescription 説明 設定
func (b *CommonServerBuilder) SetDescription(description string) {
	b.description = description
}

// WithDescription 説明 設定
func (b *CommonServerBuilder) WithDescription(description string) *CommonServerBuilder {
	b.SetDescription(description)
	return b
}

// GetIconID アイコンID 取得
func (b *CommonServerBuilder) GetIconID() int64 {
	return b.iconID
}

// SetIconID アイコンID 設定
func (b *CommonServerBuilder) SetIconID(iconID int64) {
	b.iconID = iconID
}

// WithIconID アイコンID 設定
func (b *CommonServerBuilder) WithIconID(iconID int64) *CommonServerBuilder {
	b.SetIconID(iconID)
	return b
}

// GetPrivateHostID 専有ホストID 取得
func (b *CommonServerBuilder) GetPrivateHostID() int64 {
	return b.privateHostID
}

// SetPrivateHostID 専有ホストID 設定
func (b *CommonServerBuilder) SetPrivateHostID(privateHostID int64) {
	b.privateHostID = privateHostID
}

// WithPrivateHostID 専有ホストID 設定
func (b *CommonServerBuilder) WithPrivateHostID(privateHostID int64) *CommonServerBuilder {
	b.privateHostID = privateHostID
	return b
}

// IsBootAfterCreate サーバー作成後すぐに起動フラグ 取得
func (b *CommonServerBuilder) IsBootAfterCreate() bool {
	return b.bootAfterCreate
}

// SetBootAfterCreate サーバー作成後すぐに起動フラグ 設定
func (b *CommonServerBuilder) SetBootAfterCreate(bootAfterCreate bool) {
	b.bootAfterCreate = bootAfterCreate
}

// WithBootAfterCreate サーバー作成後すぐに起動フラグ 設定
func (b *CommonServerBuilder) WithBootAfterCreate(bootAfterCreate bool) *CommonServerBuilder {
	b.SetBootAfterCreate(bootAfterCreate)
	return b
}

// GetTags タグ 取得
func (b *CommonServerBuilder) GetTags() []string {
	return b.tags
}

// SetTags タグ 設定
func (b *CommonServerBuilder) SetTags(tags []string) {
	b.tags = tags
}

// WithTags タグ 設定
func (b *CommonServerBuilder) WithTags(tags []string) *CommonServerBuilder {
	b.SetTags(tags)
	return b
}

/*---------------------------------------------------------
  for nic functioms
---------------------------------------------------------*/

// ClearNICConnections NIC接続設定 クリア
func (b *CommonServerBuilder) ClearNICConnections() {
	b.nicConnections = nil
	b.disk.SetIPAddress("")
	b.disk.SetNetworkMaskLen(0)
	b.disk.SetDefaultRoute("")
}

// WithEmptyNICConnections NIC接続設定 クリア
func (b *CommonServerBuilder) WithEmptyNICConnections() *CommonServerBuilder {
	b.ClearNICConnections()
	return b
}

// AddPublicNWConnectedNIC 共有セグメントへの接続追加(注:共有セグメントはeth0のみ接続可能)
func (b *CommonServerBuilder) AddPublicNWConnectedNIC() {
	b.nicConnections = append(b.nicConnections, "shared")
}

// WithAddPublicNWConnectedNIC 共有セグメントへの接続追加(注:共有セグメントはeth0のみ接続可能)
func (b *CommonServerBuilder) WithAddPublicNWConnectedNIC() *CommonServerBuilder {
	b.AddPublicNWConnectedNIC()
	return b
}

// AddExistsSwitchConnectedNIC スイッチ or ルーター+スイッチへの接続追加(注:ルーター+スイッチはeth0のみ接続可能)
func (b *CommonServerBuilder) AddExistsSwitchConnectedNIC(switchID string, ipaddress string, networkMaskLen int, defaultRoute string) {
	b.nicConnections = append(b.nicConnections, switchID)
	b.disk.SetIPAddress(ipaddress)
	b.disk.SetNetworkMaskLen(networkMaskLen)
	b.disk.SetDefaultRoute(defaultRoute)
}

// WithAddExistsSwitchConnectedNIC スイッチ or ルーター+スイッチへの接続追加(注:ルーター+スイッチはeth0のみ接続可能)
func (b *CommonServerBuilder) WithAddExistsSwitchConnectedNIC(switchID string, ipaddress string, networkMaskLen int, defaultRoute string) *CommonServerBuilder {
	b.AddExistsSwitchConnectedNIC(switchID, ipaddress, networkMaskLen, defaultRoute)
	return b
}

// AddDisconnectedNIC 切断されたNIC追加
func (b *CommonServerBuilder) AddDisconnectedNIC() {
	b.nicConnections = append(b.nicConnections, "")
}

// WithAddDisconnectedNIC 切断されたNIC追加
func (b *CommonServerBuilder) WithAddDisconnectedNIC() *CommonServerBuilder {
	b.AddDisconnectedNIC()
	return b
}

// GetPacketFilterIDs パケットフィルタID 取得
func (b *CommonServerBuilder) GetPacketFilterIDs() []int64 {
	return b.packetFilterIDs
}

// SetPacketFilterIDs パケットフィルタID 設定
func (b *CommonServerBuilder) SetPacketFilterIDs(ids []int64) {
	b.packetFilterIDs = ids
}

// WithPacketFilterIDs パケットフィルタID 設定
func (b *CommonServerBuilder) WithPacketFilterIDs(ids []int64) *CommonServerBuilder {
	b.packetFilterIDs = ids
	return b
}

// GetISOImageID ISOイメージ(CDROM)ID 取得
func (b *CommonServerBuilder) GetISOImageID() int64 {
	return b.isoImageID
}

// SetISOImageID ISOイメージ(CDROM)ID 設定
func (b *CommonServerBuilder) SetISOImageID(id int64) {
	b.isoImageID = id
}

// WithISOImageID ISOイメージ(CDROM)ID 設定
func (b *CommonServerBuilder) WithISOImageID(id int64) *CommonServerBuilder {
	b.SetISOImageID(id)
	return b
}

/*---------------------------------------------------------
  for disk properties
---------------------------------------------------------*/

// GetDiskSize ディスクサイズ(GB単位) 取得
func (b *CommonServerBuilder) GetDiskSize() int {
	return b.disk.GetSize()
}

// SetDiskSize ディスクサイズ(GB単位) 設定
func (b *CommonServerBuilder) SetDiskSize(diskSize int) {
	b.disk.SetSize(diskSize)
}

// WithDiskSize ディスクサイズ(GB単位) 設定
func (b *CommonServerBuilder) WithDiskSize(diskSize int) *CommonServerBuilder {
	b.SetDiskSize(diskSize)
	return b
}

// GetDistantFrom ストレージ隔離対象ディスク 取得
func (b *CommonServerBuilder) GetDistantFrom() []int64 {
	return b.disk.GetDistantFrom()
}

// SetDistantFrom ストレージ隔離対象ディスク 設定
func (b *CommonServerBuilder) SetDistantFrom(distantFrom []int64) {
	b.disk.SetDistantFrom(distantFrom)
}

// WithDistantFrom ストレージ隔離対象ディスク 設定
func (b *CommonServerBuilder) WithDistantFrom(distantFrom []int64) *CommonServerBuilder {
	b.SetDistantFrom(distantFrom)
	return b
}

// AddDistantFrom ストレージ隔離対象ディスク 追加
func (b *CommonServerBuilder) AddDistantFrom(diskID int64) {
	b.disk.AddDistantFrom(diskID)
}

// WithAddDistantFrom ストレージ隔離対象ディスク 追加
func (b *CommonServerBuilder) WithAddDistantFrom(diskID int64) *CommonServerBuilder {
	b.AddDistantFrom(diskID)
	return b
}

// ClearDistantFrom ストレージ隔離対象ディスク クリア
func (b *CommonServerBuilder) ClearDistantFrom() {
	b.disk.ClearDistantFrom()
}

// WithEmptyDistantFrom ストレージ隔離対象ディスク クリア
func (b *CommonServerBuilder) WithEmptyDistantFrom() *CommonServerBuilder {
	b.ClearDistantFrom()
	return b
}

// GetDiskPlanID ディスクプラン(SSD/HDD) 取得
func (b *CommonServerBuilder) GetDiskPlanID() sacloud.DiskPlanID {
	return b.disk.GetPlanID()
}

// SetDiskPlanID ディスクプラン(SSD/HDD) 設定
func (b *CommonServerBuilder) SetDiskPlanID(diskPlanID sacloud.DiskPlanID) {
	b.disk.SetPlanID(diskPlanID)
}

// WithDiskPlanID ディスクプラン(SSD/HDD) 設定
func (b *CommonServerBuilder) WithDiskPlanID(diskPlanID sacloud.DiskPlanID) *CommonServerBuilder {
	b.SetDiskPlanID(diskPlanID)
	return b
}

// SetDiskPlan ディスクプラン(ssd/hdd) 設定(文字列から)
func (b *CommonServerBuilder) SetDiskPlan(plan string) {
	b.disk.SetPlan(plan)
}

// WithDiskPlan ディスクプラン(ssd/hdd) 設定(文字列から)
func (b *CommonServerBuilder) WithDiskPlan(plan string) *CommonServerBuilder {
	b.SetDiskPlan(plan)
	return b
}

// GetDiskConnection ディスク接続方法(VirtIO/IDE) 取得
func (b *CommonServerBuilder) GetDiskConnection() sacloud.EDiskConnection {
	return b.disk.GetConnection()
}

// SetDiskConnection ディスク接続方法(VirtIO/IDE) 設定
func (b *CommonServerBuilder) SetDiskConnection(diskConnection sacloud.EDiskConnection) {
	b.disk.SetConnection(diskConnection)
}

// WithDiskConnection ディスク接続方法(VirtIO/IDE) 設定
func (b *CommonServerBuilder) WithDiskConnection(diskConnection sacloud.EDiskConnection) *CommonServerBuilder {
	b.SetDiskConnection(diskConnection)
	return b
}

/*---------------------------------------------------------
  for disk edit properties
---------------------------------------------------------*/

// GetSourceArchiveID ソースアーカイブID 取得
func (b *CommonServerBuilder) GetSourceArchiveID() int64 {
	return b.disk.GetSourceArchiveID()
}

// GetSourceDiskID ソースディスクID 設定
func (b *CommonServerBuilder) GetSourceDiskID() int64 {
	return b.disk.GetSourceDiskID()
}

// GetPassword パスワード 取得
func (b *CommonServerBuilder) GetPassword() string {
	return b.disk.GetPassword()
}

// SetPassword パスワード 設定
func (b *CommonServerBuilder) SetPassword(password string) {
	b.disk.SetPassword(password)
}

// WithPassword パスワード 設定
func (b *CommonServerBuilder) WithPassword(password string) *CommonServerBuilder {
	b.SetPassword(password)
	return b
}

// GetHostName ホスト名 取得
func (b *CommonServerBuilder) GetHostName() string {
	return b.disk.GetHostName()
}

// SetHostName ホスト名 設定
func (b *CommonServerBuilder) SetHostName(hostName string) {
	b.disk.SetHostName(hostName)
}

// WithHostName ホスト名 設定
func (b *CommonServerBuilder) WithHostName(hostName string) *CommonServerBuilder {
	b.SetHostName(hostName)
	return b
}

// IsDisablePWAuth パスワード認証無効化フラグ 取得
func (b *CommonServerBuilder) IsDisablePWAuth() bool {
	return b.disk.IsDisablePWAuth()
}

// SetDisablePWAuth パスワード認証無効化フラグ 設定
func (b *CommonServerBuilder) SetDisablePWAuth(disable bool) {
	b.disk.SetDisablePWAuth(disable)
}

// WithDisablePWAuth パスワード認証無効化フラグ 設定
func (b *CommonServerBuilder) WithDisablePWAuth(disable bool) *CommonServerBuilder {
	b.SetDisablePWAuth(disable)
	return b
}

// AddSSHKey 公開鍵 追加
func (b *CommonServerBuilder) AddSSHKey(sshKey string) {
	b.disk.AddSSHKey(sshKey)
}

// WithAddSSHKey 公開鍵 追加
func (b *CommonServerBuilder) WithAddSSHKey(sshKey string) *CommonServerBuilder {
	b.AddSSHKey(sshKey)
	return b
}

// ClearSSHKey 公開鍵 クリア
func (b *CommonServerBuilder) ClearSSHKey() {
	b.disk.ClearSSHKey()
}

// WithEmptySSHKey 公開鍵 クリア
func (b *CommonServerBuilder) WithEmptySSHKey() *CommonServerBuilder {
	b.ClearSSHKey()
	return b
}

// GetSSHKeys 公開鍵 取得
func (b *CommonServerBuilder) GetSSHKeys() []string {
	return b.disk.GetSSHKeys()
}

// AddSSHKeyID 公開鍵ID 追加
func (b *CommonServerBuilder) AddSSHKeyID(sshKeyID int64) {
	b.disk.AddSSHKeyID(sshKeyID)
}

// WithAddSSHKeyID 公開鍵ID 追加
func (b *CommonServerBuilder) WithAddSSHKeyID(sshKeyID int64) *CommonServerBuilder {
	b.AddSSHKeyID(sshKeyID)
	return b
}

// ClearSSHKeyIDs 公開鍵ID クリア
func (b *CommonServerBuilder) ClearSSHKeyIDs() {
	b.disk.ClearSSHKeyIDs()
}

// WithEmptySSHKeyIDs 公開鍵ID クリア
func (b *CommonServerBuilder) WithEmptySSHKeyIDs() *CommonServerBuilder {
	b.ClearSSHKeyIDs()
	return b
}

// GetSSHKeyIds 公開鍵ID 取得
func (b *CommonServerBuilder) GetSSHKeyIds() []int64 {
	return b.disk.GetSSHKeyIds()
}

// AddNote スタートアップスクリプト 追加
func (b *CommonServerBuilder) AddNote(note string) {
	b.disk.AddNote(note)
}

// WithAddNote スタートアップスクリプト 追加
func (b *CommonServerBuilder) WithAddNote(note string) *CommonServerBuilder {
	b.AddNote(note)
	return b
}

// ClearNotes スタートアップスクリプト クリア
func (b *CommonServerBuilder) ClearNotes() {
	b.disk.ClearNotes()
}

// WithEmptyNotes スタートアップスクリプト クリア
func (b *CommonServerBuilder) WithEmptyNotes() *CommonServerBuilder {
	b.ClearNotes()
	return b
}

// GetNotes スタートアップスクリプト 取得
func (b *CommonServerBuilder) GetNotes() []string {
	return b.disk.GetNotes()
}

// AddNoteID スタートアップスクリプト 追加
func (b *CommonServerBuilder) AddNoteID(noteID int64) {
	b.disk.AddNoteID(noteID)
}

// WithAddNoteID スタートアップスクリプト 追加
func (b *CommonServerBuilder) WithAddNoteID(noteID int64) *CommonServerBuilder {
	b.AddNoteID(noteID)
	return b
}

// ClearNoteIDs スタートアップスクリプト クリア
func (b *CommonServerBuilder) ClearNoteIDs() {
	b.disk.ClearNoteIDs()
}

// WithEmptyNoteIDs スタートアップスクリプト クリア
func (b *CommonServerBuilder) WithEmptyNoteIDs() *CommonServerBuilder {
	b.disk.ClearNoteIDs()
	return b
}

// GetNoteIDs スタートアップスクリプトID 取得
func (b *CommonServerBuilder) GetNoteIDs() []int64 {
	return b.disk.GetNoteIDs()
}

// IsSSHKeysEphemeral ディスク作成後の公開鍵削除フラグ 取得
func (b *CommonServerBuilder) IsSSHKeysEphemeral() bool {
	return b.disk.IsSSHKeysEphemeral()
}

// SetSSHKeysEphemeral ディスク作成後の公開鍵削除フラグ 設定
func (b *CommonServerBuilder) SetSSHKeysEphemeral(isEphemeral bool) {
	b.disk.SetSSHKeysEphemeral(isEphemeral)
}

// WithSSHKeysEphemeral ディスク作成後の公開鍵削除フラグ 設定
func (b *CommonServerBuilder) WithSSHKeysEphemeral(isEphemeral bool) *CommonServerBuilder {
	b.SetSSHKeysEphemeral(isEphemeral)
	return b
}

// IsNotesEphemeral ディスク作成後のスタートアップスクリプト削除フラグ 取得
func (b *CommonServerBuilder) IsNotesEphemeral() bool {
	return b.disk.IsNotesEphemeral()
}

// SetNotesEphemeral ディスク作成後のスタートアップスクリプト削除フラグ 設定
func (b *CommonServerBuilder) SetNotesEphemeral(isEphemeral bool) {
	b.disk.SetNotesEphemeral(isEphemeral)
}

// WithNotesEphemeral ディスク作成後のスタートアップスクリプト削除フラグ 設定
func (b *CommonServerBuilder) WithNotesEphemeral(isEphemeral bool) *CommonServerBuilder {
	b.SetNotesEphemeral(isEphemeral)
	return b
}

// GetGenerateSSHKeyName SSHキー生成 名称 取得
func (b *CommonServerBuilder) GetGenerateSSHKeyName() string {
	return b.disk.GetGenerateSSHKeyName()
}

// SetGenerateSSHKeyName SSHキー生成 名称 設定
func (b *CommonServerBuilder) SetGenerateSSHKeyName(name string) {
	b.disk.SetGenerateSSHKeyName(name)
}

// WithGenerateSSHKeyName SSHキー生成 名称 設定
func (b *CommonServerBuilder) WithGenerateSSHKeyName(name string) *CommonServerBuilder {
	b.disk.SetGenerateSSHKeyName(name)
	return b
}

// GetGenerateSSHKeyPassPhrase SSHキー生成 パスフレーズ 取得
func (b *CommonServerBuilder) GetGenerateSSHKeyPassPhrase() string {
	return b.disk.GetGenerateSSHKeyPassPhrase()
}

// SetGenerateSSHKeyPassPhrase SSHキー生成 パスフレーズ 設定
func (b *CommonServerBuilder) SetGenerateSSHKeyPassPhrase(pass string) {
	b.disk.SetGenerateSSHKeyPassPhrase(pass)
}

// WithGenerateSSHKeyPassPhrase SSHキー生成 パスフレーズ 設定
func (b *CommonServerBuilder) WithGenerateSSHKeyPassPhrase(pass string) *CommonServerBuilder {
	b.disk.SetGenerateSSHKeyPassPhrase(pass)
	return b
}

// GetGenerateSSHKeyDescription SSHキー生成 説明 取得
func (b *CommonServerBuilder) GetGenerateSSHKeyDescription() string {
	return b.disk.GetGenerateSSHKeyDescription()
}

// SetGenerateSSHKeyDescription SSHキー生成 説明 設定
func (b *CommonServerBuilder) SetGenerateSSHKeyDescription(desc string) {
	b.disk.SetGenerateSSHKeyDescription(desc)
}

// WithGenerateSSHKeyDescription SSHキー生成 説明 設定
func (b *CommonServerBuilder) WithGenerateSSHKeyDescription(desc string) *CommonServerBuilder {
	b.disk.SetGenerateSSHKeyDescription(desc)
	return b
}

/*---------------------------------------------------------
  for event handler
---------------------------------------------------------*/

// SetEventHandler イベントハンドラ 設定
func (b *CommonServerBuilder) SetEventHandler(event ServerBuildEvents, handler ServerBuildEventHandler) {
	b.buildEventHandlers[event] = handler
}

// WithEventHandler イベントハンドラ 設定
func (b *CommonServerBuilder) WithEventHandler(event ServerBuildEvents, handler ServerBuildEventHandler) *CommonServerBuilder {
	b.SetEventHandler(event, handler)
	return b
}

// ClearEventHandler イベントハンドラ クリア
func (b *CommonServerBuilder) ClearEventHandler(event ServerBuildEvents) {
	delete(b.buildEventHandlers, event)
}

// WithEmptyEventHandler イベントハンドラ クリア
func (b *CommonServerBuilder) WithEmptyEventHandler(event ServerBuildEvents) *CommonServerBuilder {
	b.ClearEventHandler(event)
	return b
}

// GetEventHandler イベントハンドラ 取得
func (b *CommonServerBuilder) GetEventHandler(event ServerBuildEvents) *ServerBuildEventHandler {
	if handler, ok := b.buildEventHandlers[event]; ok {
		return &handler
	}
	return nil
}

// SetDiskEventHandler ディスクイベントハンドラ 設定
func (b *CommonServerBuilder) SetDiskEventHandler(event DiskBuildEvents, handler DiskBuildEventHandler) {
	b.disk.SetEventHandler(event, handler)
}

// WithDiskEventHandler ディスクイベントハンドラ 設定
func (b *CommonServerBuilder) WithDiskEventHandler(event DiskBuildEvents, handler DiskBuildEventHandler) *CommonServerBuilder {
	b.SetDiskEventHandler(event, handler)
	return b
}

// ClearDiskEventHandler ディスクイベントハンドラ クリア
func (b *CommonServerBuilder) ClearDiskEventHandler(event DiskBuildEvents) {
	b.disk.ClearEventHandler(event)
}

// WithEmptyDiskEventHandler ディスクイベントハンドラ クリア
func (b *CommonServerBuilder) WithEmptyDiskEventHandler(event DiskBuildEvents) *CommonServerBuilder {
	b.ClearDiskEventHandler(event)
	return b
}

// GetDiskEventHandler ディスクイベントハンドラ 取得
func (b *CommonServerBuilder) GetDiskEventHandler(event DiskBuildEvents) *DiskBuildEventHandler {
	return b.disk.GetEventHandler(event)
}

/*---------------------------------------------------------
  for additional disks
---------------------------------------------------------*/

// AddAdditionalDisk 追加ディスク 追加
func (b *CommonServerBuilder) AddAdditionalDisk(diskBuilder *DiskBuilder) {
	b.additionalDisks = append(b.additionalDisks, diskBuilder)
}

// WithAddAdditionalDisk 追加ディスク 追加
func (b *CommonServerBuilder) WithAddAdditionalDisk(diskBuilder *DiskBuilder) *CommonServerBuilder {
	b.AddAdditionalDisk(diskBuilder)
	return b
}

// ClearAdditionalDisks 追加ディスク クリア
func (b *CommonServerBuilder) ClearAdditionalDisks() *CommonServerBuilder {
	b.additionalDisks = []*DiskBuilder{}
	return b
}

// WithEmptyAdditionalDisks 追加ディスク クリア
func (b *CommonServerBuilder) WithEmptyAdditionalDisks() *CommonServerBuilder {
	b.ClearAdditionalDisks()
	return b
}

// GetAdditionalDisks 追加ディスク 取得
func (b *CommonServerBuilder) GetAdditionalDisks() []*DiskBuilder {
	return b.additionalDisks
}
