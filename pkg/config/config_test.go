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
	os.Setenv("SAKURACLOUD_ZONES", "is1a,tk1a") //nolint:errcheck,gosec
	defer os.Unsetenv("SAKURACLOUD_ZONES")      //nolint:errcheck
	c := &Config{}
	c.loadFromEnv()
	c.fillDefaults()
	want := []string{"is1a", "tk1a", "all"}
	if !reflect.DeepEqual(c.Zones, want) {
		t.Errorf("Zones: got %v, want %v", c.Zones, want)
	}
}
