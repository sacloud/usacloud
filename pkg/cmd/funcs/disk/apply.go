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
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cmd/core"
)

// TODO 未実装
var applyCommand = &core.Command{ // nolint TODO あとでnolintを消す
	Name:         "apply",
	Category:     "basic",
	Order:        40,
	SelectorType: core.SelectorTypeNone, // TODO 将来的には作成(ID不要) + 更新(ID必要)をサポートしたい、この際にSelectorTypeを増やすべきかも

	ParameterInitializer: func() interface{} {
		return newApplyParameter()
	},
}

type applyParameter struct {
	core.ZoneParameter `cli:",squash" mapconv:",squash"`
	// core.IDParameter   `cli:",squash" mapconv:",squash"` //TODO SelectorTypeのコメント参照(将来ID使うかも)

	Name            string     `cli:",category=disk" validate:"required"`
	Description     string     `cli:",category=disk" validate:"description"`
	Tags            []string   `cli:",category=disk" validate:"tags"`
	IconID          types.ID   `cli:",category=disk"`
	DiskPlan        string     `cli:",category=disk,options=disk_plan" mapconv:"DiskPlanID,filters=disk_plan_to_id" validate:"required,disk_plan"`
	Connection      string     `cli:",category=disk,options=disk_connection" validate:"required,disk_connection"`
	SourceDiskID    types.ID   `cli:",category=disk"`
	SourceArchiveID types.ID   `cli:",category=disk"`
	ServerID        types.ID   `cli:",category=disk"`
	SizeGB          int        `cli:"size,category=disk"`
	DistantFrom     []types.ID `cli:",category=disk"`
	OSType          string     `cli:",category=disk,options=os_type" mapconv:",filters=os_type" validate:"omitempty,os_type"`

	//EditParameter editParameter `cli:",category=edit" mapconv:",omitempty"`

	core.ConfirmParameter `cli:",squash" mapconv:"-"`
	core.OutputParameter  `cli:",squash" mapconv:"-"`
}

func newApplyParameter() *applyParameter {
	return &applyParameter{
		// TODO デフォルト値はここで設定する
	}
}

func init() {
	// TODO serviceの戻り値が複数あるためコード生成してもビルドエラーになる。
	// https://github.com/sacloud/libsacloud/issues/630 の対応がされたら対応する。
	//Resource.AddCommand(applyCommand)
}
