// Copyright 2017-2026 The sacloud/usacloud Authors
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

package config

import (
	"errors"
	"io/fs"
	"slices"

	"github.com/sacloud/sacloud-sdk-go/common/saclient"
	"github.com/sacloud/usacloud/pkg/config"
)

// currentOrDefault returns the current profile name, falling back to "default"
// when the current file is missing, empty, or contains only whitespace.
func currentOrDefault(op saclient.ProfileAPI) (string, error) {
	name, err := op.GetCurrentName()
	if err == nil {
		if name == "" {
			return config.DefaultProfileName, nil
		}
		return name, nil
	}
	if errors.Is(err, fs.ErrNotExist) {
		return config.DefaultProfileName, nil
	}
	return "", err
}

// listWithDefault returns the profile list from SDK, ensuring "default" is included.
func listWithDefault(op saclient.ProfileAPI) ([]string, error) {
	names, err := op.List()
	if err != nil {
		return nil, err
	}
	if !slices.Contains(names, config.DefaultProfileName) {
		names = append([]string{config.DefaultProfileName}, names...)
	}
	return names, nil
}

// ensureDefault creates an empty default profile if it does not exist.
func ensureDefault(op saclient.ProfileAPI) error {
	_, err := op.Read(config.DefaultProfileName)
	if err == nil {
		return nil
	}
	if !errors.Is(err, fs.ErrNotExist) {
		return err
	}
	return op.Create(&saclient.Profile{Name: config.DefaultProfileName})
}

// setCurrentDefault switches current to "default", creating the profile if needed.
func setCurrentDefault(op saclient.ProfileAPI) error {
	if err := ensureDefault(op); err != nil {
		return err
	}
	return op.SetCurrentName(config.DefaultProfileName)
}
