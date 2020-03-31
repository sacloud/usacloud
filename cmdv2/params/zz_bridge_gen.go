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

// ListBridgeParam is input parameters for the sacloud API
type ListBridgeParam struct {
	Name []string
	Id   []sacloud.ID
	From int
	Max  int
	Sort []string

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

func (p *ListBridgeParam) fillValueToSkeleton() {
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

// CreateBridgeParam is input parameters for the sacloud API
type CreateBridgeParam struct {
	Name        string
	Description string

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

func (p *CreateBridgeParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.Name) {
		p.Name = ""
	}
	if utils.IsEmpty(p.Description) {
		p.Description = ""
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

// ReadBridgeParam is input parameters for the sacloud API
type ReadBridgeParam struct {
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

func (p *ReadBridgeParam) fillValueToSkeleton() {

}

func (p *ReadBridgeParam) validate() error {
	var errors []error

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

// UpdateBridgeParam is input parameters for the sacloud API
type UpdateBridgeParam struct {
	Name        string
	Description string

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

func (p *UpdateBridgeParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.Name) {
		p.Name = ""
	}
	if utils.IsEmpty(p.Description) {
		p.Description = ""
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

// DeleteBridgeParam is input parameters for the sacloud API
type DeleteBridgeParam struct {
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

func (p *DeleteBridgeParam) fillValueToSkeleton() {

}

func (p *DeleteBridgeParam) validate() error {
	var errors []error

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
