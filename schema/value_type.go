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

// ValueType is an enum of the type that can be represented by a schema.
type ValueType int //go:generate stringer -type=ValueType :: manual

const (
	TypeInvalid ValueType = iota
	TypeBool
	TypeInt
	TypeInt64
	TypeFloat
	TypeString
	TypeIntList
	TypeStringList
	TypeId
	TypeIdList
)

// IsSliceType return true if type is TypeIntList or TypeStringList
func (v ValueType) IsSliceType() bool {
	switch v {
	case TypeIntList, TypeStringList, TypeIdList:
		return true
	default:
		return false
	}
}
