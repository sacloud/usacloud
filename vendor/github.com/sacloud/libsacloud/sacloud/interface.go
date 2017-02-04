package sacloud

// Interface インターフェース(NIC)
type Interface struct {
	*Resource                   // ID
	propServer                  // サーバー
	propSwitch                  // スイッチ
	MACAddress    string        `json:",omitempty"` // MACアドレス
	IPAddress     string        `json:",omitempty"` // IPアドレス
	UserIPAddress string        `json:",omitempty"` // ユーザー指定IPアドレス
	HostName      string        `json:",omitempty"` // ホスト名
	PacketFilter  *PacketFilter `json:",omitempty"` // 適用パケットフィルタ
}
