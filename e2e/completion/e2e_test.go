// Copyright 2017-2022 The sacloud/usacloud Authors
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
	"strings"
	"testing"

	"github.com/sacloud/usacloud/e2e"
	"github.com/stretchr/testify/require"
)

func TestE2E_complete_old_iaas_command(t *testing.T) {
	output, err := e2e.UsacloudRunWithOutput(t, "completion", "bash")

	require.NoError(t, err)
	require.NotEmpty(t, output)

	// iaasコマンドの例としてauto-backupを利用する。
	// commands+=("auto-backup")という文字列がルートコマンドとiaasサブコマンドの2箇所にあるはず
	require.Equal(t, 2, strings.Count(string(output), `commands+=("auto-backup")`))
}
