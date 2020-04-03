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

// interfaceCmd represents the command to manage SAKURA Cloud Interface
func interfaceCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "interface",
		Short: "A manage commands of Interface",
		Long:  `A manage commands of Interface`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
}

func interfaceListCmd() *cobra.Command {
	interfaceListParam := params.NewListInterfaceParam()
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls", "find"},
		Short:   "List Interface",
		Long:    `List Interface`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return interfaceListParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, interfaceListParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if interfaceListParam.GenerateSkeleton {
				return generateSkeleton(ctx, interfaceListParam)
			}

			return funcs.InterfaceList(ctx, interfaceListParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&interfaceListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &interfaceListParam.Id), "id", "", "set filter by id(s)")
	fs.IntVarP(&interfaceListParam.From, "from", "", 0, "set offset")
	fs.IntVarP(&interfaceListParam.Max, "max", "", 0, "set limit")
	fs.StringSliceVarP(&interfaceListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringVarP(&interfaceListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&interfaceListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&interfaceListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&interfaceListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&interfaceListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&interfaceListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&interfaceListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&interfaceListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&interfaceListParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&interfaceListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&interfaceListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&interfaceListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	buildFlagsUsage(cmd, interfaceListFlagOrder(cmd))
	return cmd
}

func interfacePacketFilterConnectCmd() *cobra.Command {
	interfacePacketFilterConnectParam := params.NewPacketFilterConnectInterfaceParam()
	cmd := &cobra.Command{
		Use: "packet-filter-connect",

		Short: "PacketFilterConnect Interface",
		Long:  `PacketFilterConnect Interface`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return interfacePacketFilterConnectParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, interfacePacketFilterConnectParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if interfacePacketFilterConnectParam.GenerateSkeleton {
				return generateSkeleton(ctx, interfacePacketFilterConnectParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findInterfacePacketFilterConnectTargets(ctx, interfacePacketFilterConnectParam)
			if err != nil {
				return err
			}

			// confirm
			if !interfacePacketFilterConnectParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("packet-filter-connect", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				interfacePacketFilterConnectParam.SetId(id)
				go func(p *params.PacketFilterConnectInterfaceParam) {
					err := funcs.InterfacePacketFilterConnect(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(interfacePacketFilterConnectParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.VarP(newIDValue(0, &interfacePacketFilterConnectParam.PacketFilterId), "packet-filter-id", "", "set packet filter ID")
	fs.BoolVarP(&interfacePacketFilterConnectParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&interfacePacketFilterConnectParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&interfacePacketFilterConnectParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&interfacePacketFilterConnectParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&interfacePacketFilterConnectParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&interfacePacketFilterConnectParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &interfacePacketFilterConnectParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, interfacePacketFilterConnectFlagOrder(cmd))
	return cmd
}

func interfaceCreateCmd() *cobra.Command {
	interfaceCreateParam := params.NewCreateInterfaceParam()
	cmd := &cobra.Command{
		Use: "create",

		Short: "Create Interface",
		Long:  `Create Interface`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return interfaceCreateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, interfaceCreateParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if interfaceCreateParam.GenerateSkeleton {
				return generateSkeleton(ctx, interfaceCreateParam)
			}

			// confirm
			if !interfaceCreateParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("create", ctx.IO().In(), ctx.IO().Out())
				if err != nil || !result {
					return err
				}
			}

			return funcs.InterfaceCreate(ctx, interfaceCreateParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.VarP(newIDValue(0, &interfaceCreateParam.ServerId), "server-id", "", "set server ID")
	fs.BoolVarP(&interfaceCreateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&interfaceCreateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&interfaceCreateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&interfaceCreateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&interfaceCreateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&interfaceCreateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&interfaceCreateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&interfaceCreateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&interfaceCreateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&interfaceCreateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&interfaceCreateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&interfaceCreateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&interfaceCreateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	buildFlagsUsage(cmd, interfaceCreateFlagOrder(cmd))
	return cmd
}

func interfacePacketFilterDisconnectCmd() *cobra.Command {
	interfacePacketFilterDisconnectParam := params.NewPacketFilterDisconnectInterfaceParam()
	cmd := &cobra.Command{
		Use: "packet-filter-disconnect",

		Short: "PacketFilterDisconnect Interface",
		Long:  `PacketFilterDisconnect Interface`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return interfacePacketFilterDisconnectParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, interfacePacketFilterDisconnectParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if interfacePacketFilterDisconnectParam.GenerateSkeleton {
				return generateSkeleton(ctx, interfacePacketFilterDisconnectParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findInterfacePacketFilterDisconnectTargets(ctx, interfacePacketFilterDisconnectParam)
			if err != nil {
				return err
			}

			// confirm
			if !interfacePacketFilterDisconnectParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("packet-filter-disconnect", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				interfacePacketFilterDisconnectParam.SetId(id)
				go func(p *params.PacketFilterDisconnectInterfaceParam) {
					err := funcs.InterfacePacketFilterDisconnect(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(interfacePacketFilterDisconnectParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.VarP(newIDValue(0, &interfacePacketFilterDisconnectParam.PacketFilterId), "packet-filter-id", "", "set packet filter ID")
	fs.BoolVarP(&interfacePacketFilterDisconnectParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&interfacePacketFilterDisconnectParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&interfacePacketFilterDisconnectParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&interfacePacketFilterDisconnectParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&interfacePacketFilterDisconnectParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&interfacePacketFilterDisconnectParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &interfacePacketFilterDisconnectParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, interfacePacketFilterDisconnectFlagOrder(cmd))
	return cmd
}

func interfaceReadCmd() *cobra.Command {
	interfaceReadParam := params.NewReadInterfaceParam()
	cmd := &cobra.Command{
		Use: "read",

		Short: "Read Interface",
		Long:  `Read Interface`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return interfaceReadParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, interfaceReadParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if interfaceReadParam.GenerateSkeleton {
				return generateSkeleton(ctx, interfaceReadParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findInterfaceReadTargets(ctx, interfaceReadParam)
			if err != nil {
				return err
			}

			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				interfaceReadParam.SetId(id)
				go func(p *params.ReadInterfaceParam) {
					err := funcs.InterfaceRead(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(interfaceReadParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&interfaceReadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&interfaceReadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&interfaceReadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&interfaceReadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&interfaceReadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&interfaceReadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&interfaceReadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&interfaceReadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&interfaceReadParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&interfaceReadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&interfaceReadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&interfaceReadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &interfaceReadParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, interfaceReadFlagOrder(cmd))
	return cmd
}

func interfaceUpdateCmd() *cobra.Command {
	interfaceUpdateParam := params.NewUpdateInterfaceParam()
	cmd := &cobra.Command{
		Use: "update",

		Short: "Update Interface",
		Long:  `Update Interface`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return interfaceUpdateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, interfaceUpdateParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if interfaceUpdateParam.GenerateSkeleton {
				return generateSkeleton(ctx, interfaceUpdateParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findInterfaceUpdateTargets(ctx, interfaceUpdateParam)
			if err != nil {
				return err
			}

			// confirm
			if !interfaceUpdateParam.Assumeyes {
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
				interfaceUpdateParam.SetId(id)
				go func(p *params.UpdateInterfaceParam) {
					err := funcs.InterfaceUpdate(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(interfaceUpdateParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&interfaceUpdateParam.UserIpaddress, "user-ipaddress", "", "", "set user-ipaddress")
	fs.BoolVarP(&interfaceUpdateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&interfaceUpdateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&interfaceUpdateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&interfaceUpdateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&interfaceUpdateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&interfaceUpdateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&interfaceUpdateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&interfaceUpdateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&interfaceUpdateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&interfaceUpdateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&interfaceUpdateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&interfaceUpdateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&interfaceUpdateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &interfaceUpdateParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, interfaceUpdateFlagOrder(cmd))
	return cmd
}

func interfaceDeleteCmd() *cobra.Command {
	interfaceDeleteParam := params.NewDeleteInterfaceParam()
	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"rm"},
		Short:   "Delete Interface",
		Long:    `Delete Interface`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return interfaceDeleteParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, interfaceDeleteParam)
			if err != nil {
				return err
			}

			// Experiment warning
			ctx.PrintWarning("")

			if interfaceDeleteParam.GenerateSkeleton {
				return generateSkeleton(ctx, interfaceDeleteParam)
			}

			// parse ID or Name arguments(generated by tools/gen-resource-finder)
			ids, err := findInterfaceDeleteTargets(ctx, interfaceDeleteParam)
			if err != nil {
				return err
			}

			// confirm
			if !interfaceDeleteParam.Assumeyes {
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
				interfaceDeleteParam.SetId(id)
				go func(p *params.DeleteInterfaceParam) {
					err := funcs.InterfaceDelete(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(interfaceDeleteParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.BoolVarP(&interfaceDeleteParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&interfaceDeleteParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&interfaceDeleteParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&interfaceDeleteParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&interfaceDeleteParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&interfaceDeleteParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&interfaceDeleteParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&interfaceDeleteParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&interfaceDeleteParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&interfaceDeleteParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&interfaceDeleteParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&interfaceDeleteParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&interfaceDeleteParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &interfaceDeleteParam.Id), "id", "", "Set target ID")
	buildFlagsUsage(cmd, interfaceDeleteFlagOrder(cmd))
	return cmd
}

func init() {
	parent := interfaceCmd()
	parent.AddCommand(interfaceListCmd())
	parent.AddCommand(interfacePacketFilterConnectCmd())
	parent.AddCommand(interfaceCreateCmd())
	parent.AddCommand(interfacePacketFilterDisconnectCmd())
	parent.AddCommand(interfaceReadCmd())
	parent.AddCommand(interfaceUpdateCmd())
	parent.AddCommand(interfaceDeleteCmd())
	buildCommandsUsage(parent, interfaceCommandOrder(parent))
	rootCmd.AddCommand(parent)
}
