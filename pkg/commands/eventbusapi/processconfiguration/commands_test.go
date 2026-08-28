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

package processconfiguration

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"
	"time"

	"github.com/sacloud/saclient-go"
	"github.com/sacloud/sacloud-sdk-go/api/eventbus"
	v1 "github.com/sacloud/sacloud-sdk-go/api/eventbus/apis/v1"
	sdksaclient "github.com/sacloud/sacloud-sdk-go/common/saclient"
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/config"
	"github.com/sacloud/usacloud/pkg/output"
	"github.com/stretchr/testify/require"
)

type processConfigurationAPIFake struct {
	items         []v1.CommonServiceItem
	readID        string
	createRequest v1.CreateCommonServiceItemRequest
	updateID      string
	updateRequest v1.UpdateCommonServiceItemRequest
	secretID      string
	secret        v1.SetSecretRequest
	deleteID      string
}

var _ eventbus.ProcessConfigurationAPI = (*processConfigurationAPIFake)(nil)

func (f *processConfigurationAPIFake) List(context.Context) ([]v1.CommonServiceItem, error) {
	return f.items, nil
}

func (f *processConfigurationAPIFake) Read(_ context.Context, id string) (*v1.CommonServiceItem, error) {
	f.readID = id
	return &v1.CommonServiceItem{ID: id}, nil
}

func (f *processConfigurationAPIFake) Create(_ context.Context, request v1.CreateCommonServiceItemRequest) (*v1.CommonServiceItem, error) {
	f.createRequest = request
	return &v1.CommonServiceItem{ID: "created"}, nil
}

func (f *processConfigurationAPIFake) Update(_ context.Context, id string, request v1.UpdateCommonServiceItemRequest) (*v1.CommonServiceItem, error) {
	f.updateID = id
	f.updateRequest = request
	return &v1.CommonServiceItem{ID: id}, nil
}

func (f *processConfigurationAPIFake) UpdateSecret(_ context.Context, id string, secret v1.SetSecretRequest) error {
	f.secretID = id
	f.secret = secret
	return nil
}

func (f *processConfigurationAPIFake) Delete(_ context.Context, id string) error {
	f.deleteID = id
	return nil
}

func useProcessConfigurationAPIFake(t *testing.T, op eventbus.ProcessConfigurationAPI) {
	t.Helper()
	oldFactory := processConfigurationOpFactory
	processConfigurationOpFactory = func(cli.Context) (eventbus.ProcessConfigurationAPI, error) {
		return op, nil
	}
	t.Cleanup(func() {
		processConfigurationOpFactory = oldFactory
	})
}

func TestProcessConfigurationCommandsInvokeSDKInterface(t *testing.T) {
	fake := &processConfigurationAPIFake{items: []v1.CommonServiceItem{{ID: "listed"}}}
	useProcessConfigurationAPIFake(t, fake)

	results, err := listCommand.Func(nil, newListParameter())
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.IsType(t, &v1.CommonServiceItem{}, results[0])

	create := &createParameter{
		NameParameter: cflag.NameParameter{Name: "new"},
		Settings:      `{"Destination":"simplemq","Parameters":"{}"}`,
	}
	_, err = createCommand.Func(nil, create)
	require.NoError(t, err)
	require.Equal(t, v1.ProviderClassEventbusprocessconfiguration, fake.createRequest.CommonServiceItem.Provider.Class)
	require.True(t, fake.createRequest.CommonServiceItem.Settings.IsProcessConfigurationSettings())

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

	secretInput := standardInput(t, `{"APIKey":"not-a-real-secret"}`)
	_, err = updateSecretCommand.Func(&stdinContext{input: secretInput}, &updateSecretParameter{
		IDParameter: cflag.IDParameter{ID: "secret-id"},
		SecretFile:  "-",
	})
	require.NoError(t, err)
	require.Equal(t, "secret-id", fake.secretID)
	require.True(t, fake.secret.Secret.IsSimpleMQSecret())

	_, err = deleteCommand.Func(nil, &deleteParameter{IDParameter: cflag.IDParameter{ID: "delete-id"}})
	require.NoError(t, err)
	require.Equal(t, "delete-id", fake.deleteID)
}

func TestProcessConfigurationRequestBuildsTypedPatch(t *testing.T) {
	tags := []string{"tag-a"}
	settings := `{"Destination":"simplemq","Parameters":"{}"}`
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
	require.True(t, item.Settings.Value.IsProcessConfigurationSettings())
	require.Equal(t, v1.ProviderClassEventbusprocessconfiguration, item.Provider.Value.Class)
}

func TestSecretRequest(t *testing.T) {
	t.Run("Sacloud API", func(t *testing.T) {
		request, err := secretRequestFromReader(bytes.NewBufferString(
			`{"AccessToken":"not-a-real-token","AccessTokenSecret":"not-a-real-secret"}`,
		))
		require.NoError(t, err)
		require.True(t, request.Secret.IsSacloudAPISecret())
	})

	t.Run("standard input", func(t *testing.T) {
		request, err := secretRequest(&stdinContext{
			input: standardInput(t, `{"APIKey":"not-a-real-secret"}`),
		}, "-")
		require.NoError(t, err)
		require.True(t, request.Secret.IsSimpleMQSecret())
	})

	t.Run("inline values are rejected", func(t *testing.T) {
		_, err := secretRequest(nil, `{"APIKey":"not-a-real-secret"}`)
		require.Error(t, err)
	})
}

func TestProcessConfigurationUpdateRejectsDescriptionAndClear(t *testing.T) {
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

func standardInput(t *testing.T, content string) *os.File {
	t.Helper()
	input, writer, err := os.Pipe()
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, input.Close())
	})
	_, err = writer.WriteString(content)
	require.NoError(t, err)
	require.NoError(t, writer.Close())
	return input
}

type stdinContext struct {
	input *os.File
}

func (c *stdinContext) Option() *config.Config       { return &config.Config{} }
func (c *stdinContext) Output() output.Output        { return output.NewDiscardOutput() }
func (c *stdinContext) Client() interface{}          { return nil }
func (c *stdinContext) IO() cli.IO                   { return stdinIO{input: c.input} }
func (c *stdinContext) Args() []string               { return nil }
func (c *stdinContext) Saclient() saclient.ClientAPI { return nil }
func (c *stdinContext) SDKClient() (sdksaclient.ClientAPI, error) {
	return nil, nil
}
func (c *stdinContext) PlatformName() string { return "" }
func (c *stdinContext) ResourceName() string { return "" }
func (c *stdinContext) CommandName() string  { return "" }
func (c *stdinContext) WithResource(string, string, interface{}) cli.Context {
	return c
}
func (c *stdinContext) ID() string                    { return "" }
func (c *stdinContext) Zone() string                  { return "" }
func (c *stdinContext) Resource() interface{}         { return nil }
func (c *stdinContext) Deadline() (time.Time, bool)   { return time.Time{}, false }
func (c *stdinContext) Done() <-chan struct{}         { return nil }
func (c *stdinContext) Err() error                    { return nil }
func (c *stdinContext) Value(interface{}) interface{} { return nil }

type stdinIO struct {
	input *os.File
}

func (p stdinIO) In() *os.File      { return p.input }
func (stdinIO) Out() io.Writer      { return io.Discard }
func (stdinIO) Progress() io.Writer { return io.Discard }
func (stdinIO) Err() io.Writer      { return io.Discard }
