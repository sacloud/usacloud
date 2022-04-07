// Copyright 2017-2022 The Usacloud Authors
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

package old_iaas_command

import (
	"testing"

	"github.com/sacloud/usacloud/e2e"
	"github.com/stretchr/testify/require"
)

func TestE2E_oldIaaSCommand(t *testing.T) {
	// **************************************************************
	// step1: iaasサブコマンドを実行できるか?
	// **************************************************************
	err := e2e.UsacloudRun(t,
		"server", "iaas", "list", "-h")
	require.NoError(t, err, "unexpected error: %s", err)

	// **************************************************************
	// step2: root直下のサブコマンド(Hidden=trueなコマンド)を実行できるか?
	// **************************************************************
	err = e2e.UsacloudRun(t,
		"server", "list", "-h")
	require.NoError(t, err, "unexpected error: %s", err)
}

func TestE2E_oldIaaSCommandOutputs(t *testing.T) {
	defer e2e.InitializeWithTerraform(t)()

	// **************************************************************
	// step1: iaasサブコマンドでサーバの情報を取得
	// **************************************************************
	outputs1, err := e2e.UsacloudRunWithOutput(t,
		"iaas", "server", "list",
		"--names", "usacloud-e2e-old-iaas-command",
		"--format", "{{.Name}}",
		"--zone", "all",
	)
	require.NoError(t, err, "unexpected error: %s", err)

	// **************************************************************
	// step2: root直下のサブコマンド(Hidden=trueなコマンド)でサーバの情報を取得
	// **************************************************************
	outputs2, err := e2e.UsacloudRunWithOutput(t,
		"server", "list",
		"--names", "usacloud-e2e-old-iaas-command",
		"--format", "{{.Name}}",
		"--zone", "all",
	)
	require.NoError(t, err, "unexpected error: %s", err)

	// **************************************************************
	// step3: 新旧コマンドの結果を比較
	// **************************************************************
	require.Equal(t, string(outputs1), string(outputs2))
}
