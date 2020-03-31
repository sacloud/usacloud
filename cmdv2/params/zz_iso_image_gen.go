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

// ListISOImageParam is input parameters for the sacloud API
type ListISOImageParam struct {
	Scope string
	Tags  []string
	Name  []string
	Id    []sacloud.ID
	From  int
	Max   int
	Sort  []string

	input Input
}

// NewListISOImageParam return new ListISOImageParam
func NewListISOImageParam() *ListISOImageParam {
	return &ListISOImageParam{}
}

// Initialize init ListISOImageParam
func (p *ListISOImageParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ListISOImageParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *ListISOImageParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.Scope) {
		p.Scope = ""
	}
	if utils.IsEmpty(p.Tags) {
		p.Tags = []string{""}
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
	if utils.IsEmpty(p.Sort) {
		p.Sort = []string{""}
	}

}

func (p *ListISOImageParam) validate() error {
	var errors []error

	{
		validator := define.Resources["ISOImage"].Commands["list"].Params["scope"].ValidateFunc
		errs := validator("--scope", p.Scope)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["ISOImage"].Commands["list"].Params["tags"].ValidateFunc
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
		validator := define.Resources["ISOImage"].Commands["list"].Params["id"].ValidateFunc
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

func (p *ListISOImageParam) ResourceDef() *schema.Resource {
	return define.Resources["ISOImage"]
}

func (p *ListISOImageParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["list"]
}

func (p *ListISOImageParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ListISOImageParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ListISOImageParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ListISOImageParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ListISOImageParam) SetScope(v string) {
	p.Scope = v
}

func (p *ListISOImageParam) GetScope() string {
	return p.Scope
}
func (p *ListISOImageParam) SetTags(v []string) {
	p.Tags = v
}

func (p *ListISOImageParam) GetTags() []string {
	return p.Tags
}
func (p *ListISOImageParam) SetName(v []string) {
	p.Name = v
}

func (p *ListISOImageParam) GetName() []string {
	return p.Name
}
func (p *ListISOImageParam) SetId(v []sacloud.ID) {
	p.Id = v
}

func (p *ListISOImageParam) GetId() []sacloud.ID {
	return p.Id
}
func (p *ListISOImageParam) SetFrom(v int) {
	p.From = v
}

func (p *ListISOImageParam) GetFrom() int {
	return p.From
}
func (p *ListISOImageParam) SetMax(v int) {
	p.Max = v
}

func (p *ListISOImageParam) GetMax() int {
	return p.Max
}
func (p *ListISOImageParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListISOImageParam) GetSort() []string {
	return p.Sort
}

// CreateISOImageParam is input parameters for the sacloud API
type CreateISOImageParam struct {
	ISOFile     string
	Name        string
	Description string
	Tags        []string
	IconId      sacloud.ID
	Size        int

	input Input
}

// NewCreateISOImageParam return new CreateISOImageParam
func NewCreateISOImageParam() *CreateISOImageParam {
	return &CreateISOImageParam{
		Size: 5}
}

// Initialize init CreateISOImageParam
func (p *CreateISOImageParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *CreateISOImageParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *CreateISOImageParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.ISOFile) {
		p.ISOFile = ""
	}
	if utils.IsEmpty(p.Name) {
		p.Name = ""
	}
	if utils.IsEmpty(p.Description) {
		p.Description = ""
	}
	if utils.IsEmpty(p.Tags) {
		p.Tags = []string{""}
	}
	if utils.IsEmpty(p.IconId) {
		p.IconId = sacloud.ID(0)
	}
	if utils.IsEmpty(p.Size) {
		p.Size = 0
	}

}

func (p *CreateISOImageParam) validate() error {
	var errors []error

	{
		validator := define.Resources["ISOImage"].Commands["create"].Params["iso-file"].ValidateFunc
		errs := validator("--iso-file", p.ISOFile)
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
		validator := define.Resources["ISOImage"].Commands["create"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["ISOImage"].Commands["create"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["ISOImage"].Commands["create"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["ISOImage"].Commands["create"].Params["icon-id"].ValidateFunc
		errs := validator("--icon-id", p.IconId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := validateRequired
		errs := validator("--size", p.Size)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ISOImage"].Commands["create"].Params["size"].ValidateFunc
		errs := validator("--size", p.Size)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return utils.FlattenErrors(errors)
}

func (p *CreateISOImageParam) ResourceDef() *schema.Resource {
	return define.Resources["ISOImage"]
}

func (p *CreateISOImageParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["create"]
}

func (p *CreateISOImageParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *CreateISOImageParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *CreateISOImageParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *CreateISOImageParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *CreateISOImageParam) SetISOFile(v string) {
	p.ISOFile = v
}

func (p *CreateISOImageParam) GetISOFile() string {
	return p.ISOFile
}
func (p *CreateISOImageParam) SetName(v string) {
	p.Name = v
}

func (p *CreateISOImageParam) GetName() string {
	return p.Name
}
func (p *CreateISOImageParam) SetDescription(v string) {
	p.Description = v
}

func (p *CreateISOImageParam) GetDescription() string {
	return p.Description
}
func (p *CreateISOImageParam) SetTags(v []string) {
	p.Tags = v
}

func (p *CreateISOImageParam) GetTags() []string {
	return p.Tags
}
func (p *CreateISOImageParam) SetIconId(v sacloud.ID) {
	p.IconId = v
}

func (p *CreateISOImageParam) GetIconId() sacloud.ID {
	return p.IconId
}
func (p *CreateISOImageParam) SetSize(v int) {
	p.Size = v
}

func (p *CreateISOImageParam) GetSize() int {
	return p.Size
}

// ReadISOImageParam is input parameters for the sacloud API
type ReadISOImageParam struct {
	input Input
}

// NewReadISOImageParam return new ReadISOImageParam
func NewReadISOImageParam() *ReadISOImageParam {
	return &ReadISOImageParam{}
}

// Initialize init ReadISOImageParam
func (p *ReadISOImageParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ReadISOImageParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *ReadISOImageParam) fillValueToSkeleton() {

}

func (p *ReadISOImageParam) validate() error {
	var errors []error

	return utils.FlattenErrors(errors)
}

func (p *ReadISOImageParam) ResourceDef() *schema.Resource {
	return define.Resources["ISOImage"]
}

func (p *ReadISOImageParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["read"]
}

func (p *ReadISOImageParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ReadISOImageParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ReadISOImageParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ReadISOImageParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

// UpdateISOImageParam is input parameters for the sacloud API
type UpdateISOImageParam struct {
	Name        string
	Description string
	Tags        []string
	IconId      sacloud.ID

	input Input
}

// NewUpdateISOImageParam return new UpdateISOImageParam
func NewUpdateISOImageParam() *UpdateISOImageParam {
	return &UpdateISOImageParam{}
}

// Initialize init UpdateISOImageParam
func (p *UpdateISOImageParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *UpdateISOImageParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *UpdateISOImageParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.Name) {
		p.Name = ""
	}
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

func (p *UpdateISOImageParam) validate() error {
	var errors []error

	{
		validator := define.Resources["ISOImage"].Commands["update"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["ISOImage"].Commands["update"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["ISOImage"].Commands["update"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["ISOImage"].Commands["update"].Params["icon-id"].ValidateFunc
		errs := validator("--icon-id", p.IconId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return utils.FlattenErrors(errors)
}

func (p *UpdateISOImageParam) ResourceDef() *schema.Resource {
	return define.Resources["ISOImage"]
}

func (p *UpdateISOImageParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["update"]
}

func (p *UpdateISOImageParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *UpdateISOImageParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *UpdateISOImageParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *UpdateISOImageParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *UpdateISOImageParam) SetName(v string) {
	p.Name = v
}

func (p *UpdateISOImageParam) GetName() string {
	return p.Name
}
func (p *UpdateISOImageParam) SetDescription(v string) {
	p.Description = v
}

func (p *UpdateISOImageParam) GetDescription() string {
	return p.Description
}
func (p *UpdateISOImageParam) SetTags(v []string) {
	p.Tags = v
}

func (p *UpdateISOImageParam) GetTags() []string {
	return p.Tags
}
func (p *UpdateISOImageParam) SetIconId(v sacloud.ID) {
	p.IconId = v
}

func (p *UpdateISOImageParam) GetIconId() sacloud.ID {
	return p.IconId
}

// DeleteISOImageParam is input parameters for the sacloud API
type DeleteISOImageParam struct {
	input Input
}

// NewDeleteISOImageParam return new DeleteISOImageParam
func NewDeleteISOImageParam() *DeleteISOImageParam {
	return &DeleteISOImageParam{}
}

// Initialize init DeleteISOImageParam
func (p *DeleteISOImageParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *DeleteISOImageParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *DeleteISOImageParam) fillValueToSkeleton() {

}

func (p *DeleteISOImageParam) validate() error {
	var errors []error

	return utils.FlattenErrors(errors)
}

func (p *DeleteISOImageParam) ResourceDef() *schema.Resource {
	return define.Resources["ISOImage"]
}

func (p *DeleteISOImageParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["delete"]
}

func (p *DeleteISOImageParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *DeleteISOImageParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *DeleteISOImageParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *DeleteISOImageParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

// UploadISOImageParam is input parameters for the sacloud API
type UploadISOImageParam struct {
	ISOFile string

	input Input
}

// NewUploadISOImageParam return new UploadISOImageParam
func NewUploadISOImageParam() *UploadISOImageParam {
	return &UploadISOImageParam{}
}

// Initialize init UploadISOImageParam
func (p *UploadISOImageParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *UploadISOImageParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *UploadISOImageParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.ISOFile) {
		p.ISOFile = ""
	}

}

func (p *UploadISOImageParam) validate() error {
	var errors []error

	{
		validator := define.Resources["ISOImage"].Commands["upload"].Params["iso-file"].ValidateFunc
		errs := validator("--iso-file", p.ISOFile)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return utils.FlattenErrors(errors)
}

func (p *UploadISOImageParam) ResourceDef() *schema.Resource {
	return define.Resources["ISOImage"]
}

func (p *UploadISOImageParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["upload"]
}

func (p *UploadISOImageParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *UploadISOImageParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *UploadISOImageParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *UploadISOImageParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *UploadISOImageParam) SetISOFile(v string) {
	p.ISOFile = v
}

func (p *UploadISOImageParam) GetISOFile() string {
	return p.ISOFile
}

// DownloadISOImageParam is input parameters for the sacloud API
type DownloadISOImageParam struct {
	FileDestination string

	input Input
}

// NewDownloadISOImageParam return new DownloadISOImageParam
func NewDownloadISOImageParam() *DownloadISOImageParam {
	return &DownloadISOImageParam{}
}

// Initialize init DownloadISOImageParam
func (p *DownloadISOImageParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *DownloadISOImageParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *DownloadISOImageParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.FileDestination) {
		p.FileDestination = ""
	}

}

func (p *DownloadISOImageParam) validate() error {
	var errors []error

	return utils.FlattenErrors(errors)
}

func (p *DownloadISOImageParam) ResourceDef() *schema.Resource {
	return define.Resources["ISOImage"]
}

func (p *DownloadISOImageParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["download"]
}

func (p *DownloadISOImageParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *DownloadISOImageParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *DownloadISOImageParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *DownloadISOImageParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *DownloadISOImageParam) SetFileDestination(v string) {
	p.FileDestination = v
}

func (p *DownloadISOImageParam) GetFileDestination() string {
	return p.FileDestination
}

// FTPOpenISOImageParam is input parameters for the sacloud API
type FTPOpenISOImageParam struct {
	input Input
}

// NewFTPOpenISOImageParam return new FTPOpenISOImageParam
func NewFTPOpenISOImageParam() *FTPOpenISOImageParam {
	return &FTPOpenISOImageParam{}
}

// Initialize init FTPOpenISOImageParam
func (p *FTPOpenISOImageParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *FTPOpenISOImageParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *FTPOpenISOImageParam) fillValueToSkeleton() {

}

func (p *FTPOpenISOImageParam) validate() error {
	var errors []error

	return utils.FlattenErrors(errors)
}

func (p *FTPOpenISOImageParam) ResourceDef() *schema.Resource {
	return define.Resources["ISOImage"]
}

func (p *FTPOpenISOImageParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["ftp-open"]
}

func (p *FTPOpenISOImageParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *FTPOpenISOImageParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *FTPOpenISOImageParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *FTPOpenISOImageParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

// FTPCloseISOImageParam is input parameters for the sacloud API
type FTPCloseISOImageParam struct {
	input Input
}

// NewFTPCloseISOImageParam return new FTPCloseISOImageParam
func NewFTPCloseISOImageParam() *FTPCloseISOImageParam {
	return &FTPCloseISOImageParam{}
}

// Initialize init FTPCloseISOImageParam
func (p *FTPCloseISOImageParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *FTPCloseISOImageParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *FTPCloseISOImageParam) fillValueToSkeleton() {

}

func (p *FTPCloseISOImageParam) validate() error {
	var errors []error

	return utils.FlattenErrors(errors)
}

func (p *FTPCloseISOImageParam) ResourceDef() *schema.Resource {
	return define.Resources["ISOImage"]
}

func (p *FTPCloseISOImageParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["ftp-close"]
}

func (p *FTPCloseISOImageParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *FTPCloseISOImageParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *FTPCloseISOImageParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *FTPCloseISOImageParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}
