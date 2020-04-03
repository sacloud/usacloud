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

// startupScriptCmd represents the command to manage SAKURA Cloud StartupScript
func startupScriptCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "startup-script",
		Short: "A manage commands of StartupScript",
		Long:  `A manage commands of StartupScript`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
}

func startupScriptListCmd() *cobra.Command {
	startupScriptListParam := params.NewListStartupScriptParam()
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls", "find", "selector"},
		Short:   "List StartupScript",
		Long:    `List StartupScript`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return startupScriptListParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, startupScriptListParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if startupScriptListParam.GenerateSkeleton {
				return generateSkeleton(ctx, startupScriptListParam)
			}

			return funcs.StartupScriptList(ctx, startupScriptListParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&startupScriptListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &startupScriptListParam.Id), "id", "", "set filter by id(s)")
	fs.StringVarP(&startupScriptListParam.Scope, "scope", "", "", "set filter by scope('user' or 'shared')")
	fs.StringSliceVarP(&startupScriptListParam.Tags, "tags", "", []string{}, "set filter by tags(AND)")
	fs.StringSliceVarP(&startupScriptListParam.Class, "class", "", []string{}, "set filter by class(es)")
	fs.IntVarP(&startupScriptListParam.From, "from", "", 0, "set offset")
	fs.IntVarP(&startupScriptListParam.Max, "max", "", 0, "set limit")
	fs.StringSliceVarP(&startupScriptListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringVarP(&startupScriptListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&startupScriptListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&startupScriptListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&startupScriptListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&startupScriptListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&startupScriptListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&startupScriptListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&startupScriptListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&startupScriptListParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&startupScriptListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&startupScriptListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&startupScriptListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	buildFlagsUsage(cmd, startupScriptListFlagOrder(cmd))
	return cmd
}

func startupScriptCreateCmd() *cobra.Command {
	startupScriptCreateParam := params.NewCreateStartupScriptParam()
	cmd := &cobra.Command{
		Use: "create",

		Short: "Create StartupScript",
		Long:  `Create StartupScript`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return startupScriptCreateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, startupScriptCreateParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if startupScriptCreateParam.GenerateSkeleton {
				return generateSkeleton(ctx, startupScriptCreateParam)
			}

			// confirm
			if !startupScriptCreateParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("create", ctx.IO().In(), ctx.IO().Out())
				if err != nil || !result {
					return err
				}
			}

			return funcs.StartupScriptCreate(ctx, startupScriptCreateParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&startupScriptCreateParam.Script, "script", "", "", "set script from file")
	fs.StringVarP(&startupScriptCreateParam.ScriptContent, "script-content", "", "", "set script content")
	fs.StringVarP(&startupScriptCreateParam.Class, "class", "", "shell", "set script class[shell/cloud-config-yaml]")
	fs.StringVarP(&startupScriptCreateParam.Name, "name", "", "", "set resource display name")
	fs.StringSliceVarP(&startupScriptCreateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &startupScriptCreateParam.IconId), "icon-id", "", "set Icon ID")
	fs.BoolVarP(&startupScriptCreateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&startupScriptCreateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&startupScriptCreateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&startupScriptCreateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&startupScriptCreateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&startupScriptCreateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&startupScriptCreateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&startupScriptCreateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&startupScriptCreateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&startupScriptCreateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&startupScriptCreateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&startupScriptCreateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&startupScriptCreateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	buildFlagsUsage(cmd, startupScriptCreateFlagOrder(cmd))
	return cmd
}

func startupScriptReadCmd() *cobra.Command {
	startupScriptReadParam := params.NewReadStartupScriptParam()
	cmd := &cobra.Command{
		Use: "read",

		Short: "Read StartupScript",
		Long:  `Read StartupScript`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return startupScriptReadParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, startupScriptReadParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if startupScriptReadParam.GenerateSkeleton {
				return generateSkeleton(ctx, startupScriptReadParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findStartupScriptReadTargets(ctx, startupScriptReadParam)
			if err != nil {
				return err
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				startupScriptReadParam.SetId(id)
				go func(p *params.ReadStartupScriptParam) {
					err := funcs.StartupScriptRead(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(startupScriptReadParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&startupScriptReadParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&startupScriptReadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&startupScriptReadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&startupScriptReadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&startupScriptReadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&startupScriptReadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&startupScriptReadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&startupScriptReadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&startupScriptReadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&startupScriptReadParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&startupScriptReadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&startupScriptReadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&startupScriptReadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &startupScriptReadParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, startupScriptReadFlagOrder(cmd))
	return cmd
}

func startupScriptUpdateCmd() *cobra.Command {
	startupScriptUpdateParam := params.NewUpdateStartupScriptParam()
	cmd := &cobra.Command{
		Use: "update",

		Short: "Update StartupScript",
		Long:  `Update StartupScript`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return startupScriptUpdateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, startupScriptUpdateParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if startupScriptUpdateParam.GenerateSkeleton {
				return generateSkeleton(ctx, startupScriptUpdateParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findStartupScriptUpdateTargets(ctx, startupScriptUpdateParam)
			if err != nil {
				return err
			}

			// confirm
			if !startupScriptUpdateParam.Assumeyes {
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
				startupScriptUpdateParam.SetId(id)
				go func(p *params.UpdateStartupScriptParam) {
					err := funcs.StartupScriptUpdate(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(startupScriptUpdateParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&startupScriptUpdateParam.Script, "script", "", "", "set script from file")
	fs.StringVarP(&startupScriptUpdateParam.ScriptContent, "script-content", "", "", "set script content")
	fs.StringVarP(&startupScriptUpdateParam.Class, "class", "", "", "set script class[shell/cloud-config-yaml]")
	fs.StringSliceVarP(&startupScriptUpdateParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&startupScriptUpdateParam.Name, "name", "", "", "set resource display name")
	fs.StringSliceVarP(&startupScriptUpdateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &startupScriptUpdateParam.IconId), "icon-id", "", "set Icon ID")
	fs.BoolVarP(&startupScriptUpdateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&startupScriptUpdateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&startupScriptUpdateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&startupScriptUpdateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&startupScriptUpdateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&startupScriptUpdateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&startupScriptUpdateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&startupScriptUpdateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&startupScriptUpdateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&startupScriptUpdateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&startupScriptUpdateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&startupScriptUpdateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&startupScriptUpdateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &startupScriptUpdateParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, startupScriptUpdateFlagOrder(cmd))
	return cmd
}

func startupScriptDeleteCmd() *cobra.Command {
	startupScriptDeleteParam := params.NewDeleteStartupScriptParam()
	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"rm"},
		Short:   "Delete StartupScript",
		Long:    `Delete StartupScript`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return startupScriptDeleteParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, startupScriptDeleteParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if startupScriptDeleteParam.GenerateSkeleton {
				return generateSkeleton(ctx, startupScriptDeleteParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findStartupScriptDeleteTargets(ctx, startupScriptDeleteParam)
			if err != nil {
				return err
			}

			// confirm
			if !startupScriptDeleteParam.Assumeyes {
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
				startupScriptDeleteParam.SetId(id)
				go func(p *params.DeleteStartupScriptParam) {
					err := funcs.StartupScriptDelete(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(startupScriptDeleteParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&startupScriptDeleteParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&startupScriptDeleteParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&startupScriptDeleteParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&startupScriptDeleteParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&startupScriptDeleteParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&startupScriptDeleteParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&startupScriptDeleteParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&startupScriptDeleteParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&startupScriptDeleteParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&startupScriptDeleteParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&startupScriptDeleteParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&startupScriptDeleteParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&startupScriptDeleteParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&startupScriptDeleteParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &startupScriptDeleteParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, startupScriptDeleteFlagOrder(cmd))
	return cmd
}

func init() {
	parent := startupScriptCmd()
	parent.AddCommand(startupScriptListCmd())
	parent.AddCommand(startupScriptCreateCmd())
	parent.AddCommand(startupScriptReadCmd())
	parent.AddCommand(startupScriptUpdateCmd())
	parent.AddCommand(startupScriptDeleteCmd())
	buildCommandsUsage(parent, startupScriptCommandOrder(parent))
	rootCmd.AddCommand(parent)
}
