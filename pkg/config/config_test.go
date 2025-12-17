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
	"os"
	"reflect"
	"testing"

	"github.com/sacloud/api-client-go/profile"
	"github.com/sacloud/iaas-api-go"
)

func TestFillDefaults_ZonesEmpty(t *testing.T) {
	c := &Config{}
	c.fillDefaults()
	want := iaas.SakuraCloudZones
	want = append(want, "all")
	if !reflect.DeepEqual(c.Zones, want) {
		t.Errorf("Zones: got %v, want %v", c.Zones, want)
	}
}

func TestFillDefaults_ZonesWithAll(t *testing.T) {
	c := &Config{
		ConfigValue: profile.ConfigValue{
			Zones: []string{"is1a", "all"},
		},
	}
	c.fillDefaults()
	want := []string{"is1a", "all"}
	if !reflect.DeepEqual(c.Zones, want) {
		t.Errorf("Zones: got %v, want %v", c.Zones, want)
	}
}

func TestFillDefaults_ZonesWithoutAll(t *testing.T) {
	c := &Config{
		ConfigValue: profile.ConfigValue{
			Zones: []string{"is1a", "tk1a"},
		},
	}
	c.fillDefaults()
	want := []string{"is1a", "tk1a", "all"}
	if !reflect.DeepEqual(c.Zones, want) {
		t.Errorf("Zones: got %v, want %v", c.Zones, want)
	}
}

func TestFillDefaults_EnvVar(t *testing.T) {
	os.Setenv("SAKURA_ZONES", "is1a,tk1a") //nolint:errcheck,gosec
	defer os.Unsetenv("SAKURA_ZONES")      //nolint:errcheck
	c := &Config{}
	c.loadFromEnv()
	c.fillDefaults()
	want := []string{"is1a", "tk1a", "all"}
	if !reflect.DeepEqual(c.Zones, want) {
		t.Errorf("Zones: got %v, want %v", c.Zones, want)
	}
}

// got と want の差分をまとめて見るためのヘルパ
func assertConfigEqual(t *testing.T, got, want Config) {
	t.Helper()

	if got.AccessToken != want.AccessToken {
		t.Errorf("AccessToken = %q, want %q", got.AccessToken, want.AccessToken)
	}
	if got.AccessTokenSecret != want.AccessTokenSecret {
		t.Errorf("AccessTokenSecret = %q, want %q", got.AccessTokenSecret, want.AccessTokenSecret)
	}
	if got.Zone != want.Zone {
		t.Errorf("Zone = %q, want %q", got.Zone, want.Zone)
	}
	if len(got.Zones) != len(want.Zones) {
		t.Errorf("len(Zones) = %d, want %d", len(got.Zones), len(want.Zones))
	} else {
		for i := range got.Zones {
			if got.Zones[i] != want.Zones[i] {
				t.Errorf("Zones[%d] = %q, want %q", i, got.Zones[i], want.Zones[i])
			}
		}
	}
	if got.AcceptLanguage != want.AcceptLanguage {
		t.Errorf("AcceptLanguage = %q, want %q", got.AcceptLanguage, want.AcceptLanguage)
	}
	if got.RetryMax != want.RetryMax {
		t.Errorf("RetryMax = %d, want %d", got.RetryMax, want.RetryMax)
	}
	if got.RetryWaitMax != want.RetryWaitMax {
		t.Errorf("RetryWaitMax = %d, want %d", got.RetryWaitMax, want.RetryWaitMax)
	}
	if got.RetryWaitMin != want.RetryWaitMin {
		t.Errorf("RetryWaitMin = %d, want %d", got.RetryWaitMin, want.RetryWaitMin)
	}
	if got.HTTPRequestTimeout != want.HTTPRequestTimeout {
		t.Errorf("HTTPRequestTimeout = %d, want %d", got.HTTPRequestTimeout, want.HTTPRequestTimeout)
	}
	if got.HTTPRequestRateLimit != want.HTTPRequestRateLimit {
		t.Errorf("HTTPRequestRateLimit = %d, want %d", got.HTTPRequestRateLimit, want.HTTPRequestRateLimit)
	}
	if got.APIRootURL != want.APIRootURL {
		t.Errorf("APIRootURL = %q, want %q", got.APIRootURL, want.APIRootURL)
	}
	if got.DefaultZone != want.DefaultZone {
		t.Errorf("DefaultZone = %q, want %q", got.DefaultZone, want.DefaultZone)
	}
	if got.TraceMode != want.TraceMode {
		t.Errorf("TraceMode = %q, want %q", got.TraceMode, want.TraceMode)
	}
	if got.FakeMode != want.FakeMode {
		t.Errorf("FakeMode = %v, want %v", got.FakeMode, want.FakeMode)
	}
	if got.FakeStorePath != want.FakeStorePath {
		t.Errorf("FakeStorePath = %q, want %q", got.FakeStorePath, want.FakeStorePath)
	}
	if got.ProcessTimeoutSec != want.ProcessTimeoutSec {
		t.Errorf("ProcessTimeoutSec = %d, want %d", got.ProcessTimeoutSec, want.ProcessTimeoutSec)
	}
	if got.ArgumentMatchMode != want.ArgumentMatchMode {
		t.Errorf("ArgumentMatchMode = %q, want %q", got.ArgumentMatchMode, want.ArgumentMatchMode)
	}
	if got.DefaultOutputType != want.DefaultOutputType {
		t.Errorf("DefaultOutputType = %q, want %q", got.DefaultOutputType, want.DefaultOutputType)
	}
	if got.DefaultQueryDriver != want.DefaultQueryDriver {
		t.Errorf("DefaultQueryDriver = %q, want %q", got.DefaultQueryDriver, want.DefaultQueryDriver)
	}
}

// 1つのテスト関数＋テーブルで「SAKURAのみ / SAKURACLOUDのみ / 両方 / なし」を網羅
func TestConfig_loadFromEnv(t *testing.T) {
	tests := []struct {
		name string
		env  map[string]string
		want Config
	}{
		{
			name: "SAKURA_* only",
			env: map[string]string{
				"SAKURA_ACCESS_TOKEN":           "token-sakura",
				"SAKURA_ACCESS_TOKEN_SECRET":    "secret-sakura",
				"SAKURA_ZONE":                   "is1a",
				"SAKURA_ZONES":                  "is1a,is1b",
				"SAKURA_ACCEPT_LANGUAGE":        "ja-JP",
				"SAKURA_RETRY_MAX":              "10",
				"SAKURA_RETRY_WAIT_MAX":         "30",
				"SAKURA_RETRY_WAIT_MIN":         "1",
				"SAKURA_API_REQUEST_TIMEOUT":    "100",
				"SAKURA_API_REQUEST_RATE_LIMIT": "7",
				"SAKURA_API_ROOT_URL":           "https://api.sakura.example",
				"SAKURA_DEFAULT_ZONE":           "tk1a",
				"SAKURA_TRACE":                  "1",
				"SAKURA_FAKE_MODE":              "1",
				"SAKURA_FAKE_STORE_PATH":        "/tmp/sakura",
				"SAKURA_PROCESS_TIMEOUT_SEC":    "600",
				"SAKURA_ARGUMENT_MATCH_MODE":    "exact",
				"SAKURA_DEFAULT_OUTPUT_TYPE":    "json",
				"SAKURA_DEFAULT_QUERY_DRIVER":   "jmespath",
			},
			want: Config{
				ConfigValue: profile.ConfigValue{
					AccessToken:          "token-sakura",
					AccessTokenSecret:    "secret-sakura",
					Zone:                 "is1a",
					Zones:                []string{"is1a", "is1b"},
					AcceptLanguage:       "ja-JP",
					RetryMax:             10,
					RetryWaitMax:         30,
					RetryWaitMin:         1,
					HTTPRequestTimeout:   100,
					HTTPRequestRateLimit: 7,
					APIRootURL:           "https://api.sakura.example",
					DefaultZone:          "tk1a",
					TraceMode:            "1",
					FakeMode:             true,
					FakeStorePath:        "/tmp/sakura",
				},
				ProcessTimeoutSec:  600,
				ArgumentMatchMode:  "exact",
				DefaultOutputType:  "json",
				DefaultQueryDriver: "jmespath",
			},
		},
		{
			name: "SAKURACLOUD_* only",
			env: map[string]string{
				"SAKURACLOUD_ACCESS_TOKEN":           "token-cloud",
				"SAKURACLOUD_ACCESS_TOKEN_SECRET":    "secret-cloud",
				"SAKURACLOUD_ZONE":                   "is1b",
				"SAKURACLOUD_ZONES":                  "tk1a,tk1b",
				"SAKURACLOUD_ACCEPT_LANGUAGE":        "en-US",
				"SAKURACLOUD_RETRY_MAX":              "20",
				"SAKURACLOUD_RETRY_WAIT_MAX":         "40",
				"SAKURACLOUD_RETRY_WAIT_MIN":         "2",
				"SAKURACLOUD_API_REQUEST_TIMEOUT":    "200",
				"SAKURACLOUD_API_REQUEST_RATE_LIMIT": "8",
				"SAKURACLOUD_API_ROOT_URL":           "https://api.cloud.example",
				"SAKURACLOUD_DEFAULT_ZONE":           "tk1b",
				"SAKURACLOUD_TRACE":                  "2",
				"SAKURACLOUD_FAKE_MODE":              "1",
				"SAKURACLOUD_FAKE_STORE_PATH":        "/tmp/cloud",
				"SAKURACLOUD_PROCESS_TIMEOUT_SEC":    "700",
				"SAKURACLOUD_ARGUMENT_MATCH_MODE":    "prefix",
				"SAKURACLOUD_DEFAULT_OUTPUT_TYPE":    "table",
				"SAKURACLOUD_DEFAULT_QUERY_DRIVER":   "jq",
			},
			want: Config{
				ConfigValue: profile.ConfigValue{
					AccessToken:          "token-cloud",
					AccessTokenSecret:    "secret-cloud",
					Zone:                 "is1b",
					Zones:                []string{"tk1a", "tk1b"},
					AcceptLanguage:       "en-US",
					RetryMax:             20,
					RetryWaitMax:         40,
					RetryWaitMin:         2,
					HTTPRequestTimeout:   200,
					HTTPRequestRateLimit: 8,
					APIRootURL:           "https://api.cloud.example",
					DefaultZone:          "tk1b",
					TraceMode:            "2",
					FakeMode:             true,
					FakeStorePath:        "/tmp/cloud",
				},
				ProcessTimeoutSec:  700,
				ArgumentMatchMode:  "prefix",
				DefaultOutputType:  "table",
				DefaultQueryDriver: "jq",
			},
		},
		{
			name: "both SAKURA_* and SAKURACLOUD_* (SAKURA_* preferred)",
			env: func() map[string]string {
				m := map[string]string{
					"SAKURA_ACCESS_TOKEN":           "token-sakura",
					"SAKURA_ACCESS_TOKEN_SECRET":    "secret-sakura",
					"SAKURA_ZONE":                   "is1a",
					"SAKURA_ZONES":                  "is1a,is1b",
					"SAKURA_ACCEPT_LANGUAGE":        "ja-JP",
					"SAKURA_RETRY_MAX":              "10",
					"SAKURA_RETRY_WAIT_MAX":         "30",
					"SAKURA_RETRY_WAIT_MIN":         "1",
					"SAKURA_API_REQUEST_TIMEOUT":    "100",
					"SAKURA_API_REQUEST_RATE_LIMIT": "7",
					"SAKURA_API_ROOT_URL":           "https://api.sakura.example",
					"SAKURA_DEFAULT_ZONE":           "tk1a",
					"SAKURA_TRACE":                  "1",
					"SAKURA_FAKE_MODE":              "1",
					"SAKURA_FAKE_STORE_PATH":        "/tmp/sakura",
					"SAKURA_PROCESS_TIMEOUT_SEC":    "600",
					"SAKURA_ARGUMENT_MATCH_MODE":    "exact",
					"SAKURA_DEFAULT_OUTPUT_TYPE":    "json",
					"SAKURA_DEFAULT_QUERY_DRIVER":   "jmespath",
					// SAKURACLOUD_* 側は異なる値を入れておくが、結果は SAKURA_* が優先されることを期待
					"SAKURACLOUD_ACCESS_TOKEN":           "token-cloud",
					"SAKURACLOUD_ACCESS_TOKEN_SECRET":    "secret-cloud",
					"SAKURACLOUD_ZONE":                   "is1b",
					"SAKURACLOUD_ZONES":                  "tk1a,tk1b",
					"SAKURACLOUD_ACCEPT_LANGUAGE":        "en-US",
					"SAKURACLOUD_RETRY_MAX":              "20",
					"SAKURACLOUD_RETRY_WAIT_MAX":         "40",
					"SAKURACLOUD_RETRY_WAIT_MIN":         "2",
					"SAKURACLOUD_API_REQUEST_TIMEOUT":    "200",
					"SAKURACLOUD_API_REQUEST_RATE_LIMIT": "8",
					"SAKURACLOUD_API_ROOT_URL":           "https://api.cloud.example",
					"SAKURACLOUD_DEFAULT_ZONE":           "tk1b",
					"SAKURACLOUD_TRACE":                  "2",
					"SAKURACLOUD_FAKE_MODE":              "1",
					"SAKURACLOUD_FAKE_STORE_PATH":        "/tmp/cloud",
					"SAKURACLOUD_PROCESS_TIMEOUT_SEC":    "700",
					"SAKURACLOUD_ARGUMENT_MATCH_MODE":    "prefix",
					"SAKURACLOUD_DEFAULT_OUTPUT_TYPE":    "table",
					"SAKURACLOUD_DEFAULT_QUERY_DRIVER":   "jq",
				}
				return m
			}(),
			want: Config{
				ConfigValue: profile.ConfigValue{
					AccessToken:          "token-sakura",
					AccessTokenSecret:    "secret-sakura",
					Zone:                 "is1a",
					Zones:                []string{"is1a", "is1b"},
					AcceptLanguage:       "ja-JP",
					RetryMax:             10,
					RetryWaitMax:         30,
					RetryWaitMin:         1,
					HTTPRequestTimeout:   100,
					HTTPRequestRateLimit: 7,
					APIRootURL:           "https://api.sakura.example",
					DefaultZone:          "tk1a",
					TraceMode:            "1",
					FakeMode:             true,
					FakeStorePath:        "/tmp/sakura",
				},
				ProcessTimeoutSec:  600,
				ArgumentMatchMode:  "exact",
				DefaultOutputType:  "json",
				DefaultQueryDriver: "jmespath",
			},
		},
		{
			name: "no env (defaults)",
			env:  map[string]string{},
			want: Config{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clearTestEnv()
			setEnv(tt.env)

			var cfg Config
			cfg.loadFromEnv()

			assertConfigEqual(t, cfg, tt.want)
		})
	}
}

func TestConfig_loadFromEnvOverwrite(t *testing.T) {
	cfg := Config{
		ConfigValue: profile.ConfigValue{
			AccessToken:       "token from config",
			AccessTokenSecret: "secret from config",
		},
	}

	clearTestEnv()
	setEnv(map[string]string{
		"SAKURACLOUD_ACCESS_TOKEN": "token from env",
	})
	cfg.loadFromEnv()

	if cfg.AccessToken != "token from env" {
		t.Fatalf("got unexpected value: expected: %s, actual: %s", "token from env", cfg.AccessToken)
	}
	if cfg.AccessTokenSecret != "secret from config" {
		t.Fatalf("got unexpected value: expected: %s, actual: %s", "secret from config", cfg.AccessTokenSecret)
	}
}

// テストで使う環境変数一覧
var allEnvKeys = []string{
	"SAKURA_ACCESS_TOKEN", "SAKURACLOUD_ACCESS_TOKEN",
	"SAKURA_ACCESS_TOKEN_SECRET", "SAKURACLOUD_ACCESS_TOKEN_SECRET",
	"SAKURA_ZONE", "SAKURACLOUD_ZONE",
	"SAKURA_ZONES", "SAKURACLOUD_ZONES",
	"SAKURA_ACCEPT_LANGUAGE", "SAKURACLOUD_ACCEPT_LANGUAGE",
	"SAKURA_RETRY_MAX", "SAKURACLOUD_RETRY_MAX",
	"SAKURA_RETRY_WAIT_MAX", "SAKURACLOUD_RETRY_WAIT_MAX",
	"SAKURA_RETRY_WAIT_MIN", "SAKURACLOUD_RETRY_WAIT_MIN",
	"SAKURA_API_REQUEST_TIMEOUT", "SAKURACLOUD_API_REQUEST_TIMEOUT",
	"SAKURA_API_REQUEST_RATE_LIMIT", "SAKURACLOUD_API_REQUEST_RATE_LIMIT",
	"SAKURA_API_ROOT_URL", "SAKURACLOUD_API_ROOT_URL",
	"SAKURA_DEFAULT_ZONE", "SAKURACLOUD_DEFAULT_ZONE",
	"SAKURA_TRACE", "SAKURACLOUD_TRACE",
	"SAKURA_FAKE_MODE", "SAKURACLOUD_FAKE_MODE",
	"SAKURA_FAKE_STORE_PATH", "SAKURACLOUD_FAKE_STORE_PATH",
	"SAKURA_PROCESS_TIMEOUT_SEC", "SAKURACLOUD_PROCESS_TIMEOUT_SEC",
	"SAKURA_ARGUMENT_MATCH_MODE", "SAKURACLOUD_ARGUMENT_MATCH_MODE",
	"SAKURA_DEFAULT_OUTPUT_TYPE", "SAKURACLOUD_DEFAULT_OUTPUT_TYPE",
	"SAKURA_DEFAULT_QUERY_DRIVER", "SAKURACLOUD_DEFAULT_QUERY_DRIVER",
}

func clearTestEnv() {
	for _, k := range allEnvKeys {
		_ = os.Unsetenv(k)
	}
}

func setEnv(m map[string]string) {
	for k, v := range m {
		_ = os.Setenv(k, v)
	}
}
