// Copyright 2017-2023 The sacloud/usacloud Authors
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

package archive

import (
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/output"
)

var ftpOpenCommand = &core.Command{
	Name:               "ftp-open",
	Aliases:            []string{"open-ftp"},
	ServiceFuncAltName: "OpenFTP", // v0との互換用、コマンド名をopen-ftpではなくftp-openにしているためにこれが必要
	Category:           "operation",
	Order:              30,
	SelectorType:       core.SelectorTypeRequireMulti,

	ColumnDefs: []output.ColumnDef{
		{Name: "Zone"},
		{Name: "ID"},
		{Name: "HostName"},
		{Name: "IPAddress"},
		{Name: "User"},
		{Name: "Password"},
	},

	ParameterInitializer: func() interface{} {
		return newFTPOpenParameter()
	},
}

type ftpOpenParameter struct {
	cflag.ZoneParameter    `cli:",squash" mapconv:",squash"`
	cflag.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	ChangePassword bool `cli:",category=FTP"`
}

func newFTPOpenParameter() *ftpOpenParameter {
	return &ftpOpenParameter{}
}

func init() {
	Resource.AddCommand(ftpOpenCommand)
}
