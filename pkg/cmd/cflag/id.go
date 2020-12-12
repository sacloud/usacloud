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

package cflag

import "github.com/sacloud/libsacloud/v2/sacloud/types"

type IDParameterValueHandler interface {
	IDFlagValue() types.ID
	SetIDFlagValue(id types.ID)
}

func IDFlagValue(p interface{}) types.ID {
	if p == nil {
		return types.ID(0)
	}
	v, ok := p.(IDParameterValueHandler)
	if !ok {
		return types.ID(0)
	}
	return v.IDFlagValue()
}

func SetIDFlagValue(p interface{}, id types.ID) {
	if p == nil {
		return
	}
	v, ok := p.(IDParameterValueHandler)
	if !ok {
		return
	}
	v.SetIDFlagValue(id)
}

// IDParameter IDを指定して操作する必要があるリソースが実装すべきIDパラメータの定義
type IDParameter struct {
	ID types.ID `cli:"-" json:"-"` // IDは実行時にName or Tagsから検索〜設定されるケースがあるためvalidate:"required"にしない
}

func (p *IDParameter) IDFlagValue() types.ID {
	return p.ID
}

func (p *IDParameter) SetIDFlagValue(id types.ID) {
	p.ID = id
}
