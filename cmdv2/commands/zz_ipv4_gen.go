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
	ipv4ListParam      = params.NewListIPv4Param()
	ipv4PtrAddParam    = params.NewPtrAddIPv4Param()
	ipv4PtrReadParam   = params.NewPtrReadIPv4Param()
	ipv4PtrUpdateParam = params.NewPtrUpdateIPv4Param()
	ipv4PtrDeleteParam = params.NewPtrDeleteIPv4Param()
)

// ipv4Cmd represents the command to manage SAKURA Cloud IPv4
var ipv4Cmd = &cobra.Command{
	Use:   "ipv4",
	Short: "A manage commands of IPv4",
	Long:  `A manage commands of IPv4`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

var ipv4ListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "find"},
	Short:   "List IPv4",
	Long:    `List IPv4`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := ipv4ListParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("list parameter: \n%s\n", debugMarshalIndent(ipv4ListParam))
		return err
	},
}

func ipv4ListCmdInit() {
	fs := ipv4ListCmd.Flags()
	fs.IntVarP(&ipv4ListParam.Max, "max", "", 0, "set limit")
	fs.StringSliceVarP(&ipv4ListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
	fs.StringSliceVarP(&ipv4ListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &ipv4ListParam.Id), "id", "", "set filter by id(s)")
	fs.IntVarP(&ipv4ListParam.From, "from", "", 0, "set offset")
}

var ipv4PtrAddCmd = &cobra.Command{
	Use: "ptr-add",

	Short: "PtrAdd IPv4",
	Long:  `PtrAdd IPv4`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := ipv4PtrAddParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("ptr-add parameter: \n%s\n", debugMarshalIndent(ipv4PtrAddParam))
		return err
	},
}

func ipv4PtrAddCmdInit() {
	fs := ipv4PtrAddCmd.Flags()
	fs.StringVarP(&ipv4PtrAddParam.Hostname, "hostname", "", "", "set server hostname")
}

var ipv4PtrReadCmd = &cobra.Command{
	Use: "ptr-read",

	Short: "PtrRead IPv4",
	Long:  `PtrRead IPv4`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := ipv4PtrReadParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("ptr-read parameter: \n%s\n", debugMarshalIndent(ipv4PtrReadParam))
		return err
	},
}

func ipv4PtrReadCmdInit() {
}

var ipv4PtrUpdateCmd = &cobra.Command{
	Use: "ptr-update",

	Short: "PtrUpdate IPv4",
	Long:  `PtrUpdate IPv4`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := ipv4PtrUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("ptr-update parameter: \n%s\n", debugMarshalIndent(ipv4PtrUpdateParam))
		return err
	},
}

func ipv4PtrUpdateCmdInit() {
	fs := ipv4PtrUpdateCmd.Flags()
	fs.StringVarP(&ipv4PtrUpdateParam.Hostname, "hostname", "", "", "set server hostname")
}

var ipv4PtrDeleteCmd = &cobra.Command{
	Use: "ptr-delete",

	Short: "PtrDelete IPv4",
	Long:  `PtrDelete IPv4`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := ipv4PtrDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("ptr-delete parameter: \n%s\n", debugMarshalIndent(ipv4PtrDeleteParam))
		return err
	},
}

func ipv4PtrDeleteCmdInit() {
}

func init() {
	parent := ipv4Cmd

	ipv4ListCmdInit()
	parent.AddCommand(ipv4ListCmd)

	ipv4PtrAddCmdInit()
	parent.AddCommand(ipv4PtrAddCmd)

	ipv4PtrReadCmdInit()
	parent.AddCommand(ipv4PtrReadCmd)

	ipv4PtrUpdateCmdInit()
	parent.AddCommand(ipv4PtrUpdateCmd)

	ipv4PtrDeleteCmdInit()
	parent.AddCommand(ipv4PtrDeleteCmd)

	rootCmd.AddCommand(parent)
}
