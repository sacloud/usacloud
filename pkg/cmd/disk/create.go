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

package disk

import (
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cmd/base"
)

var createCommand = &base.Command{
	Name:     "create",
	Category: "basics",
	Order:    20,

	ParameterInitializer: func() interface{} {
		return newCreateParameter()
	},
}

type createParameter struct {
	base.ZoneParameter `cli:",squash" mapconv:",squash"`

	Name            string     `cli:",category=disk"`
	Description     string     `cli:",category=disk"`
	Tags            []string   `cli:",category=disk"`
	IconID          types.ID   `cli:",category=disk"`
	DiskPlan        string     `cli:",category=disk,options=disk_plan" mapconv:"DiskPlanID,filters=disk_plan_to_id" validate:"oneof=ssd hdd"`
	Connection      string     `cli:",category=disk,options=disk_connection" validate:"oneof=virtio ide"`
	SourceDiskID    types.ID   `cli:",category=disk"`
	SourceArchiveID types.ID   `cli:",category=disk"`
	ServerID        types.ID   `cli:",category=disk"`
	SizeGB          int        `cli:"size,category=disk"`
	DistantFrom     []types.ID `cli:",category=disk"`
	OSType          string     `cli:",category=disk,options=os_type" mapconv:",filters=os_type"`

	base.ConfirmParameter `cli:",squash" mapconv:"-"`
	base.OutputParameter  `cli:",squash" mapconv:"-"`
}

func newCreateParameter() *createParameter {
	return &createParameter{
		// TODO デフォルト値はここで設定する
	}
}

func init() {
	Resource.AddCommand(createCommand)
}
