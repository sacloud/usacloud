// Copyright 2017-2020 The Usacloud Authors
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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-commands'; DO NOT EDIT

package cli

import (
	"encoding/json"
	"fmt"

	"github.com/imdario/mergo"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/funcs"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/schema"
	"gopkg.in/urfave/cli.v2"
)

func init() {
	objectstorageListParam := params.NewListObjectstorageParam()
	objectstoragePutParam := params.NewPutObjectstorageParam()
	objectstorageGetParam := params.NewGetObjectstorageParam()
	objectstorageDeleteParam := params.NewDeleteObjectstorageParam()

	cliCommand := &cli.Command{
		Name:    "object-storage",
		Aliases: []string{"ojs"},
		Usage:   "A manage commands of ObjectStorage",
		Subcommands: []*cli.Command{
			{
				Name:      "list",
				Aliases:   []string{"ls"},
				Usage:     "List Objectstorage",
				ArgsUsage: "<remote path>",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "access-key",
						Usage:   "[Required] set access-key",
						EnvVars: []string{"SACLOUD_OJS_ACCESS_KEY_ID", "AWS_ACCESS_KEY_ID"},
					},
					&cli.StringFlag{
						Name:    "secret-key",
						Usage:   "[Required] set access-key",
						EnvVars: []string{"SACLOUD_OJS_SECRET_ACCESS_KEY", "AWS_SECRET_ACCESS_KEY"},
					},
					&cli.StringFlag{
						Name:    "bucket",
						Usage:   "set bucket",
						EnvVars: []string{"SACLOUD_OJS_BUCKET_NAME"},
					},
					&cli.StringFlag{
						Name:  "param-template",
						Usage: "Set input parameter from string(JSON)",
					},
					&cli.StringFlag{
						Name:  "parameters",
						Usage: "Set input parameters from JSON string",
					},
					&cli.StringFlag{
						Name:  "param-template-file",
						Usage: "Set input parameter from file",
					},
					&cli.StringFlag{
						Name:  "parameter-file",
						Usage: "Set input parameters from file",
					},
					&cli.BoolFlag{
						Name:  "generate-skeleton",
						Usage: "Output skelton of parameter JSON",
					},
					&cli.StringFlag{
						Name:    "output-type",
						Aliases: []string{"out", "o"},
						Usage:   "Output type [table/json/csv/tsv]",
					},
					&cli.StringSliceFlag{
						Name:    "column",
						Aliases: []string{"col"},
						Usage:   "Output columns(using when '--output-type' is in [csv/tsv] only)",
					},
					&cli.BoolFlag{
						Name:    "quiet",
						Aliases: []string{"q"},
						Usage:   "Only display IDs",
					},
					&cli.StringFlag{
						Name:    "format",
						Aliases: []string{"fmt"},
						Usage:   "Output format(see text/template package document for detail)",
					},
					&cli.StringFlag{
						Name:  "format-file",
						Usage: "Output format from file(see text/template package document for detail)",
					},
					&cli.StringFlag{
						Name:  "query",
						Usage: "JMESPath query(using when '--output-type' is json only)",
					},
					&cli.StringFlag{
						Name:  "query-file",
						Usage: "JMESPath query from file(using when '--output-type' is json only)",
					},
				},
				Action: func(c *cli.Context) error {

					if err := checkConfigVersion(); err != nil {
						return err
					}
					if err := applyConfigFromFile(c); err != nil {
						return err
					}

					objectstorageListParam.ParamTemplate = c.String("param-template")
					objectstorageListParam.ParamTemplateFile = c.String("param-template-file")
					strInput, err := command.GetParamTemplateValue(objectstorageListParam)
					if err != nil {
						return err
					}
					if strInput != "" {
						p := params.NewListObjectstorageParam()
						err := json.Unmarshal([]byte(strInput), p)
						if err != nil {
							return fmt.Errorf("Failed to parse JSON: %s", err)
						}
						mergo.Merge(objectstorageListParam, p, mergo.WithOverride)
					}

					// Set option values
					if c.IsSet("access-key") || command.IsEmpty(objectstorageListParam.AccessKey) {
						objectstorageListParam.AccessKey = c.String("access-key")
					}
					if c.IsSet("secret-key") || command.IsEmpty(objectstorageListParam.SecretKey) {
						objectstorageListParam.SecretKey = c.String("secret-key")
					}
					if c.IsSet("bucket") || command.IsEmpty(objectstorageListParam.Bucket) {
						objectstorageListParam.Bucket = c.String("bucket")
					}
					if c.IsSet("param-template") {
						objectstorageListParam.ParamTemplate = c.String("param-template")
					}
					if c.IsSet("parameters") {
						objectstorageListParam.Parameters = c.String("parameters")
					}
					if c.IsSet("param-template-file") {
						objectstorageListParam.ParamTemplateFile = c.String("param-template-file")
					}
					if c.IsSet("parameter-file") {
						objectstorageListParam.ParameterFile = c.String("parameter-file")
					}
					if c.IsSet("generate-skeleton") {
						objectstorageListParam.GenerateSkeleton = c.Bool("generate-skeleton")
					}
					if c.IsSet("output-type") {
						objectstorageListParam.OutputType = c.String("output-type")
					}
					if c.IsSet("column") {
						objectstorageListParam.Column = c.StringSlice("column")
					}
					if c.IsSet("quiet") {
						objectstorageListParam.Quiet = c.Bool("quiet")
					}
					if c.IsSet("format") {
						objectstorageListParam.Format = c.String("format")
					}
					if c.IsSet("format-file") {
						objectstorageListParam.FormatFile = c.String("format-file")
					}
					if c.IsSet("query") {
						objectstorageListParam.Query = c.String("query")
					}
					if c.IsSet("query-file") {
						objectstorageListParam.QueryFile = c.String("query-file")
					}

					// Validate global params
					if errors := command.GlobalOption.Validate(true); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					var outputTypeHolder interface{} = objectstorageListParam
					if v, ok := outputTypeHolder.(command.OutputTypeHolder); ok {
						if v.GetOutputType() == "" {
							v.SetOutputType(command.GlobalOption.DefaultOutputType)
						}
					}

					// Experiment warning
					printWarning("")

					// Generate skeleton
					if objectstorageListParam.GenerateSkeleton {
						objectstorageListParam.GenerateSkeleton = false
						objectstorageListParam.FillValueToSkeleton()
						d, err := json.MarshalIndent(objectstorageListParam, "", "\t")
						if err != nil {
							return fmt.Errorf("Failed to Marshal JSON: %s", err)
						}
						fmt.Fprintln(command.GlobalOption.Out, string(d))
						return nil
					}

					// Validate specific for each command params
					if errors := objectstorageListParam.Validate(); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := command.NewContext(c, c.Args().Slice(), objectstorageListParam)

					// Run command with params
					return funcs.ObjectstorageList(ctx, objectstorageListParam)

				},
			},
			{
				Name:      "put",
				Usage:     "Put Objectstorage",
				ArgsUsage: "<local file/directory> <remote path>",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "access-key",
						Usage:   "[Required] set access-key",
						EnvVars: []string{"SACLOUD_OJS_ACCESS_KEY_ID", "AWS_ACCESS_KEY_ID"},
					},
					&cli.StringFlag{
						Name:  "content-type",
						Usage: "set content-type",
						Value: "application/octet-stream",
					},
					&cli.BoolFlag{
						Name:    "recursive",
						Aliases: []string{"r"},
						Usage:   "put objects recursive",
					},
					&cli.StringFlag{
						Name:    "secret-key",
						Usage:   "[Required] set access-key",
						EnvVars: []string{"SACLOUD_OJS_SECRET_ACCESS_KEY", "AWS_SECRET_ACCESS_KEY"},
					},
					&cli.StringFlag{
						Name:    "bucket",
						Usage:   "set bucket",
						EnvVars: []string{"SACLOUD_OJS_BUCKET_NAME"},
					},
					&cli.BoolFlag{
						Name:    "assumeyes",
						Aliases: []string{"y"},
						Usage:   "Assume that the answer to any question which would be asked is yes",
					},
					&cli.StringFlag{
						Name:  "param-template",
						Usage: "Set input parameter from string(JSON)",
					},
					&cli.StringFlag{
						Name:  "parameters",
						Usage: "Set input parameters from JSON string",
					},
					&cli.StringFlag{
						Name:  "param-template-file",
						Usage: "Set input parameter from file",
					},
					&cli.StringFlag{
						Name:  "parameter-file",
						Usage: "Set input parameters from file",
					},
					&cli.BoolFlag{
						Name:  "generate-skeleton",
						Usage: "Output skelton of parameter JSON",
					},
				},
				Action: func(c *cli.Context) error {

					if err := checkConfigVersion(); err != nil {
						return err
					}
					if err := applyConfigFromFile(c); err != nil {
						return err
					}

					objectstoragePutParam.ParamTemplate = c.String("param-template")
					objectstoragePutParam.ParamTemplateFile = c.String("param-template-file")
					strInput, err := command.GetParamTemplateValue(objectstoragePutParam)
					if err != nil {
						return err
					}
					if strInput != "" {
						p := params.NewPutObjectstorageParam()
						err := json.Unmarshal([]byte(strInput), p)
						if err != nil {
							return fmt.Errorf("Failed to parse JSON: %s", err)
						}
						mergo.Merge(objectstoragePutParam, p, mergo.WithOverride)
					}

					// Set option values
					if c.IsSet("access-key") || command.IsEmpty(objectstoragePutParam.AccessKey) {
						objectstoragePutParam.AccessKey = c.String("access-key")
					}
					if c.IsSet("content-type") {
						objectstoragePutParam.ContentType = c.String("content-type")
					}
					if c.IsSet("recursive") {
						objectstoragePutParam.Recursive = c.Bool("recursive")
					}
					if c.IsSet("secret-key") || command.IsEmpty(objectstoragePutParam.SecretKey) {
						objectstoragePutParam.SecretKey = c.String("secret-key")
					}
					if c.IsSet("bucket") || command.IsEmpty(objectstoragePutParam.Bucket) {
						objectstoragePutParam.Bucket = c.String("bucket")
					}
					if c.IsSet("assumeyes") {
						objectstoragePutParam.Assumeyes = c.Bool("assumeyes")
					}
					if c.IsSet("param-template") {
						objectstoragePutParam.ParamTemplate = c.String("param-template")
					}
					if c.IsSet("parameters") {
						objectstoragePutParam.Parameters = c.String("parameters")
					}
					if c.IsSet("param-template-file") {
						objectstoragePutParam.ParamTemplateFile = c.String("param-template-file")
					}
					if c.IsSet("parameter-file") {
						objectstoragePutParam.ParameterFile = c.String("parameter-file")
					}
					if c.IsSet("generate-skeleton") {
						objectstoragePutParam.GenerateSkeleton = c.Bool("generate-skeleton")
					}

					// Validate global params
					if errors := command.GlobalOption.Validate(true); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					var outputTypeHolder interface{} = objectstoragePutParam
					if v, ok := outputTypeHolder.(command.OutputTypeHolder); ok {
						if v.GetOutputType() == "" {
							v.SetOutputType(command.GlobalOption.DefaultOutputType)
						}
					}

					// Experiment warning
					printWarning("")

					// Generate skeleton
					if objectstoragePutParam.GenerateSkeleton {
						objectstoragePutParam.GenerateSkeleton = false
						objectstoragePutParam.FillValueToSkeleton()
						d, err := json.MarshalIndent(objectstoragePutParam, "", "\t")
						if err != nil {
							return fmt.Errorf("Failed to Marshal JSON: %s", err)
						}
						fmt.Fprintln(command.GlobalOption.Out, string(d))
						return nil
					}

					// Validate specific for each command params
					if errors := objectstoragePutParam.Validate(); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := command.NewContext(c, c.Args().Slice(), objectstoragePutParam)

					// confirm
					if !objectstoragePutParam.Assumeyes {
						if !isTerminal() {
							return fmt.Errorf("When using redirect/pipe, specify --assumeyes(-y) option")
						}
						if !command.ConfirmContinue("put") {
							return nil
						}
					}

					// Run command with params
					return funcs.ObjectstoragePut(ctx, objectstoragePutParam)

				},
			},
			{
				Name:      "get",
				Usage:     "Get Objectstorage",
				ArgsUsage: "<remote path> <local file/directory>",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "access-key",
						Usage:   "[Required] set access-key",
						EnvVars: []string{"SACLOUD_OJS_ACCESS_KEY_ID", "AWS_ACCESS_KEY_ID"},
					},
					&cli.BoolFlag{
						Name:    "recursive",
						Aliases: []string{"r"},
						Usage:   "get objects recursive",
					},
					&cli.StringFlag{
						Name:    "secret-key",
						Usage:   "[Required] set access-key",
						EnvVars: []string{"SACLOUD_OJS_SECRET_ACCESS_KEY", "AWS_SECRET_ACCESS_KEY"},
					},
					&cli.StringFlag{
						Name:    "bucket",
						Usage:   "set bucket",
						EnvVars: []string{"SACLOUD_OJS_BUCKET_NAME"},
					},
					&cli.StringFlag{
						Name:  "param-template",
						Usage: "Set input parameter from string(JSON)",
					},
					&cli.StringFlag{
						Name:  "parameters",
						Usage: "Set input parameters from JSON string",
					},
					&cli.StringFlag{
						Name:  "param-template-file",
						Usage: "Set input parameter from file",
					},
					&cli.StringFlag{
						Name:  "parameter-file",
						Usage: "Set input parameters from file",
					},
					&cli.BoolFlag{
						Name:  "generate-skeleton",
						Usage: "Output skelton of parameter JSON",
					},
				},
				Action: func(c *cli.Context) error {

					if err := checkConfigVersion(); err != nil {
						return err
					}
					if err := applyConfigFromFile(c); err != nil {
						return err
					}

					objectstorageGetParam.ParamTemplate = c.String("param-template")
					objectstorageGetParam.ParamTemplateFile = c.String("param-template-file")
					strInput, err := command.GetParamTemplateValue(objectstorageGetParam)
					if err != nil {
						return err
					}
					if strInput != "" {
						p := params.NewGetObjectstorageParam()
						err := json.Unmarshal([]byte(strInput), p)
						if err != nil {
							return fmt.Errorf("Failed to parse JSON: %s", err)
						}
						mergo.Merge(objectstorageGetParam, p, mergo.WithOverride)
					}

					// Set option values
					if c.IsSet("access-key") || command.IsEmpty(objectstorageGetParam.AccessKey) {
						objectstorageGetParam.AccessKey = c.String("access-key")
					}
					if c.IsSet("recursive") {
						objectstorageGetParam.Recursive = c.Bool("recursive")
					}
					if c.IsSet("secret-key") || command.IsEmpty(objectstorageGetParam.SecretKey) {
						objectstorageGetParam.SecretKey = c.String("secret-key")
					}
					if c.IsSet("bucket") || command.IsEmpty(objectstorageGetParam.Bucket) {
						objectstorageGetParam.Bucket = c.String("bucket")
					}
					if c.IsSet("param-template") {
						objectstorageGetParam.ParamTemplate = c.String("param-template")
					}
					if c.IsSet("parameters") {
						objectstorageGetParam.Parameters = c.String("parameters")
					}
					if c.IsSet("param-template-file") {
						objectstorageGetParam.ParamTemplateFile = c.String("param-template-file")
					}
					if c.IsSet("parameter-file") {
						objectstorageGetParam.ParameterFile = c.String("parameter-file")
					}
					if c.IsSet("generate-skeleton") {
						objectstorageGetParam.GenerateSkeleton = c.Bool("generate-skeleton")
					}

					// Validate global params
					if errors := command.GlobalOption.Validate(true); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					var outputTypeHolder interface{} = objectstorageGetParam
					if v, ok := outputTypeHolder.(command.OutputTypeHolder); ok {
						if v.GetOutputType() == "" {
							v.SetOutputType(command.GlobalOption.DefaultOutputType)
						}
					}

					// Experiment warning
					printWarning("")

					// Generate skeleton
					if objectstorageGetParam.GenerateSkeleton {
						objectstorageGetParam.GenerateSkeleton = false
						objectstorageGetParam.FillValueToSkeleton()
						d, err := json.MarshalIndent(objectstorageGetParam, "", "\t")
						if err != nil {
							return fmt.Errorf("Failed to Marshal JSON: %s", err)
						}
						fmt.Fprintln(command.GlobalOption.Out, string(d))
						return nil
					}

					// Validate specific for each command params
					if errors := objectstorageGetParam.Validate(); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := command.NewContext(c, c.Args().Slice(), objectstorageGetParam)

					// Run command with params
					return funcs.ObjectstorageGet(ctx, objectstorageGetParam)

				},
			},
			{
				Name:      "delete",
				Aliases:   []string{"rm", "del"},
				Usage:     "Delete Objectstorage",
				ArgsUsage: "<remote path>",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "access-key",
						Usage:   "[Required] set access-key",
						EnvVars: []string{"SACLOUD_OJS_ACCESS_KEY_ID", "AWS_ACCESS_KEY_ID"},
					},
					&cli.BoolFlag{
						Name:    "recursive",
						Aliases: []string{"r"},
						Usage:   "delete objects recursive",
					},
					&cli.StringFlag{
						Name:    "secret-key",
						Usage:   "[Required] set access-key",
						EnvVars: []string{"SACLOUD_OJS_SECRET_ACCESS_KEY", "AWS_SECRET_ACCESS_KEY"},
					},
					&cli.StringFlag{
						Name:    "bucket",
						Usage:   "set bucket",
						EnvVars: []string{"SACLOUD_OJS_BUCKET_NAME"},
					},
					&cli.BoolFlag{
						Name:    "assumeyes",
						Aliases: []string{"y"},
						Usage:   "Assume that the answer to any question which would be asked is yes",
					},
					&cli.StringFlag{
						Name:  "param-template",
						Usage: "Set input parameter from string(JSON)",
					},
					&cli.StringFlag{
						Name:  "parameters",
						Usage: "Set input parameters from JSON string",
					},
					&cli.StringFlag{
						Name:  "param-template-file",
						Usage: "Set input parameter from file",
					},
					&cli.StringFlag{
						Name:  "parameter-file",
						Usage: "Set input parameters from file",
					},
					&cli.BoolFlag{
						Name:  "generate-skeleton",
						Usage: "Output skelton of parameter JSON",
					},
				},
				Action: func(c *cli.Context) error {

					if err := checkConfigVersion(); err != nil {
						return err
					}
					if err := applyConfigFromFile(c); err != nil {
						return err
					}

					objectstorageDeleteParam.ParamTemplate = c.String("param-template")
					objectstorageDeleteParam.ParamTemplateFile = c.String("param-template-file")
					strInput, err := command.GetParamTemplateValue(objectstorageDeleteParam)
					if err != nil {
						return err
					}
					if strInput != "" {
						p := params.NewDeleteObjectstorageParam()
						err := json.Unmarshal([]byte(strInput), p)
						if err != nil {
							return fmt.Errorf("Failed to parse JSON: %s", err)
						}
						mergo.Merge(objectstorageDeleteParam, p, mergo.WithOverride)
					}

					// Set option values
					if c.IsSet("access-key") || command.IsEmpty(objectstorageDeleteParam.AccessKey) {
						objectstorageDeleteParam.AccessKey = c.String("access-key")
					}
					if c.IsSet("recursive") {
						objectstorageDeleteParam.Recursive = c.Bool("recursive")
					}
					if c.IsSet("secret-key") || command.IsEmpty(objectstorageDeleteParam.SecretKey) {
						objectstorageDeleteParam.SecretKey = c.String("secret-key")
					}
					if c.IsSet("bucket") || command.IsEmpty(objectstorageDeleteParam.Bucket) {
						objectstorageDeleteParam.Bucket = c.String("bucket")
					}
					if c.IsSet("assumeyes") {
						objectstorageDeleteParam.Assumeyes = c.Bool("assumeyes")
					}
					if c.IsSet("param-template") {
						objectstorageDeleteParam.ParamTemplate = c.String("param-template")
					}
					if c.IsSet("parameters") {
						objectstorageDeleteParam.Parameters = c.String("parameters")
					}
					if c.IsSet("param-template-file") {
						objectstorageDeleteParam.ParamTemplateFile = c.String("param-template-file")
					}
					if c.IsSet("parameter-file") {
						objectstorageDeleteParam.ParameterFile = c.String("parameter-file")
					}
					if c.IsSet("generate-skeleton") {
						objectstorageDeleteParam.GenerateSkeleton = c.Bool("generate-skeleton")
					}

					// Validate global params
					if errors := command.GlobalOption.Validate(true); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					var outputTypeHolder interface{} = objectstorageDeleteParam
					if v, ok := outputTypeHolder.(command.OutputTypeHolder); ok {
						if v.GetOutputType() == "" {
							v.SetOutputType(command.GlobalOption.DefaultOutputType)
						}
					}

					// Experiment warning
					printWarning("")

					// Generate skeleton
					if objectstorageDeleteParam.GenerateSkeleton {
						objectstorageDeleteParam.GenerateSkeleton = false
						objectstorageDeleteParam.FillValueToSkeleton()
						d, err := json.MarshalIndent(objectstorageDeleteParam, "", "\t")
						if err != nil {
							return fmt.Errorf("Failed to Marshal JSON: %s", err)
						}
						fmt.Fprintln(command.GlobalOption.Out, string(d))
						return nil
					}

					// Validate specific for each command params
					if errors := objectstorageDeleteParam.Validate(); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := command.NewContext(c, c.Args().Slice(), objectstorageDeleteParam)

					// confirm
					if !objectstorageDeleteParam.Assumeyes {
						if !isTerminal() {
							return fmt.Errorf("When using redirect/pipe, specify --assumeyes(-y) option")
						}
						if !command.ConfirmContinue("delete") {
							return nil
						}
					}

					// Run command with params
					return funcs.ObjectstorageDelete(ctx, objectstorageDeleteParam)

				},
			},
		},
	}

	// build Category-Resource mapping
	AppendResourceCategoryMap("object-storage", &schema.Category{
		Key:         "saas",
		DisplayName: "Other services",
		Order:       80,
	})

	// build Category-Command mapping

	AppendCommandCategoryMap("object-storage", "delete", &schema.Category{
		Key:         "basics",
		DisplayName: "Basics",
		Order:       1,
	})
	AppendCommandCategoryMap("object-storage", "get", &schema.Category{
		Key:         "basics",
		DisplayName: "Basics",
		Order:       1,
	})
	AppendCommandCategoryMap("object-storage", "list", &schema.Category{
		Key:         "basics",
		DisplayName: "Basics",
		Order:       1,
	})
	AppendCommandCategoryMap("object-storage", "put", &schema.Category{
		Key:         "basics",
		DisplayName: "Basics",
		Order:       1,
	})

	// build Category-Param mapping

	AppendFlagCategoryMap("object-storage", "delete", "access-key", &schema.Category{
		Key:         "auth",
		DisplayName: "Auth options",
		Order:       1,
	})
	AppendFlagCategoryMap("object-storage", "delete", "assumeyes", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("object-storage", "delete", "bucket", &schema.Category{
		Key:         "auth",
		DisplayName: "Auth options",
		Order:       1,
	})
	AppendFlagCategoryMap("object-storage", "delete", "generate-skeleton", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("object-storage", "delete", "param-template", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("object-storage", "delete", "param-template-file", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("object-storage", "delete", "parameter-file", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("object-storage", "delete", "parameters", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("object-storage", "delete", "recursive", &schema.Category{
		Key:         "operation",
		DisplayName: "Operation options",
		Order:       1,
	})
	AppendFlagCategoryMap("object-storage", "delete", "secret-key", &schema.Category{
		Key:         "auth",
		DisplayName: "Auth options",
		Order:       1,
	})
	AppendFlagCategoryMap("object-storage", "get", "access-key", &schema.Category{
		Key:         "auth",
		DisplayName: "Auth options",
		Order:       1,
	})
	AppendFlagCategoryMap("object-storage", "get", "bucket", &schema.Category{
		Key:         "auth",
		DisplayName: "Auth options",
		Order:       1,
	})
	AppendFlagCategoryMap("object-storage", "get", "generate-skeleton", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("object-storage", "get", "param-template", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("object-storage", "get", "param-template-file", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("object-storage", "get", "parameter-file", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("object-storage", "get", "parameters", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("object-storage", "get", "recursive", &schema.Category{
		Key:         "operation",
		DisplayName: "Operation options",
		Order:       1,
	})
	AppendFlagCategoryMap("object-storage", "get", "secret-key", &schema.Category{
		Key:         "auth",
		DisplayName: "Auth options",
		Order:       1,
	})
	AppendFlagCategoryMap("object-storage", "list", "access-key", &schema.Category{
		Key:         "auth",
		DisplayName: "Auth options",
		Order:       1,
	})
	AppendFlagCategoryMap("object-storage", "list", "bucket", &schema.Category{
		Key:         "auth",
		DisplayName: "Auth options",
		Order:       1,
	})
	AppendFlagCategoryMap("object-storage", "list", "column", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("object-storage", "list", "format", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("object-storage", "list", "format-file", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("object-storage", "list", "generate-skeleton", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("object-storage", "list", "output-type", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("object-storage", "list", "param-template", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("object-storage", "list", "param-template-file", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("object-storage", "list", "parameter-file", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("object-storage", "list", "parameters", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("object-storage", "list", "query", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("object-storage", "list", "query-file", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("object-storage", "list", "quiet", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("object-storage", "list", "secret-key", &schema.Category{
		Key:         "auth",
		DisplayName: "Auth options",
		Order:       1,
	})
	AppendFlagCategoryMap("object-storage", "put", "access-key", &schema.Category{
		Key:         "auth",
		DisplayName: "Auth options",
		Order:       1,
	})
	AppendFlagCategoryMap("object-storage", "put", "assumeyes", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("object-storage", "put", "bucket", &schema.Category{
		Key:         "auth",
		DisplayName: "Auth options",
		Order:       1,
	})
	AppendFlagCategoryMap("object-storage", "put", "content-type", &schema.Category{
		Key:         "operation",
		DisplayName: "Operation options",
		Order:       1,
	})
	AppendFlagCategoryMap("object-storage", "put", "generate-skeleton", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("object-storage", "put", "param-template", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("object-storage", "put", "param-template-file", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("object-storage", "put", "parameter-file", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("object-storage", "put", "parameters", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("object-storage", "put", "recursive", &schema.Category{
		Key:         "operation",
		DisplayName: "Operation options",
		Order:       1,
	})
	AppendFlagCategoryMap("object-storage", "put", "secret-key", &schema.Category{
		Key:         "auth",
		DisplayName: "Auth options",
		Order:       1,
	})

	// append command to GlobalContext
	Commands = append(Commands, cliCommand)
}
