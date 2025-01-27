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
	"github.com/sacloud/iaas-api-go/types"
	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/examples"
	"github.com/sacloud/usacloud/pkg/util"
)

var createCommand = &core.Command{
	Name:     "create",
	Category: "basic",
	Order:    20,

	ColumnDefs: defaultColumnDefs,

	ParameterInitializer: func() interface{} {
		return newCreateParameter()
	},
}

type createParameter struct {
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag.NameParameter   `cli:",squash" mapconv:",squash"`
	cflag.DescParameter   `cli:",squash" mapconv:",squash"`
	cflag.TagsParameter   `cli:",squash" mapconv:",squash"`
	cflag.IconIDParameter `cli:",squash" mapconv:",squash"`

	Zones    []string `validate:"required"`
	Config   string   `validate:"required" mapconv:",filters=path_or_content"`
	APIKeyID string   `cli:"api-key-id" validate:"required"`

	Disabled    bool
	TriggerType string `cli:"trigger-type,options=cpu router schedule" validate:"omitempty,oneof=cpu router schedule"`

	CPUThresholdScaling    CreateCPUThresholdScaling    `mapconv:",omitempty" validate:"omitempty,dive"`
	RouterThresholdScaling CreateRouterThresholdScaling `mapconv:",omitempty" validate:"omitempty,dive"`

	ScheduleScalingData string                           `cli:"schedule-scaling" mapconv:"-" json:"-"`
	ScheduleScaling     []*iaas.AutoScaleScheduleScaling `cli:"-"`
}

type CreateCPUThresholdScaling struct {
	ServerPrefix string
	Up           int
	Down         int
}

type CreateRouterThresholdScaling struct {
	RouterPrefix string
	Direction    string `cli:",options=in out" validate:"omitempty,oneof=in out"`
	Mbps         int
}

func newCreateParameter() *createParameter {
	return &createParameter{}
}

func init() {
	Resource.AddCommand(createCommand)
}

// Customize パラメータ変換処理
func (p *createParameter) Customize(_ cli.Context) error {
	var scheduleScaling []*iaas.AutoScaleScheduleScaling
	if p.ScheduleScalingData != "" {
		if err := util.MarshalJSONFromPathOrContent(p.ScheduleScalingData, &scheduleScaling); err != nil {
			return err
		}
	}
	p.ScheduleScaling = append(p.ScheduleScaling, scheduleScaling...)
	return nil
}

func (p *createParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &createParameter{
		NameParameter:   examples.Name,
		DescParameter:   examples.Description,
		TagsParameter:   examples.Tags,
		IconIDParameter: examples.IconID,
		Zones:           []string{"is1a"},
		Config:          "...",
		APIKeyID:        "...",
		TriggerType:     "cpu | router | schedule",
		Disabled:        true,
		CPUThresholdScaling: CreateCPUThresholdScaling{
			ServerPrefix: "server-prefix",
			Up:           80,
			Down:         20,
		},
		RouterThresholdScaling: CreateRouterThresholdScaling{
			RouterPrefix: "router-prefix",
			Direction:    "in | out",
			Mbps:         100,
		},
		ScheduleScaling: []*iaas.AutoScaleScheduleScaling{
			{
				Action: "up",
				Hour:   10,
				Minute: 15,
				DayOfWeek: []types.EDayOfTheWeek{
					types.DaysOfTheWeek.Monday,
					types.DaysOfTheWeek.Tuesday,
					types.DaysOfTheWeek.Wednesday,
					types.DaysOfTheWeek.Thursday,
					types.DaysOfTheWeek.Friday,
					types.DaysOfTheWeek.Saturday,
					types.DaysOfTheWeek.Sunday,
				},
			},
		},
	}
}
