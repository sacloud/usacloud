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
	dnsListParam             = params.NewListDNSParam()
	dnsRecordInfoParam       = params.NewRecordInfoDNSParam()
	dnsRecordBulkUpdateParam = params.NewRecordBulkUpdateDNSParam()
	dnsCreateParam           = params.NewCreateDNSParam()
	dnsRecordAddParam        = params.NewRecordAddDNSParam()
	dnsReadParam             = params.NewReadDNSParam()
	dnsRecordUpdateParam     = params.NewRecordUpdateDNSParam()
	dnsRecordDeleteParam     = params.NewRecordDeleteDNSParam()
	dnsUpdateParam           = params.NewUpdateDNSParam()
	dnsDeleteParam           = params.NewDeleteDNSParam()
)

// dnsCmd represents the command to manage SAKURA Cloud DNS
var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "A manage commands of DNS",
	Long:  `A manage commands of DNS`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

var dnsListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "find", "selector"},
	Short:   "List DNS",
	Long:    `List DNS`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := dnsListParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("list parameter: \n%s\n", debugMarshalIndent(dnsListParam))
		return err
	},
}

func dnsListCmdInit() {
	fs := dnsListCmd.Flags()
	fs.IntVarP(&dnsListParam.From, "from", "", 0, "set offset")
	fs.IntVarP(&dnsListParam.Max, "max", "", 0, "set limit")
	fs.StringSliceVarP(&dnsListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringSliceVarP(&dnsListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &dnsListParam.Id), "id", "", "set filter by id(s)")
	fs.StringSliceVarP(&dnsListParam.Tags, "tags", "", []string{}, "set filter by tags(AND)")
}

var dnsRecordInfoCmd = &cobra.Command{
	Use:     "record-info",
	Aliases: []string{"record-list"},
	Short:   "RecordInfo DNS",
	Long:    `RecordInfo DNS`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := dnsRecordInfoParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("record-info parameter: \n%s\n", debugMarshalIndent(dnsRecordInfoParam))
		return err
	},
}

func dnsRecordInfoCmdInit() {
	fs := dnsRecordInfoCmd.Flags()
	fs.StringVarP(&dnsRecordInfoParam.Name, "name", "", "", "set name")
	fs.StringVarP(&dnsRecordInfoParam.Type, "type", "", "", "set record type[A/AAAA/ALIAS/NS/CNAME/MX/TXT/SRV/CAA/PTR]")
}

var dnsRecordBulkUpdateCmd = &cobra.Command{
	Use: "record-bulk-update",

	Short: "RecordBulkUpdate DNS",
	Long:  `RecordBulkUpdate DNS`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := dnsRecordBulkUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("record-bulk-update parameter: \n%s\n", debugMarshalIndent(dnsRecordBulkUpdateParam))
		return err
	},
}

func dnsRecordBulkUpdateCmdInit() {
	fs := dnsRecordBulkUpdateCmd.Flags()
	fs.StringVarP(&dnsRecordBulkUpdateParam.File, "file", "", "", "set name")
	fs.StringVarP(&dnsRecordBulkUpdateParam.Mode, "mode", "", "upsert-only", "set name")
}

var dnsCreateCmd = &cobra.Command{
	Use: "create",

	Short: "Create DNS",
	Long:  `Create DNS`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := dnsCreateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("create parameter: \n%s\n", debugMarshalIndent(dnsCreateParam))
		return err
	},
}

func dnsCreateCmdInit() {
	fs := dnsCreateCmd.Flags()
	fs.StringVarP(&dnsCreateParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&dnsCreateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &dnsCreateParam.IconId), "icon-id", "", "set Icon ID")
	fs.StringVarP(&dnsCreateParam.Name, "name", "", "", "set DNS zone name")
}

var dnsRecordAddCmd = &cobra.Command{
	Use: "record-add",

	Short: "RecordAdd DNS",
	Long:  `RecordAdd DNS`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := dnsRecordAddParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("record-add parameter: \n%s\n", debugMarshalIndent(dnsRecordAddParam))
		return err
	},
}

func dnsRecordAddCmdInit() {
	fs := dnsRecordAddCmd.Flags()
	fs.StringVarP(&dnsRecordAddParam.Name, "name", "", "", "set name")
	fs.IntVarP(&dnsRecordAddParam.MxPriority, "mx-priority", "", 10, "set MX priority")
	fs.IntVarP(&dnsRecordAddParam.SrvWeight, "srv-weight", "", 0, "set SRV priority")
	fs.IntVarP(&dnsRecordAddParam.SrvPort, "srv-port", "", 0, "set SRV priority")
	fs.StringVarP(&dnsRecordAddParam.SrvTarget, "srv-target", "", "", "set SRV priority")
	fs.StringVarP(&dnsRecordAddParam.Type, "type", "", "", "set record type[A/AAAA/ALIAS/NS/CNAME/MX/TXT/SRV/CAA/PTR]")
	fs.StringVarP(&dnsRecordAddParam.Value, "value", "", "", "set record data")
	fs.IntVarP(&dnsRecordAddParam.Ttl, "ttl", "", 3600, "set ttl")
	fs.IntVarP(&dnsRecordAddParam.SrvPriority, "srv-priority", "", 0, "set SRV priority")
}

var dnsReadCmd = &cobra.Command{
	Use: "read",

	Short: "Read DNS",
	Long:  `Read DNS`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := dnsReadParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("read parameter: \n%s\n", debugMarshalIndent(dnsReadParam))
		return err
	},
}

func dnsReadCmdInit() {
}

var dnsRecordUpdateCmd = &cobra.Command{
	Use: "record-update",

	Short: "RecordUpdate DNS",
	Long:  `RecordUpdate DNS`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := dnsRecordUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("record-update parameter: \n%s\n", debugMarshalIndent(dnsRecordUpdateParam))
		return err
	},
}

func dnsRecordUpdateCmdInit() {
	fs := dnsRecordUpdateCmd.Flags()
	fs.IntVarP(&dnsRecordUpdateParam.Index, "index", "", 0, "index of target record")
	fs.StringVarP(&dnsRecordUpdateParam.Name, "name", "", "", "set name")
	fs.IntVarP(&dnsRecordUpdateParam.SrvPriority, "srv-priority", "", 0, "set SRV priority")
	fs.IntVarP(&dnsRecordUpdateParam.SrvWeight, "srv-weight", "", 0, "set SRV priority")
	fs.StringVarP(&dnsRecordUpdateParam.Type, "type", "", "", "set record type[A/AAAA/ALIAS/NS/CNAME/MX/TXT/SRV/CAA/PTR]")
	fs.StringVarP(&dnsRecordUpdateParam.Value, "value", "", "", "set record data")
	fs.IntVarP(&dnsRecordUpdateParam.Ttl, "ttl", "", 0, "set ttl")
	fs.IntVarP(&dnsRecordUpdateParam.MxPriority, "mx-priority", "", 0, "set MX priority")
	fs.IntVarP(&dnsRecordUpdateParam.SrvPort, "srv-port", "", 0, "set SRV priority")
	fs.StringVarP(&dnsRecordUpdateParam.SrvTarget, "srv-target", "", "", "set SRV priority")
}

var dnsRecordDeleteCmd = &cobra.Command{
	Use: "record-delete",

	Short: "RecordDelete DNS",
	Long:  `RecordDelete DNS`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := dnsRecordDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("record-delete parameter: \n%s\n", debugMarshalIndent(dnsRecordDeleteParam))
		return err
	},
}

func dnsRecordDeleteCmdInit() {
	fs := dnsRecordDeleteCmd.Flags()
	fs.IntVarP(&dnsRecordDeleteParam.Index, "index", "", 0, "index of target record")
}

var dnsUpdateCmd = &cobra.Command{
	Use: "update",

	Short: "Update DNS",
	Long:  `Update DNS`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := dnsUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("update parameter: \n%s\n", debugMarshalIndent(dnsUpdateParam))
		return err
	},
}

func dnsUpdateCmdInit() {
	fs := dnsUpdateCmd.Flags()
	fs.StringVarP(&dnsUpdateParam.Description, "description", "", "", "set resource description")
	fs.StringSliceVarP(&dnsUpdateParam.Tags, "tags", "", []string{}, "set resource tags")
	fs.VarP(newIDValue(0, &dnsUpdateParam.IconId), "icon-id", "", "set Icon ID")
}

var dnsDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"rm"},
	Short:   "Delete DNS",
	Long:    `Delete DNS`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := dnsDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("delete parameter: \n%s\n", debugMarshalIndent(dnsDeleteParam))
		return err
	},
}

func dnsDeleteCmdInit() {
}

func init() {
	parent := dnsCmd

	dnsListCmdInit()
	parent.AddCommand(dnsListCmd)

	dnsRecordInfoCmdInit()
	parent.AddCommand(dnsRecordInfoCmd)

	dnsRecordBulkUpdateCmdInit()
	parent.AddCommand(dnsRecordBulkUpdateCmd)

	dnsCreateCmdInit()
	parent.AddCommand(dnsCreateCmd)

	dnsRecordAddCmdInit()
	parent.AddCommand(dnsRecordAddCmd)

	dnsReadCmdInit()
	parent.AddCommand(dnsReadCmd)

	dnsRecordUpdateCmdInit()
	parent.AddCommand(dnsRecordUpdateCmd)

	dnsRecordDeleteCmdInit()
	parent.AddCommand(dnsRecordDeleteCmd)

	dnsUpdateCmdInit()
	parent.AddCommand(dnsUpdateCmd)

	dnsDeleteCmdInit()
	parent.AddCommand(dnsDeleteCmd)

	rootCmd.AddCommand(parent)
}
