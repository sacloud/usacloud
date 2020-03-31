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

// authStatusCmd represents the command to manage SAKURA Cloud AuthStatus
func authStatusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "auth-status",
		Short: "A manage commands of AuthStatus",
		Long:  `A manage commands of AuthStatus`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO not implements: call show func as default
		},
	}
}

func authStatusShowCmd() *cobra.Command {
	authStatusShowParam := params.NewShowAuthStatusParam()
	cmd := &cobra.Command{
		Use: "show",

		Short: "Show AuthStatus (default)",
		Long:  `Show AuthStatus (default)`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return authStatusShowParam.Initialize(newParamsAdapter(cmd.Flags()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, authStatusShowParam)
			if err != nil {
				return err
			}
			return funcs.AuthStatusShow(ctx, authStatusShowParam.ToV0())
		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&authStatusShowParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&authStatusShowParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&authStatusShowParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&authStatusShowParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&authStatusShowParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&authStatusShowParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&authStatusShowParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&authStatusShowParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&authStatusShowParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&authStatusShowParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&authStatusShowParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&authStatusShowParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	return cmd
}

func init() {
	parent := authStatusCmd()
	parent.AddCommand(authStatusShowCmd())
	rootCmd.AddCommand(parent)
}
