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
	"github.com/sacloud/usacloud/command/funcs"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/pkg/utils"
	"github.com/spf13/cobra"
)

// ipv4Cmd represents the command to manage SAKURA Cloud IPv4
func ipv4Cmd() *cobra.Command {
	return &cobra.Command{
		Use: "ipv4",

		Short: "A manage commands of IPv4",
		Long:  `A manage commands of IPv4`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.HelpFunc()(cmd, args)
			return nil
		},
	}
}

func ipv4ListCmd() *cobra.Command {
	ipv4ListParam := params.NewListIPv4Param()
	cmd := &cobra.Command{
		Use:          "list",
		Aliases:      []string{"ls", "find"},
		Short:        "List IPv4",
		Long:         `List IPv4`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return ipv4ListParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, ipv4ListParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if ipv4ListParam.GenerateSkeleton {
				return generateSkeleton(ctx, ipv4ListParam)
			}

			return funcs.IPv4List(ctx, ipv4ListParam)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&ipv4ListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &ipv4ListParam.Id), "id", "", "set filter by id(s)")
	fs.IntVarP(&ipv4ListParam.From, "from", "", 0, "set offset (aliases: offset)")
	fs.IntVarP(&ipv4ListParam.Max, "max", "", 0, "set limit (aliases: limit)")
	fs.StringSliceVarP(&ipv4ListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringVarP(&ipv4ListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&ipv4ListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&ipv4ListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&ipv4ListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&ipv4ListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&ipv4ListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&ipv4ListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&ipv4ListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&ipv4ListParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&ipv4ListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&ipv4ListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&ipv4ListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.SetNormalizeFunc(ipv4ListNormalizeFlagNames)
	buildFlagsUsage(cmd, ipv4ListFlagOrder(cmd))
	return cmd
}

func ipv4PtrAddCmd() *cobra.Command {
	ipv4PtrAddParam := params.NewPtrAddIPv4Param()
	cmd := &cobra.Command{
		Use: "ptr-add",

		Short:        "PtrAdd IPv4",
		Long:         `PtrAdd IPv4`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return ipv4PtrAddParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, ipv4PtrAddParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if ipv4PtrAddParam.GenerateSkeleton {
				return generateSkeleton(ctx, ipv4PtrAddParam)
			}

			// confirm
			if !ipv4PtrAddParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("ptr-add", ctx.IO().In(), ctx.IO().Out())
				if err != nil || !result {
					return err
				}
			}

			return funcs.IPv4PtrAdd(ctx, ipv4PtrAddParam)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&ipv4PtrAddParam.Hostname, "hostname", "", "", "set server hostname")
	fs.BoolVarP(&ipv4PtrAddParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&ipv4PtrAddParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&ipv4PtrAddParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&ipv4PtrAddParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&ipv4PtrAddParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&ipv4PtrAddParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&ipv4PtrAddParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&ipv4PtrAddParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&ipv4PtrAddParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&ipv4PtrAddParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&ipv4PtrAddParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&ipv4PtrAddParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&ipv4PtrAddParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.SetNormalizeFunc(ipv4PtrAddNormalizeFlagNames)
	buildFlagsUsage(cmd, ipv4PtrAddFlagOrder(cmd))
	return cmd
}

func ipv4PtrReadCmd() *cobra.Command {
	ipv4PtrReadParam := params.NewPtrReadIPv4Param()
	cmd := &cobra.Command{
		Use: "ptr-read",

		Short:        "PtrRead IPv4",
		Long:         `PtrRead IPv4`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return ipv4PtrReadParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, ipv4PtrReadParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if ipv4PtrReadParam.GenerateSkeleton {
				return generateSkeleton(ctx, ipv4PtrReadParam)
			}

			return funcs.IPv4PtrRead(ctx, ipv4PtrReadParam)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&ipv4PtrReadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&ipv4PtrReadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&ipv4PtrReadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&ipv4PtrReadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&ipv4PtrReadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&ipv4PtrReadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&ipv4PtrReadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&ipv4PtrReadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&ipv4PtrReadParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&ipv4PtrReadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&ipv4PtrReadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&ipv4PtrReadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.SetNormalizeFunc(ipv4PtrReadNormalizeFlagNames)
	buildFlagsUsage(cmd, ipv4PtrReadFlagOrder(cmd))
	return cmd
}

func ipv4PtrUpdateCmd() *cobra.Command {
	ipv4PtrUpdateParam := params.NewPtrUpdateIPv4Param()
	cmd := &cobra.Command{
		Use: "ptr-update",

		Short:        "PtrUpdate IPv4",
		Long:         `PtrUpdate IPv4`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return ipv4PtrUpdateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, ipv4PtrUpdateParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if ipv4PtrUpdateParam.GenerateSkeleton {
				return generateSkeleton(ctx, ipv4PtrUpdateParam)
			}

			// confirm
			if !ipv4PtrUpdateParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("ptr-update", ctx.IO().In(), ctx.IO().Out())
				if err != nil || !result {
					return err
				}
			}

			return funcs.IPv4PtrUpdate(ctx, ipv4PtrUpdateParam)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&ipv4PtrUpdateParam.Hostname, "hostname", "", "", "set server hostname")
	fs.BoolVarP(&ipv4PtrUpdateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&ipv4PtrUpdateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&ipv4PtrUpdateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&ipv4PtrUpdateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&ipv4PtrUpdateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&ipv4PtrUpdateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&ipv4PtrUpdateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&ipv4PtrUpdateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&ipv4PtrUpdateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&ipv4PtrUpdateParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&ipv4PtrUpdateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&ipv4PtrUpdateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&ipv4PtrUpdateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.SetNormalizeFunc(ipv4PtrUpdateNormalizeFlagNames)
	buildFlagsUsage(cmd, ipv4PtrUpdateFlagOrder(cmd))
	return cmd
}

func ipv4PtrDeleteCmd() *cobra.Command {
	ipv4PtrDeleteParam := params.NewPtrDeleteIPv4Param()
	cmd := &cobra.Command{
		Use: "ptr-delete",

		Short:        "PtrDelete IPv4",
		Long:         `PtrDelete IPv4`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return ipv4PtrDeleteParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, ipv4PtrDeleteParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if ipv4PtrDeleteParam.GenerateSkeleton {
				return generateSkeleton(ctx, ipv4PtrDeleteParam)
			}

			// confirm
			if !ipv4PtrDeleteParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("ptr-delete", ctx.IO().In(), ctx.IO().Out())
				if err != nil || !result {
					return err
				}
			}

			return funcs.IPv4PtrDelete(ctx, ipv4PtrDeleteParam)

		},
	}

	fs := cmd.Flags()
	fs.BoolVarP(&ipv4PtrDeleteParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&ipv4PtrDeleteParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&ipv4PtrDeleteParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&ipv4PtrDeleteParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&ipv4PtrDeleteParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&ipv4PtrDeleteParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&ipv4PtrDeleteParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&ipv4PtrDeleteParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&ipv4PtrDeleteParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&ipv4PtrDeleteParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&ipv4PtrDeleteParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&ipv4PtrDeleteParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&ipv4PtrDeleteParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.SetNormalizeFunc(ipv4PtrDeleteNormalizeFlagNames)
	buildFlagsUsage(cmd, ipv4PtrDeleteFlagOrder(cmd))
	return cmd
}
