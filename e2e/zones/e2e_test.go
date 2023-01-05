// Copyright 2017-2023 The sacloud/usacloud Authors
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

package zones

import (
	"strings"
	"testing"

	"github.com/sacloud/packages-go/e2e"
	usacloudE2E "github.com/sacloud/usacloud/e2e"
	"github.com/stretchr/testify/require"
)

func TestE2E_multiZoneAction(t *testing.T) {
	defer e2e.InitializeWithTerraform(t)()

	// **************************************************************
	// step1: テスト対象の全サーバの電源状態を取得
	// **************************************************************
	outputs, err := usacloudE2E.UsacloudRunWithOutput(t,
		"server", "list",
		"--names", "usacloud-e2e-zones-server",
		"--format", "{{.InstanceStatus}}",
		"--zone", "all",
	)
	lines := strings.Split(string(outputs), "\n")

	require.NoError(t, err)
	require.Len(t, lines, 5) // 4ゾーン分+最後の改行分
	for _, line := range lines {
		if line == "" {
			continue
		}
		require.Equal(t, "up", line)
	}

	// **************************************************************
	// step2: 全サーバの電源OFF
	// **************************************************************
	_, err = usacloudE2E.UsacloudRunWithOutput(t,
		"server", "shutdown",
		"-f", "-y",
		"--zone", "all",
		"usacloud-e2e-zones-server",
	)
	require.NoError(t, err)

	// **************************************************************
	// step2: テスト対象の全サーバの電源状態を取得
	// **************************************************************
	outputs, err = usacloudE2E.UsacloudRunWithOutput(t,
		"server", "list",
		"--names", "usacloud-e2e-zones-server",
		"--format", "{{.InstanceStatus}}",
		"--zone", "all",
	)
	lines = strings.Split(string(outputs), "\n")

	require.NoError(t, err)
	require.Len(t, lines, 5) // 4ゾーン分+最後の改行分
	for _, line := range lines {
		if line == "" {
			continue
		}
		require.Equal(t, "down", line)
	}
}
