// Copyright 2017-2020 The Usacloud Authors
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

	"github.com/sacloud/usacloud/pkg/util"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/profile"
	"github.com/sacloud/usacloud/pkg/validation"
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
}

// InitConfig 指定のFlagSetにConfigへ値を設定するためのフラグを登録する
func InitConfig(flags *pflag.FlagSet) {
	initCredentialConfig(flags)
	initOutputConfig(flags)
	initDebugConfig(flags)
}

// LoadConfigValue 指定のフラグセットからフラグを読み取り*Flagsを組み立てて返す
func LoadConfigValue(flags *pflag.FlagSet, errW io.Writer) (*Config, error) {
	o := &Config{}
	o.loadConfig(flags, errW)

	return o, util.FlattenErrors(o.Validate(true))
}

func initCredentialConfig(fs *pflag.FlagSet) {
	fs.StringP("profile", "", "default", "the name of saved credentials")
	fs.StringP("token", "", "", "the API token used when calling SAKURA Cloud API")
	fs.StringP("secret", "", "", "the API secret used when calling SAKURA Cloud API")
	fs.StringSliceP("zones", "", []string{}, "permitted zone names")
}

func initOutputConfig(fs *pflag.FlagSet) {
	fs.BoolP("no-color", "", false, "disable ANSI color output")
}

func initDebugConfig(fs *pflag.FlagSet) {
	fs.BoolP("trace", "", false, "enable trace logs for API calling")
	fs.BoolP("fake", "", false, "enable fake API driver")
	fs.StringP("fake-store", "", "", "path to file store used by the fake API driver")
}

func (o *Config) IsEmpty() bool {
	return o.AccessToken == "" &&
		o.AccessTokenSecret == "" &&
		o.Zone == "" && o.DefaultOutputType == ""
}

func (o *Config) loadConfig(flags *pflag.FlagSet, errW io.Writer) {
	o.loadFromEnv()
	o.loadFromProfile(errW)
	o.loadFromFlags(flags, errW)
	o.fillDefaults()
}

func (o *Config) fillDefaults() {
	if len(o.Zones) == 0 {
		o.Zones = sacloud.SakuraCloudZones
	}
}

func (o *Config) loadFromEnv() {
	o.Profile = stringFromEnv("SAKURACLOUD_PROFILE", "default")
	o.AccessToken = stringFromEnv("SAKURACLOUD_ACCESS_TOKEN", "")
	o.AccessTokenSecret = stringFromEnv("SAKURACLOUD_ACCESS_TOKEN_SECRET", "")
	o.Zone = stringFromEnv("SAKURACLOUD_ZONE", "")
	o.Zones = stringSliceFromEnv("SAKURACLOUD_ZONES", append(sacloud.SakuraCloudZones, "all"))
	o.AcceptLanguage = stringFromEnv("SAKURACLOUD_ACCEPT_LANGUAGE", "")
	o.RetryMax = intFromEnv("SAKURACLOUD_RETRY_MAX", sacloud.APIDefaultRetryMax)
	o.RetryWaitMax = intFromEnv("SAKURACLOUD_RETRY_WAIT_MAX", 64)
	o.RetryWaitMin = intFromEnv("SAKURACLOUD_RETRY_WAIT_MIN", 1)
	o.HTTPRequestTimeout = intFromEnv("SAKURACLOUD_API_REQUEST_TIMEOUT", 300)
	o.HTTPRequestRateLimit = intFromEnv("SAKURACLOUD_API_REQUEST_RATE_LIMIT", 5) // デフォルト5ゾーン分(is1a/is1b/tk1a/tk1b/tk1v)
	o.APIRootURL = stringFromEnv("SAKURACLOUD_API_ROOT_URL", sacloud.SakuraCloudAPIRoot)
	o.DefaultZone = stringFromEnv("SAKURACLOUD_DEFAULT_ZONE", sacloud.APIDefaultZone)
	o.TraceMode = stringFromEnv("SAKURACLOUD_TRACE", "")
	o.FakeMode = os.Getenv("SAKURACLOUD_FAKE_MODE") != ""
	o.FakeStorePath = stringFromEnv("SAKURACLOUD_FAKE_STORE_PATH", "")
}

func (o *Config) loadFromProfile(errW io.Writer) {
	if o.Profile != "" {
		if err := profile.Load(o.Profile, o); err != nil {
			fmt.Fprintf(errW, "[WARN] loading profile %q is failed: %s", o.Profile, err) // nolint
			return
		}
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
}

func stringFromEnv(key, defaultValue string) string {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	return v
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

func (o *Config) Validate(skipCred bool) []error {
	var errs []error

	if !skipCred {
		errs = append(errs, validation.Required("token", o.AccessToken)...)
		errs = append(errs, validation.Required("secret", o.AccessTokenSecret)...)
	}
	errs = append(errs, validation.StringInSlice("default-output-type", o.DefaultOutputType, []string{"table", "json", "yaml"})...)

	return errs
}
