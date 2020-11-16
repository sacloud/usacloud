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
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/validate"
)

var createCommand = &core.Command{
	Name:     "create",
	Category: "basic",
	Order:    20,

	ColumnDefs: defaultColumnDefs,

	ParameterInitializer: func() interface{} {
		return newCreateParameter()
	},
	ValidateFunc: validateCreateParameter,
}

type createParameter struct {
	cflag.ZoneParameter `cli:",squash" mapconv:",squash"`

	Name            string     `cli:",category=disk" validate:"required"`
	Description     string     `cli:",category=disk" validate:"description"`
	Tags            []string   `cli:",category=disk" validate:"tags"`
	IconID          types.ID   `cli:",category=disk"`
	DiskPlan        string     `cli:",category=disk,options=disk_plan" mapconv:"DiskPlanID,filters=disk_plan_to_value" validate:"required,disk_plan"`
	Connection      string     `cli:",category=disk,options=disk_connection" validate:"required,disk_connection"`
	SourceDiskID    types.ID   `cli:",category=disk"`
	SourceArchiveID types.ID   `cli:",category=disk"`
	ServerID        types.ID   `cli:",category=disk"`
	SizeGB          int        `cli:"size,category=disk"`
	DistantFrom     []types.ID `cli:",category=disk"`
	OSType          string     `cli:",category=disk,options=os_type" mapconv:",filters=os_type_to_value" validate:"omitempty,os_type"`

	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`
}

func validateCreateParameter(ctx cli.Context, parameter interface{}) error {
	if err := validate.Exec(parameter); err != nil {
		return err
	}
	p := parameter.(*createParameter)

	var errs []error
	// OSTypeとSourceXXXが指定されていたらエラー
	if p.OSType != "" && (!p.SourceArchiveID.IsEmpty() || !p.SourceDiskID.IsEmpty()) {
		errs = append(errs, validate.NewFlagError("--os-type & --source-archive-id & --source-disk-id", "only one of them can be specified"))
	}

	// SourceXXXが両方指定されていたらエラーとする
	if !p.SourceArchiveID.IsEmpty() && !p.SourceDiskID.IsEmpty() {
		errs = append(errs, validate.NewFlagError("--source-archive-id & --source-disk-id", "only one of them can be specified"))
	}

	return validate.NewValidationError(errs...)
}

func newCreateParameter() *createParameter {
	return &createParameter{
		DiskPlan:   "ssd",
		Connection: "virtio",
		SizeGB:     20,
	}
}

func init() {
	Resource.AddCommand(createCommand)
}
