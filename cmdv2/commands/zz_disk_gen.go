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
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/cmdv2/params"
	"github.com/spf13/cobra"
)

// diskCmd represents the command to manage SAKURA Cloud Disk
func diskCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "disk",
		Short: "A manage commands of Disk",
		Long:  `A manage commands of Disk`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
}

func diskListCmd() *cobra.Command {
	diskListParam := params.NewListDiskParam()
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls", "find", "selector"},
		Short:   "List Disk",
		Long:    `List Disk`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return diskListParam.Initialize(newParamsAdapter(cmd.Flags()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), diskListParam)
			if err != nil {
				return err
			}

			// TODO DEBUG
			fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
			fmt.Printf("list local parameter: \n%s\n", debugMarshalIndent(diskListParam))
			return nil
		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&diskListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &diskListParam.Id), "id", "", "set filter by id(s)")
	fs.StringVarP(&diskListParam.Scope, "scope", "", "", "set filter by scope('user' or 'shared')")
	fs.StringSliceVarP(&diskListParam.Tags, "tags", "", []string{}, "set filter by tags(AND)")
	fs.VarP(newIDValue(0, &diskListParam.SourceArchiveId), "source-archive-id", "", "set filter by source-archive-id")
	fs.VarP(newIDValue(0, &diskListParam.SourceDiskId), "source-disk-id", "", "set filter by source-disk-id")
	fs.StringVarP(&diskListParam.Storage, "storage", "", "", "set filter by storage-name")
	fs.IntVarP(&diskListParam.From, "from", "", 0, "set offset")
	fs.IntVarP(&diskListParam.Max, "max", "", 0, "set limit")
	fs.StringSliceVarP(&diskListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringVarP(&diskListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&diskListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&diskListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&diskListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&diskListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&diskListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&diskListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&diskListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&diskListParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&diskListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&diskListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&diskListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	return cmd
}

func diskCreateCmd() *cobra.Command {
	diskCreateParam := params.NewCreateDiskParam()
	cmd := &cobra.Command{
		Use: "create",

		Short: "Create Disk",
		Long:  `Create Disk`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return diskCreateParam.Initialize(newParamsAdapter(cmd.Flags()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), diskCreateParam)
			if err != nil {
				return err
			}

			// TODO DEBUG
			fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
			fmt.Printf("create local parameter: \n%s\n", debugMarshalIndent(diskCreateParam))
			return nil
		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&diskCreateParam.Plan, "plan", "", "ssd", "set disk plan('hdd' or 'ssd')")
	fs.StringVarP(&diskCreateParam.Connection, "connection", "", "virtio", "set disk connection('virtio' or 'ide')")
	fs.VarP(newIDValue(0, &diskCreateParam.SourceArchiveId), "source-archive-id", "", "set source disk ID")
	fs.VarP(newIDValue(0, &diskCreateParam.SourceDiskId), "source-disk-id", "", "set source disk ID")
	fs.IntVarP(&diskCreateParam.Size, "size", "", 20, "set disk size(GB)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &diskCreateParam.DistantFrom), "distant-from", "", "set distant from disk IDs")
	fs.StringVarP(&diskCreateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&diskCreateParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&diskCreateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &diskCreateParam.IconId), "icon-id", "", "set Icon ID")
	fs.BoolVarP(&diskCreateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&diskCreateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&diskCreateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&diskCreateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&diskCreateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&diskCreateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&diskCreateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&diskCreateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&diskCreateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&diskCreateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&diskCreateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&diskCreateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&diskCreateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	return cmd
}

func diskReadCmd() *cobra.Command {
	diskReadParam := params.NewReadDiskParam()
	cmd := &cobra.Command{
		Use: "read",

		Short: "Read Disk",
		Long:  `Read Disk`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return diskReadParam.Initialize(newParamsAdapter(cmd.Flags()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), diskReadParam)
			if err != nil {
				return err
			}

			// TODO DEBUG
			fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
			fmt.Printf("read local parameter: \n%s\n", debugMarshalIndent(diskReadParam))
			return nil
		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&diskReadParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&diskReadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&diskReadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&diskReadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&diskReadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&diskReadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&diskReadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&diskReadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&diskReadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&diskReadParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&diskReadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&diskReadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&diskReadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &diskReadParam.Id), "id", "", "Set target ID")
	return cmd
}

func diskUpdateCmd() *cobra.Command {
	diskUpdateParam := params.NewUpdateDiskParam()
	cmd := &cobra.Command{
		Use: "update",

		Short: "Update Disk",
		Long:  `Update Disk`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return diskUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), diskUpdateParam)
			if err != nil {
				return err
			}

			// TODO DEBUG
			fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
			fmt.Printf("update local parameter: \n%s\n", debugMarshalIndent(diskUpdateParam))
			return nil
		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&diskUpdateParam.Connection, "connection", "", "", "set disk connection('virtio' or 'ide')")
	fs.StringSliceVarP(&diskUpdateParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&diskUpdateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&diskUpdateParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&diskUpdateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &diskUpdateParam.IconId), "icon-id", "", "set Icon ID")
	fs.BoolVarP(&diskUpdateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&diskUpdateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&diskUpdateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&diskUpdateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&diskUpdateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&diskUpdateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&diskUpdateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&diskUpdateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&diskUpdateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&diskUpdateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&diskUpdateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&diskUpdateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&diskUpdateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &diskUpdateParam.Id), "id", "", "Set target ID")
	return cmd
}

func diskDeleteCmd() *cobra.Command {
	diskDeleteParam := params.NewDeleteDiskParam()
	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"rm"},
		Short:   "Delete Disk",
		Long:    `Delete Disk`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return diskDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), diskDeleteParam)
			if err != nil {
				return err
			}

			// TODO DEBUG
			fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
			fmt.Printf("delete local parameter: \n%s\n", debugMarshalIndent(diskDeleteParam))
			return nil
		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&diskDeleteParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&diskDeleteParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&diskDeleteParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&diskDeleteParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&diskDeleteParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&diskDeleteParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&diskDeleteParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&diskDeleteParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&diskDeleteParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&diskDeleteParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&diskDeleteParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&diskDeleteParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&diskDeleteParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&diskDeleteParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &diskDeleteParam.Id), "id", "", "Set target ID")
	return cmd
}

func diskEditCmd() *cobra.Command {
	diskEditParam := params.NewEditDiskParam()
	cmd := &cobra.Command{
		Use:     "edit",
		Aliases: []string{"config"},
		Short:   "Edit Disk",
		Long:    `Edit Disk`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return diskEditParam.Initialize(newParamsAdapter(cmd.Flags()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), diskEditParam)
			if err != nil {
				return err
			}

			// TODO DEBUG
			fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
			fmt.Printf("edit local parameter: \n%s\n", debugMarshalIndent(diskEditParam))
			return nil
		},
	}

	fs := cmd.Flags()
	fs.StringVarP(&diskEditParam.Hostname, "hostname", "", "", "set hostname")
	fs.StringVarP(&diskEditParam.Password, "password", "", "", "set password")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &diskEditParam.SSHKeyIds), "ssh-key-ids", "", "set ssh-key ID(s)")
	fs.BoolVarP(&diskEditParam.DisablePasswordAuth, "disable-password-auth", "", false, "disable password auth on SSH")
	fs.StringVarP(&diskEditParam.Ipaddress, "ipaddress", "", "", "set ipaddress")
	fs.StringVarP(&diskEditParam.DefaultRoute, "default-route", "", "", "set default gateway")
	fs.IntVarP(&diskEditParam.NwMasklen, "nw-masklen", "", 24, "set ipaddress  prefix")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &diskEditParam.StartupScriptIds), "startup-script-ids", "", "set startup-script ID(s)")
	fs.StringSliceVarP(&diskEditParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&diskEditParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&diskEditParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&diskEditParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&diskEditParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&diskEditParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&diskEditParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&diskEditParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&diskEditParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&diskEditParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&diskEditParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&diskEditParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&diskEditParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&diskEditParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &diskEditParam.Id), "id", "", "Set target ID")
	return cmd
}

func diskResizePartitionCmd() *cobra.Command {
	diskResizePartitionParam := params.NewResizePartitionDiskParam()
	cmd := &cobra.Command{
		Use: "resize-partition",

		Short: "ResizePartition Disk",
		Long:  `ResizePartition Disk`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return diskResizePartitionParam.Initialize(newParamsAdapter(cmd.Flags()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), diskResizePartitionParam)
			if err != nil {
				return err
			}

			// TODO DEBUG
			fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
			fmt.Printf("resize-partition local parameter: \n%s\n", debugMarshalIndent(diskResizePartitionParam))
			return nil
		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&diskResizePartitionParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&diskResizePartitionParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&diskResizePartitionParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&diskResizePartitionParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&diskResizePartitionParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&diskResizePartitionParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&diskResizePartitionParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&diskResizePartitionParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&diskResizePartitionParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&diskResizePartitionParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&diskResizePartitionParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&diskResizePartitionParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&diskResizePartitionParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&diskResizePartitionParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &diskResizePartitionParam.Id), "id", "", "Set target ID")
	return cmd
}

func diskReinstallFromArchiveCmd() *cobra.Command {
	diskReinstallFromArchiveParam := params.NewReinstallFromArchiveDiskParam()
	cmd := &cobra.Command{
		Use: "reinstall-from-archive",

		Short: "ReinstallFromArchive Disk",
		Long:  `ReinstallFromArchive Disk`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return diskReinstallFromArchiveParam.Initialize(newParamsAdapter(cmd.Flags()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), diskReinstallFromArchiveParam)
			if err != nil {
				return err
			}

			// TODO DEBUG
			fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
			fmt.Printf("reinstall-from-archive local parameter: \n%s\n", debugMarshalIndent(diskReinstallFromArchiveParam))
			return nil
		},
	}

	fs := cmd.Flags()
	fs.VarP(newIDValue(0, &diskReinstallFromArchiveParam.SourceArchiveId), "source-archive-id", "", "set source archive ID")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &diskReinstallFromArchiveParam.DistantFrom), "distant-from", "", "set distant from disk IDs")
	fs.StringSliceVarP(&diskReinstallFromArchiveParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&diskReinstallFromArchiveParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&diskReinstallFromArchiveParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&diskReinstallFromArchiveParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&diskReinstallFromArchiveParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&diskReinstallFromArchiveParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&diskReinstallFromArchiveParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &diskReinstallFromArchiveParam.Id), "id", "", "Set target ID")
	return cmd
}

func diskReinstallFromDiskCmd() *cobra.Command {
	diskReinstallFromDiskParam := params.NewReinstallFromDiskDiskParam()
	cmd := &cobra.Command{
		Use: "reinstall-from-disk",

		Short: "ReinstallFromDisk Disk",
		Long:  `ReinstallFromDisk Disk`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return diskReinstallFromDiskParam.Initialize(newParamsAdapter(cmd.Flags()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), diskReinstallFromDiskParam)
			if err != nil {
				return err
			}

			// TODO DEBUG
			fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
			fmt.Printf("reinstall-from-disk local parameter: \n%s\n", debugMarshalIndent(diskReinstallFromDiskParam))
			return nil
		},
	}

	fs := cmd.Flags()
	fs.VarP(newIDValue(0, &diskReinstallFromDiskParam.SourceDiskId), "source-disk-id", "", "set source disk ID")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &diskReinstallFromDiskParam.DistantFrom), "distant-from", "", "set distant from disk IDs")
	fs.StringSliceVarP(&diskReinstallFromDiskParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&diskReinstallFromDiskParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&diskReinstallFromDiskParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&diskReinstallFromDiskParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&diskReinstallFromDiskParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&diskReinstallFromDiskParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&diskReinstallFromDiskParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &diskReinstallFromDiskParam.Id), "id", "", "Set target ID")
	return cmd
}

func diskReinstallToBlankCmd() *cobra.Command {
	diskReinstallToBlankParam := params.NewReinstallToBlankDiskParam()
	cmd := &cobra.Command{
		Use: "reinstall-to-blank",

		Short: "ReinstallToBlank Disk",
		Long:  `ReinstallToBlank Disk`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return diskReinstallToBlankParam.Initialize(newParamsAdapter(cmd.Flags()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), diskReinstallToBlankParam)
			if err != nil {
				return err
			}

			// TODO DEBUG
			fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
			fmt.Printf("reinstall-to-blank local parameter: \n%s\n", debugMarshalIndent(diskReinstallToBlankParam))
			return nil
		},
	}

	fs := cmd.Flags()
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &diskReinstallToBlankParam.DistantFrom), "distant-from", "", "set distant from disk IDs")
	fs.StringSliceVarP(&diskReinstallToBlankParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&diskReinstallToBlankParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&diskReinstallToBlankParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&diskReinstallToBlankParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&diskReinstallToBlankParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&diskReinstallToBlankParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&diskReinstallToBlankParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &diskReinstallToBlankParam.Id), "id", "", "Set target ID")
	return cmd
}

func diskServerConnectCmd() *cobra.Command {
	diskServerConnectParam := params.NewServerConnectDiskParam()
	cmd := &cobra.Command{
		Use: "server-connect",

		Short: "ServerConnect Disk",
		Long:  `ServerConnect Disk`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return diskServerConnectParam.Initialize(newParamsAdapter(cmd.Flags()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), diskServerConnectParam)
			if err != nil {
				return err
			}

			// TODO DEBUG
			fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
			fmt.Printf("server-connect local parameter: \n%s\n", debugMarshalIndent(diskServerConnectParam))
			return nil
		},
	}

	fs := cmd.Flags()
	fs.VarP(newIDValue(0, &diskServerConnectParam.ServerId), "server-id", "", "set target server ID")
	fs.StringSliceVarP(&diskServerConnectParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&diskServerConnectParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&diskServerConnectParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&diskServerConnectParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&diskServerConnectParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&diskServerConnectParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&diskServerConnectParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &diskServerConnectParam.Id), "id", "", "Set target ID")
	return cmd
}

func diskServerDisconnectCmd() *cobra.Command {
	diskServerDisconnectParam := params.NewServerDisconnectDiskParam()
	cmd := &cobra.Command{
		Use: "server-disconnect",

		Short: "ServerDisconnect Disk",
		Long:  `ServerDisconnect Disk`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return diskServerDisconnectParam.Initialize(newParamsAdapter(cmd.Flags()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), diskServerDisconnectParam)
			if err != nil {
				return err
			}

			// TODO DEBUG
			fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
			fmt.Printf("server-disconnect local parameter: \n%s\n", debugMarshalIndent(diskServerDisconnectParam))
			return nil
		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&diskServerDisconnectParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&diskServerDisconnectParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&diskServerDisconnectParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&diskServerDisconnectParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&diskServerDisconnectParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&diskServerDisconnectParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&diskServerDisconnectParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &diskServerDisconnectParam.Id), "id", "", "Set target ID")
	return cmd
}

func diskMonitorCmd() *cobra.Command {
	diskMonitorParam := params.NewMonitorDiskParam()
	cmd := &cobra.Command{
		Use: "monitor",

		Short: "Monitor Disk",
		Long:  `Monitor Disk`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return diskMonitorParam.Initialize(newParamsAdapter(cmd.Flags()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), diskMonitorParam)
			if err != nil {
				return err
			}

			// TODO DEBUG
			fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
			fmt.Printf("monitor local parameter: \n%s\n", debugMarshalIndent(diskMonitorParam))
			return nil
		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&diskMonitorParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&diskMonitorParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&diskMonitorParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&diskMonitorParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&diskMonitorParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&diskMonitorParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&diskMonitorParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&diskMonitorParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&diskMonitorParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&diskMonitorParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&diskMonitorParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&diskMonitorParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&diskMonitorParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.StringVarP(&diskMonitorParam.End, "end", "", "", "set end-time")
	fs.VarP(newIDValue(0, &diskMonitorParam.Id), "id", "", "Set target ID")
	fs.StringVarP(&diskMonitorParam.KeyFormat, "key-format", "", "sakuracloud.disk.{{.ID}}.disk", "set monitoring value key-format")
	fs.StringVarP(&diskMonitorParam.Start, "start", "", "", "set start-time")
	return cmd
}

func diskWaitForCopyCmd() *cobra.Command {
	diskWaitForCopyParam := params.NewWaitForCopyDiskParam()
	cmd := &cobra.Command{
		Use:     "wait-for-copy",
		Aliases: []string{"wait"},
		Short:   "WaitForCopy Disk",
		Long:    `WaitForCopy Disk`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return diskWaitForCopyParam.Initialize(newParamsAdapter(cmd.Flags()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := newCLIContext(globalFlags(), diskWaitForCopyParam)
			if err != nil {
				return err
			}

			// TODO DEBUG
			fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
			fmt.Printf("wait-for-copy local parameter: \n%s\n", debugMarshalIndent(diskWaitForCopyParam))
			return nil
		},
	}

	fs := cmd.Flags()
	fs.StringSliceVarP(&diskWaitForCopyParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&diskWaitForCopyParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&diskWaitForCopyParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&diskWaitForCopyParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&diskWaitForCopyParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&diskWaitForCopyParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &diskWaitForCopyParam.Id), "id", "", "Set target ID")
	return cmd
}

func init() {
	parent := diskCmd()
	parent.AddCommand(diskListCmd())
	parent.AddCommand(diskCreateCmd())
	parent.AddCommand(diskReadCmd())
	parent.AddCommand(diskUpdateCmd())
	parent.AddCommand(diskDeleteCmd())
	parent.AddCommand(diskEditCmd())
	parent.AddCommand(diskResizePartitionCmd())
	parent.AddCommand(diskReinstallFromArchiveCmd())
	parent.AddCommand(diskReinstallFromDiskCmd())
	parent.AddCommand(diskReinstallToBlankCmd())
	parent.AddCommand(diskServerConnectCmd())
	parent.AddCommand(diskServerDisconnectCmd())
	parent.AddCommand(diskMonitorCmd())
	parent.AddCommand(diskWaitForCopyCmd())
	rootCmd.AddCommand(parent)
}
