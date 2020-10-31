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

package base

import "github.com/sacloud/libsacloud/v2/sacloud/types"

// Input 入力値を保持/参照するためのインターフェース
type Input interface {
	Changed(name string) bool
	Bool(name string) (bool, error)
	String(name string) (string, error)
	StringSlice(name string) ([]string, error)
	Int(name string) (int, error)
	IntSlice(name string) ([]int, error)
	Int64(name string) (int64, error)
	Int64Slice(name string) ([]int64, error)
	ID(name string) (types.ID, error)
	IDSlice(name string) ([]types.ID, error)
}
