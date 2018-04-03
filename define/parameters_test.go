package define

import (
	"testing"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/stretchr/testify/assert"
)

func TestFilterByTagFunc(t *testing.T) {

	t.Run("With single word", func(t *testing.T) {

		var sw1 = &sacloud.Switch{Resource: sacloud.NewResource(1)}
		var sw2 = &sacloud.Switch{Resource: sacloud.NewResource(2)}

		items := []interface{}{sw1, sw2}

		sw1.SetTags([]string{"tag1", "tag2"})
		sw2.SetTags([]string{"tag11", "tag2"})

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

		var sw1 = &sacloud.Switch{Resource: sacloud.NewResource(1)}
		var sw2 = &sacloud.Switch{Resource: sacloud.NewResource(2)}
		var sw3 = &sacloud.Switch{Resource: sacloud.NewResource(3)}

		items := []interface{}{sw1, sw2, sw3}

		sw1.SetTags([]string{"tag1", "tag2"})
		sw2.SetTags([]string{"tag1", "tag3"})
		sw3.SetTags([]string{"tag1", "tag2", "tag3"})

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
