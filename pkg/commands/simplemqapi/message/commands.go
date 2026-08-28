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
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/sacloud/sacloud-sdk-go/api/simplemq"
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/commands/simplemqapi/common"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/validate"
)

var messageOpFactory = common.NewMessageOp

var (
	queueNamePattern = regexp.MustCompile(`^[0-9a-zA-Z]+(-[0-9a-zA-Z]+)*$`)
	messageIDPattern = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	contentPattern   = regexp.MustCompile(`^[0-9a-zA-Z+/=]*$`)
)

var sendCommand = &core.Command{
	Name:       "send",
	Usage:      "Send a message to a SimpleMQ queue",
	Category:   "operation",
	Order:      10,
	NoProgress: true,
	ColumnDefs: newMessageColumnDefs,
	ParameterInitializer: func() interface{} {
		return newSendParameter()
	},
	ValidateFunc: validateSendParameter,
	Func:         sendFunc,
}

type sendParameter struct {
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	QueueName  string `cli:"queue-name,desc=Queue name (5-64 alphanumeric characters separated by single hyphens)" validate:"required,min=5,max=64"`
	APIKeyFile string `cli:"api-key-file,desc=Path to the queue API key file or - to read from standard input" validate:"required"`
	Content    string `cli:"content,desc=Base64 message content" validate:"max=256000"`
}

func newSendParameter() *sendParameter {
	return &sendParameter{}
}

var receiveCommand = &core.Command{
	Name:       "receive",
	Usage:      "Receive messages from a SimpleMQ queue",
	LongUsage:  "Receive messages from a SimpleMQ queue. Receiving changes message visibility until the queue's visibility timeout.",
	Category:   "operation",
	Order:      20,
	NoProgress: true,
	ColumnDefs: messageColumnDefs,
	ParameterInitializer: func() interface{} {
		return newReceiveParameter()
	},
	ValidateFunc: validateMessageAccessParameter,
	Func:         receiveFunc,
}

type receiveParameter struct {
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	QueueName  string `cli:"queue-name,desc=Queue name (5-64 alphanumeric characters separated by single hyphens)" validate:"required,min=5,max=64"`
	APIKeyFile string `cli:"api-key-file,desc=Path to the queue API key file or - to read from standard input" validate:"required"`
}

func newReceiveParameter() *receiveParameter {
	return &receiveParameter{}
}

var extendTimeoutCommand = &core.Command{
	Name:       "extend-timeout",
	Usage:      "Extend a message visibility timeout",
	Category:   "operation",
	Order:      30,
	NoProgress: true,
	ColumnDefs: messageColumnDefs,
	ParameterInitializer: func() interface{} {
		return newExtendTimeoutParameter()
	},
	ValidateFunc: validateMessageIDParameter,
	Func:         extendTimeoutFunc,
}

type extendTimeoutParameter struct {
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	QueueName  string `cli:"queue-name,desc=Queue name (5-64 alphanumeric characters separated by single hyphens)" validate:"required,min=5,max=64"`
	APIKeyFile string `cli:"api-key-file,desc=Path to the queue API key file or - to read from standard input" validate:"required"`
	MessageID  string `cli:"message-id,desc=Lowercase UUID message ID" validate:"required,len=36"`
}

func newExtendTimeoutParameter() *extendTimeoutParameter {
	return &extendTimeoutParameter{}
}

var deleteCommand = &core.Command{
	Name:     "delete",
	Aliases:  []string{"rm"},
	Usage:    "Delete a processed message",
	Category: "operation",
	Order:    40,
	ParameterInitializer: func() interface{} {
		return newDeleteParameter()
	},
	ValidateFunc: validateMessageIDParameter,
	Func:         deleteFunc,
}

type deleteParameter struct {
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	QueueName  string `cli:"queue-name,desc=Queue name (5-64 alphanumeric characters separated by single hyphens)" validate:"required,min=5,max=64"`
	APIKeyFile string `cli:"api-key-file,desc=Path to the queue API key file or - to read from standard input" validate:"required"`
	MessageID  string `cli:"message-id,desc=Lowercase UUID message ID" validate:"required,len=36"`
}

func newDeleteParameter() *deleteParameter {
	return &deleteParameter{}
}

func init() {
	Resource.AddCommand(sendCommand)
	Resource.AddCommand(receiveCommand)
	Resource.AddCommand(extendTimeoutCommand)
	Resource.AddCommand(deleteCommand)
}

func sendFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	p, ok := parameter.(*sendParameter)
	if !ok {
		return nil, invalidParameterError(parameter)
	}
	op, err := newMessageOp(ctx, p.QueueName, p.APIKeyFile)
	if err != nil {
		return nil, err
	}
	result, err := op.Send(ctx, p.Content)
	if err != nil {
		return nil, err
	}
	return []interface{}{result}, nil
}

func receiveFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	p, ok := parameter.(*receiveParameter)
	if !ok {
		return nil, invalidParameterError(parameter)
	}
	op, err := newMessageOp(ctx, p.QueueName, p.APIKeyFile)
	if err != nil {
		return nil, err
	}
	items, err := op.Receive(ctx)
	if err != nil {
		return nil, err
	}
	results := make([]interface{}, 0, len(items))
	for i := range items {
		results = append(results, &items[i])
	}
	return results, nil
}

func extendTimeoutFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	p, ok := parameter.(*extendTimeoutParameter)
	if !ok {
		return nil, invalidParameterError(parameter)
	}
	op, err := newMessageOp(ctx, p.QueueName, p.APIKeyFile)
	if err != nil {
		return nil, err
	}
	result, err := op.ExtendTimeout(ctx, p.MessageID)
	if err != nil {
		return nil, err
	}
	return []interface{}{result}, nil
}

func deleteFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	p, ok := parameter.(*deleteParameter)
	if !ok {
		return nil, invalidParameterError(parameter)
	}
	op, err := newMessageOp(ctx, p.QueueName, p.APIKeyFile)
	if err != nil {
		return nil, err
	}
	if err := op.Delete(ctx, p.MessageID); err != nil {
		return nil, err
	}
	return nil, nil
}

func newMessageOp(ctx cli.Context, queueName, apiKeyFile string) (simplemq.MessageAPI, error) {
	apiKey, err := readAPIKey(ctx, apiKeyFile)
	if err != nil {
		return nil, err
	}
	return messageOpFactory(ctx, apiKey, queueName)
}

func readAPIKey(ctx cli.Context, path string) (string, error) {
	var data []byte
	var err error
	if path == "-" {
		data, err = io.ReadAll(ctx.IO().In())
	} else {
		data, err = os.ReadFile(path) //nolint:gosec // Path is explicitly supplied through --api-key-file.
	}
	if err != nil {
		return "", fmt.Errorf("read API key: %w", err)
	}
	apiKey := strings.TrimSpace(string(data))
	if apiKey == "" {
		return "", fmt.Errorf("API key file is empty")
	}
	return apiKey, nil
}

func validateSendParameter(_ cli.Context, parameter interface{}) error {
	p, ok := parameter.(*sendParameter)
	if !ok {
		return invalidParameterError(parameter)
	}
	if err := validate.Exec(p); err != nil {
		return err
	}
	if err := validateQueueName(p.QueueName); err != nil {
		return err
	}
	if !contentPattern.MatchString(p.Content) {
		return validate.NewFlagError("--content", "must match the SDK Base64 character pattern")
	}
	return nil
}

func validateMessageAccessParameter(_ cli.Context, parameter interface{}) error {
	p, ok := parameter.(*receiveParameter)
	if !ok {
		return invalidParameterError(parameter)
	}
	if err := validate.Exec(p); err != nil {
		return err
	}
	return validateQueueName(p.QueueName)
}

func validateMessageIDParameter(_ cli.Context, parameter interface{}) error {
	var queueName, messageID string
	switch p := parameter.(type) {
	case *extendTimeoutParameter:
		if err := validate.Exec(p); err != nil {
			return err
		}
		queueName, messageID = p.QueueName, p.MessageID
	case *deleteParameter:
		if err := validate.Exec(p); err != nil {
			return err
		}
		queueName, messageID = p.QueueName, p.MessageID
	default:
		return invalidParameterError(parameter)
	}
	if err := validateQueueName(queueName); err != nil {
		return err
	}
	if !messageIDPattern.MatchString(messageID) {
		return validate.NewFlagError("--message-id", "must be a lowercase UUID")
	}
	return nil
}

func validateQueueName(queueName string) error {
	if !queueNamePattern.MatchString(queueName) {
		return validate.NewFlagError("--queue-name", "must contain alphanumeric groups separated by single hyphens")
	}
	return nil
}

func invalidParameterError(parameter interface{}) error {
	return fmt.Errorf("got invalid parameter type: %#v", parameter)
}
