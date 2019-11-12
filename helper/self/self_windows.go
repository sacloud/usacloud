// Copyright 2017-2019 The Usacloud Authors
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

// +build windows

package self

import (
	"fmt"
	"os/exec"
	"regexp"
)

func ID() (string, error) {
	o, err := exec.Command("wmic", "path", "win32_computersystemproduct", "get", "IdentifyingNumber", "/VALUE").Output()
	if err != nil {
		return "", err
	}
	r := regexp.MustCompile(`IdentifyingNumber=([0-9]{12})`)
	groups := r.FindStringSubmatch(string(o))
	if len(groups) > 1 {
		return groups[1], nil
	}
	return "", fmt.Errorf("Can't find IdentifyingNumber(WMI)")
}
