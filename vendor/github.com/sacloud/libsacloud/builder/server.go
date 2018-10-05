package builder

import (
	"fmt"

	"github.com/sacloud/libsacloud/api"
	"github.com/sacloud/libsacloud/sacloud"
)

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

	// type info
	hasCommonProperty           bool
	hasNetworkInterfaceProperty bool
	hasDiskProperty             bool
	hasAdditionalDiskProperty   bool
	hasServerEventProperty      bool
	hasDiskEventProperty        bool
	hasDiskSourceProperty       bool
	hasDiskEditProperty         bool
}

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
  Inner functions
---------------------------------------------------------*/

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

	if !b.canAutoBoot() {
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
	plan, err := b.client.Product.Server.GetBySpec(b.core, b.memory, sacloud.PlanDefault)
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

	if b.bootAfterCreate && b.canAutoBoot() {
		s.SetWaitDiskMigration(true)
		b.disk.SetAutoBoot(true)
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
	if b.bootAfterCreate && b.canAutoBoot() {
		if err := b.client.Server.SleepUntilUp(b.currentBuildResult.Server.ID, b.client.DefaultTimeoutDuration); err != nil {
			return err
		}
		// refresh CurrentBildResult.Server
		s, err := b.client.Server.Read(b.currentBuildResult.Server.ID)
		if err != nil {
			return err
		}
		b.currentBuildResult.Server = s
	} else {
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

func (b *serverBuilder) canAutoBoot() bool {
	if b.disk == nil ||
		len(b.additionalDisks) > 0 ||
		b.isoImageID > 0 ||
		len(b.packetFilterIDs) > 0 {

		return false
	}

	return true
}

/*---------------------------------------------------------
  server builder common properties
---------------------------------------------------------*/

// GetServerName サーバー名 取得
func (b *serverBuilder) GetServerName() string {
	return b.serverName
}

// SetServerName サーバー名 設定
func (b *serverBuilder) SetServerName(serverName string) {
	b.serverName = serverName
}

// GetCore CPUコア数 取得
func (b *serverBuilder) GetCore() int {
	return b.core
}

// SetCore CPUコア数 設定
func (b *serverBuilder) SetCore(core int) {
	b.core = core
}

// GetMemory メモリサイズ(GB単位) 取得
func (b *serverBuilder) GetMemory() int {
	return b.memory
}

// SetMemory メモリサイズ(GB単位) 設定
func (b *serverBuilder) SetMemory(memory int) {
	b.memory = memory
}

// GetInterfaceDriver インターフェースドライバ 取得
func (b *serverBuilder) GetInterfaceDriver() sacloud.EInterfaceDriver {
	return b.interfaceDriver
}

// SetInterfaceDriver インターフェースドライバ 設定
func (b *serverBuilder) SetInterfaceDriver(interfaceDriver sacloud.EInterfaceDriver) {
	b.interfaceDriver = interfaceDriver
}

// GetDescription 説明 取得
func (b *serverBuilder) GetDescription() string {
	return b.description
}

// SetDescription 説明 設定
func (b *serverBuilder) SetDescription(description string) {
	b.description = description
}

// GetIconID アイコンID 取得
func (b *serverBuilder) GetIconID() int64 {
	return b.iconID
}

// SetIconID アイコンID 設定
func (b *serverBuilder) SetIconID(iconID int64) {
	b.iconID = iconID
}

// GetPrivateHostID 専有ホストID 取得
func (b *serverBuilder) GetPrivateHostID() int64 {
	return b.privateHostID
}

// SetPrivateHostID 専有ホストID 設定
func (b *serverBuilder) SetPrivateHostID(privateHostID int64) {
	b.privateHostID = privateHostID
}

// IsBootAfterCreate サーバー作成後すぐに起動フラグ 取得
func (b *serverBuilder) IsBootAfterCreate() bool {
	return b.bootAfterCreate
}

// SetBootAfterCreate サーバー作成後すぐに起動フラグ 設定
func (b *serverBuilder) SetBootAfterCreate(bootAfterCreate bool) {
	b.bootAfterCreate = bootAfterCreate
}

// GetTags タグ 取得
func (b *serverBuilder) GetTags() []string {
	return b.tags
}

// SetTags タグ 設定
func (b *serverBuilder) SetTags(tags []string) {
	b.tags = tags
}

// GetISOImageID ISOイメージ(CDROM)ID 取得
func (b *serverBuilder) GetISOImageID() int64 {
	return b.isoImageID
}

// SetISOImageID ISOイメージ(CDROM)ID 設定
func (b *serverBuilder) SetISOImageID(id int64) {
	b.isoImageID = id
}

/*---------------------------------------------------------
  for nic functioms
---------------------------------------------------------*/

// ClearNICConnections NIC接続設定 クリア
func (b *serverBuilder) ClearNICConnections() {
	b.nicConnections = nil
	b.disk.SetIPAddress("")
	b.disk.SetNetworkMaskLen(0)
	b.disk.SetDefaultRoute("")
}

// AddPublicNWConnectedNIC 共有セグメントへの接続追加(注:共有セグメントはeth0のみ接続可能)
func (b *serverBuilder) AddPublicNWConnectedNIC() {
	b.nicConnections = append(b.nicConnections, "shared")
}

// AddExistsSwitchConnectedNIC スイッチ or ルーター+スイッチへの接続追加(注:ルーター+スイッチはeth0のみ接続可能)
func (b *serverBuilder) AddExistsSwitchConnectedNIC(switchID string) {
	b.nicConnections = append(b.nicConnections, switchID)
}

// AddDisconnectedNIC 切断されたNIC追加
func (b *serverBuilder) AddDisconnectedNIC() {
	b.nicConnections = append(b.nicConnections, "")
}

// GetPacketFilterIDs パケットフィルタID 取得
func (b *serverBuilder) GetPacketFilterIDs() []int64 {
	return b.packetFilterIDs
}

// SetPacketFilterIDs パケットフィルタID 設定
func (b *serverBuilder) SetPacketFilterIDs(ids []int64) {
	b.packetFilterIDs = ids
}

/*---------------------------------------------------------
  for disk properties
---------------------------------------------------------*/

// GetDiskSize ディスクサイズ(GB単位) 取得
func (b *serverBuilder) GetDiskSize() int {
	return b.disk.GetSize()
}

// SetDiskSize ディスクサイズ(GB単位) 設定
func (b *serverBuilder) SetDiskSize(diskSize int) {
	b.disk.SetSize(diskSize)
}

// GetDistantFrom ストレージ隔離対象ディスク 取得
func (b *serverBuilder) GetDistantFrom() []int64 {
	return b.disk.GetDistantFrom()
}

// SetDistantFrom ストレージ隔離対象ディスク 設定
func (b *serverBuilder) SetDistantFrom(distantFrom []int64) {
	b.disk.SetDistantFrom(distantFrom)
}

// AddDistantFrom ストレージ隔離対象ディスク 追加
func (b *serverBuilder) AddDistantFrom(diskID int64) {
	b.disk.AddDistantFrom(diskID)
}

// ClearDistantFrom ストレージ隔離対象ディスク クリア
func (b *serverBuilder) ClearDistantFrom() {
	b.disk.ClearDistantFrom()
}

// GetDiskPlanID ディスクプラン(SSD/HDD) 取得
func (b *serverBuilder) GetDiskPlanID() sacloud.DiskPlanID {
	return b.disk.GetPlanID()
}

// SetDiskPlanID ディスクプラン(SSD/HDD) 設定
func (b *serverBuilder) SetDiskPlanID(diskPlanID sacloud.DiskPlanID) {
	b.disk.SetPlanID(diskPlanID)
}

// SetDiskPlan ディスクプラン(ssd/hdd) 設定(文字列から)
func (b *serverBuilder) SetDiskPlan(plan string) {
	b.disk.SetPlan(plan)
}

// GetDiskConnection ディスク接続方法(VirtIO/IDE) 取得
func (b *serverBuilder) GetDiskConnection() sacloud.EDiskConnection {
	return b.disk.GetConnection()
}

// SetDiskConnection ディスク接続方法(VirtIO/IDE) 設定
func (b *serverBuilder) SetDiskConnection(diskConnection sacloud.EDiskConnection) {
	b.disk.SetConnection(diskConnection)
}

/*---------------------------------------------------------
  for disk edit properties
---------------------------------------------------------*/

// GetSourceArchiveID ソースアーカイブID 取得
func (b *serverBuilder) GetSourceArchiveID() int64 {
	return b.disk.GetSourceArchiveID()
}

// GetSourceDiskID ソースディスクID 設定
func (b *serverBuilder) GetSourceDiskID() int64 {
	return b.disk.GetSourceDiskID()
}

// GetPassword パスワード 取得
func (b *serverBuilder) GetPassword() string {
	return b.disk.GetPassword()
}

// SetPassword パスワード 設定
func (b *serverBuilder) SetPassword(password string) {
	b.disk.SetPassword(password)
}

// GetHostName ホスト名 取得
func (b *serverBuilder) GetHostName() string {
	return b.disk.GetHostName()
}

// SetHostName ホスト名 設定
func (b *serverBuilder) SetHostName(hostName string) {
	b.disk.SetHostName(hostName)
}

// GetIPAddress IPアドレス 取得
func (b *serverBuilder) GetIPAddress() string {
	return b.disk.ipAddress
}

// SetIPAddress IPアドレス 設定
func (b *serverBuilder) SetIPAddress(ip string) {
	b.disk.ipAddress = ip
}

// GetNetworkMaskLen ネットワークマスク長 取得
func (b *serverBuilder) GetNetworkMaskLen() int {
	return b.disk.networkMaskLen
}

// SetNetworkMaskLen ネットワークマスク長 設定
func (b *serverBuilder) SetNetworkMaskLen(masklen int) {
	b.disk.networkMaskLen = masklen
}

// GetDefaultRoute デフォルトルート 取得
func (b *serverBuilder) GetDefaultRoute() string {
	return b.disk.defaultRoute
}

// SetDefaultRoute デフォルトルート 設定
func (b *serverBuilder) SetDefaultRoute(route string) {
	b.disk.defaultRoute = route
}

// IsDisablePWAuth パスワード認証無効化フラグ 取得
func (b *serverBuilder) IsDisablePWAuth() bool {
	return b.disk.IsDisablePWAuth()
}

// SetDisablePWAuth パスワード認証無効化フラグ 設定
func (b *serverBuilder) SetDisablePWAuth(disable bool) {
	b.disk.SetDisablePWAuth(disable)
}

// AddSSHKey 公開鍵 追加
func (b *serverBuilder) AddSSHKey(sshKey string) {
	b.disk.AddSSHKey(sshKey)
}

// ClearSSHKey 公開鍵 クリア
func (b *serverBuilder) ClearSSHKey() {
	b.disk.ClearSSHKey()
}

// GetSSHKeys 公開鍵 取得
func (b *serverBuilder) GetSSHKeys() []string {
	return b.disk.GetSSHKeys()
}

// AddSSHKeyID 公開鍵ID 追加
func (b *serverBuilder) AddSSHKeyID(sshKeyID int64) {
	b.disk.AddSSHKeyID(sshKeyID)
}

// ClearSSHKeyIDs 公開鍵ID クリア
func (b *serverBuilder) ClearSSHKeyIDs() {
	b.disk.ClearSSHKeyIDs()
}

// GetSSHKeyIds 公開鍵ID 取得
func (b *serverBuilder) GetSSHKeyIds() []int64 {
	return b.disk.GetSSHKeyIds()
}

// AddNote スタートアップスクリプト 追加
func (b *serverBuilder) AddNote(note string) {
	b.disk.AddNote(note)
}

// ClearNotes スタートアップスクリプト クリア
func (b *serverBuilder) ClearNotes() {
	b.disk.ClearNotes()
}

// GetNotes スタートアップスクリプト 取得
func (b *serverBuilder) GetNotes() []string {
	return b.disk.GetNotes()
}

// AddNoteID スタートアップスクリプト 追加
func (b *serverBuilder) AddNoteID(noteID int64) {
	b.disk.AddNoteID(noteID)
}

// ClearNoteIDs スタートアップスクリプト クリア
func (b *serverBuilder) ClearNoteIDs() {
	b.disk.ClearNoteIDs()
}

// GetNoteIDs スタートアップスクリプトID 取得
func (b *serverBuilder) GetNoteIDs() []int64 {
	return b.disk.GetNoteIDs()
}

// IsSSHKeysEphemeral ディスク作成後の公開鍵削除フラグ 取得
func (b *serverBuilder) IsSSHKeysEphemeral() bool {
	return b.disk.IsSSHKeysEphemeral()
}

// SetSSHKeysEphemeral ディスク作成後の公開鍵削除フラグ 設定
func (b *serverBuilder) SetSSHKeysEphemeral(isEphemeral bool) {
	b.disk.SetSSHKeysEphemeral(isEphemeral)
}

// IsNotesEphemeral ディスク作成後のスタートアップスクリプト削除フラグ 取得
func (b *serverBuilder) IsNotesEphemeral() bool {
	return b.disk.IsNotesEphemeral()
}

// SetNotesEphemeral ディスク作成後のスタートアップスクリプト削除フラグ 設定
func (b *serverBuilder) SetNotesEphemeral(isEphemeral bool) {
	b.disk.SetNotesEphemeral(isEphemeral)
}

// GetGenerateSSHKeyName SSHキー生成 名称 取得
func (b *serverBuilder) GetGenerateSSHKeyName() string {
	return b.disk.GetGenerateSSHKeyName()
}

// SetGenerateSSHKeyName SSHキー生成 名称 設定
func (b *serverBuilder) SetGenerateSSHKeyName(name string) {
	b.disk.SetGenerateSSHKeyName(name)
}

// GetGenerateSSHKeyPassPhrase SSHキー生成 パスフレーズ 取得
func (b *serverBuilder) GetGenerateSSHKeyPassPhrase() string {
	return b.disk.GetGenerateSSHKeyPassPhrase()
}

// SetGenerateSSHKeyPassPhrase SSHキー生成 パスフレーズ 設定
func (b *serverBuilder) SetGenerateSSHKeyPassPhrase(pass string) {
	b.disk.SetGenerateSSHKeyPassPhrase(pass)
}

// GetGenerateSSHKeyDescription SSHキー生成 説明 取得
func (b *serverBuilder) GetGenerateSSHKeyDescription() string {
	return b.disk.GetGenerateSSHKeyDescription()
}

// SetGenerateSSHKeyDescription SSHキー生成 説明 設定
func (b *serverBuilder) SetGenerateSSHKeyDescription(desc string) {
	b.disk.SetGenerateSSHKeyDescription(desc)
}

/*---------------------------------------------------------
  for event handler
---------------------------------------------------------*/

// SetEventHandler イベントハンドラ 設定
func (b *serverBuilder) SetEventHandler(event ServerBuildEvents, handler ServerBuildEventHandler) {
	b.buildEventHandlers[event] = handler
}

// ClearEventHandler イベントハンドラ クリア
func (b *serverBuilder) ClearEventHandler(event ServerBuildEvents) {
	delete(b.buildEventHandlers, event)
}

// GetEventHandler イベントハンドラ 取得
func (b *serverBuilder) GetEventHandler(event ServerBuildEvents) *ServerBuildEventHandler {
	if handler, ok := b.buildEventHandlers[event]; ok {
		return &handler
	}
	return nil
}

// SetDiskEventHandler ディスクイベントハンドラ 設定
func (b *serverBuilder) SetDiskEventHandler(event DiskBuildEvents, handler DiskBuildEventHandler) {
	b.disk.SetEventHandler(event, handler)
}

// ClearDiskEventHandler ディスクイベントハンドラ クリア
func (b *serverBuilder) ClearDiskEventHandler(event DiskBuildEvents) {
	b.disk.ClearEventHandler(event)
}

// GetDiskEventHandler ディスクイベントハンドラ 取得
func (b *serverBuilder) GetDiskEventHandler(event DiskBuildEvents) *DiskBuildEventHandler {
	return b.disk.GetEventHandler(event)
}

/*---------------------------------------------------------
  for additional disks
---------------------------------------------------------*/

// AddAdditionalDisk 追加ディスク 追加
func (b *serverBuilder) AddAdditionalDisk(diskBuilder *DiskBuilder) {
	b.additionalDisks = append(b.additionalDisks, diskBuilder)
}

// ClearAdditionalDisks 追加ディスク クリア
func (b *serverBuilder) ClearAdditionalDisks() {
	b.additionalDisks = []*DiskBuilder{}
}

// GetAdditionalDisks 追加ディスク 取得
func (b *serverBuilder) GetAdditionalDisks() []*DiskBuilder {
	return b.additionalDisks
}

/*---------------------------------------------------------
  type info
---------------------------------------------------------*/

// HasCommonProperty 汎用プロパティを保持しているか
func (b *serverBuilder) HasCommonProperty() bool {
	return b.hasCommonProperty
}

// HasNetworkInterfaceProperty NIC関連プロパティを保持しているか
func (b *serverBuilder) HasNetworkInterfaceProperty() bool {
	return b.hasNetworkInterfaceProperty
}

// HasDiskProperty ディスク関連プロパティを保持しているか
func (b *serverBuilder) HasDiskProperty() bool {
	return b.hasDiskProperty
}

// HasAdditionalDiskProperty 追加ディスク関連プロパティを保持しているか
func (b *serverBuilder) HasAdditionalDiskProperty() bool {
	return b.hasAdditionalDiskProperty
}

// HasServerEventProperty サーバ構築イベント関連プロパティを保持しているか
func (b *serverBuilder) HasServerEventProperty() bool {
	return b.hasServerEventProperty
}

// HasDiskEventProperty ディスク構築イベント関連プロパティを保持しているか
func (b *serverBuilder) HasDiskEventProperty() bool {
	return b.hasDiskEditProperty
}

// HasDiskSourceProperty ディスクコピー元関連プロパティを保持しているか
func (b *serverBuilder) HasDiskSourceProperty() bool {
	return b.hasDiskSourceProperty
}

// HasDiskEditProperty ディスク修正関連プロパティを保持しているか
func (b *serverBuilder) HasDiskEditProperty() bool {
	return b.hasDiskEditProperty
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
