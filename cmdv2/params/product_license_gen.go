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

// ListProductlicenseParam is input parameters for the sacloud API
type ListProductlicenseParam struct {
	Sort []string
	Name []string
	Id   []sacloud.ID
	From int
	Max  int

	input Input
}

// NewListProductlicenseParam return new ListProductlicenseParam
func NewListProductlicenseParam() *ListProductlicenseParam {
	return &ListProductlicenseParam{}
}

// Initialize init ListProductlicenseParam
func (p *ListProductlicenseParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ListProductlicenseParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *ListProductlicenseParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.Sort) {
		p.Sort = []string{""}
	}
	if utils.IsEmpty(p.Name) {
		p.Name = []string{""}
	}
	if utils.IsEmpty(p.Id) {
		p.Id = []sacloud.ID{}
	}
	if utils.IsEmpty(p.From) {
		p.From = 0
	}
	if utils.IsEmpty(p.Max) {
		p.Max = 0
	}

}

func (p *ListProductlicenseParam) validate() error {
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
		validator := define.Resources["ProductLicense"].Commands["list"].Params["id"].ValidateFunc
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

func (p *ListProductlicenseParam) ResourceDef() *schema.Resource {
	return define.Resources["ProductLicense"]
}

func (p *ListProductlicenseParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["list"]
}

func (p *ListProductlicenseParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ListProductlicenseParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ListProductlicenseParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ListProductlicenseParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ListProductlicenseParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListProductlicenseParam) GetSort() []string {
	return p.Sort
}
func (p *ListProductlicenseParam) SetName(v []string) {
	p.Name = v
}

func (p *ListProductlicenseParam) GetName() []string {
	return p.Name
}
func (p *ListProductlicenseParam) SetId(v []sacloud.ID) {
	p.Id = v
}

func (p *ListProductlicenseParam) GetId() []sacloud.ID {
	return p.Id
}
func (p *ListProductlicenseParam) SetFrom(v int) {
	p.From = v
}

func (p *ListProductlicenseParam) GetFrom() int {
	return p.From
}
func (p *ListProductlicenseParam) SetMax(v int) {
	p.Max = v
}

func (p *ListProductlicenseParam) GetMax() int {
	return p.Max
}

// ReadProductlicenseParam is input parameters for the sacloud API
type ReadProductlicenseParam struct {
	Id sacloud.ID

	input Input
}

// NewReadProductlicenseParam return new ReadProductlicenseParam
func NewReadProductlicenseParam() *ReadProductlicenseParam {
	return &ReadProductlicenseParam{}
}

// Initialize init ReadProductlicenseParam
func (p *ReadProductlicenseParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ReadProductlicenseParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *ReadProductlicenseParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.Id) {
		p.Id = sacloud.ID(0)
	}

}

func (p *ReadProductlicenseParam) validate() error {
	var errors []error

	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ProductLicense"].Commands["read"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return utils.FlattenErrors(errors)
}

func (p *ReadProductlicenseParam) ResourceDef() *schema.Resource {
	return define.Resources["ProductLicense"]
}

func (p *ReadProductlicenseParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["read"]
}

func (p *ReadProductlicenseParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ReadProductlicenseParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ReadProductlicenseParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ReadProductlicenseParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ReadProductlicenseParam) SetId(v sacloud.ID) {
	p.Id = v
}

func (p *ReadProductlicenseParam) GetId() sacloud.ID {
	return p.Id
}
