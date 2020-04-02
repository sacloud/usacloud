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

// selfCmd represents the command to manage SAKURA Cloud Self
func selfCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "self",
		Short: "Show self info",
		Long:  `Show self info`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO not implements: call info func as default
		},
	}
}

func selfInfoCmd() *cobra.Command {
	selfInfoParam := params.NewInfoSelfParam()
	cmd := &cobra.Command{
		Use: "info",

		Short: "Info Self (default)",
		Long:  `Info Self (default)`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return selfInfoParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, selfInfoParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if selfInfoParam.GenerateSkeleton {
				return generateSkeleton(ctx, selfInfoParam)
			}

			return funcs.SelfInfo(ctx, selfInfoParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&selfInfoParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&selfInfoParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&selfInfoParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&selfInfoParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&selfInfoParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	return cmd
}

func init() {
	parent := selfCmd()
	parent.AddCommand(selfInfoCmd())
	rootCmd.AddCommand(parent)
}
