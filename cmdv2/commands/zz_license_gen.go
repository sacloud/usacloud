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

// licenseCmd represents the command to manage SAKURA Cloud License
func licenseCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "license",
		Short: "A manage commands of License",
		Long:  `A manage commands of License`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
}

func licenseListCmd() *cobra.Command {
	licenseListParam := params.NewListLicenseParam()
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls", "find"},
		Short:   "List License",
		Long:    `List License`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return licenseListParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, licenseListParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if licenseListParam.GenerateSkeleton {
				return generateSkeleton(ctx, licenseListParam)
			}

			return funcs.LicenseList(ctx, licenseListParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&licenseListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &licenseListParam.Id), "id", "", "set filter by id(s)")
	fs.IntVarP(&licenseListParam.From, "from", "", 0, "set offset")
	fs.IntVarP(&licenseListParam.Max, "max", "", 0, "set limit")
	fs.StringSliceVarP(&licenseListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringVarP(&licenseListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&licenseListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&licenseListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&licenseListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&licenseListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&licenseListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&licenseListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&licenseListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&licenseListParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&licenseListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&licenseListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&licenseListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	buildFlagsUsage(cmd, licenseListFlagOrder(cmd))
	return cmd
}

func licenseCreateCmd() *cobra.Command {
	licenseCreateParam := params.NewCreateLicenseParam()
	cmd := &cobra.Command{
		Use: "create",

		Short: "Create License",
		Long:  `Create License`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return licenseCreateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, licenseCreateParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if licenseCreateParam.GenerateSkeleton {
				return generateSkeleton(ctx, licenseCreateParam)
			}

			// confirm
			if !licenseCreateParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("create", ctx.IO().In(), ctx.IO().Out())
				if err != nil || !result {
					return err
				}
			}

			return funcs.LicenseCreate(ctx, licenseCreateParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.VarP(newIDValue(0, &licenseCreateParam.LicenseInfoId), "license-info-id", "", "set LicenseInfo ID")
	fs.StringVarP(&licenseCreateParam.Name, "name", "", "", "set resource display name")
	fs.BoolVarP(&licenseCreateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&licenseCreateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&licenseCreateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&licenseCreateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&licenseCreateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&licenseCreateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&licenseCreateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&licenseCreateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&licenseCreateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&licenseCreateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&licenseCreateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&licenseCreateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&licenseCreateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	buildFlagsUsage(cmd, licenseCreateFlagOrder(cmd))
	return cmd
}

func licenseReadCmd() *cobra.Command {
	licenseReadParam := params.NewReadLicenseParam()
	cmd := &cobra.Command{
		Use: "read",

		Short: "Read License",
		Long:  `Read License`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return licenseReadParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, licenseReadParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if licenseReadParam.GenerateSkeleton {
				return generateSkeleton(ctx, licenseReadParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findLicenseReadTargets(ctx, licenseReadParam)
			if err != nil {
				return err
			}

			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				licenseReadParam.SetId(id)
				go func(p *params.ReadLicenseParam) {
					err := funcs.LicenseRead(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(licenseReadParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&licenseReadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&licenseReadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&licenseReadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&licenseReadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&licenseReadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&licenseReadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&licenseReadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&licenseReadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&licenseReadParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&licenseReadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&licenseReadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&licenseReadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &licenseReadParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, licenseReadFlagOrder(cmd))
	return cmd
}

func licenseUpdateCmd() *cobra.Command {
	licenseUpdateParam := params.NewUpdateLicenseParam()
	cmd := &cobra.Command{
		Use: "update",

		Short: "Update License",
		Long:  `Update License`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return licenseUpdateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, licenseUpdateParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if licenseUpdateParam.GenerateSkeleton {
				return generateSkeleton(ctx, licenseUpdateParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findLicenseUpdateTargets(ctx, licenseUpdateParam)
			if err != nil {
				return err
			}

			// confirm
			if !licenseUpdateParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("update", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				licenseUpdateParam.SetId(id)
				go func(p *params.UpdateLicenseParam) {
					err := funcs.LicenseUpdate(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(licenseUpdateParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&licenseUpdateParam.Name, "name", "", "", "set resource display name")
	fs.BoolVarP(&licenseUpdateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&licenseUpdateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&licenseUpdateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&licenseUpdateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&licenseUpdateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&licenseUpdateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&licenseUpdateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&licenseUpdateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&licenseUpdateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&licenseUpdateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&licenseUpdateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&licenseUpdateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&licenseUpdateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &licenseUpdateParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, licenseUpdateFlagOrder(cmd))
	return cmd
}

func licenseDeleteCmd() *cobra.Command {
	licenseDeleteParam := params.NewDeleteLicenseParam()
	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"rm"},
		Short:   "Delete License",
		Long:    `Delete License`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return licenseDeleteParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, licenseDeleteParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if licenseDeleteParam.GenerateSkeleton {
				return generateSkeleton(ctx, licenseDeleteParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findLicenseDeleteTargets(ctx, licenseDeleteParam)
			if err != nil {
				return err
			}

			// confirm
			if !licenseDeleteParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("delete", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				licenseDeleteParam.SetId(id)
				go func(p *params.DeleteLicenseParam) {
					err := funcs.LicenseDelete(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(licenseDeleteParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.BoolVarP(&licenseDeleteParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&licenseDeleteParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&licenseDeleteParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&licenseDeleteParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&licenseDeleteParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&licenseDeleteParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&licenseDeleteParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&licenseDeleteParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&licenseDeleteParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&licenseDeleteParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&licenseDeleteParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&licenseDeleteParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&licenseDeleteParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &licenseDeleteParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, licenseDeleteFlagOrder(cmd))
	return cmd
}

func init() {
	parent := licenseCmd()
	parent.AddCommand(licenseListCmd())
	parent.AddCommand(licenseCreateCmd())
	parent.AddCommand(licenseReadCmd())
	parent.AddCommand(licenseUpdateCmd())
	parent.AddCommand(licenseDeleteCmd())
	buildCommandsUsage(parent, licenseCommandOrder(parent))
	rootCmd.AddCommand(parent)
}
