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
	simListParam          = params.NewListSIMParam()
	simCreateParam        = params.NewCreateSIMParam()
	simReadParam          = params.NewReadSIMParam()
	simUpdateParam        = params.NewUpdateSIMParam()
	simDeleteParam        = params.NewDeleteSIMParam()
	simCarrierInfoParam   = params.NewCarrierInfoSIMParam()
	simCarrierUpdateParam = params.NewCarrierUpdateSIMParam()
	simActivateParam      = params.NewActivateSIMParam()
	simDeactivateParam    = params.NewDeactivateSIMParam()
	simImeiLockParam      = params.NewImeiLockSIMParam()
	simIpAddParam         = params.NewIpAddSIMParam()
	simImeiUnlockParam    = params.NewImeiUnlockSIMParam()
	simIpDeleteParam      = params.NewIpDeleteSIMParam()
	simLogsParam          = params.NewLogsSIMParam()
	simMonitorParam       = params.NewMonitorSIMParam()
)

// simCmd represents the command to manage SAKURA Cloud SIM
var simCmd = &cobra.Command{
	Use:   "sim",
	Short: "A manage commands of SIM",
	Long:  `A manage commands of SIM`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

var simListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "find", "selector"},
	Short:   "List SIM",
	Long:    `List SIM`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return simListParam.Initialize(newParamsAdapter(cmd.Flags()))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := newCLIContext(globalFlags(), simListParam)
		if err != nil {
			return err
		}

		// TODO DEBUG
		fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
		fmt.Printf("list local parameter: \n%s\n", debugMarshalIndent(simListParam))
		return nil
	},
}

func simListCmdInit() {
	fs := simListCmd.Flags()
	fs.StringSliceVarP(&simListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &simListParam.Id), "id", "", "set filter by id(s)")
	fs.StringSliceVarP(&simListParam.Tags, "tags", "", []string{}, "set filter by tags(AND)")
	fs.IntVarP(&simListParam.From, "from", "", 0, "set offset")
	fs.IntVarP(&simListParam.Max, "max", "", 0, "set limit")
	fs.StringSliceVarP(&simListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringVarP(&simListParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&simListParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&simListParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&simListParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&simListParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&simListParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&simListParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&simListParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&simListParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&simListParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&simListParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&simListParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
}

var simCreateCmd = &cobra.Command{
	Use: "create",

	Short: "Create SIM",
	Long:  `Create SIM`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return simCreateParam.Initialize(newParamsAdapter(cmd.Flags()))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := newCLIContext(globalFlags(), simCreateParam)
		if err != nil {
			return err
		}

		// TODO DEBUG
		fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
		fmt.Printf("create local parameter: \n%s\n", debugMarshalIndent(simCreateParam))
		return nil
	},
}

func simCreateCmdInit() {
	fs := simCreateCmd.Flags()
	fs.StringVarP(&simCreateParam.Iccid, "iccid", "", "", "")
	fs.StringVarP(&simCreateParam.Passcode, "passcode", "", "", "")
	fs.BoolVarP(&simCreateParam.Disabled, "disabled", "", false, "")
	fs.StringVarP(&simCreateParam.Imei, "imei", "", "", "")
	fs.StringSliceVarP(&simCreateParam.Carrier, "carrier", "", []string{}, "")
	fs.StringVarP(&simCreateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&simCreateParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&simCreateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &simCreateParam.IconId), "icon-id", "", "set Icon ID")
	fs.BoolVarP(&simCreateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&simCreateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&simCreateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&simCreateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&simCreateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&simCreateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&simCreateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&simCreateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&simCreateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&simCreateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&simCreateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&simCreateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&simCreateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
}

var simReadCmd = &cobra.Command{
	Use: "read",

	Short: "Read SIM",
	Long:  `Read SIM`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return simReadParam.Initialize(newParamsAdapter(cmd.Flags()))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := newCLIContext(globalFlags(), simReadParam)
		if err != nil {
			return err
		}

		// TODO DEBUG
		fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
		fmt.Printf("read local parameter: \n%s\n", debugMarshalIndent(simReadParam))
		return nil
	},
}

func simReadCmdInit() {
	fs := simReadCmd.Flags()
	fs.StringSliceVarP(&simReadParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&simReadParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&simReadParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&simReadParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&simReadParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&simReadParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&simReadParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&simReadParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&simReadParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&simReadParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&simReadParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&simReadParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&simReadParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &simReadParam.Id), "id", "", "Set target ID")
}

var simUpdateCmd = &cobra.Command{
	Use: "update",

	Short: "Update SIM",
	Long:  `Update SIM`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return simUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := newCLIContext(globalFlags(), simUpdateParam)
		if err != nil {
			return err
		}

		// TODO DEBUG
		fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
		fmt.Printf("update local parameter: \n%s\n", debugMarshalIndent(simUpdateParam))
		return nil
	},
}

func simUpdateCmdInit() {
	fs := simUpdateCmd.Flags()
	fs.StringSliceVarP(&simUpdateParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&simUpdateParam.Name, "name", "", "", "set resource display name")
	fs.StringVarP(&simUpdateParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&simUpdateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &simUpdateParam.IconId), "icon-id", "", "set Icon ID")
	fs.BoolVarP(&simUpdateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&simUpdateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&simUpdateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&simUpdateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&simUpdateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&simUpdateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&simUpdateParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&simUpdateParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&simUpdateParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&simUpdateParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&simUpdateParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&simUpdateParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&simUpdateParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &simUpdateParam.Id), "id", "", "Set target ID")
}

var simDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"rm"},
	Short:   "Delete SIM",
	Long:    `Delete SIM`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return simDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := newCLIContext(globalFlags(), simDeleteParam)
		if err != nil {
			return err
		}

		// TODO DEBUG
		fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
		fmt.Printf("delete local parameter: \n%s\n", debugMarshalIndent(simDeleteParam))
		return nil
	},
}

func simDeleteCmdInit() {
	fs := simDeleteCmd.Flags()
	fs.BoolVarP(&simDeleteParam.Force, "force", "f", false, "forced-delete flag if SIM is still activating")
	fs.StringSliceVarP(&simDeleteParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&simDeleteParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&simDeleteParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&simDeleteParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&simDeleteParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&simDeleteParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&simDeleteParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &simDeleteParam.Id), "id", "", "Set target ID")
}

var simCarrierInfoCmd = &cobra.Command{
	Use:     "carrier-info",
	Aliases: []string{"carrier-list"},
	Short:   "CarrierInfo SIM",
	Long:    `CarrierInfo SIM`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return simCarrierInfoParam.Initialize(newParamsAdapter(cmd.Flags()))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := newCLIContext(globalFlags(), simCarrierInfoParam)
		if err != nil {
			return err
		}

		// TODO DEBUG
		fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
		fmt.Printf("carrier-info local parameter: \n%s\n", debugMarshalIndent(simCarrierInfoParam))
		return nil
	},
}

func simCarrierInfoCmdInit() {
	fs := simCarrierInfoCmd.Flags()
	fs.StringSliceVarP(&simCarrierInfoParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&simCarrierInfoParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&simCarrierInfoParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&simCarrierInfoParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&simCarrierInfoParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&simCarrierInfoParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&simCarrierInfoParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&simCarrierInfoParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&simCarrierInfoParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&simCarrierInfoParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&simCarrierInfoParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&simCarrierInfoParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&simCarrierInfoParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &simCarrierInfoParam.Id), "id", "", "Set target ID")
}

var simCarrierUpdateCmd = &cobra.Command{
	Use: "carrier-update",

	Short: "CarrierUpdate SIM",
	Long:  `CarrierUpdate SIM`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return simCarrierUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := newCLIContext(globalFlags(), simCarrierUpdateParam)
		if err != nil {
			return err
		}

		// TODO DEBUG
		fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
		fmt.Printf("carrier-update local parameter: \n%s\n", debugMarshalIndent(simCarrierUpdateParam))
		return nil
	},
}

func simCarrierUpdateCmdInit() {
	fs := simCarrierUpdateCmd.Flags()
	fs.StringSliceVarP(&simCarrierUpdateParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&simCarrierUpdateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&simCarrierUpdateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&simCarrierUpdateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&simCarrierUpdateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&simCarrierUpdateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&simCarrierUpdateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &simCarrierUpdateParam.Id), "id", "", "Set target ID")
	fs.StringSliceVarP(&simCarrierUpdateParam.Carrier, "carrier", "", []string{}, "")
}

var simActivateCmd = &cobra.Command{
	Use: "activate",

	Short: "Activate SIM",
	Long:  `Activate SIM`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return simActivateParam.Initialize(newParamsAdapter(cmd.Flags()))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := newCLIContext(globalFlags(), simActivateParam)
		if err != nil {
			return err
		}

		// TODO DEBUG
		fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
		fmt.Printf("activate local parameter: \n%s\n", debugMarshalIndent(simActivateParam))
		return nil
	},
}

func simActivateCmdInit() {
	fs := simActivateCmd.Flags()
	fs.StringSliceVarP(&simActivateParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&simActivateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&simActivateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&simActivateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&simActivateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&simActivateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&simActivateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &simActivateParam.Id), "id", "", "Set target ID")
}

var simDeactivateCmd = &cobra.Command{
	Use: "deactivate",

	Short: "Deactivate SIM",
	Long:  `Deactivate SIM`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return simDeactivateParam.Initialize(newParamsAdapter(cmd.Flags()))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := newCLIContext(globalFlags(), simDeactivateParam)
		if err != nil {
			return err
		}

		// TODO DEBUG
		fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
		fmt.Printf("deactivate local parameter: \n%s\n", debugMarshalIndent(simDeactivateParam))
		return nil
	},
}

func simDeactivateCmdInit() {
	fs := simDeactivateCmd.Flags()
	fs.StringSliceVarP(&simDeactivateParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&simDeactivateParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&simDeactivateParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&simDeactivateParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&simDeactivateParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&simDeactivateParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&simDeactivateParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &simDeactivateParam.Id), "id", "", "Set target ID")
}

var simImeiLockCmd = &cobra.Command{
	Use: "imei-lock",

	Short: "ImeiLock SIM",
	Long:  `ImeiLock SIM`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return simImeiLockParam.Initialize(newParamsAdapter(cmd.Flags()))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := newCLIContext(globalFlags(), simImeiLockParam)
		if err != nil {
			return err
		}

		// TODO DEBUG
		fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
		fmt.Printf("imei-lock local parameter: \n%s\n", debugMarshalIndent(simImeiLockParam))
		return nil
	},
}

func simImeiLockCmdInit() {
	fs := simImeiLockCmd.Flags()
	fs.StringSliceVarP(&simImeiLockParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&simImeiLockParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&simImeiLockParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&simImeiLockParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&simImeiLockParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&simImeiLockParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&simImeiLockParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &simImeiLockParam.Id), "id", "", "Set target ID")
	fs.StringVarP(&simImeiLockParam.Imei, "imei", "", "", "")
}

var simIpAddCmd = &cobra.Command{
	Use: "ip-add",

	Short: "IpAdd SIM",
	Long:  `IpAdd SIM`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return simIpAddParam.Initialize(newParamsAdapter(cmd.Flags()))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := newCLIContext(globalFlags(), simIpAddParam)
		if err != nil {
			return err
		}

		// TODO DEBUG
		fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
		fmt.Printf("ip-add local parameter: \n%s\n", debugMarshalIndent(simIpAddParam))
		return nil
	},
}

func simIpAddCmdInit() {
	fs := simIpAddCmd.Flags()
	fs.StringSliceVarP(&simIpAddParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&simIpAddParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&simIpAddParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&simIpAddParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&simIpAddParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&simIpAddParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&simIpAddParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &simIpAddParam.Id), "id", "", "Set target ID")
	fs.StringVarP(&simIpAddParam.Ip, "ip", "", "", "")
}

var simImeiUnlockCmd = &cobra.Command{
	Use: "imei-unlock",

	Short: "ImeiUnlock SIM",
	Long:  `ImeiUnlock SIM`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return simImeiUnlockParam.Initialize(newParamsAdapter(cmd.Flags()))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := newCLIContext(globalFlags(), simImeiUnlockParam)
		if err != nil {
			return err
		}

		// TODO DEBUG
		fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
		fmt.Printf("imei-unlock local parameter: \n%s\n", debugMarshalIndent(simImeiUnlockParam))
		return nil
	},
}

func simImeiUnlockCmdInit() {
	fs := simImeiUnlockCmd.Flags()
	fs.StringSliceVarP(&simImeiUnlockParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&simImeiUnlockParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&simImeiUnlockParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&simImeiUnlockParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&simImeiUnlockParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&simImeiUnlockParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&simImeiUnlockParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &simImeiUnlockParam.Id), "id", "", "Set target ID")
}

var simIpDeleteCmd = &cobra.Command{
	Use:     "ip-delete",
	Aliases: []string{"ip-del"},
	Short:   "IpDelete SIM",
	Long:    `IpDelete SIM`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return simIpDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := newCLIContext(globalFlags(), simIpDeleteParam)
		if err != nil {
			return err
		}

		// TODO DEBUG
		fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
		fmt.Printf("ip-delete local parameter: \n%s\n", debugMarshalIndent(simIpDeleteParam))
		return nil
	},
}

func simIpDeleteCmdInit() {
	fs := simIpDeleteCmd.Flags()
	fs.StringSliceVarP(&simIpDeleteParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.BoolVarP(&simIpDeleteParam.Assumeyes, "assumeyes", "y", false, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&simIpDeleteParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&simIpDeleteParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&simIpDeleteParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&simIpDeleteParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&simIpDeleteParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.VarP(newIDValue(0, &simIpDeleteParam.Id), "id", "", "Set target ID")
}

var simLogsCmd = &cobra.Command{
	Use: "logs",

	Short: "Logs SIM",
	Long:  `Logs SIM`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return simLogsParam.Initialize(newParamsAdapter(cmd.Flags()))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := newCLIContext(globalFlags(), simLogsParam)
		if err != nil {
			return err
		}

		// TODO DEBUG
		fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
		fmt.Printf("logs local parameter: \n%s\n", debugMarshalIndent(simLogsParam))
		return nil
	},
}

func simLogsCmdInit() {
	fs := simLogsCmd.Flags()
	fs.BoolVarP(&simLogsParam.Follow, "follow", "f", false, "follow log output")
	fs.Int64VarP(&simLogsParam.RefreshInterval, "refresh-interval", "", 3, "log refresh interval second")
	fs.StringSliceVarP(&simLogsParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&simLogsParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&simLogsParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&simLogsParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&simLogsParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&simLogsParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&simLogsParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&simLogsParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&simLogsParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&simLogsParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&simLogsParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&simLogsParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&simLogsParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &simLogsParam.Id), "id", "", "Set target ID")
}

var simMonitorCmd = &cobra.Command{
	Use: "monitor",

	Short: "Monitor SIM",
	Long:  `Monitor SIM`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return simMonitorParam.Initialize(newParamsAdapter(cmd.Flags()))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := newCLIContext(globalFlags(), simMonitorParam)
		if err != nil {
			return err
		}

		// TODO DEBUG
		fmt.Printf("global parameter: \n%s\n", debugMarshalIndent(ctx.Option()))
		fmt.Printf("monitor local parameter: \n%s\n", debugMarshalIndent(simMonitorParam))
		return nil
	},
}

func simMonitorCmdInit() {
	fs := simMonitorCmd.Flags()
	fs.StringVarP(&simMonitorParam.Start, "start", "", "", "set start-time")
	fs.StringVarP(&simMonitorParam.End, "end", "", "", "set end-time")
	fs.StringVarP(&simMonitorParam.KeyFormat, "key-format", "", "sakuracloud.sim.{{.ID}}", "set monitoring value key-format")
	fs.StringSliceVarP(&simMonitorParam.Selector, "selector", "", []string{}, "Set target filter by tag")
	fs.StringVarP(&simMonitorParam.ParamTemplate, "param-template", "", "", "Set input parameter from string(JSON)")
	fs.StringVarP(&simMonitorParam.Parameters, "parameters", "", "", "Set input parameters from JSON string")
	fs.StringVarP(&simMonitorParam.ParamTemplateFile, "param-template-file", "", "", "Set input parameter from file")
	fs.StringVarP(&simMonitorParam.ParameterFile, "parameter-file", "", "", "Set input parameters from file")
	fs.BoolVarP(&simMonitorParam.GenerateSkeleton, "generate-skeleton", "", false, "Output skelton of parameter JSON")
	fs.StringVarP(&simMonitorParam.OutputType, "output-type", "o", "", "Output type [table/json/csv/tsv]")
	fs.StringSliceVarP(&simMonitorParam.Column, "column", "", []string{}, "Output columns(using when '--output-type' is in [csv/tsv] only)")
	fs.BoolVarP(&simMonitorParam.Quiet, "quiet", "q", false, "Only display IDs")
	fs.StringVarP(&simMonitorParam.Format, "format", "", "", "Output format(see text/template package document for detail)")
	fs.StringVarP(&simMonitorParam.FormatFile, "format-file", "", "", "Output format from file(see text/template package document for detail)")
	fs.StringVarP(&simMonitorParam.Query, "query", "", "", "JMESPath query(using when '--output-type' is json only)")
	fs.StringVarP(&simMonitorParam.QueryFile, "query-file", "", "", "JMESPath query from file(using when '--output-type' is json only)")
	fs.VarP(newIDValue(0, &simMonitorParam.Id), "id", "", "Set target ID")
}

func init() {
	parent := simCmd

	simListCmdInit()
	parent.AddCommand(simListCmd)

	simCreateCmdInit()
	parent.AddCommand(simCreateCmd)

	simReadCmdInit()
	parent.AddCommand(simReadCmd)

	simUpdateCmdInit()
	parent.AddCommand(simUpdateCmd)

	simDeleteCmdInit()
	parent.AddCommand(simDeleteCmd)

	simCarrierInfoCmdInit()
	parent.AddCommand(simCarrierInfoCmd)

	simCarrierUpdateCmdInit()
	parent.AddCommand(simCarrierUpdateCmd)

	simActivateCmdInit()
	parent.AddCommand(simActivateCmd)

	simDeactivateCmdInit()
	parent.AddCommand(simDeactivateCmd)

	simImeiLockCmdInit()
	parent.AddCommand(simImeiLockCmd)

	simIpAddCmdInit()
	parent.AddCommand(simIpAddCmd)

	simImeiUnlockCmdInit()
	parent.AddCommand(simImeiUnlockCmd)

	simIpDeleteCmdInit()
	parent.AddCommand(simIpDeleteCmd)

	simLogsCmdInit()
	parent.AddCommand(simLogsCmd)

	simMonitorCmdInit()
	parent.AddCommand(simMonitorCmd)

	rootCmd.AddCommand(parent)
}
