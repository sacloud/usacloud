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
