// Copyright 2017-2021 The Usacloud Authors
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

package common

import (
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/util"
)

type EditRequest struct {
	HostName string `cli:",category=diskedit,order=10" json:",omitempty"`
	Password string `cli:",category=diskedit,order=20" json:",omitempty"`

	IPAddress      string `cli:"ip-address,category=diskedit,order=30" json:",omitempty"`
	NetworkMaskLen int    `cli:"netmask,aliases=network-mask-len,category=diskedit,order=31" json:",omitempty"`
	DefaultRoute   string `cli:"gateway,aliases=default-route,category=diskedit,order=32" json:",omitempty"`

	DisablePWAuth       bool `cli:"disable-pw-auth,category=diskedit,order=40" json:",omitempty"`
	EnableDHCP          bool `cli:"enable-dhcp,category=diskedit,order=50" json:",omitempty"`
	ChangePartitionUUID bool `cli:"change-partition-uuid,category=diskedit,order=60" json:",omitempty"`

	SSHKeys            []string   `cli:"ssh-keys,category=diskedit,order=70" json:",omitempty"`
	SSHKeyIDs          []types.ID `cli:"ssh-key-ids,category=diskedit,order=71" json:",omitempty"`
	IsSSHKeysEphemeral bool       `cli:"make-ssh-keys-ephemeral,category=diskedit,order=72" json:",omitempty"`

	NoteIDs          []types.ID              `cli:"note-ids,category=diskedit,order=80" mapconv:"-" json:",omitempty"`
	NotesData        string                  `cli:"notes,category=diskedit,order=81" mapconv:"-" json:"-"`
	IsNotesEphemeral bool                    `cli:"make-notes-ephemeral,category=diskedit,order=82" json:",omitempty"`
	Notes            []*sacloud.DiskEditNote `cli:"-" json:",omitempty"` // --parametersでファイルからパラメータ指定する場合向け
}

// Customize パラメータ変換処理
func (p *EditRequest) Customize(_ cli.Context) error {
	var notes []*sacloud.DiskEditNote
	if p.NotesData != "" {
		if err := util.MarshalJSONFromPathOrContent(p.NotesData, &notes); err != nil {
			return err
		}
	}

	for _, id := range p.NoteIDs {
		notes = append(notes, &sacloud.DiskEditNote{ID: id})
	}

	p.Notes = append(p.Notes, notes...)

	for i := range p.SSHKeys {
		key, err := util.StringFromPathOrContent(p.SSHKeys[i])
		if err != nil {
			return err
		}
		p.SSHKeys[i] = key
	}
	return nil
}
