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

package base

import (
	"strings"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

type IDFlag types.ID

func NewIDFlag(val *types.ID, p *types.ID) *IDFlag {
	*p = *val
	return (*IDFlag)(p)
}

func (v *IDFlag) String() string {
	return types.ID(*v).String()
}

func (v *IDFlag) Set(id string) error {
	*v = IDFlag(types.StringID(id))
	return nil
}

func (v *IDFlag) Type() string {
	return "types.ID"
}

type IDSliceFlag []types.ID

func NewIDSliceFlag(val *[]types.ID, p *[]types.ID) *IDSliceFlag {
	*p = *val
	return (*IDSliceFlag)(p)
}

func (v *IDSliceFlag) String() string {
	var ids []string
	for _, id := range []types.ID(*v) {
		ids = append(ids, id.String())
	}
	return strings.Join(ids, ",")
}

func (v *IDSliceFlag) Set(ids string) error {
	values := strings.Split(ids, ",")
	*v = []types.ID{}
	for _, val := range values {
		*v = append(*v, types.StringID(val))
	}
	return nil
}

func (v *IDSliceFlag) Type() string {
	return "types.IDSlice"
}
