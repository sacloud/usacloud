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

package search

import (
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/search"
	"github.com/sacloud/libsacloud/v2/sacloud/search/keys"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/util"
)

func FindCondition(in interface{}) *sacloud.FindCondition {
	fc := &sacloud.FindCondition{
		Filter: search.Filter{},
	}
	if v, ok := in.(paramName); ok && !util.IsEmpty(v.GetName()) {
		fc.Filter[search.Key("Name")] = search.PartialMatch(v.GetName()...)
	}
	if v, ok := in.(paramID); ok && !util.IsEmpty(v.GetId()) {
		fc.Filter[search.Key("ID")] = search.ExactMatch(util.StringIDs(v.GetId())...)
	}
	if v, ok := in.(paramTags); ok && !util.IsEmpty(v.GetTags()) {
		fc.Filter[search.Key(keys.Tags)] = search.TagsAndEqual(v.GetTags()...)
	}
	if v, ok := in.(paramScope); ok && !util.IsEmpty(v.GetScope()) {
		fc.Filter[search.Key("Scope")] = search.ExactMatch(v.GetScope())
	}
	if v, ok := in.(paramFrom); ok && !util.IsEmpty(v.GetFrom()) {
		fc.From = v.GetFrom()
	}
	if v, ok := in.(paramMax); ok && !util.IsEmpty(v.GetMax()) {
		fc.Count = v.GetMax()
	}
	if v, ok := in.(paramSort); ok && !util.IsEmpty(v.GetSort()) {
		fc.Sort = SliceToSortKeys(v.GetSort())
	}
	return fc
}

type paramName interface {
	GetName() []string
}

type paramID interface {
	GetId() []types.ID
}

type paramTags interface {
	GetTags() []string
}

type paramFrom interface {
	GetFrom() int
}

type paramMax interface {
	GetMax() int
}

type paramSort interface {
	GetSort() []string
}

type paramScope interface {
	GetScope() string
}
