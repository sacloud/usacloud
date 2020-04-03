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
	"sync"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/cmdv2/params"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/funcs"
	"github.com/sacloud/usacloud/pkg/utils"
	"github.com/spf13/cobra"
)

// bridgeCmd represents the command to manage SAKURA Cloud Bridge
func bridgeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "bridge",
		Short: "A manage commands of Bridge",
		Long:  `A manage commands of Bridge`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
}

func bridgeListCmd() *cobra.Command {
	bridgeListParam := params.NewListBridgeParam()
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls", "find"},
		Short:   "List Bridge",
		Long:    `List Bridge`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return bridgeListParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, bridgeListParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if bridgeListParam.GenerateSkeleton {
				return generateSkeleton(ctx, bridgeListParam)
			}

			return funcs.BridgeList(ctx, bridgeListParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&bridgeListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &bridgeListParam.Id), "id", "", "set filter by id(s)")
	fs.IntVarP(&bridgeListParam.From, "from", "", 0, "set offset")
	fs.IntVarP(&bridgeListParam.Max, "max", "", 0, "set limit")
	fs.StringSliceVarP(&bridgeListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringVarP(&bridgeListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&bridgeListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&bridgeListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&bridgeListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&bridgeListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&bridgeListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&bridgeListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&bridgeListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&bridgeListParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&bridgeListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&bridgeListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&bridgeListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	buildFlagsUsage(cmd, bridgeListFlagOrder(cmd))
	return cmd
}

func bridgeCreateCmd() *cobra.Command {
	bridgeCreateParam := params.NewCreateBridgeParam()
	cmd := &cobra.Command{
		Use: "create",

		Short: "Create Bridge",
		Long:  `Create Bridge`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return bridgeCreateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, bridgeCreateParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if bridgeCreateParam.GenerateSkeleton {
				return generateSkeleton(ctx, bridgeCreateParam)
			}

			// confirm
			if !bridgeCreateParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("create", ctx.IO().In(), ctx.IO().Out())
				if err != nil || !result {
					return err
				}
			}

			return funcs.BridgeCreate(ctx, bridgeCreateParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&bridgeCreateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&bridgeCreateParam.Description, "description", "", "", "set resource description")
	fs.BoolVarP(&bridgeCreateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&bridgeCreateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&bridgeCreateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&bridgeCreateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&bridgeCreateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&bridgeCreateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&bridgeCreateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&bridgeCreateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&bridgeCreateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&bridgeCreateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&bridgeCreateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&bridgeCreateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&bridgeCreateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	buildFlagsUsage(cmd, bridgeCreateFlagOrder(cmd))
	return cmd
}

func bridgeReadCmd() *cobra.Command {
	bridgeReadParam := params.NewReadBridgeParam()
	cmd := &cobra.Command{
		Use: "read",

		Short: "Read Bridge",
		Long:  `Read Bridge`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return bridgeReadParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, bridgeReadParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if bridgeReadParam.GenerateSkeleton {
				return generateSkeleton(ctx, bridgeReadParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findBridgeReadTargets(ctx, bridgeReadParam)
			if err != nil {
				return err
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				bridgeReadParam.SetId(id)
				go func(p *params.ReadBridgeParam) {
					err := funcs.BridgeRead(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(bridgeReadParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&bridgeReadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&bridgeReadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&bridgeReadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&bridgeReadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&bridgeReadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&bridgeReadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&bridgeReadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&bridgeReadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&bridgeReadParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&bridgeReadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&bridgeReadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&bridgeReadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &bridgeReadParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, bridgeReadFlagOrder(cmd))
	return cmd
}

func bridgeUpdateCmd() *cobra.Command {
	bridgeUpdateParam := params.NewUpdateBridgeParam()
	cmd := &cobra.Command{
		Use: "update",

		Short: "Update Bridge",
		Long:  `Update Bridge`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return bridgeUpdateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, bridgeUpdateParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if bridgeUpdateParam.GenerateSkeleton {
				return generateSkeleton(ctx, bridgeUpdateParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findBridgeUpdateTargets(ctx, bridgeUpdateParam)
			if err != nil {
				return err
			}

			// confirm
			if !bridgeUpdateParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("update", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				bridgeUpdateParam.SetId(id)
				go func(p *params.UpdateBridgeParam) {
					err := funcs.BridgeUpdate(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(bridgeUpdateParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&bridgeUpdateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&bridgeUpdateParam.Description, "description", "", "", "set resource description")
	fs.BoolVarP(&bridgeUpdateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&bridgeUpdateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&bridgeUpdateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&bridgeUpdateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&bridgeUpdateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&bridgeUpdateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&bridgeUpdateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&bridgeUpdateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&bridgeUpdateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&bridgeUpdateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&bridgeUpdateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&bridgeUpdateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&bridgeUpdateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &bridgeUpdateParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, bridgeUpdateFlagOrder(cmd))
	return cmd
}

func bridgeDeleteCmd() *cobra.Command {
	bridgeDeleteParam := params.NewDeleteBridgeParam()
	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"rm"},
		Short:   "Delete Bridge",
		Long:    `Delete Bridge`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return bridgeDeleteParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, bridgeDeleteParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if bridgeDeleteParam.GenerateSkeleton {
				return generateSkeleton(ctx, bridgeDeleteParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findBridgeDeleteTargets(ctx, bridgeDeleteParam)
			if err != nil {
				return err
			}

			// confirm
			if !bridgeDeleteParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("delete", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				bridgeDeleteParam.SetId(id)
				go func(p *params.DeleteBridgeParam) {
					err := funcs.BridgeDelete(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(bridgeDeleteParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.BoolVarP(&bridgeDeleteParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&bridgeDeleteParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&bridgeDeleteParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&bridgeDeleteParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&bridgeDeleteParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&bridgeDeleteParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&bridgeDeleteParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&bridgeDeleteParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&bridgeDeleteParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&bridgeDeleteParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&bridgeDeleteParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&bridgeDeleteParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&bridgeDeleteParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &bridgeDeleteParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, bridgeDeleteFlagOrder(cmd))
	return cmd
}

func init() {
	parent := bridgeCmd()
	parent.AddCommand(bridgeListCmd())
	parent.AddCommand(bridgeCreateCmd())
	parent.AddCommand(bridgeReadCmd())
	parent.AddCommand(bridgeUpdateCmd())
	parent.AddCommand(bridgeDeleteCmd())
	buildCommandsUsage(parent, bridgeCommandOrder(parent))
	rootCmd.AddCommand(parent)
}
