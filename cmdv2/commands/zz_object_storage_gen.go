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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-v2-commands'; DO NOT EDIT

package commands

import (
	"errors"

	"github.com/sacloud/usacloud/cmdv2/params"
	"github.com/sacloud/usacloud/command/funcs"
	"github.com/sacloud/usacloud/pkg/utils"
	"github.com/spf13/cobra"
)

// objectStorageCmd represents the command to manage SAKURA Cloud ObjectStorage
func objectStorageCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "object-storage",
		Short: "A manage commands of ObjectStorage",
		Long:  `A manage commands of ObjectStorage`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
}

func objectStorageListCmd() *cobra.Command {
	objectStorageListParam := params.NewListObjectStorageParam()
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List ObjectStorage",
		Long:    `List ObjectStorage`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return objectStorageListParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, objectStorageListParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if objectStorageListParam.GenerateSkeleton {
				return generateSkeleton(ctx, objectStorageListParam)
			}

			return funcs.ObjectStorageList(ctx, objectStorageListParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&objectStorageListParam.AccessKey, "access-key", "", "", "set access-key")
	fs.StringVarP(&objectStorageListParam.SecretKey, "secret-key", "", "", "set access-key")
	fs.StringVarP(&objectStorageListParam.Bucket, "bucket", "", "", "set bucket")
	fs.StringVarP(&objectStorageListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&objectStorageListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&objectStorageListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&objectStorageListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&objectStorageListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&objectStorageListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&objectStorageListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&objectStorageListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&objectStorageListParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&objectStorageListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&objectStorageListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&objectStorageListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	setFlagsUsage(cmd, buildFlagsUsage(objectStorageListFlagOrder(cmd)))
	return cmd
}

func objectStoragePutCmd() *cobra.Command {
	objectStoragePutParam := params.NewPutObjectStorageParam()
	cmd := &cobra.Command{
		Use: "put",

		Short: "Put ObjectStorage",
		Long:  `Put ObjectStorage`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return objectStoragePutParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, objectStoragePutParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if objectStoragePutParam.GenerateSkeleton {
				return generateSkeleton(ctx, objectStoragePutParam)
			}

			// confirm
			if !objectStoragePutParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("put", ctx.IO().In(), ctx.IO().Out())
				if err != nil || !result {
					return err
				}
			}

			return funcs.ObjectStoragePut(ctx, objectStoragePutParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&objectStoragePutParam.AccessKey, "access-key", "", "", "set access-key")
	fs.StringVarP(&objectStoragePutParam.ContentType, "content-type", "", "application/octet-stream", "set content-type")
	fs.BoolVarP(&objectStoragePutParam.Recursive, "recursive", "r", false, "put objects recursive")
	fs.StringVarP(&objectStoragePutParam.SecretKey, "secret-key", "", "", "set access-key")
	fs.StringVarP(&objectStoragePutParam.Bucket, "bucket", "", "", "set bucket")
	fs.BoolVarP(&objectStoragePutParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&objectStoragePutParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&objectStoragePutParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&objectStoragePutParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&objectStoragePutParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&objectStoragePutParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	setFlagsUsage(cmd, buildFlagsUsage(objectStoragePutFlagOrder(cmd)))
	return cmd
}

func objectStorageGetCmd() *cobra.Command {
	objectStorageGetParam := params.NewGetObjectStorageParam()
	cmd := &cobra.Command{
		Use: "get",

		Short: "Get ObjectStorage",
		Long:  `Get ObjectStorage`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return objectStorageGetParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, objectStorageGetParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if objectStorageGetParam.GenerateSkeleton {
				return generateSkeleton(ctx, objectStorageGetParam)
			}

			return funcs.ObjectStorageGet(ctx, objectStorageGetParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&objectStorageGetParam.AccessKey, "access-key", "", "", "set access-key")
	fs.BoolVarP(&objectStorageGetParam.Recursive, "recursive", "r", false, "get objects recursive")
	fs.StringVarP(&objectStorageGetParam.SecretKey, "secret-key", "", "", "set access-key")
	fs.StringVarP(&objectStorageGetParam.Bucket, "bucket", "", "", "set bucket")
	fs.StringVarP(&objectStorageGetParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&objectStorageGetParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&objectStorageGetParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&objectStorageGetParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&objectStorageGetParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	setFlagsUsage(cmd, buildFlagsUsage(objectStorageGetFlagOrder(cmd)))
	return cmd
}

func objectStorageDeleteCmd() *cobra.Command {
	objectStorageDeleteParam := params.NewDeleteObjectStorageParam()
	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"rm", "del"},
		Short:   "Delete ObjectStorage",
		Long:    `Delete ObjectStorage`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return objectStorageDeleteParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, objectStorageDeleteParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if objectStorageDeleteParam.GenerateSkeleton {
				return generateSkeleton(ctx, objectStorageDeleteParam)
			}

			// confirm
			if !objectStorageDeleteParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("delete", ctx.IO().In(), ctx.IO().Out())
				if err != nil || !result {
					return err
				}
			}

			return funcs.ObjectStorageDelete(ctx, objectStorageDeleteParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&objectStorageDeleteParam.AccessKey, "access-key", "", "", "set access-key")
	fs.BoolVarP(&objectStorageDeleteParam.Recursive, "recursive", "r", false, "delete objects recursive")
	fs.StringVarP(&objectStorageDeleteParam.SecretKey, "secret-key", "", "", "set access-key")
	fs.StringVarP(&objectStorageDeleteParam.Bucket, "bucket", "", "", "set bucket")
	fs.BoolVarP(&objectStorageDeleteParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&objectStorageDeleteParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&objectStorageDeleteParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&objectStorageDeleteParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&objectStorageDeleteParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&objectStorageDeleteParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	setFlagsUsage(cmd, buildFlagsUsage(objectStorageDeleteFlagOrder(cmd)))
	return cmd
}

func init() {
	parent := objectStorageCmd()
	parent.AddCommand(objectStorageListCmd())
	parent.AddCommand(objectStoragePutCmd())
	parent.AddCommand(objectStorageGetCmd())
	parent.AddCommand(objectStorageDeleteCmd())
	rootCmd.AddCommand(parent)
}
