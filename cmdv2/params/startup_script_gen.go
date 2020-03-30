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

// ListStartupscriptParam is input parameters for the sacloud API
type ListStartupscriptParam struct {
	Tags  []string
	Class []string
	Max   int
	Sort  []string
	Name  []string
	Id    []sacloud.ID
	From  int
	Scope string

	input Input
}

// NewListStartupscriptParam return new ListStartupscriptParam
func NewListStartupscriptParam() *ListStartupscriptParam {
	return &ListStartupscriptParam{}
}

// Initialize init ListStartupscriptParam
func (p *ListStartupscriptParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ListStartupscriptParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *ListStartupscriptParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.Tags) {
		p.Tags = []string{""}
	}
	if utils.IsEmpty(p.Class) {
		p.Class = []string{""}
	}
	if utils.IsEmpty(p.Max) {
		p.Max = 0
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
	if utils.IsEmpty(p.From) {
		p.From = 0
	}
	if utils.IsEmpty(p.Scope) {
		p.Scope = ""
	}

}

func (p *ListStartupscriptParam) validate() error {
	var errors []error

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

	{
		validator := define.Resources["StartupScript"].Commands["list"].Params["scope"].ValidateFunc
		errs := validator("--scope", p.Scope)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return utils.FlattenErrors(errors)
}

func (p *ListStartupscriptParam) ResourceDef() *schema.Resource {
	return define.Resources["StartupScript"]
}

func (p *ListStartupscriptParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["list"]
}

func (p *ListStartupscriptParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ListStartupscriptParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ListStartupscriptParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ListStartupscriptParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ListStartupscriptParam) SetTags(v []string) {
	p.Tags = v
}

func (p *ListStartupscriptParam) GetTags() []string {
	return p.Tags
}
func (p *ListStartupscriptParam) SetClass(v []string) {
	p.Class = v
}

func (p *ListStartupscriptParam) GetClass() []string {
	return p.Class
}
func (p *ListStartupscriptParam) SetMax(v int) {
	p.Max = v
}

func (p *ListStartupscriptParam) GetMax() int {
	return p.Max
}
func (p *ListStartupscriptParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListStartupscriptParam) GetSort() []string {
	return p.Sort
}
func (p *ListStartupscriptParam) SetName(v []string) {
	p.Name = v
}

func (p *ListStartupscriptParam) GetName() []string {
	return p.Name
}
func (p *ListStartupscriptParam) SetId(v []sacloud.ID) {
	p.Id = v
}

func (p *ListStartupscriptParam) GetId() []sacloud.ID {
	return p.Id
}
func (p *ListStartupscriptParam) SetFrom(v int) {
	p.From = v
}

func (p *ListStartupscriptParam) GetFrom() int {
	return p.From
}
func (p *ListStartupscriptParam) SetScope(v string) {
	p.Scope = v
}

func (p *ListStartupscriptParam) GetScope() string {
	return p.Scope
}

// CreateStartupscriptParam is input parameters for the sacloud API
type CreateStartupscriptParam struct {
	Script        string
	Class         string
	Name          string
	Tags          []string
	IconId        sacloud.ID
	ScriptContent string

	input Input
}

// NewCreateStartupscriptParam return new CreateStartupscriptParam
func NewCreateStartupscriptParam() *CreateStartupscriptParam {
	return &CreateStartupscriptParam{
		Class: "shell"}
}

// Initialize init CreateStartupscriptParam
func (p *CreateStartupscriptParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *CreateStartupscriptParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *CreateStartupscriptParam) fillValueToSkeleton() {
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
	if utils.IsEmpty(p.ScriptContent) {
		p.ScriptContent = ""
	}

}

func (p *CreateStartupscriptParam) validate() error {
	var errors []error

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

	return utils.FlattenErrors(errors)
}

func (p *CreateStartupscriptParam) ResourceDef() *schema.Resource {
	return define.Resources["StartupScript"]
}

func (p *CreateStartupscriptParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["create"]
}

func (p *CreateStartupscriptParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *CreateStartupscriptParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *CreateStartupscriptParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *CreateStartupscriptParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *CreateStartupscriptParam) SetScript(v string) {
	p.Script = v
}

func (p *CreateStartupscriptParam) GetScript() string {
	return p.Script
}
func (p *CreateStartupscriptParam) SetClass(v string) {
	p.Class = v
}

func (p *CreateStartupscriptParam) GetClass() string {
	return p.Class
}
func (p *CreateStartupscriptParam) SetName(v string) {
	p.Name = v
}

func (p *CreateStartupscriptParam) GetName() string {
	return p.Name
}
func (p *CreateStartupscriptParam) SetTags(v []string) {
	p.Tags = v
}

func (p *CreateStartupscriptParam) GetTags() []string {
	return p.Tags
}
func (p *CreateStartupscriptParam) SetIconId(v sacloud.ID) {
	p.IconId = v
}

func (p *CreateStartupscriptParam) GetIconId() sacloud.ID {
	return p.IconId
}
func (p *CreateStartupscriptParam) SetScriptContent(v string) {
	p.ScriptContent = v
}

func (p *CreateStartupscriptParam) GetScriptContent() string {
	return p.ScriptContent
}

// ReadStartupscriptParam is input parameters for the sacloud API
type ReadStartupscriptParam struct {
	input Input
}

// NewReadStartupscriptParam return new ReadStartupscriptParam
func NewReadStartupscriptParam() *ReadStartupscriptParam {
	return &ReadStartupscriptParam{}
}

// Initialize init ReadStartupscriptParam
func (p *ReadStartupscriptParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ReadStartupscriptParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *ReadStartupscriptParam) fillValueToSkeleton() {

}

func (p *ReadStartupscriptParam) validate() error {
	var errors []error

	return utils.FlattenErrors(errors)
}

func (p *ReadStartupscriptParam) ResourceDef() *schema.Resource {
	return define.Resources["StartupScript"]
}

func (p *ReadStartupscriptParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["read"]
}

func (p *ReadStartupscriptParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ReadStartupscriptParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ReadStartupscriptParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ReadStartupscriptParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

// UpdateStartupscriptParam is input parameters for the sacloud API
type UpdateStartupscriptParam struct {
	Class         string
	Name          string
	Tags          []string
	IconId        sacloud.ID
	ScriptContent string
	Script        string

	input Input
}

// NewUpdateStartupscriptParam return new UpdateStartupscriptParam
func NewUpdateStartupscriptParam() *UpdateStartupscriptParam {
	return &UpdateStartupscriptParam{}
}

// Initialize init UpdateStartupscriptParam
func (p *UpdateStartupscriptParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *UpdateStartupscriptParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *UpdateStartupscriptParam) fillValueToSkeleton() {
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
	if utils.IsEmpty(p.ScriptContent) {
		p.ScriptContent = ""
	}
	if utils.IsEmpty(p.Script) {
		p.Script = ""
	}

}

func (p *UpdateStartupscriptParam) validate() error {
	var errors []error

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

	return utils.FlattenErrors(errors)
}

func (p *UpdateStartupscriptParam) ResourceDef() *schema.Resource {
	return define.Resources["StartupScript"]
}

func (p *UpdateStartupscriptParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["update"]
}

func (p *UpdateStartupscriptParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *UpdateStartupscriptParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *UpdateStartupscriptParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *UpdateStartupscriptParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *UpdateStartupscriptParam) SetClass(v string) {
	p.Class = v
}

func (p *UpdateStartupscriptParam) GetClass() string {
	return p.Class
}
func (p *UpdateStartupscriptParam) SetName(v string) {
	p.Name = v
}

func (p *UpdateStartupscriptParam) GetName() string {
	return p.Name
}
func (p *UpdateStartupscriptParam) SetTags(v []string) {
	p.Tags = v
}

func (p *UpdateStartupscriptParam) GetTags() []string {
	return p.Tags
}
func (p *UpdateStartupscriptParam) SetIconId(v sacloud.ID) {
	p.IconId = v
}

func (p *UpdateStartupscriptParam) GetIconId() sacloud.ID {
	return p.IconId
}
func (p *UpdateStartupscriptParam) SetScriptContent(v string) {
	p.ScriptContent = v
}

func (p *UpdateStartupscriptParam) GetScriptContent() string {
	return p.ScriptContent
}
func (p *UpdateStartupscriptParam) SetScript(v string) {
	p.Script = v
}

func (p *UpdateStartupscriptParam) GetScript() string {
	return p.Script
}

// DeleteStartupscriptParam is input parameters for the sacloud API
type DeleteStartupscriptParam struct {
	input Input
}

// NewDeleteStartupscriptParam return new DeleteStartupscriptParam
func NewDeleteStartupscriptParam() *DeleteStartupscriptParam {
	return &DeleteStartupscriptParam{}
}

// Initialize init DeleteStartupscriptParam
func (p *DeleteStartupscriptParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *DeleteStartupscriptParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *DeleteStartupscriptParam) fillValueToSkeleton() {

}

func (p *DeleteStartupscriptParam) validate() error {
	var errors []error

	return utils.FlattenErrors(errors)
}

func (p *DeleteStartupscriptParam) ResourceDef() *schema.Resource {
	return define.Resources["StartupScript"]
}

func (p *DeleteStartupscriptParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["delete"]
}

func (p *DeleteStartupscriptParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *DeleteStartupscriptParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *DeleteStartupscriptParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *DeleteStartupscriptParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}
