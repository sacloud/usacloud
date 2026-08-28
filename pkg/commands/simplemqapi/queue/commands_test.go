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

package queue

import (
	"context"
	"testing"

	"github.com/sacloud/sacloud-sdk-go/api/simplemq"
	sdkqueue "github.com/sacloud/sacloud-sdk-go/api/simplemq/apis/v1/queue"
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/output"
	"github.com/sacloud/usacloud/pkg/validate"
	"github.com/stretchr/testify/require"
)

func init() {
	validate.InitializeValidator(nil)
}

type queueAPIFake struct {
	items               []sdkqueue.CommonServiceItem
	readID              string
	createRequest       sdkqueue.CreateQueueRequest
	configID            string
	configRequest       sdkqueue.ConfigQueueRequest
	deleteID            string
	countMessagesID     string
	countMessagesResult int
	rotateAPIKeyID      string
	rotateAPIKeyResult  string
	clearMessagesID     string
}

var _ simplemq.QueueAPI = (*queueAPIFake)(nil)

func (f *queueAPIFake) List(context.Context) ([]sdkqueue.CommonServiceItem, error) {
	return f.items, nil
}

func (f *queueAPIFake) Read(_ context.Context, id string) (*sdkqueue.CommonServiceItem, error) {
	f.readID = id
	return &sdkqueue.CommonServiceItem{Status: sdkqueue.Status{QueueName: "read-queue"}}, nil
}

func (f *queueAPIFake) Create(_ context.Context, request sdkqueue.CreateQueueRequest) (*sdkqueue.CommonServiceItem, error) {
	f.createRequest = request
	return &sdkqueue.CommonServiceItem{Status: sdkqueue.Status{QueueName: string(request.CommonServiceItem.Name)}}, nil
}

func (f *queueAPIFake) Config(_ context.Context, id string, request sdkqueue.ConfigQueueRequest) (*sdkqueue.CommonServiceItem, error) {
	f.configID = id
	f.configRequest = request
	return &sdkqueue.CommonServiceItem{Status: sdkqueue.Status{QueueName: "configured-queue"}}, nil
}

func (f *queueAPIFake) Delete(_ context.Context, id string) error {
	f.deleteID = id
	return nil
}

func (f *queueAPIFake) CountMessages(_ context.Context, id string) (int, error) {
	f.countMessagesID = id
	return f.countMessagesResult, nil
}

func (f *queueAPIFake) RotateAPIKey(_ context.Context, id string) (string, error) {
	f.rotateAPIKeyID = id
	return f.rotateAPIKeyResult, nil
}

func (f *queueAPIFake) ClearMessages(_ context.Context, id string) error {
	f.clearMessagesID = id
	return nil
}

func useQueueAPIFake(t *testing.T, op simplemq.QueueAPI) {
	t.Helper()
	oldFactory := queueOpFactory
	queueOpFactory = func(cli.Context) (simplemq.QueueAPI, error) {
		return op, nil
	}
	t.Cleanup(func() {
		queueOpFactory = oldFactory
	})
}

func TestQueueCommandsInvokeSDKInterface(t *testing.T) {
	fake := &queueAPIFake{
		items:               []sdkqueue.CommonServiceItem{{Status: sdkqueue.Status{QueueName: "listed-queue"}}},
		countMessagesResult: 7,
		rotateAPIKeyResult:  "dummy-value",
	}
	useQueueAPIFake(t, fake)

	results, err := listCommand.Func(nil, newListParameter())
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.IsType(t, &sdkqueue.CommonServiceItem{}, results[0])

	results, err = createCommand.Func(nil, &createParameter{
		Name:        "created-queue",
		Description: "description",
		Tags:        []string{"tag-a"},
	})
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.Equal(t, sdkqueue.QueueName("created-queue"), fake.createRequest.CommonServiceItem.Name)
	require.Equal(t, "description", fake.createRequest.CommonServiceItem.Description.Value)
	require.Equal(t, []string{"tag-a"}, fake.createRequest.CommonServiceItem.Tags)
	require.False(t, fake.createRequest.CommonServiceItem.Icon.Set)

	_, err = readCommand.Func(nil, &readParameter{IDParameter: cflag.IDParameter{ID: "read-id"}})
	require.NoError(t, err)
	require.Equal(t, "read-id", fake.readID)

	_, err = configCommand.Func(nil, &configParameter{
		IDParameter:              cflag.IDParameter{ID: "config-id"},
		VisibilityTimeoutSeconds: 30,
		ExpireSeconds:            3600,
		Description:              "new description",
		Tags:                     []string{"tag-b"},
	})
	require.NoError(t, err)
	require.Equal(t, "config-id", fake.configID)
	require.Equal(t, sdkqueue.VisibilityTimeoutSeconds(30), fake.configRequest.CommonServiceItem.Settings.VisibilityTimeoutSeconds)
	require.Equal(t, sdkqueue.ExpireSeconds(3600), fake.configRequest.CommonServiceItem.Settings.ExpireSeconds)
	require.Equal(t, "new description", fake.configRequest.CommonServiceItem.Description.Value)
	require.Equal(t, []string{"tag-b"}, fake.configRequest.CommonServiceItem.Tags)
	require.False(t, fake.configRequest.CommonServiceItem.Icon.Set)

	_, err = deleteCommand.Func(nil, &deleteParameter{IDParameter: cflag.IDParameter{ID: "delete-id"}})
	require.NoError(t, err)
	require.Equal(t, "delete-id", fake.deleteID)

	results, err = countMessagesCommand.Func(nil, &countMessagesParameter{IDParameter: cflag.IDParameter{ID: "count-id"}})
	require.NoError(t, err)
	require.Equal(t, "count-id", fake.countMessagesID)
	require.Equal(t, &countMessagesResult{MessageCount: 7}, results[0])

	results, err = rotateAPIKeyCommand.Func(nil, &rotateAPIKeyParameter{IDParameter: cflag.IDParameter{ID: "rotate-id"}})
	require.NoError(t, err)
	require.Equal(t, "rotate-id", fake.rotateAPIKeyID)
	require.Equal(t, &rotateAPIKeyResult{APIKey: "dummy-value"}, results[0])

	_, err = clearMessagesCommand.Func(nil, &clearMessagesParameter{IDParameter: cflag.IDParameter{ID: "clear-id"}})
	require.NoError(t, err)
	require.Equal(t, "clear-id", fake.clearMessagesID)
}

func TestQueueCommandMetadata(t *testing.T) {
	require.True(t, listCommand.NoProgress)
	require.NotEmpty(t, listCommand.ColumnDefs)

	for _, command := range []*core.Command{readCommand, countMessagesCommand, rotateAPIKeyCommand} {
		require.Equal(t, core.SelectorType(core.SelectorTypeRequireSingle), command.SelectorType)
		require.NotNil(t, command.ListAllFunc)
	}
	for _, command := range []*core.Command{configCommand, deleteCommand, clearMessagesCommand} {
		require.Equal(t, core.SelectorType(core.SelectorTypeRequireMulti), command.SelectorType)
		require.NotNil(t, command.ListAllFunc)
		require.Implements(t, (*cflag.ConfirmParameterValueHandler)(nil), command.ParameterInitializer())
	}
	require.Implements(t, (*cflag.ConfirmParameterValueHandler)(nil), createCommand.ParameterInitializer())
	require.Implements(t, (*cflag.ConfirmParameterValueHandler)(nil), rotateAPIKeyCommand.ParameterInitializer())
	require.Contains(t, createCommand.LongUsage, "does not return a Message API key")
	require.Contains(t, createCommand.Example, "rotate-api-key")
	require.Equal(t, []string{"MessageCount"}, columnNames(countMessagesCommand.ColumnDefs))
	require.Equal(t, []string{"APIKey"}, columnNames(rotateAPIKeyCommand.ColumnDefs))
	require.Contains(t, rotateAPIKeyCommand.LongUsage, "invalidates")
}

func TestQueueValidation(t *testing.T) {
	for _, name := range []string{"valid-queue", "ABCDE", "a234567890123456789012345678901234567890123456789012345678901234"} {
		require.NoError(t, validateQueueNameParameter(nil, &createParameter{Name: name}))
	}
	for _, name := range []string{"abcd", "-invalid", "invalid-", "invalid--queue", "invalid_queue"} {
		require.Error(t, validateQueueNameParameter(nil, &createParameter{Name: name}))
	}

	valid := &configParameter{VisibilityTimeoutSeconds: 5, ExpireSeconds: 60}
	require.NoError(t, validate.Exec(valid))
	valid.VisibilityTimeoutSeconds = 900
	valid.ExpireSeconds = 1209600
	require.NoError(t, validate.Exec(valid))

	for _, parameter := range []*configParameter{
		{VisibilityTimeoutSeconds: 4, ExpireSeconds: 60},
		{VisibilityTimeoutSeconds: 901, ExpireSeconds: 60},
		{VisibilityTimeoutSeconds: 5, ExpireSeconds: 59},
		{VisibilityTimeoutSeconds: 5, ExpireSeconds: 1209601},
	} {
		require.Error(t, validate.Exec(parameter))
	}
}

func columnNames(columns []output.ColumnDef) []string {
	names := make([]string, 0, len(columns))
	for _, column := range columns {
		names = append(names, column.Name)
	}
	return names
}
