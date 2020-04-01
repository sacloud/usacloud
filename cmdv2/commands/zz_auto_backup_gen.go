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
	"github.com/sacloud/usacloud/cmdv2/params"
	"github.com/sacloud/usacloud/command/funcs"
	"github.com/sacloud/usacloud/pkg/utils"
	"github.com/spf13/cobra"
)

// autoBackupCmd represents the command to manage SAKURA Cloud AutoBackup
func autoBackupCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "auto-backup",
		Short: "A manage commands of AutoBackup",
		Long:  `A manage commands of AutoBackup`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
}

func autoBackupListCmd() *cobra.Command {
	autoBackupListParam := params.NewListAutoBackupParam()
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls", "find", "selector"},
		Short:   "List AutoBackup",
		Long:    `List AutoBackup`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return autoBackupListParam.Initialize(newParamsAdapter(cmd.Flags()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, autoBackupListParam)
			if err != nil {
				return err
			}

			if autoBackupListParam.GenerateSkeleton {
				return generateSkeleton(ctx, autoBackupListParam)
			}

			// TODO implements ID parameter handling

			// Run
			return funcs.AutoBackupList(ctx, autoBackupListParam.ToV0())
		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&autoBackupListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &autoBackupListParam.Id), "id", "", "set filter by id(s)")
	fs.StringSliceVarP(&autoBackupListParam.Tags, "tags", "", []string{}, "set filter by tags(AND)")
	fs.IntVarP(&autoBackupListParam.From, "from", "", 0, "set offset")
	fs.IntVarP(&autoBackupListParam.Max, "max", "", 0, "set limit")
	fs.StringSliceVarP(&autoBackupListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringVarP(&autoBackupListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&autoBackupListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&autoBackupListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&autoBackupListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&autoBackupListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&autoBackupListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&autoBackupListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&autoBackupListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&autoBackupListParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&autoBackupListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&autoBackupListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&autoBackupListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	return cmd
}

func autoBackupCreateCmd() *cobra.Command {
	autoBackupCreateParam := params.NewCreateAutoBackupParam()
	cmd := &cobra.Command{
		Use: "create",

		Short: "Create AutoBackup",
		Long:  `Create AutoBackup`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return autoBackupCreateParam.Initialize(newParamsAdapter(cmd.Flags()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, autoBackupCreateParam)
			if err != nil {
				return err
			}

			if autoBackupCreateParam.GenerateSkeleton {
				return generateSkeleton(ctx, autoBackupCreateParam)
			}

			// TODO implements ID parameter handling

			// confirm
			if !autoBackupCreateParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("create", ctx.IO().In(), ctx.IO().Out()) // TODO idハンドリング
				if err != nil {
					return err
				}
				if !result {
					return nil // canceled
				}
			}

			// Run
			return funcs.AutoBackupCreate(ctx, autoBackupCreateParam.ToV0())
		},
	}

	fs := cmd.Flags()
	fs.VarP(newIDValue(0, &autoBackupCreateParam.DiskId), "disk-id", "", "set target diskID ")
	fs.StringSliceVarP(&autoBackupCreateParam.Weekdays, "weekdays", "", []string{"all"}, "set backup target weekdays[all or mon/tue/wed/thu/fri/sat/sun]")
	fs.IntVarP(&autoBackupCreateParam.Generation, "generation", "", 1, "set backup generation[1-10]")
	fs.StringVarP(&autoBackupCreateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&autoBackupCreateParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&autoBackupCreateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &autoBackupCreateParam.IconId), "icon-id", "", "set Icon ID")
	fs.BoolVarP(&autoBackupCreateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&autoBackupCreateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&autoBackupCreateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&autoBackupCreateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&autoBackupCreateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&autoBackupCreateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&autoBackupCreateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&autoBackupCreateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&autoBackupCreateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&autoBackupCreateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&autoBackupCreateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&autoBackupCreateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&autoBackupCreateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	return cmd
}

func autoBackupReadCmd() *cobra.Command {
	autoBackupReadParam := params.NewReadAutoBackupParam()
	cmd := &cobra.Command{
		Use: "read",

		Short: "Read AutoBackup",
		Long:  `Read AutoBackup`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return autoBackupReadParam.Initialize(newParamsAdapter(cmd.Flags()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, autoBackupReadParam)
			if err != nil {
				return err
			}

			if autoBackupReadParam.GenerateSkeleton {
				return generateSkeleton(ctx, autoBackupReadParam)
			}

			// TODO implements ID parameter handling

			// Run
			return funcs.AutoBackupRead(ctx, autoBackupReadParam.ToV0())
		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&autoBackupReadParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&autoBackupReadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&autoBackupReadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&autoBackupReadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&autoBackupReadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&autoBackupReadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&autoBackupReadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&autoBackupReadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&autoBackupReadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&autoBackupReadParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&autoBackupReadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&autoBackupReadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&autoBackupReadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &autoBackupReadParam.Id), "id", "", "Set target ID")
	return cmd
}

func autoBackupUpdateCmd() *cobra.Command {
	autoBackupUpdateParam := params.NewUpdateAutoBackupParam()
	cmd := &cobra.Command{
		Use: "update",

		Short: "Update AutoBackup",
		Long:  `Update AutoBackup`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return autoBackupUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, autoBackupUpdateParam)
			if err != nil {
				return err
			}

			if autoBackupUpdateParam.GenerateSkeleton {
				return generateSkeleton(ctx, autoBackupUpdateParam)
			}

			// TODO implements ID parameter handling

			// confirm
			if !autoBackupUpdateParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("update", ctx.IO().In(), ctx.IO().Out()) // TODO idハンドリング
				if err != nil {
					return err
				}
				if !result {
					return nil // canceled
				}
			}

			// Run
			return funcs.AutoBackupUpdate(ctx, autoBackupUpdateParam.ToV0())
		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&autoBackupUpdateParam.Weekdays, "weekdays", "", []string{}, "set backup target weekdays[all or mon/tue/wed/thu/fri/sat/sun]")
	fs.IntVarP(&autoBackupUpdateParam.Generation, "generation", "", 0, "set backup generation[1-10]")
	fs.StringSliceVarP(&autoBackupUpdateParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&autoBackupUpdateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&autoBackupUpdateParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&autoBackupUpdateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &autoBackupUpdateParam.IconId), "icon-id", "", "set Icon ID")
	fs.BoolVarP(&autoBackupUpdateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&autoBackupUpdateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&autoBackupUpdateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&autoBackupUpdateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&autoBackupUpdateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&autoBackupUpdateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&autoBackupUpdateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&autoBackupUpdateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&autoBackupUpdateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&autoBackupUpdateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&autoBackupUpdateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&autoBackupUpdateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&autoBackupUpdateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &autoBackupUpdateParam.Id), "id", "", "Set target ID")
	return cmd
}

func autoBackupDeleteCmd() *cobra.Command {
	autoBackupDeleteParam := params.NewDeleteAutoBackupParam()
	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"rm"},
		Short:   "Delete AutoBackup",
		Long:    `Delete AutoBackup`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return autoBackupDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, autoBackupDeleteParam)
			if err != nil {
				return err
			}

			if autoBackupDeleteParam.GenerateSkeleton {
				return generateSkeleton(ctx, autoBackupDeleteParam)
			}

			// TODO implements ID parameter handling

			// confirm
			if !autoBackupDeleteParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("delete", ctx.IO().In(), ctx.IO().Out()) // TODO idハンドリング
				if err != nil {
					return err
				}
				if !result {
					return nil // canceled
				}
			}

			// Run
			return funcs.AutoBackupDelete(ctx, autoBackupDeleteParam.ToV0())
		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&autoBackupDeleteParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&autoBackupDeleteParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&autoBackupDeleteParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&autoBackupDeleteParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&autoBackupDeleteParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&autoBackupDeleteParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&autoBackupDeleteParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&autoBackupDeleteParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&autoBackupDeleteParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&autoBackupDeleteParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&autoBackupDeleteParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&autoBackupDeleteParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&autoBackupDeleteParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&autoBackupDeleteParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &autoBackupDeleteParam.Id), "id", "", "Set target ID")
	return cmd
}

func init() {
	parent := autoBackupCmd()
	parent.AddCommand(autoBackupListCmd())
	parent.AddCommand(autoBackupCreateCmd())
	parent.AddCommand(autoBackupReadCmd())
	parent.AddCommand(autoBackupUpdateCmd())
	parent.AddCommand(autoBackupDeleteCmd())
	rootCmd.AddCommand(parent)
}
