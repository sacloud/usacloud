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

// ListSSHKeyParam is input parameters for the sacloud API
type ListSSHKeyParam struct {
	Sort []string
	Name []string
	Id   []sacloud.ID
	From int
	Max  int

	input Input
}

// NewListSSHKeyParam return new ListSSHKeyParam
func NewListSSHKeyParam() *ListSSHKeyParam {
	return &ListSSHKeyParam{}
}

// Initialize init ListSSHKeyParam
func (p *ListSSHKeyParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ListSSHKeyParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *ListSSHKeyParam) fillValueToSkeleton() {
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

func (p *ListSSHKeyParam) validate() error {
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
		validator := define.Resources["SSHKey"].Commands["list"].Params["id"].ValidateFunc
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

func (p *ListSSHKeyParam) ResourceDef() *schema.Resource {
	return define.Resources["SSHKey"]
}

func (p *ListSSHKeyParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["list"]
}

func (p *ListSSHKeyParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ListSSHKeyParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ListSSHKeyParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ListSSHKeyParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ListSSHKeyParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListSSHKeyParam) GetSort() []string {
	return p.Sort
}
func (p *ListSSHKeyParam) SetName(v []string) {
	p.Name = v
}

func (p *ListSSHKeyParam) GetName() []string {
	return p.Name
}
func (p *ListSSHKeyParam) SetId(v []sacloud.ID) {
	p.Id = v
}

func (p *ListSSHKeyParam) GetId() []sacloud.ID {
	return p.Id
}
func (p *ListSSHKeyParam) SetFrom(v int) {
	p.From = v
}

func (p *ListSSHKeyParam) GetFrom() int {
	return p.From
}
func (p *ListSSHKeyParam) SetMax(v int) {
	p.Max = v
}

func (p *ListSSHKeyParam) GetMax() int {
	return p.Max
}

// CreateSSHKeyParam is input parameters for the sacloud API
type CreateSSHKeyParam struct {
	PublicKeyContent string
	PublicKey        string
	Name             string
	Description      string

	input Input
}

// NewCreateSSHKeyParam return new CreateSSHKeyParam
func NewCreateSSHKeyParam() *CreateSSHKeyParam {
	return &CreateSSHKeyParam{}
}

// Initialize init CreateSSHKeyParam
func (p *CreateSSHKeyParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *CreateSSHKeyParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *CreateSSHKeyParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.PublicKeyContent) {
		p.PublicKeyContent = ""
	}
	if utils.IsEmpty(p.PublicKey) {
		p.PublicKey = ""
	}
	if utils.IsEmpty(p.Name) {
		p.Name = ""
	}
	if utils.IsEmpty(p.Description) {
		p.Description = ""
	}

}

func (p *CreateSSHKeyParam) validate() error {
	var errors []error

	{
		errs := validation.ConflictsWith("--public-key-content", p.PublicKeyContent, map[string]interface{}{

			"--public-key": p.PublicKey,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["SSHKey"].Commands["create"].Params["public-key"].ValidateFunc
		errs := validator("--public-key", p.PublicKey)
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
		validator := define.Resources["SSHKey"].Commands["create"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["SSHKey"].Commands["create"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return utils.FlattenErrors(errors)
}

func (p *CreateSSHKeyParam) ResourceDef() *schema.Resource {
	return define.Resources["SSHKey"]
}

func (p *CreateSSHKeyParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["create"]
}

func (p *CreateSSHKeyParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *CreateSSHKeyParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *CreateSSHKeyParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *CreateSSHKeyParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *CreateSSHKeyParam) SetPublicKeyContent(v string) {
	p.PublicKeyContent = v
}

func (p *CreateSSHKeyParam) GetPublicKeyContent() string {
	return p.PublicKeyContent
}
func (p *CreateSSHKeyParam) SetPublicKey(v string) {
	p.PublicKey = v
}

func (p *CreateSSHKeyParam) GetPublicKey() string {
	return p.PublicKey
}
func (p *CreateSSHKeyParam) SetName(v string) {
	p.Name = v
}

func (p *CreateSSHKeyParam) GetName() string {
	return p.Name
}
func (p *CreateSSHKeyParam) SetDescription(v string) {
	p.Description = v
}

func (p *CreateSSHKeyParam) GetDescription() string {
	return p.Description
}

// ReadSSHKeyParam is input parameters for the sacloud API
type ReadSSHKeyParam struct {
	input Input
}

// NewReadSSHKeyParam return new ReadSSHKeyParam
func NewReadSSHKeyParam() *ReadSSHKeyParam {
	return &ReadSSHKeyParam{}
}

// Initialize init ReadSSHKeyParam
func (p *ReadSSHKeyParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ReadSSHKeyParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *ReadSSHKeyParam) fillValueToSkeleton() {

}

func (p *ReadSSHKeyParam) validate() error {
	var errors []error

	return utils.FlattenErrors(errors)
}

func (p *ReadSSHKeyParam) ResourceDef() *schema.Resource {
	return define.Resources["SSHKey"]
}

func (p *ReadSSHKeyParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["read"]
}

func (p *ReadSSHKeyParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ReadSSHKeyParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ReadSSHKeyParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ReadSSHKeyParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

// UpdateSSHKeyParam is input parameters for the sacloud API
type UpdateSSHKeyParam struct {
	Name        string
	Description string

	input Input
}

// NewUpdateSSHKeyParam return new UpdateSSHKeyParam
func NewUpdateSSHKeyParam() *UpdateSSHKeyParam {
	return &UpdateSSHKeyParam{}
}

// Initialize init UpdateSSHKeyParam
func (p *UpdateSSHKeyParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *UpdateSSHKeyParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *UpdateSSHKeyParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.Name) {
		p.Name = ""
	}
	if utils.IsEmpty(p.Description) {
		p.Description = ""
	}

}

func (p *UpdateSSHKeyParam) validate() error {
	var errors []error

	{
		validator := define.Resources["SSHKey"].Commands["update"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["SSHKey"].Commands["update"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return utils.FlattenErrors(errors)
}

func (p *UpdateSSHKeyParam) ResourceDef() *schema.Resource {
	return define.Resources["SSHKey"]
}

func (p *UpdateSSHKeyParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["update"]
}

func (p *UpdateSSHKeyParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *UpdateSSHKeyParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *UpdateSSHKeyParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *UpdateSSHKeyParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *UpdateSSHKeyParam) SetName(v string) {
	p.Name = v
}

func (p *UpdateSSHKeyParam) GetName() string {
	return p.Name
}
func (p *UpdateSSHKeyParam) SetDescription(v string) {
	p.Description = v
}

func (p *UpdateSSHKeyParam) GetDescription() string {
	return p.Description
}

// DeleteSSHKeyParam is input parameters for the sacloud API
type DeleteSSHKeyParam struct {
	input Input
}

// NewDeleteSSHKeyParam return new DeleteSSHKeyParam
func NewDeleteSSHKeyParam() *DeleteSSHKeyParam {
	return &DeleteSSHKeyParam{}
}

// Initialize init DeleteSSHKeyParam
func (p *DeleteSSHKeyParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *DeleteSSHKeyParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *DeleteSSHKeyParam) fillValueToSkeleton() {

}

func (p *DeleteSSHKeyParam) validate() error {
	var errors []error

	return utils.FlattenErrors(errors)
}

func (p *DeleteSSHKeyParam) ResourceDef() *schema.Resource {
	return define.Resources["SSHKey"]
}

func (p *DeleteSSHKeyParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["delete"]
}

func (p *DeleteSSHKeyParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *DeleteSSHKeyParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *DeleteSSHKeyParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *DeleteSSHKeyParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

// GenerateSSHKeyParam is input parameters for the sacloud API
type GenerateSSHKeyParam struct {
	PassPhrase       string
	PrivateKeyOutput string
	Name             string
	Description      string

	input Input
}

// NewGenerateSSHKeyParam return new GenerateSSHKeyParam
func NewGenerateSSHKeyParam() *GenerateSSHKeyParam {
	return &GenerateSSHKeyParam{}
}

// Initialize init GenerateSSHKeyParam
func (p *GenerateSSHKeyParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *GenerateSSHKeyParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *GenerateSSHKeyParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.PassPhrase) {
		p.PassPhrase = ""
	}
	if utils.IsEmpty(p.PrivateKeyOutput) {
		p.PrivateKeyOutput = ""
	}
	if utils.IsEmpty(p.Name) {
		p.Name = ""
	}
	if utils.IsEmpty(p.Description) {
		p.Description = ""
	}

}

func (p *GenerateSSHKeyParam) validate() error {
	var errors []error

	{
		validator := define.Resources["SSHKey"].Commands["generate"].Params["pass-phrase"].ValidateFunc
		errs := validator("--pass-phrase", p.PassPhrase)
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
		validator := define.Resources["SSHKey"].Commands["generate"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["SSHKey"].Commands["generate"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return utils.FlattenErrors(errors)
}

func (p *GenerateSSHKeyParam) ResourceDef() *schema.Resource {
	return define.Resources["SSHKey"]
}

func (p *GenerateSSHKeyParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["generate"]
}

func (p *GenerateSSHKeyParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *GenerateSSHKeyParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *GenerateSSHKeyParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *GenerateSSHKeyParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *GenerateSSHKeyParam) SetPassPhrase(v string) {
	p.PassPhrase = v
}

func (p *GenerateSSHKeyParam) GetPassPhrase() string {
	return p.PassPhrase
}
func (p *GenerateSSHKeyParam) SetPrivateKeyOutput(v string) {
	p.PrivateKeyOutput = v
}

func (p *GenerateSSHKeyParam) GetPrivateKeyOutput() string {
	return p.PrivateKeyOutput
}
func (p *GenerateSSHKeyParam) SetName(v string) {
	p.Name = v
}

func (p *GenerateSSHKeyParam) GetName() string {
	return p.Name
}
func (p *GenerateSSHKeyParam) SetDescription(v string) {
	p.Description = v
}

func (p *GenerateSSHKeyParam) GetDescription() string {
	return p.Description
}
