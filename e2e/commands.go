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

package e2e

import (
	"os"
	"testing"

	"github.com/sacloud/packages-go/e2e"
)

func UsacloudRun(t *testing.T, args ...string) error {
	return e2e.RunCommand(t, usacloudCmdName(t), args...)
}

func UsacloudRunWithOutput(t *testing.T, args ...string) ([]byte, error) {
	return e2e.RunCommandWithOutput(t, usacloudCmdName(t), args...)
}

func UsacloudRunWithCombinedOutput(t *testing.T, args ...string) ([]byte, error) {
	return e2e.RunCommandWithCombinedOutput(t, usacloudCmdName(t), args...)
}

func usacloudCmdName(t *testing.T) string {
	cmd := "usacloud"
	if overwrite := os.Getenv("USACLOUD_COMMAND"); overwrite != "" {
		cmd = overwrite
		t.Logf("using `usacloud` from custom path: %s", cmd)
	}
	return cmd
}
