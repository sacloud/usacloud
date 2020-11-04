// Copyright 2016-2020 The Libsacloud Authors
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

package pointer

// NewStringSlice returns a pointer to the given tags value
func NewStringSlice(v []string) *[]string { return &v }

// NewIntSlice returns a pointer to the given tags value
func NewIntSlice(v []int) *[]int { return &v }

// NewInt64Slice returns a pointer to the given tags value
func NewInt64Slice(v []int64) *[]int64 { return &v }

// NewUintSlice returns a pointer to the given tags value
func NewUintSlice(v []uint) *[]uint { return &v }

// NewUint64Slice returns a pointer to the given tags value
func NewUint64Slice(v []uint64) *[]uint64 { return &v }

// NewByteSlice returns a pointer to the given tags value
func NewByteSlice(v []byte) *[]byte { return &v }
