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
	"github.com/sacloud/usacloud/cmdv2/params"
	"github.com/sacloud/usacloud/command/funcs"
	"github.com/spf13/cobra"
)

// couponCmd represents the command to manage SAKURA Cloud Coupon
func couponCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "coupon",
		Short: "A manage commands of Coupon",
		Long:  `A manage commands of Coupon`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO not implements: call list func as default
		},
	}
}

func couponListCmd() *cobra.Command {
	couponListParam := params.NewListCouponParam()
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls", "find"},
		Short:   "List Coupon (default)",
		Long:    `List Coupon (default)`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return couponListParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, couponListParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if couponListParam.GenerateSkeleton {
				return generateSkeleton(ctx, couponListParam)
			}

			return funcs.CouponList(ctx, couponListParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&couponListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&couponListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&couponListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&couponListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&couponListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&couponListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.BoolVarP(&couponListParam.Usable, "usable", "", false, "show usable coupons only")
	fs.StringSliceVarP(&couponListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&couponListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&couponListParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&couponListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&couponListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&couponListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	return cmd
}

func init() {
	parent := couponCmd()
	parent.AddCommand(couponListCmd())
	rootCmd.AddCommand(parent)
}
