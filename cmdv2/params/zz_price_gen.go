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

// ListPriceParam is input parameters for the sacloud API
type ListPriceParam struct {
	Name              []string
	Id                []sacloud.ID
	From              int
	Max               int
	Sort              []string
	ParamTemplate     string
	Parameters        string
	ParamTemplateFile string
	ParameterFile     string
	GenerateSkeleton  bool
	OutputType        string
	Column            []string
	Quiet             bool
	Format            string
	FormatFile        string
	Query             string
	QueryFile         string

	input Input
}

// NewListPriceParam return new ListPriceParam
func NewListPriceParam() *ListPriceParam {
	return &ListPriceParam{}
}

// Initialize init ListPriceParam
func (p *ListPriceParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ListPriceParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *ListPriceParam) fillValueToSkeleton() {
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
	if utils.IsEmpty(p.Sort) {
		p.Sort = []string{""}
	}
	if utils.IsEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if utils.IsEmpty(p.Parameters) {
		p.Parameters = ""
	}
	if utils.IsEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if utils.IsEmpty(p.ParameterFile) {
		p.ParameterFile = ""
	}
	if utils.IsEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if utils.IsEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if utils.IsEmpty(p.Column) {
		p.Column = []string{""}
	}
	if utils.IsEmpty(p.Quiet) {
		p.Quiet = false
	}
	if utils.IsEmpty(p.Format) {
		p.Format = ""
	}
	if utils.IsEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if utils.IsEmpty(p.Query) {
		p.Query = ""
	}
	if utils.IsEmpty(p.QueryFile) {
		p.QueryFile = ""
	}

}

func (p *ListPriceParam) validate() error {
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
		validator := define.Resources["Price"].Commands["list"].Params["id"].ValidateFunc
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

	{
		validator := schema.ValidateInStrValues(define.AllowOutputTypes...)
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateInputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	return utils.FlattenErrors(errors)
}

func (p *ListPriceParam) ResourceDef() *schema.Resource {
	return define.Resources["Price"]
}

func (p *ListPriceParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["list"]
}

func (p *ListPriceParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ListPriceParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ListPriceParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ListPriceParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ListPriceParam) SetName(v []string) {
	p.Name = v
}

func (p *ListPriceParam) GetName() []string {
	return p.Name
}
func (p *ListPriceParam) SetId(v []sacloud.ID) {
	p.Id = v
}

func (p *ListPriceParam) GetId() []sacloud.ID {
	return p.Id
}
func (p *ListPriceParam) SetFrom(v int) {
	p.From = v
}

func (p *ListPriceParam) GetFrom() int {
	return p.From
}
func (p *ListPriceParam) SetMax(v int) {
	p.Max = v
}

func (p *ListPriceParam) GetMax() int {
	return p.Max
}
func (p *ListPriceParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListPriceParam) GetSort() []string {
	return p.Sort
}
func (p *ListPriceParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ListPriceParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ListPriceParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *ListPriceParam) GetParameters() string {
	return p.Parameters
}
func (p *ListPriceParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ListPriceParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ListPriceParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *ListPriceParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *ListPriceParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ListPriceParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ListPriceParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ListPriceParam) GetOutputType() string {
	return p.OutputType
}
func (p *ListPriceParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ListPriceParam) GetColumn() []string {
	return p.Column
}
func (p *ListPriceParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ListPriceParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ListPriceParam) SetFormat(v string) {
	p.Format = v
}

func (p *ListPriceParam) GetFormat() string {
	return p.Format
}
func (p *ListPriceParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ListPriceParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ListPriceParam) SetQuery(v string) {
	p.Query = v
}

func (p *ListPriceParam) GetQuery() string {
	return p.Query
}
func (p *ListPriceParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ListPriceParam) GetQueryFile() string {
	return p.QueryFile
}
