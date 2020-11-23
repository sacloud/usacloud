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

package cdrom

import (
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
)

var ftpCloseCommand = &core.Command{
	Name:               "ftp-close",
	Aliases:            []string{"close-ftp"},
	ServiceFuncAltName: "CloseFTP", // v0との互換用、コマンド名をclose-ftpではなくftp-closeにしているためにこれが必要
	Category:           "operation",
	Order:              40,
	SelectorType:       core.SelectorTypeRequireMulti,

	ParameterInitializer: func() interface{} {
		return newFTPCloseParameter()
	},
}

type ftpCloseParameter struct {
	cflag.ZoneParameter   `cli:",squash" mapconv:",squash"`
	cflag.IDParameter     `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter `cli:",squash" mapconv:"-"`

	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
}

func newFTPCloseParameter() *ftpCloseParameter {
	return &ftpCloseParameter{}
}

func init() {
	Resource.AddCommand(ftpCloseCommand)
}
