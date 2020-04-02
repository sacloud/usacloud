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

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/cmdv2/params"
	"github.com/sacloud/usacloud/command/funcs"
	"github.com/sacloud/usacloud/pkg/utils"
	"github.com/spf13/cobra"
)

// regionCmd represents the command to manage SAKURA Cloud Region
func regionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "region",
		Short: "A manage commands of Region",
		Long:  `A manage commands of Region`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO not implements: call list func as default
		},
	}
}

func regionListCmd() *cobra.Command {
	regionListParam := params.NewListRegionParam()
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls", "find"},
		Short:   "List Region (default)",
		Long:    `List Region (default)`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return regionListParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, regionListParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if regionListParam.GenerateSkeleton {
				return generateSkeleton(ctx, regionListParam)
			}

			return funcs.RegionList(ctx, regionListParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&regionListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &regionListParam.Id), "id", "", "set filter by id(s)")
	fs.IntVarP(&regionListParam.From, "from", "", 0, "set offset")
	fs.IntVarP(&regionListParam.Max, "max", "", 0, "set limit")
	fs.StringSliceVarP(&regionListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringVarP(&regionListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&regionListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&regionListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&regionListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&regionListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&regionListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&regionListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&regionListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&regionListParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&regionListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&regionListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&regionListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	return cmd
}

func regionReadCmd() *cobra.Command {
	regionReadParam := params.NewReadRegionParam()
	cmd := &cobra.Command{
		Use: "read",

		Short: "Read Region",
		Long:  `Read Region`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return regionReadParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, regionReadParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if regionReadParam.GenerateSkeleton {
				return generateSkeleton(ctx, regionReadParam)
			}

			// confirm
			if !regionReadParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("read", ctx.IO().In(), ctx.IO().Out())
				if err != nil || !result {
					return err
				}
			}

			return funcs.RegionRead(ctx, regionReadParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.BoolVarP(&regionReadParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&regionReadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&regionReadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&regionReadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&regionReadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&regionReadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&regionReadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&regionReadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&regionReadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&regionReadParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&regionReadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&regionReadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&regionReadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &regionReadParam.Id), "id", "", "set resource ID")
	return cmd
}

func init() {
	parent := regionCmd()
	parent.AddCommand(regionListCmd())
	parent.AddCommand(regionReadCmd())
	rootCmd.AddCommand(parent)
}
