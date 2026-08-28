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

package schedule

import (
	"context"
	"testing"

	"github.com/sacloud/sacloud-sdk-go/api/eventbus"
	v1 "github.com/sacloud/sacloud-sdk-go/api/eventbus/apis/v1"
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/stretchr/testify/require"
)

type scheduleAPIFake struct {
	items         []v1.CommonServiceItem
	readID        string
	createRequest v1.CreateCommonServiceItemRequest
	updateID      string
	updateRequest v1.UpdateCommonServiceItemRequest
	deleteID      string
}

var _ eventbus.ScheduleAPI = (*scheduleAPIFake)(nil)

func (f *scheduleAPIFake) List(context.Context) ([]v1.CommonServiceItem, error) {
	return f.items, nil
}

func (f *scheduleAPIFake) Read(_ context.Context, id string) (*v1.CommonServiceItem, error) {
	f.readID = id
	return &v1.CommonServiceItem{ID: id}, nil
}

func (f *scheduleAPIFake) Create(_ context.Context, request v1.CreateCommonServiceItemRequest) (*v1.CommonServiceItem, error) {
	f.createRequest = request
	return &v1.CommonServiceItem{ID: "created"}, nil
}

func (f *scheduleAPIFake) Update(_ context.Context, id string, request v1.UpdateCommonServiceItemRequest) (*v1.CommonServiceItem, error) {
	f.updateID = id
	f.updateRequest = request
	return &v1.CommonServiceItem{ID: id}, nil
}

func (f *scheduleAPIFake) Delete(_ context.Context, id string) error {
	f.deleteID = id
	return nil
}

func useScheduleAPIFake(t *testing.T, op eventbus.ScheduleAPI) {
	t.Helper()
	oldFactory := scheduleOpFactory
	scheduleOpFactory = func(cli.Context) (eventbus.ScheduleAPI, error) {
		return op, nil
	}
	t.Cleanup(func() {
		scheduleOpFactory = oldFactory
	})
}

func TestScheduleCommandsInvokeSDKInterface(t *testing.T) {
	fake := &scheduleAPIFake{items: []v1.CommonServiceItem{{ID: "listed"}}}
	useScheduleAPIFake(t, fake)

	results, err := listCommand.Func(nil, newListParameter())
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.IsType(t, &v1.CommonServiceItem{}, results[0])

	create := &createParameter{
		NameParameter: cflag.NameParameter{Name: "new"},
		Settings:      `{"ProcessConfigurationID":"pc-id","RecurringStep":5,"RecurringUnit":"min","StartsAt":1}`,
	}
	_, err = createCommand.Func(nil, create)
	require.NoError(t, err)
	require.Equal(t, v1.ProviderClassEventbusschedule, fake.createRequest.CommonServiceItem.Provider.Class)
	require.True(t, fake.createRequest.CommonServiceItem.Settings.IsScheduleSettings())

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

func TestScheduleRequestBuildsTypedPatch(t *testing.T) {
	tags := []string{"tag-a"}
	settings := `{"ProcessConfigurationID":"pc-id","RecurringStep":5,"RecurringUnit":"min","StartsAt":1}`
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
	require.True(t, item.Settings.Value.IsScheduleSettings())
	require.Equal(t, v1.ProviderClassEventbusschedule, item.Provider.Value.Class)
}

func TestScheduleSettingsRejectsResponseTimestampFormat(t *testing.T) {
	_, err := scheduleSettings(`{"ProcessConfigurationID":"pc-id","StartsAt":"2026-08-28T00:00:00Z"}`)
	require.EqualError(t, err, "Schedule settings StartsAt must be an epoch timestamp in milliseconds")
}

func TestScheduleCreateHelpDocumentsSettings(t *testing.T) {
	command := createCommand.CLICommand()

	require.Contains(t, command.Long, "ProcessConfigurationID")
	require.Contains(t, command.Long, "Unix epoch milliseconds")
	require.Contains(t, command.Long, "min, hour, or day")
	require.Contains(t, command.Example, `"RecurringUnit":"day"`)
	require.Contains(t, command.Example, "./schedule-settings.json")
}

func TestScheduleUpdateRejectsDescriptionAndClear(t *testing.T) {
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
