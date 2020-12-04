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

package disk

import (
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/commands/common"
	"github.com/sacloud/usacloud/pkg/cmd/core"
)

var editCommand = &core.Command{
	Name:         "edit",
	Category:     "operation",
	Order:        10,
	SelectorType: core.SelectorTypeRequireMulti,

	ParameterInitializer: func() interface{} {
		return newEditParameter()
	},
}

type editParameter struct {
	cflag.ZoneParameter    `cli:",squash" mapconv:",squash"`
	cflag.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	EditDisk common.EditRequest `cli:",squash" mapconv:",squash"`
	NoWait   bool               `request:"-"` // trueの場合ディスクの修正完了まで待たずに即時復帰する
}

func newEditParameter() *editParameter {
	return &editParameter{
		// TODO デフォルト値はここで設定する
	}
}

func init() {
	Resource.AddCommand(editCommand)
}

// Customize パラメータ変換処理
func (p *editParameter) Customize(ctx cli.Context) error {
	return p.EditDisk.Customize(ctx)
}
