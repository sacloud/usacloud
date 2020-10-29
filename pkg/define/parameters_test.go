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

package define

import (
	"testing"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/stretchr/testify/assert"
)

func TestFilterByTagFunc(t *testing.T) {
	t.Run("With single word", func(t *testing.T) {
		var sw1 = &sacloud.Switch{ID: 1, Tags: types.Tags{"tag1", "tag2"}}
		var sw2 = &sacloud.Switch{ID: 2, Tags: types.Tags{"tag11", "tag2"}}

		items := []interface{}{sw1, sw2}

		expects := map[string][]bool{
			"tag1":  {true, false},
			"tag2":  {true, true},
			"tag11": {false, true},
		}

		for k, v := range expects {
			for i := range items {
				res := filterListByTags(items, items[i], []string{k})
				assert.Equal(t, v[i], res,
					"filterListByTags(%q), expect:%t, but got %t", k, v[i], res)
			}
		}
	})

	t.Run("With multiple word", func(t *testing.T) {
		var sw1 = &sacloud.Switch{ID: 1, Tags: types.Tags{"tag1", "tag2"}}
		var sw2 = &sacloud.Switch{ID: 2, Tags: types.Tags{"tag1", "tag3"}}
		var sw3 = &sacloud.Switch{ID: 2, Tags: types.Tags{"tag1", "tag2", "tag3"}}

		items := []interface{}{sw1, sw2, sw3}

		type expect struct {
			words   []string
			results []bool
		}

		expects := []expect{
			{
				words:   []string{"tag1"},
				results: []bool{true, true, true},
			},
			{
				words:   []string{"tag1", "tag2"},
				results: []bool{true, false, true},
			},
			{
				words:   []string{"tag1", "tag3"},
				results: []bool{false, true, true},
			},
			{
				words:   []string{"tag2", "tag3"},
				results: []bool{false, false, true},
			},
			{
				words:   []string{"tag1", "tag2", "tag3"},
				results: []bool{false, false, true},
			},
			{
				words:   []string{"tag1", "dummy"},
				results: []bool{false, false, false},
			},
		}

		for _, v := range expects {
			for i := range items {
				res := filterListByTags(items, items[i], v.words)
				assert.Equal(t, v.results[i], res,
					"filterListByTags(%q), expect:%t, but got %t", v.words, v.results[i], res)
			}
		}
	})
}
