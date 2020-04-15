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

package cli

import (
	"errors"
	"sync"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/funcs"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/pkg/utils"
	"github.com/spf13/cobra"
)

// dnsCmd represents the command to manage SAKURA Cloud DNS
func dnsCmd() *cobra.Command {
	return &cobra.Command{
		Use: "dns",

		Short: "A manage commands of DNS",
		Long:  `A manage commands of DNS`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.HelpFunc()(cmd, args)
			return nil
		},
	}
}

func dnsListCmd() *cobra.Command {
	dnsListParam := params.NewListDNSParam()
	cmd := &cobra.Command{
		Use:          "list",
		Aliases:      []string{"ls", "find", "select"},
		Short:        "List DNS",
		Long:         `List DNS`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return dnsListParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, dnsListParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if dnsListParam.GenerateSkeleton {
				return generateSkeleton(ctx, dnsListParam)
			}

			return funcs.DNSList(ctx, dnsListParam)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&dnsListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &dnsListParam.Id), "id", "", "set filter by id(s)")
	fs.StringSliceVarP(&dnsListParam.Tags, "tags", "", []string{}, "set filter by tags(AND) (aliases: selector)")
	fs.IntVarP(&dnsListParam.From, "from", "", 0, "set offset (aliases: offset)")
	fs.IntVarP(&dnsListParam.Max, "max", "", 0, "set limit (aliases: limit)")
	fs.StringSliceVarP(&dnsListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringVarP(&dnsListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&dnsListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&dnsListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&dnsListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&dnsListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&dnsListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&dnsListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&dnsListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&dnsListParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&dnsListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&dnsListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&dnsListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.SetNormalizeFunc(dnsListNormalizeFlagNames)
	buildFlagsUsage(cmd, dnsListFlagOrder(cmd))
	return cmd
}

func dnsRecordInfoCmd() *cobra.Command {
	dnsRecordInfoParam := params.NewRecordInfoDNSParam()
	cmd := &cobra.Command{
		Use:          "record-info",
		Aliases:      []string{"record-list"},
		Short:        "RecordInfo DNS",
		Long:         `RecordInfo DNS`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return dnsRecordInfoParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, dnsRecordInfoParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if dnsRecordInfoParam.GenerateSkeleton {
				return generateSkeleton(ctx, dnsRecordInfoParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findDNSRecordInfoTargets(ctx, dnsRecordInfoParam)
			if err != nil {
				return err
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				dnsRecordInfoParam.SetId(id)
				go func(p *params.RecordInfoDNSParam) {
					err := funcs.DNSRecordInfo(ctx, p)
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(dnsRecordInfoParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&dnsRecordInfoParam.Name, "name", "", "", "set name")
	fs.StringVarP(&dnsRecordInfoParam.Type, "type", "", "", "set record type[A/AAAA/ALIAS/NS/CNAME/MX/TXT/SRV/CAA/PTR]")
	fs.StringSliceVarP(&dnsRecordInfoParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&dnsRecordInfoParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&dnsRecordInfoParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&dnsRecordInfoParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&dnsRecordInfoParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&dnsRecordInfoParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&dnsRecordInfoParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&dnsRecordInfoParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&dnsRecordInfoParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&dnsRecordInfoParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&dnsRecordInfoParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&dnsRecordInfoParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&dnsRecordInfoParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &dnsRecordInfoParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(dnsRecordInfoNormalizeFlagNames)
	buildFlagsUsage(cmd, dnsRecordInfoFlagOrder(cmd))
	return cmd
}

func dnsRecordBulkUpdateCmd() *cobra.Command {
	dnsRecordBulkUpdateParam := params.NewRecordBulkUpdateDNSParam()
	cmd := &cobra.Command{
		Use: "record-bulk-update",

		Short:        "RecordBulkUpdate DNS",
		Long:         `RecordBulkUpdate DNS`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return dnsRecordBulkUpdateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, dnsRecordBulkUpdateParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if dnsRecordBulkUpdateParam.GenerateSkeleton {
				return generateSkeleton(ctx, dnsRecordBulkUpdateParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findDNSRecordBulkUpdateTargets(ctx, dnsRecordBulkUpdateParam)
			if err != nil {
				return err
			}

			// confirm
			if !dnsRecordBulkUpdateParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("record-bulk-update", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				dnsRecordBulkUpdateParam.SetId(id)
				go func(p *params.RecordBulkUpdateDNSParam) {
					err := funcs.DNSRecordBulkUpdate(ctx, p)
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(dnsRecordBulkUpdateParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&dnsRecordBulkUpdateParam.File, "file", "", "", "set name")
	fs.StringVarP(&dnsRecordBulkUpdateParam.Mode, "mode", "", "upsert-only", "set name")
	fs.StringSliceVarP(&dnsRecordBulkUpdateParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&dnsRecordBulkUpdateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&dnsRecordBulkUpdateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&dnsRecordBulkUpdateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&dnsRecordBulkUpdateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&dnsRecordBulkUpdateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&dnsRecordBulkUpdateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&dnsRecordBulkUpdateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&dnsRecordBulkUpdateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&dnsRecordBulkUpdateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&dnsRecordBulkUpdateParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&dnsRecordBulkUpdateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&dnsRecordBulkUpdateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&dnsRecordBulkUpdateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &dnsRecordBulkUpdateParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(dnsRecordBulkUpdateNormalizeFlagNames)
	buildFlagsUsage(cmd, dnsRecordBulkUpdateFlagOrder(cmd))
	return cmd
}

func dnsCreateCmd() *cobra.Command {
	dnsCreateParam := params.NewCreateDNSParam()
	cmd := &cobra.Command{
		Use: "create",

		Short:        "Create DNS",
		Long:         `Create DNS`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return dnsCreateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, dnsCreateParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if dnsCreateParam.GenerateSkeleton {
				return generateSkeleton(ctx, dnsCreateParam)
			}

			// confirm
			if !dnsCreateParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("create", ctx.IO().In(), ctx.IO().Out())
				if err != nil || !result {
					return err
				}
			}

			return funcs.DNSCreate(ctx, dnsCreateParam)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&dnsCreateParam.Name, "name", "", "", "set DNS zone name")
	fs.StringVarP(&dnsCreateParam.Description, "description", "", "", "set resource description (aliases: desc)")
	fs.StringSliceVarP(&dnsCreateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &dnsCreateParam.IconId), "icon-id", "", "set Icon ID")
	fs.BoolVarP(&dnsCreateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&dnsCreateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&dnsCreateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&dnsCreateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&dnsCreateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&dnsCreateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&dnsCreateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&dnsCreateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&dnsCreateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&dnsCreateParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&dnsCreateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&dnsCreateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&dnsCreateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.SetNormalizeFunc(dnsCreateNormalizeFlagNames)
	buildFlagsUsage(cmd, dnsCreateFlagOrder(cmd))
	return cmd
}

func dnsRecordAddCmd() *cobra.Command {
	dnsRecordAddParam := params.NewRecordAddDNSParam()
	cmd := &cobra.Command{
		Use: "record-add",

		Short:        "RecordAdd DNS",
		Long:         `RecordAdd DNS`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return dnsRecordAddParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, dnsRecordAddParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if dnsRecordAddParam.GenerateSkeleton {
				return generateSkeleton(ctx, dnsRecordAddParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findDNSRecordAddTargets(ctx, dnsRecordAddParam)
			if err != nil {
				return err
			}

			// confirm
			if !dnsRecordAddParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("record-add", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				dnsRecordAddParam.SetId(id)
				go func(p *params.RecordAddDNSParam) {
					err := funcs.DNSRecordAdd(ctx, p)
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(dnsRecordAddParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&dnsRecordAddParam.Name, "name", "", "", "set name")
	fs.StringVarP(&dnsRecordAddParam.Type, "type", "", "", "set record type[A/AAAA/ALIAS/NS/CNAME/MX/TXT/SRV/CAA/PTR]")
	fs.StringVarP(&dnsRecordAddParam.Value, "value", "", "", "set record data")
	fs.IntVarP(&dnsRecordAddParam.Ttl, "ttl", "", 3600, "set ttl")
	fs.IntVarP(&dnsRecordAddParam.MxPriority, "mx-priority", "", 10, "set MX priority")
	fs.IntVarP(&dnsRecordAddParam.SrvPriority, "srv-priority", "", 0, "set SRV priority")
	fs.IntVarP(&dnsRecordAddParam.SrvWeight, "srv-weight", "", 0, "set SRV priority")
	fs.IntVarP(&dnsRecordAddParam.SrvPort, "srv-port", "", 0, "set SRV priority")
	fs.StringVarP(&dnsRecordAddParam.SrvTarget, "srv-target", "", "", "set SRV priority")
	fs.StringSliceVarP(&dnsRecordAddParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&dnsRecordAddParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&dnsRecordAddParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&dnsRecordAddParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&dnsRecordAddParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&dnsRecordAddParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&dnsRecordAddParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&dnsRecordAddParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&dnsRecordAddParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&dnsRecordAddParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&dnsRecordAddParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&dnsRecordAddParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&dnsRecordAddParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&dnsRecordAddParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &dnsRecordAddParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(dnsRecordAddNormalizeFlagNames)
	buildFlagsUsage(cmd, dnsRecordAddFlagOrder(cmd))
	return cmd
}

func dnsReadCmd() *cobra.Command {
	dnsReadParam := params.NewReadDNSParam()
	cmd := &cobra.Command{
		Use: "read",

		Short:        "Read DNS",
		Long:         `Read DNS`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return dnsReadParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, dnsReadParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if dnsReadParam.GenerateSkeleton {
				return generateSkeleton(ctx, dnsReadParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findDNSReadTargets(ctx, dnsReadParam)
			if err != nil {
				return err
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				dnsReadParam.SetId(id)
				go func(p *params.ReadDNSParam) {
					err := funcs.DNSRead(ctx, p)
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(dnsReadParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&dnsReadParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&dnsReadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&dnsReadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&dnsReadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&dnsReadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&dnsReadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&dnsReadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&dnsReadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&dnsReadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&dnsReadParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&dnsReadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&dnsReadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&dnsReadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &dnsReadParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(dnsReadNormalizeFlagNames)
	buildFlagsUsage(cmd, dnsReadFlagOrder(cmd))
	return cmd
}

func dnsRecordUpdateCmd() *cobra.Command {
	dnsRecordUpdateParam := params.NewRecordUpdateDNSParam()
	cmd := &cobra.Command{
		Use: "record-update",

		Short:        "RecordUpdate DNS",
		Long:         `RecordUpdate DNS`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return dnsRecordUpdateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, dnsRecordUpdateParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if dnsRecordUpdateParam.GenerateSkeleton {
				return generateSkeleton(ctx, dnsRecordUpdateParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findDNSRecordUpdateTargets(ctx, dnsRecordUpdateParam)
			if err != nil {
				return err
			}

			// confirm
			if !dnsRecordUpdateParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("record-update", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				dnsRecordUpdateParam.SetId(id)
				go func(p *params.RecordUpdateDNSParam) {
					err := funcs.DNSRecordUpdate(ctx, p)
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(dnsRecordUpdateParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.IntVarP(&dnsRecordUpdateParam.Index, "index", "", 0, "index of target record")
	fs.StringVarP(&dnsRecordUpdateParam.Name, "name", "", "", "set name")
	fs.StringVarP(&dnsRecordUpdateParam.Type, "type", "", "", "set record type[A/AAAA/ALIAS/NS/CNAME/MX/TXT/SRV/CAA/PTR]")
	fs.StringVarP(&dnsRecordUpdateParam.Value, "value", "", "", "set record data")
	fs.IntVarP(&dnsRecordUpdateParam.Ttl, "ttl", "", 0, "set ttl")
	fs.IntVarP(&dnsRecordUpdateParam.MxPriority, "mx-priority", "", 0, "set MX priority")
	fs.IntVarP(&dnsRecordUpdateParam.SrvPriority, "srv-priority", "", 0, "set SRV priority")
	fs.IntVarP(&dnsRecordUpdateParam.SrvWeight, "srv-weight", "", 0, "set SRV priority")
	fs.IntVarP(&dnsRecordUpdateParam.SrvPort, "srv-port", "", 0, "set SRV priority")
	fs.StringVarP(&dnsRecordUpdateParam.SrvTarget, "srv-target", "", "", "set SRV priority")
	fs.StringSliceVarP(&dnsRecordUpdateParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&dnsRecordUpdateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&dnsRecordUpdateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&dnsRecordUpdateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&dnsRecordUpdateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&dnsRecordUpdateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&dnsRecordUpdateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&dnsRecordUpdateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&dnsRecordUpdateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&dnsRecordUpdateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&dnsRecordUpdateParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&dnsRecordUpdateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&dnsRecordUpdateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&dnsRecordUpdateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &dnsRecordUpdateParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(dnsRecordUpdateNormalizeFlagNames)
	buildFlagsUsage(cmd, dnsRecordUpdateFlagOrder(cmd))
	return cmd
}

func dnsRecordDeleteCmd() *cobra.Command {
	dnsRecordDeleteParam := params.NewRecordDeleteDNSParam()
	cmd := &cobra.Command{
		Use: "record-delete",

		Short:        "RecordDelete DNS",
		Long:         `RecordDelete DNS`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return dnsRecordDeleteParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, dnsRecordDeleteParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if dnsRecordDeleteParam.GenerateSkeleton {
				return generateSkeleton(ctx, dnsRecordDeleteParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findDNSRecordDeleteTargets(ctx, dnsRecordDeleteParam)
			if err != nil {
				return err
			}

			// confirm
			if !dnsRecordDeleteParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("delete record", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				dnsRecordDeleteParam.SetId(id)
				go func(p *params.RecordDeleteDNSParam) {
					err := funcs.DNSRecordDelete(ctx, p)
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(dnsRecordDeleteParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.IntVarP(&dnsRecordDeleteParam.Index, "index", "", 0, "index of target record")
	fs.StringSliceVarP(&dnsRecordDeleteParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&dnsRecordDeleteParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&dnsRecordDeleteParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&dnsRecordDeleteParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&dnsRecordDeleteParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&dnsRecordDeleteParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&dnsRecordDeleteParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&dnsRecordDeleteParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&dnsRecordDeleteParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&dnsRecordDeleteParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&dnsRecordDeleteParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&dnsRecordDeleteParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&dnsRecordDeleteParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&dnsRecordDeleteParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &dnsRecordDeleteParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(dnsRecordDeleteNormalizeFlagNames)
	buildFlagsUsage(cmd, dnsRecordDeleteFlagOrder(cmd))
	return cmd
}

func dnsUpdateCmd() *cobra.Command {
	dnsUpdateParam := params.NewUpdateDNSParam()
	cmd := &cobra.Command{
		Use: "update",

		Short:        "Update DNS",
		Long:         `Update DNS`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return dnsUpdateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, dnsUpdateParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if dnsUpdateParam.GenerateSkeleton {
				return generateSkeleton(ctx, dnsUpdateParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findDNSUpdateTargets(ctx, dnsUpdateParam)
			if err != nil {
				return err
			}

			// confirm
			if !dnsUpdateParam.Assumeyes {
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
				dnsUpdateParam.SetId(id)
				go func(p *params.UpdateDNSParam) {
					err := funcs.DNSUpdate(ctx, p)
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(dnsUpdateParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&dnsUpdateParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&dnsUpdateParam.Description, "description", "", "", "set resource description (aliases: desc)")
	fs.StringSliceVarP(&dnsUpdateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &dnsUpdateParam.IconId), "icon-id", "", "set Icon ID")
	fs.BoolVarP(&dnsUpdateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&dnsUpdateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&dnsUpdateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&dnsUpdateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&dnsUpdateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&dnsUpdateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&dnsUpdateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&dnsUpdateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&dnsUpdateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&dnsUpdateParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&dnsUpdateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&dnsUpdateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&dnsUpdateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &dnsUpdateParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(dnsUpdateNormalizeFlagNames)
	buildFlagsUsage(cmd, dnsUpdateFlagOrder(cmd))
	return cmd
}

func dnsDeleteCmd() *cobra.Command {
	dnsDeleteParam := params.NewDeleteDNSParam()
	cmd := &cobra.Command{
		Use:          "delete",
		Aliases:      []string{"rm"},
		Short:        "Delete DNS",
		Long:         `Delete DNS`,
		SilenceUsage: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return dnsDeleteParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, dnsDeleteParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if dnsDeleteParam.GenerateSkeleton {
				return generateSkeleton(ctx, dnsDeleteParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findDNSDeleteTargets(ctx, dnsDeleteParam)
			if err != nil {
				return err
			}

			// confirm
			if !dnsDeleteParam.Assumeyes {
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
				dnsDeleteParam.SetId(id)
				go func(p *params.DeleteDNSParam) {
					err := funcs.DNSDelete(ctx, p)
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(dnsDeleteParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&dnsDeleteParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&dnsDeleteParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&dnsDeleteParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&dnsDeleteParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&dnsDeleteParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&dnsDeleteParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&dnsDeleteParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&dnsDeleteParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv] (aliases: out)")
	fs.StringSliceVarP(&dnsDeleteParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only) (aliases: col)")
	fs.BoolVarP(&dnsDeleteParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&dnsDeleteParam.Format, "format", "", "", "Output format(see text/template package document for detail) (aliases: fmt)")
	fs.StringVarP(&dnsDeleteParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&dnsDeleteParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&dnsDeleteParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &dnsDeleteParam.Id), "id", "", "Set target ID")
	fs.SetNormalizeFunc(dnsDeleteNormalizeFlagNames)
	buildFlagsUsage(cmd, dnsDeleteFlagOrder(cmd))
	return cmd
}