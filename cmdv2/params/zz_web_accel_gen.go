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

// ListWebAccelParam is input parameters for the sacloud API
type ListWebAccelParam struct {
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

// NewListWebAccelParam return new ListWebAccelParam
func NewListWebAccelParam() *ListWebAccelParam {
	return &ListWebAccelParam{}
}

// Initialize init ListWebAccelParam
func (p *ListWebAccelParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ListWebAccelParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *ListWebAccelParam) fillValueToSkeleton() {
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

func (p *ListWebAccelParam) validate() error {
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

func (p *ListWebAccelParam) ResourceDef() *schema.Resource {
	return define.Resources["WebAccel"]
}

func (p *ListWebAccelParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["list"]
}

func (p *ListWebAccelParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ListWebAccelParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ListWebAccelParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ListWebAccelParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ListWebAccelParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ListWebAccelParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ListWebAccelParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *ListWebAccelParam) GetParameters() string {
	return p.Parameters
}
func (p *ListWebAccelParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ListWebAccelParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ListWebAccelParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *ListWebAccelParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *ListWebAccelParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ListWebAccelParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ListWebAccelParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ListWebAccelParam) GetOutputType() string {
	return p.OutputType
}
func (p *ListWebAccelParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ListWebAccelParam) GetColumn() []string {
	return p.Column
}
func (p *ListWebAccelParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ListWebAccelParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ListWebAccelParam) SetFormat(v string) {
	p.Format = v
}

func (p *ListWebAccelParam) GetFormat() string {
	return p.Format
}
func (p *ListWebAccelParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ListWebAccelParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ListWebAccelParam) SetQuery(v string) {
	p.Query = v
}

func (p *ListWebAccelParam) GetQuery() string {
	return p.Query
}
func (p *ListWebAccelParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ListWebAccelParam) GetQueryFile() string {
	return p.QueryFile
}

// ReadWebAccelParam is input parameters for the sacloud API
type ReadWebAccelParam struct {
	Selector          []string
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

// NewReadWebAccelParam return new ReadWebAccelParam
func NewReadWebAccelParam() *ReadWebAccelParam {
	return &ReadWebAccelParam{}
}

// Initialize init ReadWebAccelParam
func (p *ReadWebAccelParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ReadWebAccelParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *ReadWebAccelParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.Selector) {
		p.Selector = []string{""}
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

func (p *ReadWebAccelParam) validate() error {
	var errors []error

	{
		validator := validateSakuraID
		errs := validator("--id", p.Id)
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

func (p *ReadWebAccelParam) ResourceDef() *schema.Resource {
	return define.Resources["WebAccel"]
}

func (p *ReadWebAccelParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["read"]
}

func (p *ReadWebAccelParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ReadWebAccelParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ReadWebAccelParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ReadWebAccelParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ReadWebAccelParam) SetSelector(v []string) {
	p.Selector = v
}

func (p *ReadWebAccelParam) GetSelector() []string {
	return p.Selector
}
func (p *ReadWebAccelParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ReadWebAccelParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ReadWebAccelParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *ReadWebAccelParam) GetParameters() string {
	return p.Parameters
}
func (p *ReadWebAccelParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ReadWebAccelParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ReadWebAccelParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *ReadWebAccelParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *ReadWebAccelParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ReadWebAccelParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ReadWebAccelParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ReadWebAccelParam) GetOutputType() string {
	return p.OutputType
}
func (p *ReadWebAccelParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ReadWebAccelParam) GetColumn() []string {
	return p.Column
}
func (p *ReadWebAccelParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ReadWebAccelParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ReadWebAccelParam) SetFormat(v string) {
	p.Format = v
}

func (p *ReadWebAccelParam) GetFormat() string {
	return p.Format
}
func (p *ReadWebAccelParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ReadWebAccelParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ReadWebAccelParam) SetQuery(v string) {
	p.Query = v
}

func (p *ReadWebAccelParam) GetQuery() string {
	return p.Query
}
func (p *ReadWebAccelParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ReadWebAccelParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *ReadWebAccelParam) SetId(v sacloud.ID) {
	p.Id = v
}

func (p *ReadWebAccelParam) GetId() sacloud.ID {
	return p.Id
}

// CertificateInfoWebAccelParam is input parameters for the sacloud API
type CertificateInfoWebAccelParam struct {
	Selector          []string
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

// NewCertificateInfoWebAccelParam return new CertificateInfoWebAccelParam
func NewCertificateInfoWebAccelParam() *CertificateInfoWebAccelParam {
	return &CertificateInfoWebAccelParam{}
}

// Initialize init CertificateInfoWebAccelParam
func (p *CertificateInfoWebAccelParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *CertificateInfoWebAccelParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *CertificateInfoWebAccelParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.Selector) {
		p.Selector = []string{""}
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

func (p *CertificateInfoWebAccelParam) validate() error {
	var errors []error

	{
		validator := validateSakuraID
		errs := validator("--id", p.Id)
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

func (p *CertificateInfoWebAccelParam) ResourceDef() *schema.Resource {
	return define.Resources["WebAccel"]
}

func (p *CertificateInfoWebAccelParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["certificate-info"]
}

func (p *CertificateInfoWebAccelParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *CertificateInfoWebAccelParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *CertificateInfoWebAccelParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *CertificateInfoWebAccelParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *CertificateInfoWebAccelParam) SetSelector(v []string) {
	p.Selector = v
}

func (p *CertificateInfoWebAccelParam) GetSelector() []string {
	return p.Selector
}
func (p *CertificateInfoWebAccelParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *CertificateInfoWebAccelParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *CertificateInfoWebAccelParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *CertificateInfoWebAccelParam) GetParameters() string {
	return p.Parameters
}
func (p *CertificateInfoWebAccelParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *CertificateInfoWebAccelParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *CertificateInfoWebAccelParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *CertificateInfoWebAccelParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *CertificateInfoWebAccelParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *CertificateInfoWebAccelParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *CertificateInfoWebAccelParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *CertificateInfoWebAccelParam) GetOutputType() string {
	return p.OutputType
}
func (p *CertificateInfoWebAccelParam) SetColumn(v []string) {
	p.Column = v
}

func (p *CertificateInfoWebAccelParam) GetColumn() []string {
	return p.Column
}
func (p *CertificateInfoWebAccelParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *CertificateInfoWebAccelParam) GetQuiet() bool {
	return p.Quiet
}
func (p *CertificateInfoWebAccelParam) SetFormat(v string) {
	p.Format = v
}

func (p *CertificateInfoWebAccelParam) GetFormat() string {
	return p.Format
}
func (p *CertificateInfoWebAccelParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *CertificateInfoWebAccelParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *CertificateInfoWebAccelParam) SetQuery(v string) {
	p.Query = v
}

func (p *CertificateInfoWebAccelParam) GetQuery() string {
	return p.Query
}
func (p *CertificateInfoWebAccelParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *CertificateInfoWebAccelParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *CertificateInfoWebAccelParam) SetId(v sacloud.ID) {
	p.Id = v
}

func (p *CertificateInfoWebAccelParam) GetId() sacloud.ID {
	return p.Id
}

// CertificateNewWebAccelParam is input parameters for the sacloud API
type CertificateNewWebAccelParam struct {
	Cert              string
	Key               string
	CertContent       string
	KeyContent        string
	Selector          []string
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

// NewCertificateNewWebAccelParam return new CertificateNewWebAccelParam
func NewCertificateNewWebAccelParam() *CertificateNewWebAccelParam {
	return &CertificateNewWebAccelParam{}
}

// Initialize init CertificateNewWebAccelParam
func (p *CertificateNewWebAccelParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *CertificateNewWebAccelParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *CertificateNewWebAccelParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.Cert) {
		p.Cert = ""
	}
	if utils.IsEmpty(p.Key) {
		p.Key = ""
	}
	if utils.IsEmpty(p.CertContent) {
		p.CertContent = ""
	}
	if utils.IsEmpty(p.KeyContent) {
		p.KeyContent = ""
	}
	if utils.IsEmpty(p.Selector) {
		p.Selector = []string{""}
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

func (p *CertificateNewWebAccelParam) validate() error {
	var errors []error

	{
		validator := define.Resources["WebAccel"].Commands["certificate-new"].Params["cert"].ValidateFunc
		errs := validator("--cert", p.Cert)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["WebAccel"].Commands["certificate-new"].Params["key"].ValidateFunc
		errs := validator("--key", p.Key)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		errs := validation.ConflictsWith("--cert-content", p.CertContent, map[string]interface{}{

			"--cert": p.Cert,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		errs := validation.ConflictsWith("--key-content", p.KeyContent, map[string]interface{}{

			"--key": p.Key,
		})
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

func (p *CertificateNewWebAccelParam) ResourceDef() *schema.Resource {
	return define.Resources["WebAccel"]
}

func (p *CertificateNewWebAccelParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["certificate-new"]
}

func (p *CertificateNewWebAccelParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *CertificateNewWebAccelParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *CertificateNewWebAccelParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *CertificateNewWebAccelParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *CertificateNewWebAccelParam) SetCert(v string) {
	p.Cert = v
}

func (p *CertificateNewWebAccelParam) GetCert() string {
	return p.Cert
}
func (p *CertificateNewWebAccelParam) SetKey(v string) {
	p.Key = v
}

func (p *CertificateNewWebAccelParam) GetKey() string {
	return p.Key
}
func (p *CertificateNewWebAccelParam) SetCertContent(v string) {
	p.CertContent = v
}

func (p *CertificateNewWebAccelParam) GetCertContent() string {
	return p.CertContent
}
func (p *CertificateNewWebAccelParam) SetKeyContent(v string) {
	p.KeyContent = v
}

func (p *CertificateNewWebAccelParam) GetKeyContent() string {
	return p.KeyContent
}
func (p *CertificateNewWebAccelParam) SetSelector(v []string) {
	p.Selector = v
}

func (p *CertificateNewWebAccelParam) GetSelector() []string {
	return p.Selector
}
func (p *CertificateNewWebAccelParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *CertificateNewWebAccelParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *CertificateNewWebAccelParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *CertificateNewWebAccelParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *CertificateNewWebAccelParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *CertificateNewWebAccelParam) GetParameters() string {
	return p.Parameters
}
func (p *CertificateNewWebAccelParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *CertificateNewWebAccelParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *CertificateNewWebAccelParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *CertificateNewWebAccelParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *CertificateNewWebAccelParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *CertificateNewWebAccelParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *CertificateNewWebAccelParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *CertificateNewWebAccelParam) GetOutputType() string {
	return p.OutputType
}
func (p *CertificateNewWebAccelParam) SetColumn(v []string) {
	p.Column = v
}

func (p *CertificateNewWebAccelParam) GetColumn() []string {
	return p.Column
}
func (p *CertificateNewWebAccelParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *CertificateNewWebAccelParam) GetQuiet() bool {
	return p.Quiet
}
func (p *CertificateNewWebAccelParam) SetFormat(v string) {
	p.Format = v
}

func (p *CertificateNewWebAccelParam) GetFormat() string {
	return p.Format
}
func (p *CertificateNewWebAccelParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *CertificateNewWebAccelParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *CertificateNewWebAccelParam) SetQuery(v string) {
	p.Query = v
}

func (p *CertificateNewWebAccelParam) GetQuery() string {
	return p.Query
}
func (p *CertificateNewWebAccelParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *CertificateNewWebAccelParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *CertificateNewWebAccelParam) SetId(v sacloud.ID) {
	p.Id = v
}

func (p *CertificateNewWebAccelParam) GetId() sacloud.ID {
	return p.Id
}

// CertificateUpdateWebAccelParam is input parameters for the sacloud API
type CertificateUpdateWebAccelParam struct {
	Cert              string
	Key               string
	CertContent       string
	KeyContent        string
	Selector          []string
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

// NewCertificateUpdateWebAccelParam return new CertificateUpdateWebAccelParam
func NewCertificateUpdateWebAccelParam() *CertificateUpdateWebAccelParam {
	return &CertificateUpdateWebAccelParam{}
}

// Initialize init CertificateUpdateWebAccelParam
func (p *CertificateUpdateWebAccelParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *CertificateUpdateWebAccelParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *CertificateUpdateWebAccelParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.Cert) {
		p.Cert = ""
	}
	if utils.IsEmpty(p.Key) {
		p.Key = ""
	}
	if utils.IsEmpty(p.CertContent) {
		p.CertContent = ""
	}
	if utils.IsEmpty(p.KeyContent) {
		p.KeyContent = ""
	}
	if utils.IsEmpty(p.Selector) {
		p.Selector = []string{""}
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

func (p *CertificateUpdateWebAccelParam) validate() error {
	var errors []error

	{
		validator := define.Resources["WebAccel"].Commands["certificate-update"].Params["cert"].ValidateFunc
		errs := validator("--cert", p.Cert)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["WebAccel"].Commands["certificate-update"].Params["key"].ValidateFunc
		errs := validator("--key", p.Key)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		errs := validation.ConflictsWith("--cert-content", p.CertContent, map[string]interface{}{

			"--cert": p.Cert,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		errs := validation.ConflictsWith("--key-content", p.KeyContent, map[string]interface{}{

			"--key": p.Key,
		})
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

func (p *CertificateUpdateWebAccelParam) ResourceDef() *schema.Resource {
	return define.Resources["WebAccel"]
}

func (p *CertificateUpdateWebAccelParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["certificate-update"]
}

func (p *CertificateUpdateWebAccelParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *CertificateUpdateWebAccelParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *CertificateUpdateWebAccelParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *CertificateUpdateWebAccelParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *CertificateUpdateWebAccelParam) SetCert(v string) {
	p.Cert = v
}

func (p *CertificateUpdateWebAccelParam) GetCert() string {
	return p.Cert
}
func (p *CertificateUpdateWebAccelParam) SetKey(v string) {
	p.Key = v
}

func (p *CertificateUpdateWebAccelParam) GetKey() string {
	return p.Key
}
func (p *CertificateUpdateWebAccelParam) SetCertContent(v string) {
	p.CertContent = v
}

func (p *CertificateUpdateWebAccelParam) GetCertContent() string {
	return p.CertContent
}
func (p *CertificateUpdateWebAccelParam) SetKeyContent(v string) {
	p.KeyContent = v
}

func (p *CertificateUpdateWebAccelParam) GetKeyContent() string {
	return p.KeyContent
}
func (p *CertificateUpdateWebAccelParam) SetSelector(v []string) {
	p.Selector = v
}

func (p *CertificateUpdateWebAccelParam) GetSelector() []string {
	return p.Selector
}
func (p *CertificateUpdateWebAccelParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *CertificateUpdateWebAccelParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *CertificateUpdateWebAccelParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *CertificateUpdateWebAccelParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *CertificateUpdateWebAccelParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *CertificateUpdateWebAccelParam) GetParameters() string {
	return p.Parameters
}
func (p *CertificateUpdateWebAccelParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *CertificateUpdateWebAccelParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *CertificateUpdateWebAccelParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *CertificateUpdateWebAccelParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *CertificateUpdateWebAccelParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *CertificateUpdateWebAccelParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *CertificateUpdateWebAccelParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *CertificateUpdateWebAccelParam) GetOutputType() string {
	return p.OutputType
}
func (p *CertificateUpdateWebAccelParam) SetColumn(v []string) {
	p.Column = v
}

func (p *CertificateUpdateWebAccelParam) GetColumn() []string {
	return p.Column
}
func (p *CertificateUpdateWebAccelParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *CertificateUpdateWebAccelParam) GetQuiet() bool {
	return p.Quiet
}
func (p *CertificateUpdateWebAccelParam) SetFormat(v string) {
	p.Format = v
}

func (p *CertificateUpdateWebAccelParam) GetFormat() string {
	return p.Format
}
func (p *CertificateUpdateWebAccelParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *CertificateUpdateWebAccelParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *CertificateUpdateWebAccelParam) SetQuery(v string) {
	p.Query = v
}

func (p *CertificateUpdateWebAccelParam) GetQuery() string {
	return p.Query
}
func (p *CertificateUpdateWebAccelParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *CertificateUpdateWebAccelParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *CertificateUpdateWebAccelParam) SetId(v sacloud.ID) {
	p.Id = v
}

func (p *CertificateUpdateWebAccelParam) GetId() sacloud.ID {
	return p.Id
}

// DeleteCacheWebAccelParam is input parameters for the sacloud API
type DeleteCacheWebAccelParam struct {
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

// NewDeleteCacheWebAccelParam return new DeleteCacheWebAccelParam
func NewDeleteCacheWebAccelParam() *DeleteCacheWebAccelParam {
	return &DeleteCacheWebAccelParam{}
}

// Initialize init DeleteCacheWebAccelParam
func (p *DeleteCacheWebAccelParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *DeleteCacheWebAccelParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *DeleteCacheWebAccelParam) fillValueToSkeleton() {
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

func (p *DeleteCacheWebAccelParam) validate() error {
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

func (p *DeleteCacheWebAccelParam) ResourceDef() *schema.Resource {
	return define.Resources["WebAccel"]
}

func (p *DeleteCacheWebAccelParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["delete-cache"]
}

func (p *DeleteCacheWebAccelParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *DeleteCacheWebAccelParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *DeleteCacheWebAccelParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *DeleteCacheWebAccelParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *DeleteCacheWebAccelParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *DeleteCacheWebAccelParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *DeleteCacheWebAccelParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *DeleteCacheWebAccelParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *DeleteCacheWebAccelParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *DeleteCacheWebAccelParam) GetParameters() string {
	return p.Parameters
}
func (p *DeleteCacheWebAccelParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *DeleteCacheWebAccelParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *DeleteCacheWebAccelParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *DeleteCacheWebAccelParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *DeleteCacheWebAccelParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *DeleteCacheWebAccelParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *DeleteCacheWebAccelParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *DeleteCacheWebAccelParam) GetOutputType() string {
	return p.OutputType
}
func (p *DeleteCacheWebAccelParam) SetColumn(v []string) {
	p.Column = v
}

func (p *DeleteCacheWebAccelParam) GetColumn() []string {
	return p.Column
}
func (p *DeleteCacheWebAccelParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *DeleteCacheWebAccelParam) GetQuiet() bool {
	return p.Quiet
}
func (p *DeleteCacheWebAccelParam) SetFormat(v string) {
	p.Format = v
}

func (p *DeleteCacheWebAccelParam) GetFormat() string {
	return p.Format
}
func (p *DeleteCacheWebAccelParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *DeleteCacheWebAccelParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *DeleteCacheWebAccelParam) SetQuery(v string) {
	p.Query = v
}

func (p *DeleteCacheWebAccelParam) GetQuery() string {
	return p.Query
}
func (p *DeleteCacheWebAccelParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *DeleteCacheWebAccelParam) GetQueryFile() string {
	return p.QueryFile
}
