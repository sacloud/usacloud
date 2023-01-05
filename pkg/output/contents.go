// Copyright 2017-2023 The sacloud/usacloud Authors
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

package output

import (
	"sort"

	"github.com/sacloud/iaas-api-go/accessor"
)

type Content struct {
	Zone  string
	ID    string
	Value interface{}
}

type Contents []*Content

func (c *Contents) Values() []interface{} {
	var results []interface{}
	for _, v := range *c {
		results = append(results, v.Value)
	}
	return results
}

func (c *Contents) Sort(zones []string) {
	// ゾーン配列からmapに変換
	zoneOrder := map[string]int{}
	for i, zone := range zones {
		zoneOrder[zone] = i + 1 // ゾーンなしの場合が0のため+1しておく
	}

	values := *c
	// 優先度: ゾーン => ID(Content内の) => ID(Content.Valueの)
	sort.Slice(values, func(i, j int) bool {
		o1 := zoneOrder[values[i].Zone]
		o2 := zoneOrder[values[j].Zone]
		if o1 != o2 {
			return o1 < o2
		}

		cid1 := values[i].ID
		cid2 := values[j].ID
		if cid1 != cid2 {
			return cid1 < cid2
		}

		id1 := c.forceID(values[i].Value)
		id2 := c.forceID(values[j].Value)
		return id1 < id2
	})
}

func (c *Contents) forceID(v interface{}) int64 {
	id, ok := v.(accessor.ID) // TODO IaaS以外にはどう対応すべきか?
	if !ok {
		return -1
	}
	return id.GetID().Int64()
}
