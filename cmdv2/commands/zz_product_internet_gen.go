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

var (
	productInternetListParam = params.NewListProductInternetParam()
	productInternetReadParam = params.NewReadProductInternetParam()
)

// productInternetCmd represents the command to manage SAKURA Cloud ProductInternet
var productInternetCmd = &cobra.Command{
	Use:   "product-internet",
	Short: "A manage commands of ProductInternet",
	Long:  `A manage commands of ProductInternet`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO not implements: call list func as default
	},
}

var productInternetListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "find"},
	Short:   "List ProductInternet (default)",
	Long:    `List ProductInternet (default)`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return productInternetListParam.Initialize(newParamsAdapter(cmd.Flags()))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := newCLIContext(globalFlags(), productInternetListParam)
		if err != nil {
			return err
		}

		// TODO DEBUG
		fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
		fmt.Printf("list local parameter: \n%s\n", debugMarshalIndent(productInternetListParam))
		return nil
	},
}

func productInternetListCmdInit() {
	fs := productInternetListCmd.Flags()
	fs.StringSliceVarP(&productInternetListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &productInternetListParam.Id), "id", "", "set filter by id(s)")
	fs.IntVarP(&productInternetListParam.From, "from", "", 0, "set offset")
	fs.IntVarP(&productInternetListParam.Max, "max", "", 0, "set limit")
	fs.StringSliceVarP(&productInternetListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringVarP(&productInternetListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&productInternetListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&productInternetListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&productInternetListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&productInternetListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&productInternetListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&productInternetListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&productInternetListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&productInternetListParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&productInternetListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&productInternetListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&productInternetListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
}

var productInternetReadCmd = &cobra.Command{
	Use: "read",

	Short: "Read ProductInternet",
	Long:  `Read ProductInternet`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return productInternetReadParam.Initialize(newParamsAdapter(cmd.Flags()))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := newCLIContext(globalFlags(), productInternetReadParam)
		if err != nil {
			return err
		}

		// TODO DEBUG
		fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
		fmt.Printf("read local parameter: \n%s\n", debugMarshalIndent(productInternetReadParam))
		return nil
	},
}

func productInternetReadCmdInit() {
	fs := productInternetReadCmd.Flags()
	fs.BoolVarP(&productInternetReadParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&productInternetReadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&productInternetReadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&productInternetReadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&productInternetReadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&productInternetReadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&productInternetReadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&productInternetReadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&productInternetReadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&productInternetReadParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&productInternetReadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&productInternetReadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&productInternetReadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &productInternetReadParam.Id), "id", "", "set resource ID")
}

func init() {
	parent := productInternetCmd

	productInternetListCmdInit()
	parent.AddCommand(productInternetListCmd)

	productInternetReadCmdInit()
	parent.AddCommand(productInternetReadCmd)

	rootCmd.AddCommand(parent)
}
