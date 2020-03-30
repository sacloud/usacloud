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
	ipv6ListParam      = params.NewListIpv6Param()
	ipv6PtrAddParam    = params.NewPtrAddIpv6Param()
	ipv6PtrReadParam   = params.NewPtrReadIpv6Param()
	ipv6PtrUpdateParam = params.NewPtrUpdateIpv6Param()
	ipv6PtrDeleteParam = params.NewPtrDeleteIpv6Param()
)

// ipv6Cmd represents the command to manage SAKURA Cloud IPv6
var ipv6Cmd = &cobra.Command{
	Use:   "ipv6",
	Short: "A manage commands of IPv6",
	Long:  `A manage commands of IPv6`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

var ipv6ListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "find"},
	Short:   "List Ipv6",
	Long:    `List Ipv6`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := ipv6ListParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("list parameter: \n%s\n", debugMarshalIndent(ipv6ListParam))
		return err
	},
}

func ipv6ListCmdInit() {
	fs := ipv6ListCmd.Flags()
	fs.VarP(newIDValue(0, &ipv6ListParam.Ipv6netId), "ipv-6net-id", "", "set filter by ipv6net-id")
	fs.VarP(newIDValue(0, &ipv6ListParam.InternetId), "internet-id", "", "set filter by internet-id")
	fs.StringSliceVarP(&ipv6ListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &ipv6ListParam.Id), "id", "", "set filter by id(s)")
	fs.IntVarP(&ipv6ListParam.From, "from", "", 0, "set offset")
	fs.IntVarP(&ipv6ListParam.Max, "max", "", 0, "set limit")
	fs.StringSliceVarP(&ipv6ListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
}

var ipv6PtrAddCmd = &cobra.Command{
	Use: "ptr-add",

	Short: "PtrAdd Ipv6",
	Long:  `PtrAdd Ipv6`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := ipv6PtrAddParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("ptr-add parameter: \n%s\n", debugMarshalIndent(ipv6PtrAddParam))
		return err
	},
}

func ipv6PtrAddCmdInit() {
	fs := ipv6PtrAddCmd.Flags()
	fs.StringVarP(&ipv6PtrAddParam.Hostname, "hostname", "", "", "set server hostname")
}

var ipv6PtrReadCmd = &cobra.Command{
	Use: "ptr-read",

	Short: "PtrRead Ipv6",
	Long:  `PtrRead Ipv6`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := ipv6PtrReadParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("ptr-read parameter: \n%s\n", debugMarshalIndent(ipv6PtrReadParam))
		return err
	},
}

func ipv6PtrReadCmdInit() {
}

var ipv6PtrUpdateCmd = &cobra.Command{
	Use: "ptr-update",

	Short: "PtrUpdate Ipv6",
	Long:  `PtrUpdate Ipv6`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := ipv6PtrUpdateParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("ptr-update parameter: \n%s\n", debugMarshalIndent(ipv6PtrUpdateParam))
		return err
	},
}

func ipv6PtrUpdateCmdInit() {
	fs := ipv6PtrUpdateCmd.Flags()
	fs.StringVarP(&ipv6PtrUpdateParam.Hostname, "hostname", "", "", "set server hostname")
}

var ipv6PtrDeleteCmd = &cobra.Command{
	Use: "ptr-delete",

	Short: "PtrDelete Ipv6",
	Long:  `PtrDelete Ipv6`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := ipv6PtrDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("ptr-delete parameter: \n%s\n", debugMarshalIndent(ipv6PtrDeleteParam))
		return err
	},
}

func ipv6PtrDeleteCmdInit() {
}

func init() {
	parent := ipv6Cmd

	ipv6ListCmdInit()
	parent.AddCommand(ipv6ListCmd)

	ipv6PtrAddCmdInit()
	parent.AddCommand(ipv6PtrAddCmd)

	ipv6PtrReadCmdInit()
	parent.AddCommand(ipv6PtrReadCmd)

	ipv6PtrUpdateCmdInit()
	parent.AddCommand(ipv6PtrUpdateCmd)

	ipv6PtrDeleteCmdInit()
	parent.AddCommand(ipv6PtrDeleteCmd)

	rootCmd.AddCommand(parent)
}
