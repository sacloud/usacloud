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
	"encoding/json"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
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
	cflag.ZoneParameter   `cli:",squash" mapconv:",squash"`
	cflag.IDParameter     `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter `cli:",squash" mapconv:"-"`

	NoWait bool `request:"-"` // trueの場合ディスクの修正完了まで待たずに即時復帰する

	HostName string
	Password string

	DisablePWAuth       bool
	EnableDHCP          bool
	ChangePartitionUUID bool

	IPAddress      string
	NetworkMaskLen int
	DefaultRoute   string

	SSHKeys   []string   `cli:"ssh-keys"`
	SSHKeyIDs []types.ID `cli:"ssh-key-ids"`

	NoteIDs   []types.ID              `cli:"note-ids" mapconv:"-"`
	NotesData string                  `cli:"notes" mapconv:"-"`
	Notes     []*sacloud.DiskEditNote `cli:"-"` // --parametersでファイルからパラメータ指定する場合向け

	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`
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
func (p *editParameter) Customize() error {
	var notes []*sacloud.DiskEditNote
	if p.NotesData != "" {
		if err := json.Unmarshal([]byte(p.NotesData), &notes); err != nil {
			return err
		}
	}

	for _, id := range p.NoteIDs {
		notes = append(notes, &sacloud.DiskEditNote{ID: id})
	}

	p.Notes = append(p.Notes, notes...)
	return nil
}
