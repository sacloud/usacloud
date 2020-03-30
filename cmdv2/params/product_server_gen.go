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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package params

import (
	"io"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/pkg/utils"
	"github.com/sacloud/usacloud/pkg/validation"
	"github.com/sacloud/usacloud/schema"
)

// ListProductServerParam is input parameters for the sacloud API
type ListProductServerParam struct {
	From int
	Max  int
	Sort []string
	Name []string
	Id   []sacloud.ID

	input Input
}

// NewListProductServerParam return new ListProductServerParam
func NewListProductServerParam() *ListProductServerParam {
	return &ListProductServerParam{}
}

// Initialize init ListProductServerParam
func (p *ListProductServerParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ListProductServerParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *ListProductServerParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.From) {
		p.From = 0
	}
	if utils.IsEmpty(p.Max) {
		p.Max = 0
	}
	if utils.IsEmpty(p.Sort) {
		p.Sort = []string{""}
	}
	if utils.IsEmpty(p.Name) {
		p.Name = []string{""}
	}
	if utils.IsEmpty(p.Id) {
		p.Id = []sacloud.ID{}
	}

}

func (p *ListProductServerParam) validate() error {
	var errors []error

	{
		errs := validation.ConflictsWith("--name", p.Name, map[string]interface{}{

			"--id": p.Id,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["ProductServer"].Commands["list"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validation.ConflictsWith("--id", p.Id, map[string]interface{}{

			"--name": p.Name,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return utils.FlattenErrors(errors)
}

func (p *ListProductServerParam) ResourceDef() *schema.Resource {
	return define.Resources["ProductServer"]
}

func (p *ListProductServerParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["list"]
}

func (p *ListProductServerParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ListProductServerParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ListProductServerParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ListProductServerParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ListProductServerParam) SetFrom(v int) {
	p.From = v
}

func (p *ListProductServerParam) GetFrom() int {
	return p.From
}
func (p *ListProductServerParam) SetMax(v int) {
	p.Max = v
}

func (p *ListProductServerParam) GetMax() int {
	return p.Max
}
func (p *ListProductServerParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListProductServerParam) GetSort() []string {
	return p.Sort
}
func (p *ListProductServerParam) SetName(v []string) {
	p.Name = v
}

func (p *ListProductServerParam) GetName() []string {
	return p.Name
}
func (p *ListProductServerParam) SetId(v []sacloud.ID) {
	p.Id = v
}

func (p *ListProductServerParam) GetId() []sacloud.ID {
	return p.Id
}

// ReadProductServerParam is input parameters for the sacloud API
type ReadProductServerParam struct {
	Id sacloud.ID

	input Input
}

// NewReadProductServerParam return new ReadProductServerParam
func NewReadProductServerParam() *ReadProductServerParam {
	return &ReadProductServerParam{}
}

// Initialize init ReadProductServerParam
func (p *ReadProductServerParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ReadProductServerParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *ReadProductServerParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.Id) {
		p.Id = sacloud.ID(0)
	}

}

func (p *ReadProductServerParam) validate() error {
	var errors []error

	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ProductServer"].Commands["read"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return utils.FlattenErrors(errors)
}

func (p *ReadProductServerParam) ResourceDef() *schema.Resource {
	return define.Resources["ProductServer"]
}

func (p *ReadProductServerParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["read"]
}

func (p *ReadProductServerParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ReadProductServerParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ReadProductServerParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ReadProductServerParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ReadProductServerParam) SetId(v sacloud.ID) {
	p.Id = v
}

func (p *ReadProductServerParam) GetId() sacloud.ID {
	return p.Id
}
