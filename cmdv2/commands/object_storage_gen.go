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

	"github.com/sacloud/usacloud/cmdv2/params"
	"github.com/spf13/cobra"
)

var (
	objectStorageListParam   = params.NewListObjectStorageParam()
	objectStoragePutParam    = params.NewPutObjectStorageParam()
	objectStorageGetParam    = params.NewGetObjectStorageParam()
	objectStorageDeleteParam = params.NewDeleteObjectStorageParam()
)

// objectStorageCmd represents the command to manage SAKURA Cloud ObjectStorage
var objectStorageCmd = &cobra.Command{
	Use:   "objectStorage",
	Short: "A manage commands of ObjectStorage",
	Long:  `A manage commands of ObjectStorage`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

var objectStorageListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List ObjectStorage",
	Long:    `List ObjectStorage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := objectStorageListParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("list parameter: \n%s\n", debugMarshalIndent(objectStorageListParam))
		return err
	},
}

func objectStorageListCmdInit() {
	fs := objectStorageListCmd.Flags()
	fs.StringVarP(&objectStorageListParam.Bucket, "bucket", "", "", "set bucket")
	fs.StringVarP(&objectStorageListParam.AccessKey, "access-key", "", "", "set access-key")
	fs.StringVarP(&objectStorageListParam.SecretKey, "secret-key", "", "", "set access-key")
}

var objectStoragePutCmd = &cobra.Command{
	Use: "put",

	Short: "Put ObjectStorage",
	Long:  `Put ObjectStorage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := objectStoragePutParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("put parameter: \n%s\n", debugMarshalIndent(objectStoragePutParam))
		return err
	},
}

func objectStoragePutCmdInit() {
	fs := objectStoragePutCmd.Flags()
	fs.BoolVarP(&objectStoragePutParam.Recursive, "recursive", "r", false, "put objects recursive")
	fs.StringVarP(&objectStoragePutParam.AccessKey, "access-key", "", "", "set access-key")
	fs.StringVarP(&objectStoragePutParam.SecretKey, "secret-key", "", "", "set access-key")
	fs.StringVarP(&objectStoragePutParam.Bucket, "bucket", "", "", "set bucket")
	fs.StringVarP(&objectStoragePutParam.ContentType, "content-type", "", "application/octet-stream", "set content-type")
}

var objectStorageGetCmd = &cobra.Command{
	Use: "get",

	Short: "Get ObjectStorage",
	Long:  `Get ObjectStorage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := objectStorageGetParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("get parameter: \n%s\n", debugMarshalIndent(objectStorageGetParam))
		return err
	},
}

func objectStorageGetCmdInit() {
	fs := objectStorageGetCmd.Flags()
	fs.StringVarP(&objectStorageGetParam.AccessKey, "access-key", "", "", "set access-key")
	fs.StringVarP(&objectStorageGetParam.SecretKey, "secret-key", "", "", "set access-key")
	fs.StringVarP(&objectStorageGetParam.Bucket, "bucket", "", "", "set bucket")
	fs.BoolVarP(&objectStorageGetParam.Recursive, "recursive", "r", false, "get objects recursive")
}

var objectStorageDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"rm", "del"},
	Short:   "Delete ObjectStorage",
	Long:    `Delete ObjectStorage`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := objectStorageDeleteParam.Initialize(newParamsAdapter(cmd.Flags()))
		// TODO DEBUG
		fmt.Printf("delete parameter: \n%s\n", debugMarshalIndent(objectStorageDeleteParam))
		return err
	},
}

func objectStorageDeleteCmdInit() {
	fs := objectStorageDeleteCmd.Flags()
	fs.StringVarP(&objectStorageDeleteParam.Bucket, "bucket", "", "", "set bucket")
	fs.BoolVarP(&objectStorageDeleteParam.Recursive, "recursive", "r", false, "delete objects recursive")
	fs.StringVarP(&objectStorageDeleteParam.AccessKey, "access-key", "", "", "set access-key")
	fs.StringVarP(&objectStorageDeleteParam.SecretKey, "secret-key", "", "", "set access-key")
}

func init() {
	parent := objectStorageCmd

	objectStorageListCmdInit()
	parent.AddCommand(objectStorageListCmd)

	objectStoragePutCmdInit()
	parent.AddCommand(objectStoragePutCmd)

	objectStorageGetCmdInit()
	parent.AddCommand(objectStorageGetCmd)

	objectStorageDeleteCmdInit()
	parent.AddCommand(objectStorageDeleteCmd)

	rootCmd.AddCommand(parent)
}
