package builder

import (
	"fmt"
	"github.com/sacloud/libsacloud/api"
	"github.com/sacloud/libsacloud/sacloud"
	"time"
)

/**********************************************************
  Type : DiskBuildEvents
**********************************************************/

// DiskBuildEvents ディスク構築時イベント種別
type DiskBuildEvents int

const (
	// DiskBuildOnStart ディスク作成 開始
	DiskBuildOnStart DiskBuildEvents = iota

	// DiskBuildOnCreateSSHKeyBefore SSHキー作成 開始時
	DiskBuildOnCreateSSHKeyBefore

	// DiskBuildOnCreateSSHKeyAfter SSHキー作成 終了時
	DiskBuildOnCreateSSHKeyAfter

	// DiskBuildOnCreateNoteBefore スタートアップスクリプト作成 開始時
	DiskBuildOnCreateNoteBefore

	// DiskBuildOnCreateNoteAfter スタートアップスクリプト作成 終了時
	DiskBuildOnCreateNoteAfter

	// DiskBuildOnCreateDiskBefore ディスク作成 開始時
	DiskBuildOnCreateDiskBefore

	// DiskBuildOnCreateDiskAfter ディスク作成 終了時
	DiskBuildOnCreateDiskAfter

	// DiskBuildOnCleanupSSHKeyBefore SSHキークリーンアップ 開始時
	DiskBuildOnCleanupSSHKeyBefore

	// DiskBuildOnCleanupSSHKeyAfter SSHキークリーンアップ 終了時
	DiskBuildOnCleanupSSHKeyAfter

	// DiskBuildOnCleanupNoteBefore スタートアップスクリプトクリーンアップ 開始時
	DiskBuildOnCleanupNoteBefore

	// DiskBuildOnCleanupNoteAfter スタートアップスクリプトクリーンアップ 終了時
	DiskBuildOnCleanupNoteAfter

	// DiskBuildOnComplete ディスク作成 完了
	DiskBuildOnComplete
)

// DiskBuildEventHandler ディスク構築時イベントハンドラ型
type DiskBuildEventHandler func(value *DiskBuildValue, result *DiskBuildResult)

/**********************************************************
  Type : DiskBuilder
**********************************************************/

// DiskBuilder ディスクビルダー
type DiskBuilder struct {
	*baseBuilder
	buildEventHandlers map[DiskBuildEvents]DiskBuildEventHandler

	name            string
	size            int
	distantFrom     []int64
	planID          sacloud.DiskPlanID
	connection      sacloud.EDiskConnection
	sourceArchiveID int64
	sourceDiskID    int64
	description     string
	tags            []string
	iconID          int64
	serverID        int64

	ipAddress          string
	networkMaskLen     int
	defaultRoute       string
	password           string
	hostName           string
	disablePWAuth      bool
	sshKeys            []string
	isSSHKeysEphemeral bool
	sshKeyIDs          []int64
	notes              []string
	isNotesEphemeral   bool
	noteIDs            []int64

	// for sshkey generate
	generateSSHKeyName        string
	generateSSHKeyPassPhrase  string
	generateSSHKeyDescription string

	forceEditDisk bool // windowsなどの場合にディスクの修正を強制するためのフラグ

	autoBoot bool

	currentDiskBuildValue  *DiskBuildValue
	currentDiskBuildResult *DiskBuildResult
}

const (
	// DefaultDiskPlanID ディスクプラン(デフォルト値)
	DefaultDiskPlanID = sacloud.DiskPlanSSDID
	// DefaultDiskConnection ディスク接続方法(デフォルト値)
	DefaultDiskConnection = sacloud.DiskConnectionVirtio
	// DefaultDiskSize ディスクサイズ(デフォルト値)
	DefaultDiskSize = 20
	// DefaultDiskIsSSHKeysEphemeral ディスク作成後の公開鍵削除フラグ(デフォルト値)
	DefaultDiskIsSSHKeysEphemeral = true
	// DefaultDiskIsNotesEphemeral ディスク作成後のスタートアップスクリプト削除フラグ(デフォルト値)
	DefaultDiskIsNotesEphemeral = true
)

// Disk ディスクビルダーの作成
func Disk(client *api.Client, name string) *DiskBuilder {
	return &DiskBuilder{
		baseBuilder: &baseBuilder{
			client: client,
		},
		buildEventHandlers: map[DiskBuildEvents]DiskBuildEventHandler{},
		name:               name,
		size:               DefaultDiskSize,
		planID:             DefaultDiskPlanID,
		connection:         DefaultDiskConnection,
		isSSHKeysEphemeral: DefaultDiskIsSSHKeysEphemeral,
		isNotesEphemeral:   DefaultDiskIsNotesEphemeral,
	}
}

// GetName ディスク名 取得
func (b *DiskBuilder) GetName() string {
	return b.name
}

// SetName ディスク名 設定
func (b *DiskBuilder) SetName(name string) {
	b.name = name
}

// GetSize ディスクサイズ(GB単位) 取得
func (b *DiskBuilder) GetSize() int {
	return b.size
}

// SetSize ディスクサイズ(GB単位) 設定
func (b *DiskBuilder) SetSize(size int) {
	b.size = size
}

// GetDistantFrom ストレージ隔離対象ディスク 取得
func (b *DiskBuilder) GetDistantFrom() []int64 {
	return b.distantFrom
}

// SetDistantFrom ストレージ隔離対象ディスク 設定
func (b *DiskBuilder) SetDistantFrom(diskIDs []int64) {
	b.distantFrom = diskIDs
}

// AddDistantFrom ストレージ隔離対象ディスク 追加
func (b *DiskBuilder) AddDistantFrom(diskID int64) *DiskBuilder {
	b.distantFrom = append(b.distantFrom, diskID)
	return b
}

// ClearDistantFrom ストレージ隔離対象ディスク クリア
func (b *DiskBuilder) ClearDistantFrom() {
	b.distantFrom = []int64{}
}

// GetPlanID ディスクプラン(SSD/HDD) 取得
func (b *DiskBuilder) GetPlanID() sacloud.DiskPlanID {
	return b.planID
}

// SetPlanID ディスクプラン(SSD/HDD) 設定
func (b *DiskBuilder) SetPlanID(planID sacloud.DiskPlanID) {
	b.planID = planID
}

// SetPlan ディスクプラン(ssd/hdd) 設定(文字列から)
func (b *DiskBuilder) SetPlan(plan string) {
	switch plan {
	case "ssd":
		b.SetPlanID(sacloud.DiskPlanSSDID)
	case "hdd":
		b.SetPlanID(sacloud.DiskPlanHDDID)
	default:
		panic(fmt.Errorf("Invalid plan:%s", plan))
	}
}

// GetConnection ディスク接続方法(VirtIO/IDE) 取得
func (b *DiskBuilder) GetConnection() sacloud.EDiskConnection {
	return b.connection
}

// SetConnection ディスク接続方法(VirtIO/IDE) 設定
func (b *DiskBuilder) SetConnection(connection sacloud.EDiskConnection) {
	b.connection = connection
}

// GetSourceArchiveID ソースアーカイブID 取得
func (b *DiskBuilder) GetSourceArchiveID() int64 {
	return b.sourceArchiveID
}

// SetSourceArchiveID ソースアーカイブID 設定
func (b *DiskBuilder) SetSourceArchiveID(id int64) {
	b.sourceArchiveID = id
	b.sourceDiskID = 0
}

// GetSourceDiskID ソースディスクID 取得
func (b *DiskBuilder) GetSourceDiskID() int64 {
	return b.sourceDiskID
}

// SetSourceDiskID ソースディスクID 設定
func (b *DiskBuilder) SetSourceDiskID(id int64) {
	b.sourceArchiveID = 0
	b.sourceDiskID = id
}

// GetDescription 説明 取得
func (b *DiskBuilder) GetDescription() string {
	return b.description
}

// SetDescription 説明 設定
func (b *DiskBuilder) SetDescription(desc string) {
	b.description = desc
}

// GetTags タグ 取得
func (b *DiskBuilder) GetTags() []string {
	return b.tags
}

// SetTags タグ 設定
func (b *DiskBuilder) SetTags(tags []string) {
	b.tags = tags
}

// GetIconID アイコンID 取得
func (b *DiskBuilder) GetIconID() int64 {
	return b.iconID
}

// SetIconID アイコンID 設定
func (b *DiskBuilder) SetIconID(id int64) {
	b.iconID = id
}

// GetServerID サーバーID 取得
func (b *DiskBuilder) GetServerID() int64 {
	return b.serverID
}

// SetServerID サーバーID 設定
func (b *DiskBuilder) SetServerID(id int64) {
	b.serverID = id
}

// GetIPAddress IPアドレス 取得
func (b *DiskBuilder) GetIPAddress() string {
	return b.ipAddress
}

// SetIPAddress IPアドレス 設定
func (b *DiskBuilder) SetIPAddress(ip string) {
	b.ipAddress = ip
}

// GetNetworkMaskLen ネットワークマスク長 取得
func (b *DiskBuilder) GetNetworkMaskLen() int {
	return b.networkMaskLen
}

// SetNetworkMaskLen ネットワークマスク長 設定
func (b *DiskBuilder) SetNetworkMaskLen(masklen int) {
	b.networkMaskLen = masklen
}

// GetDefaultRoute デフォルトルート 取得
func (b *DiskBuilder) GetDefaultRoute() string {
	return b.defaultRoute
}

// SetDefaultRoute デフォルトルート 設定
func (b *DiskBuilder) SetDefaultRoute(route string) {
	b.defaultRoute = route
}

// GetPassword パスワード 取得
func (b *DiskBuilder) GetPassword() string {
	return b.password
}

// SetPassword パスワード 設定
func (b *DiskBuilder) SetPassword(password string) {
	b.password = password
}

// GetHostName ホスト名 取得
func (b *DiskBuilder) GetHostName() string {
	return b.hostName
}

// SetHostName ホスト名 設定
func (b *DiskBuilder) SetHostName(name string) {
	b.hostName = name
}

// IsDisablePWAuth パスワード認証無効化フラグ 取得
func (b *DiskBuilder) IsDisablePWAuth() bool {
	return b.disablePWAuth
}

// SetDisablePWAuth パスワード認証無効化フラグ 設定
func (b *DiskBuilder) SetDisablePWAuth(disable bool) {
	b.disablePWAuth = disable
}

// AddSSHKeyID 公開鍵ID 追加
func (b *DiskBuilder) AddSSHKeyID(sshKeyID int64) {
	b.sshKeyIDs = append(b.sshKeyIDs, sshKeyID)
}

// ClearSSHKeyIDs 公開鍵ID クリア
func (b *DiskBuilder) ClearSSHKeyIDs() {
	b.sshKeyIDs = []int64{}
}

// GetSSHKeyIds 公開鍵ID 取得
func (b *DiskBuilder) GetSSHKeyIds() []int64 {
	return b.sshKeyIDs
}

// AddSSHKey 公開鍵 追加
func (b *DiskBuilder) AddSSHKey(sshKey string) {
	b.sshKeys = append(b.sshKeys, sshKey)
}

// ClearSSHKey 公開鍵 クリア
func (b *DiskBuilder) ClearSSHKey() {
	b.sshKeys = []string{}
}

// GetSSHKeys 公開鍵 取得
func (b *DiskBuilder) GetSSHKeys() []string {
	return b.sshKeys
}

// AddNote スタートアップスクリプト 追加
func (b *DiskBuilder) AddNote(note string) {
	b.notes = append(b.notes, note)
}

// ClearNotes スタートアップスクリプト クリア
func (b *DiskBuilder) ClearNotes() {
	b.notes = []string{}
}

// GetNotes スタートアップスクリプト 取得
func (b *DiskBuilder) GetNotes() []string {
	return b.notes
}

// AddNoteID スタートアップスクリプトID 追加
func (b *DiskBuilder) AddNoteID(noteID int64) {
	b.noteIDs = append(b.noteIDs, noteID)
}

// ClearNoteIDs スタートアップスクリプトID クリア
func (b *DiskBuilder) ClearNoteIDs() {
	b.noteIDs = []int64{}
}

// GetNoteIDs スタートアップスクリプトID 取得
func (b *DiskBuilder) GetNoteIDs() []int64 {
	return b.noteIDs
}

// IsSSHKeysEphemeral ディスク作成後の公開鍵削除フラグ 取得
func (b *DiskBuilder) IsSSHKeysEphemeral() bool {
	return b.isSSHKeysEphemeral
}

// SetSSHKeysEphemeral ディスク作成後の公開鍵削除フラグ 設定
func (b *DiskBuilder) SetSSHKeysEphemeral(isEphemeral bool) {
	b.isSSHKeysEphemeral = isEphemeral
}

// IsNotesEphemeral ディスク作成後のスタートアップスクリプト削除フラグ 取得
func (b *DiskBuilder) IsNotesEphemeral() bool {
	return b.isNotesEphemeral
}

// SetNotesEphemeral ディスク作成後のスタートアップスクリプト削除フラグ 設定
func (b *DiskBuilder) SetNotesEphemeral(isEphemeral bool) {
	b.isNotesEphemeral = isEphemeral
}

// GetGenerateSSHKeyName SSHキー生成 名称 取得
func (b *DiskBuilder) GetGenerateSSHKeyName() string {
	return b.generateSSHKeyName
}

// SetGenerateSSHKeyName SSHキー生成 名称 設定
func (b *DiskBuilder) SetGenerateSSHKeyName(name string) {
	b.generateSSHKeyName = name
}

// GetGenerateSSHKeyPassPhrase SSHキー生成 パスフレーズ 取得
func (b *DiskBuilder) GetGenerateSSHKeyPassPhrase() string {
	return b.generateSSHKeyPassPhrase
}

// SetGenerateSSHKeyPassPhrase SSHキー生成 パスフレーズ 設定
func (b *DiskBuilder) SetGenerateSSHKeyPassPhrase(pass string) {
	b.generateSSHKeyPassPhrase = pass
}

// GetGenerateSSHKeyDescription SSHキー生成 説明 取得
func (b *DiskBuilder) GetGenerateSSHKeyDescription() string {
	return b.generateSSHKeyDescription
}

// SetGenerateSSHKeyDescription SSHキー生成 説明 設定
func (b *DiskBuilder) SetGenerateSSHKeyDescription(desc string) {
	b.generateSSHKeyDescription = desc
}

// GetAutoBoot サーバ自動起動 取得
func (b *DiskBuilder) GetAutoBoot() bool {
	return b.autoBoot
}

// SetAutoBoot サーバ自動起動 設定
func (b *DiskBuilder) SetAutoBoot(autoBoot bool) {
	b.autoBoot = autoBoot
}

// SetEventHandler イベントハンドラ 登録
func (b *DiskBuilder) SetEventHandler(event DiskBuildEvents, handler DiskBuildEventHandler) {
	b.buildEventHandlers[event] = handler
}

// ClearEventHandler イベントハンドラ クリア
func (b *DiskBuilder) ClearEventHandler(event DiskBuildEvents) {
	delete(b.buildEventHandlers, event)
}

// GetEventHandler イベントハンドラ取得
func (b *DiskBuilder) GetEventHandler(event DiskBuildEvents) *DiskBuildEventHandler {
	if handler, ok := b.buildEventHandlers[event]; ok {
		return &handler
	}
	return nil
}

// Build ディスクの構築
func (b *DiskBuilder) Build() (*DiskBuildResult, error) {

	// start
	b.callEventHandlerIfExists(DiskBuildOnStart)

	// create parameters
	b.currentDiskBuildValue = &DiskBuildValue{}
	b.currentDiskBuildResult = &DiskBuildResult{}

	if err := b.buildDiskParams(); err != nil {
		return nil, err
	}

	// create disk
	b.callEventHandlerIfExists(DiskBuildOnCreateDiskBefore)
	if err := b.createDisk(b.currentDiskBuildValue.Disk, b.currentDiskBuildValue.Edit); err != nil {
		return b.currentDiskBuildResult, err
	}
	b.callEventHandlerIfExists(DiskBuildOnCreateDiskAfter)

	// cleanup
	if b.isSSHKeysEphemeral && (len(b.currentDiskBuildResult.SSHKeys) > 0 || b.currentDiskBuildResult.GeneratedSSHKey != nil) {
		b.callEventHandlerIfExists(DiskBuildOnCleanupSSHKeyBefore)

		// created keys
		for _, key := range b.currentDiskBuildResult.SSHKeys {
			_, err := b.client.SSHKey.Delete(key.ID)
			if err != nil {
				return b.currentDiskBuildResult, err
			}
		}

		// generated key
		if b.currentDiskBuildResult.GeneratedSSHKey != nil {
			_, err := b.client.SSHKey.Delete(b.currentDiskBuildResult.GeneratedSSHKey.ID)
			if err != nil {
				return b.currentDiskBuildResult, err
			}
		}

		b.callEventHandlerIfExists(DiskBuildOnCleanupSSHKeyAfter)
	}
	if b.isNotesEphemeral && len(b.currentDiskBuildResult.Notes) > 0 {
		b.callEventHandlerIfExists(DiskBuildOnCleanupNoteBefore)

		for _, note := range b.currentDiskBuildResult.Notes {
			_, err := b.client.Note.Delete(note.ID)
			if err != nil {
				return b.currentDiskBuildResult, err
			}
		}

		b.callEventHandlerIfExists(DiskBuildOnCleanupNoteAfter)
	}

	b.callEventHandlerIfExists(DiskBuildOnComplete)

	return b.currentDiskBuildResult, nil
}

func (b *DiskBuilder) buildDiskParams() error {
	if err := b.buildDiskParam(); err != nil {
		return err
	}

	if !b.isNeedDiskEdit() {
		return nil
	}

	return b.buildDiskEditParam()
}

func (b *DiskBuilder) buildDiskParam() error {
	v := b.currentDiskBuildValue

	v.Disk = b.client.Disk.New()
	d := v.Disk
	d.Name = b.name
	d.SizeMB = b.size * 1024
	d.DistantFrom = b.distantFrom
	d.Plan = b.planID.ToResource()
	d.Connection = b.connection
	d.Description = b.description
	d.Tags = b.GetTags()
	if b.iconID > 0 {
		d.Icon = &sacloud.Icon{Resource: sacloud.NewResource(b.iconID)}
	}
	if b.serverID > 0 {
		d.Server = &sacloud.Server{Resource: sacloud.NewResource(b.serverID)}
	}

	if b.sourceArchiveID > 0 {
		d.SetSourceArchive(b.sourceArchiveID)
		d.SourceDisk = nil
	}
	if b.sourceDiskID > 0 {
		d.SourceArchive = nil
		d.SetSourceDisk(b.sourceDiskID)
	}

	return nil
}

func (b *DiskBuilder) buildDiskEditParam() error {
	v := b.currentDiskBuildValue

	// for DiskEditValue( POST /disk/config )
	v.Edit = b.client.Disk.NewCondig()
	e := v.Edit
	if b.ipAddress != "" {
		e.SetUserIPAddress(b.ipAddress)
	}
	if b.networkMaskLen > 0 {
		e.SetNetworkMaskLen(fmt.Sprintf("%d", b.networkMaskLen))
	}
	if b.defaultRoute != "" {
		e.SetDefaultRoute(b.defaultRoute)
	}
	if b.password != "" {
		e.SetPassword(b.password)
	}
	if b.hostName != "" {
		e.SetHostName(b.hostName)
	}
	e.SetDisablePWAuth(b.disablePWAuth)

	sshKeyIDs := []string{}
	if len(b.sshKeyIDs) > 0 {
		sshKeyIDs = append(sshKeyIDs, b.getStrSSHKeyIDs()...)
	}
	if len(b.sshKeys) > 0 {

		createdIDs, err := b.createSSHKeys()
		if err != nil {
			return err
		}
		sshKeyIDs = append(sshKeyIDs, createdIDs...)
	}
	if b.generateSSHKeyName != "" {
		key, err := b.generateSSHKey(b.generateSSHKeyName, b.generateSSHKeyPassPhrase, b.generateSSHKeyDescription)
		if err != nil {
			return err
		}
		sshKeyIDs = append(sshKeyIDs, fmt.Sprintf("%d", key.ID))
	}
	if len(sshKeyIDs) > 0 {
		e.SetSSHKeys(sshKeyIDs)
	}

	noteIDs := []string{}
	if len(b.noteIDs) > 0 {
		noteIDs = append(noteIDs, b.getStrNoteIDs()...)
	}
	if len(b.notes) > 0 {
		createdIDs, err := b.createNotes()
		if err != nil {
			return err
		}
		noteIDs = append(noteIDs, createdIDs...)

	}
	e.SetNotes(noteIDs)

	return nil
}

func (b *DiskBuilder) createSSHKeys() ([]string, error) {
	createdIDs := []string{}
	for _, strKey := range b.sshKeys {
		key, err := b.createSSHKey(strKey)
		if err != nil {
			return createdIDs, err
		}
		createdIDs = append(createdIDs, key.GetStrID())
	}
	return createdIDs, nil
}

func (b *DiskBuilder) createSSHKey(strKey string) (*sacloud.SSHKey, error) {

	// raise events
	b.callEventHandlerIfExists(DiskBuildOnCreateSSHKeyBefore)

	keyReq := b.client.SSHKey.New()
	keyReq.Name = fmt.Sprintf("publickey-%s", time.Now().Format(time.RFC3339))
	keyReq.PublicKey = strKey

	key, err := b.client.SSHKey.Create(keyReq)
	if err != nil {
		return nil, err
	}
	b.currentDiskBuildResult.addSSHKey(key)

	// raise events
	b.callEventHandlerIfExists(DiskBuildOnCreateSSHKeyAfter)

	return key, nil

}

func (b *DiskBuilder) generateSSHKey(name string, passPhrase string, desc string) (*sacloud.SSHKeyGenerated, error) {

	key, err := b.client.SSHKey.Generate(name, passPhrase, desc)
	if err != nil {
		return nil, err
	}
	b.currentDiskBuildResult.GeneratedSSHKey = key
	return key, nil

}

func (b *DiskBuilder) createNotes() ([]string, error) {
	createdIDs := []string{}
	for _, strNote := range b.notes {
		note, err := b.createNote(strNote)
		if err != nil {
			return createdIDs, err
		}
		createdIDs = append(createdIDs, note.GetStrID())
	}
	return createdIDs, nil

}

func (b *DiskBuilder) createNote(strNote string) (*sacloud.Note, error) {

	// raise events
	b.callEventHandlerIfExists(DiskBuildOnCreateNoteBefore)

	noteReq := b.client.Note.New()
	noteReq.Name = fmt.Sprintf("note-%s", time.Now().Format(time.RFC3339))
	noteReq.Content = strNote

	note, err := b.client.Note.Create(noteReq)
	if err != nil {
		return nil, err
	}
	b.currentDiskBuildResult.addNote(note)

	// raise events
	b.callEventHandlerIfExists(DiskBuildOnCreateNoteAfter)

	return note, nil

}

func (b *DiskBuilder) createDisk(diskReq *sacloud.Disk, editReq *sacloud.DiskEditValue) error {
	var disk *sacloud.Disk
	var err error
	if editReq == nil {
		disk, err = b.client.Disk.Create(diskReq)
	} else {
		disk, err = b.client.Disk.CreateWithConfig(diskReq, editReq, b.autoBoot)
	}

	if err != nil {
		return err
	}

	b.currentDiskBuildResult.Disk = disk
	//wait
	return b.client.Disk.SleepWhileCopying(disk.ID, b.client.DefaultTimeoutDuration)
}

func (b *DiskBuilder) isNeedDiskEdit() bool {
	if b.sourceArchiveID == 0 && b.sourceDiskID == 0 {
		// blank disk
		return false
	}

	return b.forceEditDisk ||
		b.ipAddress != "" ||
		b.networkMaskLen > 0 ||
		b.defaultRoute != "" ||
		b.password != "" ||
		b.hostName != "" ||
		len(b.sshKeyIDs) > 0 ||
		len(b.sshKeys) > 0 ||
		len(b.noteIDs) > 0 ||
		len(b.notes) > 0
}

func (b *DiskBuilder) getStrSSHKeyIDs() []string {
	return b.toStringList(b.sshKeyIDs)
}

func (b *DiskBuilder) getStrNoteIDs() []string {
	return b.toStringList(b.noteIDs)
}

func (b *DiskBuilder) callEventHandlerIfExists(event DiskBuildEvents) {
	if handler, ok := b.buildEventHandlers[event]; ok {
		handler(b.currentDiskBuildValue, b.currentDiskBuildResult)
	}
}

/**********************************************************
  Type : DiskBuildValue
**********************************************************/

// DiskBuildValue ディスク構築用パラメータ
type DiskBuildValue struct {
	// Disk ディスク作成用パラメータ
	Disk *sacloud.Disk
	// Edit ディスクの編集用パラメータ
	Edit *sacloud.DiskEditValue
}

/**********************************************************
  Type : DiskBuildResult
**********************************************************/

// DiskBuildResult ディスク構築結果
type DiskBuildResult struct {
	// Disk ディスク
	Disk *sacloud.Disk
	// Notes スタートアップスクリプト
	Notes []*sacloud.Note
	// SSHKeys 公開鍵
	SSHKeys []*sacloud.SSHKey
	// GeneratedSSHKey 生成されたSSHキー
	GeneratedSSHKey *sacloud.SSHKeyGenerated
}

func (d *DiskBuildResult) addNote(note *sacloud.Note) {
	d.Notes = append(d.Notes, note)
}

func (d *DiskBuildResult) addSSHKey(key *sacloud.SSHKey) {
	d.SSHKeys = append(d.SSHKeys, key)
}
