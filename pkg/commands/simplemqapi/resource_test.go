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

package simplemqapi

import (
	"reflect"
	"sort"
	"testing"

	"github.com/sacloud/sacloud-sdk-go/api/simplemq"
	sdkqueue "github.com/sacloud/sacloud-sdk-go/api/simplemq/apis/v1/queue"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/naming"
	"github.com/stretchr/testify/require"
)

func TestResource(t *testing.T) {
	require.True(t, Resource.IsGlobalResource)
	require.Empty(t, Resource.PlatformName)
	require.Equal(t, "api", Resource.Category.Key)

	children := Resource.Children()
	require.Len(t, children, 2)
	require.Equal(t, []string{"queue", "message"}, []string{children[0].Name, children[1].Name})
	for _, child := range children {
		require.True(t, child.IsGlobalResource)
		require.Empty(t, child.PlatformName)
	}
}

func TestResourceHelpExplainsSimpleMQModel(t *testing.T) {
	command := Resource.CLICommand()

	require.Contains(t, command.Short, "[Experimental]")
	require.Contains(t, command.Long, "experimental support")
	require.Contains(t, command.Long, "may change without notice")
	require.Contains(t, command.Long, "asynchronous")
	require.Contains(t, command.Long, "at-least-once")
	require.Contains(t, command.Long, "Create a queue before")
	require.Contains(t, command.Long, "global")
	require.Contains(t, command.Example, "queue create --help")
	require.Contains(t, command.Example, "message send --help")

	queue, _, err := command.Find([]string{"queue"})
	require.NoError(t, err)
	require.Contains(t, queue.Long, "experimental support")

	create, _, err := command.Find([]string{"queue", "create"})
	require.NoError(t, err)
	require.Contains(t, create.Long, "may change without notice")

	expectedUsage := map[string]string{
		"queue":   "Manage queues that store SimpleMQ messages",
		"message": "Send, receive, extend, and delete queue messages",
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
		"queue":   reflect.TypeOf((*simplemq.QueueAPI)(nil)).Elem(),
		"message": reflect.TypeOf((*simplemq.MessageAPI)(nil)).Elem(),
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
	stringID := sdkqueue.NewStringCommonServiceItemID("queue-id")
	item := &sdkqueue.CommonServiceItem{
		ID:     stringID,
		Status: sdkqueue.Status{QueueName: "queue-name"},
		Tags:   []string{"tag-a", "tag-b"},
	}
	require.Equal(t, &core.Labels{
		Id:   "queue-id",
		Name: "queue-name",
		Tags: []string{"tag-a", "tag-b"},
	}, extractLabels(item))
	require.Nil(t, extractLabels(sdkqueue.CommonServiceItem{}))
}
