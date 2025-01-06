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

package autoscale

import (
	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/examples"
	"github.com/sacloud/usacloud/pkg/util"
)

var updateCommand = &core.Command{
	Name:         "update",
	Category:     "basic",
	Order:        40,
	SelectorType: core.SelectorTypeRequireMulti,

	ColumnDefs: defaultColumnDefs,

	ParameterInitializer: func() interface{} {
		return newUpdateParameter()
	},
}

type updateParameter struct {
	cflag.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag.NameUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.DescUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.TagsUpdateParameter   `cli:",squash" mapconv:",omitempty,squash"`
	cflag.IconIDUpdateParameter `cli:",squash" mapconv:",omitempty,squash"`

	Zones  *[]string `validate:"omitempty,required"`
	Config *string   `validate:"omitempty,required" mapconv:",omitempty,filters=dereference path_or_content"`

	Disabled               bool
	TriggerType            *string                      `cli:"trigger-type,options=cpu router schedule" validate:"omitempty,oneof=cpu router schedule" mapconv:",omitempty"`
	CPUThresholdScaling    UpdateCPUThresholdScaling    `validate:"omitempty,dive"`
	RouterThresholdScaling UpdateRouterThresholdScaling `validate:"omitempty,dive"`

	ScheduleScalingData *string                           `cli:"schedule-scaling" mapconv:"-" json:"-"`
	ScheduleScaling     *[]*iaas.AutoScaleScheduleScaling `cli:"-"`
}

type UpdateCPUThresholdScaling struct {
	ServerPrefix *string
	Up           *int
	Down         *int
}

type UpdateRouterThresholdScaling struct {
	RouterPrefix *string
	Direction    *string `validate:"omitempty,oneof=in out"`
	Mbps         *int
}

func newUpdateParameter() *updateParameter {
	return &updateParameter{}
}

func init() {
	Resource.AddCommand(updateCommand)
}

// Customize パラメータ変換処理
func (p *updateParameter) Customize(_ cli.Context) error {
	var scheduleScaling []*iaas.AutoScaleScheduleScaling
	if p.ScheduleScalingData != nil && *p.ScheduleScalingData != "" {
		if err := util.MarshalJSONFromPathOrContent(*p.ScheduleScalingData, &scheduleScaling); err != nil {
			return err
		}
		p.ScheduleScaling = &scheduleScaling
	}
	return nil
}

func (p *updateParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &updateParameter{
		NameUpdateParameter:   examples.NameUpdate,
		DescUpdateParameter:   examples.DescriptionUpdate,
		TagsUpdateParameter:   examples.TagsUpdate,
		IconIDUpdateParameter: examples.IconIDUpdate,
	}
}
