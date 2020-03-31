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

var (
	sshKeyListParam     = params.NewListSSHKeyParam()
	sshKeyCreateParam   = params.NewCreateSSHKeyParam()
	sshKeyReadParam     = params.NewReadSSHKeyParam()
	sshKeyUpdateParam   = params.NewUpdateSSHKeyParam()
	sshKeyDeleteParam   = params.NewDeleteSSHKeyParam()
	sshKeyGenerateParam = params.NewGenerateSSHKeyParam()
)

// sshKeyCmd represents the command to manage SAKURA Cloud SSHKey
var sshKeyCmd = &cobra.Command{
	Use:   "ssh-key",
	Short: "A manage commands of SSHKey",
	Long:  `A manage commands of SSHKey`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

var sshKeyListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "find"},
	Short:   "List SSHKey",
	Long:    `List SSHKey`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return sshKeyListParam.Initialize(newParamsAdapter(cmd.Flags()))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := newCLIContext(globalFlags(), sshKeyListParam)
		if err != nil {
			return err
		}

		// TODO DEBUG
		fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
		fmt.Printf("list local parameter: \n%s\n", debugMarshalIndent(sshKeyListParam))
		return nil
	},
}

func sshKeyListCmdInit() {
	fs := sshKeyListCmd.Flags()
	fs.StringSliceVarP(&sshKeyListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &sshKeyListParam.Id), "id", "", "set filter by id(s)")
	fs.IntVarP(&sshKeyListParam.From, "from", "", 0, "set offset")
	fs.IntVarP(&sshKeyListParam.Max, "max", "", 0, "set limit")
	fs.StringSliceVarP(&sshKeyListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringVarP(&sshKeyListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&sshKeyListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&sshKeyListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&sshKeyListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&sshKeyListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&sshKeyListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&sshKeyListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&sshKeyListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&sshKeyListParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&sshKeyListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&sshKeyListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&sshKeyListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
}

var sshKeyCreateCmd = &cobra.Command{
	Use: "create",

	Short: "Create SSHKey",
	Long:  `Create SSHKey`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return sshKeyCreateParam.Initialize(newParamsAdapter(cmd.Flags()))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := newCLIContext(globalFlags(), sshKeyCreateParam)
		if err != nil {
			return err
		}

		// TODO DEBUG
		fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
		fmt.Printf("create local parameter: \n%s\n", debugMarshalIndent(sshKeyCreateParam))
		return nil
	},
}

func sshKeyCreateCmdInit() {
	fs := sshKeyCreateCmd.Flags()
	fs.StringVarP(&sshKeyCreateParam.PublicKey, "public-key", "", "", "set public-key from file")
	fs.StringVarP(&sshKeyCreateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&sshKeyCreateParam.Description, "description", "", "", "set resource description")
	fs.BoolVarP(&sshKeyCreateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&sshKeyCreateParam.PublicKeyContent, "public-key-content", "", "", "set public-key")
	fs.StringVarP(&sshKeyCreateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&sshKeyCreateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&sshKeyCreateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&sshKeyCreateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&sshKeyCreateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&sshKeyCreateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&sshKeyCreateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&sshKeyCreateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&sshKeyCreateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&sshKeyCreateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&sshKeyCreateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&sshKeyCreateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
}

var sshKeyReadCmd = &cobra.Command{
	Use: "read",

	Short: "Read SSHKey",
	Long:  `Read SSHKey`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return sshKeyReadParam.Initialize(newParamsAdapter(cmd.Flags()))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := newCLIContext(globalFlags(), sshKeyReadParam)
		if err != nil {
			return err
		}

		// TODO DEBUG
		fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
		fmt.Printf("read local parameter: \n%s\n", debugMarshalIndent(sshKeyReadParam))
		return nil
	},
}

func sshKeyReadCmdInit() {
	fs := sshKeyReadCmd.Flags()
	fs.StringVarP(&sshKeyReadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&sshKeyReadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&sshKeyReadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&sshKeyReadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&sshKeyReadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&sshKeyReadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&sshKeyReadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&sshKeyReadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&sshKeyReadParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&sshKeyReadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&sshKeyReadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&sshKeyReadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &sshKeyReadParam.Id), "id", "", "Set target ID")
}

var sshKeyUpdateCmd = &cobra.Command{
	Use: "update",

	Short: "Update SSHKey",
	Long:  `Update SSHKey`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return sshKeyUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := newCLIContext(globalFlags(), sshKeyUpdateParam)
		if err != nil {
			return err
		}

		// TODO DEBUG
		fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
		fmt.Printf("update local parameter: \n%s\n", debugMarshalIndent(sshKeyUpdateParam))
		return nil
	},
}

func sshKeyUpdateCmdInit() {
	fs := sshKeyUpdateCmd.Flags()
	fs.StringVarP(&sshKeyUpdateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&sshKeyUpdateParam.Description, "description", "", "", "set resource description")
	fs.BoolVarP(&sshKeyUpdateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&sshKeyUpdateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&sshKeyUpdateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&sshKeyUpdateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&sshKeyUpdateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&sshKeyUpdateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&sshKeyUpdateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&sshKeyUpdateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&sshKeyUpdateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&sshKeyUpdateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&sshKeyUpdateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&sshKeyUpdateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&sshKeyUpdateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &sshKeyUpdateParam.Id), "id", "", "Set target ID")
}

var sshKeyDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"rm"},
	Short:   "Delete SSHKey",
	Long:    `Delete SSHKey`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return sshKeyDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := newCLIContext(globalFlags(), sshKeyDeleteParam)
		if err != nil {
			return err
		}

		// TODO DEBUG
		fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
		fmt.Printf("delete local parameter: \n%s\n", debugMarshalIndent(sshKeyDeleteParam))
		return nil
	},
}

func sshKeyDeleteCmdInit() {
	fs := sshKeyDeleteCmd.Flags()
	fs.BoolVarP(&sshKeyDeleteParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&sshKeyDeleteParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&sshKeyDeleteParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&sshKeyDeleteParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&sshKeyDeleteParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&sshKeyDeleteParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&sshKeyDeleteParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&sshKeyDeleteParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&sshKeyDeleteParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&sshKeyDeleteParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&sshKeyDeleteParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&sshKeyDeleteParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&sshKeyDeleteParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &sshKeyDeleteParam.Id), "id", "", "Set target ID")
}

var sshKeyGenerateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen"},
	Short:   "Generate SSHKey",
	Long:    `Generate SSHKey`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return sshKeyGenerateParam.Initialize(newParamsAdapter(cmd.Flags()))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := newCLIContext(globalFlags(), sshKeyGenerateParam)
		if err != nil {
			return err
		}

		// TODO DEBUG
		fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
		fmt.Printf("generate local parameter: \n%s\n", debugMarshalIndent(sshKeyGenerateParam))
		return nil
	},
}

func sshKeyGenerateCmdInit() {
	fs := sshKeyGenerateCmd.Flags()
	fs.StringVarP(&sshKeyGenerateParam.PassPhrase, "pass-phrase", "", "", "set ssh-key pass phrase")
	fs.StringVarP(&sshKeyGenerateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&sshKeyGenerateParam.Description, "description", "", "", "set resource description")
	fs.BoolVarP(&sshKeyGenerateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&sshKeyGenerateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&sshKeyGenerateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&sshKeyGenerateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&sshKeyGenerateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&sshKeyGenerateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&sshKeyGenerateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringVarP(&sshKeyGenerateParam.PrivateKeyOutput, "private-key-output", "", "", "set ssh-key privatekey output path")
	fs.StringSliceVarP(&sshKeyGenerateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&sshKeyGenerateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&sshKeyGenerateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&sshKeyGenerateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&sshKeyGenerateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&sshKeyGenerateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
}

func init() {
	parent := sshKeyCmd

	sshKeyListCmdInit()
	parent.AddCommand(sshKeyListCmd)

	sshKeyCreateCmdInit()
	parent.AddCommand(sshKeyCreateCmd)

	sshKeyReadCmdInit()
	parent.AddCommand(sshKeyReadCmd)

	sshKeyUpdateCmdInit()
	parent.AddCommand(sshKeyUpdateCmd)

	sshKeyDeleteCmdInit()
	parent.AddCommand(sshKeyDeleteCmd)

	sshKeyGenerateCmdInit()
	parent.AddCommand(sshKeyGenerateCmd)

	rootCmd.AddCommand(parent)
}
