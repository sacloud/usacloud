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
	"github.com/sacloud/usacloud/schema"
)

// CsvBillParam is input parameters for the sacloud API
type CsvBillParam struct {
	BillId     sacloud.ID
	NoHeader   bool
	BillOutput string

	input Input
}

// NewCsvBillParam return new CsvBillParam
func NewCsvBillParam() *CsvBillParam {
	return &CsvBillParam{}
}

// Initialize init CsvBillParam
func (p *CsvBillParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *CsvBillParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *CsvBillParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.BillId) {
		p.BillId = sacloud.ID(0)
	}
	if utils.IsEmpty(p.NoHeader) {
		p.NoHeader = false
	}
	if utils.IsEmpty(p.BillOutput) {
		p.BillOutput = ""
	}

}

func (p *CsvBillParam) validate() error {
	var errors []error

	{
		validator := define.Resources["Bill"].Commands["csv"].Params["bill-id"].ValidateFunc
		errs := validator("--bill-id", p.BillId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return utils.FlattenErrors(errors)
}

func (p *CsvBillParam) ResourceDef() *schema.Resource {
	return define.Resources["Bill"]
}

func (p *CsvBillParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["csv"]
}

func (p *CsvBillParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *CsvBillParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *CsvBillParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *CsvBillParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *CsvBillParam) SetBillId(v sacloud.ID) {
	p.BillId = v
}

func (p *CsvBillParam) GetBillId() sacloud.ID {
	return p.BillId
}
func (p *CsvBillParam) SetNoHeader(v bool) {
	p.NoHeader = v
}

func (p *CsvBillParam) GetNoHeader() bool {
	return p.NoHeader
}
func (p *CsvBillParam) SetBillOutput(v string) {
	p.BillOutput = v
}

func (p *CsvBillParam) GetBillOutput() string {
	return p.BillOutput
}

// ListBillParam is input parameters for the sacloud API
type ListBillParam struct {
	Year  int
	Month int

	input Input
}

// NewListBillParam return new ListBillParam
func NewListBillParam() *ListBillParam {
	return &ListBillParam{}
}

// Initialize init ListBillParam
func (p *ListBillParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ListBillParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *ListBillParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.Year) {
		p.Year = 0
	}
	if utils.IsEmpty(p.Month) {
		p.Month = 0
	}

}

func (p *ListBillParam) validate() error {
	var errors []error

	{
		validator := define.Resources["Bill"].Commands["list"].Params["year"].ValidateFunc
		errs := validator("--year", p.Year)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["Bill"].Commands["list"].Params["month"].ValidateFunc
		errs := validator("--month", p.Month)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return utils.FlattenErrors(errors)
}

func (p *ListBillParam) ResourceDef() *schema.Resource {
	return define.Resources["Bill"]
}

func (p *ListBillParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["list"]
}

func (p *ListBillParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ListBillParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ListBillParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ListBillParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ListBillParam) SetYear(v int) {
	p.Year = v
}

func (p *ListBillParam) GetYear() int {
	return p.Year
}
func (p *ListBillParam) SetMonth(v int) {
	p.Month = v
}

func (p *ListBillParam) GetMonth() int {
	return p.Month
}
