// Copyright 2017-2025 The sacloud/usacloud Authors
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

package eventbusapi

import (
	v1 "github.com/sacloud/sacloud-sdk-go/api/eventbus/apis/v1"
	"github.com/sacloud/usacloud/pkg/core"
)

func init() {
	core.LabelsExtractors = append(core.LabelsExtractors, extractLabels)
}

func extractLabels(v interface{}) *core.Labels {
	item, ok := v.(*v1.CommonServiceItem)
	if !ok {
		return nil
	}
	return &core.Labels{
		Id:   item.ID,
		Name: item.Name,
		Tags: item.Tags,
	}
}
