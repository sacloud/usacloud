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

// ListIPv6Param is input parameters for the sacloud API
type ListIPv6Param struct {
	Name              []string
	Id                []sacloud.ID
	IPv6netId         sacloud.ID
	InternetId        sacloud.ID
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

// NewListIPv6Param return new ListIPv6Param
func NewListIPv6Param() *ListIPv6Param {
	return &ListIPv6Param{}
}

// Initialize init ListIPv6Param
func (p *ListIPv6Param) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ListIPv6Param) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *ListIPv6Param) fillValueToSkeleton() {
	if utils.IsEmpty(p.Name) {
		p.Name = []string{""}
	}
	if utils.IsEmpty(p.Id) {
		p.Id = []sacloud.ID{}
	}
	if utils.IsEmpty(p.IPv6netId) {
		p.IPv6netId = sacloud.ID(0)
	}
	if utils.IsEmpty(p.InternetId) {
		p.InternetId = sacloud.ID(0)
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

func (p *ListIPv6Param) validate() error {
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
		validator := define.Resources["IPv6"].Commands["list"].Params["id"].ValidateFunc
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
		validator := define.Resources["IPv6"].Commands["list"].Params["ipv6net-id"].ValidateFunc
		errs := validator("--ipv6net-id", p.IPv6netId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["IPv6"].Commands["list"].Params["internet-id"].ValidateFunc
		errs := validator("--internet-id", p.InternetId)
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

func (p *ListIPv6Param) ResourceDef() *schema.Resource {
	return define.Resources["IPv6"]
}

func (p *ListIPv6Param) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["list"]
}

func (p *ListIPv6Param) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ListIPv6Param) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ListIPv6Param) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ListIPv6Param) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ListIPv6Param) SetName(v []string) {
	p.Name = v
}

func (p *ListIPv6Param) GetName() []string {
	return p.Name
}
func (p *ListIPv6Param) SetId(v []sacloud.ID) {
	p.Id = v
}

func (p *ListIPv6Param) GetId() []sacloud.ID {
	return p.Id
}
func (p *ListIPv6Param) SetIPv6netId(v sacloud.ID) {
	p.IPv6netId = v
}

func (p *ListIPv6Param) GetIPv6netId() sacloud.ID {
	return p.IPv6netId
}
func (p *ListIPv6Param) SetInternetId(v sacloud.ID) {
	p.InternetId = v
}

func (p *ListIPv6Param) GetInternetId() sacloud.ID {
	return p.InternetId
}
func (p *ListIPv6Param) SetFrom(v int) {
	p.From = v
}

func (p *ListIPv6Param) GetFrom() int {
	return p.From
}
func (p *ListIPv6Param) SetMax(v int) {
	p.Max = v
}

func (p *ListIPv6Param) GetMax() int {
	return p.Max
}
func (p *ListIPv6Param) SetSort(v []string) {
	p.Sort = v
}

func (p *ListIPv6Param) GetSort() []string {
	return p.Sort
}
func (p *ListIPv6Param) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ListIPv6Param) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ListIPv6Param) SetParameters(v string) {
	p.Parameters = v
}

func (p *ListIPv6Param) GetParameters() string {
	return p.Parameters
}
func (p *ListIPv6Param) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ListIPv6Param) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ListIPv6Param) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *ListIPv6Param) GetParameterFile() string {
	return p.ParameterFile
}
func (p *ListIPv6Param) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ListIPv6Param) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ListIPv6Param) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ListIPv6Param) GetOutputType() string {
	return p.OutputType
}
func (p *ListIPv6Param) SetColumn(v []string) {
	p.Column = v
}

func (p *ListIPv6Param) GetColumn() []string {
	return p.Column
}
func (p *ListIPv6Param) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ListIPv6Param) GetQuiet() bool {
	return p.Quiet
}
func (p *ListIPv6Param) SetFormat(v string) {
	p.Format = v
}

func (p *ListIPv6Param) GetFormat() string {
	return p.Format
}
func (p *ListIPv6Param) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ListIPv6Param) GetFormatFile() string {
	return p.FormatFile
}
func (p *ListIPv6Param) SetQuery(v string) {
	p.Query = v
}

func (p *ListIPv6Param) GetQuery() string {
	return p.Query
}
func (p *ListIPv6Param) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ListIPv6Param) GetQueryFile() string {
	return p.QueryFile
}

// PtrAddIPv6Param is input parameters for the sacloud API
type PtrAddIPv6Param struct {
	Hostname          string
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

// NewPtrAddIPv6Param return new PtrAddIPv6Param
func NewPtrAddIPv6Param() *PtrAddIPv6Param {
	return &PtrAddIPv6Param{}
}

// Initialize init PtrAddIPv6Param
func (p *PtrAddIPv6Param) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *PtrAddIPv6Param) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *PtrAddIPv6Param) fillValueToSkeleton() {
	if utils.IsEmpty(p.Hostname) {
		p.Hostname = ""
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

func (p *PtrAddIPv6Param) validate() error {
	var errors []error

	{
		validator := validateRequired
		errs := validator("--hostname", p.Hostname)
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

func (p *PtrAddIPv6Param) ResourceDef() *schema.Resource {
	return define.Resources["IPv6"]
}

func (p *PtrAddIPv6Param) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["ptr-add"]
}

func (p *PtrAddIPv6Param) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *PtrAddIPv6Param) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *PtrAddIPv6Param) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *PtrAddIPv6Param) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *PtrAddIPv6Param) SetHostname(v string) {
	p.Hostname = v
}

func (p *PtrAddIPv6Param) GetHostname() string {
	return p.Hostname
}
func (p *PtrAddIPv6Param) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *PtrAddIPv6Param) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *PtrAddIPv6Param) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *PtrAddIPv6Param) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *PtrAddIPv6Param) SetParameters(v string) {
	p.Parameters = v
}

func (p *PtrAddIPv6Param) GetParameters() string {
	return p.Parameters
}
func (p *PtrAddIPv6Param) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *PtrAddIPv6Param) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *PtrAddIPv6Param) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *PtrAddIPv6Param) GetParameterFile() string {
	return p.ParameterFile
}
func (p *PtrAddIPv6Param) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *PtrAddIPv6Param) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *PtrAddIPv6Param) SetOutputType(v string) {
	p.OutputType = v
}

func (p *PtrAddIPv6Param) GetOutputType() string {
	return p.OutputType
}
func (p *PtrAddIPv6Param) SetColumn(v []string) {
	p.Column = v
}

func (p *PtrAddIPv6Param) GetColumn() []string {
	return p.Column
}
func (p *PtrAddIPv6Param) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *PtrAddIPv6Param) GetQuiet() bool {
	return p.Quiet
}
func (p *PtrAddIPv6Param) SetFormat(v string) {
	p.Format = v
}

func (p *PtrAddIPv6Param) GetFormat() string {
	return p.Format
}
func (p *PtrAddIPv6Param) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *PtrAddIPv6Param) GetFormatFile() string {
	return p.FormatFile
}
func (p *PtrAddIPv6Param) SetQuery(v string) {
	p.Query = v
}

func (p *PtrAddIPv6Param) GetQuery() string {
	return p.Query
}
func (p *PtrAddIPv6Param) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *PtrAddIPv6Param) GetQueryFile() string {
	return p.QueryFile
}

// PtrReadIPv6Param is input parameters for the sacloud API
type PtrReadIPv6Param struct {
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

// NewPtrReadIPv6Param return new PtrReadIPv6Param
func NewPtrReadIPv6Param() *PtrReadIPv6Param {
	return &PtrReadIPv6Param{}
}

// Initialize init PtrReadIPv6Param
func (p *PtrReadIPv6Param) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *PtrReadIPv6Param) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *PtrReadIPv6Param) fillValueToSkeleton() {
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

func (p *PtrReadIPv6Param) validate() error {
	var errors []error

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

func (p *PtrReadIPv6Param) ResourceDef() *schema.Resource {
	return define.Resources["IPv6"]
}

func (p *PtrReadIPv6Param) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["ptr-read"]
}

func (p *PtrReadIPv6Param) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *PtrReadIPv6Param) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *PtrReadIPv6Param) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *PtrReadIPv6Param) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *PtrReadIPv6Param) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *PtrReadIPv6Param) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *PtrReadIPv6Param) SetParameters(v string) {
	p.Parameters = v
}

func (p *PtrReadIPv6Param) GetParameters() string {
	return p.Parameters
}
func (p *PtrReadIPv6Param) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *PtrReadIPv6Param) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *PtrReadIPv6Param) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *PtrReadIPv6Param) GetParameterFile() string {
	return p.ParameterFile
}
func (p *PtrReadIPv6Param) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *PtrReadIPv6Param) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *PtrReadIPv6Param) SetOutputType(v string) {
	p.OutputType = v
}

func (p *PtrReadIPv6Param) GetOutputType() string {
	return p.OutputType
}
func (p *PtrReadIPv6Param) SetColumn(v []string) {
	p.Column = v
}

func (p *PtrReadIPv6Param) GetColumn() []string {
	return p.Column
}
func (p *PtrReadIPv6Param) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *PtrReadIPv6Param) GetQuiet() bool {
	return p.Quiet
}
func (p *PtrReadIPv6Param) SetFormat(v string) {
	p.Format = v
}

func (p *PtrReadIPv6Param) GetFormat() string {
	return p.Format
}
func (p *PtrReadIPv6Param) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *PtrReadIPv6Param) GetFormatFile() string {
	return p.FormatFile
}
func (p *PtrReadIPv6Param) SetQuery(v string) {
	p.Query = v
}

func (p *PtrReadIPv6Param) GetQuery() string {
	return p.Query
}
func (p *PtrReadIPv6Param) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *PtrReadIPv6Param) GetQueryFile() string {
	return p.QueryFile
}

// PtrUpdateIPv6Param is input parameters for the sacloud API
type PtrUpdateIPv6Param struct {
	Hostname          string
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

// NewPtrUpdateIPv6Param return new PtrUpdateIPv6Param
func NewPtrUpdateIPv6Param() *PtrUpdateIPv6Param {
	return &PtrUpdateIPv6Param{}
}

// Initialize init PtrUpdateIPv6Param
func (p *PtrUpdateIPv6Param) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *PtrUpdateIPv6Param) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *PtrUpdateIPv6Param) fillValueToSkeleton() {
	if utils.IsEmpty(p.Hostname) {
		p.Hostname = ""
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

func (p *PtrUpdateIPv6Param) validate() error {
	var errors []error

	{
		validator := validateRequired
		errs := validator("--hostname", p.Hostname)
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

func (p *PtrUpdateIPv6Param) ResourceDef() *schema.Resource {
	return define.Resources["IPv6"]
}

func (p *PtrUpdateIPv6Param) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["ptr-update"]
}

func (p *PtrUpdateIPv6Param) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *PtrUpdateIPv6Param) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *PtrUpdateIPv6Param) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *PtrUpdateIPv6Param) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *PtrUpdateIPv6Param) SetHostname(v string) {
	p.Hostname = v
}

func (p *PtrUpdateIPv6Param) GetHostname() string {
	return p.Hostname
}
func (p *PtrUpdateIPv6Param) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *PtrUpdateIPv6Param) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *PtrUpdateIPv6Param) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *PtrUpdateIPv6Param) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *PtrUpdateIPv6Param) SetParameters(v string) {
	p.Parameters = v
}

func (p *PtrUpdateIPv6Param) GetParameters() string {
	return p.Parameters
}
func (p *PtrUpdateIPv6Param) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *PtrUpdateIPv6Param) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *PtrUpdateIPv6Param) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *PtrUpdateIPv6Param) GetParameterFile() string {
	return p.ParameterFile
}
func (p *PtrUpdateIPv6Param) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *PtrUpdateIPv6Param) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *PtrUpdateIPv6Param) SetOutputType(v string) {
	p.OutputType = v
}

func (p *PtrUpdateIPv6Param) GetOutputType() string {
	return p.OutputType
}
func (p *PtrUpdateIPv6Param) SetColumn(v []string) {
	p.Column = v
}

func (p *PtrUpdateIPv6Param) GetColumn() []string {
	return p.Column
}
func (p *PtrUpdateIPv6Param) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *PtrUpdateIPv6Param) GetQuiet() bool {
	return p.Quiet
}
func (p *PtrUpdateIPv6Param) SetFormat(v string) {
	p.Format = v
}

func (p *PtrUpdateIPv6Param) GetFormat() string {
	return p.Format
}
func (p *PtrUpdateIPv6Param) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *PtrUpdateIPv6Param) GetFormatFile() string {
	return p.FormatFile
}
func (p *PtrUpdateIPv6Param) SetQuery(v string) {
	p.Query = v
}

func (p *PtrUpdateIPv6Param) GetQuery() string {
	return p.Query
}
func (p *PtrUpdateIPv6Param) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *PtrUpdateIPv6Param) GetQueryFile() string {
	return p.QueryFile
}

// PtrDeleteIPv6Param is input parameters for the sacloud API
type PtrDeleteIPv6Param struct {
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

// NewPtrDeleteIPv6Param return new PtrDeleteIPv6Param
func NewPtrDeleteIPv6Param() *PtrDeleteIPv6Param {
	return &PtrDeleteIPv6Param{}
}

// Initialize init PtrDeleteIPv6Param
func (p *PtrDeleteIPv6Param) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *PtrDeleteIPv6Param) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *PtrDeleteIPv6Param) fillValueToSkeleton() {
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

func (p *PtrDeleteIPv6Param) validate() error {
	var errors []error

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

func (p *PtrDeleteIPv6Param) ResourceDef() *schema.Resource {
	return define.Resources["IPv6"]
}

func (p *PtrDeleteIPv6Param) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["ptr-delete"]
}

func (p *PtrDeleteIPv6Param) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *PtrDeleteIPv6Param) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *PtrDeleteIPv6Param) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *PtrDeleteIPv6Param) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *PtrDeleteIPv6Param) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *PtrDeleteIPv6Param) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *PtrDeleteIPv6Param) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *PtrDeleteIPv6Param) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *PtrDeleteIPv6Param) SetParameters(v string) {
	p.Parameters = v
}

func (p *PtrDeleteIPv6Param) GetParameters() string {
	return p.Parameters
}
func (p *PtrDeleteIPv6Param) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *PtrDeleteIPv6Param) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *PtrDeleteIPv6Param) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *PtrDeleteIPv6Param) GetParameterFile() string {
	return p.ParameterFile
}
func (p *PtrDeleteIPv6Param) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *PtrDeleteIPv6Param) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *PtrDeleteIPv6Param) SetOutputType(v string) {
	p.OutputType = v
}

func (p *PtrDeleteIPv6Param) GetOutputType() string {
	return p.OutputType
}
func (p *PtrDeleteIPv6Param) SetColumn(v []string) {
	p.Column = v
}

func (p *PtrDeleteIPv6Param) GetColumn() []string {
	return p.Column
}
func (p *PtrDeleteIPv6Param) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *PtrDeleteIPv6Param) GetQuiet() bool {
	return p.Quiet
}
func (p *PtrDeleteIPv6Param) SetFormat(v string) {
	p.Format = v
}

func (p *PtrDeleteIPv6Param) GetFormat() string {
	return p.Format
}
func (p *PtrDeleteIPv6Param) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *PtrDeleteIPv6Param) GetFormatFile() string {
	return p.FormatFile
}
func (p *PtrDeleteIPv6Param) SetQuery(v string) {
	p.Query = v
}

func (p *PtrDeleteIPv6Param) GetQuery() string {
	return p.Query
}
func (p *PtrDeleteIPv6Param) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *PtrDeleteIPv6Param) GetQueryFile() string {
	return p.QueryFile
}
