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

package trigger

import (
	"bytes"
	"context"
	"testing"

	"github.com/sacloud/sacloud-sdk-go/api/eventbus"
	v1 "github.com/sacloud/sacloud-sdk-go/api/eventbus/apis/v1"
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/output"
	"github.com/stretchr/testify/require"
)

type triggerAPIFake struct {
	items         []v1.CommonServiceItem
	readID        string
	createRequest v1.CreateCommonServiceItemRequest
	updateID      string
	updateRequest v1.UpdateCommonServiceItemRequest
	deleteID      string
}

var _ eventbus.TriggerAPI = (*triggerAPIFake)(nil)

func (f *triggerAPIFake) List(context.Context) ([]v1.CommonServiceItem, error) {
	return f.items, nil
}

func (f *triggerAPIFake) Read(_ context.Context, id string) (*v1.CommonServiceItem, error) {
	f.readID = id
	return &v1.CommonServiceItem{ID: id}, nil
}

func (f *triggerAPIFake) Create(_ context.Context, request v1.CreateCommonServiceItemRequest) (*v1.CommonServiceItem, error) {
	f.createRequest = request
	return &v1.CommonServiceItem{ID: "created"}, nil
}

func (f *triggerAPIFake) Update(_ context.Context, id string, request v1.UpdateCommonServiceItemRequest) (*v1.CommonServiceItem, error) {
	f.updateID = id
	f.updateRequest = request
	return &v1.CommonServiceItem{ID: id}, nil
}

func (f *triggerAPIFake) Delete(_ context.Context, id string) error {
	f.deleteID = id
	return nil
}

func useTriggerAPIFake(t *testing.T, op eventbus.TriggerAPI) {
	t.Helper()
	oldFactory := triggerOpFactory
	triggerOpFactory = func(cli.Context) (eventbus.TriggerAPI, error) {
		return op, nil
	}
	t.Cleanup(func() {
		triggerOpFactory = oldFactory
	})
}

func TestTriggerCommandsInvokeSDKInterface(t *testing.T) {
	fake := &triggerAPIFake{items: []v1.CommonServiceItem{{ID: "listed"}}}
	useTriggerAPIFake(t, fake)

	results, err := listCommand.Func(nil, newListParameter())
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.IsType(t, &v1.CommonServiceItem{}, results[0])

	create := &createParameter{
		NameParameter: cflag.NameParameter{Name: "new"},
		Settings:      `{"Source":"source","Types":["type-a"],"Conditions":[],"ProcessConfigurationID":"pc-id"}`,
	}
	results, err = createCommand.Func(nil, create)
	require.NoError(t, err)
	require.Equal(t, v1.ProviderClassEventbustrigger, fake.createRequest.CommonServiceItem.Provider.Class)
	require.True(t, fake.createRequest.CommonServiceItem.Settings.IsTriggerSettings())
	requireQuietOutput(t, results, "created\n")

	_, err = readCommand.Func(nil, &readParameter{IDParameter: cflag.IDParameter{ID: "read-id"}})
	require.NoError(t, err)
	require.Equal(t, "read-id", fake.readID)

	name := "updated"
	_, err = updateCommand.Func(nil, &updateParameter{
		IDParameter:         cflag.IDParameter{ID: "update-id"},
		NameUpdateParameter: cflag.NameUpdateParameter{Name: &name},
	})
	require.NoError(t, err)
	require.Equal(t, "update-id", fake.updateID)
	require.True(t, fake.updateRequest.CommonServiceItem.Name.IsSet())

	_, err = deleteCommand.Func(nil, &deleteParameter{IDParameter: cflag.IDParameter{ID: "delete-id"}})
	require.NoError(t, err)
	require.Equal(t, "delete-id", fake.deleteID)
}

func TestTriggerRequestBuildsTypedPatch(t *testing.T) {
	tags := []string{"tag-a"}
	settings := `{"Source":"source","Types":["type-a"],"Conditions":[],"ProcessConfigurationID":"pc-id"}`
	request, err := updateRequest(&updateParameter{
		NameUpdateParameter: cflag.NameUpdateParameter{Name: stringPtr("new-name")},
		TagsUpdateParameter: cflag.TagsUpdateParameter{Tags: &tags},
		ClearDescription:    true,
		Settings:            &settings,
	})
	require.NoError(t, err)

	item := request.CommonServiceItem
	require.True(t, item.Name.IsSet())
	require.True(t, item.Description.IsNull())
	require.Equal(t, tags, item.Tags)
	require.True(t, item.Settings.IsSet())
	require.True(t, item.Settings.Value.IsTriggerSettings())
	require.Equal(t, v1.ProviderClassEventbustrigger, item.Provider.Value.Class)
}

func TestTriggerUpdateRejectsDescriptionAndClear(t *testing.T) {
	description := "description"
	err := validateUpdateParameter(nil, &updateParameter{
		DescUpdateParameter: cflag.DescUpdateParameter{Description: &description},
		ClearDescription:    true,
	})
	require.EqualError(t, err, "--description and --clear-description cannot be used together")
}

func stringPtr(value string) *string {
	return &value
}

func requireQuietOutput(t *testing.T, results []interface{}, expected string) {
	t.Helper()

	contents := make(output.Contents, 0, len(results))
	for _, result := range results {
		contents = append(contents, &output.Content{Value: result})
	}
	var out bytes.Buffer
	require.NoError(t, output.NewIDOutput(&out, nil).Print(contents))
	require.Equal(t, expected, out.String())
}
