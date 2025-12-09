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

//go:build e2e
// +build e2e

package minimum

import (
	"os"
	"testing"

	"github.com/sacloud/saclient-go"
	usacloudE2E "github.com/sacloud/usacloud/e2e"
	"github.com/stretchr/testify/require"
)

func TestE2E_CreateConfigWithoutTTY(t *testing.T) {
	// setup
	profileDir := t.TempDir()
	os.Setenv("SAKURACLOUD_PROFILE_DIR", profileDir)

	profileName := "foobar"

	err := usacloudE2E.UsacloudRun(t, "config", "create", profileName, "--token", "aaa", "--secret", "bbb", "--default-output-type", "table", "--zone", "is1b", "--use")
	require.NoError(t, err)

	profileOp := saclient.NewProfileOp(os.Environ())
	profile, err := profileOp.Read(profileName)
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, "aaa", profile.Attributes["AccessToken"])
	require.Equal(t, "bbb", profile.Attributes["AccessTokenSecret"])
	require.Equal(t, "table", profile.Attributes["DefaultOutputType"])
	require.Equal(t, "is1b", profile.Attributes["Zone"])
}
