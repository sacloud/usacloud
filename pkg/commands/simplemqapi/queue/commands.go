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
	"fmt"
	"regexp"

	sdkqueue "github.com/sacloud/sacloud-sdk-go/api/simplemq/apis/v1/queue"
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/commands/simplemqapi/common"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/validate"
)

var queueOpFactory = common.NewQueueOp

var queueNamePattern = regexp.MustCompile(`^[0-9a-zA-Z]+(-[0-9a-zA-Z]+)*$`)

var listCommand = &core.Command{
	Name:       "list",
	Aliases:    []string{"ls", "find", "select"},
	Usage:      "List SimpleMQ queues",
	Category:   "basic",
	Order:      10,
	NoProgress: true,
	ColumnDefs: defaultColumnDefs,
	ParameterInitializer: func() interface{} {
		return newListParameter()
	},
	Func: listFunc,
}

type listParameter struct {
	cflag.CommonParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter `cli:",squash" mapconv:"-"`
}

func newListParameter() *listParameter {
	return &listParameter{}
}

var createCommand = &core.Command{
	Name:      "create",
	Usage:     "Create a SimpleMQ queue",
	LongUsage: "Create a SimpleMQ queue. The create API does not return a Message API key; issue one separately with rotate-api-key.",
	Example: `  usacloud simplemq-api queue create --name example-queue -y
  usacloud simplemq-api queue rotate-api-key example-queue -y`,
	Category:   "basic",
	Order:      20,
	ColumnDefs: defaultColumnDefs,
	ParameterInitializer: func() interface{} {
		return newCreateParameter()
	},
	ValidateFunc: validateQueueNameParameter,
	Func:         createFunc,
}

type createParameter struct {
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	Name        string   `cli:"name,desc=Queue name (5-64 alphanumeric characters separated by single hyphens)" validate:"required,min=5,max=64"`
	Description string   `cli:"description"`
	Tags        []string `cli:"tags"`
}

func newCreateParameter() *createParameter {
	return &createParameter{}
}

var readCommand = &core.Command{
	Name:         "read",
	Aliases:      []string{"show"},
	Usage:        "Read a SimpleMQ queue",
	Category:     "basic",
	Order:        30,
	NoProgress:   true,
	SelectorType: core.SelectorTypeRequireSingle,
	ColumnDefs:   defaultColumnDefs,
	ParameterInitializer: func() interface{} {
		return newReadParameter()
	},
	ListAllFunc: listFunc,
	Func:        readFunc,
}

type readParameter struct {
	cflag.IDParameter     `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter `cli:",squash" mapconv:"-"`
}

func newReadParameter() *readParameter {
	return &readParameter{}
}

var configCommand = &core.Command{
	Name:         "config",
	Usage:        "Configure SimpleMQ queues",
	LongUsage:    "Configure one or more SimpleMQ queues with a complete configuration. Description and tags replace their current values.",
	Category:     "basic",
	Order:        40,
	SelectorType: core.SelectorTypeRequireMulti,
	ColumnDefs:   defaultColumnDefs,
	ParameterInitializer: func() interface{} {
		return newConfigParameter()
	},
	ListAllFunc: listFunc,
	Func:        configFunc,
}

type configParameter struct {
	cflag.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	VisibilityTimeoutSeconds int      `cli:"visibility-timeout-seconds,desc=Visibility timeout in seconds" validate:"required,min=5,max=900"`
	ExpireSeconds            int      `cli:"expire-seconds,desc=Unprocessed message retention in seconds" validate:"required,min=60,max=1209600"`
	Description              string   `cli:"description"`
	Tags                     []string `cli:"tags"`
}

func newConfigParameter() *configParameter {
	return &configParameter{}
}

var deleteCommand = &core.Command{
	Name:         "delete",
	Aliases:      []string{"rm"},
	Usage:        "Delete SimpleMQ queues",
	Category:     "basic",
	Order:        50,
	SelectorType: core.SelectorTypeRequireMulti,
	ParameterInitializer: func() interface{} {
		return newDeleteParameter()
	},
	ListAllFunc: listFunc,
	Func:        deleteFunc,
}

type deleteParameter struct {
	cflag.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`
}

func newDeleteParameter() *deleteParameter {
	return &deleteParameter{}
}

var countMessagesCommand = &core.Command{
	Name:         "count-messages",
	Usage:        "Count messages in a SimpleMQ queue",
	Category:     "operation",
	Order:        10,
	NoProgress:   true,
	SelectorType: core.SelectorTypeRequireSingle,
	ColumnDefs:   countMessagesColumnDefs,
	ParameterInitializer: func() interface{} {
		return newCountMessagesParameter()
	},
	ListAllFunc: listFunc,
	Func:        countMessagesFunc,
}

type countMessagesParameter struct {
	cflag.IDParameter     `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter `cli:",squash" mapconv:"-"`
}

func newCountMessagesParameter() *countMessagesParameter {
	return &countMessagesParameter{}
}

type countMessagesResult struct {
	MessageCount int
}

var rotateAPIKeyCommand = &core.Command{
	Name:           "rotate-api-key",
	Usage:          "Rotate a SimpleMQ queue API key",
	LongUsage:      "Rotate the Message API key for one queue. Rotation invalidates the old API key.",
	Category:       "operation",
	Order:          20,
	NoProgress:     true,
	ConfirmMessage: "rotate API key and invalidate the old key",
	SelectorType:   core.SelectorTypeRequireSingle,
	ColumnDefs:     rotateAPIKeyColumnDefs,
	ParameterInitializer: func() interface{} {
		return newRotateAPIKeyParameter()
	},
	ListAllFunc: listFunc,
	Func:        rotateAPIKeyFunc,
}

type rotateAPIKeyParameter struct {
	cflag.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`
}

func newRotateAPIKeyParameter() *rotateAPIKeyParameter {
	return &rotateAPIKeyParameter{}
}

type rotateAPIKeyResult struct {
	APIKey string
}

var clearMessagesCommand = &core.Command{
	Name:         "clear-messages",
	Usage:        "Delete all messages from SimpleMQ queues",
	Category:     "operation",
	Order:        30,
	SelectorType: core.SelectorTypeRequireMulti,
	ParameterInitializer: func() interface{} {
		return newClearMessagesParameter()
	},
	ListAllFunc: listFunc,
	Func:        clearMessagesFunc,
}

type clearMessagesParameter struct {
	cflag.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`
}

func newClearMessagesParameter() *clearMessagesParameter {
	return &clearMessagesParameter{}
}

func init() {
	Resource.AddCommand(listCommand)
	Resource.AddCommand(createCommand)
	Resource.AddCommand(readCommand)
	Resource.AddCommand(configCommand)
	Resource.AddCommand(deleteCommand)
	Resource.AddCommand(countMessagesCommand)
	Resource.AddCommand(rotateAPIKeyCommand)
	Resource.AddCommand(clearMessagesCommand)
}

func listFunc(ctx cli.Context, _ interface{}) ([]interface{}, error) {
	op, err := queueOpFactory(ctx)
	if err != nil {
		return nil, err
	}
	items, err := op.List(ctx)
	if err != nil {
		return nil, err
	}
	results := make([]interface{}, 0, len(items))
	for i := range items {
		results = append(results, &items[i])
	}
	return results, nil
}

func createFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	p, ok := parameter.(*createParameter)
	if !ok {
		return nil, invalidParameterError(parameter)
	}
	op, err := queueOpFactory(ctx)
	if err != nil {
		return nil, err
	}
	result, err := op.Create(ctx, createRequest(p))
	if err != nil {
		return nil, err
	}
	return []interface{}{result}, nil
}

func readFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	p, ok := parameter.(*readParameter)
	if !ok {
		return nil, invalidParameterError(parameter)
	}
	op, err := queueOpFactory(ctx)
	if err != nil {
		return nil, err
	}
	result, err := op.Read(ctx, p.ID)
	if err != nil {
		return nil, err
	}
	return []interface{}{result}, nil
}

func configFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	p, ok := parameter.(*configParameter)
	if !ok {
		return nil, invalidParameterError(parameter)
	}
	op, err := queueOpFactory(ctx)
	if err != nil {
		return nil, err
	}
	result, err := op.Config(ctx, p.ID, configRequest(p))
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
	op, err := queueOpFactory(ctx)
	if err != nil {
		return nil, err
	}
	if err := op.Delete(ctx, p.ID); err != nil {
		return nil, err
	}
	return nil, nil
}

func countMessagesFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	p, ok := parameter.(*countMessagesParameter)
	if !ok {
		return nil, invalidParameterError(parameter)
	}
	op, err := queueOpFactory(ctx)
	if err != nil {
		return nil, err
	}
	count, err := op.CountMessages(ctx, p.ID)
	if err != nil {
		return nil, err
	}
	return []interface{}{&countMessagesResult{MessageCount: count}}, nil
}

func rotateAPIKeyFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	p, ok := parameter.(*rotateAPIKeyParameter)
	if !ok {
		return nil, invalidParameterError(parameter)
	}
	op, err := queueOpFactory(ctx)
	if err != nil {
		return nil, err
	}
	apiKey, err := op.RotateAPIKey(ctx, p.ID)
	if err != nil {
		return nil, err
	}
	return []interface{}{&rotateAPIKeyResult{APIKey: apiKey}}, nil
}

func clearMessagesFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	p, ok := parameter.(*clearMessagesParameter)
	if !ok {
		return nil, invalidParameterError(parameter)
	}
	op, err := queueOpFactory(ctx)
	if err != nil {
		return nil, err
	}
	if err := op.ClearMessages(ctx, p.ID); err != nil {
		return nil, err
	}
	return nil, nil
}

func createRequest(p *createParameter) sdkqueue.CreateQueueRequest {
	return sdkqueue.CreateQueueRequest{
		CommonServiceItem: sdkqueue.CreateQueueRequestCommonServiceItem{
			Name:        sdkqueue.QueueName(p.Name),
			Description: sdkqueue.NewOptString(p.Description),
			Tags:        p.Tags,
		},
	}
}

func configRequest(p *configParameter) sdkqueue.ConfigQueueRequest {
	return sdkqueue.ConfigQueueRequest{
		CommonServiceItem: sdkqueue.ConfigQueueRequestCommonServiceItem{
			Description: sdkqueue.NewOptString(p.Description),
			Settings: sdkqueue.Settings{
				VisibilityTimeoutSeconds: sdkqueue.VisibilityTimeoutSeconds(p.VisibilityTimeoutSeconds),
				ExpireSeconds:            sdkqueue.ExpireSeconds(p.ExpireSeconds),
			},
			Tags: p.Tags,
		},
	}
}

func validateQueueNameParameter(_ cli.Context, parameter interface{}) error {
	p, ok := parameter.(*createParameter)
	if !ok {
		return invalidParameterError(parameter)
	}
	if err := validate.Exec(p); err != nil {
		return err
	}
	if !queueNamePattern.MatchString(p.Name) {
		return validate.NewFlagError("--name", "must contain alphanumeric groups separated by single hyphens")
	}
	return nil
}

func invalidParameterError(parameter interface{}) error {
	return fmt.Errorf("got invalid parameter type: %#v", parameter)
}
