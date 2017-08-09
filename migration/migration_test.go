package migration

import (
	"github.com/sacloud/usacloud/command/profile"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestMigrateConfig(t *testing.T) {

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
