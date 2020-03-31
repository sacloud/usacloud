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

// ListRegionParam is input parameters for the sacloud API
type ListRegionParam struct {
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

// NewListRegionParam return new ListRegionParam
func NewListRegionParam() *ListRegionParam {
	return &ListRegionParam{}
}

// Initialize init ListRegionParam
func (p *ListRegionParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ListRegionParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *ListRegionParam) fillValueToSkeleton() {
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

func (p *ListRegionParam) validate() error {
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
		validator := define.Resources["Region"].Commands["list"].Params["id"].ValidateFunc
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
		errs := validateParameterOptions(p)
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
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	return utils.FlattenErrors(errors)
}

func (p *ListRegionParam) ResourceDef() *schema.Resource {
	return define.Resources["Region"]
}

func (p *ListRegionParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["list"]
}

func (p *ListRegionParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ListRegionParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ListRegionParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ListRegionParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ListRegionParam) SetName(v []string) {
	p.Name = v
}

func (p *ListRegionParam) GetName() []string {
	return p.Name
}
func (p *ListRegionParam) SetId(v []sacloud.ID) {
	p.Id = v
}

func (p *ListRegionParam) GetId() []sacloud.ID {
	return p.Id
}
func (p *ListRegionParam) SetFrom(v int) {
	p.From = v
}

func (p *ListRegionParam) GetFrom() int {
	return p.From
}
func (p *ListRegionParam) SetMax(v int) {
	p.Max = v
}

func (p *ListRegionParam) GetMax() int {
	return p.Max
}
func (p *ListRegionParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListRegionParam) GetSort() []string {
	return p.Sort
}
func (p *ListRegionParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ListRegionParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ListRegionParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *ListRegionParam) GetParameters() string {
	return p.Parameters
}
func (p *ListRegionParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ListRegionParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ListRegionParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *ListRegionParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *ListRegionParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ListRegionParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ListRegionParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ListRegionParam) GetOutputType() string {
	return p.OutputType
}
func (p *ListRegionParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ListRegionParam) GetColumn() []string {
	return p.Column
}
func (p *ListRegionParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ListRegionParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ListRegionParam) SetFormat(v string) {
	p.Format = v
}

func (p *ListRegionParam) GetFormat() string {
	return p.Format
}
func (p *ListRegionParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ListRegionParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ListRegionParam) SetQuery(v string) {
	p.Query = v
}

func (p *ListRegionParam) GetQuery() string {
	return p.Query
}
func (p *ListRegionParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ListRegionParam) GetQueryFile() string {
	return p.QueryFile
}

// Changed usacloud v0系との互換性維持のための実装
func (p *ListRegionParam) Changed(name string) bool {
	return p.input.Changed(name)
}

// ReadRegionParam is input parameters for the sacloud API
type ReadRegionParam struct {
	Assumeyes         bool
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
	Id                sacloud.ID

	input Input
}

// NewReadRegionParam return new ReadRegionParam
func NewReadRegionParam() *ReadRegionParam {
	return &ReadRegionParam{}
}

// Initialize init ReadRegionParam
func (p *ReadRegionParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ReadRegionParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *ReadRegionParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.Assumeyes) {
		p.Assumeyes = false
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
	if utils.IsEmpty(p.Id) {
		p.Id = sacloud.ID(0)
	}

}

func (p *ReadRegionParam) validate() error {
	var errors []error

	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Region"].Commands["read"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		errs := validateParameterOptions(p)
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
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	return utils.FlattenErrors(errors)
}

func (p *ReadRegionParam) ResourceDef() *schema.Resource {
	return define.Resources["Region"]
}

func (p *ReadRegionParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["read"]
}

func (p *ReadRegionParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ReadRegionParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ReadRegionParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ReadRegionParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ReadRegionParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *ReadRegionParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *ReadRegionParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ReadRegionParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ReadRegionParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *ReadRegionParam) GetParameters() string {
	return p.Parameters
}
func (p *ReadRegionParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ReadRegionParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ReadRegionParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *ReadRegionParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *ReadRegionParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ReadRegionParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ReadRegionParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ReadRegionParam) GetOutputType() string {
	return p.OutputType
}
func (p *ReadRegionParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ReadRegionParam) GetColumn() []string {
	return p.Column
}
func (p *ReadRegionParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ReadRegionParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ReadRegionParam) SetFormat(v string) {
	p.Format = v
}

func (p *ReadRegionParam) GetFormat() string {
	return p.Format
}
func (p *ReadRegionParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ReadRegionParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ReadRegionParam) SetQuery(v string) {
	p.Query = v
}

func (p *ReadRegionParam) GetQuery() string {
	return p.Query
}
func (p *ReadRegionParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ReadRegionParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *ReadRegionParam) SetId(v sacloud.ID) {
	p.Id = v
}

func (p *ReadRegionParam) GetId() sacloud.ID {
	return p.Id
}

// Changed usacloud v0系との互換性維持のための実装
func (p *ReadRegionParam) Changed(name string) bool {
	return p.input.Changed(name)
}
