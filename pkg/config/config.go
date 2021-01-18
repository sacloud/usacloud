// Copyright 2017-2021 The Usacloud Authors
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
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/sacloud/usacloud/pkg/query"

	"github.com/sacloud/usacloud/pkg/validate"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/profile"
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
}

var DefaultProcessTimeoutSec = 60 * 60 * 2 // 2時間
var DefaultQueryDriver = query.DriverJMESPath

// LoadConfigValue 指定のフラグセットからフラグを読み取り*Flagsを組み立てて返す
func LoadConfigValue(flags *pflag.FlagSet, errW io.Writer, skipLoadingProfile bool) (*Config, error) {
	o := &Config{
		ConfigValue: profile.ConfigValue{
			Zones: append(sacloud.SakuraCloudZones, "all"),
		},
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
		o.Profile = stringFromEnvMulti([]string{"SAKURACLOUD_PROFILE", "USACLOUD_PROFILE"}, "default")
	}

	o.loadFromProfile(flags, errW)
	o.loadFromEnv()
	o.loadFromFlags(flags, errW)
	o.fillDefaults()
}

func (o *Config) fillDefaults() {
	if len(o.Zones) == 0 {
		o.Zones = sacloud.SakuraCloudZones
	}
}

func (o *Config) loadFromEnv() {
	if o.AccessToken == "" {
		o.AccessToken = stringFromEnv("SAKURACLOUD_ACCESS_TOKEN", "")
	}
	if o.AccessTokenSecret == "" {
		o.AccessTokenSecret = stringFromEnv("SAKURACLOUD_ACCESS_TOKEN_SECRET", "")
	}
	if o.Zone == "" {
		o.Zone = stringFromEnv("SAKURACLOUD_ZONE", "")
	}
	if len(o.Zones) == 0 {
		o.Zones = stringSliceFromEnv("SAKURACLOUD_ZONES", append(sacloud.SakuraCloudZones, "all"))
	}
	if o.AcceptLanguage == "" {
		o.AcceptLanguage = stringFromEnv("SAKURACLOUD_ACCEPT_LANGUAGE", "")
	}
	if o.RetryMax <= 0 {
		o.RetryMax = intFromEnv("SAKURACLOUD_RETRY_MAX", sacloud.APIDefaultRetryMax)
	}
	if o.RetryWaitMax <= 0 {
		o.RetryWaitMax = intFromEnv("SAKURACLOUD_RETRY_WAIT_MAX", 64)
	}
	if o.RetryWaitMin <= 0 {
		o.RetryWaitMin = intFromEnv("SAKURACLOUD_RETRY_WAIT_MIN", 1)
	}
	if o.HTTPRequestTimeout <= 0 {
		o.HTTPRequestTimeout = intFromEnv("SAKURACLOUD_API_REQUEST_TIMEOUT", 300)
	}
	if o.HTTPRequestRateLimit <= 0 {
		o.HTTPRequestRateLimit = intFromEnv("SAKURACLOUD_API_REQUEST_RATE_LIMIT", 5) // デフォルト5ゾーン分(is1a/is1b/tk1a/tk1b/tk1v)
	}
	if o.APIRootURL == "" {
		o.APIRootURL = stringFromEnv("SAKURACLOUD_API_ROOT_URL", sacloud.SakuraCloudAPIRoot)
	}
	if o.DefaultZone == "" {
		o.DefaultZone = stringFromEnv("SAKURACLOUD_DEFAULT_ZONE", sacloud.APIDefaultZone)
	}
	if o.TraceMode == "" {
		o.TraceMode = stringFromEnv("SAKURACLOUD_TRACE", "")
	}
	if !o.FakeMode {
		o.FakeMode = os.Getenv("SAKURACLOUD_FAKE_MODE") != ""
	}
	if o.FakeStorePath == "" {
		o.FakeStorePath = stringFromEnv("SAKURACLOUD_FAKE_STORE_PATH", "")
	}
	if o.ProcessTimeoutSec <= 0 {
		o.ProcessTimeoutSec = intFromEnv("SAKURACLOUD_PROCESS_TIMEOUT_SEC", DefaultProcessTimeoutSec)
	}
	if o.ArgumentMatchMode == "" {
		o.ArgumentMatchMode = stringFromEnv("SAKURACLOUD_ARGUMENT_MATCH_MODE", "partial")
	}
}

func (o *Config) loadFromFlags(flags *pflag.FlagSet, errW io.Writer) {
	if flags.Changed("token") {
		v, err := flags.GetString("token")
		if err != nil {
			fmt.Fprintf(errW, "[WARN] reading value of %q flag is failed: %s", "token", err) // nolint
			return
		}
		o.AccessToken = v
	}
	if flags.Changed("secret") {
		v, err := flags.GetString("secret")
		if err != nil {
			fmt.Fprintf(errW, "[WARN] reading value of %q flag is failed: %s", "secret", err) // nolint
			return
		}
		o.AccessTokenSecret = v
	}
	if flags.Changed("zones") {
		v, err := flags.GetStringSlice("zones")
		if err != nil {
			fmt.Fprintf(errW, "[WARN] reading value of %q flag is failed: %s", "zones", err) // nolint
			return
		}
		o.Zones = v
	}
	if flags.Changed("no-color") {
		v, err := flags.GetBool("no-color")
		if err != nil {
			fmt.Fprintf(errW, "[WARN] reading value of %q flag is failed: %s", "no-color", err) // nolint
			return
		}
		o.NoColor = v
	}
	if flags.Changed("trace") {
		v, err := flags.GetBool("trace")
		if err != nil {
			fmt.Fprintf(errW, "[WARN] reading value of %q flag is failed: %s", "trace", err) // nolint
			return
		}
		if v {
			o.TraceMode = "all"
		}
	}
	if flags.Changed("fake") {
		v, err := flags.GetBool("fake")
		if err != nil {
			fmt.Fprintf(errW, "[WARN] reading value of %q flag is failed: %s", "fake", err) // nolint
			return
		}
		o.FakeMode = v
	}
	if flags.Changed("fake-store") {
		v, err := flags.GetString("fake-store")
		if err != nil {
			fmt.Fprintf(errW, "[WARN] reading value of %q flag is failed: %s", "fake-store", err) // nolint
			return
		}
		o.FakeStorePath = v
	}
	if flags.Changed("process-timeout-sec") {
		v, err := flags.GetInt("process-timeout-sec")
		if err != nil {
			fmt.Fprintf(errW, "[WARN] reading value of %q flag is failed: %s", "process-timeout-sec", err) // nolint
			return
		}
		o.ProcessTimeoutSec = v
	}
	if flags.Changed("argument-match-mode") {
		v, err := flags.GetString("argument-match-mode")
		if err != nil {
			fmt.Fprintf(errW, "[WARN] reading value of %q flag is failed: %s", "argument-match-mode", err) // nolint
			return
		}
		o.ArgumentMatchMode = v
	}
}

func stringFromEnv(key, defaultValue string) string {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	return v
}

func stringFromEnvMulti(keys []string, defaultValue string) string {
	for _, key := range keys {
		v := os.Getenv(key)
		if v != "" {
			return v
		}
	}
	return defaultValue
}

func stringSliceFromEnv(key string, defaultValue []string) []string {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	values := strings.Split(v, ",")
	for i := range values {
		values[i] = strings.Trim(values[i], " ")
	}
	return values
}

func intFromEnv(key string, defaultValue int) int {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return defaultValue
	}
	return int(i)
}

func (o *Config) Validate(skipCred bool) error {
	var errs []error

	if !skipCred {
		if o.AccessToken == "" {
			errs = append(errs, validate.NewFlagError("--token", "required"))
		}
		if o.AccessTokenSecret == "" {
			errs = append(errs, validate.NewFlagError("--secret", "required"))
		}
	}
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
