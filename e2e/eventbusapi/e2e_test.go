// Copyright 2017-2025 The sacloud/usacloud Authors
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

package eventbusapi

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/sacloud/sakumock/eventbus"
	usacloudE2E "github.com/sacloud/usacloud/e2e"
	"github.com/stretchr/testify/require"
)

func TestE2E_EventBusAPI(t *testing.T) {
	server := eventbus.NewTestServer(eventbus.Config{})
	t.Cleanup(server.Close)
	t.Setenv("SAKURA_ACCESS_TOKEN", "dummy-token")
	t.Setenv("SAKURA_ACCESS_TOKEN_SECRET", "dummy-secret")
	// The trailing slash is required by EventBus clients which construct list
	// URLs with url.JoinPath.
	t.Setenv("SAKURA_ENDPOINTS_EVENTBUS", server.TestURL()+"/")

	const (
		processConfigurationName        = "usacloud-e2e-process-configuration"
		updatedProcessConfigurationName = "usacloud-e2e-process-configuration-updated"
		scheduleName                    = "usacloud-e2e-schedule"
		updatedScheduleName             = "usacloud-e2e-schedule-updated"
		triggerName                     = "usacloud-e2e-trigger"
		updatedTriggerName              = "usacloud-e2e-trigger-updated"
	)

	processConfigurationSettings := jsonArgument(t, map[string]any{
		"Destination": "simplemq",
		"Parameters":  `{"queue_name":"usacloud-e2e","content":"created"}`,
	})

	// Mutating commands must not proceed when confirmation is omitted.
	err := usacloudE2E.UsacloudRun(t,
		"eventbus-api", "process-configuration", "create",
		"--name", "usacloud-e2e-process-configuration-unconfirmed",
		"--settings", processConfigurationSettings,
	)
	require.Error(t, err)

	var processConfigurationID, scheduleID, triggerID string
	t.Cleanup(func() {
		// Child resources must be removed before their process configuration.
		deleteEventBusResource(t, "trigger", triggerID)
		deleteEventBusResource(t, "schedule", scheduleID)
		deleteEventBusResource(t, "process-configuration", processConfigurationID)
	})

	processConfiguration := onlyEventBusResult(t, runEventBusJSON(t,
		"eventbus-api", "process-configuration", "create",
		"--name", processConfigurationName,
		"--description", "created process configuration",
		"--tags", "eventbus,e2e",
		"--settings", processConfigurationSettings,
		"-y",
	))
	processConfigurationID = requiredString(t, processConfiguration, "ID")
	assertEventBusItem(t, processConfiguration, processConfigurationName, "created process configuration", []string{"eventbus", "e2e"}, "eventbusprocessconfiguration")
	assertSettings(t, processConfiguration, map[string]any{
		"Destination": "simplemq",
		"Parameters":  `{"queue_name":"usacloud-e2e","content":"created"}`,
	})

	processConfigurations := runEventBusJSON(t,
		"eventbus-api", "process-configuration", "list",
	)
	require.Len(t, processConfigurations, 1, "the unconfirmed create must not have mutated the mock")
	require.Equal(t, processConfigurationID, requiredString(t, processConfigurations[0], "ID"))

	readProcessConfiguration := onlyEventBusResult(t, runEventBusJSON(t,
		"eventbus-api", "process-configuration", "read", processConfigurationName,
	))
	require.Equal(t, processConfigurationID, requiredString(t, readProcessConfiguration, "ID"))
	assertEventBusItem(t, readProcessConfiguration, processConfigurationName, "created process configuration", []string{"eventbus", "e2e"}, "eventbusprocessconfiguration")

	updatedProcessConfigurationSettings := jsonArgument(t, map[string]any{
		"Destination": "simplemq",
		"Parameters":  `{"queue_name":"usacloud-e2e","content":"updated"}`,
	})
	updatedProcessConfiguration := onlyEventBusResult(t, runEventBusJSON(t,
		"eventbus-api", "process-configuration", "update", processConfigurationID,
		"--name", updatedProcessConfigurationName,
		"--description", "updated process configuration",
		"--tags", "eventbus,updated",
		"--settings", updatedProcessConfigurationSettings,
		"-y",
	))
	assertEventBusItem(t, updatedProcessConfiguration, updatedProcessConfigurationName, "updated process configuration", []string{"eventbus", "updated"}, "eventbusprocessconfiguration")
	assertSettings(t, updatedProcessConfiguration, map[string]any{
		"Destination": "simplemq",
		"Parameters":  `{"queue_name":"usacloud-e2e","content":"updated"}`,
	})

	secretFile := filepath.Join(t.TempDir(), "process-configuration-secret.json")
	secret := []byte(`{"AccessToken":"process-token","AccessTokenSecret":"process-secret"}`)
	require.NoError(t, os.WriteFile(secretFile, secret, 0o600))
	require.NoError(t, runEventBus(t,
		"eventbus-api", "process-configuration", "update-secret", processConfigurationID,
		"--secret-file", secretFile,
		"-y",
	))
	storedSecret, ok := server.Secret(processConfigurationID)
	require.True(t, ok, "update-secret must write the secret to EventBus")
	var storedSecretValue map[string]any
	require.NoError(t, json.Unmarshal(storedSecret, &storedSecretValue))
	require.Equal(t, "process-token", requiredString(t, storedSecretValue, "AccessToken"))
	require.Equal(t, "process-secret", requiredString(t, storedSecretValue, "AccessTokenSecret"))

	scheduleSettings := jsonArgument(t, map[string]any{
		"ProcessConfigurationID": processConfigurationID,
		"StartsAt":               1700000000000,
		"RecurringStep":          1,
		"RecurringUnit":          "day",
	})
	schedule := onlyEventBusResult(t, runEventBusJSON(t,
		"eventbus-api", "schedule", "create",
		"--name", scheduleName,
		"--description", "created schedule",
		"--tags", "eventbus,e2e",
		"--settings", scheduleSettings,
		"-y",
	))
	scheduleID = requiredString(t, schedule, "ID")
	assertEventBusItem(t, schedule, scheduleName, "created schedule", []string{"eventbus", "e2e"}, "eventbusschedule")
	assertSettings(t, schedule, map[string]any{
		"ProcessConfigurationID": processConfigurationID,
		"StartsAt":               "1700000000000",
		"RecurringStep":          float64(1),
		"RecurringUnit":          "day",
	})

	schedules := runEventBusJSON(t, "eventbus-api", "schedule", "list")
	require.Len(t, schedules, 1)
	require.Equal(t, scheduleID, requiredString(t, schedules[0], "ID"))
	readSchedule := onlyEventBusResult(t, runEventBusJSON(t,
		"eventbus-api", "schedule", "read", scheduleName,
	))
	require.Equal(t, scheduleID, requiredString(t, readSchedule, "ID"))
	assertEventBusItem(t, readSchedule, scheduleName, "created schedule", []string{"eventbus", "e2e"}, "eventbusschedule")

	updatedScheduleSettings := jsonArgument(t, map[string]any{
		"ProcessConfigurationID": processConfigurationID,
		"StartsAt":               1700003600000,
		"Crontab":                "*/15 * * * *",
	})
	updatedSchedule := onlyEventBusResult(t, runEventBusJSON(t,
		"eventbus-api", "schedule", "update", scheduleID,
		"--name", updatedScheduleName,
		"--description", "updated schedule",
		"--tags", "eventbus,updated",
		"--settings", updatedScheduleSettings,
		"-y",
	))
	assertEventBusItem(t, updatedSchedule, updatedScheduleName, "updated schedule", []string{"eventbus", "updated"}, "eventbusschedule")
	assertSettings(t, updatedSchedule, map[string]any{
		"ProcessConfigurationID": processConfigurationID,
		"StartsAt":               "1700003600000",
		"Crontab":                "*/15 * * * *",
	})

	triggerSettings := jsonArgument(t, map[string]any{
		"Source":                 "//usacloud-e2e/source",
		"Types":                  []string{"usacloud-e2e.created"},
		"ProcessConfigurationID": processConfigurationID,
	})
	trigger := onlyEventBusResult(t, runEventBusJSON(t,
		"eventbus-api", "trigger", "create",
		"--name", triggerName,
		"--description", "created trigger",
		"--tags", "eventbus,e2e",
		"--settings", triggerSettings,
		"-y",
	))
	triggerID = requiredString(t, trigger, "ID")
	assertEventBusItem(t, trigger, triggerName, "created trigger", []string{"eventbus", "e2e"}, "eventbustrigger")
	assertSettings(t, trigger, map[string]any{
		"Source":                 "//usacloud-e2e/source",
		"Types":                  []any{"usacloud-e2e.created"},
		"ProcessConfigurationID": processConfigurationID,
	})

	triggers := runEventBusJSON(t, "eventbus-api", "trigger", "list")
	require.Len(t, triggers, 1)
	require.Equal(t, triggerID, requiredString(t, triggers[0], "ID"))
	readTrigger := onlyEventBusResult(t, runEventBusJSON(t,
		"eventbus-api", "trigger", "read", triggerName,
	))
	require.Equal(t, triggerID, requiredString(t, readTrigger, "ID"))
	assertEventBusItem(t, readTrigger, triggerName, "created trigger", []string{"eventbus", "e2e"}, "eventbustrigger")

	updatedTriggerSettings := jsonArgument(t, map[string]any{
		"Source":                 "//usacloud-e2e/updated-source",
		"Types":                  []string{"usacloud-e2e.updated"},
		"ProcessConfigurationID": processConfigurationID,
	})
	updatedTrigger := onlyEventBusResult(t, runEventBusJSON(t,
		"eventbus-api", "trigger", "update", triggerID,
		"--name", updatedTriggerName,
		"--description", "updated trigger",
		"--tags", "eventbus,updated",
		"--settings", updatedTriggerSettings,
		"-y",
	))
	assertEventBusItem(t, updatedTrigger, updatedTriggerName, "updated trigger", []string{"eventbus", "updated"}, "eventbustrigger")
	assertSettings(t, updatedTrigger, map[string]any{
		"Source":                 "//usacloud-e2e/updated-source",
		"Types":                  []any{"usacloud-e2e.updated"},
		"ProcessConfigurationID": processConfigurationID,
	})

	require.NoError(t, runEventBus(t, "eventbus-api", "trigger", "delete", triggerID, "-y"))
	triggerID = ""
	require.Empty(t, runEventBusJSON(t, "eventbus-api", "trigger", "list"))
	require.NoError(t, runEventBus(t, "eventbus-api", "schedule", "delete", scheduleID, "-y"))
	scheduleID = ""
	require.Empty(t, runEventBusJSON(t, "eventbus-api", "schedule", "list"))
	require.NoError(t, runEventBus(t, "eventbus-api", "process-configuration", "delete", processConfigurationID, "-y"))
	processConfigurationID = ""
	require.Empty(t, runEventBusJSON(t, "eventbus-api", "process-configuration", "list"))
}

func runEventBus(t *testing.T, args ...string) error {
	t.Helper()
	args = append(args, "--output-type", "json")
	return usacloudE2E.UsacloudRun(t, args...)
}

func runEventBusJSON(t *testing.T, args ...string) []map[string]any {
	t.Helper()
	args = append(args, "--output-type", "json")
	output, err := usacloudE2E.UsacloudRunWithCombinedOutput(t, args...)
	require.NoError(t, err, "command output: %s", output)
	if len(bytes.TrimSpace(output)) == 0 {
		return nil
	}

	var values []map[string]any
	require.NoError(t, json.Unmarshal(output, &values), "command output must be JSON: %s", output)
	for i := range values {
		values[i] = normalizeSDKJSON(values[i]).(map[string]any)
	}
	return values
}

func normalizeSDKJSON(value any) any {
	switch value := value.(type) {
	case []any:
		result := make([]any, len(value))
		for i := range value {
			result[i] = normalizeSDKJSON(value[i])
		}
		return result
	case map[string]any:
		if set, ok := value["Set"].(bool); ok {
			if !set || value["Null"] == true {
				return nil
			}
			return normalizeSDKJSON(value["Value"])
		}
		if variant, ok := value["Type"].(string); ok {
			switch variant {
			case "ProcessConfigurationSettings", "ScheduleSettings", "TriggerSettings":
				return normalizeSDKJSON(value[variant])
			case "int64":
				return normalizeSDKJSON(value["Int64"])
			case "string":
				return normalizeSDKJSON(value["String"])
			}
		}
		result := make(map[string]any, len(value))
		for key, item := range value {
			result[key] = normalizeSDKJSON(item)
		}
		return result
	default:
		return value
	}
}

func deleteEventBusResource(t *testing.T, resource, id string) {
	t.Helper()
	if id == "" {
		return
	}
	if err := runEventBus(t, "eventbus-api", resource, "delete", id, "-y"); err != nil {
		t.Errorf("cleaning up EventBus %s %q: %v", resource, id, err)
	}
}

func onlyEventBusResult(t *testing.T, values []map[string]any) map[string]any {
	t.Helper()
	require.Len(t, values, 1)
	return values[0]
}

func jsonArgument(t *testing.T, value any) string {
	t.Helper()
	data, err := json.Marshal(value)
	require.NoError(t, err)
	return string(data)
}

func assertEventBusItem(t *testing.T, item map[string]any, name, description string, tags []string, providerClass string) {
	t.Helper()
	require.Equal(t, name, requiredString(t, item, "Name"))
	require.Equal(t, description, requiredString(t, item, "Description"))
	require.NotEmpty(t, requiredString(t, item, "ID"))

	rawTags, ok := item["Tags"].([]any)
	require.True(t, ok, "Tags must be an array: %#v", item["Tags"])
	actualTags := make([]string, len(rawTags))
	for i, tag := range rawTags {
		actualTags[i] = tag.(string)
	}
	require.Equal(t, tags, actualTags)

	provider, ok := item["Provider"].(map[string]any)
	require.True(t, ok, "Provider must be an object: %#v", item["Provider"])
	require.Equal(t, providerClass, requiredString(t, provider, "Class"))
}

func assertSettings(t *testing.T, item map[string]any, expected map[string]any) {
	t.Helper()
	settings, ok := item["Settings"].(map[string]any)
	require.True(t, ok, "Settings must be an object: %#v", item["Settings"])
	for key, value := range expected {
		require.Equal(t, value, settings[key], "unexpected Settings.%s", key)
	}
}

func requiredString(t *testing.T, value map[string]any, key string) string {
	t.Helper()
	result, ok := value[key].(string)
	require.True(t, ok, "%s must be a string: %#v", key, value[key])
	return result
}
