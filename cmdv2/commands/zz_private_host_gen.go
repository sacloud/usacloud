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
	privateHostListParam         = params.NewListPrivateHostParam()
	privateHostCreateParam       = params.NewCreatePrivateHostParam()
	privateHostReadParam         = params.NewReadPrivateHostParam()
	privateHostUpdateParam       = params.NewUpdatePrivateHostParam()
	privateHostDeleteParam       = params.NewDeletePrivateHostParam()
	privateHostServerInfoParam   = params.NewServerInfoPrivateHostParam()
	privateHostServerAddParam    = params.NewServerAddPrivateHostParam()
	privateHostServerDeleteParam = params.NewServerDeletePrivateHostParam()
)

// privateHostCmd represents the command to manage SAKURA Cloud PrivateHost
var privateHostCmd = &cobra.Command{
	Use:   "private-host",
	Short: "A manage commands of PrivateHost",
	Long:  `A manage commands of PrivateHost`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

var privateHostListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "find", "selector"},
	Short:   "List PrivateHost",
	Long:    `List PrivateHost`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := privateHostListParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("list parameter: \n%s\n", debugMarshalIndent(privateHostListParam))
		return err
	},
}

func privateHostListCmdInit() {
	fs := privateHostListCmd.Flags()
	fs.StringSliceVarP(&privateHostListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &privateHostListParam.Id), "id", "", "set filter by id(s)")
	fs.StringSliceVarP(&privateHostListParam.Tags, "tags", "", []string{}, "set filter by tags(AND)")
	fs.IntVarP(&privateHostListParam.From, "from", "", 0, "set offset")
	fs.IntVarP(&privateHostListParam.Max, "max", "", 0, "set limit")
	fs.StringSliceVarP(&privateHostListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringVarP(&privateHostListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&privateHostListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&privateHostListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&privateHostListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&privateHostListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&privateHostListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&privateHostListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&privateHostListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&privateHostListParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&privateHostListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&privateHostListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&privateHostListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
}

var privateHostCreateCmd = &cobra.Command{
	Use: "create",

	Short: "Create PrivateHost",
	Long:  `Create PrivateHost`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := privateHostCreateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("create parameter: \n%s\n", debugMarshalIndent(privateHostCreateParam))
		return err
	},
}

func privateHostCreateCmdInit() {
	fs := privateHostCreateCmd.Flags()
	fs.StringVarP(&privateHostCreateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&privateHostCreateParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&privateHostCreateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &privateHostCreateParam.IconId), "icon-id", "", "set Icon ID")
	fs.BoolVarP(&privateHostCreateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&privateHostCreateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&privateHostCreateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&privateHostCreateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&privateHostCreateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&privateHostCreateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&privateHostCreateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&privateHostCreateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&privateHostCreateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&privateHostCreateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&privateHostCreateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&privateHostCreateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&privateHostCreateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
}

var privateHostReadCmd = &cobra.Command{
	Use: "read",

	Short: "Read PrivateHost",
	Long:  `Read PrivateHost`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := privateHostReadParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("read parameter: \n%s\n", debugMarshalIndent(privateHostReadParam))
		return err
	},
}

func privateHostReadCmdInit() {
	fs := privateHostReadCmd.Flags()
	fs.StringSliceVarP(&privateHostReadParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&privateHostReadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&privateHostReadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&privateHostReadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&privateHostReadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&privateHostReadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&privateHostReadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&privateHostReadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&privateHostReadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&privateHostReadParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&privateHostReadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&privateHostReadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&privateHostReadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &privateHostReadParam.Id), "id", "", "Set target ID")
}

var privateHostUpdateCmd = &cobra.Command{
	Use: "update",

	Short: "Update PrivateHost",
	Long:  `Update PrivateHost`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := privateHostUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("update parameter: \n%s\n", debugMarshalIndent(privateHostUpdateParam))
		return err
	},
}

func privateHostUpdateCmdInit() {
	fs := privateHostUpdateCmd.Flags()
	fs.StringSliceVarP(&privateHostUpdateParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&privateHostUpdateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&privateHostUpdateParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&privateHostUpdateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &privateHostUpdateParam.IconId), "icon-id", "", "set Icon ID")
	fs.BoolVarP(&privateHostUpdateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&privateHostUpdateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&privateHostUpdateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&privateHostUpdateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&privateHostUpdateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&privateHostUpdateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&privateHostUpdateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&privateHostUpdateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&privateHostUpdateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&privateHostUpdateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&privateHostUpdateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&privateHostUpdateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&privateHostUpdateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &privateHostUpdateParam.Id), "id", "", "Set target ID")
}

var privateHostDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"rm"},
	Short:   "Delete PrivateHost",
	Long:    `Delete PrivateHost`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := privateHostDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("delete parameter: \n%s\n", debugMarshalIndent(privateHostDeleteParam))
		return err
	},
}

func privateHostDeleteCmdInit() {
	fs := privateHostDeleteCmd.Flags()
	fs.StringSliceVarP(&privateHostDeleteParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&privateHostDeleteParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&privateHostDeleteParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&privateHostDeleteParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&privateHostDeleteParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&privateHostDeleteParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&privateHostDeleteParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&privateHostDeleteParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&privateHostDeleteParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&privateHostDeleteParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&privateHostDeleteParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&privateHostDeleteParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&privateHostDeleteParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&privateHostDeleteParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &privateHostDeleteParam.Id), "id", "", "Set target ID")
}

var privateHostServerInfoCmd = &cobra.Command{
	Use:     "server-info",
	Aliases: []string{"server-list"},
	Short:   "ServerInfo PrivateHost",
	Long:    `ServerInfo PrivateHost`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := privateHostServerInfoParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("server-info parameter: \n%s\n", debugMarshalIndent(privateHostServerInfoParam))
		return err
	},
}

func privateHostServerInfoCmdInit() {
	fs := privateHostServerInfoCmd.Flags()
	fs.StringSliceVarP(&privateHostServerInfoParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&privateHostServerInfoParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&privateHostServerInfoParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&privateHostServerInfoParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&privateHostServerInfoParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&privateHostServerInfoParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&privateHostServerInfoParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&privateHostServerInfoParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&privateHostServerInfoParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&privateHostServerInfoParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&privateHostServerInfoParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&privateHostServerInfoParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&privateHostServerInfoParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &privateHostServerInfoParam.Id), "id", "", "Set target ID")
}

var privateHostServerAddCmd = &cobra.Command{
	Use: "server-add",

	Short: "ServerAdd PrivateHost",
	Long:  `ServerAdd PrivateHost`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := privateHostServerAddParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("server-add parameter: \n%s\n", debugMarshalIndent(privateHostServerAddParam))
		return err
	},
}

func privateHostServerAddCmdInit() {
	fs := privateHostServerAddCmd.Flags()
	fs.VarP(newIDValue(0, &privateHostServerAddParam.ServerId), "server-id", "", "set server ID")
	fs.StringSliceVarP(&privateHostServerAddParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&privateHostServerAddParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&privateHostServerAddParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&privateHostServerAddParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&privateHostServerAddParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&privateHostServerAddParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&privateHostServerAddParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&privateHostServerAddParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&privateHostServerAddParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&privateHostServerAddParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&privateHostServerAddParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&privateHostServerAddParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&privateHostServerAddParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&privateHostServerAddParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &privateHostServerAddParam.Id), "id", "", "Set target ID")
}

var privateHostServerDeleteCmd = &cobra.Command{
	Use: "server-delete",

	Short: "ServerDelete PrivateHost",
	Long:  `ServerDelete PrivateHost`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := privateHostServerDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("server-delete parameter: \n%s\n", debugMarshalIndent(privateHostServerDeleteParam))
		return err
	},
}

func privateHostServerDeleteCmdInit() {
	fs := privateHostServerDeleteCmd.Flags()
	fs.VarP(newIDValue(0, &privateHostServerDeleteParam.ServerId), "server-id", "", "set server ID")
	fs.StringSliceVarP(&privateHostServerDeleteParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&privateHostServerDeleteParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&privateHostServerDeleteParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&privateHostServerDeleteParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&privateHostServerDeleteParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&privateHostServerDeleteParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&privateHostServerDeleteParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&privateHostServerDeleteParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&privateHostServerDeleteParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&privateHostServerDeleteParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&privateHostServerDeleteParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&privateHostServerDeleteParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&privateHostServerDeleteParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&privateHostServerDeleteParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &privateHostServerDeleteParam.Id), "id", "", "Set target ID")
}

func init() {
	parent := privateHostCmd

	privateHostListCmdInit()
	parent.AddCommand(privateHostListCmd)

	privateHostCreateCmdInit()
	parent.AddCommand(privateHostCreateCmd)

	privateHostReadCmdInit()
	parent.AddCommand(privateHostReadCmd)

	privateHostUpdateCmdInit()
	parent.AddCommand(privateHostUpdateCmd)

	privateHostDeleteCmdInit()
	parent.AddCommand(privateHostDeleteCmd)

	privateHostServerInfoCmdInit()
	parent.AddCommand(privateHostServerInfoCmd)

	privateHostServerAddCmdInit()
	parent.AddCommand(privateHostServerAddCmd)

	privateHostServerDeleteCmdInit()
	parent.AddCommand(privateHostServerDeleteCmd)

	rootCmd.AddCommand(parent)
}
