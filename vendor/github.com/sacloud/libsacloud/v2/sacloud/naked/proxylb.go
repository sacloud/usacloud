// Copyright 2016-2020 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package naked

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// ProxyLB エンハンスドロードバランサ
type ProxyLB struct {
	ID           types.ID            `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string              `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description  string              `yaml:"description"`
	Tags         types.Tags          `yaml:"tags"`
	Icon         *Icon               `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt    *time.Time          `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt   *time.Time          `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
	Availability types.EAvailability `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	Provider     *Provider           `json:",omitempty" yaml:"provider,omitempty" structs:",omitempty"`
	Settings     *ProxyLBSettings    `json:",omitempty" yaml:"settings,omitempty" structs:",omitempty"`
	SettingsHash string              `json:",omitempty" yaml:"settings_hash,omitempty" structs:",omitempty"`
	Status       *ProxyLBStatus      `json:",omitempty" yaml:"status,omitempty" structs:",omitempty"`

	// ServiceClass [HACK] プランはServiceClassとして文字列で表す必要がある。 詳細はtypes.EProxyLBPlanのコメント参照
	ServiceClass types.EProxyLBPlan `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
}

// ProxyLBSettingsUpdate エンハンスドロードバランサ
type ProxyLBSettingsUpdate struct {
	Settings     *ProxyLBSettings `json:",omitempty" yaml:"settings,omitempty" structs:",omitempty"`
	SettingsHash string           `json:",omitempty" yaml:"settings_hash,omitempty" structs:",omitempty"`
}

// ProxyLBSettings エンハンスドロードバランサ設定
type ProxyLBSettings struct {
	ProxyLB *ProxyLBSetting `json:",omitempty" yaml:"proxy_lb,omitempty" structs:",omitempty"`
}

// ProxyLBSetting エンハンスドロードバランサ設定
type ProxyLBSetting struct {
	HealthCheck   ProxyLBHealthCheck   `yaml:"health_check"`                                                  // ヘルスチェック
	SorryServer   ProxyLBSorryServer   `yaml:"sorry_server"`                                                  // ソーリーサーバー
	BindPorts     []*ProxyLBBindPorts  `yaml:"bind_ports"`                                                    // プロキシ方式(プロトコル&ポート)
	Servers       []ProxyLBServer      `yaml:"servers"`                                                       // サーバー
	Rules         []ProxyLBRule        `yaml:"rules"`                                                         // 振り分けルール
	LetsEncrypt   *ProxyLBACMESetting  `json:",omitempty" yaml:"lets_encrypt,omitempty" structs:",omitempty"` // Let's encryptでの証明書取得設定
	StickySession ProxyLBStickySession `yaml:"sticky_session"`                                                // StickySession
	Timeout       ProxyLBTimeout       `json:",omitempty" yaml:"timeout,omitempty" structs:",omitempty"`      // タイムアウト
}

// MarshalJSON nullの場合に空配列を出力するための実装
func (s ProxyLBSetting) MarshalJSON() ([]byte, error) {
	if s.BindPorts == nil {
		s.BindPorts = make([]*ProxyLBBindPorts, 0)
	}
	if s.Servers == nil {
		s.Servers = make([]ProxyLBServer, 0)
	}
	if s.Rules == nil {
		s.Rules = make([]ProxyLBRule, 0)
	}

	type alias ProxyLBSetting
	tmp := alias(s)
	return json.Marshal(&tmp)
}

// ProxyLBHealthCheck ヘルスチェック
type ProxyLBHealthCheck struct {
	Protocol  types.EProxyLBHealthCheckProtocol `json:",omitempty" yaml:"protocol,omitempty" structs:",omitempty"`
	Path      string                            `json:",omitempty" yaml:"path,omitempty" structs:",omitempty"`
	Host      string                            `json:",omitempty" yaml:"host,omitempty" structs:",omitempty"`
	DelayLoop int                               `json:",omitempty" yaml:"delay_loop,omitempty" structs:",omitempty"`
}

// ProxyLBSorryServer ソーリーサーバ設定
type ProxyLBSorryServer struct {
	IPAddress string `yaml:"ip_address"`
	Port      *int   `yaml:"port"`
}

// ProxyLBBindPorts プロキシ方式
type ProxyLBBindPorts struct {
	ProxyMode         types.EProxyLBProxyMode  `json:",omitempty" yaml:"proxy_mode,omitempty" structs:",omitempty"`          // モード(プロトコル)
	Port              int                      `json:",omitempty" yaml:"port,omitempty" structs:",omitempty"`                // ポート
	RedirectToHTTPS   bool                     `json:"RedirectToHttps" yaml:"redirect_to_https"`                             // HTTPSへのリダイレクト(モードがhttpの場合のみ)
	SupportHTTP2      bool                     `json:"SupportHttp2" yaml:"support_http2"`                                    // HTTP/2のサポート(モードがhttpsの場合のみ)
	AddResponseHeader []*ProxyLBResponseHeader `json:",omitempty" yaml:"add_response_header,omitempty" structs:",omitempty"` // レスポンスヘッダ
}

// ProxyLBResponseHeader ポートごとの追加レスポンスヘッダ
type ProxyLBResponseHeader struct {
	Header string // ヘッダ名称(英字, 数字, ハイフン)
	Value  string // 値(英字, 数字, 半角スペース, 一部記号（!#$%&'()*+,-./:;<=>?@[]^_`{|}~）)
}

// ProxyLBServer ProxyLB配下のサーバー
type ProxyLBServer struct {
	IPAddress   string `json:",omitempty" yaml:"ip_address,omitempty" structs:",omitempty"` // IPアドレス
	Port        int    `json:",omitempty" yaml:"port,omitempty" structs:",omitempty"`       // ポート
	ServerGroup string `yaml:"server_group"`                                                // サーバグループ
	Enabled     bool   // 有効/無効
}

// ProxyLBRule ProxyLBの振り分けルール
type ProxyLBRule struct {
	Host        string `json:",omitempty" yaml:"host,omitempty" structs:",omitempty"` // ホストヘッダのパターン(ワイルドカードとして?と*が利用可能)
	Path        string `json:",omitempty" yaml:"path,omitempty" structs:",omitempty"` // パス
	ServerGroup string `yaml:"server_group"`
}

// ProxyLBACMESetting Let's Encryptでの証明書取得設定
type ProxyLBACMESetting struct {
	Enabled    bool
	CommonName string `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
}

// ProxyLBStickySession セッション維持(Sticky session)設定
type ProxyLBStickySession struct {
	Enabled bool
	Method  string `json:",omitempty" yaml:"method,omitempty" structs:",omitempty"`
}

// ProxyLBTimeout 実サーバの通信タイムアウト
type ProxyLBTimeout struct {
	InactiveSec int `json:",omitempty" yaml:"inactive_sec" structs:",omitempty"` // 10から600まで1秒刻みで設定可
}

// ProxyLBStatus ステータス
type ProxyLBStatus struct {
	UseVIPFailover bool                 `yaml:"use_vip_failover"`
	Region         types.EProxyLBRegion `json:",omitempty" yaml:"region,omitempty" structs:",omitempty"`
	ProxyNetworks  []string             `json:",omitempty" yaml:"proxy_networks,omitempty" structs:",omitempty"`
	FQDN           string               `json:",omitempty" yaml:"fqdn,omitempty" structs:",omitempty"`
}

// ProxyLBAdditionalCerts additional certificates
type ProxyLBAdditionalCerts []*ProxyLBCertificate

// ProxyLBCertificates ProxyLBのSSL証明書
type ProxyLBCertificates struct {
	PrimaryCert     *ProxyLBCertificate    `yaml:"primary_cert"`
	AdditionalCerts ProxyLBAdditionalCerts `yaml:"additional_certs"`
}

// UnmarshalJSON UnmarshalJSON(AdditionalCertsが空の場合に空文字を返す問題への対応)
func (p *ProxyLBAdditionalCerts) UnmarshalJSON(data []byte) error {
	targetData := strings.Replace(strings.Replace(string(data), " ", "", -1), "\n", "", -1)
	if targetData == `` {
		return nil
	}

	var certs []*ProxyLBCertificate
	if err := json.Unmarshal(data, &certs); err != nil {
		return err
	}

	*p = certs
	return nil
}

// ProxyLBCertificate ProxyLBのSSL証明書詳細
type ProxyLBCertificate struct {
	ServerCertificate       string     `yaml:"server_certificate"`                                                       // サーバ証明書
	IntermediateCertificate string     `yaml:"intermediate_certificate"`                                                 // 中間証明書
	PrivateKey              string     `yaml:"private_key"`                                                              // 秘密鍵
	CertificateEndDate      *time.Time `json:",omitempty" yaml:"certificate_end_date,omitempty" structs:",omitempty"`    // 有効期限
	CertificateCommonName   string     `json:",omitempty" yaml:"certificate_common_name,omitempty" structs:",omitempty"` // CommonName
}

// UnmarshalJSON UnmarshalJSON(CertificateEndDateのtime.TimeへのUnmarshal対応)
func (p *ProxyLBCertificate) UnmarshalJSON(data []byte) error {
	var tmp map[string]interface{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	p.ServerCertificate = tmp["ServerCertificate"].(string)
	p.IntermediateCertificate = tmp["IntermediateCertificate"].(string)
	p.PrivateKey = tmp["PrivateKey"].(string)
	p.CertificateCommonName = tmp["CertificateCommonName"].(string)
	endDate := tmp["CertificateEndDate"].(string)
	if endDate != "" {
		date, err := time.Parse("Jan _2 15:04:05 2006 MST", endDate)
		if err != nil {
			return err
		}
		p.CertificateEndDate = &date
	}

	return nil
}

// ProxyLBHealth ProxyLBのヘルスチェック戻り値
type ProxyLBHealth struct {
	ActiveConn int                    `json:",omitempty" yaml:"active_conn,omitempty" structs:",omitempty"` // アクティブなコネクション数
	CPS        int                    `json:",omitempty" yaml:"cps,omitempty" structs:",omitempty"`         // 秒あたりコネクション数
	Servers    []*ProxyLBHealthServer `json:",omitempty" yaml:"servers,omitempty" structs:",omitempty"`     // 実サーバのステータス
	CurrentVIP string                 `json:",omitempty" yaml:"current_vip,omitempty" structs:",omitempty"` // 現在のVIP
}

// ProxyLBHealthServer ProxyLBの実サーバのステータス
type ProxyLBHealthServer struct {
	ActiveConn int                         `json:",omitempty" yaml:"active_conn,omitempty" structs:",omitempty"` // アクティブなコネクション数
	Status     types.EServerInstanceStatus `json:",omitempty" yaml:"status,omitempty" structs:",omitempty"`      // ステータス(UP or DOWN)
	IPAddress  string                      `json:",omitempty" yaml:"ip_address,omitempty" structs:",omitempty"`  // IPアドレス
	Port       string                      `json:",omitempty" yaml:"port,omitempty" structs:",omitempty"`        // ポート
	CPS        int                         `json:",omitempty" yaml:"cps,omitempty" structs:",omitempty"`         // 秒あたりコネクション数
}
