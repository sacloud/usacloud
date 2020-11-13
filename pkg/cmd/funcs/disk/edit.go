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
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cmd/core"
)

// TODO 未実装
var editCommand = &core.Command{ // nolint TODO あとでnolintを消す
	Name:         "edit",
	Category:     "operation",
	Order:        10,
	SelectorType: core.SelectorTypeRequireSingle,

	ParameterInitializer: func() interface{} {
		return newEditParameter()
	},
}

type editParameter struct {
	core.ZoneParameter `cli:",squash" mapconv:",squash"`
	core.IDParameter   `cli:",squash" mapconv:",squash"`

	EditParameter editServiceParameter `cli:",squash,category=edit" mapconv:",omitempty"`

	core.ConfirmParameter `cli:",squash" mapconv:"-"`
	core.OutputParameter  `cli:",squash" mapconv:"-"`
}

type editServiceParameter struct {
	HostName string `cli:",category=edit"`
	Password string `cli:",category=edit"`

	DisablePWAuth       bool `cli:",category=edit"`
	EnableDHCP          bool `cli:",category=edit"`
	ChangePartitionUUID bool `cli:",category=edit"`

	IPAddress      string `cli:",category=edit"`
	NetworkMaskLen int    `cli:",category=edit"`
	DefaultRoute   string `cli:",category=edit"`

	SSHKeys   []string   `cli:",category=edit"`
	SSHKeyIDs []types.ID `cli:",category=edit"`

	Notes []*sacloud.DiskEditNote // TODO 2段階以上にネストしたパラメータをどう扱うか? => https://github.com/sacloud/usacloud/issues/568
}

func newEditParameter() *editParameter {
	return &editParameter{
		// TODO デフォルト値はここで設定する
	}
}

func init() {
	// TODO libsacloud service側でeditが未実装
	//Resource.AddCommand(editCommand)
}
