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

// ListDNSParam is input parameters for the sacloud API
type ListDNSParam struct {
	From int
	Max  int
	Tags []string
	Sort []string
	Name []string
	Id   []sacloud.ID

	input Input
}

// NewListDNSParam return new ListDNSParam
func NewListDNSParam() *ListDNSParam {
	return &ListDNSParam{}
}

// Initialize init ListDNSParam
func (p *ListDNSParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ListDNSParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *ListDNSParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.From) {
		p.From = 0
	}
	if utils.IsEmpty(p.Max) {
		p.Max = 0
	}
	if utils.IsEmpty(p.Tags) {
		p.Tags = []string{""}
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

func (p *ListDNSParam) validate() error {
	var errors []error

	{
		validator := define.Resources["DNS"].Commands["list"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		errs := validation.ConflictsWith("--name", p.Name, map[string]interface{}{

			"--id": p.Id,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["DNS"].Commands["list"].Params["id"].ValidateFunc
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

func (p *ListDNSParam) ResourceDef() *schema.Resource {
	return define.Resources["DNS"]
}

func (p *ListDNSParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["list"]
}

func (p *ListDNSParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ListDNSParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ListDNSParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ListDNSParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ListDNSParam) SetFrom(v int) {
	p.From = v
}

func (p *ListDNSParam) GetFrom() int {
	return p.From
}
func (p *ListDNSParam) SetMax(v int) {
	p.Max = v
}

func (p *ListDNSParam) GetMax() int {
	return p.Max
}
func (p *ListDNSParam) SetTags(v []string) {
	p.Tags = v
}

func (p *ListDNSParam) GetTags() []string {
	return p.Tags
}
func (p *ListDNSParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListDNSParam) GetSort() []string {
	return p.Sort
}
func (p *ListDNSParam) SetName(v []string) {
	p.Name = v
}

func (p *ListDNSParam) GetName() []string {
	return p.Name
}
func (p *ListDNSParam) SetId(v []sacloud.ID) {
	p.Id = v
}

func (p *ListDNSParam) GetId() []sacloud.ID {
	return p.Id
}

// RecordInfoDNSParam is input parameters for the sacloud API
type RecordInfoDNSParam struct {
	Name string
	Type string

	input Input
}

// NewRecordInfoDNSParam return new RecordInfoDNSParam
func NewRecordInfoDNSParam() *RecordInfoDNSParam {
	return &RecordInfoDNSParam{}
}

// Initialize init RecordInfoDNSParam
func (p *RecordInfoDNSParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *RecordInfoDNSParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *RecordInfoDNSParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.Name) {
		p.Name = ""
	}
	if utils.IsEmpty(p.Type) {
		p.Type = ""
	}

}

func (p *RecordInfoDNSParam) validate() error {
	var errors []error

	{
		validator := define.Resources["DNS"].Commands["record-info"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["DNS"].Commands["record-info"].Params["type"].ValidateFunc
		errs := validator("--type", p.Type)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return utils.FlattenErrors(errors)
}

func (p *RecordInfoDNSParam) ResourceDef() *schema.Resource {
	return define.Resources["DNS"]
}

func (p *RecordInfoDNSParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["record-info"]
}

func (p *RecordInfoDNSParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *RecordInfoDNSParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *RecordInfoDNSParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *RecordInfoDNSParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *RecordInfoDNSParam) SetName(v string) {
	p.Name = v
}

func (p *RecordInfoDNSParam) GetName() string {
	return p.Name
}
func (p *RecordInfoDNSParam) SetType(v string) {
	p.Type = v
}

func (p *RecordInfoDNSParam) GetType() string {
	return p.Type
}

// RecordBulkUpdateDNSParam is input parameters for the sacloud API
type RecordBulkUpdateDNSParam struct {
	File string
	Mode string

	input Input
}

// NewRecordBulkUpdateDNSParam return new RecordBulkUpdateDNSParam
func NewRecordBulkUpdateDNSParam() *RecordBulkUpdateDNSParam {
	return &RecordBulkUpdateDNSParam{
		Mode: "upsert-only"}
}

// Initialize init RecordBulkUpdateDNSParam
func (p *RecordBulkUpdateDNSParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *RecordBulkUpdateDNSParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *RecordBulkUpdateDNSParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.File) {
		p.File = ""
	}
	if utils.IsEmpty(p.Mode) {
		p.Mode = ""
	}

}

func (p *RecordBulkUpdateDNSParam) validate() error {
	var errors []error

	{
		validator := validateRequired
		errs := validator("--file", p.File)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["DNS"].Commands["record-bulk-update"].Params["file"].ValidateFunc
		errs := validator("--file", p.File)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := validateRequired
		errs := validator("--mode", p.Mode)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["DNS"].Commands["record-bulk-update"].Params["mode"].ValidateFunc
		errs := validator("--mode", p.Mode)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return utils.FlattenErrors(errors)
}

func (p *RecordBulkUpdateDNSParam) ResourceDef() *schema.Resource {
	return define.Resources["DNS"]
}

func (p *RecordBulkUpdateDNSParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["record-bulk-update"]
}

func (p *RecordBulkUpdateDNSParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *RecordBulkUpdateDNSParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *RecordBulkUpdateDNSParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *RecordBulkUpdateDNSParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *RecordBulkUpdateDNSParam) SetFile(v string) {
	p.File = v
}

func (p *RecordBulkUpdateDNSParam) GetFile() string {
	return p.File
}
func (p *RecordBulkUpdateDNSParam) SetMode(v string) {
	p.Mode = v
}

func (p *RecordBulkUpdateDNSParam) GetMode() string {
	return p.Mode
}

// CreateDNSParam is input parameters for the sacloud API
type CreateDNSParam struct {
	Description string
	Tags        []string
	IconId      sacloud.ID
	Name        string

	input Input
}

// NewCreateDNSParam return new CreateDNSParam
func NewCreateDNSParam() *CreateDNSParam {
	return &CreateDNSParam{}
}

// Initialize init CreateDNSParam
func (p *CreateDNSParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *CreateDNSParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *CreateDNSParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.Description) {
		p.Description = ""
	}
	if utils.IsEmpty(p.Tags) {
		p.Tags = []string{""}
	}
	if utils.IsEmpty(p.IconId) {
		p.IconId = sacloud.ID(0)
	}
	if utils.IsEmpty(p.Name) {
		p.Name = ""
	}

}

func (p *CreateDNSParam) validate() error {
	var errors []error

	{
		validator := define.Resources["DNS"].Commands["create"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["DNS"].Commands["create"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["DNS"].Commands["create"].Params["icon-id"].ValidateFunc
		errs := validator("--icon-id", p.IconId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := validateRequired
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["DNS"].Commands["create"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return utils.FlattenErrors(errors)
}

func (p *CreateDNSParam) ResourceDef() *schema.Resource {
	return define.Resources["DNS"]
}

func (p *CreateDNSParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["create"]
}

func (p *CreateDNSParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *CreateDNSParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *CreateDNSParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *CreateDNSParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *CreateDNSParam) SetDescription(v string) {
	p.Description = v
}

func (p *CreateDNSParam) GetDescription() string {
	return p.Description
}
func (p *CreateDNSParam) SetTags(v []string) {
	p.Tags = v
}

func (p *CreateDNSParam) GetTags() []string {
	return p.Tags
}
func (p *CreateDNSParam) SetIconId(v sacloud.ID) {
	p.IconId = v
}

func (p *CreateDNSParam) GetIconId() sacloud.ID {
	return p.IconId
}
func (p *CreateDNSParam) SetName(v string) {
	p.Name = v
}

func (p *CreateDNSParam) GetName() string {
	return p.Name
}

// RecordAddDNSParam is input parameters for the sacloud API
type RecordAddDNSParam struct {
	SrvPort     int
	Name        string
	Ttl         int
	MxPriority  int
	SrvPriority int
	SrvWeight   int
	Type        string
	Value       string
	SrvTarget   string

	input Input
}

// NewRecordAddDNSParam return new RecordAddDNSParam
func NewRecordAddDNSParam() *RecordAddDNSParam {
	return &RecordAddDNSParam{
		Ttl: 3600, MxPriority: 10}
}

// Initialize init RecordAddDNSParam
func (p *RecordAddDNSParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *RecordAddDNSParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *RecordAddDNSParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.SrvPort) {
		p.SrvPort = 0
	}
	if utils.IsEmpty(p.Name) {
		p.Name = ""
	}
	if utils.IsEmpty(p.Ttl) {
		p.Ttl = 0
	}
	if utils.IsEmpty(p.MxPriority) {
		p.MxPriority = 0
	}
	if utils.IsEmpty(p.SrvPriority) {
		p.SrvPriority = 0
	}
	if utils.IsEmpty(p.SrvWeight) {
		p.SrvWeight = 0
	}
	if utils.IsEmpty(p.Type) {
		p.Type = ""
	}
	if utils.IsEmpty(p.Value) {
		p.Value = ""
	}
	if utils.IsEmpty(p.SrvTarget) {
		p.SrvTarget = ""
	}

}

func (p *RecordAddDNSParam) validate() error {
	var errors []error

	{
		validator := define.Resources["DNS"].Commands["record-add"].Params["srv-port"].ValidateFunc
		errs := validator("--srv-port", p.SrvPort)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := validateRequired
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["DNS"].Commands["record-add"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["DNS"].Commands["record-add"].Params["ttl"].ValidateFunc
		errs := validator("--ttl", p.Ttl)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["DNS"].Commands["record-add"].Params["mx-priority"].ValidateFunc
		errs := validator("--mx-priority", p.MxPriority)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["DNS"].Commands["record-add"].Params["srv-priority"].ValidateFunc
		errs := validator("--srv-priority", p.SrvPriority)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["DNS"].Commands["record-add"].Params["srv-weight"].ValidateFunc
		errs := validator("--srv-weight", p.SrvWeight)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := validateRequired
		errs := validator("--type", p.Type)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["DNS"].Commands["record-add"].Params["type"].ValidateFunc
		errs := validator("--type", p.Type)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["DNS"].Commands["record-add"].Params["srv-target"].ValidateFunc
		errs := validator("--srv-target", p.SrvTarget)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return utils.FlattenErrors(errors)
}

func (p *RecordAddDNSParam) ResourceDef() *schema.Resource {
	return define.Resources["DNS"]
}

func (p *RecordAddDNSParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["record-add"]
}

func (p *RecordAddDNSParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *RecordAddDNSParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *RecordAddDNSParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *RecordAddDNSParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *RecordAddDNSParam) SetSrvPort(v int) {
	p.SrvPort = v
}

func (p *RecordAddDNSParam) GetSrvPort() int {
	return p.SrvPort
}
func (p *RecordAddDNSParam) SetName(v string) {
	p.Name = v
}

func (p *RecordAddDNSParam) GetName() string {
	return p.Name
}
func (p *RecordAddDNSParam) SetTtl(v int) {
	p.Ttl = v
}

func (p *RecordAddDNSParam) GetTtl() int {
	return p.Ttl
}
func (p *RecordAddDNSParam) SetMxPriority(v int) {
	p.MxPriority = v
}

func (p *RecordAddDNSParam) GetMxPriority() int {
	return p.MxPriority
}
func (p *RecordAddDNSParam) SetSrvPriority(v int) {
	p.SrvPriority = v
}

func (p *RecordAddDNSParam) GetSrvPriority() int {
	return p.SrvPriority
}
func (p *RecordAddDNSParam) SetSrvWeight(v int) {
	p.SrvWeight = v
}

func (p *RecordAddDNSParam) GetSrvWeight() int {
	return p.SrvWeight
}
func (p *RecordAddDNSParam) SetType(v string) {
	p.Type = v
}

func (p *RecordAddDNSParam) GetType() string {
	return p.Type
}
func (p *RecordAddDNSParam) SetValue(v string) {
	p.Value = v
}

func (p *RecordAddDNSParam) GetValue() string {
	return p.Value
}
func (p *RecordAddDNSParam) SetSrvTarget(v string) {
	p.SrvTarget = v
}

func (p *RecordAddDNSParam) GetSrvTarget() string {
	return p.SrvTarget
}

// ReadDNSParam is input parameters for the sacloud API
type ReadDNSParam struct {
	input Input
}

// NewReadDNSParam return new ReadDNSParam
func NewReadDNSParam() *ReadDNSParam {
	return &ReadDNSParam{}
}

// Initialize init ReadDNSParam
func (p *ReadDNSParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ReadDNSParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *ReadDNSParam) fillValueToSkeleton() {

}

func (p *ReadDNSParam) validate() error {
	var errors []error

	return utils.FlattenErrors(errors)
}

func (p *ReadDNSParam) ResourceDef() *schema.Resource {
	return define.Resources["DNS"]
}

func (p *ReadDNSParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["read"]
}

func (p *ReadDNSParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ReadDNSParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ReadDNSParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ReadDNSParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

// RecordUpdateDNSParam is input parameters for the sacloud API
type RecordUpdateDNSParam struct {
	Index       int
	Name        string
	MxPriority  int
	SrvPriority int
	SrvWeight   int
	SrvTarget   string
	Type        string
	Value       string
	Ttl         int
	SrvPort     int

	input Input
}

// NewRecordUpdateDNSParam return new RecordUpdateDNSParam
func NewRecordUpdateDNSParam() *RecordUpdateDNSParam {
	return &RecordUpdateDNSParam{}
}

// Initialize init RecordUpdateDNSParam
func (p *RecordUpdateDNSParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *RecordUpdateDNSParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *RecordUpdateDNSParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.Index) {
		p.Index = 0
	}
	if utils.IsEmpty(p.Name) {
		p.Name = ""
	}
	if utils.IsEmpty(p.MxPriority) {
		p.MxPriority = 0
	}
	if utils.IsEmpty(p.SrvPriority) {
		p.SrvPriority = 0
	}
	if utils.IsEmpty(p.SrvWeight) {
		p.SrvWeight = 0
	}
	if utils.IsEmpty(p.SrvTarget) {
		p.SrvTarget = ""
	}
	if utils.IsEmpty(p.Type) {
		p.Type = ""
	}
	if utils.IsEmpty(p.Value) {
		p.Value = ""
	}
	if utils.IsEmpty(p.Ttl) {
		p.Ttl = 0
	}
	if utils.IsEmpty(p.SrvPort) {
		p.SrvPort = 0
	}

}

func (p *RecordUpdateDNSParam) validate() error {
	var errors []error

	{
		validator := validateRequired
		errs := validator("--index", p.Index)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["DNS"].Commands["record-update"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["DNS"].Commands["record-update"].Params["mx-priority"].ValidateFunc
		errs := validator("--mx-priority", p.MxPriority)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["DNS"].Commands["record-update"].Params["srv-priority"].ValidateFunc
		errs := validator("--srv-priority", p.SrvPriority)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["DNS"].Commands["record-update"].Params["srv-weight"].ValidateFunc
		errs := validator("--srv-weight", p.SrvWeight)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["DNS"].Commands["record-update"].Params["srv-target"].ValidateFunc
		errs := validator("--srv-target", p.SrvTarget)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["DNS"].Commands["record-update"].Params["type"].ValidateFunc
		errs := validator("--type", p.Type)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["DNS"].Commands["record-update"].Params["ttl"].ValidateFunc
		errs := validator("--ttl", p.Ttl)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["DNS"].Commands["record-update"].Params["srv-port"].ValidateFunc
		errs := validator("--srv-port", p.SrvPort)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return utils.FlattenErrors(errors)
}

func (p *RecordUpdateDNSParam) ResourceDef() *schema.Resource {
	return define.Resources["DNS"]
}

func (p *RecordUpdateDNSParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["record-update"]
}

func (p *RecordUpdateDNSParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *RecordUpdateDNSParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *RecordUpdateDNSParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *RecordUpdateDNSParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *RecordUpdateDNSParam) SetIndex(v int) {
	p.Index = v
}

func (p *RecordUpdateDNSParam) GetIndex() int {
	return p.Index
}
func (p *RecordUpdateDNSParam) SetName(v string) {
	p.Name = v
}

func (p *RecordUpdateDNSParam) GetName() string {
	return p.Name
}
func (p *RecordUpdateDNSParam) SetMxPriority(v int) {
	p.MxPriority = v
}

func (p *RecordUpdateDNSParam) GetMxPriority() int {
	return p.MxPriority
}
func (p *RecordUpdateDNSParam) SetSrvPriority(v int) {
	p.SrvPriority = v
}

func (p *RecordUpdateDNSParam) GetSrvPriority() int {
	return p.SrvPriority
}
func (p *RecordUpdateDNSParam) SetSrvWeight(v int) {
	p.SrvWeight = v
}

func (p *RecordUpdateDNSParam) GetSrvWeight() int {
	return p.SrvWeight
}
func (p *RecordUpdateDNSParam) SetSrvTarget(v string) {
	p.SrvTarget = v
}

func (p *RecordUpdateDNSParam) GetSrvTarget() string {
	return p.SrvTarget
}
func (p *RecordUpdateDNSParam) SetType(v string) {
	p.Type = v
}

func (p *RecordUpdateDNSParam) GetType() string {
	return p.Type
}
func (p *RecordUpdateDNSParam) SetValue(v string) {
	p.Value = v
}

func (p *RecordUpdateDNSParam) GetValue() string {
	return p.Value
}
func (p *RecordUpdateDNSParam) SetTtl(v int) {
	p.Ttl = v
}

func (p *RecordUpdateDNSParam) GetTtl() int {
	return p.Ttl
}
func (p *RecordUpdateDNSParam) SetSrvPort(v int) {
	p.SrvPort = v
}

func (p *RecordUpdateDNSParam) GetSrvPort() int {
	return p.SrvPort
}

// RecordDeleteDNSParam is input parameters for the sacloud API
type RecordDeleteDNSParam struct {
	Index int

	input Input
}

// NewRecordDeleteDNSParam return new RecordDeleteDNSParam
func NewRecordDeleteDNSParam() *RecordDeleteDNSParam {
	return &RecordDeleteDNSParam{}
}

// Initialize init RecordDeleteDNSParam
func (p *RecordDeleteDNSParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *RecordDeleteDNSParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *RecordDeleteDNSParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.Index) {
		p.Index = 0
	}

}

func (p *RecordDeleteDNSParam) validate() error {
	var errors []error

	{
		validator := validateRequired
		errs := validator("--index", p.Index)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return utils.FlattenErrors(errors)
}

func (p *RecordDeleteDNSParam) ResourceDef() *schema.Resource {
	return define.Resources["DNS"]
}

func (p *RecordDeleteDNSParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["record-delete"]
}

func (p *RecordDeleteDNSParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *RecordDeleteDNSParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *RecordDeleteDNSParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *RecordDeleteDNSParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *RecordDeleteDNSParam) SetIndex(v int) {
	p.Index = v
}

func (p *RecordDeleteDNSParam) GetIndex() int {
	return p.Index
}

// UpdateDNSParam is input parameters for the sacloud API
type UpdateDNSParam struct {
	Description string
	Tags        []string
	IconId      sacloud.ID

	input Input
}

// NewUpdateDNSParam return new UpdateDNSParam
func NewUpdateDNSParam() *UpdateDNSParam {
	return &UpdateDNSParam{}
}

// Initialize init UpdateDNSParam
func (p *UpdateDNSParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *UpdateDNSParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *UpdateDNSParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.Description) {
		p.Description = ""
	}
	if utils.IsEmpty(p.Tags) {
		p.Tags = []string{""}
	}
	if utils.IsEmpty(p.IconId) {
		p.IconId = sacloud.ID(0)
	}

}

func (p *UpdateDNSParam) validate() error {
	var errors []error

	{
		validator := define.Resources["DNS"].Commands["update"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["DNS"].Commands["update"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["DNS"].Commands["update"].Params["icon-id"].ValidateFunc
		errs := validator("--icon-id", p.IconId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return utils.FlattenErrors(errors)
}

func (p *UpdateDNSParam) ResourceDef() *schema.Resource {
	return define.Resources["DNS"]
}

func (p *UpdateDNSParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["update"]
}

func (p *UpdateDNSParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *UpdateDNSParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *UpdateDNSParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *UpdateDNSParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *UpdateDNSParam) SetDescription(v string) {
	p.Description = v
}

func (p *UpdateDNSParam) GetDescription() string {
	return p.Description
}
func (p *UpdateDNSParam) SetTags(v []string) {
	p.Tags = v
}

func (p *UpdateDNSParam) GetTags() []string {
	return p.Tags
}
func (p *UpdateDNSParam) SetIconId(v sacloud.ID) {
	p.IconId = v
}

func (p *UpdateDNSParam) GetIconId() sacloud.ID {
	return p.IconId
}

// DeleteDNSParam is input parameters for the sacloud API
type DeleteDNSParam struct {
	input Input
}

// NewDeleteDNSParam return new DeleteDNSParam
func NewDeleteDNSParam() *DeleteDNSParam {
	return &DeleteDNSParam{}
}

// Initialize init DeleteDNSParam
func (p *DeleteDNSParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *DeleteDNSParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *DeleteDNSParam) fillValueToSkeleton() {

}

func (p *DeleteDNSParam) validate() error {
	var errors []error

	return utils.FlattenErrors(errors)
}

func (p *DeleteDNSParam) ResourceDef() *schema.Resource {
	return define.Resources["DNS"]
}

func (p *DeleteDNSParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["delete"]
}

func (p *DeleteDNSParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *DeleteDNSParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *DeleteDNSParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *DeleteDNSParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}
