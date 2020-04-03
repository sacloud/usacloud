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

// switchCmd represents the command to manage SAKURA Cloud Switch
func switchCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "switch",
		Short: "A manage commands of Switch",
		Long:  `A manage commands of Switch`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
}

func switchListCmd() *cobra.Command {
	switchListParam := params.NewListSwitchParam()
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls", "find", "selector"},
		Short:   "List Switch",
		Long:    `List Switch`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return switchListParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, switchListParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if switchListParam.GenerateSkeleton {
				return generateSkeleton(ctx, switchListParam)
			}

			return funcs.SwitchList(ctx, switchListParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&switchListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &switchListParam.Id), "id", "", "set filter by id(s)")
	fs.StringSliceVarP(&switchListParam.Tags, "tags", "", []string{}, "set filter by tags(AND)")
	fs.IntVarP(&switchListParam.From, "from", "", 0, "set offset")
	fs.IntVarP(&switchListParam.Max, "max", "", 0, "set limit")
	fs.StringSliceVarP(&switchListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringVarP(&switchListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&switchListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&switchListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&switchListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&switchListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&switchListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&switchListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&switchListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&switchListParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&switchListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&switchListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&switchListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	buildFlagsUsage(cmd, switchListFlagOrder(cmd))
	return cmd
}

func switchCreateCmd() *cobra.Command {
	switchCreateParam := params.NewCreateSwitchParam()
	cmd := &cobra.Command{
		Use: "create",

		Short: "Create Switch",
		Long:  `Create Switch`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return switchCreateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, switchCreateParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if switchCreateParam.GenerateSkeleton {
				return generateSkeleton(ctx, switchCreateParam)
			}

			// confirm
			if !switchCreateParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("create", ctx.IO().In(), ctx.IO().Out())
				if err != nil || !result {
					return err
				}
			}

			return funcs.SwitchCreate(ctx, switchCreateParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&switchCreateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&switchCreateParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&switchCreateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &switchCreateParam.IconId), "icon-id", "", "set Icon ID")
	fs.BoolVarP(&switchCreateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&switchCreateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&switchCreateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&switchCreateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&switchCreateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&switchCreateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&switchCreateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&switchCreateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&switchCreateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&switchCreateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&switchCreateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&switchCreateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&switchCreateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	buildFlagsUsage(cmd, switchCreateFlagOrder(cmd))
	return cmd
}

func switchReadCmd() *cobra.Command {
	switchReadParam := params.NewReadSwitchParam()
	cmd := &cobra.Command{
		Use: "read",

		Short: "Read Switch",
		Long:  `Read Switch`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return switchReadParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, switchReadParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if switchReadParam.GenerateSkeleton {
				return generateSkeleton(ctx, switchReadParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findSwitchReadTargets(ctx, switchReadParam)
			if err != nil {
				return err
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				switchReadParam.SetId(id)
				go func(p *params.ReadSwitchParam) {
					err := funcs.SwitchRead(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(switchReadParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&switchReadParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&switchReadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&switchReadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&switchReadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&switchReadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&switchReadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&switchReadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&switchReadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&switchReadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&switchReadParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&switchReadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&switchReadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&switchReadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &switchReadParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, switchReadFlagOrder(cmd))
	return cmd
}

func switchUpdateCmd() *cobra.Command {
	switchUpdateParam := params.NewUpdateSwitchParam()
	cmd := &cobra.Command{
		Use: "update",

		Short: "Update Switch",
		Long:  `Update Switch`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return switchUpdateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, switchUpdateParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if switchUpdateParam.GenerateSkeleton {
				return generateSkeleton(ctx, switchUpdateParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findSwitchUpdateTargets(ctx, switchUpdateParam)
			if err != nil {
				return err
			}

			// confirm
			if !switchUpdateParam.Assumeyes {
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
				switchUpdateParam.SetId(id)
				go func(p *params.UpdateSwitchParam) {
					err := funcs.SwitchUpdate(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(switchUpdateParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&switchUpdateParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&switchUpdateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&switchUpdateParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&switchUpdateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &switchUpdateParam.IconId), "icon-id", "", "set Icon ID")
	fs.BoolVarP(&switchUpdateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&switchUpdateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&switchUpdateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&switchUpdateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&switchUpdateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&switchUpdateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&switchUpdateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&switchUpdateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&switchUpdateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&switchUpdateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&switchUpdateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&switchUpdateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&switchUpdateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &switchUpdateParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, switchUpdateFlagOrder(cmd))
	return cmd
}

func switchDeleteCmd() *cobra.Command {
	switchDeleteParam := params.NewDeleteSwitchParam()
	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"rm"},
		Short:   "Delete Switch",
		Long:    `Delete Switch`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return switchDeleteParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, switchDeleteParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if switchDeleteParam.GenerateSkeleton {
				return generateSkeleton(ctx, switchDeleteParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findSwitchDeleteTargets(ctx, switchDeleteParam)
			if err != nil {
				return err
			}

			// confirm
			if !switchDeleteParam.Assumeyes {
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
				switchDeleteParam.SetId(id)
				go func(p *params.DeleteSwitchParam) {
					err := funcs.SwitchDelete(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(switchDeleteParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&switchDeleteParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&switchDeleteParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&switchDeleteParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&switchDeleteParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&switchDeleteParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&switchDeleteParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&switchDeleteParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&switchDeleteParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&switchDeleteParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&switchDeleteParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&switchDeleteParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&switchDeleteParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&switchDeleteParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&switchDeleteParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &switchDeleteParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, switchDeleteFlagOrder(cmd))
	return cmd
}

func switchBridgeConnectCmd() *cobra.Command {
	switchBridgeConnectParam := params.NewBridgeConnectSwitchParam()
	cmd := &cobra.Command{
		Use: "bridge-connect",

		Short: "BridgeConnect Switch",
		Long:  `BridgeConnect Switch`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return switchBridgeConnectParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, switchBridgeConnectParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if switchBridgeConnectParam.GenerateSkeleton {
				return generateSkeleton(ctx, switchBridgeConnectParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findSwitchBridgeConnectTargets(ctx, switchBridgeConnectParam)
			if err != nil {
				return err
			}

			// confirm
			if !switchBridgeConnectParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("bridge-connect", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				switchBridgeConnectParam.SetId(id)
				go func(p *params.BridgeConnectSwitchParam) {
					err := funcs.SwitchBridgeConnect(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(switchBridgeConnectParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.VarP(newIDValue(0, &switchBridgeConnectParam.BridgeId), "bridge-id", "", "set Bridge ID")
	fs.StringSliceVarP(&switchBridgeConnectParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&switchBridgeConnectParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&switchBridgeConnectParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&switchBridgeConnectParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&switchBridgeConnectParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&switchBridgeConnectParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&switchBridgeConnectParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &switchBridgeConnectParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, switchBridgeConnectFlagOrder(cmd))
	return cmd
}

func switchBridgeDisconnectCmd() *cobra.Command {
	switchBridgeDisconnectParam := params.NewBridgeDisconnectSwitchParam()
	cmd := &cobra.Command{
		Use: "bridge-disconnect",

		Short: "BridgeDisconnect Switch",
		Long:  `BridgeDisconnect Switch`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return switchBridgeDisconnectParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, switchBridgeDisconnectParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if switchBridgeDisconnectParam.GenerateSkeleton {
				return generateSkeleton(ctx, switchBridgeDisconnectParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findSwitchBridgeDisconnectTargets(ctx, switchBridgeDisconnectParam)
			if err != nil {
				return err
			}

			// confirm
			if !switchBridgeDisconnectParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("bridge-disconnect", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			// TODO v1で置き換えるまでの暫定実装
			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				switchBridgeDisconnectParam.SetId(id)
				go func(p *params.BridgeDisconnectSwitchParam) {
					err := funcs.SwitchBridgeDisconnect(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(switchBridgeDisconnectParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&switchBridgeDisconnectParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&switchBridgeDisconnectParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&switchBridgeDisconnectParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&switchBridgeDisconnectParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&switchBridgeDisconnectParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&switchBridgeDisconnectParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&switchBridgeDisconnectParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &switchBridgeDisconnectParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, switchBridgeDisconnectFlagOrder(cmd))
	return cmd
}

func init() {
	parent := switchCmd()
	parent.AddCommand(switchListCmd())
	parent.AddCommand(switchCreateCmd())
	parent.AddCommand(switchReadCmd())
	parent.AddCommand(switchUpdateCmd())
	parent.AddCommand(switchDeleteCmd())
	parent.AddCommand(switchBridgeConnectCmd())
	parent.AddCommand(switchBridgeDisconnectCmd())
	buildCommandsUsage(parent, switchCommandOrder(parent))
	rootCmd.AddCommand(parent)
}
