// Copyright 2017-2022 The Usacloud Authors
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

package iaas

import (
	"github.com/sacloud/iaas-api-go/accessor"
	"github.com/sacloud/usacloud/pkg/core"
)

func init() {
	core.LabelsExtractors = append(core.LabelsExtractors, extractLabels)
}

func extractLabels(v interface{}) *core.Labels {
	if v, ok := v.(accessor.ID); ok {
		labels := &core.Labels{Id: v.GetID().String()}

		// Name(部分一致)
		if name, ok := v.(accessor.Name); ok {
			labels.Name = name.GetName()
		}

		// Tags
		if tags, ok := v.(accessor.Tags); ok {
			labels.Tags = tags.GetTags()
		}
		return labels
	}
	return nil
}
