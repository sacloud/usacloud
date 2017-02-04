package sacloud

// License ライセンス
type License struct {
	*Resource       // ID
	propName        // 名称
	propDescription // 説明
	propCreatedAt   // 作成日時
	PropModifiedAt  // 変更日時

	LicenseInfo *ProductLicense `json:",omitempty"` // ライセンス情報
}
