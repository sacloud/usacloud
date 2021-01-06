// Copyright 2016-2021 The Libsacloud Authors
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
	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type EditRequest struct {
	Zone string   `request:"-" validate:"required"`
	ID   types.ID `request:"-"`

	NoWait bool `request:"-"` // trueの場合ディスクの修正完了まで待たずに即時復帰する

	HostName string
	Password string

	DisablePWAuth       bool
	EnableDHCP          bool
	ChangePartitionUUID bool

	IPAddress      string
	NetworkMaskLen int
	DefaultRoute   string

	SSHKeys   []string
	SSHKeyIDs []types.ID

	Notes []*sacloud.DiskEditNote // スタートアップスクリプトをIDで指定(変数や埋め込むAPIキーを指定可能)
}

func (req *EditRequest) Validate() error {
	return validate.Struct(req)
}

func (req *EditRequest) ToRequestParameter() (*sacloud.DiskEditRequest, error) {
	// TODO builderからコピーしたもの。builderとの統合後に整理する
	editReq := &sacloud.DiskEditRequest{
		Background:          true,
		Password:            req.Password,
		DisablePWAuth:       req.DisablePWAuth,
		EnableDHCP:          req.EnableDHCP,
		ChangePartitionUUID: req.ChangePartitionUUID,
		HostName:            req.HostName,
	}

	if req.IPAddress != "" {
		editReq.UserIPAddress = req.IPAddress
	}
	if req.NetworkMaskLen > 0 || req.DefaultRoute != "" {
		editReq.UserSubnet = &sacloud.DiskEditUserSubnet{
			NetworkMaskLen: req.NetworkMaskLen,
			DefaultRoute:   req.DefaultRoute,
		}
	}

	// ssh key
	var sshKeys []*sacloud.DiskEditSSHKey
	for _, key := range req.SSHKeys {
		sshKeys = append(sshKeys, &sacloud.DiskEditSSHKey{
			PublicKey: key,
		})
	}
	for _, id := range req.SSHKeyIDs {
		sshKeys = append(sshKeys, &sacloud.DiskEditSSHKey{
			ID: id,
		})
	}
	editReq.SSHKeys = sshKeys

	// startup script
	var notes []*sacloud.DiskEditNote
	for _, note := range req.Notes {
		notes = append(notes, &sacloud.DiskEditNote{
			ID:        note.ID,
			APIKeyID:  note.APIKeyID,
			Variables: note.Variables,
		})
	}
	editReq.Notes = notes

	return editReq, nil
}
