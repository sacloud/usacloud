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

// packetFilterCmd represents the command to manage SAKURA Cloud PacketFilter
func packetFilterCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "packet-filter",
		Short: "A manage commands of PacketFilter",
		Long:  `A manage commands of PacketFilter`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
}

func packetFilterListCmd() *cobra.Command {
	packetFilterListParam := params.NewListPacketFilterParam()
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls", "find"},
		Short:   "List PacketFilter",
		Long:    `List PacketFilter`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return packetFilterListParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, packetFilterListParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if packetFilterListParam.GenerateSkeleton {
				return generateSkeleton(ctx, packetFilterListParam)
			}

			return funcs.PacketFilterList(ctx, packetFilterListParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&packetFilterListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &packetFilterListParam.Id), "id", "", "set filter by id(s)")
	fs.IntVarP(&packetFilterListParam.From, "from", "", 0, "set offset")
	fs.IntVarP(&packetFilterListParam.Max, "max", "", 0, "set limit")
	fs.StringSliceVarP(&packetFilterListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringVarP(&packetFilterListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&packetFilterListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&packetFilterListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&packetFilterListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&packetFilterListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&packetFilterListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&packetFilterListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&packetFilterListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&packetFilterListParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&packetFilterListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&packetFilterListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&packetFilterListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	buildFlagsUsage(cmd, packetFilterListFlagOrder(cmd))
	return cmd
}

func packetFilterCreateCmd() *cobra.Command {
	packetFilterCreateParam := params.NewCreatePacketFilterParam()
	cmd := &cobra.Command{
		Use: "create",

		Short: "Create PacketFilter",
		Long:  `Create PacketFilter`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return packetFilterCreateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, packetFilterCreateParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if packetFilterCreateParam.GenerateSkeleton {
				return generateSkeleton(ctx, packetFilterCreateParam)
			}

			// confirm
			if !packetFilterCreateParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("create", ctx.IO().In(), ctx.IO().Out())
				if err != nil || !result {
					return err
				}
			}

			return funcs.PacketFilterCreate(ctx, packetFilterCreateParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&packetFilterCreateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&packetFilterCreateParam.Description, "description", "", "", "set resource description")
	fs.BoolVarP(&packetFilterCreateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&packetFilterCreateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&packetFilterCreateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&packetFilterCreateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&packetFilterCreateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&packetFilterCreateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&packetFilterCreateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&packetFilterCreateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&packetFilterCreateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&packetFilterCreateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&packetFilterCreateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&packetFilterCreateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&packetFilterCreateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	buildFlagsUsage(cmd, packetFilterCreateFlagOrder(cmd))
	return cmd
}

func packetFilterReadCmd() *cobra.Command {
	packetFilterReadParam := params.NewReadPacketFilterParam()
	cmd := &cobra.Command{
		Use: "read",

		Short: "Read PacketFilter",
		Long:  `Read PacketFilter`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return packetFilterReadParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, packetFilterReadParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if packetFilterReadParam.GenerateSkeleton {
				return generateSkeleton(ctx, packetFilterReadParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findPacketFilterReadTargets(ctx, packetFilterReadParam)
			if err != nil {
				return err
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				packetFilterReadParam.SetId(id)
				go func(p *params.ReadPacketFilterParam) {
					err := funcs.PacketFilterRead(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(packetFilterReadParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&packetFilterReadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&packetFilterReadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&packetFilterReadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&packetFilterReadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&packetFilterReadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&packetFilterReadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&packetFilterReadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&packetFilterReadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&packetFilterReadParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&packetFilterReadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&packetFilterReadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&packetFilterReadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &packetFilterReadParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, packetFilterReadFlagOrder(cmd))
	return cmd
}

func packetFilterUpdateCmd() *cobra.Command {
	packetFilterUpdateParam := params.NewUpdatePacketFilterParam()
	cmd := &cobra.Command{
		Use: "update",

		Short: "Update PacketFilter",
		Long:  `Update PacketFilter`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return packetFilterUpdateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, packetFilterUpdateParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if packetFilterUpdateParam.GenerateSkeleton {
				return generateSkeleton(ctx, packetFilterUpdateParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findPacketFilterUpdateTargets(ctx, packetFilterUpdateParam)
			if err != nil {
				return err
			}

			// confirm
			if !packetFilterUpdateParam.Assumeyes {
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
				packetFilterUpdateParam.SetId(id)
				go func(p *params.UpdatePacketFilterParam) {
					err := funcs.PacketFilterUpdate(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(packetFilterUpdateParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&packetFilterUpdateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&packetFilterUpdateParam.Description, "description", "", "", "set resource description")
	fs.BoolVarP(&packetFilterUpdateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&packetFilterUpdateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&packetFilterUpdateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&packetFilterUpdateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&packetFilterUpdateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&packetFilterUpdateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&packetFilterUpdateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&packetFilterUpdateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&packetFilterUpdateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&packetFilterUpdateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&packetFilterUpdateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&packetFilterUpdateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&packetFilterUpdateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &packetFilterUpdateParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, packetFilterUpdateFlagOrder(cmd))
	return cmd
}

func packetFilterDeleteCmd() *cobra.Command {
	packetFilterDeleteParam := params.NewDeletePacketFilterParam()
	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"rm"},
		Short:   "Delete PacketFilter",
		Long:    `Delete PacketFilter`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return packetFilterDeleteParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, packetFilterDeleteParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if packetFilterDeleteParam.GenerateSkeleton {
				return generateSkeleton(ctx, packetFilterDeleteParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findPacketFilterDeleteTargets(ctx, packetFilterDeleteParam)
			if err != nil {
				return err
			}

			// confirm
			if !packetFilterDeleteParam.Assumeyes {
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
				packetFilterDeleteParam.SetId(id)
				go func(p *params.DeletePacketFilterParam) {
					err := funcs.PacketFilterDelete(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(packetFilterDeleteParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.BoolVarP(&packetFilterDeleteParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&packetFilterDeleteParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&packetFilterDeleteParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&packetFilterDeleteParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&packetFilterDeleteParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&packetFilterDeleteParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&packetFilterDeleteParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&packetFilterDeleteParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&packetFilterDeleteParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&packetFilterDeleteParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&packetFilterDeleteParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&packetFilterDeleteParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&packetFilterDeleteParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &packetFilterDeleteParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, packetFilterDeleteFlagOrder(cmd))
	return cmd
}

func packetFilterRuleInfoCmd() *cobra.Command {
	packetFilterRuleInfoParam := params.NewRuleInfoPacketFilterParam()
	cmd := &cobra.Command{
		Use:     "rule-info",
		Aliases: []string{"rules", "rule-list"},
		Short:   "RuleInfo PacketFilter",
		Long:    `RuleInfo PacketFilter`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return packetFilterRuleInfoParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, packetFilterRuleInfoParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if packetFilterRuleInfoParam.GenerateSkeleton {
				return generateSkeleton(ctx, packetFilterRuleInfoParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findPacketFilterRuleInfoTargets(ctx, packetFilterRuleInfoParam)
			if err != nil {
				return err
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				packetFilterRuleInfoParam.SetId(id)
				go func(p *params.RuleInfoPacketFilterParam) {
					err := funcs.PacketFilterRuleInfo(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(packetFilterRuleInfoParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&packetFilterRuleInfoParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&packetFilterRuleInfoParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&packetFilterRuleInfoParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&packetFilterRuleInfoParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&packetFilterRuleInfoParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&packetFilterRuleInfoParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&packetFilterRuleInfoParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&packetFilterRuleInfoParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&packetFilterRuleInfoParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&packetFilterRuleInfoParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&packetFilterRuleInfoParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&packetFilterRuleInfoParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &packetFilterRuleInfoParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, packetFilterRuleInfoFlagOrder(cmd))
	return cmd
}

func packetFilterRuleAddCmd() *cobra.Command {
	packetFilterRuleAddParam := params.NewRuleAddPacketFilterParam()
	cmd := &cobra.Command{
		Use: "rule-add",

		Short: "RuleAdd PacketFilter",
		Long:  `RuleAdd PacketFilter`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return packetFilterRuleAddParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, packetFilterRuleAddParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if packetFilterRuleAddParam.GenerateSkeleton {
				return generateSkeleton(ctx, packetFilterRuleAddParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findPacketFilterRuleAddTargets(ctx, packetFilterRuleAddParam)
			if err != nil {
				return err
			}

			// confirm
			if !packetFilterRuleAddParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("rule-add", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				packetFilterRuleAddParam.SetId(id)
				go func(p *params.RuleAddPacketFilterParam) {
					err := funcs.PacketFilterRuleAdd(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(packetFilterRuleAddParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.IntVarP(&packetFilterRuleAddParam.Index, "index", "", 1, "index to insert rule into")
	fs.StringVarP(&packetFilterRuleAddParam.Protocol, "protocol", "", "", "set target protocol[tcp/udp/icmp/fragment/ip]")
	fs.StringVarP(&packetFilterRuleAddParam.SourceNetwork, "source-network", "", "", "set source network[A.A.A.A] or [A.A.A.A/N (N=1..31)] or [A.A.A.A/M.M.M.M]")
	fs.StringVarP(&packetFilterRuleAddParam.SourcePort, "source-port", "", "", "set source port[N (N=0..65535)] or [N-N (N=0..65535)] or [0xPPPP/0xMMMM]")
	fs.StringVarP(&packetFilterRuleAddParam.DestinationPort, "destination-port", "", "", "set destination port[N (N=0..65535)] or [N-N (N=0..65535)] or [0xPPPP/0xMMMM]")
	fs.StringVarP(&packetFilterRuleAddParam.Action, "action", "", "", "set action[allow/deny]")
	fs.StringVarP(&packetFilterRuleAddParam.Description, "description", "", "", "set resource description")
	fs.BoolVarP(&packetFilterRuleAddParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&packetFilterRuleAddParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&packetFilterRuleAddParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&packetFilterRuleAddParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&packetFilterRuleAddParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&packetFilterRuleAddParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&packetFilterRuleAddParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&packetFilterRuleAddParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&packetFilterRuleAddParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&packetFilterRuleAddParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&packetFilterRuleAddParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&packetFilterRuleAddParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&packetFilterRuleAddParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &packetFilterRuleAddParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, packetFilterRuleAddFlagOrder(cmd))
	return cmd
}

func packetFilterRuleUpdateCmd() *cobra.Command {
	packetFilterRuleUpdateParam := params.NewRuleUpdatePacketFilterParam()
	cmd := &cobra.Command{
		Use: "rule-update",

		Short: "RuleUpdate PacketFilter",
		Long:  `RuleUpdate PacketFilter`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return packetFilterRuleUpdateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, packetFilterRuleUpdateParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if packetFilterRuleUpdateParam.GenerateSkeleton {
				return generateSkeleton(ctx, packetFilterRuleUpdateParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findPacketFilterRuleUpdateTargets(ctx, packetFilterRuleUpdateParam)
			if err != nil {
				return err
			}

			// confirm
			if !packetFilterRuleUpdateParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("rule-update", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				packetFilterRuleUpdateParam.SetId(id)
				go func(p *params.RuleUpdatePacketFilterParam) {
					err := funcs.PacketFilterRuleUpdate(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(packetFilterRuleUpdateParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.IntVarP(&packetFilterRuleUpdateParam.Index, "index", "", 0, "index of target rule")
	fs.StringVarP(&packetFilterRuleUpdateParam.Protocol, "protocol", "", "", "set target protocol[tcp/udp/icmp/fragment/ip]")
	fs.StringVarP(&packetFilterRuleUpdateParam.SourceNetwork, "source-network", "", "", "set source network[A.A.A.A] or [A.A.A.A/N (N=1..31)] or [A.A.A.A/M.M.M.M]")
	fs.StringVarP(&packetFilterRuleUpdateParam.SourcePort, "source-port", "", "", "set source port[N (N=0..65535)] or [N-N (N=0..65535)] or [0xPPPP/0xMMMM]")
	fs.StringVarP(&packetFilterRuleUpdateParam.DestinationPort, "destination-port", "", "", "set destination port[N (N=0..65535)] or [N-N (N=0..65535)] or [0xPPPP/0xMMMM]")
	fs.StringVarP(&packetFilterRuleUpdateParam.Action, "action", "", "", "set action[allow/deny]")
	fs.StringVarP(&packetFilterRuleUpdateParam.Description, "description", "", "", "set resource description")
	fs.BoolVarP(&packetFilterRuleUpdateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&packetFilterRuleUpdateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&packetFilterRuleUpdateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&packetFilterRuleUpdateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&packetFilterRuleUpdateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&packetFilterRuleUpdateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&packetFilterRuleUpdateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&packetFilterRuleUpdateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&packetFilterRuleUpdateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&packetFilterRuleUpdateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&packetFilterRuleUpdateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&packetFilterRuleUpdateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&packetFilterRuleUpdateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &packetFilterRuleUpdateParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, packetFilterRuleUpdateFlagOrder(cmd))
	return cmd
}

func packetFilterRuleDeleteCmd() *cobra.Command {
	packetFilterRuleDeleteParam := params.NewRuleDeletePacketFilterParam()
	cmd := &cobra.Command{
		Use: "rule-delete",

		Short: "RuleDelete PacketFilter",
		Long:  `RuleDelete PacketFilter`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return packetFilterRuleDeleteParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, packetFilterRuleDeleteParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if packetFilterRuleDeleteParam.GenerateSkeleton {
				return generateSkeleton(ctx, packetFilterRuleDeleteParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findPacketFilterRuleDeleteTargets(ctx, packetFilterRuleDeleteParam)
			if err != nil {
				return err
			}

			// confirm
			if !packetFilterRuleDeleteParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("delete rule", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				packetFilterRuleDeleteParam.SetId(id)
				go func(p *params.RuleDeletePacketFilterParam) {
					err := funcs.PacketFilterRuleDelete(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(packetFilterRuleDeleteParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.IntVarP(&packetFilterRuleDeleteParam.Index, "index", "", 0, "index of target rule")
	fs.BoolVarP(&packetFilterRuleDeleteParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&packetFilterRuleDeleteParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&packetFilterRuleDeleteParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&packetFilterRuleDeleteParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&packetFilterRuleDeleteParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&packetFilterRuleDeleteParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&packetFilterRuleDeleteParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&packetFilterRuleDeleteParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&packetFilterRuleDeleteParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&packetFilterRuleDeleteParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&packetFilterRuleDeleteParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&packetFilterRuleDeleteParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&packetFilterRuleDeleteParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &packetFilterRuleDeleteParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, packetFilterRuleDeleteFlagOrder(cmd))
	return cmd
}

func packetFilterInterfaceConnectCmd() *cobra.Command {
	packetFilterInterfaceConnectParam := params.NewInterfaceConnectPacketFilterParam()
	cmd := &cobra.Command{
		Use: "interface-connect",

		Short: "InterfaceConnect PacketFilter",
		Long:  `InterfaceConnect PacketFilter`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return packetFilterInterfaceConnectParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, packetFilterInterfaceConnectParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if packetFilterInterfaceConnectParam.GenerateSkeleton {
				return generateSkeleton(ctx, packetFilterInterfaceConnectParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findPacketFilterInterfaceConnectTargets(ctx, packetFilterInterfaceConnectParam)
			if err != nil {
				return err
			}

			// confirm
			if !packetFilterInterfaceConnectParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("interface-connect", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				packetFilterInterfaceConnectParam.SetId(id)
				go func(p *params.InterfaceConnectPacketFilterParam) {
					err := funcs.PacketFilterInterfaceConnect(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(packetFilterInterfaceConnectParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.VarP(newIDValue(0, &packetFilterInterfaceConnectParam.InterfaceId), "interface-id", "", "set interface ID")
	fs.BoolVarP(&packetFilterInterfaceConnectParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&packetFilterInterfaceConnectParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&packetFilterInterfaceConnectParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&packetFilterInterfaceConnectParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&packetFilterInterfaceConnectParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&packetFilterInterfaceConnectParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &packetFilterInterfaceConnectParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, packetFilterInterfaceConnectFlagOrder(cmd))
	return cmd
}

func packetFilterInterfaceDisconnectCmd() *cobra.Command {
	packetFilterInterfaceDisconnectParam := params.NewInterfaceDisconnectPacketFilterParam()
	cmd := &cobra.Command{
		Use: "interface-disconnect",

		Short: "InterfaceDisconnect PacketFilter",
		Long:  `InterfaceDisconnect PacketFilter`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return packetFilterInterfaceDisconnectParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, packetFilterInterfaceDisconnectParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if packetFilterInterfaceDisconnectParam.GenerateSkeleton {
				return generateSkeleton(ctx, packetFilterInterfaceDisconnectParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findPacketFilterInterfaceDisconnectTargets(ctx, packetFilterInterfaceDisconnectParam)
			if err != nil {
				return err
			}

			// confirm
			if !packetFilterInterfaceDisconnectParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("interface-disconnect", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				packetFilterInterfaceDisconnectParam.SetId(id)
				go func(p *params.InterfaceDisconnectPacketFilterParam) {
					err := funcs.PacketFilterInterfaceDisconnect(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(packetFilterInterfaceDisconnectParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.VarP(newIDValue(0, &packetFilterInterfaceDisconnectParam.InterfaceId), "interface-id", "", "set interface ID")
	fs.BoolVarP(&packetFilterInterfaceDisconnectParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&packetFilterInterfaceDisconnectParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&packetFilterInterfaceDisconnectParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&packetFilterInterfaceDisconnectParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&packetFilterInterfaceDisconnectParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&packetFilterInterfaceDisconnectParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &packetFilterInterfaceDisconnectParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, packetFilterInterfaceDisconnectFlagOrder(cmd))
	return cmd
}
