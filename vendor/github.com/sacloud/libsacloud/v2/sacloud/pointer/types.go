// Copyright 2016-2021 The Libsacloud Authors
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

import "github.com/sacloud/libsacloud/v2/sacloud/types"

// NewID returns a pointer to the given value
func NewID(id types.ID) *types.ID { return &id }

// NewIDSlice returns a pointer to the given value
func NewIDSlice(ids []types.ID) *[]types.ID { return &ids }

// NewTags returns a pointer to the given tags value
func NewTags(tags types.Tags) *types.Tags { return &tags }

// NewDNSRecordType returns a pointer to the given type value
func NewDNSRecordType(t types.EDNSRecordType) *types.EDNSRecordType { return &t }
