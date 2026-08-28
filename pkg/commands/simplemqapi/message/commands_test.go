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

package message

import (
	"context"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/sacloud/saclient-go"
	"github.com/sacloud/sacloud-sdk-go/api/simplemq"
	sdkmessage "github.com/sacloud/sacloud-sdk-go/api/simplemq/apis/v1/message"
	sdksaclient "github.com/sacloud/sacloud-sdk-go/common/saclient"
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/config"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/output"
	"github.com/sacloud/usacloud/pkg/validate"
	"github.com/stretchr/testify/require"
)

func init() {
	validate.InitializeValidator(nil)
}

type messageAPIFake struct {
	sendContent     string
	receiveResult   []sdkmessage.Message
	extendMessageID string
	deleteMessageID string
}

var _ simplemq.MessageAPI = (*messageAPIFake)(nil)

func (f *messageAPIFake) Send(_ context.Context, content string) (*sdkmessage.NewMessage, error) {
	f.sendContent = content
	return &sdkmessage.NewMessage{ID: "00000000-0000-0000-0000-000000000000", Content: sdkmessage.MessageContent(content)}, nil
}

func (f *messageAPIFake) Receive(context.Context) ([]sdkmessage.Message, error) {
	return f.receiveResult, nil
}

func (f *messageAPIFake) ExtendTimeout(_ context.Context, messageID string) (*sdkmessage.Message, error) {
	f.extendMessageID = messageID
	return &sdkmessage.Message{ID: sdkmessage.MessageId(messageID)}, nil
}

func (f *messageAPIFake) Delete(_ context.Context, messageID string) error {
	f.deleteMessageID = messageID
	return nil
}

func useMessageAPIFake(t *testing.T, op simplemq.MessageAPI, expectedQueueName string) {
	t.Helper()
	oldFactory := messageOpFactory
	messageOpFactory = func(_ cli.Context, apiKey, queueName string) (simplemq.MessageAPI, error) {
		require.NotEmpty(t, apiKey)
		require.Equal(t, expectedQueueName, queueName)
		return op, nil
	}
	t.Cleanup(func() {
		messageOpFactory = oldFactory
	})
}

func TestMessageCommandsInvokeSDKInterface(t *testing.T) {
	const (
		queueName = "test-queue"
		messageID = "00000000-0000-0000-0000-000000000000"
	)
	fake := &messageAPIFake{
		receiveResult: []sdkmessage.Message{{ID: messageID, Content: "YQ=="}},
	}
	useMessageAPIFake(t, fake, queueName)

	results, err := sendCommand.Func(newInputContext(t, " dummy-value \n"), &sendParameter{
		QueueName:  queueName,
		APIKeyFile: "-",
		Content:    "YQ==",
	})
	require.NoError(t, err)
	require.Equal(t, "YQ==", fake.sendContent)
	require.Len(t, results, 1)
	require.IsType(t, &sdkmessage.NewMessage{}, results[0])

	results, err = receiveCommand.Func(newInputContext(t, "dummy-value"), &receiveParameter{
		QueueName:  queueName,
		APIKeyFile: "-",
	})
	require.NoError(t, err)
	require.Len(t, results, 1)
	require.IsType(t, &sdkmessage.Message{}, results[0])

	results, err = extendTimeoutCommand.Func(newInputContext(t, "dummy-value"), &extendTimeoutParameter{
		QueueName:  queueName,
		APIKeyFile: "-",
		MessageID:  messageID,
	})
	require.NoError(t, err)
	require.Equal(t, messageID, fake.extendMessageID)
	require.Len(t, results, 1)

	results, err = deleteCommand.Func(newInputContext(t, "dummy-value"), &deleteParameter{
		QueueName:  queueName,
		APIKeyFile: "-",
		MessageID:  messageID,
	})
	require.NoError(t, err)
	require.Equal(t, messageID, fake.deleteMessageID)
	require.Nil(t, results)
}

func TestReadAPIKey(t *testing.T) {
	t.Run("standard input trims surrounding whitespace", func(t *testing.T) {
		value, err := readAPIKey(newInputContext(t, " \n dummy-value\t"), "-")
		require.NoError(t, err)
		require.Equal(t, "dummy-value", value)
	})

	t.Run("file", func(t *testing.T) {
		const path = ".simplemq-api-key-test"
		require.NoError(t, os.WriteFile(path, []byte("\tdummy-value\n"), 0o600))
		t.Cleanup(func() {
			require.NoError(t, os.Remove(path))
		})

		value, err := readAPIKey(nil, path)
		require.NoError(t, err)
		require.Equal(t, "dummy-value", value)
	})

	t.Run("empty", func(t *testing.T) {
		_, err := readAPIKey(newInputContext(t, " \n\t"), "-")
		require.EqualError(t, err, "API key file is empty")
	})
}

func TestMessageValidation(t *testing.T) {
	require.NoError(t, validateSendParameter(nil, &sendParameter{
		QueueName:  "valid-queue",
		APIKeyFile: "-",
		Content:    "YQ==",
	}))
	require.NoError(t, validateSendParameter(nil, &sendParameter{
		QueueName:  "valid-queue",
		APIKeyFile: "-",
		Content:    "",
	}))
	require.Error(t, validateSendParameter(nil, &sendParameter{
		QueueName:  "invalid_queue",
		APIKeyFile: "-",
		Content:    "YQ==",
	}))
	require.Error(t, validateSendParameter(nil, &sendParameter{
		QueueName:  "valid-queue",
		APIKeyFile: "-",
		Content:    "not*base64",
	}))
	require.Error(t, validateSendParameter(nil, &sendParameter{
		QueueName:  "valid-queue",
		APIKeyFile: "-",
		Content:    strings.Repeat("A", 256001),
	}))

	validID := "0193b878-1b25-7775-87f5-9c698206a7e7"
	require.NoError(t, validateMessageIDParameter(nil, &extendTimeoutParameter{
		QueueName:  "valid-queue",
		APIKeyFile: "-",
		MessageID:  validID,
	}))
	require.Error(t, validateMessageIDParameter(nil, &deleteParameter{
		QueueName:  "valid-queue",
		APIKeyFile: "-",
		MessageID:  strings.ToUpper(validID),
	}))
}

func TestMessageCommandMetadata(t *testing.T) {
	for _, command := range []*core.Command{sendCommand, receiveCommand, extendTimeoutCommand, deleteCommand} {
		require.Equal(t, core.SelectorType(core.SelectorTypeNone), command.SelectorType)
		require.Nil(t, command.ListAllFunc)
		require.Implements(t, (*cflag.ConfirmParameterValueHandler)(nil), command.ParameterInitializer())
	}
	for _, command := range []*core.Command{sendCommand, receiveCommand, extendTimeoutCommand} {
		require.True(t, command.NoProgress)
		require.NotEmpty(t, command.ColumnDefs)
	}
	require.Contains(t, receiveCommand.LongUsage, "changes message visibility")
}

func newInputContext(t *testing.T, content string) cli.Context {
	t.Helper()
	input, writer, err := os.Pipe()
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, input.Close())
	})
	_, err = writer.WriteString(content)
	require.NoError(t, err)
	require.NoError(t, writer.Close())
	return &inputContext{input: input}
}

type inputContext struct {
	input *os.File
}

func (c *inputContext) Option() *config.Config       { return &config.Config{} }
func (c *inputContext) Output() output.Output        { return output.NewDiscardOutput() }
func (c *inputContext) Client() interface{}          { return nil }
func (c *inputContext) IO() cli.IO                   { return inputIO{input: c.input} }
func (c *inputContext) Args() []string               { return nil }
func (c *inputContext) Saclient() saclient.ClientAPI { return nil }
func (c *inputContext) SDKClient() (sdksaclient.ClientAPI, error) {
	return nil, nil
}
func (c *inputContext) PlatformName() string { return "" }
func (c *inputContext) ResourceName() string { return "" }
func (c *inputContext) CommandName() string  { return "" }
func (c *inputContext) WithResource(string, string, interface{}) cli.Context {
	return c
}
func (c *inputContext) ID() string                    { return "" }
func (c *inputContext) Zone() string                  { return "" }
func (c *inputContext) Resource() interface{}         { return nil }
func (c *inputContext) Deadline() (time.Time, bool)   { return time.Time{}, false }
func (c *inputContext) Done() <-chan struct{}         { return nil }
func (c *inputContext) Err() error                    { return nil }
func (c *inputContext) Value(interface{}) interface{} { return nil }

type inputIO struct {
	input *os.File
}

func (p inputIO) In() *os.File      { return p.input }
func (inputIO) Out() io.Writer      { return io.Discard }
func (inputIO) Progress() io.Writer { return io.Discard }
func (inputIO) Err() io.Writer      { return io.Discard }
