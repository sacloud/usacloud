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

package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestInventoryPackageGolden(t *testing.T) {
	meta := packageMeta{
		ImportPath: "fixture",
		Dir:        filepath.Join("testdata", "sdkfixture"),
		GoFiles:    []string{"api.go"},
	}
	got, err := inventoryPackage(meta, "")
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.MarshalIndent(got, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	want, err := os.ReadFile(filepath.Join("testdata", "interfaces.golden.json"))
	if err != nil {
		t.Fatal(err)
	}
	if string(data)+"\n" != string(want) {
		t.Errorf("inventory mismatch\nwant:\n%s\ngot:\n%s", want, data)
	}
}

func TestInventoryPackageFiltersInterface(t *testing.T) {
	meta := packageMeta{
		ImportPath: "fixture",
		Dir:        filepath.Join("testdata", "sdkfixture"),
		GoFiles:    []string{"api.go"},
	}
	got, err := inventoryPackage(meta, "WidgetAPI")
	if err != nil {
		t.Fatal(err)
	}
	if len(got.Interfaces) != 1 || got.Interfaces[0].Interface != "WidgetAPI" || got.Interfaces[0].Resource != "widget" {
		t.Fatalf("unexpected filtered inventory: %#v", got.Interfaces)
	}
}

func TestClassifyAndCommandName(t *testing.T) {
	cases := map[string]struct{ operation, command string }{
		"ListWidgets":  {"list", "list-widgets"},
		"GetWidget":    {"read", "read-widget"},
		"DeleteWidget": {"delete", "delete-widget"},
		"UpdateSecret": {"update", "update-secret"},
		"RotateAPIKey": {"action", "rotate-api-key"},
		"Rebuild":      {"action", "rebuild"},
	}
	for input, want := range cases {
		if got := classify(input); got != want.operation {
			t.Errorf("classify(%q) = %q, want %q", input, got, want.operation)
		}
		if got := commandName(input, want.operation); got != want.command {
			t.Errorf("commandName(%q) = %q, want %q", input, got, want.command)
		}
		if got := resourceName("ProcessConfigurationAPI"); got != "process-configuration" {
			t.Errorf("resourceName(ProcessConfigurationAPI) = %q, want process-configuration", got)
		}
	}
}
