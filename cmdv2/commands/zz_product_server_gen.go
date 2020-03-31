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
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/cmdv2/params"
	"github.com/spf13/cobra"
)

// productServerCmd represents the command to manage SAKURA Cloud ProductServer
func productServerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "product-server",
		Short: "A manage commands of ProductServer",
		Long:  `A manage commands of ProductServer`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO not implements: call list func as default
		},
	}
}

func productServerListCmd() *cobra.Command {
	productServerListParam := params.NewListProductServerParam()
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls", "find"},
		Short:   "List ProductServer (default)",
		Long:    `List ProductServer (default)`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return productServerListParam.Initialize(newParamsAdapter(cmd.Flags()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), productServerListParam)
			if err != nil {
				return err
			}

			// TODO DEBUG
			fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
			fmt.Printf("list local parameter: \n%s\n", debugMarshalIndent(productServerListParam))
			return nil
		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&productServerListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &productServerListParam.Id), "id", "", "set filter by id(s)")
	fs.IntVarP(&productServerListParam.From, "from", "", 0, "set offset")
	fs.IntVarP(&productServerListParam.Max, "max", "", 0, "set limit")
	fs.StringSliceVarP(&productServerListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringVarP(&productServerListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&productServerListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&productServerListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&productServerListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&productServerListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&productServerListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&productServerListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&productServerListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&productServerListParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&productServerListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&productServerListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&productServerListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	return cmd
}

func productServerReadCmd() *cobra.Command {
	productServerReadParam := params.NewReadProductServerParam()
	cmd := &cobra.Command{
		Use: "read",

		Short: "Read ProductServer",
		Long:  `Read ProductServer`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return productServerReadParam.Initialize(newParamsAdapter(cmd.Flags()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), productServerReadParam)
			if err != nil {
				return err
			}

			// TODO DEBUG
			fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
			fmt.Printf("read local parameter: \n%s\n", debugMarshalIndent(productServerReadParam))
			return nil
		},
	}

	fs := cmd.Flags()
	fs.BoolVarP(&productServerReadParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&productServerReadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&productServerReadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&productServerReadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&productServerReadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&productServerReadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&productServerReadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&productServerReadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&productServerReadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&productServerReadParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&productServerReadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&productServerReadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&productServerReadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &productServerReadParam.Id), "id", "", "set resource ID")
	return cmd
}

func init() {
	parent := productServerCmd()
	parent.AddCommand(productServerListCmd())
	parent.AddCommand(productServerReadCmd())
	rootCmd.AddCommand(parent)
}
