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

package cli

import (
	"errors"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command/funcs"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/pkg/utils"
	"github.com/spf13/cobra"
)

// productDiskCmd represents the command to manage SAKURA Cloud ProductDisk
func productDiskCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "product-disk",
		Aliases: []string{"disk-plan"},
		Short:   "A manage commands of ProductDisk",
		Long:    `A manage commands of ProductDisk`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDefaultCmd(cmd, args, "list")
		},
	}
}

func productDiskListCmd() *cobra.Command {
	productDiskListParam := params.NewListProductDiskParam()
	cmd := &cobra.Command{
		Use:          "list",
		Aliases:      []string{"ls", "find"},
		Short:        "List ProductDisk (default)",
		Long:         `List ProductDisk (default)`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return productDiskListParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, productDiskListParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if productDiskListParam.GenerateSkeleton {
				return generateSkeleton(ctx, productDiskListParam)
			}

			return funcs.ProductDiskList(ctx, productDiskListParam)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&productDiskListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &productDiskListParam.Id), "id", "", "set filter by id(s)")
	fs.IntVarP(&productDiskListParam.From, "from", "", 0, "set offset (aliases: offset)")
	fs.IntVarP(&productDiskListParam.Max, "max", "", 0, "set limit (aliases: limit)")
	fs.StringSliceVarP(&productDiskListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringVarP(&productDiskListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&productDiskListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&productDiskListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&productDiskListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&productDiskListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&productDiskListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&productDiskListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&productDiskListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&productDiskListParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&productDiskListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&productDiskListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&productDiskListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.SetNormalizeFunc(productDiskListNormalizeFlagNames)
	buildFlagsUsage(cmd, productDiskListFlagOrder(cmd))
	return cmd
}

func productDiskReadCmd() *cobra.Command {
	productDiskReadParam := params.NewReadProductDiskParam()
	cmd := &cobra.Command{
		Use: "read",

		Short:        "Read ProductDisk",
		Long:         `Read ProductDisk`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return productDiskReadParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, productDiskReadParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if productDiskReadParam.GenerateSkeleton {
				return generateSkeleton(ctx, productDiskReadParam)
			}

			// confirm
			if !productDiskReadParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("read", ctx.IO().In(), ctx.IO().Out())
				if err != nil || !result {
					return err
				}
			}

			return funcs.ProductDiskRead(ctx, productDiskReadParam)

		},
	}

	fs := cmd.Flags()
	fs.BoolVarP(&productDiskReadParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&productDiskReadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&productDiskReadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&productDiskReadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&productDiskReadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&productDiskReadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&productDiskReadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&productDiskReadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&productDiskReadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&productDiskReadParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&productDiskReadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&productDiskReadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&productDiskReadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &productDiskReadParam.Id), "id", "", "set resource ID")
	fs.SetNormalizeFunc(productDiskReadNormalizeFlagNames)
	buildFlagsUsage(cmd, productDiskReadFlagOrder(cmd))
	return cmd
}