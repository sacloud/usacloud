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
	priceListParam = params.NewListPriceParam()
)

// priceCmd represents the command to manage SAKURA Cloud Price
var priceCmd = &cobra.Command{
	Use:   "price",
	Short: "A manage commands of Price",
	Long:  `A manage commands of Price`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO not implements: call list func as default
	},
}

var priceListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "find"},
	Short:   "List Price (default)",
	Long:    `List Price (default)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := priceListParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("list parameter: \n%s\n", debugMarshalIndent(priceListParam))
		return err
	},
}

func priceListCmdInit() {
	fs := priceListCmd.Flags()
	fs.StringSliceVarP(&priceListParam.Name, "name", "", []string{}, "set filter by name(s)")
	fs.VarP(newIDSliceValue([]sacloud.ID{}, &priceListParam.Id), "id", "", "set filter by id(s)")
	fs.IntVarP(&priceListParam.From, "from", "", 0, "set offset")
	fs.IntVarP(&priceListParam.Max, "max", "", 0, "set limit")
	fs.StringSliceVarP(&priceListParam.Sort, "sort", "", []string{}, "set field(s) for sort")
}

func init() {
	parent := priceCmd

	priceListCmdInit()
	parent.AddCommand(priceListCmd)

	rootCmd.AddCommand(parent)
}
