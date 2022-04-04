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

package minimum

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestE2E_minimum(t *testing.T) {
	cmd := exec.Command("usacloud", "-h")
	output, err := cmd.Output()

	require.NoError(t, err)
	require.NotEmpty(t, output)
}

func TestE2E_invalidSubCommand(t *testing.T) {
	cmd := exec.Command("usacloud", "invalid subcommand")
	err := cmd.Run()

	require.Error(t, err)
}
