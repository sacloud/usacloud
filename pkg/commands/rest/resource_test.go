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

package rest

import (
	"strings"
	"testing"
)

func TestHelpIncludesRESTRequestUsage(t *testing.T) {
	cmd := Resource.CLICommand()

	if got, want := cmd.Use, "rest URL"; got != want {
		t.Errorf("resource usage = %q, want %q", got, want)
	}
	if got, want := cmd.Long, Resource.Usage; got != want {
		t.Errorf("resource help = %q, want %q", got, want)
	}
	if want := "usacloud rest https://secure.sakura.ad.jp/cloud/zone/is1a/api/cloud/1.1/server"; !strings.Contains(cmd.Long, want) {
		t.Errorf("resource help does not include absolute URL example %q", want)
	}

	request, _, err := cmd.Find([]string{"request"})
	if err != nil {
		t.Fatalf("finding request command: %v", err)
	}
	if got, want := request.Use, "request URL"; got != want {
		t.Errorf("request usage = %q, want %q", got, want)
	}
	for name, want := range map[string]string{
		"data":   "JSON request body or path to a JSON file",
		"method": "HTTP method options: [get/post/patch/put/delete/GET/POST/PATCH/PUT/DELETE]",
		"zone":   "Zone for a relative URL (defaults to the active profile zone)",
	} {
		if got := request.Flags().Lookup(name).Usage; got != want {
			t.Errorf("--%s description = %q, want %q", name, got, want)
		}
	}
}

func TestMissingURLMessageIncludesUsage(t *testing.T) {
	err := missingURLError(0)

	for _, want := range []string{
		"URL: required (received 0 arguments)",
		"/server",
		"https://secure.sakura.ad.jp/cloud/zone/is1a/api/cloud/1.1/server",
		"usacloud rest --help",
	} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("missing URL error does not include %q: %s", want, err)
		}
	}
}
