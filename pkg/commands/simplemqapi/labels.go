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

package simplemqapi

import (
	"github.com/sacloud/sacloud-sdk-go/api/simplemq"
	sdkqueue "github.com/sacloud/sacloud-sdk-go/api/simplemq/apis/v1/queue"
	"github.com/sacloud/usacloud/pkg/core"
)

func init() {
	core.LabelsExtractors = append(core.LabelsExtractors, extractLabels)
}

func extractLabels(v interface{}) *core.Labels {
	item, ok := v.(*sdkqueue.CommonServiceItem)
	if !ok {
		return nil
	}
	return &core.Labels{
		Id:   simplemq.GetQueueID(item),
		Name: simplemq.GetQueueName(item),
		Tags: item.Tags,
	}
}
