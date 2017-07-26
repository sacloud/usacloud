package internal

import (
	"crypto/md5"
	"fmt"
)

type HashQueue struct {
	m     map[string]bool
	order []string
	Size  int
}

func NewHashQueue(size int) *HashQueue {
	return &HashQueue{
		m:     map[string]bool{},
		order: []string{},
		Size:  size,
	}
}

func (q *HashQueue) PutIfAbsent(value string) bool {
	key := fmt.Sprintf("%x", md5.Sum([]byte(value)))
	if _, ok := q.m[key]; ok {
		return false // exists
	}
	q.m[key] = true
	q.order = append(q.order, key)
	if q.Size > 0 && len(q.order) > q.Size {
		delete(q.m, q.order[0])
		q.order = q.order[1:]
	}
	return true
}
