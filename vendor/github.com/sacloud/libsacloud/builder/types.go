package builder

import "github.com/sacloud/libsacloud/sacloud"

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

// CommonProperty ビルダー共通のプロパティ
type CommonProperty interface {
	// GetServerName サーバー名 取得
	GetServerName() string
	// SetServerName サーバー名 設定
	SetServerName(serverName string)

	// GetCore CPUコア数 取得
	GetCore() int
	// SetCore CPUコア数 設定
	SetCore(core int)

	// GetMemory メモリサイズ(GB単位) 取得
	GetMemory() int
	// SetMemory メモリサイズ(GB単位) 設定
	SetMemory(memory int)

	// GetInterfaceDriver インターフェースドライバ 取得
	GetInterfaceDriver() sacloud.EInterfaceDriver
	// SetInterfaceDriver インターフェースドライバ 設定
	SetInterfaceDriver(interfaceDriver sacloud.EInterfaceDriver)

	// GetDescription 説明 取得
	GetDescription() string
	// SetDescription 説明 設定
	SetDescription(description string)

	// GetIconID アイコンID 取得
	GetIconID() int64
	// SetIconID アイコンID 設定
	SetIconID(iconID int64)

	// GetPrivateHostID アイコンID 取得
	GetPrivateHostID() int64
	// SetPrivateHostID アイコンID 設定
	SetPrivateHostID(privateHostID int64)

	// IsBootAfterCreate サーバー作成後すぐに起動フラグ 取得
	IsBootAfterCreate() bool
	// SetBootAfterCreate サーバー作成後すぐに起動フラグ 設定
	SetBootAfterCreate(bootAfterCreate bool)

	// GetTags タグ 取得
	GetTags() []string
	// SetTags タグ 設定
	SetTags(tags []string)

	// GetISOImageID ISOイメージ(CDROM)ID 取得
	GetISOImageID() int64
	// SetISOImageID ISOイメージ(CDROM)ID 設定
	SetISOImageID(id int64)
}

// NetworkInterfaceProperty NIC関連プロパティ
type NetworkInterfaceProperty interface {
	// ClearNICConnections NIC接続設定 クリア
	ClearNICConnections()

	// AddPublicNWConnectedNIC 共有セグメントへの接続追加(注:共有セグメントはeth0のみ接続可能)
	AddPublicNWConnectedNIC()

	// AddExistsSwitchConnectedNIC スイッチ or ルーター+スイッチへの接続追加(注:ルーター+スイッチはeth0のみ接続可能)
	AddExistsSwitchConnectedNIC(switchID string)

	// AddDisconnectedNIC 切断されたNIC追加
	AddDisconnectedNIC()

	// GetPacketFilterIDs パケットフィルタID 取得
	GetPacketFilterIDs() []int64
	// SetPacketFilterIDs パケットフィルタID 設定
	SetPacketFilterIDs(ids []int64)
}

// DiskProperty ディスク関連プロパティ
type DiskProperty interface {
	// GetDiskSize ディスクサイズ(GB単位) 取得
	GetDiskSize() int
	// SetDiskSize ディスクサイズ(GB単位) 設定
	SetDiskSize(diskSize int)

	// GetDistantFrom ストレージ隔離対象ディスク 取得
	GetDistantFrom() []int64
	// SetDistantFrom ストレージ隔離対象ディスク 設定
	SetDistantFrom(distantFrom []int64)
	// AddDistantFrom ストレージ隔離対象ディスク 追加
	AddDistantFrom(diskID int64)
	// ClearDistantFrom ストレージ隔離対象ディスク クリア
	ClearDistantFrom()

	// GetDiskPlanID ディスクプラン(SSD/HDD) 取得
	GetDiskPlanID() sacloud.DiskPlanID
	// SetDiskPlanID ディスクプラン(SSD/HDD) 設定
	SetDiskPlanID(diskPlanID sacloud.DiskPlanID)
	// WithDiskPlanID ディスクプラン(SSD/HDD) 設定
	SetDiskPlan(plan string)

	// GetDiskConnection ディスク接続方法(VirtIO/IDE) 取得
	GetDiskConnection() sacloud.EDiskConnection
	// SetDiskConnection ディスク接続方法(VirtIO/IDE) 設定
	SetDiskConnection(diskConnection sacloud.EDiskConnection)
}

// AdditionalDiskProperty 追加ディスクプロパティ
type AdditionalDiskProperty interface {
	// AddAdditionalDisk 追加ディスク 追加
	AddAdditionalDisk(diskBuilder *DiskBuilder)
	// ClearAdditionalDisks 追加ディスク クリア
	ClearAdditionalDisks()
	// GetAdditionalDisks 追加ディスク 取得
	GetAdditionalDisks() []*DiskBuilder
}

// ServerEventProperty サーバ構築時イベントプロパティ
type ServerEventProperty interface {
	// GetEventHandler イベントハンドラ 取得
	GetEventHandler(event ServerBuildEvents) *ServerBuildEventHandler
	// SetEventHandler イベントハンドラ 設定
	SetEventHandler(event ServerBuildEvents, handler ServerBuildEventHandler)
	// ClearEventHandler イベントハンドラ クリア
	ClearEventHandler(event ServerBuildEvents)
}

// DiskEventProperty ディスク作成時イベントプロパティ
type DiskEventProperty interface {
	// GetDiskEventHandler ディスクイベントハンドラ 取得
	GetDiskEventHandler(event DiskBuildEvents) *DiskBuildEventHandler
	// SetDiskEventHandler ディスクイベントハンドラ 設定
	SetDiskEventHandler(event DiskBuildEvents, handler DiskBuildEventHandler)
	// ClearDiskEventHandler ディスクイベントハンドラ クリア
	ClearDiskEventHandler(event DiskBuildEvents)
}

// DiskSourceProperty コピー元アーカイブ/ディスクプロパティ
type DiskSourceProperty interface {
	// GetSourceArchiveID ソースアーカイブID 取得
	GetSourceArchiveID() int64

	// GetSourceDiskID ソースディスクID 設定
	GetSourceDiskID() int64
}

// DiskEditProperty ディスクの修正関連プロパティ
type DiskEditProperty interface {

	// GetPassword パスワード 取得
	GetPassword() string
	// SetPassword パスワード 設定
	SetPassword(password string)

	// GetHostName ホスト名 取得
	GetHostName() string
	// SetHostName ホスト名 設定
	SetHostName(hostName string)

	// GetIPAddress IPアドレス 取得
	GetIPAddress() string
	// SetIPAddress IPアドレス 設定
	SetIPAddress(ipAddress string)

	// NetworkMaskLen ネットワークマスク長 取得
	GetNetworkMaskLen() int
	// SetNetworkMaskLen ネットワークマスク長 設定
	SetNetworkMaskLen(maskLen int)

	// GetDefaultRoute デフォルトルート 取得
	GetDefaultRoute() string
	// SetDefaultRoute デフォルトルート 設定
	SetDefaultRoute(route string)

	// IsDisablePWAuth パスワード認証無効化フラグ 取得
	IsDisablePWAuth() bool
	// SetDisablePWAuth パスワード認証無効化フラグ 設定
	SetDisablePWAuth(disable bool)

	// GetSSHKeys 公開鍵 取得
	GetSSHKeys() []string
	// GetSSHKeyIds 公開鍵ID 取得
	GetSSHKeyIds() []int64
	// AddSSHKey 公開鍵 追加
	AddSSHKey(sshKey string)
	// AddSSHKeyID 公開鍵ID 追加
	AddSSHKeyID(sshKeyID int64)
	// ClearSSHKey 公開鍵 クリア
	ClearSSHKey()
	// ClearSSHKeyIDs 公開鍵ID クリア
	ClearSSHKeyIDs()

	// AddNote スタートアップスクリプト 追加
	AddNote(note string)
	// ClearNotes スタートアップスクリプト クリア
	ClearNotes()
	// GetNotes スタートアップスクリプト 取得
	GetNotes() []string
	// AddNoteID スタートアップスクリプト 追加
	AddNoteID(noteID int64)
	// ClearNoteIDs スタートアップスクリプト クリア
	ClearNoteIDs()
	// GetNoteIDs スタートアップスクリプトID 取得
	GetNoteIDs() []int64

	// IsSSHKeysEphemeral ディスク作成後の公開鍵削除フラグ 取得
	IsSSHKeysEphemeral() bool
	// SetSSHKeysEphemeral ディスク作成後の公開鍵削除フラグ 設定
	SetSSHKeysEphemeral(isEphemeral bool)

	// IsNotesEphemeral ディスク作成後のスタートアップスクリプト削除フラグ 取得
	IsNotesEphemeral() bool
	// SetNotesEphemeral ディスク作成後のスタートアップスクリプト削除フラグ 設定
	SetNotesEphemeral(isEphemeral bool)

	// GetGenerateSSHKeyName SSHキー生成 名称 取得
	GetGenerateSSHKeyName() string
	// SetGenerateSSHKeyName SSHキー生成 名称 設定
	SetGenerateSSHKeyName(name string)

	// GetGenerateSSHKeyPassPhrase SSHキー生成 パスフレーズ 取得
	GetGenerateSSHKeyPassPhrase() string
	// SetGenerateSSHKeyPassPhrase SSHキー生成 パスフレーズ 設定
	SetGenerateSSHKeyPassPhrase(pass string)

	// GetGenerateSSHKeyDescription SSHキー生成 説明 取得
	GetGenerateSSHKeyDescription() string
	// SetGenerateSSHKeyDescription SSHキー生成 説明 設定
	SetGenerateSSHKeyDescription(desc string)
}

//-----------------------------------------------------------------------------
// Builders
//-----------------------------------------------------------------------------

// Builder ビルダー基本インターフェース
type Builder interface {
	Build() (*ServerBuildResult, error)
	HasCommonProperty() bool
	HasNetworkInterfaceProperty() bool
	HasDiskProperty() bool
	HasAdditionalDiskProperty() bool
	HasServerEventProperty() bool
	HasDiskEventProperty() bool
	HasDiskSourceProperty() bool
	HasDiskEditProperty() bool
}

// BlankDiskServerBuilder ブランクディスクを利用して構築を行うサーバービルダー
//
// 空のディスクを持ちます。基本的なディスク設定やディスクの追加に対応していますが、ディスクの修正機能には非対応です。
type BlankDiskServerBuilder interface {
	Builder
	CommonProperty
	NetworkInterfaceProperty
	DiskProperty
	ServerEventProperty
	DiskEventProperty
	AdditionalDiskProperty
}

// CommonServerBuilder 既存のアーカイブ or ディスクを利用して構築を行うサーバービルダー
//
// 基本的なディスク設定やディスクの追加、ディスクの修正機能に対応しています。
// ただし、ディスクの修正機能はソースアーカイブ or ソースディスクが対応していない場合は
// さくらのクラウドAPIコール時にエラーとなるため、適切にハンドリングするように実装する必要があります。
type CommonServerBuilder interface {
	Builder
	CommonProperty
	NetworkInterfaceProperty
	DiskProperty
	DiskEditProperty
	DiskSourceProperty
	ServerEventProperty
	DiskEventProperty
	AdditionalDiskProperty
}

// ConnectDiskServerBuilder ブランクディスクを利用して構築を行うサーバービルダー
//
// すでに存在するディスクを持ちます。ディスクの追加に対応していますが、ディスクの修正機能には非対応です。
type ConnectDiskServerBuilder interface {
	Builder
	CommonProperty
	NetworkInterfaceProperty
	AdditionalDiskProperty
}

// DisklessServerBuilder ディスクレス サーバービルダー
//
// ディスクレスのサーバーを構築します。 ディスク関連の設定に非対応です。
type DisklessServerBuilder interface {
	Builder
	CommonProperty
	NetworkInterfaceProperty
	ServerEventProperty
}

// FixedUnixArchiveServerBuilder ディスクの修正不可なUNIX系パブリックアーカイブを利用して構築を行うサーバービルダー
//
// 基本的なディスク設定やディスクの追加に対応していますが、ディスクの修正機能には非対応です。
type FixedUnixArchiveServerBuilder interface {
	Builder
	CommonProperty
	NetworkInterfaceProperty
	DiskProperty
	DiskSourceProperty
	ServerEventProperty
	DiskEventProperty
	AdditionalDiskProperty
}

// PublicArchiveUnixServerBuilder Linux(Unix)系パブリックアーカイブを利用して構築を行うサーバービルダー
//
// 基本的なディスク設定やディスクの追加、ディスクの修正機能に対応しています。
type PublicArchiveUnixServerBuilder interface {
	Builder
	CommonProperty
	NetworkInterfaceProperty
	DiskProperty
	DiskSourceProperty
	DiskEditProperty
	ServerEventProperty
	DiskEventProperty
	AdditionalDiskProperty
}

// PublicArchiveWindowsServerBuilder Windows系パブリックアーカイブを利用して構築を行うサーバービルダー
//
// 基本的なディスク設定やディスクの追加に対応していますが、ディスクの修正機能には非対応です。
type PublicArchiveWindowsServerBuilder interface {
	Builder
	CommonProperty
	NetworkInterfaceProperty
	DiskProperty
	DiskSourceProperty
	ServerEventProperty
	DiskEventProperty
	AdditionalDiskProperty
}
