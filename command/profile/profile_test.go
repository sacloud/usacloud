// +build !windows

package profile

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
)

func TestGetProfileBaseDir(t *testing.T) {
	os.Unsetenv(ProfileDirEnv)
	defer func() {
		os.Unsetenv(ProfileDirEnv)
	}()

	t.Run(fmt.Sprintf("Without %s env", ProfileDirEnv), func(t *testing.T) {
		homeDir, err := homedir.Dir()
		assert.NoError(t, err)

		baseDir, err := getProfileBaseDir()
		assert.NoError(t, err)

		assert.EqualValues(t, homeDir, baseDir)
	})

	t.Run(fmt.Sprintf("With %s env", ProfileDirEnv), func(t *testing.T) {
		t.Run("Valid env var", func(t *testing.T) {
			testDir := "/test"
			err := os.Setenv(ProfileDirEnv, testDir)
			assert.NoError(t, err)

			baseDir, err := getProfileBaseDir()
			assert.NoError(t, err)

			assert.EqualValues(t, testDir, baseDir)
		})

		t.Run("Redundant path", func(t *testing.T) {
			testDir := "/test/../test1/../test"
			expect := "/test"

			err := os.Setenv(ProfileDirEnv, testDir)
			assert.NoError(t, err)

			baseDir, err := getProfileBaseDir()
			assert.NoError(t, err)

			assert.EqualValues(t, expect, baseDir)

		})
		t.Run("Invalid env var", func(t *testing.T) {
			// - with filepath.ListSeparator
			testDir := "/test" + string([]rune{filepath.ListSeparator}) + "test"
			err := os.Setenv(ProfileDirEnv, testDir)
			assert.NoError(t, err)

			_, err = getProfileBaseDir()
			assert.Error(t, err)
		})
	})
}

func TestGetConfigFilePath(t *testing.T) {

	homeDir, err := homedir.Dir()
	assert.NoError(t, err)

	expects := []struct {
		profileName string
		filePath    string
	}{
		{
			profileName: "default",
			filePath:    "/.usacloud/default/config.json",
		},
		{
			profileName: "test1",
			filePath:    "/.usacloud/test1/config.json",
		},
		{
			profileName: "",
			filePath:    "/.usacloud/default/config.json",
		},
	}

	t.Run("Valid profiles", func(t *testing.T) {
		for _, expect := range expects {
			path, err := GetConfigFilePath(expect.profileName)
			assert.NoError(t, err)

			p := strings.Replace(path, homeDir, "", 1)
			assert.EqualValues(t, expect.filePath, p)
		}
	})

	t.Run("Invalid profiles", func(t *testing.T) {
		// - with filepath.Separator
		_, err = GetConfigFilePath("test" + string([]rune{filepath.Separator}) + "test")
		assert.Error(t, err)
		// - with filepath.ListSeparator
		_, err = GetConfigFilePath("test" + string([]rune{filepath.ListSeparator}) + "test")
		assert.Error(t, err)
	})
}

type loadConfigExpects struct {
	profileName string
	isValid     bool
	body        string
}

func testTargetProfiles() []loadConfigExpects {
	return []loadConfigExpects{
		{
			profileName: "default",
			isValid:     true,
			body:        fmt.Sprintf(confTemplate, "default", "default", "default"),
		},
		{
			profileName: "for-usacloud-unit-test1",
			isValid:     true,
			body:        fmt.Sprintf(confTemplate, "for-usacloud-unit-test1", "for-usacloud-unit-test1", "for-usacloud-unit-test1"),
		},
		{
			profileName: "for-usacloud-unit-test2",
			isValid:     true,
			body:        fmt.Sprintf(confTemplate, "for-usacloud-unit-test2", "for-usacloud-unit-test2", "for-usacloud-unit-test2"),
		},
		{
			profileName: " for-usacloud-unit-test3\n\n",
			isValid:     true,
			body:        fmt.Sprintf(confTemplate, "for-usacloud-unit-test3", "for-usacloud-unit-test3", "for-usacloud-unit-test3"),
		},
		{
			profileName: "invalid-json",
			isValid:     false,
			body:        "{",
		},
		{
			profileName: "empty-body",
			isValid:     false,
			body:        "",
		},
	}
}

func TestLoadConfigFile(t *testing.T) {
	defer initConfigFiles()()

	t.Run("Valid profiles", func(t *testing.T) {
		for _, prof := range testTargetProfiles() {
			conf, err := LoadConfigFile(prof.profileName)
			if prof.isValid {
				assert.NoError(t, err)
				pname := cleanupProfileName(prof.profileName)
				assert.EqualValues(t, pname, conf.AccessToken)
				assert.EqualValues(t, pname, conf.AccessTokenSecret)
				assert.EqualValues(t, pname, conf.Zone)
			} else {
				assert.Error(t, err)
			}
		}
	})

	t.Run("Not exists profile", func(t *testing.T) {
		// not exists profile
		_, err := LoadConfigFile("not-exists-profile-name")
		assert.Error(t, err)
	})

	t.Run("Invalid profile names", func(t *testing.T) {
		// - with filepath.Separator
		_, err := LoadConfigFile("test" + string([]rune{filepath.Separator}) + "test")
		assert.Error(t, err)
		// - with filepath.ListSeparator
		_, err = LoadConfigFile("test" + string([]rune{filepath.ListSeparator}) + "test")
		assert.Error(t, err)
	})

}

func initConfigFiles() func() {

	for _, prof := range testTargetProfiles() {
		p, _ := GetConfigFilePath(prof.profileName)
		if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
			panic(err)
		}
		if err := ioutil.WriteFile(p, []byte(prof.body), 0600); err != nil {
			panic(err)
		}
	}

	return func() {
		for _, prof := range testTargetProfiles() {
			p, _ := GetConfigFilePath(prof.profileName)
			os.RemoveAll(filepath.Dir(p))
		}
	}
}

var confTemplate = `
{
        "AccessToken": "%s",
        "AccessTokenSecret": "%s",
        "Zone": "%s"
}`

func TestGetCurrentName(t *testing.T) {
	homeDir, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	configDir := filepath.Join(homeDir, configDirName)
	profNameFile := filepath.Join(homeDir, configDirName, currentFileName)

	os.Mkdir(configDir, 0755)
	os.Remove(profNameFile)
	defer func() {
		os.Remove(profNameFile)
	}()

	t.Run("Should use default", func(t *testing.T) {
		n, err := GetCurrentName()
		assert.NoError(t, err)
		assert.Equal(t, "default", n)
	})

	t.Run("Should use profile file", func(t *testing.T) {
		// create profile name
		if err := ioutil.WriteFile(profNameFile, []byte("usacloud-unit-test1"), 0600); err != nil {
			panic(err)
		}
		n, err := GetCurrentName()
		assert.NoError(t, err)
		assert.Equal(t, "usacloud-unit-test1", n)
	})

	t.Run("Invalid name in profile file", func(t *testing.T) {
		// - with filepath.Separator
		if err := ioutil.WriteFile(profNameFile, []byte("test"+string([]rune{filepath.Separator})+"test"), 0600); err != nil {
			panic(err)
		}
		_, err := GetCurrentName()
		assert.Error(t, err)

		// - with filepath.ListSeparator
		if err := ioutil.WriteFile(profNameFile, []byte("test"+string([]rune{filepath.ListSeparator})+"test"), 0600); err != nil {
			panic(err)
		}
		_, err = GetCurrentName()
		assert.Error(t, err)
	})
}

func TestSetCurrentName(t *testing.T) {

	homeDir, err := homedir.Dir()
	if err != nil {
		panic(err)
	}

	configDir := filepath.Join(homeDir, configDirName)
	profNameFile := filepath.Join(homeDir, configDirName, currentFileName)

	os.Mkdir(configDir, 0755)
	os.Remove(profNameFile)
	defer func() {
		os.Remove(profNameFile)
	}()

	t.Run("Default profile", func(t *testing.T) {

		// profile dir isnot exists
		configFilePath, err := GetConfigFilePath("default")
		assert.NoError(t, err)
		profileDirExists := false
		if _, err := os.Stat(configFilePath); err == nil {
			profileDirExists = true
		}
		assert.False(t, profileDirExists)

		err = SetCurrentName("default")
		assert.NoError(t, err)

		data, err := ioutil.ReadFile(profNameFile)
		assert.NoError(t, err)
		assert.Equal(t, "default", string(data))
	})

	t.Run("Exists profile", func(t *testing.T) {

		defer initConfigFiles()()

		err = SetCurrentName("for-usacloud-unit-test1")
		assert.NoError(t, err)

		data, err := ioutil.ReadFile(profNameFile)
		assert.NoError(t, err)
		assert.Equal(t, "for-usacloud-unit-test1", string(data))
	})

	t.Run("Not exists profile", func(t *testing.T) {

		defer initConfigFiles()()

		err := SetCurrentName("for-usacloud-unit-test1")
		assert.NoError(t, err)

		err = SetCurrentName("not-exists")
		assert.Error(t, err)

		data, err := ioutil.ReadFile(profNameFile)
		assert.NoError(t, err)
		assert.Equal(t, "for-usacloud-unit-test1", string(data))
	})

	t.Run("Invalid name ", func(t *testing.T) {
		// - with filepath.Separator
		err := SetCurrentName("test" + string([]rune{filepath.Separator}) + "test")
		assert.Error(t, err)

		// - with filepath.ListSeparator
		err = SetCurrentName("test" + string([]rune{filepath.ListSeparator}) + "test")
		assert.Error(t, err)
	})
}

func TestSaveConfigFile(t *testing.T) {

	testProfileName := "for-usacloud-unit-test1"
	cleanupProfile(testProfileName)
	defer cleanupProfile(testProfileName)

	fileExists := func(path string) bool {
		_, err := os.Stat(path)
		return err == nil
	}

	t.Run("Valid profile", func(t *testing.T) {
		defer cleanupProfile(testProfileName)

		val := &ConfigFileValue{
			AccessToken:       "test-token",
			AccessTokenSecret: "test-secret",
			Zone:              "tk1v",
		}

		err := SaveConfigFile(testProfileName, val)
		assert.NoError(t, err)

		path, err := GetConfigFilePath(testProfileName)
		assert.NoError(t, err)

		assert.True(t, fileExists(filepath.Dir(path)))
		assert.True(t, fileExists(path))

	})
	t.Run("Empty profile", func(t *testing.T) {
		defer cleanupProfile(testProfileName)

		val := &ConfigFileValue{}

		err := SaveConfigFile(testProfileName, val)
		assert.Error(t, err)

		path, err := GetConfigFilePath(testProfileName)
		assert.NoError(t, err)

		assert.False(t, fileExists(path))

	})
	t.Run("Invalid profile name", func(t *testing.T) {
		defer cleanupProfile(testProfileName)

		val := &ConfigFileValue{
			AccessToken:       "test-token",
			AccessTokenSecret: "test-secret",
			Zone:              "tk1v",
		}
		// - with filepath.Separator
		profileName := "test" + string([]rune{filepath.Separator}) + "test"

		err := SaveConfigFile(profileName, val)
		assert.Error(t, err)

		path, err := GetConfigFilePath(testProfileName)
		assert.NoError(t, err)

		assert.False(t, fileExists(path))

		// - with filepath.ListSeparator
		profileName = "test" + string([]rune{filepath.ListSeparator}) + "test"
		err = SaveConfigFile(profileName, val)
		assert.Error(t, err)

		path, err = GetConfigFilePath(testProfileName)
		assert.NoError(t, err)

		assert.False(t, fileExists(path))

	})
}

func TestRemoveConfigFile(t *testing.T) {
	testProfileName := "for-usacloud-unit-test1"

	initFunc := func() func() {
		val := &ConfigFileValue{
			AccessToken:       "test-token",
			AccessTokenSecret: "test-secret",
			Zone:              "tk1v",
		}

		err := SaveConfigFile(testProfileName, val)
		if err != nil {
			panic(err)
		}

		err = SetCurrentName(testProfileName)
		if err != nil {
			panic(err)
		}

		return func() {
			cleanupProfile(testProfileName)
		}
	}
	fileExists := func(path string) bool {
		_, err := os.Stat(path)
		return err == nil
	}
	t.Run("Profile exists", func(t *testing.T) {
		defer initFunc()()
		err := RemoveConfigFile(testProfileName)

		assert.NoError(t, err)

		path, err := GetConfigFilePath(testProfileName)
		assert.NoError(t, err)

		assert.False(t, fileExists(filepath.Dir(path)))
		assert.False(t, fileExists(path))

		current, err := GetCurrentName()
		assert.NoError(t, err)
		assert.EqualValues(t, DefaultProfileName, current)
	})
	t.Run("Profile exists with other file", func(t *testing.T) {
		defer initFunc()()

		// create file in ProfileDir
		path, err := GetConfigFilePath(testProfileName)
		assert.NoError(t, err)

		testOtherFile := filepath.Join(filepath.Dir(path), "test")
		err = ioutil.WriteFile(testOtherFile, []byte{}, 0600)
		if err != nil {
			panic(err)
		}

		err = RemoveConfigFile(testProfileName)

		assert.NoError(t, err)

		assert.True(t, fileExists(filepath.Dir(path)))
		assert.True(t, fileExists(testOtherFile))
		assert.False(t, fileExists(path))
	})
	t.Run("Profile not exists", func(t *testing.T) {
		defer initFunc()()
		err := RemoveConfigFile("NotExistsProfileName")

		assert.Error(t, err)

		path, err := GetConfigFilePath(testProfileName)
		assert.NoError(t, err)

		assert.True(t, fileExists(filepath.Dir(path)))
		assert.True(t, fileExists(path))

		current, err := GetCurrentName()
		assert.NoError(t, err)
		assert.EqualValues(t, testProfileName, current)
	})
	t.Run("Invalid profile name", func(t *testing.T) {
		defer initFunc()()

		// - with filepath.Separator
		profileName := "test" + string([]rune{filepath.Separator}) + "test"

		err := RemoveConfigFile(profileName)
		assert.Error(t, err)

		path, err := GetConfigFilePath(testProfileName)
		assert.NoError(t, err)
		assert.True(t, fileExists(filepath.Dir(path)))
		assert.True(t, fileExists(path))

		current, err := GetCurrentName()
		assert.NoError(t, err)
		assert.EqualValues(t, testProfileName, current)

		// - with filepath.ListSeparator
		profileName = "test" + string([]rune{filepath.ListSeparator}) + "test"
		err = RemoveConfigFile(profileName)
		assert.Error(t, err)

		path, err = GetConfigFilePath(testProfileName)
		assert.NoError(t, err)
		assert.True(t, fileExists(filepath.Dir(path)))
		assert.True(t, fileExists(path))

		current, err = GetCurrentName()
		assert.NoError(t, err)
		assert.EqualValues(t, testProfileName, current)
	})
}

func cleanupProfile(profile string) {
	path, err := GetConfigFilePath(profile)
	if err != nil {
		panic(err)
	}
	os.RemoveAll(filepath.Dir(path))
}

// cleanupAllProfiles remove all entries under profile-base-dir(includes "current" file)
func cleanupAllProfiles() {
	dir, err := getProfileBaseDir()
	if err != nil {
		panic(err)
	}

	// dir is exists?
	configDirPath := filepath.Join(dir, configDirName)
	if _, err := os.Stat(configDirPath); err == nil {
		err := os.RemoveAll(configDirPath)
		if err != nil {
			panic(err)
		}
	}

}

func TestList(t *testing.T) {

	initFunc := func() func() {
		cleanupAllProfiles()
		return func() {
			cleanupAllProfiles()
		}
	}
	defer initFunc()()

	t.Run("Default only", func(t *testing.T) {
		defer initFunc()()

		profiles, err := List()
		assert.NoError(t, err)
		assert.Len(t, profiles, 1)
		assert.EqualValues(t, DefaultProfileName, profiles[0])

	})
	t.Run("Multiple profile", func(t *testing.T) {
		defer initFunc()()

		// create profile
		testProfileNames := []string{"test2", "test1"}
		val := &ConfigFileValue{
			AccessToken:       "test-token",
			AccessTokenSecret: "test-secret",
			Zone:              "tk1v",
		}
		for _, n := range testProfileNames {
			err := SaveConfigFile(n, val)
			if err != nil {
				panic(err)
			}
		}

		profiles, err := List()
		assert.NoError(t, err)
		assert.Len(t, profiles, 3) // default + test1 + test2
		assert.EqualValues(t, DefaultProfileName, profiles[0])
		assert.EqualValues(t, "test1", profiles[1]) // sorted by name(except "default")
		assert.EqualValues(t, "test2", profiles[2])

	})
	t.Run("With invalid profile", func(t *testing.T) {
		defer initFunc()()
		defer initConfigFiles()()

		profiles, err := List()
		assert.NoError(t, err)

		targets := testTargetProfiles()
		validCount := 0
		for _, p := range targets {
			if p.isValid {
				validCount++
			}
		}

		assert.Len(t, profiles, validCount)

	})
	t.Run("With empty dir", func(t *testing.T) {
		defer initFunc()()

		// create profile
		testProfileNames := []string{"test2", "test1"}
		val := &ConfigFileValue{
			AccessToken:       "test-token",
			AccessTokenSecret: "test-secret",
			Zone:              "tk1v",
		}
		for _, n := range testProfileNames {
			err := SaveConfigFile(n, val)
			if err != nil {
				panic(err)
			}
		}

		// create empty dir
		dir, _ := getProfileBaseDir()
		err := os.MkdirAll(filepath.Join(dir, configDirName, "test3"), 0755)
		if err != nil {
			panic(err)
		}

		profiles, err := List()
		assert.NoError(t, err)
		assert.Len(t, profiles, 3) // default + test1 + test2 ( without test3 )
		assert.EqualValues(t, DefaultProfileName, profiles[0])
		assert.EqualValues(t, "test1", profiles[1]) // sorted by name(except "default")
		assert.EqualValues(t, "test2", profiles[2])
	})
}
