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

package commands

import (
	"strings"

	"github.com/sacloud/libsacloud/sacloud"
)

type idValue sacloud.ID

func newIDValue(val sacloud.ID, p *sacloud.ID) *idValue {
	*p = val
	return (*idValue)(p)
}

func (v *idValue) String() string {
	return sacloud.ID(*v).String()
}

func (v *idValue) Set(id string) error {
	*v = idValue(sacloud.StringID(id))
	return nil
}

func (v *idValue) Type() string {
	return "sacloud.ID"
}

type idSliceValue []sacloud.ID

func newIDSliceValue(val []sacloud.ID, p *[]sacloud.ID) *idSliceValue {
	*p = val
	return (*idSliceValue)(p)
}

func (v *idSliceValue) String() string {
	var ids []string
	for _, id := range []sacloud.ID(*v) {
		ids = append(ids, id.String())
	}
	return strings.Join(ids, ",")
}

func (v *idSliceValue) Set(ids string) error {
	values := strings.Split(ids, ",")
	*v = []sacloud.ID{}
	for _, val := range values {
		*v = append(*v, sacloud.StringID(val))
	}
	return nil
}

func (v *idSliceValue) Type() string {
	return "sacloud.IDSlice"
}

//type int64SliceValue []int64
//
//func newInt64SliceValue(val []int64, p *[]int64) *int64SliceValue {
//	*p = val
//	return (*int64SliceValue)(p)
//}
//
//func (v *int64SliceValue) String() string {
//	var ids []string
//	for _, id := range []int64(*v) {
//		ids = append(ids, fmt.Sprintf("%d", id))
//	}
//	return strings.Join(ids, ",")
//}
//
//func (v *int64SliceValue) Set(ids string) error {
//	values := strings.Split(ids, ",")
//	*v = []int64{}
//	for _, val := range values {
//		id, err := strconv.ParseInt(val, 10, 64)
//		if err != nil {
//			return err
//		}
//		*v = append(*v, id)
//	}
//	return nil
//}
//
//func (v *int64SliceValue) Type() string {
//	return "sacloud.IDSlice"
//}
