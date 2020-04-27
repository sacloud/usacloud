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

package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogHashBuffer_PutIfAbsent(t *testing.T) {

	b := NewHashQueue(2)

	// put 1 value
	assert.True(t, b.PutIfAbsent("test1"))
	assert.False(t, b.PutIfAbsent("test1"))

	// put 2 value
	assert.True(t, b.PutIfAbsent("test2"))

	assert.False(t, b.PutIfAbsent("test1"))
	assert.False(t, b.PutIfAbsent("test2"))

	// put 3 value (dequeue first value)
	assert.True(t, b.PutIfAbsent("test3"))
	assert.False(t, b.PutIfAbsent("test2"))
	assert.False(t, b.PutIfAbsent("test3"))
	assert.True(t, b.PutIfAbsent("test1"))

}
