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

package eventbusapi

import (
	"reflect"
	"sort"
	"testing"

	"github.com/sacloud/sacloud-sdk-go/api/eventbus"
	v1 "github.com/sacloud/sacloud-sdk-go/api/eventbus/apis/v1"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/naming"
	"github.com/stretchr/testify/require"
)

func TestResource(t *testing.T) {
	require.True(t, Resource.IsGlobalResource)
	require.Empty(t, Resource.PlatformName)
	require.Empty(t, Resource.DefaultCommandName)

	children := Resource.Children()
	require.Len(t, children, 3)
	require.Equal(t, []string{"schedule", "trigger", "process-configuration"}, []string{
		children[0].Name,
		children[1].Name,
		children[2].Name,
	})
	for _, child := range children {
		require.True(t, child.IsGlobalResource)
		require.Empty(t, child.PlatformName)
		require.Empty(t, child.DefaultCommandName)
	}
}

func TestResourceHelpExplainsEventBusModel(t *testing.T) {
	command := Resource.CLICommand()

	require.Contains(t, command.Short, "[Experimental]")
	require.Contains(t, command.Long, "experimental support")
	require.Contains(t, command.Long, "may change without notice")
	require.Contains(t, command.Long, "Create a process configuration before")
	require.Contains(t, command.Long, "global")
	require.Contains(t, command.Long, "best effort")

	schedule, _, err := command.Find([]string{"schedule"})
	require.NoError(t, err)
	require.Contains(t, schedule.Long, "experimental support")

	create, _, err := command.Find([]string{"schedule", "create"})
	require.NoError(t, err)
	require.Contains(t, create.Long, "may change without notice")

	expectedUsage := map[string]string{
		"process-configuration": "Define the job that EventBus runs",
		"schedule":              "Run a process configuration at specified times",
		"trigger":               "Run a process configuration when an event is detected",
	}
	for resourceName, usage := range expectedUsage {
		child, _, err := command.Find([]string{resourceName})
		require.NoError(t, err)
		require.Equal(t, usage, child.Short)
		require.NotEqual(t, child.Short, child.Long)
		require.NotEmpty(t, child.Example)
	}
}

func TestSDKInterfaceCoverage(t *testing.T) {
	apiTypes := map[string]reflect.Type{
		"schedule":              reflect.TypeOf((*eventbus.ScheduleAPI)(nil)).Elem(),
		"trigger":               reflect.TypeOf((*eventbus.TriggerAPI)(nil)).Elem(),
		"process-configuration": reflect.TypeOf((*eventbus.ProcessConfigurationAPI)(nil)).Elem(),
	}
	resources := make(map[string]*core.Resource)
	for _, resource := range Resource.Children() {
		resources[resource.Name] = resource
	}

	for resourceName, apiType := range apiTypes {
		t.Run(resourceName, func(t *testing.T) {
			resource := resources[resourceName]
			require.NotNil(t, resource)

			var sdkCommands []string
			for i := 0; i < apiType.NumMethod(); i++ {
				sdkCommands = append(sdkCommands, naming.ToKebabCase(apiType.Method(i).Name))
			}
			var cliCommands []string
			for _, command := range resource.Commands() {
				require.NotNil(t, command.Func, "%s must use an explicit Func", command.Name)
				cliCommands = append(cliCommands, command.Name)
			}
			sort.Strings(sdkCommands)
			sort.Strings(cliCommands)
			require.Equal(t, sdkCommands, cliCommands)
		})
	}
}

func TestExtractLabels(t *testing.T) {
	item := &v1.CommonServiceItem{ID: "item-id", Name: "item-name", Tags: []string{"tag-a", "tag-b"}}
	require.Equal(t, "item-id", extractLabels(item).Id)
	require.Equal(t, "item-name", extractLabels(item).Name)
	require.Equal(t, []string{"tag-a", "tag-b"}, extractLabels(item).Tags)
	require.Nil(t, extractLabels(v1.CommonServiceItem{}))
}
