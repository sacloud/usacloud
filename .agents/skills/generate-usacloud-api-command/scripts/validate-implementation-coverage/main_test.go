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
	"path/filepath"
	"strings"
	"testing"
)

func TestCommandNames(t *testing.T) {
	got, err := commandNames("./" + filepath.Join("testdata", "commands"))
	if err != nil {
		t.Fatal(err)
	}
	if !got[commandIdentity{Resource: "widget", Command: "list"}] ||
		!got[commandIdentity{Resource: "widget", Command: "delete"}] ||
		!got[commandIdentity{Resource: "gadget", Command: "list"}] {
		t.Errorf("expected core.Resource-derived command identities, got %#v", got)
	}
	if got[commandIdentity{Resource: "gadget", Command: "delete"}] {
		t.Errorf("gadget delete without Func must not be covered: %#v", got)
	}
}

func TestValidateManifest(t *testing.T) {
	err := validateManifest(manifest{
		Interfaces: []interfaceManifest{{
			Interface: "WidgetAPI",
			Resource:  "widget",
			Methods:   []method{{Name: "Read", Command: "read"}},
		}},
	}, true)
	if err != nil {
		t.Fatal(err)
	}
	if err := validateManifest(manifest{
		Interfaces: []interfaceManifest{{
			Interface: "WidgetAPI",
			Resource:  "widget",
			Methods:   []method{{Name: "Read"}},
		}},
	}, true); err == nil {
		t.Fatal("expected strict command-name rejection")
	}
}

func TestValidateCoverageReportsMissingResourceCommand(t *testing.T) {
	m := manifest{Interfaces: []interfaceManifest{
		{
			Interface: "WidgetAPI",
			Resource:  "widget",
			Methods:   []method{{Name: "List", Command: "list"}, {Name: "Delete", Command: "delete"}},
		},
		{
			Interface: "GadgetAPI",
			Resource:  "gadget",
			Methods:   []method{{Name: "List", Command: "list"}, {Name: "Delete", Command: "delete"}},
		},
	}}
	err := validateCoverage(m, "./"+filepath.Join("testdata", "commands"), true)
	if err == nil {
		t.Fatal("expected missing gadget delete command")
	}
	if !strings.Contains(err.Error(), "GadgetAPI/gadget: Delete -> delete") {
		t.Fatalf("missing identity from error: %v", err)
	}
}

func TestValidateManifestAllowsKnownExcludedMethod(t *testing.T) {
	m := manifest{Interfaces: []interfaceManifest{{
		Interface: "WidgetAPI",
		Resource:  "widget",
		Methods:   []method{{Name: "List", Command: "list"}, {Name: "Delete", Command: "delete"}},
		Excluded:  []excluded{{Method: "Delete", Reason: "read-only fixture"}},
	}}}
	if err := validateManifest(m, true); err != nil {
		t.Fatalf("known excluded method should be allowed: %v", err)
	}
	if got := methodCount(m); got != 1 {
		t.Errorf("methodCount() = %d, want 1 after exclusion", got)
	}
	m.Interfaces[0].Excluded[0].Method = "Unknown"
	if err := validateManifest(m, true); err == nil {
		t.Fatal("expected unknown excluded method rejection")
	}
}

func TestValidateCoverageSkipsExcludedMethod(t *testing.T) {
	m := manifest{Interfaces: []interfaceManifest{{
		Interface: "GadgetAPI",
		Resource:  "gadget",
		Methods:   []method{{Name: "Delete", Command: "delete"}},
		Excluded:  []excluded{{Method: "Delete", Reason: "not exposed by the CLI"}},
	}}}
	if err := validateManifest(m, true); err != nil {
		t.Fatal(err)
	}
	if err := validateCoverage(m, "./"+filepath.Join("testdata", "commands"), true); err != nil {
		t.Fatalf("excluded method must not require a command: %v", err)
	}
}
