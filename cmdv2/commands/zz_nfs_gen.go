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

// nfsCmd represents the command to manage SAKURA Cloud NFS
func nfsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "nfs",
		Short: "A manage commands of NFS",
		Long:  `A manage commands of NFS`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
}

func nfsListCmd() *cobra.Command {
	nfsListParam := params.NewListNFSParam()
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls", "find", "selector"},
		Short:   "List NFS",
		Long:    `List NFS`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nfsListParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, nfsListParam)
			if err != nil {
				return err
			}

			if nfsListParam.GenerateSkeleton {
				return generateSkeleton(ctx, nfsListParam)
			}

			return funcs.NFSList(ctx, nfsListParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&nfsListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &nfsListParam.Id), "id", "", "set filter by id(s)")
	fs.StringSliceVarP(&nfsListParam.Tags, "tags", "", []string{}, "set filter by tags(AND)")
	fs.IntVarP(&nfsListParam.From, "from", "", 0, "set offset")
	fs.IntVarP(&nfsListParam.Max, "max", "", 0, "set limit")
	fs.StringSliceVarP(&nfsListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringVarP(&nfsListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&nfsListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&nfsListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&nfsListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&nfsListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&nfsListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&nfsListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&nfsListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&nfsListParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&nfsListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&nfsListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&nfsListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	return cmd
}

func nfsCreateCmd() *cobra.Command {
	nfsCreateParam := params.NewCreateNFSParam()
	cmd := &cobra.Command{
		Use: "create",

		Short: "Create NFS",
		Long:  `Create NFS`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nfsCreateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, nfsCreateParam)
			if err != nil {
				return err
			}

			if nfsCreateParam.GenerateSkeleton {
				return generateSkeleton(ctx, nfsCreateParam)
			}

			// confirm
			if !nfsCreateParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("create", ctx.IO().In(), ctx.IO().Out())
				if err != nil || !result {
					return err
				}
			}

			return funcs.NFSCreate(ctx, nfsCreateParam.ToV0())

		},
	}

	fs := cmd.Flags()
	fs.VarP(newIDValue(0, &nfsCreateParam.SwitchId), "switch-id", "", "set connect switch ID")
	fs.StringVarP(&nfsCreateParam.Plan, "plan", "", "hdd", "set plan[ssd/hdd]")
	fs.IntVarP(&nfsCreateParam.Size, "size", "", 100, "set plan[100/500/1024/2048/4096/8192/12288]")
	fs.StringVarP(&nfsCreateParam.Ipaddress, "ipaddress", "", "", "set ipaddress(#)")
	fs.IntVarP(&nfsCreateParam.NwMaskLen, "nw-mask-len", "", 0, "set network mask length")
	fs.StringVarP(&nfsCreateParam.DefaultRoute, "default-route", "", "", "set default route")
	fs.StringVarP(&nfsCreateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&nfsCreateParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&nfsCreateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &nfsCreateParam.IconId), "icon-id", "", "set Icon ID")
	fs.BoolVarP(&nfsCreateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&nfsCreateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&nfsCreateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&nfsCreateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&nfsCreateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&nfsCreateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&nfsCreateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&nfsCreateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&nfsCreateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&nfsCreateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&nfsCreateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&nfsCreateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&nfsCreateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	return cmd
}

func nfsReadCmd() *cobra.Command {
	nfsReadParam := params.NewReadNFSParam()
	cmd := &cobra.Command{
		Use: "read",

		Short: "Read NFS",
		Long:  `Read NFS`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nfsReadParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, nfsReadParam)
			if err != nil {
				return err
			}

			if nfsReadParam.GenerateSkeleton {
				return generateSkeleton(ctx, nfsReadParam)
			}

			// parse ID or Name arguments
			ids, err := findNFSReadTargets(ctx, nfsReadParam)
			if err != nil {
				return err
			}

			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				nfsReadParam.SetId(id)
				go func(p *params.ReadNFSParam) {
					err := funcs.NFSRead(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(nfsReadParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&nfsReadParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&nfsReadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&nfsReadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&nfsReadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&nfsReadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&nfsReadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&nfsReadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&nfsReadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&nfsReadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&nfsReadParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&nfsReadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&nfsReadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&nfsReadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &nfsReadParam.Id), "id", "", "Set target ID")
	return cmd
}

func nfsUpdateCmd() *cobra.Command {
	nfsUpdateParam := params.NewUpdateNFSParam()
	cmd := &cobra.Command{
		Use: "update",

		Short: "Update NFS",
		Long:  `Update NFS`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nfsUpdateParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, nfsUpdateParam)
			if err != nil {
				return err
			}

			if nfsUpdateParam.GenerateSkeleton {
				return generateSkeleton(ctx, nfsUpdateParam)
			}

			// parse ID or Name arguments
			ids, err := findNFSUpdateTargets(ctx, nfsUpdateParam)
			if err != nil {
				return err
			}

			// confirm
			if !nfsUpdateParam.Assumeyes {
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
				nfsUpdateParam.SetId(id)
				go func(p *params.UpdateNFSParam) {
					err := funcs.NFSUpdate(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(nfsUpdateParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&nfsUpdateParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&nfsUpdateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&nfsUpdateParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&nfsUpdateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &nfsUpdateParam.IconId), "icon-id", "", "set Icon ID")
	fs.BoolVarP(&nfsUpdateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&nfsUpdateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&nfsUpdateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&nfsUpdateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&nfsUpdateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&nfsUpdateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&nfsUpdateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&nfsUpdateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&nfsUpdateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&nfsUpdateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&nfsUpdateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&nfsUpdateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&nfsUpdateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &nfsUpdateParam.Id), "id", "", "Set target ID")
	return cmd
}

func nfsDeleteCmd() *cobra.Command {
	nfsDeleteParam := params.NewDeleteNFSParam()
	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"rm"},
		Short:   "Delete NFS",
		Long:    `Delete NFS`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nfsDeleteParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, nfsDeleteParam)
			if err != nil {
				return err
			}

			if nfsDeleteParam.GenerateSkeleton {
				return generateSkeleton(ctx, nfsDeleteParam)
			}

			// parse ID or Name arguments
			ids, err := findNFSDeleteTargets(ctx, nfsDeleteParam)
			if err != nil {
				return err
			}

			// confirm
			if !nfsDeleteParam.Assumeyes {
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
				nfsDeleteParam.SetId(id)
				go func(p *params.DeleteNFSParam) {
					err := funcs.NFSDelete(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(nfsDeleteParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.BoolVarP(&nfsDeleteParam.Force, "force", "f", false, "forced-shutdown flag if server is running")
	fs.StringSliceVarP(&nfsDeleteParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&nfsDeleteParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&nfsDeleteParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&nfsDeleteParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&nfsDeleteParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&nfsDeleteParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&nfsDeleteParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&nfsDeleteParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&nfsDeleteParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&nfsDeleteParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&nfsDeleteParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&nfsDeleteParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&nfsDeleteParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&nfsDeleteParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &nfsDeleteParam.Id), "id", "", "Set target ID")
	return cmd
}

func nfsBootCmd() *cobra.Command {
	nfsBootParam := params.NewBootNFSParam()
	cmd := &cobra.Command{
		Use:     "boot",
		Aliases: []string{"power-on"},
		Short:   "Boot NFS",
		Long:    `Boot NFS`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nfsBootParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, nfsBootParam)
			if err != nil {
				return err
			}

			if nfsBootParam.GenerateSkeleton {
				return generateSkeleton(ctx, nfsBootParam)
			}

			// parse ID or Name arguments
			ids, err := findNFSBootTargets(ctx, nfsBootParam)
			if err != nil {
				return err
			}

			// confirm
			if !nfsBootParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("boot", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				nfsBootParam.SetId(id)
				go func(p *params.BootNFSParam) {
					err := funcs.NFSBoot(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(nfsBootParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&nfsBootParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&nfsBootParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&nfsBootParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&nfsBootParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&nfsBootParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&nfsBootParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&nfsBootParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &nfsBootParam.Id), "id", "", "Set target ID")
	return cmd
}

func nfsShutdownCmd() *cobra.Command {
	nfsShutdownParam := params.NewShutdownNFSParam()
	cmd := &cobra.Command{
		Use:     "shutdown",
		Aliases: []string{"power-off"},
		Short:   "Shutdown NFS",
		Long:    `Shutdown NFS`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nfsShutdownParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, nfsShutdownParam)
			if err != nil {
				return err
			}

			if nfsShutdownParam.GenerateSkeleton {
				return generateSkeleton(ctx, nfsShutdownParam)
			}

			// parse ID or Name arguments
			ids, err := findNFSShutdownTargets(ctx, nfsShutdownParam)
			if err != nil {
				return err
			}

			// confirm
			if !nfsShutdownParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("shutdown", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				nfsShutdownParam.SetId(id)
				go func(p *params.ShutdownNFSParam) {
					err := funcs.NFSShutdown(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(nfsShutdownParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&nfsShutdownParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&nfsShutdownParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&nfsShutdownParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&nfsShutdownParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&nfsShutdownParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&nfsShutdownParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&nfsShutdownParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &nfsShutdownParam.Id), "id", "", "Set target ID")
	return cmd
}

func nfsShutdownForceCmd() *cobra.Command {
	nfsShutdownForceParam := params.NewShutdownForceNFSParam()
	cmd := &cobra.Command{
		Use:     "shutdown-force",
		Aliases: []string{"stop"},
		Short:   "ShutdownForce NFS",
		Long:    `ShutdownForce NFS`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nfsShutdownForceParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, nfsShutdownForceParam)
			if err != nil {
				return err
			}

			if nfsShutdownForceParam.GenerateSkeleton {
				return generateSkeleton(ctx, nfsShutdownForceParam)
			}

			// parse ID or Name arguments
			ids, err := findNFSShutdownForceTargets(ctx, nfsShutdownForceParam)
			if err != nil {
				return err
			}

			// confirm
			if !nfsShutdownForceParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("shutdown-force", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				nfsShutdownForceParam.SetId(id)
				go func(p *params.ShutdownForceNFSParam) {
					err := funcs.NFSShutdownForce(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(nfsShutdownForceParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&nfsShutdownForceParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&nfsShutdownForceParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&nfsShutdownForceParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&nfsShutdownForceParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&nfsShutdownForceParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&nfsShutdownForceParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&nfsShutdownForceParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &nfsShutdownForceParam.Id), "id", "", "Set target ID")
	return cmd
}

func nfsResetCmd() *cobra.Command {
	nfsResetParam := params.NewResetNFSParam()
	cmd := &cobra.Command{
		Use: "reset",

		Short: "Reset NFS",
		Long:  `Reset NFS`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nfsResetParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, nfsResetParam)
			if err != nil {
				return err
			}

			if nfsResetParam.GenerateSkeleton {
				return generateSkeleton(ctx, nfsResetParam)
			}

			// parse ID or Name arguments
			ids, err := findNFSResetTargets(ctx, nfsResetParam)
			if err != nil {
				return err
			}

			// confirm
			if !nfsResetParam.Assumeyes {
				if !utils.IsTerminal() {
					return errors.New("the confirm dialog cannot be used without the terminal. Please use --assumeyes(-y) option")
				}
				result, err := utils.ConfirmContinue("reset", ctx.IO().In(), ctx.IO().Out(), ids...)
				if err != nil || !result {
					return err
				}
			}

			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				nfsResetParam.SetId(id)
				go func(p *params.ResetNFSParam) {
					err := funcs.NFSReset(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(nfsResetParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&nfsResetParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&nfsResetParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&nfsResetParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&nfsResetParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&nfsResetParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&nfsResetParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&nfsResetParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &nfsResetParam.Id), "id", "", "Set target ID")
	return cmd
}

func nfsWaitForBootCmd() *cobra.Command {
	nfsWaitForBootParam := params.NewWaitForBootNFSParam()
	cmd := &cobra.Command{
		Use: "wait-for-boot",

		Short: "Wait until boot is completed",
		Long:  `Wait until boot is completed`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nfsWaitForBootParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, nfsWaitForBootParam)
			if err != nil {
				return err
			}

			if nfsWaitForBootParam.GenerateSkeleton {
				return generateSkeleton(ctx, nfsWaitForBootParam)
			}

			// parse ID or Name arguments
			ids, err := findNFSWaitForBootTargets(ctx, nfsWaitForBootParam)
			if err != nil {
				return err
			}

			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				nfsWaitForBootParam.SetId(id)
				go func(p *params.WaitForBootNFSParam) {
					err := funcs.NFSWaitForBoot(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(nfsWaitForBootParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&nfsWaitForBootParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&nfsWaitForBootParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&nfsWaitForBootParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&nfsWaitForBootParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&nfsWaitForBootParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&nfsWaitForBootParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &nfsWaitForBootParam.Id), "id", "", "Set target ID")
	return cmd
}

func nfsWaitForDownCmd() *cobra.Command {
	nfsWaitForDownParam := params.NewWaitForDownNFSParam()
	cmd := &cobra.Command{
		Use: "wait-for-down",

		Short: "Wait until shutdown is completed",
		Long:  `Wait until shutdown is completed`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nfsWaitForDownParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, nfsWaitForDownParam)
			if err != nil {
				return err
			}

			if nfsWaitForDownParam.GenerateSkeleton {
				return generateSkeleton(ctx, nfsWaitForDownParam)
			}

			// parse ID or Name arguments
			ids, err := findNFSWaitForDownTargets(ctx, nfsWaitForDownParam)
			if err != nil {
				return err
			}

			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				nfsWaitForDownParam.SetId(id)
				go func(p *params.WaitForDownNFSParam) {
					err := funcs.NFSWaitForDown(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(nfsWaitForDownParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&nfsWaitForDownParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&nfsWaitForDownParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&nfsWaitForDownParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&nfsWaitForDownParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&nfsWaitForDownParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&nfsWaitForDownParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &nfsWaitForDownParam.Id), "id", "", "Set target ID")
	return cmd
}

func nfsMonitorNicCmd() *cobra.Command {
	nfsMonitorNicParam := params.NewMonitorNicNFSParam()
	cmd := &cobra.Command{
		Use: "monitor-nic",

		Short: "Collect NIC(s) monitor values",
		Long:  `Collect NIC(s) monitor values`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nfsMonitorNicParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, nfsMonitorNicParam)
			if err != nil {
				return err
			}

			if nfsMonitorNicParam.GenerateSkeleton {
				return generateSkeleton(ctx, nfsMonitorNicParam)
			}

			// parse ID or Name arguments
			ids, err := findNFSMonitorNicTargets(ctx, nfsMonitorNicParam)
			if err != nil {
				return err
			}

			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				nfsMonitorNicParam.SetId(id)
				go func(p *params.MonitorNicNFSParam) {
					err := funcs.NFSMonitorNic(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(nfsMonitorNicParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&nfsMonitorNicParam.Start, "start", "", "", "set start-time")
	fs.StringVarP(&nfsMonitorNicParam.End, "end", "", "", "set end-time")
	fs.StringVarP(&nfsMonitorNicParam.KeyFormat, "key-format", "", "sakuracloud.disk.{{.ID}}.nic", "set monitoring value key-format")
	fs.StringSliceVarP(&nfsMonitorNicParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&nfsMonitorNicParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&nfsMonitorNicParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&nfsMonitorNicParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&nfsMonitorNicParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&nfsMonitorNicParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&nfsMonitorNicParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&nfsMonitorNicParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&nfsMonitorNicParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&nfsMonitorNicParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&nfsMonitorNicParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&nfsMonitorNicParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&nfsMonitorNicParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &nfsMonitorNicParam.Id), "id", "", "Set target ID")
	return cmd
}

func nfsMonitorFreeDiskSizeCmd() *cobra.Command {
	nfsMonitorFreeDiskSizeParam := params.NewMonitorFreeDiskSizeNFSParam()
	cmd := &cobra.Command{
		Use: "monitor-free-disk-size",

		Short: "Collect system-disk monitor values(IO)",
		Long:  `Collect system-disk monitor values(IO)`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return nfsMonitorFreeDiskSizeParam.Initialize(newParamsAdapter(cmd.Flags()), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), args, nfsMonitorFreeDiskSizeParam)
			if err != nil {
				return err
			}

			if nfsMonitorFreeDiskSizeParam.GenerateSkeleton {
				return generateSkeleton(ctx, nfsMonitorFreeDiskSizeParam)
			}

			// parse ID or Name arguments
			ids, err := findNFSMonitorFreeDiskSizeTargets(ctx, nfsMonitorFreeDiskSizeParam)
			if err != nil {
				return err
			}

			var wg sync.WaitGroup
			var errs []error
			for _, id := range ids {
				wg.Add(1)
				nfsMonitorFreeDiskSizeParam.SetId(id)
				go func(p *params.MonitorFreeDiskSizeNFSParam) {
					err := funcs.NFSMonitorFreeDiskSize(ctx, p.ToV0())
					if err != nil {
						errs = append(errs, err)
					}
					wg.Done()
				}(nfsMonitorFreeDiskSizeParam)
			}
			wg.Wait()
			return command.FlattenErrors(errs)

		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&nfsMonitorFreeDiskSizeParam.Start, "start", "", "", "set start-time")
	fs.StringVarP(&nfsMonitorFreeDiskSizeParam.End, "end", "", "", "set end-time")
	fs.StringVarP(&nfsMonitorFreeDiskSizeParam.KeyFormat, "key-format", "", "sakuracloud.disk.{{.ID}}.free-disk-size", "set monitoring value key-format")
	fs.StringSliceVarP(&nfsMonitorFreeDiskSizeParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&nfsMonitorFreeDiskSizeParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&nfsMonitorFreeDiskSizeParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&nfsMonitorFreeDiskSizeParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&nfsMonitorFreeDiskSizeParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&nfsMonitorFreeDiskSizeParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&nfsMonitorFreeDiskSizeParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&nfsMonitorFreeDiskSizeParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&nfsMonitorFreeDiskSizeParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&nfsMonitorFreeDiskSizeParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&nfsMonitorFreeDiskSizeParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&nfsMonitorFreeDiskSizeParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&nfsMonitorFreeDiskSizeParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &nfsMonitorFreeDiskSizeParam.Id), "id", "", "Set target ID")
	return cmd
}

func init() {
	parent := nfsCmd()
	parent.AddCommand(nfsListCmd())
	parent.AddCommand(nfsCreateCmd())
	parent.AddCommand(nfsReadCmd())
	parent.AddCommand(nfsUpdateCmd())
	parent.AddCommand(nfsDeleteCmd())
	parent.AddCommand(nfsBootCmd())
	parent.AddCommand(nfsShutdownCmd())
	parent.AddCommand(nfsShutdownForceCmd())
	parent.AddCommand(nfsResetCmd())
	parent.AddCommand(nfsWaitForBootCmd())
	parent.AddCommand(nfsWaitForDownCmd())
	parent.AddCommand(nfsMonitorNicCmd())
	parent.AddCommand(nfsMonitorFreeDiskSizeCmd())
	rootCmd.AddCommand(parent)
}
