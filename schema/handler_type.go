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

// HandlerType is an enum of the type that can be represented by a schema.
type HandlerType int //go:generate stringer -type=HandlerType :: manual

const (
	HandlerPathThrough HandlerType = iota
	HandlerPathThroughEach
	HandlerSort
	HandlerFilterBy
	HandlerAndParams
	HandlerOrParams
	HandlerCustomFunc
	HandlerFilterFunc
	HandlerNoop
)

// IsWhenListOnly return true when HandlerType is able to use with CommandList
func (h HandlerType) IsWhenListOnly() bool {
	switch h {
	case HandlerSort, HandlerAndParams, HandlerOrParams, HandlerFilterFunc:
		return true
	default:
		return false
	}
}

// IsNeedSliceValue return true when HandlerType is need Slice ValueType
func (h HandlerType) IsNeedSliceValue() bool {
	switch h {
	case HandlerPathThroughEach, HandlerSort, HandlerAndParams, HandlerOrParams:
		return true
	default:
		return false
	}
}

// CanSetDestinationProp return true when HandlerType is able to have DestinationProp
func (h HandlerType) CanSetDestinationProp() bool {
	switch h {
	case HandlerPathThrough, HandlerPathThroughEach, HandlerFilterBy, HandlerAndParams, HandlerOrParams:
		return true
	default:
		return false
	}
}
