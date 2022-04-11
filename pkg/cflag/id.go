// Copyright 2017-2022 The Usacloud Authors
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

type IDParameterValueHandler interface {
	IDFlagValue() string
	SetIDFlagValue(id string)
}

func IDFlagValue(p interface{}) string {
	if p == nil {
		return ""
	}
	v, ok := p.(IDParameterValueHandler)
	if !ok {
		return ""
	}
	return v.IDFlagValue()
}

func SetIDFlagValue(p interface{}, id string) {
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
	ID string `cli:"-" json:"-"` // IDは実行時にName or Tagsから検索〜設定されるケースがあるためvalidate:"required"にしない
}

func (p *IDParameter) IDFlagValue() string {
	return p.ID
}

func (p *IDParameter) SetIDFlagValue(id string) {
	p.ID = id
}
