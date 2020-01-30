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

package schema

// CommandType is an enum of the type that can be represented by a schema.
type CommandType int //go:generate stringer -type=CommandType :: manual

const (
	CommandInvalid CommandType = iota
	CommandList
	CommandCreate
	CommandRead
	CommandUpdate
	CommandDelete
	CommandManipulateMulti
	CommandManipulateSingle
	CommandManipulateIDOnly
	CommandCustom
)

// IsRequiredIDType 引数に対象リソースの指定を要求するコマンドタイプであるか
func (c CommandType) IsRequiredIDType() bool {
	switch c {
	case CommandRead, CommandUpdate, CommandDelete, CommandManipulateMulti, CommandManipulateSingle, CommandManipulateIDOnly:
		return true
	default:
		return false
	}
}

// IsNeedSingleIDType 引数に単一の対象リソースを要求するコマンドタイプであるか
func (c CommandType) IsNeedSingleIDType() bool {
	return c == CommandManipulateSingle || c == CommandRead
}

// IsNeedIDOnlyType 引数にIDのみ受け付けるコマンドタイプであるか
func (c CommandType) IsNeedIDOnlyType() bool {
	return c == CommandManipulateIDOnly
}

// IsNeedConfirmType コマンド実行時に確認ダイアログが必要なコマンドタイプであるか
func (c CommandType) IsNeedConfirmType() bool {
	switch c {
	case CommandCreate, CommandUpdate, CommandDelete, CommandManipulateMulti, CommandManipulateSingle, CommandManipulateIDOnly, CommandCustom:
		return true
	default:
		return false
	}
}

// CanUseSelector タグなどでのセレクターが利用可能なコマンドタイプであるか
func (c CommandType) CanUseSelector() bool {
	return c.IsRequiredIDType() && !c.IsNeedIDOnlyType()
}
