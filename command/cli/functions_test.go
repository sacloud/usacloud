package cli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/sacloud/usacloud/command/profile"
	"github.com/stretchr/testify/assert"
)

type mockFlagHandler struct {
	val     map[string]interface{}
	initVal map[string]interface{}
}

var envKeyValMap = map[string]string{
	"SAKURACLOUD_ACCESS_TOKEN":        "token",
	"SAKURACLOUD_ACCESS_TOKEN_SECRET": "secret",
	"SAKURACLOUD_ZONE":                "zone",
}

func newMockFlagHandler(val map[string]interface{}) *mockFlagHandler {
	h := &mockFlagHandler{
		val:     val,
		initVal: map[string]interface{}{},
	}
	for k, v := range val {
		h.initVal[k] = v
	}

	for key, valKey := range envKeyValMap {
		if v, ok := os.LookupEnv(key); ok {
			if s, ok := h.val[valKey]; !ok || s.(string) == "" {
				h.val[valKey] = v
			}
		}
	}
	return h
}

func (h *mockFlagHandler) IsSet(name string) bool {
	_, ok := h.initVal[name]
	return ok
}
func (h *mockFlagHandler) Set(name, value string) error {
	h.val[name] = value
	return nil
}

func (h *mockFlagHandler) String(name string) string {
	if v, ok := h.val[name]; ok {
		return v.(string)
	}
	return ""
}
func (h *mockFlagHandler) StringSlice(name string) []string {
	if v, ok := h.val[name]; ok {
		return v.([]string)
	}
	return []string{}
}

func TestApplyConfigFromFile(t *testing.T) {

	testProfileName := "testProfile"

	profiles := []string{testProfileName}
	initFunc := func() func() {
		clearAPIKeyEnvVars()
		clearProfiles(profiles...)

		return func() {
			clearAPIKeyEnvVars()
			clearProfiles(profiles...)
		}
	}

	defer initFunc()()

	t.Run("From cli-flag only", func(t *testing.T) {
		defer initFunc()()

		flagHandler := newMockFlagHandler(
			map[string]interface{}{
				"token":  "test-token",
				"secret": "test-secret",
				"zone":   "tk1v",
				//"profile": "",
			},
		)

		err := applyConfigFromFile(flagHandler)
		assert.NoError(t, err)

		assert.EqualValues(t, "test-token", flagHandler.String("token"))
		assert.EqualValues(t, "test-secret", flagHandler.String("secret"))
		assert.EqualValues(t, "tk1v", flagHandler.String("zone"))
	})

	t.Run("From config-file only", func(t *testing.T) {
		defer initFunc()()

		profile.SaveConfigFile(testProfileName, &profile.ConfigFileValue{
			AccessToken:       "test-token",
			AccessTokenSecret: "test-secret",
			Zone:              "tk1v",
		})

		flagHandler := newMockFlagHandler(
			map[string]interface{}{
				"profile": testProfileName,
			},
		)
		err := applyConfigFromFile(flagHandler)
		assert.NoError(t, err)

		assert.EqualValues(t, "test-token", flagHandler.String("token"))
		assert.EqualValues(t, "test-secret", flagHandler.String("secret"))
		assert.EqualValues(t, "tk1v", flagHandler.String("zone"))
	})

	t.Run("From env-var only", func(t *testing.T) {
		defer initFunc()()

		os.Setenv("SAKURACLOUD_ACCESS_TOKEN", "test-token")
		os.Setenv("SAKURACLOUD_ACCESS_TOKEN_SECRET", "test-secret")
		os.Setenv("SAKURACLOUD_ZONE", "tk1v")

		flagHandler := newMockFlagHandler(map[string]interface{}{})

		err := applyConfigFromFile(flagHandler)
		assert.NoError(t, err)

		assert.EqualValues(t, "test-token", flagHandler.String("token"))
		assert.EqualValues(t, "test-secret", flagHandler.String("secret"))
		assert.EqualValues(t, "tk1v", flagHandler.String("zone"))

	})

	t.Run("From cli-flag and config-file", func(t *testing.T) {
		defer initFunc()()

		profile.SaveConfigFile(testProfileName, &profile.ConfigFileValue{
			AccessToken:       "fromConfigFile",
			AccessTokenSecret: "fromConfigFile",
			Zone:              "fromConfigFile",
		})

		flagHandler := newMockFlagHandler(
			map[string]interface{}{
				"token":   "test-token",
				"secret":  "test-secret",
				"zone":    "tk1v",
				"profile": testProfileName,
			},
		)

		err := applyConfigFromFile(flagHandler)
		assert.NoError(t, err)

		assert.EqualValues(t, "test-token", flagHandler.String("token"))
		assert.EqualValues(t, "test-secret", flagHandler.String("secret"))
		assert.EqualValues(t, "tk1v", flagHandler.String("zone"))
	})

	t.Run("From cli-flag and env-var", func(t *testing.T) {
		defer initFunc()()

		os.Setenv("SAKURACLOUD_ACCESS_TOKEN", "fromEnv")
		os.Setenv("SAKURACLOUD_ACCESS_TOKEN_SECRET", "fromEnv")
		os.Setenv("SAKURACLOUD_ZONE", "fromEnv")

		flagHandler := newMockFlagHandler(
			map[string]interface{}{
				"token":  "test-token",
				"secret": "test-secret",
				"zone":   "tk1v",
			},
		)

		err := applyConfigFromFile(flagHandler)
		assert.NoError(t, err)

		assert.EqualValues(t, "test-token", flagHandler.String("token"))
		assert.EqualValues(t, "test-secret", flagHandler.String("secret"))
		assert.EqualValues(t, "tk1v", flagHandler.String("zone"))
	})
	t.Run("From config-file and env-var", func(t *testing.T) {
		defer initFunc()()

		profile.SaveConfigFile(testProfileName, &profile.ConfigFileValue{
			AccessToken:       "test-token",
			AccessTokenSecret: "test-secret",
			Zone:              "tk1v",
		})

		os.Setenv("SAKURACLOUD_ACCESS_TOKEN", "fromEnv")
		os.Setenv("SAKURACLOUD_ACCESS_TOKEN_SECRET", "fromEnv")
		os.Setenv("SAKURACLOUD_ZONE", "fromEnv")

		flagHandler := newMockFlagHandler(
			map[string]interface{}{
				"profile": testProfileName,
			},
		)

		err := applyConfigFromFile(flagHandler)
		assert.NoError(t, err)

		assert.EqualValues(t, "test-token", flagHandler.String("token"))
		assert.EqualValues(t, "test-secret", flagHandler.String("secret"))
		assert.EqualValues(t, "tk1v", flagHandler.String("zone"))
	})
	t.Run("All exists", func(t *testing.T) {
		defer initFunc()()

		os.Setenv("SAKURACLOUD_ACCESS_TOKEN", "fromEnv")
		os.Setenv("SAKURACLOUD_ACCESS_TOKEN_SECRET", "fromEnv")
		os.Setenv("SAKURACLOUD_ZONE", "fromEnv")

		profile.SaveConfigFile(testProfileName, &profile.ConfigFileValue{
			AccessToken:       "fromConfigFile",
			AccessTokenSecret: "fromConfigFile",
			Zone:              "fromConfigFile",
		})

		flagHandler := newMockFlagHandler(
			map[string]interface{}{
				"token":   "test-token",
				"secret":  "test-secret",
				"zone":    "tk1v",
				"profile": testProfileName,
			},
		)

		err := applyConfigFromFile(flagHandler)
		assert.NoError(t, err)

		assert.EqualValues(t, "test-token", flagHandler.String("token"))
		assert.EqualValues(t, "test-secret", flagHandler.String("secret"))
		assert.EqualValues(t, "tk1v", flagHandler.String("zone"))
	})
}

func clearCurrentProfile() {
	profile.SetCurrentName(profile.DefaultProfileName)
}

func clearAPIKeyEnvVars() {
	targets := []string{
		"SAKURACLOUD_ACCESS_TOKEN",
		"SAKURACLOUD_ACCESS_TOKEN_SECRET",
		"SAKURACLOUD_ZONE",
		"USACLOUD_PROFILE",
	}

	for _, key := range targets {
		if _, ok := os.LookupEnv(key); ok {
			err := os.Unsetenv(key)
			if err != nil {
				panic(err)
			}
		}
	}
}

func clearProfiles(targets ...string) {
	os.Unsetenv(profile.ProfileDirEnv)
	for _, profileName := range targets {
		path, err := profile.GetConfigFilePath(profileName)
		if err != nil {
			panic(err)
		}
		os.RemoveAll(filepath.Dir(path))
	}
}
