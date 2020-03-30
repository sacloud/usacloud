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

// ListStartupScriptParam is input parameters for the sacloud API
type ListStartupScriptParam struct {
	Max   int
	Sort  []string
	Scope string
	Tags  []string
	Class []string
	Name  []string
	Id    []sacloud.ID
	From  int

	input Input
}

// NewListStartupScriptParam return new ListStartupScriptParam
func NewListStartupScriptParam() *ListStartupScriptParam {
	return &ListStartupScriptParam{}
}

// Initialize init ListStartupScriptParam
func (p *ListStartupScriptParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ListStartupScriptParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *ListStartupScriptParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.Max) {
		p.Max = 0
	}
	if utils.IsEmpty(p.Sort) {
		p.Sort = []string{""}
	}
	if utils.IsEmpty(p.Scope) {
		p.Scope = ""
	}
	if utils.IsEmpty(p.Tags) {
		p.Tags = []string{""}
	}
	if utils.IsEmpty(p.Class) {
		p.Class = []string{""}
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

}

func (p *ListStartupScriptParam) validate() error {
	var errors []error

	{
		validator := define.Resources["StartupScript"].Commands["list"].Params["scope"].ValidateFunc
		errs := validator("--scope", p.Scope)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["StartupScript"].Commands["list"].Params["tags"].ValidateFunc
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
		validator := define.Resources["StartupScript"].Commands["list"].Params["id"].ValidateFunc
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

func (p *ListStartupScriptParam) ResourceDef() *schema.Resource {
	return define.Resources["StartupScript"]
}

func (p *ListStartupScriptParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["list"]
}

func (p *ListStartupScriptParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ListStartupScriptParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ListStartupScriptParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ListStartupScriptParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ListStartupScriptParam) SetMax(v int) {
	p.Max = v
}

func (p *ListStartupScriptParam) GetMax() int {
	return p.Max
}
func (p *ListStartupScriptParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListStartupScriptParam) GetSort() []string {
	return p.Sort
}
func (p *ListStartupScriptParam) SetScope(v string) {
	p.Scope = v
}

func (p *ListStartupScriptParam) GetScope() string {
	return p.Scope
}
func (p *ListStartupScriptParam) SetTags(v []string) {
	p.Tags = v
}

func (p *ListStartupScriptParam) GetTags() []string {
	return p.Tags
}
func (p *ListStartupScriptParam) SetClass(v []string) {
	p.Class = v
}

func (p *ListStartupScriptParam) GetClass() []string {
	return p.Class
}
func (p *ListStartupScriptParam) SetName(v []string) {
	p.Name = v
}

func (p *ListStartupScriptParam) GetName() []string {
	return p.Name
}
func (p *ListStartupScriptParam) SetId(v []sacloud.ID) {
	p.Id = v
}

func (p *ListStartupScriptParam) GetId() []sacloud.ID {
	return p.Id
}
func (p *ListStartupScriptParam) SetFrom(v int) {
	p.From = v
}

func (p *ListStartupScriptParam) GetFrom() int {
	return p.From
}

// CreateStartupScriptParam is input parameters for the sacloud API
type CreateStartupScriptParam struct {
	Name          string
	Tags          []string
	IconId        sacloud.ID
	ScriptContent string
	Script        string
	Class         string

	input Input
}

// NewCreateStartupScriptParam return new CreateStartupScriptParam
func NewCreateStartupScriptParam() *CreateStartupScriptParam {
	return &CreateStartupScriptParam{
		Class: "shell"}
}

// Initialize init CreateStartupScriptParam
func (p *CreateStartupScriptParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *CreateStartupScriptParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *CreateStartupScriptParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.Name) {
		p.Name = ""
	}
	if utils.IsEmpty(p.Tags) {
		p.Tags = []string{""}
	}
	if utils.IsEmpty(p.IconId) {
		p.IconId = sacloud.ID(0)
	}
	if utils.IsEmpty(p.ScriptContent) {
		p.ScriptContent = ""
	}
	if utils.IsEmpty(p.Script) {
		p.Script = ""
	}
	if utils.IsEmpty(p.Class) {
		p.Class = ""
	}

}

func (p *CreateStartupScriptParam) validate() error {
	var errors []error

	{
		validator := validateRequired
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["StartupScript"].Commands["create"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["StartupScript"].Commands["create"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["StartupScript"].Commands["create"].Params["icon-id"].ValidateFunc
		errs := validator("--icon-id", p.IconId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		errs := validation.ConflictsWith("--script-content", p.ScriptContent, map[string]interface{}{

			"--script": p.Script,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["StartupScript"].Commands["create"].Params["script"].ValidateFunc
		errs := validator("--script", p.Script)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := validateRequired
		errs := validator("--class", p.Class)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["StartupScript"].Commands["create"].Params["class"].ValidateFunc
		errs := validator("--class", p.Class)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return utils.FlattenErrors(errors)
}

func (p *CreateStartupScriptParam) ResourceDef() *schema.Resource {
	return define.Resources["StartupScript"]
}

func (p *CreateStartupScriptParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["create"]
}

func (p *CreateStartupScriptParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *CreateStartupScriptParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *CreateStartupScriptParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *CreateStartupScriptParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *CreateStartupScriptParam) SetName(v string) {
	p.Name = v
}

func (p *CreateStartupScriptParam) GetName() string {
	return p.Name
}
func (p *CreateStartupScriptParam) SetTags(v []string) {
	p.Tags = v
}

func (p *CreateStartupScriptParam) GetTags() []string {
	return p.Tags
}
func (p *CreateStartupScriptParam) SetIconId(v sacloud.ID) {
	p.IconId = v
}

func (p *CreateStartupScriptParam) GetIconId() sacloud.ID {
	return p.IconId
}
func (p *CreateStartupScriptParam) SetScriptContent(v string) {
	p.ScriptContent = v
}

func (p *CreateStartupScriptParam) GetScriptContent() string {
	return p.ScriptContent
}
func (p *CreateStartupScriptParam) SetScript(v string) {
	p.Script = v
}

func (p *CreateStartupScriptParam) GetScript() string {
	return p.Script
}
func (p *CreateStartupScriptParam) SetClass(v string) {
	p.Class = v
}

func (p *CreateStartupScriptParam) GetClass() string {
	return p.Class
}

// ReadStartupScriptParam is input parameters for the sacloud API
type ReadStartupScriptParam struct {
	input Input
}

// NewReadStartupScriptParam return new ReadStartupScriptParam
func NewReadStartupScriptParam() *ReadStartupScriptParam {
	return &ReadStartupScriptParam{}
}

// Initialize init ReadStartupScriptParam
func (p *ReadStartupScriptParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ReadStartupScriptParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *ReadStartupScriptParam) fillValueToSkeleton() {

}

func (p *ReadStartupScriptParam) validate() error {
	var errors []error

	return utils.FlattenErrors(errors)
}

func (p *ReadStartupScriptParam) ResourceDef() *schema.Resource {
	return define.Resources["StartupScript"]
}

func (p *ReadStartupScriptParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["read"]
}

func (p *ReadStartupScriptParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ReadStartupScriptParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ReadStartupScriptParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ReadStartupScriptParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

// UpdateStartupScriptParam is input parameters for the sacloud API
type UpdateStartupScriptParam struct {
	ScriptContent string
	Script        string
	Class         string
	Name          string
	Tags          []string
	IconId        sacloud.ID

	input Input
}

// NewUpdateStartupScriptParam return new UpdateStartupScriptParam
func NewUpdateStartupScriptParam() *UpdateStartupScriptParam {
	return &UpdateStartupScriptParam{}
}

// Initialize init UpdateStartupScriptParam
func (p *UpdateStartupScriptParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *UpdateStartupScriptParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *UpdateStartupScriptParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.ScriptContent) {
		p.ScriptContent = ""
	}
	if utils.IsEmpty(p.Script) {
		p.Script = ""
	}
	if utils.IsEmpty(p.Class) {
		p.Class = ""
	}
	if utils.IsEmpty(p.Name) {
		p.Name = ""
	}
	if utils.IsEmpty(p.Tags) {
		p.Tags = []string{""}
	}
	if utils.IsEmpty(p.IconId) {
		p.IconId = sacloud.ID(0)
	}

}

func (p *UpdateStartupScriptParam) validate() error {
	var errors []error

	{
		errs := validation.ConflictsWith("--script-content", p.ScriptContent, map[string]interface{}{

			"--script": p.Script,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["StartupScript"].Commands["update"].Params["script"].ValidateFunc
		errs := validator("--script", p.Script)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["StartupScript"].Commands["update"].Params["class"].ValidateFunc
		errs := validator("--class", p.Class)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["StartupScript"].Commands["update"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["StartupScript"].Commands["update"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["StartupScript"].Commands["update"].Params["icon-id"].ValidateFunc
		errs := validator("--icon-id", p.IconId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return utils.FlattenErrors(errors)
}

func (p *UpdateStartupScriptParam) ResourceDef() *schema.Resource {
	return define.Resources["StartupScript"]
}

func (p *UpdateStartupScriptParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["update"]
}

func (p *UpdateStartupScriptParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *UpdateStartupScriptParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *UpdateStartupScriptParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *UpdateStartupScriptParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *UpdateStartupScriptParam) SetScriptContent(v string) {
	p.ScriptContent = v
}

func (p *UpdateStartupScriptParam) GetScriptContent() string {
	return p.ScriptContent
}
func (p *UpdateStartupScriptParam) SetScript(v string) {
	p.Script = v
}

func (p *UpdateStartupScriptParam) GetScript() string {
	return p.Script
}
func (p *UpdateStartupScriptParam) SetClass(v string) {
	p.Class = v
}

func (p *UpdateStartupScriptParam) GetClass() string {
	return p.Class
}
func (p *UpdateStartupScriptParam) SetName(v string) {
	p.Name = v
}

func (p *UpdateStartupScriptParam) GetName() string {
	return p.Name
}
func (p *UpdateStartupScriptParam) SetTags(v []string) {
	p.Tags = v
}

func (p *UpdateStartupScriptParam) GetTags() []string {
	return p.Tags
}
func (p *UpdateStartupScriptParam) SetIconId(v sacloud.ID) {
	p.IconId = v
}

func (p *UpdateStartupScriptParam) GetIconId() sacloud.ID {
	return p.IconId
}

// DeleteStartupScriptParam is input parameters for the sacloud API
type DeleteStartupScriptParam struct {
	input Input
}

// NewDeleteStartupScriptParam return new DeleteStartupScriptParam
func NewDeleteStartupScriptParam() *DeleteStartupScriptParam {
	return &DeleteStartupScriptParam{}
}

// Initialize init DeleteStartupScriptParam
func (p *DeleteStartupScriptParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *DeleteStartupScriptParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *DeleteStartupScriptParam) fillValueToSkeleton() {

}

func (p *DeleteStartupScriptParam) validate() error {
	var errors []error

	return utils.FlattenErrors(errors)
}

func (p *DeleteStartupScriptParam) ResourceDef() *schema.Resource {
	return define.Resources["StartupScript"]
}

func (p *DeleteStartupScriptParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["delete"]
}

func (p *DeleteStartupScriptParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *DeleteStartupScriptParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *DeleteStartupScriptParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *DeleteStartupScriptParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}
