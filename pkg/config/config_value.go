// Copyright 2017-2025 The sacloud/usacloud Authors
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

package config

import "strings"

const (
	// DefaultProfileName デフォルトのプロファイル名
	DefaultProfileName = "default"

	// EnableAPITraceWord TraceModeに設定する、APIトレースを有効化するためのキーワード
	EnableAPITraceWord = "api"

	// EnableHTTPTraceWord TraceModeに設定する、HTTPトレースを有効化するためのキーワード
	EnableHTTPTraceWord = "http"
)

// ConfigValue プロファイル コンフィグ
//
// 以前は github.com/sacloud/api-client-go/profile.ConfigValue を利用していたが、
// sacloud-sdk-goでは当該パッケージがinternalとなったためusacloud側で定義している。
type ConfigValue struct {
	// AccessToken アクセストークン
	AccessToken string
	// AccessTokenSecret アクセスシークレット
	AccessTokenSecret string

	// Zone デフォルトゾーン
	Zone string
	// Zones 利用可能なゾーン
	Zones []string

	// UserAgent ユーザーエージェント
	UserAgent string `json:",omitempty"`
	// AcceptLanguage リクエスト時のAccept-Languageヘッダ
	AcceptLanguage string
	// Gzip Gzip圧縮の有効化
	Gzip bool

	// RetryMax 423/503時のリトライ回数
	RetryMax int
	// RetryWaitMin 423/503時のリトライ間隔(最小) 単位:秒
	RetryWaitMin int
	// RetryWaitMax 423/503時のリトライ間隔(最大) 単位:秒
	RetryWaitMax int

	// StatePollingTimeout StatePollWaiterでのタイムアウト 単位:秒
	StatePollingTimeout int
	// StatePollingInterval StatePollWaiterでのポーリング間隔 単位:秒
	StatePollingInterval int

	// HTTPRequestTimeout APIリクエスト時のHTTPタイムアウト 単位:秒
	HTTPRequestTimeout int
	// HTTPRequestRateLimit APIリクエスト時の1秒あたりのリクエスト上限数
	HTTPRequestRateLimit int

	// APIRootURL APIのルートURL
	APIRootURL string

	// DefaultZone グローバルリソースAPIを呼ぶ際に指定するゾーン
	DefaultZone string

	// TraceMode トレースモード
	TraceMode string
	// FakeMode フェイクモード有効化
	FakeMode bool
	// FakeStorePath フェイクモードでのファイルストアパス
	FakeStorePath string
}

func (o *ConfigValue) EnableHTTPTrace() bool {
	return EnableHTTPTrace(o.TraceMode)
}

func (o *ConfigValue) EnableAPITrace() bool {
	return EnableAPITrace(o.TraceMode)
}

func traceModeValue(strTraceMode string) string {
	return strings.ToLower(strings.TrimSpace(strTraceMode))
}

// EnableHTTPTrace 指定のトレースモード文字列がHTTPトレースを有効にするか判定する
func EnableHTTPTrace(strTraceMode string) bool {
	traceMode := traceModeValue(strTraceMode)
	if traceMode == "" {
		return false
	}

	// TraceModeが"api"の場合はfalseにする(TraceMode=1などの場合はAPI/HTTP両方が有効になる)
	if traceMode == EnableAPITraceWord {
		return false
	}
	return true
}

// EnableAPITrace 指定のトレースモード文字列がAPIトレースを有効にするか判定する
func EnableAPITrace(strTraceMode string) bool {
	traceMode := traceModeValue(strTraceMode)
	if traceMode == "" {
		return false
	}

	// TraceModeが"http"の場合はfalseにする(TraceMode=1などの場合はAPI/HTTP両方が有効になる)
	if traceMode == EnableHTTPTraceWord || traceMode == "error" {
		return false
	}
	return true
}
