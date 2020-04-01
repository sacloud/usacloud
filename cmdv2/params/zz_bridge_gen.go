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
	v0params "github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/pkg/utils"
	"github.com/sacloud/usacloud/pkg/validation"
	"github.com/sacloud/usacloud/schema"
)

// ListBridgeParam is input parameters for the sacloud API
type ListBridgeParam struct {
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

// NewListBridgeParam return new ListBridgeParam
func NewListBridgeParam() *ListBridgeParam {
	return &ListBridgeParam{}
}

// Initialize init ListBridgeParam
func (p *ListBridgeParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ListBridgeParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

// FillValueToSkeleton fills empty value to the parameter
func (p *ListBridgeParam) FillValueToSkeleton() {
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

func (p *ListBridgeParam) validate() error {
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
		validator := define.Resources["Bridge"].Commands["list"].Params["id"].ValidateFunc
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

func (p *ListBridgeParam) ResourceDef() *schema.Resource {
	return define.Resources["Bridge"]
}

func (p *ListBridgeParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["list"]
}

func (p *ListBridgeParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ListBridgeParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ListBridgeParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ListBridgeParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

/*
 * v0系との互換性維持のための実装
 */
func (p *ListBridgeParam) GetResourceDef() *schema.Resource {
	return define.Resources["Bridge"]
}

func (p *ListBridgeParam) GetCommandDef() *schema.Command {
	return p.ResourceDef().Commands["list"]
}

func (p *ListBridgeParam) GetIncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ListBridgeParam) GetExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ListBridgeParam) GetTableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ListBridgeParam) GetColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ListBridgeParam) SetName(v []string) {
	p.Name = v
}

func (p *ListBridgeParam) GetName() []string {
	return p.Name
}
func (p *ListBridgeParam) SetId(v []sacloud.ID) {
	p.Id = v
}

func (p *ListBridgeParam) GetId() []sacloud.ID {
	return p.Id
}
func (p *ListBridgeParam) SetFrom(v int) {
	p.From = v
}

func (p *ListBridgeParam) GetFrom() int {
	return p.From
}
func (p *ListBridgeParam) SetMax(v int) {
	p.Max = v
}

func (p *ListBridgeParam) GetMax() int {
	return p.Max
}
func (p *ListBridgeParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListBridgeParam) GetSort() []string {
	return p.Sort
}
func (p *ListBridgeParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ListBridgeParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ListBridgeParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *ListBridgeParam) GetParameters() string {
	return p.Parameters
}
func (p *ListBridgeParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ListBridgeParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ListBridgeParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *ListBridgeParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *ListBridgeParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ListBridgeParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ListBridgeParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ListBridgeParam) GetOutputType() string {
	return p.OutputType
}
func (p *ListBridgeParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ListBridgeParam) GetColumn() []string {
	return p.Column
}
func (p *ListBridgeParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ListBridgeParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ListBridgeParam) SetFormat(v string) {
	p.Format = v
}

func (p *ListBridgeParam) GetFormat() string {
	return p.Format
}
func (p *ListBridgeParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ListBridgeParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ListBridgeParam) SetQuery(v string) {
	p.Query = v
}

func (p *ListBridgeParam) GetQuery() string {
	return p.Query
}
func (p *ListBridgeParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ListBridgeParam) GetQueryFile() string {
	return p.QueryFile
}

// Changed usacloud v0系との互換性維持のための実装
func (p *ListBridgeParam) Changed(name string) bool {
	return p.input.Changed(name)
}

func (p *ListBridgeParam) ToV0() *v0params.ListBridgeParam {
	return &v0params.ListBridgeParam{
		Name:              p.Name,
		Id:                p.Id,
		From:              p.From,
		Max:               p.Max,
		Sort:              p.Sort,
		ParamTemplate:     p.ParamTemplate,
		Parameters:        p.Parameters,
		ParamTemplateFile: p.ParamTemplateFile,
		ParameterFile:     p.ParameterFile,
		GenerateSkeleton:  p.GenerateSkeleton,
		OutputType:        p.OutputType,
		Column:            p.Column,
		Quiet:             p.Quiet,
		Format:            p.Format,
		FormatFile:        p.FormatFile,
		Query:             p.Query,
		QueryFile:         p.QueryFile,
	}
}

// CreateBridgeParam is input parameters for the sacloud API
type CreateBridgeParam struct {
	Name              string
	Description       string
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

	input Input
}

// NewCreateBridgeParam return new CreateBridgeParam
func NewCreateBridgeParam() *CreateBridgeParam {
	return &CreateBridgeParam{}
}

// Initialize init CreateBridgeParam
func (p *CreateBridgeParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *CreateBridgeParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

// FillValueToSkeleton fills empty value to the parameter
func (p *CreateBridgeParam) FillValueToSkeleton() {
	if utils.IsEmpty(p.Name) {
		p.Name = ""
	}
	if utils.IsEmpty(p.Description) {
		p.Description = ""
	}
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

}

func (p *CreateBridgeParam) validate() error {
	var errors []error

	{
		validator := validateRequired
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Bridge"].Commands["create"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["Bridge"].Commands["create"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
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

func (p *CreateBridgeParam) ResourceDef() *schema.Resource {
	return define.Resources["Bridge"]
}

func (p *CreateBridgeParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["create"]
}

func (p *CreateBridgeParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *CreateBridgeParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *CreateBridgeParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *CreateBridgeParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

/*
 * v0系との互換性維持のための実装
 */
func (p *CreateBridgeParam) GetResourceDef() *schema.Resource {
	return define.Resources["Bridge"]
}

func (p *CreateBridgeParam) GetCommandDef() *schema.Command {
	return p.ResourceDef().Commands["create"]
}

func (p *CreateBridgeParam) GetIncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *CreateBridgeParam) GetExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *CreateBridgeParam) GetTableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *CreateBridgeParam) GetColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *CreateBridgeParam) SetName(v string) {
	p.Name = v
}

func (p *CreateBridgeParam) GetName() string {
	return p.Name
}
func (p *CreateBridgeParam) SetDescription(v string) {
	p.Description = v
}

func (p *CreateBridgeParam) GetDescription() string {
	return p.Description
}
func (p *CreateBridgeParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *CreateBridgeParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *CreateBridgeParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *CreateBridgeParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *CreateBridgeParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *CreateBridgeParam) GetParameters() string {
	return p.Parameters
}
func (p *CreateBridgeParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *CreateBridgeParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *CreateBridgeParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *CreateBridgeParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *CreateBridgeParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *CreateBridgeParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *CreateBridgeParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *CreateBridgeParam) GetOutputType() string {
	return p.OutputType
}
func (p *CreateBridgeParam) SetColumn(v []string) {
	p.Column = v
}

func (p *CreateBridgeParam) GetColumn() []string {
	return p.Column
}
func (p *CreateBridgeParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *CreateBridgeParam) GetQuiet() bool {
	return p.Quiet
}
func (p *CreateBridgeParam) SetFormat(v string) {
	p.Format = v
}

func (p *CreateBridgeParam) GetFormat() string {
	return p.Format
}
func (p *CreateBridgeParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *CreateBridgeParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *CreateBridgeParam) SetQuery(v string) {
	p.Query = v
}

func (p *CreateBridgeParam) GetQuery() string {
	return p.Query
}
func (p *CreateBridgeParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *CreateBridgeParam) GetQueryFile() string {
	return p.QueryFile
}

// Changed usacloud v0系との互換性維持のための実装
func (p *CreateBridgeParam) Changed(name string) bool {
	return p.input.Changed(name)
}

func (p *CreateBridgeParam) ToV0() *v0params.CreateBridgeParam {
	return &v0params.CreateBridgeParam{
		Name:              p.Name,
		Description:       p.Description,
		Assumeyes:         p.Assumeyes,
		ParamTemplate:     p.ParamTemplate,
		Parameters:        p.Parameters,
		ParamTemplateFile: p.ParamTemplateFile,
		ParameterFile:     p.ParameterFile,
		GenerateSkeleton:  p.GenerateSkeleton,
		OutputType:        p.OutputType,
		Column:            p.Column,
		Quiet:             p.Quiet,
		Format:            p.Format,
		FormatFile:        p.FormatFile,
		Query:             p.Query,
		QueryFile:         p.QueryFile,
	}
}

// ReadBridgeParam is input parameters for the sacloud API
type ReadBridgeParam struct {
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

// NewReadBridgeParam return new ReadBridgeParam
func NewReadBridgeParam() *ReadBridgeParam {
	return &ReadBridgeParam{}
}

// Initialize init ReadBridgeParam
func (p *ReadBridgeParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ReadBridgeParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

// FillValueToSkeleton fills empty value to the parameter
func (p *ReadBridgeParam) FillValueToSkeleton() {
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

func (p *ReadBridgeParam) validate() error {
	var errors []error

	{
		validator := validateSakuraID
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

func (p *ReadBridgeParam) ResourceDef() *schema.Resource {
	return define.Resources["Bridge"]
}

func (p *ReadBridgeParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["read"]
}

func (p *ReadBridgeParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ReadBridgeParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ReadBridgeParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ReadBridgeParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

/*
 * v0系との互換性維持のための実装
 */
func (p *ReadBridgeParam) GetResourceDef() *schema.Resource {
	return define.Resources["Bridge"]
}

func (p *ReadBridgeParam) GetCommandDef() *schema.Command {
	return p.ResourceDef().Commands["read"]
}

func (p *ReadBridgeParam) GetIncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ReadBridgeParam) GetExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ReadBridgeParam) GetTableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ReadBridgeParam) GetColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ReadBridgeParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ReadBridgeParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ReadBridgeParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *ReadBridgeParam) GetParameters() string {
	return p.Parameters
}
func (p *ReadBridgeParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ReadBridgeParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ReadBridgeParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *ReadBridgeParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *ReadBridgeParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ReadBridgeParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ReadBridgeParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ReadBridgeParam) GetOutputType() string {
	return p.OutputType
}
func (p *ReadBridgeParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ReadBridgeParam) GetColumn() []string {
	return p.Column
}
func (p *ReadBridgeParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ReadBridgeParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ReadBridgeParam) SetFormat(v string) {
	p.Format = v
}

func (p *ReadBridgeParam) GetFormat() string {
	return p.Format
}
func (p *ReadBridgeParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ReadBridgeParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ReadBridgeParam) SetQuery(v string) {
	p.Query = v
}

func (p *ReadBridgeParam) GetQuery() string {
	return p.Query
}
func (p *ReadBridgeParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ReadBridgeParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *ReadBridgeParam) SetId(v sacloud.ID) {
	p.Id = v
}

func (p *ReadBridgeParam) GetId() sacloud.ID {
	return p.Id
}

// Changed usacloud v0系との互換性維持のための実装
func (p *ReadBridgeParam) Changed(name string) bool {
	return p.input.Changed(name)
}

func (p *ReadBridgeParam) ToV0() *v0params.ReadBridgeParam {
	return &v0params.ReadBridgeParam{
		ParamTemplate:     p.ParamTemplate,
		Parameters:        p.Parameters,
		ParamTemplateFile: p.ParamTemplateFile,
		ParameterFile:     p.ParameterFile,
		GenerateSkeleton:  p.GenerateSkeleton,
		OutputType:        p.OutputType,
		Column:            p.Column,
		Quiet:             p.Quiet,
		Format:            p.Format,
		FormatFile:        p.FormatFile,
		Query:             p.Query,
		QueryFile:         p.QueryFile,
		Id:                p.Id,
	}
}

// UpdateBridgeParam is input parameters for the sacloud API
type UpdateBridgeParam struct {
	Name              string
	Description       string
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

// NewUpdateBridgeParam return new UpdateBridgeParam
func NewUpdateBridgeParam() *UpdateBridgeParam {
	return &UpdateBridgeParam{}
}

// Initialize init UpdateBridgeParam
func (p *UpdateBridgeParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *UpdateBridgeParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

// FillValueToSkeleton fills empty value to the parameter
func (p *UpdateBridgeParam) FillValueToSkeleton() {
	if utils.IsEmpty(p.Name) {
		p.Name = ""
	}
	if utils.IsEmpty(p.Description) {
		p.Description = ""
	}
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

func (p *UpdateBridgeParam) validate() error {
	var errors []error

	{
		validator := define.Resources["Bridge"].Commands["update"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["Bridge"].Commands["update"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := validateSakuraID
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

func (p *UpdateBridgeParam) ResourceDef() *schema.Resource {
	return define.Resources["Bridge"]
}

func (p *UpdateBridgeParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["update"]
}

func (p *UpdateBridgeParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *UpdateBridgeParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *UpdateBridgeParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *UpdateBridgeParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

/*
 * v0系との互換性維持のための実装
 */
func (p *UpdateBridgeParam) GetResourceDef() *schema.Resource {
	return define.Resources["Bridge"]
}

func (p *UpdateBridgeParam) GetCommandDef() *schema.Command {
	return p.ResourceDef().Commands["update"]
}

func (p *UpdateBridgeParam) GetIncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *UpdateBridgeParam) GetExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *UpdateBridgeParam) GetTableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *UpdateBridgeParam) GetColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *UpdateBridgeParam) SetName(v string) {
	p.Name = v
}

func (p *UpdateBridgeParam) GetName() string {
	return p.Name
}
func (p *UpdateBridgeParam) SetDescription(v string) {
	p.Description = v
}

func (p *UpdateBridgeParam) GetDescription() string {
	return p.Description
}
func (p *UpdateBridgeParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *UpdateBridgeParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *UpdateBridgeParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *UpdateBridgeParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *UpdateBridgeParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *UpdateBridgeParam) GetParameters() string {
	return p.Parameters
}
func (p *UpdateBridgeParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *UpdateBridgeParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *UpdateBridgeParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *UpdateBridgeParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *UpdateBridgeParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *UpdateBridgeParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *UpdateBridgeParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *UpdateBridgeParam) GetOutputType() string {
	return p.OutputType
}
func (p *UpdateBridgeParam) SetColumn(v []string) {
	p.Column = v
}

func (p *UpdateBridgeParam) GetColumn() []string {
	return p.Column
}
func (p *UpdateBridgeParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *UpdateBridgeParam) GetQuiet() bool {
	return p.Quiet
}
func (p *UpdateBridgeParam) SetFormat(v string) {
	p.Format = v
}

func (p *UpdateBridgeParam) GetFormat() string {
	return p.Format
}
func (p *UpdateBridgeParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *UpdateBridgeParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *UpdateBridgeParam) SetQuery(v string) {
	p.Query = v
}

func (p *UpdateBridgeParam) GetQuery() string {
	return p.Query
}
func (p *UpdateBridgeParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *UpdateBridgeParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *UpdateBridgeParam) SetId(v sacloud.ID) {
	p.Id = v
}

func (p *UpdateBridgeParam) GetId() sacloud.ID {
	return p.Id
}

// Changed usacloud v0系との互換性維持のための実装
func (p *UpdateBridgeParam) Changed(name string) bool {
	return p.input.Changed(name)
}

func (p *UpdateBridgeParam) ToV0() *v0params.UpdateBridgeParam {
	return &v0params.UpdateBridgeParam{
		Name:              p.Name,
		Description:       p.Description,
		Assumeyes:         p.Assumeyes,
		ParamTemplate:     p.ParamTemplate,
		Parameters:        p.Parameters,
		ParamTemplateFile: p.ParamTemplateFile,
		ParameterFile:     p.ParameterFile,
		GenerateSkeleton:  p.GenerateSkeleton,
		OutputType:        p.OutputType,
		Column:            p.Column,
		Quiet:             p.Quiet,
		Format:            p.Format,
		FormatFile:        p.FormatFile,
		Query:             p.Query,
		QueryFile:         p.QueryFile,
		Id:                p.Id,
	}
}

// DeleteBridgeParam is input parameters for the sacloud API
type DeleteBridgeParam struct {
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

// NewDeleteBridgeParam return new DeleteBridgeParam
func NewDeleteBridgeParam() *DeleteBridgeParam {
	return &DeleteBridgeParam{}
}

// Initialize init DeleteBridgeParam
func (p *DeleteBridgeParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *DeleteBridgeParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

// FillValueToSkeleton fills empty value to the parameter
func (p *DeleteBridgeParam) FillValueToSkeleton() {
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

func (p *DeleteBridgeParam) validate() error {
	var errors []error

	{
		validator := validateSakuraID
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

func (p *DeleteBridgeParam) ResourceDef() *schema.Resource {
	return define.Resources["Bridge"]
}

func (p *DeleteBridgeParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["delete"]
}

func (p *DeleteBridgeParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *DeleteBridgeParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *DeleteBridgeParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *DeleteBridgeParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

/*
 * v0系との互換性維持のための実装
 */
func (p *DeleteBridgeParam) GetResourceDef() *schema.Resource {
	return define.Resources["Bridge"]
}

func (p *DeleteBridgeParam) GetCommandDef() *schema.Command {
	return p.ResourceDef().Commands["delete"]
}

func (p *DeleteBridgeParam) GetIncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *DeleteBridgeParam) GetExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *DeleteBridgeParam) GetTableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *DeleteBridgeParam) GetColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *DeleteBridgeParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *DeleteBridgeParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *DeleteBridgeParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *DeleteBridgeParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *DeleteBridgeParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *DeleteBridgeParam) GetParameters() string {
	return p.Parameters
}
func (p *DeleteBridgeParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *DeleteBridgeParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *DeleteBridgeParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *DeleteBridgeParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *DeleteBridgeParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *DeleteBridgeParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *DeleteBridgeParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *DeleteBridgeParam) GetOutputType() string {
	return p.OutputType
}
func (p *DeleteBridgeParam) SetColumn(v []string) {
	p.Column = v
}

func (p *DeleteBridgeParam) GetColumn() []string {
	return p.Column
}
func (p *DeleteBridgeParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *DeleteBridgeParam) GetQuiet() bool {
	return p.Quiet
}
func (p *DeleteBridgeParam) SetFormat(v string) {
	p.Format = v
}

func (p *DeleteBridgeParam) GetFormat() string {
	return p.Format
}
func (p *DeleteBridgeParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *DeleteBridgeParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *DeleteBridgeParam) SetQuery(v string) {
	p.Query = v
}

func (p *DeleteBridgeParam) GetQuery() string {
	return p.Query
}
func (p *DeleteBridgeParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *DeleteBridgeParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *DeleteBridgeParam) SetId(v sacloud.ID) {
	p.Id = v
}

func (p *DeleteBridgeParam) GetId() sacloud.ID {
	return p.Id
}

// Changed usacloud v0系との互換性維持のための実装
func (p *DeleteBridgeParam) Changed(name string) bool {
	return p.input.Changed(name)
}

func (p *DeleteBridgeParam) ToV0() *v0params.DeleteBridgeParam {
	return &v0params.DeleteBridgeParam{
		Assumeyes:         p.Assumeyes,
		ParamTemplate:     p.ParamTemplate,
		Parameters:        p.Parameters,
		ParamTemplateFile: p.ParamTemplateFile,
		ParameterFile:     p.ParameterFile,
		GenerateSkeleton:  p.GenerateSkeleton,
		OutputType:        p.OutputType,
		Column:            p.Column,
		Quiet:             p.Quiet,
		Format:            p.Format,
		FormatFile:        p.FormatFile,
		Query:             p.Query,
		QueryFile:         p.QueryFile,
		Id:                p.Id,
	}
}
