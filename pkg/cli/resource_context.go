// Copyright 2017-2021 The Usacloud Authors
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

package cli

import (
	"fmt"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// ResourceContext 現在処理中のリソースの情報
type ResourceContext struct {
	ID       types.ID
	Zone     string
	Resource interface{} // 対象のリソースそのもの
}

func (r *ResourceContext) String() string {
	if r.ID.IsEmpty() {
		return ""
	}
	if r.Zone == "" {
		return r.ID.String()
	}
	return fmt.Sprintf("[%s] %s", r.Zone, r.ID.String())
}

type ResourceContexts []ResourceContext

func (r *ResourceContexts) Append(values ...ResourceContext) {
	for _, rc := range values {
		exists := false
		for _, v := range *r {
			if v.Zone == rc.Zone && v.ID == rc.ID {
				exists = true
				break
			}
		}
		if !exists {
			*r = append(*r, rc)
		}
	}
}

func (r *ResourceContexts) IDs() []types.ID {
	var ids []types.ID
	for _, v := range *r {
		if !v.ID.IsEmpty() {
			ids = append(ids, v.ID)
		}
	}
	return ids
}

func (r *ResourceContexts) Strings() []string {
	var res []string
	for _, v := range *r {
		s := v.String()
		if s != "" {
			res = append(res, s)
		}
	}
	return res
}
