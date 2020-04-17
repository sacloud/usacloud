// Copyright 2017-2020 The Usacloud Authors
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

package migration

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/sacloud/usacloud/pkg/cli"

	"github.com/sacloud/usacloud/pkg/profile"
	"github.com/stretchr/testify/assert"
)

func TestMigrateConfig(t *testing.T) {
	// TODO グローバルオプションの扱いが確定したら修正する
	cli.GlobalOption = &cli.CLIOptions{}

	initFunc := func() func() {
		confirmMigrateFunc = func() bool { return true }
		confirmOverwriteFunc = func(s string) bool { return true }

		cleanupDefaultProfile()
		cleanupOldConfig()
		return func() {
			cleanupDefaultProfile()
			cleanupOldConfig()
		}
	}
	defer initFunc()()

	t.Run("Valid old config", func(t *testing.T) {
		defer initFunc()()

		createOldConfig()

		// do migration
		err := MigrateConfig()
		assert.NoError(t, err)

		v, err := profile.LoadConfigFile(profile.DefaultProfileName)
		assert.NoError(t, err)

		assert.EqualValues(t, "test-token", v.AccessToken)
		assert.EqualValues(t, "test-secret", v.AccessTokenSecret)
		assert.EqualValues(t, "tk1v", v.Zone)

		src, err := getOldConfigPath()
		assert.NoError(t, err)

		fileExists := false
		if _, err := os.Stat(src); err == nil {
			fileExists = true
		}
		assert.False(t, fileExists)
	})
	t.Run("Answer 'No' to confirm", func(t *testing.T) {
		defer initFunc()()
		confirmMigrateFunc = func() bool { return false }
		createOldConfig()

		// do migration
		err := MigrateConfig()
		assert.NoError(t, err)

		// src isnot deleted
		src, err := getOldConfigPath()
		assert.NoError(t, err)

		fileExists := false
		if _, err := os.Stat(src); err == nil {
			fileExists = true
		}
		assert.True(t, fileExists)

		// dest isnot created
		dest, err := profile.GetConfigFilePath(profile.DefaultProfileName)
		assert.NoError(t, err)

		fileExists = false
		if _, err := os.Stat(dest); err == nil {
			fileExists = true
		}
		assert.False(t, fileExists)

	})
	t.Run("Answer 'No' to overwrite confirm", func(t *testing.T) {
		defer initFunc()()
		confirmOverwriteFunc = func(s string) bool { return false }
		createOldConfig()
		createDefaultProfile()

		// do migration
		err := MigrateConfig()
		assert.NoError(t, err)

		// src isnot deleted
		src, err := getOldConfigPath()
		assert.NoError(t, err)

		fileExists := false
		if _, err := os.Stat(src); err == nil {
			fileExists = true
		}
		assert.True(t, fileExists)

		// dest isnot changed
		v, err := profile.LoadConfigFile(profile.DefaultProfileName)
		assert.NoError(t, err)
		assert.EqualValues(t, "default-token", v.AccessToken)
		assert.EqualValues(t, "default-secret", v.AccessTokenSecret)
		assert.EqualValues(t, "tk1v", v.Zone)

	})
	t.Run("Overwrite default config", func(t *testing.T) {
		defer initFunc()()
		createOldConfig()
		createDefaultProfile()

		// do migration
		err := MigrateConfig()
		assert.NoError(t, err)

		// src isnot deleted
		src, err := getOldConfigPath()
		assert.NoError(t, err)

		fileExists := false
		if _, err := os.Stat(src); err == nil {
			fileExists = true
		}
		assert.False(t, fileExists)

		// dest isnot changed
		v, err := profile.LoadConfigFile(profile.DefaultProfileName)
		assert.NoError(t, err)
		assert.EqualValues(t, "test-token", v.AccessToken)
		assert.EqualValues(t, "test-secret", v.AccessTokenSecret)
		assert.EqualValues(t, "tk1v", v.Zone)
	})
}

func createOldConfig() {
	path, err := getOldConfigPath()
	if err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(path, []byte(testOldConfigBody), 0600); err != nil {
		panic(err)
	}
}

func createDefaultProfile() {
	profile.SaveConfigFile(profile.DefaultProfileName, &profile.ConfigFileValue{
		AccessToken:       "default-token",
		AccessTokenSecret: "default-secret",
		Zone:              "tk1v",
	})
}

func cleanupDefaultProfile() {
	path, err := profile.GetConfigFilePath(profile.DefaultProfileName)
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(path); err == nil {
		if err := profile.RemoveConfigFile(profile.DefaultProfileName); err != nil {
			panic(err)
		}
	}
}

func cleanupOldConfig() {
	path, err := getOldConfigPath()
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(path); err == nil {
		err := os.Remove(path)
		if err != nil {
			panic(err)
		}
	}
}

var testOldConfigBody = `{
	"AccessToken": "test-token",
	"AccessTokenSecret": "test-secret",
	"Zone": "tk1v"
}`
