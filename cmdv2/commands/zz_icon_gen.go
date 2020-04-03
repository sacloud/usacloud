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

// iconCmd represents the command to manage SAKURA Cloud Icon
func iconCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "icon",
		Short: "A manage commands of Icon",
		Long:  `A manage commands of Icon`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
}

func iconListCmd() *cobra.Command {
	iconListParam := params.NewListIconParam()
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls", "find", "selector"},
		Short:   "List Icon",
		Long:    `List Icon`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return iconListParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, iconListParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if iconListParam.GenerateSkeleton {
				return generateSkeleton(ctx, iconListParam)
			}

			return funcs.IconList(ctx, iconListParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&iconListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &iconListParam.Id), "id", "", "set filter by id(s)")
	fs.StringVarP(&iconListParam.Scope, "scope", "", "", "set filter by scope('user' or 'shared')")
	fs.StringSliceVarP(&iconListParam.Tags, "tags", "", []string{}, "set filter by tags(AND)")
	fs.IntVarP(&iconListParam.From, "from", "", 0, "set offset")
	fs.IntVarP(&iconListParam.Max, "max", "", 0, "set limit")
	fs.StringSliceVarP(&iconListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringVarP(&iconListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&iconListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&iconListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&iconListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&iconListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&iconListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&iconListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&iconListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&iconListParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&iconListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&iconListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&iconListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	buildFlagsUsage(cmd, iconListFlagOrder(cmd))
	return cmd
}

func iconCreateCmd() *cobra.Command {
	iconCreateParam := params.NewCreateIconParam()
	cmd := &cobra.Command{
		Use: "create",

		Short: "Create Icon",
		Long:  `Create Icon`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return iconCreateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, iconCreateParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if iconCreateParam.GenerateSkeleton {
				return generateSkeleton(ctx, iconCreateParam)
			}

			// confirm
			if !iconCreateParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("create", ctx.IO().In(), ctx.IO().Out())
				if err != nil || !result {
					return err
				}
			}

			return funcs.IconCreate(ctx, iconCreateParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&iconCreateParam.Image, "image", "", "", "set file path for upload")
	fs.StringVarP(&iconCreateParam.Name, "name", "", "", "set resource display name")
	fs.StringSliceVarP(&iconCreateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.BoolVarP(&iconCreateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&iconCreateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&iconCreateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&iconCreateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&iconCreateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&iconCreateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&iconCreateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&iconCreateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&iconCreateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&iconCreateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&iconCreateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&iconCreateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&iconCreateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	buildFlagsUsage(cmd, iconCreateFlagOrder(cmd))
	return cmd
}

func iconReadCmd() *cobra.Command {
	iconReadParam := params.NewReadIconParam()
	cmd := &cobra.Command{
		Use: "read",

		Short: "Read Icon",
		Long:  `Read Icon`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return iconReadParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, iconReadParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if iconReadParam.GenerateSkeleton {
				return generateSkeleton(ctx, iconReadParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findIconReadTargets(ctx, iconReadParam)
			if err != nil {
				return err
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				iconReadParam.SetId(id)
				go func(p *params.ReadIconParam) {
					err := funcs.IconRead(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(iconReadParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&iconReadParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&iconReadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&iconReadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&iconReadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&iconReadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&iconReadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&iconReadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&iconReadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&iconReadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&iconReadParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&iconReadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&iconReadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&iconReadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &iconReadParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, iconReadFlagOrder(cmd))
	return cmd
}

func iconUpdateCmd() *cobra.Command {
	iconUpdateParam := params.NewUpdateIconParam()
	cmd := &cobra.Command{
		Use: "update",

		Short: "Update Icon",
		Long:  `Update Icon`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return iconUpdateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, iconUpdateParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if iconUpdateParam.GenerateSkeleton {
				return generateSkeleton(ctx, iconUpdateParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findIconUpdateTargets(ctx, iconUpdateParam)
			if err != nil {
				return err
			}

			// confirm
			if !iconUpdateParam.Assumeyes {
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
				iconUpdateParam.SetId(id)
				go func(p *params.UpdateIconParam) {
					err := funcs.IconUpdate(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(iconUpdateParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&iconUpdateParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&iconUpdateParam.Name, "name", "", "", "set resource display name")
	fs.StringSliceVarP(&iconUpdateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.BoolVarP(&iconUpdateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&iconUpdateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&iconUpdateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&iconUpdateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&iconUpdateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&iconUpdateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&iconUpdateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&iconUpdateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&iconUpdateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&iconUpdateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&iconUpdateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&iconUpdateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&iconUpdateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &iconUpdateParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, iconUpdateFlagOrder(cmd))
	return cmd
}

func iconDeleteCmd() *cobra.Command {
	iconDeleteParam := params.NewDeleteIconParam()
	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"rm"},
		Short:   "Delete Icon",
		Long:    `Delete Icon`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return iconDeleteParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, iconDeleteParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if iconDeleteParam.GenerateSkeleton {
				return generateSkeleton(ctx, iconDeleteParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findIconDeleteTargets(ctx, iconDeleteParam)
			if err != nil {
				return err
			}

			// confirm
			if !iconDeleteParam.Assumeyes {
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
				iconDeleteParam.SetId(id)
				go func(p *params.DeleteIconParam) {
					err := funcs.IconDelete(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(iconDeleteParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&iconDeleteParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&iconDeleteParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&iconDeleteParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&iconDeleteParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&iconDeleteParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&iconDeleteParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&iconDeleteParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&iconDeleteParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&iconDeleteParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&iconDeleteParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&iconDeleteParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&iconDeleteParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&iconDeleteParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&iconDeleteParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &iconDeleteParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, iconDeleteFlagOrder(cmd))
	return cmd
}
