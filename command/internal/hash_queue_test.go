package internal

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
