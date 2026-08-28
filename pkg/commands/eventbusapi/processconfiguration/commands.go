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
	"fmt"
	"io"
	"os"

	"github.com/sacloud/sacloud-sdk-go/api/eventbus"
	v1 "github.com/sacloud/sacloud-sdk-go/api/eventbus/apis/v1"
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/commands/eventbusapi/common"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/validate"
)

var processConfigurationOpFactory = func(ctx cli.Context) (eventbus.ProcessConfigurationAPI, error) {
	client, err := common.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return eventbus.NewProcessConfigurationOp(client), nil
}

const processConfigurationSettingsHelp = `Process configuration settings are a JSON object or a path to a JSON file.

Required fields:
  Destination  One of simplenotification, simplemq, or autoscale
  Parameters   Destination-specific parameters encoded as a JSON string`

var listCommand = &core.Command{
	Name:       "list",
	Aliases:    []string{"ls", "find", "select"},
	Usage:      "List EventBus process configurations",
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
	Usage:     "Create an EventBus process configuration",
	LongUsage: "Create an EventBus process configuration.\n\n" + processConfigurationSettingsHelp,
	Example:   `  usacloud eventbus-api process-configuration create --name queue-job --settings '{"Destination":"simplemq","Parameters":"{\"queue_name\":\"my-queue\",\"content\":\"hello\"}"}' -y`,
	Category:  "basic",
	Order:     20,

	ColumnDefs: defaultColumnDefs,
	ParameterInitializer: func() interface{} {
		return newCreateParameter()
	},
	Func: createFunc,
}

type createParameter struct {
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag.NameParameter `cli:",squash" mapconv:",squash"`
	cflag.DescParameter `cli:",squash" mapconv:",squash"`
	cflag.TagsParameter `cli:",squash" mapconv:",squash"`

	Settings string `cli:"settings,desc=Process configuration settings JSON or file path; see command help for fields and examples" validate:"required" mapconv:"-" json:"-"`
}

func newCreateParameter() *createParameter {
	return &createParameter{}
}

var readCommand = &core.Command{
	Name:         "read",
	Aliases:      []string{"show"},
	Usage:        "Read an EventBus process configuration",
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

var updateCommand = &core.Command{
	Name:         "update",
	Usage:        "Update EventBus process configurations",
	LongUsage:    "Update one or more EventBus process configurations.\n\n" + processConfigurationSettingsHelp,
	Example:      `  usacloud eventbus-api process-configuration update configuration-id --settings '{"Destination":"simplemq","Parameters":"{\"queue_name\":\"my-queue\",\"content\":\"updated\"}"}' -y`,
	Category:     "basic",
	Order:        40,
	SelectorType: core.SelectorTypeRequireMulti,
	ColumnDefs:   defaultColumnDefs,
	ParameterInitializer: func() interface{} {
		return newUpdateParameter()
	},
	ValidateFunc: validateUpdateParameter,
	ListAllFunc:  listFunc,
	Func:         updateFunc,
}

type updateParameter struct {
	cflag.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag.NameUpdateParameter `cli:",squash" mapconv:",omitempty,squash"`
	cflag.DescUpdateParameter `cli:",squash" mapconv:",omitempty,squash"`
	cflag.TagsUpdateParameter `cli:",squash" mapconv:",omitempty,squash"`

	ClearDescription bool    `cli:"clear-description,desc=Clear the description" mapconv:"-"`
	Settings         *string `cli:"settings,desc=Process configuration settings JSON or file path; see command help for fields and examples" mapconv:"-" json:"-"`
}

func newUpdateParameter() *updateParameter {
	return &updateParameter{}
}

var updateSecretCommand = &core.Command{
	Name:         "update-secret",
	Usage:        "Update credentials used by a process configuration",
	Category:     "basic",
	Order:        45,
	SelectorType: core.SelectorTypeRequireMulti,
	ParameterInitializer: func() interface{} {
		return newUpdateSecretParameter()
	},
	ListAllFunc: listFunc,
	Func:        updateSecretFunc,
}

type updateSecretParameter struct {
	cflag.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	SecretFile string `cli:"secret-file,desc=Path to a JSON secret file or - to read from standard input" validate:"required" mapconv:"-" json:"-"`
}

func newUpdateSecretParameter() *updateSecretParameter {
	return &updateSecretParameter{}
}

var deleteCommand = &core.Command{
	Name:         "delete",
	Aliases:      []string{"rm"},
	Usage:        "Delete EventBus process configurations",
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

func init() {
	Resource.AddCommand(listCommand)
	Resource.AddCommand(createCommand)
	Resource.AddCommand(readCommand)
	Resource.AddCommand(updateCommand)
	Resource.AddCommand(updateSecretCommand)
	Resource.AddCommand(deleteCommand)
}

func listFunc(ctx cli.Context, _ interface{}) ([]interface{}, error) {
	op, err := processConfigurationOpFactory(ctx)
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
	request, err := createRequest(p)
	if err != nil {
		return nil, err
	}

	op, err := processConfigurationOpFactory(ctx)
	if err != nil {
		return nil, err
	}
	result, err := op.Create(ctx, request)
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

	op, err := processConfigurationOpFactory(ctx)
	if err != nil {
		return nil, err
	}
	result, err := op.Read(ctx, p.ID)
	if err != nil {
		return nil, err
	}
	return []interface{}{result}, nil
}

func updateFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	p, ok := parameter.(*updateParameter)
	if !ok {
		return nil, invalidParameterError(parameter)
	}
	request, err := updateRequest(p)
	if err != nil {
		return nil, err
	}

	op, err := processConfigurationOpFactory(ctx)
	if err != nil {
		return nil, err
	}
	result, err := op.Update(ctx, p.ID, request)
	if err != nil {
		return nil, err
	}
	return []interface{}{result}, nil
}

func updateSecretFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	p, ok := parameter.(*updateSecretParameter)
	if !ok {
		return nil, invalidParameterError(parameter)
	}
	secret, err := secretRequest(ctx, p.SecretFile)
	if err != nil {
		return nil, err
	}

	op, err := processConfigurationOpFactory(ctx)
	if err != nil {
		return nil, err
	}
	if err := op.UpdateSecret(ctx, p.ID, secret); err != nil {
		return nil, err
	}
	return nil, nil
}

func deleteFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	p, ok := parameter.(*deleteParameter)
	if !ok {
		return nil, invalidParameterError(parameter)
	}

	op, err := processConfigurationOpFactory(ctx)
	if err != nil {
		return nil, err
	}
	if err := op.Delete(ctx, p.ID); err != nil {
		return nil, err
	}
	return nil, nil
}

func createRequest(p *createParameter) (v1.CreateCommonServiceItemRequest, error) {
	settings, err := processConfigurationSettings(p.Settings)
	if err != nil {
		return v1.CreateCommonServiceItemRequest{}, err
	}
	return v1.CreateCommonServiceItemRequest{
		CommonServiceItem: v1.CreateCommonServiceItemRequestCommonServiceItem{
			Name:        p.Name,
			Description: v1.NewOptNilString(p.Description),
			Tags:        p.Tags,
			Settings:    v1.NewProcessConfigurationSettingsSettings(settings),
			Provider:    v1.Provider{Class: v1.ProviderClassEventbusprocessconfiguration},
		},
	}, nil
}

func updateRequest(p *updateParameter) (v1.UpdateCommonServiceItemRequest, error) {
	item := v1.UpdateCommonServiceItemRequestCommonServiceItem{
		Provider: v1.NewOptProvider(v1.Provider{Class: v1.ProviderClassEventbusprocessconfiguration}),
	}
	if p.Name != nil {
		item.Name = v1.NewOptString(*p.Name)
	}
	if p.ClearDescription {
		item.Description.SetToNull()
	} else if p.Description != nil {
		item.Description.SetTo(*p.Description)
	}
	if p.Tags != nil {
		item.Tags = *p.Tags
	}
	if p.Settings != nil {
		settings, err := processConfigurationSettings(*p.Settings)
		if err != nil {
			return v1.UpdateCommonServiceItemRequest{}, err
		}
		item.Settings = v1.NewOptSettings(v1.NewProcessConfigurationSettingsSettings(settings))
	}
	return v1.UpdateCommonServiceItemRequest{CommonServiceItem: item}, nil
}

func processConfigurationSettings(data string) (v1.ProcessConfigurationSettings, error) {
	var settings v1.ProcessConfigurationSettings
	if err := common.DecodeJSONPathOrContent(data, &settings); err != nil {
		return v1.ProcessConfigurationSettings{}, err
	}
	return settings, nil
}

func secretRequest(ctx cli.Context, path string) (v1.SetSecretRequest, error) {
	if path == "-" {
		return secretRequestFromReader(ctx.IO().In())
	}
	data, err := os.ReadFile(path) //nolint:gosec // Path is provided explicitly through --secret-file.
	if err != nil {
		return v1.SetSecretRequest{}, err
	}
	return secretRequestFromReader(bytes.NewReader(data))
}

func secretRequestFromReader(reader io.Reader) (v1.SetSecretRequest, error) {
	var secret v1.SetSecretRequestSecret
	if err := common.DecodeJSON(reader, &secret); err != nil {
		return v1.SetSecretRequest{}, fmt.Errorf("decode secret JSON: %w", err)
	}
	return v1.SetSecretRequest{Secret: secret}, nil
}

func validateUpdateParameter(_ cli.Context, parameter interface{}) error {
	p, ok := parameter.(*updateParameter)
	if !ok {
		return invalidParameterError(parameter)
	}
	if p.ClearDescription && p.Description != nil {
		return fmt.Errorf("--description and --clear-description cannot be used together")
	}
	return validate.Exec(p)
}

func invalidParameterError(parameter interface{}) error {
	return fmt.Errorf("got invalid parameter type: %#v", parameter)
}
