// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package command

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// ListStartupScriptParam is input parameters for the sacloud API
type ListStartupScriptParam struct {
	Sort  []string
	Name  []string
	Scope string
	Id    []int64
	From  int
	Max   int
}

// NewListStartupScriptParam return new ListStartupScriptParam
func NewListStartupScriptParam() *ListStartupScriptParam {
	return &ListStartupScriptParam{}
}

// Validate checks current values in model
func (p *ListStartupScriptParam) Validate() []error {
	errors := []error{}
	{
		errs := validateConflicts("--name", p.Name, map[string]interface{}{

			"--id": p.Id,
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
	{
		validator := define.Resources["StartupScript"].Commands["list"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateConflicts("--id", p.Id, map[string]interface{}{

			"--name": p.Name,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ListStartupScriptParam) getResourceDef() *schema.Resource {
	return define.Resources["StartupScript"]
}

func (p *ListStartupScriptParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["list"]
}

func (p *ListStartupScriptParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ListStartupScriptParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ListStartupScriptParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ListStartupScriptParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ListStartupScriptParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListStartupScriptParam) GetSort() []string {
	return p.Sort
}
func (p *ListStartupScriptParam) SetName(v []string) {
	p.Name = v
}

func (p *ListStartupScriptParam) GetName() []string {
	return p.Name
}
func (p *ListStartupScriptParam) SetScope(v string) {
	p.Scope = v
}

func (p *ListStartupScriptParam) GetScope() string {
	return p.Scope
}
func (p *ListStartupScriptParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListStartupScriptParam) GetId() []int64 {
	return p.Id
}
func (p *ListStartupScriptParam) SetFrom(v int) {
	p.From = v
}

func (p *ListStartupScriptParam) GetFrom() int {
	return p.From
}
func (p *ListStartupScriptParam) SetMax(v int) {
	p.Max = v
}

func (p *ListStartupScriptParam) GetMax() int {
	return p.Max
}

// CreateStartupScriptParam is input parameters for the sacloud API
type CreateStartupScriptParam struct {
	Name          string
	Tags          []string
	IconId        int64
	ScriptContent string
	Script        string
}

// NewCreateStartupScriptParam return new CreateStartupScriptParam
func NewCreateStartupScriptParam() *CreateStartupScriptParam {
	return &CreateStartupScriptParam{}
}

// Validate checks current values in model
func (p *CreateStartupScriptParam) Validate() []error {
	errors := []error{}
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
		errs := validateConflicts("--script-content", p.ScriptContent, map[string]interface{}{

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

	return errors
}

func (p *CreateStartupScriptParam) getResourceDef() *schema.Resource {
	return define.Resources["StartupScript"]
}

func (p *CreateStartupScriptParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["create"]
}

func (p *CreateStartupScriptParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *CreateStartupScriptParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *CreateStartupScriptParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *CreateStartupScriptParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
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
func (p *CreateStartupScriptParam) SetIconId(v int64) {
	p.IconId = v
}

func (p *CreateStartupScriptParam) GetIconId() int64 {
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

// ReadStartupScriptParam is input parameters for the sacloud API
type ReadStartupScriptParam struct {
	Id int64
}

// NewReadStartupScriptParam return new ReadStartupScriptParam
func NewReadStartupScriptParam() *ReadStartupScriptParam {
	return &ReadStartupScriptParam{}
}

// Validate checks current values in model
func (p *ReadStartupScriptParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["StartupScript"].Commands["read"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ReadStartupScriptParam) getResourceDef() *schema.Resource {
	return define.Resources["StartupScript"]
}

func (p *ReadStartupScriptParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["read"]
}

func (p *ReadStartupScriptParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ReadStartupScriptParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ReadStartupScriptParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ReadStartupScriptParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ReadStartupScriptParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadStartupScriptParam) GetId() int64 {
	return p.Id
}

// UpdateStartupScriptParam is input parameters for the sacloud API
type UpdateStartupScriptParam struct {
	Tags          []string
	IconId        int64
	ScriptContent string
	Script        string
	Id            int64
	Name          string
}

// NewUpdateStartupScriptParam return new UpdateStartupScriptParam
func NewUpdateStartupScriptParam() *UpdateStartupScriptParam {
	return &UpdateStartupScriptParam{}
}

// Validate checks current values in model
func (p *UpdateStartupScriptParam) Validate() []error {
	errors := []error{}
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
		errs := validateConflicts("--script-content", p.ScriptContent, map[string]interface{}{

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
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["StartupScript"].Commands["update"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
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
		validator := define.Resources["StartupScript"].Commands["update"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *UpdateStartupScriptParam) getResourceDef() *schema.Resource {
	return define.Resources["StartupScript"]
}

func (p *UpdateStartupScriptParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["update"]
}

func (p *UpdateStartupScriptParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *UpdateStartupScriptParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *UpdateStartupScriptParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *UpdateStartupScriptParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *UpdateStartupScriptParam) SetTags(v []string) {
	p.Tags = v
}

func (p *UpdateStartupScriptParam) GetTags() []string {
	return p.Tags
}
func (p *UpdateStartupScriptParam) SetIconId(v int64) {
	p.IconId = v
}

func (p *UpdateStartupScriptParam) GetIconId() int64 {
	return p.IconId
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
func (p *UpdateStartupScriptParam) SetId(v int64) {
	p.Id = v
}

func (p *UpdateStartupScriptParam) GetId() int64 {
	return p.Id
}
func (p *UpdateStartupScriptParam) SetName(v string) {
	p.Name = v
}

func (p *UpdateStartupScriptParam) GetName() string {
	return p.Name
}

// DeleteStartupScriptParam is input parameters for the sacloud API
type DeleteStartupScriptParam struct {
	Id int64
}

// NewDeleteStartupScriptParam return new DeleteStartupScriptParam
func NewDeleteStartupScriptParam() *DeleteStartupScriptParam {
	return &DeleteStartupScriptParam{}
}

// Validate checks current values in model
func (p *DeleteStartupScriptParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["StartupScript"].Commands["delete"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *DeleteStartupScriptParam) getResourceDef() *schema.Resource {
	return define.Resources["StartupScript"]
}

func (p *DeleteStartupScriptParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["delete"]
}

func (p *DeleteStartupScriptParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *DeleteStartupScriptParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *DeleteStartupScriptParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *DeleteStartupScriptParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *DeleteStartupScriptParam) SetId(v int64) {
	p.Id = v
}

func (p *DeleteStartupScriptParam) GetId() int64 {
	return p.Id
}
