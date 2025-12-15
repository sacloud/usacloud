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

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"slices"
	"strings"
	"time"

	"github.com/sacloud/api-client-go/profile"
	sacloudhttp "github.com/sacloud/go-http"
	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/packages-go/envvar"
	saclient "github.com/sacloud/saclient-go"
	"github.com/sacloud/usacloud/pkg/query"
	"github.com/sacloud/usacloud/pkg/validate"
	"github.com/spf13/pflag"
)

// Config CLI全コマンドが利用するフラグ
type Config struct {
	profile.ConfigValue

	// Profile プロファイル名
	Profile string `json:"-"`

	// DefaultOutputType デフォルトアウトプットタイプ
	DefaultOutputType string

	// NoColor ANSIエスケープシーケンスによる色つけを無効化
	NoColor bool

	// ProcessTimeoutSec コマンド全体の実行タイムアウトまでの秒数
	ProcessTimeoutSec int

	// ArgumentMatchMode 引数でリソースを特定する際にリソース名と引数を比較する方法を指定
	// 有効な値:
	//   - `partial`(デフォルト): 部分一致
	//   - `exact`: 完全一致
	// Note: 引数はID or Name or Tagsと比較されるが、この項目はNameとの比較にのみ影響する。IDとTagsは常に完全一致となる。
	ArgumentMatchMode string

	// DefaultQueryDriver 各コマンドで--query-driverが省略された場合のデフォルト値
	// 有効な値:
	//   - `jmespath`(デフォルト): JMESPath
	//   - `jq` : gojq
	DefaultQueryDriver string

	// Fields related to service principal keys:

	// TokenEndpoint specifes IAM token request endpoint (URL)
	TokenEndpoint string `json:",omitempty"`

	// ServicePrincipalID is a resource ID of the service principal itself
	ServicePrincipalID string `json:",omitempty"`

	// ServicePrincipalKeyID is a Key's ID; note that a servie principal can have multiple keys.
	ServicePrincipalKeyID string `json:",omitempty"`

	// PrivateKeyPEMPath is a path to a local file which stores private key of ServicePrincipalKeyID.
	PrivateKeyPEMPath string `json:",omitempty"`
}

var TheClient saclient.Client
var DefaultProcessTimeoutSec = 60 * 60 * 2 // 2時間
var DefaultQueryDriver = query.DriverJMESPath

// LoadConfigValue 指定のフラグセットからフラグを読み取り*Flagsを組み立てて返す
func LoadConfigValue(flags *pflag.FlagSet, errW io.Writer, skipLoadingProfile bool) (*Config, error) {
	o := &Config{
		ConfigValue: profile.ConfigValue{},
	}
	if skipLoadingProfile {
		return o, nil
	}

	o.loadConfig(flags, errW)
	return o, o.Validate(false)
}

func (o *Config) IsEmpty() bool {
	return o.AccessToken == "" &&
		o.AccessTokenSecret == "" &&
		o.Zone == "" && o.DefaultOutputType == ""
}

func (o *Config) loadConfig(flags *pflag.FlagSet, errW io.Writer) {
	// プロファイルだけ先に環境変数を読んでおく
	if o.Profile == "" {
		o.Profile = envvar.StringFromEnvMulti([]string{"SAKURA_PROFILE", "SAKURACLOUD_PROFILE", "USACLOUD_PROFILE"}, "")
	}

	o.loadFromProfile(flags, errW)
	o.loadFromEnv()
	o.loadFromFlags(flags, errW)
	o.fillDefaults()
}

func (o *Config) fillDefaults() {
	if len(o.Zones) == 0 {
		o.Zones = iaas.SakuraCloudZones
	}
	if !slices.Contains(o.Zones, "all") {
		o.Zones = append(o.Zones, "all")
	}
}

func (o *Config) loadFromEnv() {
	if o.AccessToken == "" {
		o.AccessToken = envvar.StringFromEnvMulti([]string{"SAKURA_ACCESS_TOKEN", "SAKURACLOUD_ACCESS_TOKEN"}, "")
	}
	if o.AccessTokenSecret == "" {
		o.AccessTokenSecret = envvar.StringFromEnvMulti([]string{"SAKURA_ACCESS_TOKEN_SECRET", "SAKURACLOUD_ACCESS_TOKEN_SECRET"}, "")
	}
	if o.Zone == "" {
		o.Zone = envvar.StringFromEnvMulti([]string{"SAKURA_ZONE", "SAKURACLOUD_ZONE"}, "")
	}
	if len(o.Zones) == 0 {
		o.Zones = envvar.StringSliceFromEnvMulti([]string{"SAKURA_ZONES", "SAKURACLOUD_ZONES"}, []string{})
	}
	if o.AcceptLanguage == "" {
		o.AcceptLanguage = envvar.StringFromEnvMulti([]string{"SAKURA_ACCEPT_LANGUAGE", "SAKURACLOUD_ACCEPT_LANGUAGE"}, "")
	}
	if o.RetryMax <= 0 {
		o.RetryMax = envvar.IntFromEnvMulti([]string{"SAKURA_RETRY_MAX", "SAKURACLOUD_RETRY_MAX"}, sacloudhttp.DefaultRetryMax)
	}
	if o.RetryWaitMax <= 0 {
		o.RetryWaitMax = envvar.IntFromEnvMulti([]string{"SAKURA_RETRY_WAIT_MAX", "SAKURACLOUD_RETRY_WAIT_MAX"}, int(sacloudhttp.DefaultRetryWaitMax.Seconds()))
	}
	if o.RetryWaitMin <= 0 {
		o.RetryWaitMin = envvar.IntFromEnvMulti([]string{"SAKURA_RETRY_WAIT_MIN", "SAKURACLOUD_RETRY_WAIT_MIN"}, int(sacloudhttp.DefaultRetryWaitMin.Seconds()))
	}
	if o.HTTPRequestTimeout <= 0 {
		o.HTTPRequestTimeout = envvar.IntFromEnvMulti([]string{"SAKURA_API_REQUEST_TIMEOUT", "SAKURACLOUD_API_REQUEST_TIMEOUT"}, 300)
	}
	if o.HTTPRequestRateLimit <= 0 {
		o.HTTPRequestRateLimit = envvar.IntFromEnvMulti([]string{"SAKURA_API_REQUEST_RATE_LIMIT", "SAKURACLOUD_API_REQUEST_RATE_LIMIT"}, 5) // デフォルト5ゾーン分(is1a/is1b/tk1a/tk1b/tk1v)
	}
	if o.APIRootURL == "" {
		o.APIRootURL = envvar.StringFromEnvMulti([]string{"SAKURA_API_ROOT_URL", "SAKURACLOUD_API_ROOT_URL"}, iaas.SakuraCloudAPIRoot)
	}
	if o.DefaultZone == "" {
		o.DefaultZone = envvar.StringFromEnvMulti([]string{"SAKURA_DEFAULT_ZONE", "SAKURACLOUD_DEFAULT_ZONE"}, iaas.APIDefaultZone)
	}
	if o.TraceMode == "" {
		o.TraceMode = envvar.StringFromEnvMulti([]string{"SAKURA_TRACE", "SAKURACLOUD_TRACE"}, "")
	}
	if !o.FakeMode {
		o.FakeMode = envvar.StringFromEnvMulti([]string{"SAKURA_FAKE_MODE", "SAKURACLOUD_FAKE_MODE"}, "") != ""
	}
	if o.FakeStorePath == "" {
		o.FakeStorePath = envvar.StringFromEnvMulti([]string{"SAKURA_FAKE_STORE_PATH", "SAKURACLOUD_FAKE_STORE_PATH"}, "")
	}
	if o.ProcessTimeoutSec <= 0 {
		o.ProcessTimeoutSec = envvar.IntFromEnvMulti([]string{"SAKURA_PROCESS_TIMEOUT_SEC", "SAKURACLOUD_PROCESS_TIMEOUT_SEC"}, DefaultProcessTimeoutSec)
	}
	if o.ArgumentMatchMode == "" {
		o.ArgumentMatchMode = envvar.StringFromEnvMulti([]string{"SAKURA_ARGUMENT_MATCH_MODE", "SAKURACLOUD_ARGUMENT_MATCH_MODE"}, "partial")
	}
	if o.DefaultOutputType == "" {
		o.DefaultOutputType = envvar.StringFromEnvMulti([]string{"SAKURA_DEFAULT_OUTPUT_TYPE", "SAKURACLOUD_DEFAULT_OUTPUT_TYPE"}, "")
	}
	if o.DefaultQueryDriver == "" {
		o.DefaultQueryDriver = envvar.StringFromEnvMulti([]string{"SAKURA_DEFAULT_QUERY_DRIVER", "SAKURACLOUD_DEFAULT_QUERY_DRIVER"}, "")
	}
}

func (o *Config) loadFromFlags(flags *pflag.FlagSet, errW io.Writer) {
	var argv []string

	if flags.Changed("token") {
		v, err := flags.GetString("token")
		if err != nil {
			fmt.Fprintf(errW, "[WARN] reading value of %q flag is failed: %s", "token", err)
			return
		}
		o.AccessToken = v
		argv = append(argv, "--token", v)
	}
	if flags.Changed("secret") {
		v, err := flags.GetString("secret")
		if err != nil {
			fmt.Fprintf(errW, "[WARN] reading value of %q flag is failed: %s", "secret", err)
			return
		}
		o.AccessTokenSecret = v
		argv = append(argv, "--secret", v)
	}
	if flags.Changed("zones") {
		v, err := flags.GetStringSlice("zones")
		if err != nil {
			fmt.Fprintf(errW, "[WARN] reading value of %q flag is failed: %s", "zones", err)
			return
		}
		o.Zones = v
		var buf strings.Builder
		if err = csv.NewWriter(&buf).Write(v); err != nil {
			fmt.Fprintf(errW, "[WARN] reading value of %q flag is failed: %s", "zones", err)
			return
		}
		argv = append(argv, "--zones", buf.String())
	}
	if flags.Changed("no-color") {
		v, err := flags.GetBool("no-color")
		if err != nil {
			fmt.Fprintf(errW, "[WARN] reading value of %q flag is failed: %s", "no-color", err)
			return
		}
		o.NoColor = v
	}
	if flags.Changed("trace") {
		v, err := flags.GetBool("trace")
		if err != nil {
			fmt.Fprintf(errW, "[WARN] reading value of %q flag is failed: %s", "trace", err)
			return
		}
		if v {
			o.TraceMode = "all"
			argv = append(argv, "--trace")
		}
	}
	if flags.Changed("fake") {
		v, err := flags.GetBool("fake")
		if err != nil {
			fmt.Fprintf(errW, "[WARN] reading value of %q flag is failed: %s", "fake", err)
			return
		}
		o.FakeMode = v
	}
	if flags.Changed("fake-store") {
		v, err := flags.GetString("fake-store")
		if err != nil {
			fmt.Fprintf(errW, "[WARN] reading value of %q flag is failed: %s", "fake-store", err)
			return
		}
		o.FakeStorePath = v
	}
	if flags.Changed("process-timeout-sec") {
		v, err := flags.GetInt("process-timeout-sec")
		if err != nil {
			fmt.Fprintf(errW, "[WARN] reading value of %q flag is failed: %s", "process-timeout-sec", err)
			return
		}
		o.ProcessTimeoutSec = v
	}
	if flags.Changed("argument-match-mode") {
		v, err := flags.GetString("argument-match-mode")
		if err != nil {
			fmt.Fprintf(errW, "[WARN] reading value of %q flag is failed: %s", "argument-match-mode", err)
			return
		}
		o.ArgumentMatchMode = v
	}

	if flags.Changed("profile") {
		v, err := flags.GetString("profile")
		if err != nil {
			fmt.Fprintf(errW, "[WARN] reading value of %q flag is failed: %s", "profile", err)
			return
		}
		argv = append(argv, "--profile", v)
	}
	if flags.Changed("private-key-path") {
		v, err := flags.GetString("private-key-path")
		if err != nil {
			fmt.Fprintf(errW, "[WARN] reading value of %q flag is failed: %s", "profile", err)
			return
		}
		argv = append(argv, "--private-key-path", v)
	}
	if flags.Changed("service-principal-id") {
		v, err := flags.GetString("service-principal-id")
		if err != nil {
			fmt.Fprintf(errW, "[WARN] reading value of %q flag is failed: %s", "profile", err)
			return
		}
		argv = append(argv, "--service-principal-id", v)
	}
	if flags.Changed("service-principal-key-id") {
		v, err := flags.GetString("service-principal-key-id")
		if err != nil {
			fmt.Fprintf(errW, "[WARN] reading value of %q flag is failed: %s", "profile", err)
			return
		}
		argv = append(argv, "--service-principal-key-id", v)
	}

	if err := TheClient.FlagSet(flag.ContinueOnError).Parse(argv); err != nil {
		fmt.Fprintf(errW, "[WARN] argv reconstrcution failed: %s", err)
		return
	}
}

func (o *Config) LoadFromAttributes(p *saclient.Profile) error {
	if buf, err := json.Marshal(p.Attributes); err != nil {
		return err
	} else if err := json.Unmarshal(buf, o); err != nil {
		return err
	} else {
		o.Profile = p.Name
		o.fillDefaults()
		return nil
	}
}

func (o *Config) IntoAttributes() (*saclient.Profile, error) {
	ret := saclient.Profile{
		Name:       o.Profile,
		Attributes: make(map[string]any),
	}

	if buf, err := json.Marshal(o); err != nil {
		return nil, err
	} else if err := json.Unmarshal(buf, &ret.Attributes); err != nil {
		return nil, err
	} else {
		return &ret, nil
	}
}

func (o *Config) Validate(skipCred bool) error {
	var errs []error

	switch o.DefaultOutputType {
	case "", "table", "json", "yaml":
		// noop
	default:
		errs = append(errs, validate.NewFlagError("profile.DefaultOutputType", "must be one of [table/json/yaml]"))
	}

	switch o.ArgumentMatchMode {
	case "", "partial", "exact":
		// noop
	default:
		errs = append(errs, validate.NewFlagError("--argument-match-mode", "must be one of [partial/exact]"))
	}

	switch o.DefaultQueryDriver {
	case "", query.DriverJMESPath, query.DriverGoJQ:
		// noop
	default:
		errs = append(errs, validate.NewFlagError("profile.DefaultQueryDriver", "must be one of [jmespath/jq]"))
	}

	return validate.NewValidationError(errs...)
}

func (o *Config) ProcessTimeout() time.Duration {
	sec := o.ProcessTimeoutSec
	if sec <= 0 {
		sec = DefaultProcessTimeoutSec
	}
	return time.Duration(sec) * time.Second
}

func (o *Config) ArgumentMatchModeValue() string {
	if o.ArgumentMatchMode == "" {
		return "partial"
	}
	return o.ArgumentMatchMode
}
