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

// productLicenseCmd represents the command to manage SAKURA Cloud ProductLicense
func productLicenseCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "product-license",
		Short: "A manage commands of ProductLicense",
		Long:  `A manage commands of ProductLicense`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO not implements: call list func as default
		},
	}
}

func productLicenseListCmd() *cobra.Command {
	productLicenseListParam := params.NewListProductLicenseParam()
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls", "find"},
		Short:   "List ProductLicense (default)",
		Long:    `List ProductLicense (default)`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return productLicenseListParam.Initialize(newParamsAdapter(cmd.Flags()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, productLicenseListParam)
			if err != nil {
				return err
			}

			// TODO DEBUG
			fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
			fmt.Printf("list local parameter: \n%s\n", debugMarshalIndent(productLicenseListParam))
			return nil
		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&productLicenseListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &productLicenseListParam.Id), "id", "", "set filter by id(s)")
	fs.IntVarP(&productLicenseListParam.From, "from", "", 0, "set offset")
	fs.IntVarP(&productLicenseListParam.Max, "max", "", 0, "set limit")
	fs.StringSliceVarP(&productLicenseListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringVarP(&productLicenseListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&productLicenseListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&productLicenseListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&productLicenseListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&productLicenseListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&productLicenseListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&productLicenseListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&productLicenseListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&productLicenseListParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&productLicenseListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&productLicenseListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&productLicenseListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	return cmd
}

func productLicenseReadCmd() *cobra.Command {
	productLicenseReadParam := params.NewReadProductLicenseParam()
	cmd := &cobra.Command{
		Use: "read",

		Short: "Read ProductLicense",
		Long:  `Read ProductLicense`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return productLicenseReadParam.Initialize(newParamsAdapter(cmd.Flags()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, productLicenseReadParam)
			if err != nil {
				return err
			}

			// TODO DEBUG
			fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
			fmt.Printf("read local parameter: \n%s\n", debugMarshalIndent(productLicenseReadParam))
			return nil
		},
	}

	fs := cmd.Flags()
	fs.BoolVarP(&productLicenseReadParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&productLicenseReadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&productLicenseReadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&productLicenseReadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&productLicenseReadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&productLicenseReadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&productLicenseReadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&productLicenseReadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&productLicenseReadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&productLicenseReadParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&productLicenseReadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&productLicenseReadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&productLicenseReadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &productLicenseReadParam.Id), "id", "", "set resource ID")
	return cmd
}

func init() {
	parent := productLicenseCmd()
	parent.AddCommand(productLicenseListCmd())
	parent.AddCommand(productLicenseReadCmd())
	rootCmd.AddCommand(parent)
}
