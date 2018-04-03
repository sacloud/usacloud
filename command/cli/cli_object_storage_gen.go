// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-commands'; DO NOT EDIT

package cli

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/imdario/mergo"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/completion"
	"github.com/sacloud/usacloud/command/funcs"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/schema"
	"gopkg.in/urfave/cli.v2"
)

func init() {
	listParam := params.NewListObjectStorageParam()
	putParam := params.NewPutObjectStorageParam()
	getParam := params.NewGetObjectStorageParam()
	deleteParam := params.NewDeleteObjectStorageParam()

	cliCommand := &cli.Command{
		Name:    "object-storage",
		Aliases: []string{"ojs"},
		Usage:   "A manage commands of ObjectStorage",
		Subcommands: []*cli.Command{
			{
				Name:      "list",
				Aliases:   []string{"ls"},
				Usage:     "List ObjectStorage",
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
						Name:  "param-template-file",
						Usage: "Set input parameter from file",
					},
					&cli.BoolFlag{
						Name:  "generate-skeleton",
						Usage: "Output skelton of parameter JSON",
					},
					&cli.StringFlag{
						Name:    "output-type",
						Aliases: []string{"out"},
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
				},
				ShellComplete: func(c *cli.Context) {

					if c.NArg() < 3 { // invalid args
						return
					}

					if err := checkConfigVersion(); err != nil {
						return
					}
					if err := applyConfigFromFile(c); err != nil {
						return
					}

					// c.Args() == arg1 arg2 arg3 -- [cur] [prev] [commandName]
					args := c.Args().Slice()
					commandName := args[c.NArg()-1]
					prev := args[c.NArg()-2]
					cur := args[c.NArg()-3]

					// set real args
					realArgs := args[0 : c.NArg()-3]

					// Validate global params
					command.GlobalOption.Validate(false)

					// set default output-type
					// when params have output-type option and have empty value
					var outputTypeHolder interface{} = listParam
					if v, ok := outputTypeHolder.(command.OutputTypeHolder); ok {
						if v.GetOutputType() == "" {
							v.SetOutputType(command.GlobalOption.DefaultOutputType)
						}
					}

					// build command context
					ctx := command.NewContext(c, realArgs, listParam)

					// Set option values
					if c.IsSet("access-key") || command.IsEmpty(listParam.AccessKey) {
						listParam.AccessKey = c.String("access-key")
					}
					if c.IsSet("secret-key") || command.IsEmpty(listParam.SecretKey) {
						listParam.SecretKey = c.String("secret-key")
					}
					if c.IsSet("bucket") || command.IsEmpty(listParam.Bucket) {
						listParam.Bucket = c.String("bucket")
					}
					if c.IsSet("param-template") {
						listParam.ParamTemplate = c.String("param-template")
					}
					if c.IsSet("param-template-file") {
						listParam.ParamTemplateFile = c.String("param-template-file")
					}
					if c.IsSet("generate-skeleton") {
						listParam.GenerateSkeleton = c.Bool("generate-skeleton")
					}
					if c.IsSet("output-type") {
						listParam.OutputType = c.String("output-type")
					}
					if c.IsSet("column") {
						listParam.Column = c.StringSlice("column")
					}
					if c.IsSet("quiet") {
						listParam.Quiet = c.Bool("quiet")
					}
					if c.IsSet("format") {
						listParam.Format = c.String("format")
					}
					if c.IsSet("format-file") {
						listParam.FormatFile = c.String("format-file")
					}

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completion.FlagNames(c, commandName)
								return
							} else {
								completion.ObjectStorageListCompleteArgs(ctx, listParam, cur, prev, commandName)
								return
							}
						}

						// cleanup flag name
						name := prev
						for {
							if !strings.HasPrefix(name, "-") {
								break
							}
							name = strings.Replace(name, "-", "", 1)
						}

						// flag is exists? , is BoolFlag?
						exists := false
						for _, flag := range c.App.Command(commandName).Flags {

							for _, n := range flag.Names() {
								if n == name {
									exists = true
									break
								}
							}

							if exists {
								if _, ok := flag.(*cli.BoolFlag); ok {
									if strings.HasPrefix(cur, "-") {
										completion.FlagNames(c, commandName)
										return
									} else {
										completion.ObjectStorageListCompleteArgs(ctx, listParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									completion.ObjectStorageListCompleteFlags(ctx, listParam, name, cur)
									return
								}
							}
						}
						// here, prev is wrong, so noop.
					} else {
						if strings.HasPrefix(cur, "-") {
							completion.FlagNames(c, commandName)
							return
						} else {
							completion.ObjectStorageListCompleteArgs(ctx, listParam, cur, prev, commandName)
							return
						}
					}
				},
				Action: func(c *cli.Context) error {

					if err := checkConfigVersion(); err != nil {
						return err
					}
					if err := applyConfigFromFile(c); err != nil {
						return err
					}

					listParam.ParamTemplate = c.String("param-template")
					listParam.ParamTemplateFile = c.String("param-template-file")
					strInput, err := command.GetParamTemplateValue(listParam)
					if err != nil {
						return err
					}
					if strInput != "" {
						p := params.NewListObjectStorageParam()
						err := json.Unmarshal([]byte(strInput), p)
						if err != nil {
							return fmt.Errorf("Failed to parse JSON: %s", err)
						}
						mergo.MergeWithOverwrite(listParam, p)
					}

					// Set option values
					if c.IsSet("access-key") || command.IsEmpty(listParam.AccessKey) {
						listParam.AccessKey = c.String("access-key")
					}
					if c.IsSet("secret-key") || command.IsEmpty(listParam.SecretKey) {
						listParam.SecretKey = c.String("secret-key")
					}
					if c.IsSet("bucket") || command.IsEmpty(listParam.Bucket) {
						listParam.Bucket = c.String("bucket")
					}
					if c.IsSet("param-template") {
						listParam.ParamTemplate = c.String("param-template")
					}
					if c.IsSet("param-template-file") {
						listParam.ParamTemplateFile = c.String("param-template-file")
					}
					if c.IsSet("generate-skeleton") {
						listParam.GenerateSkeleton = c.Bool("generate-skeleton")
					}
					if c.IsSet("output-type") {
						listParam.OutputType = c.String("output-type")
					}
					if c.IsSet("column") {
						listParam.Column = c.StringSlice("column")
					}
					if c.IsSet("quiet") {
						listParam.Quiet = c.Bool("quiet")
					}
					if c.IsSet("format") {
						listParam.Format = c.String("format")
					}
					if c.IsSet("format-file") {
						listParam.FormatFile = c.String("format-file")
					}

					// Validate global params
					if errors := command.GlobalOption.Validate(true); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					var outputTypeHolder interface{} = listParam
					if v, ok := outputTypeHolder.(command.OutputTypeHolder); ok {
						if v.GetOutputType() == "" {
							v.SetOutputType(command.GlobalOption.DefaultOutputType)
						}
					}

					// Generate skeleton
					if listParam.GenerateSkeleton {
						listParam.GenerateSkeleton = false
						listParam.FillValueToSkeleton()
						d, err := json.MarshalIndent(listParam, "", "\t")
						if err != nil {
							return fmt.Errorf("Failed to Marshal JSON: %s", err)
						}
						fmt.Fprintln(command.GlobalOption.Out, string(d))
						return nil
					}

					// Validate specific for each command params
					if errors := listParam.Validate(); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := command.NewContext(c, c.Args().Slice(), listParam)

					// Run command with params
					return funcs.ObjectStorageList(ctx, listParam)

				},
			},
			{
				Name:      "put",
				Usage:     "Put ObjectStorage",
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
						Name:  "param-template-file",
						Usage: "Set input parameter from file",
					},
					&cli.BoolFlag{
						Name:  "generate-skeleton",
						Usage: "Output skelton of parameter JSON",
					},
				},
				ShellComplete: func(c *cli.Context) {

					if c.NArg() < 3 { // invalid args
						return
					}

					if err := checkConfigVersion(); err != nil {
						return
					}
					if err := applyConfigFromFile(c); err != nil {
						return
					}

					// c.Args() == arg1 arg2 arg3 -- [cur] [prev] [commandName]
					args := c.Args().Slice()
					commandName := args[c.NArg()-1]
					prev := args[c.NArg()-2]
					cur := args[c.NArg()-3]

					// set real args
					realArgs := args[0 : c.NArg()-3]

					// Validate global params
					command.GlobalOption.Validate(false)

					// set default output-type
					// when params have output-type option and have empty value
					var outputTypeHolder interface{} = putParam
					if v, ok := outputTypeHolder.(command.OutputTypeHolder); ok {
						if v.GetOutputType() == "" {
							v.SetOutputType(command.GlobalOption.DefaultOutputType)
						}
					}

					// build command context
					ctx := command.NewContext(c, realArgs, putParam)

					// Set option values
					if c.IsSet("access-key") || command.IsEmpty(putParam.AccessKey) {
						putParam.AccessKey = c.String("access-key")
					}
					if c.IsSet("content-type") {
						putParam.ContentType = c.String("content-type")
					}
					if c.IsSet("recursive") {
						putParam.Recursive = c.Bool("recursive")
					}
					if c.IsSet("secret-key") || command.IsEmpty(putParam.SecretKey) {
						putParam.SecretKey = c.String("secret-key")
					}
					if c.IsSet("bucket") || command.IsEmpty(putParam.Bucket) {
						putParam.Bucket = c.String("bucket")
					}
					if c.IsSet("assumeyes") {
						putParam.Assumeyes = c.Bool("assumeyes")
					}
					if c.IsSet("param-template") {
						putParam.ParamTemplate = c.String("param-template")
					}
					if c.IsSet("param-template-file") {
						putParam.ParamTemplateFile = c.String("param-template-file")
					}
					if c.IsSet("generate-skeleton") {
						putParam.GenerateSkeleton = c.Bool("generate-skeleton")
					}

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completion.FlagNames(c, commandName)
								return
							} else {
								completion.ObjectStoragePutCompleteArgs(ctx, putParam, cur, prev, commandName)
								return
							}
						}

						// cleanup flag name
						name := prev
						for {
							if !strings.HasPrefix(name, "-") {
								break
							}
							name = strings.Replace(name, "-", "", 1)
						}

						// flag is exists? , is BoolFlag?
						exists := false
						for _, flag := range c.App.Command(commandName).Flags {

							for _, n := range flag.Names() {
								if n == name {
									exists = true
									break
								}
							}

							if exists {
								if _, ok := flag.(*cli.BoolFlag); ok {
									if strings.HasPrefix(cur, "-") {
										completion.FlagNames(c, commandName)
										return
									} else {
										completion.ObjectStoragePutCompleteArgs(ctx, putParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									completion.ObjectStoragePutCompleteFlags(ctx, putParam, name, cur)
									return
								}
							}
						}
						// here, prev is wrong, so noop.
					} else {
						if strings.HasPrefix(cur, "-") {
							completion.FlagNames(c, commandName)
							return
						} else {
							completion.ObjectStoragePutCompleteArgs(ctx, putParam, cur, prev, commandName)
							return
						}
					}
				},
				Action: func(c *cli.Context) error {

					if err := checkConfigVersion(); err != nil {
						return err
					}
					if err := applyConfigFromFile(c); err != nil {
						return err
					}

					putParam.ParamTemplate = c.String("param-template")
					putParam.ParamTemplateFile = c.String("param-template-file")
					strInput, err := command.GetParamTemplateValue(putParam)
					if err != nil {
						return err
					}
					if strInput != "" {
						p := params.NewPutObjectStorageParam()
						err := json.Unmarshal([]byte(strInput), p)
						if err != nil {
							return fmt.Errorf("Failed to parse JSON: %s", err)
						}
						mergo.MergeWithOverwrite(putParam, p)
					}

					// Set option values
					if c.IsSet("access-key") || command.IsEmpty(putParam.AccessKey) {
						putParam.AccessKey = c.String("access-key")
					}
					if c.IsSet("content-type") {
						putParam.ContentType = c.String("content-type")
					}
					if c.IsSet("recursive") {
						putParam.Recursive = c.Bool("recursive")
					}
					if c.IsSet("secret-key") || command.IsEmpty(putParam.SecretKey) {
						putParam.SecretKey = c.String("secret-key")
					}
					if c.IsSet("bucket") || command.IsEmpty(putParam.Bucket) {
						putParam.Bucket = c.String("bucket")
					}
					if c.IsSet("assumeyes") {
						putParam.Assumeyes = c.Bool("assumeyes")
					}
					if c.IsSet("param-template") {
						putParam.ParamTemplate = c.String("param-template")
					}
					if c.IsSet("param-template-file") {
						putParam.ParamTemplateFile = c.String("param-template-file")
					}
					if c.IsSet("generate-skeleton") {
						putParam.GenerateSkeleton = c.Bool("generate-skeleton")
					}

					// Validate global params
					if errors := command.GlobalOption.Validate(true); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					var outputTypeHolder interface{} = putParam
					if v, ok := outputTypeHolder.(command.OutputTypeHolder); ok {
						if v.GetOutputType() == "" {
							v.SetOutputType(command.GlobalOption.DefaultOutputType)
						}
					}

					// Generate skeleton
					if putParam.GenerateSkeleton {
						putParam.GenerateSkeleton = false
						putParam.FillValueToSkeleton()
						d, err := json.MarshalIndent(putParam, "", "\t")
						if err != nil {
							return fmt.Errorf("Failed to Marshal JSON: %s", err)
						}
						fmt.Fprintln(command.GlobalOption.Out, string(d))
						return nil
					}

					// Validate specific for each command params
					if errors := putParam.Validate(); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := command.NewContext(c, c.Args().Slice(), putParam)

					// confirm
					if !putParam.Assumeyes {
						if !isTerminal() {
							return fmt.Errorf("When using redirect/pipe, specify --assumeyes(-y) option")
						}
						if !command.ConfirmContinue("put") {
							return nil
						}
					}

					// Run command with params
					return funcs.ObjectStoragePut(ctx, putParam)

				},
			},
			{
				Name:      "get",
				Usage:     "Get ObjectStorage",
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
						Name:  "param-template-file",
						Usage: "Set input parameter from file",
					},
					&cli.BoolFlag{
						Name:  "generate-skeleton",
						Usage: "Output skelton of parameter JSON",
					},
				},
				ShellComplete: func(c *cli.Context) {

					if c.NArg() < 3 { // invalid args
						return
					}

					if err := checkConfigVersion(); err != nil {
						return
					}
					if err := applyConfigFromFile(c); err != nil {
						return
					}

					// c.Args() == arg1 arg2 arg3 -- [cur] [prev] [commandName]
					args := c.Args().Slice()
					commandName := args[c.NArg()-1]
					prev := args[c.NArg()-2]
					cur := args[c.NArg()-3]

					// set real args
					realArgs := args[0 : c.NArg()-3]

					// Validate global params
					command.GlobalOption.Validate(false)

					// set default output-type
					// when params have output-type option and have empty value
					var outputTypeHolder interface{} = getParam
					if v, ok := outputTypeHolder.(command.OutputTypeHolder); ok {
						if v.GetOutputType() == "" {
							v.SetOutputType(command.GlobalOption.DefaultOutputType)
						}
					}

					// build command context
					ctx := command.NewContext(c, realArgs, getParam)

					// Set option values
					if c.IsSet("access-key") || command.IsEmpty(getParam.AccessKey) {
						getParam.AccessKey = c.String("access-key")
					}
					if c.IsSet("recursive") {
						getParam.Recursive = c.Bool("recursive")
					}
					if c.IsSet("secret-key") || command.IsEmpty(getParam.SecretKey) {
						getParam.SecretKey = c.String("secret-key")
					}
					if c.IsSet("bucket") || command.IsEmpty(getParam.Bucket) {
						getParam.Bucket = c.String("bucket")
					}
					if c.IsSet("param-template") {
						getParam.ParamTemplate = c.String("param-template")
					}
					if c.IsSet("param-template-file") {
						getParam.ParamTemplateFile = c.String("param-template-file")
					}
					if c.IsSet("generate-skeleton") {
						getParam.GenerateSkeleton = c.Bool("generate-skeleton")
					}

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completion.FlagNames(c, commandName)
								return
							} else {
								completion.ObjectStorageGetCompleteArgs(ctx, getParam, cur, prev, commandName)
								return
							}
						}

						// cleanup flag name
						name := prev
						for {
							if !strings.HasPrefix(name, "-") {
								break
							}
							name = strings.Replace(name, "-", "", 1)
						}

						// flag is exists? , is BoolFlag?
						exists := false
						for _, flag := range c.App.Command(commandName).Flags {

							for _, n := range flag.Names() {
								if n == name {
									exists = true
									break
								}
							}

							if exists {
								if _, ok := flag.(*cli.BoolFlag); ok {
									if strings.HasPrefix(cur, "-") {
										completion.FlagNames(c, commandName)
										return
									} else {
										completion.ObjectStorageGetCompleteArgs(ctx, getParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									completion.ObjectStorageGetCompleteFlags(ctx, getParam, name, cur)
									return
								}
							}
						}
						// here, prev is wrong, so noop.
					} else {
						if strings.HasPrefix(cur, "-") {
							completion.FlagNames(c, commandName)
							return
						} else {
							completion.ObjectStorageGetCompleteArgs(ctx, getParam, cur, prev, commandName)
							return
						}
					}
				},
				Action: func(c *cli.Context) error {

					if err := checkConfigVersion(); err != nil {
						return err
					}
					if err := applyConfigFromFile(c); err != nil {
						return err
					}

					getParam.ParamTemplate = c.String("param-template")
					getParam.ParamTemplateFile = c.String("param-template-file")
					strInput, err := command.GetParamTemplateValue(getParam)
					if err != nil {
						return err
					}
					if strInput != "" {
						p := params.NewGetObjectStorageParam()
						err := json.Unmarshal([]byte(strInput), p)
						if err != nil {
							return fmt.Errorf("Failed to parse JSON: %s", err)
						}
						mergo.MergeWithOverwrite(getParam, p)
					}

					// Set option values
					if c.IsSet("access-key") || command.IsEmpty(getParam.AccessKey) {
						getParam.AccessKey = c.String("access-key")
					}
					if c.IsSet("recursive") {
						getParam.Recursive = c.Bool("recursive")
					}
					if c.IsSet("secret-key") || command.IsEmpty(getParam.SecretKey) {
						getParam.SecretKey = c.String("secret-key")
					}
					if c.IsSet("bucket") || command.IsEmpty(getParam.Bucket) {
						getParam.Bucket = c.String("bucket")
					}
					if c.IsSet("param-template") {
						getParam.ParamTemplate = c.String("param-template")
					}
					if c.IsSet("param-template-file") {
						getParam.ParamTemplateFile = c.String("param-template-file")
					}
					if c.IsSet("generate-skeleton") {
						getParam.GenerateSkeleton = c.Bool("generate-skeleton")
					}

					// Validate global params
					if errors := command.GlobalOption.Validate(true); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					var outputTypeHolder interface{} = getParam
					if v, ok := outputTypeHolder.(command.OutputTypeHolder); ok {
						if v.GetOutputType() == "" {
							v.SetOutputType(command.GlobalOption.DefaultOutputType)
						}
					}

					// Generate skeleton
					if getParam.GenerateSkeleton {
						getParam.GenerateSkeleton = false
						getParam.FillValueToSkeleton()
						d, err := json.MarshalIndent(getParam, "", "\t")
						if err != nil {
							return fmt.Errorf("Failed to Marshal JSON: %s", err)
						}
						fmt.Fprintln(command.GlobalOption.Out, string(d))
						return nil
					}

					// Validate specific for each command params
					if errors := getParam.Validate(); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := command.NewContext(c, c.Args().Slice(), getParam)

					// Run command with params
					return funcs.ObjectStorageGet(ctx, getParam)

				},
			},
			{
				Name:      "delete",
				Aliases:   []string{"rm", "del"},
				Usage:     "Delete ObjectStorage",
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
						Name:  "param-template-file",
						Usage: "Set input parameter from file",
					},
					&cli.BoolFlag{
						Name:  "generate-skeleton",
						Usage: "Output skelton of parameter JSON",
					},
				},
				ShellComplete: func(c *cli.Context) {

					if c.NArg() < 3 { // invalid args
						return
					}

					if err := checkConfigVersion(); err != nil {
						return
					}
					if err := applyConfigFromFile(c); err != nil {
						return
					}

					// c.Args() == arg1 arg2 arg3 -- [cur] [prev] [commandName]
					args := c.Args().Slice()
					commandName := args[c.NArg()-1]
					prev := args[c.NArg()-2]
					cur := args[c.NArg()-3]

					// set real args
					realArgs := args[0 : c.NArg()-3]

					// Validate global params
					command.GlobalOption.Validate(false)

					// set default output-type
					// when params have output-type option and have empty value
					var outputTypeHolder interface{} = deleteParam
					if v, ok := outputTypeHolder.(command.OutputTypeHolder); ok {
						if v.GetOutputType() == "" {
							v.SetOutputType(command.GlobalOption.DefaultOutputType)
						}
					}

					// build command context
					ctx := command.NewContext(c, realArgs, deleteParam)

					// Set option values
					if c.IsSet("access-key") || command.IsEmpty(deleteParam.AccessKey) {
						deleteParam.AccessKey = c.String("access-key")
					}
					if c.IsSet("recursive") {
						deleteParam.Recursive = c.Bool("recursive")
					}
					if c.IsSet("secret-key") || command.IsEmpty(deleteParam.SecretKey) {
						deleteParam.SecretKey = c.String("secret-key")
					}
					if c.IsSet("bucket") || command.IsEmpty(deleteParam.Bucket) {
						deleteParam.Bucket = c.String("bucket")
					}
					if c.IsSet("assumeyes") {
						deleteParam.Assumeyes = c.Bool("assumeyes")
					}
					if c.IsSet("param-template") {
						deleteParam.ParamTemplate = c.String("param-template")
					}
					if c.IsSet("param-template-file") {
						deleteParam.ParamTemplateFile = c.String("param-template-file")
					}
					if c.IsSet("generate-skeleton") {
						deleteParam.GenerateSkeleton = c.Bool("generate-skeleton")
					}

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completion.FlagNames(c, commandName)
								return
							} else {
								completion.ObjectStorageDeleteCompleteArgs(ctx, deleteParam, cur, prev, commandName)
								return
							}
						}

						// cleanup flag name
						name := prev
						for {
							if !strings.HasPrefix(name, "-") {
								break
							}
							name = strings.Replace(name, "-", "", 1)
						}

						// flag is exists? , is BoolFlag?
						exists := false
						for _, flag := range c.App.Command(commandName).Flags {

							for _, n := range flag.Names() {
								if n == name {
									exists = true
									break
								}
							}

							if exists {
								if _, ok := flag.(*cli.BoolFlag); ok {
									if strings.HasPrefix(cur, "-") {
										completion.FlagNames(c, commandName)
										return
									} else {
										completion.ObjectStorageDeleteCompleteArgs(ctx, deleteParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									completion.ObjectStorageDeleteCompleteFlags(ctx, deleteParam, name, cur)
									return
								}
							}
						}
						// here, prev is wrong, so noop.
					} else {
						if strings.HasPrefix(cur, "-") {
							completion.FlagNames(c, commandName)
							return
						} else {
							completion.ObjectStorageDeleteCompleteArgs(ctx, deleteParam, cur, prev, commandName)
							return
						}
					}
				},
				Action: func(c *cli.Context) error {

					if err := checkConfigVersion(); err != nil {
						return err
					}
					if err := applyConfigFromFile(c); err != nil {
						return err
					}

					deleteParam.ParamTemplate = c.String("param-template")
					deleteParam.ParamTemplateFile = c.String("param-template-file")
					strInput, err := command.GetParamTemplateValue(deleteParam)
					if err != nil {
						return err
					}
					if strInput != "" {
						p := params.NewDeleteObjectStorageParam()
						err := json.Unmarshal([]byte(strInput), p)
						if err != nil {
							return fmt.Errorf("Failed to parse JSON: %s", err)
						}
						mergo.MergeWithOverwrite(deleteParam, p)
					}

					// Set option values
					if c.IsSet("access-key") || command.IsEmpty(deleteParam.AccessKey) {
						deleteParam.AccessKey = c.String("access-key")
					}
					if c.IsSet("recursive") {
						deleteParam.Recursive = c.Bool("recursive")
					}
					if c.IsSet("secret-key") || command.IsEmpty(deleteParam.SecretKey) {
						deleteParam.SecretKey = c.String("secret-key")
					}
					if c.IsSet("bucket") || command.IsEmpty(deleteParam.Bucket) {
						deleteParam.Bucket = c.String("bucket")
					}
					if c.IsSet("assumeyes") {
						deleteParam.Assumeyes = c.Bool("assumeyes")
					}
					if c.IsSet("param-template") {
						deleteParam.ParamTemplate = c.String("param-template")
					}
					if c.IsSet("param-template-file") {
						deleteParam.ParamTemplateFile = c.String("param-template-file")
					}
					if c.IsSet("generate-skeleton") {
						deleteParam.GenerateSkeleton = c.Bool("generate-skeleton")
					}

					// Validate global params
					if errors := command.GlobalOption.Validate(true); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					var outputTypeHolder interface{} = deleteParam
					if v, ok := outputTypeHolder.(command.OutputTypeHolder); ok {
						if v.GetOutputType() == "" {
							v.SetOutputType(command.GlobalOption.DefaultOutputType)
						}
					}

					// Generate skeleton
					if deleteParam.GenerateSkeleton {
						deleteParam.GenerateSkeleton = false
						deleteParam.FillValueToSkeleton()
						d, err := json.MarshalIndent(deleteParam, "", "\t")
						if err != nil {
							return fmt.Errorf("Failed to Marshal JSON: %s", err)
						}
						fmt.Fprintln(command.GlobalOption.Out, string(d))
						return nil
					}

					// Validate specific for each command params
					if errors := deleteParam.Validate(); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := command.NewContext(c, c.Args().Slice(), deleteParam)

					// confirm
					if !deleteParam.Assumeyes {
						if !isTerminal() {
							return fmt.Errorf("When using redirect/pipe, specify --assumeyes(-y) option")
						}
						if !command.ConfirmContinue("delete") {
							return nil
						}
					}

					// Run command with params
					return funcs.ObjectStorageDelete(ctx, deleteParam)

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
