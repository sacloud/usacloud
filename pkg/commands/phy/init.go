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

package phy

import (
	"fmt"

	v1 "github.com/sacloud/phy-api-go/apis/v1"
	"github.com/sacloud/phy-service-go/server"
	"github.com/sacloud/usacloud/pkg/core"
	_ "github.com/sacloud/usacloud/pkg/services/phy"
)

func init() {
	core.LabelsExtractors = append(core.LabelsExtractors, labelsExtractorForPhy)
}

func labelsExtractorForPhy(v interface{}) *core.Labels {
	switch v := v.(type) {
	case *v1.Service:
		return &core.Labels{Id: v.ServiceId, Name: v.Nickname, Tags: v1TagsToStrings(&v.Tags)}
	case *v1.Server:
		return &core.Labels{Id: v.ServerId, Name: v.Service.Nickname, Tags: v1TagsToStrings(v.Service.Tags)}
	case *server.Server:
		return &core.Labels{Id: v.ServerId, Name: v.Service.Nickname, Tags: v1TagsToStrings(v.Service.Tags)}
	case *v1.InterfacePort:
		return &core.Labels{Id: fmt.Sprintf("%d", v.PortId), Name: v.Nickname}
	}
	return nil
}

func v1TagsToStrings(tags *[]v1.Tag) []string {
	if tags == nil {
		return nil
	}
	var res []string
	for _, t := range *tags {
		res = append(res, t.Label)
	}
	return res
}
