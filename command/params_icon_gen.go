// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package command

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// ListIconParam is input parameters for the sacloud API
type ListIconParam struct {
	Name  []string
	Id    []int64
	From  int
	Max   int
	Sort  []string
	Scope string
}

// NewListIconParam return new ListIconParam
func NewListIconParam() *ListIconParam {
	return &ListIconParam{}
}

// Validate checks current values in model
func (p *ListIconParam) Validate() []error {
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
		validator := define.Resources["Icon"].Commands["list"].Params["id"].ValidateFunc
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
	{
		validator := define.Resources["Icon"].Commands["list"].Params["scope"].ValidateFunc
		errs := validator("--scope", p.Scope)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ListIconParam) getResourceDef() *schema.Resource {
	return define.Resources["Icon"]
}

func (p *ListIconParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["list"]
}

func (p *ListIconParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ListIconParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ListIconParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ListIconParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ListIconParam) SetName(v []string) {
	p.Name = v
}

func (p *ListIconParam) GetName() []string {
	return p.Name
}
func (p *ListIconParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListIconParam) GetId() []int64 {
	return p.Id
}
func (p *ListIconParam) SetFrom(v int) {
	p.From = v
}

func (p *ListIconParam) GetFrom() int {
	return p.From
}
func (p *ListIconParam) SetMax(v int) {
	p.Max = v
}

func (p *ListIconParam) GetMax() int {
	return p.Max
}
func (p *ListIconParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListIconParam) GetSort() []string {
	return p.Sort
}
func (p *ListIconParam) SetScope(v string) {
	p.Scope = v
}

func (p *ListIconParam) GetScope() string {
	return p.Scope
}

// CreateIconParam is input parameters for the sacloud API
type CreateIconParam struct {
	Name  string
	Tags  []string
	Image string
}

// NewCreateIconParam return new CreateIconParam
func NewCreateIconParam() *CreateIconParam {
	return &CreateIconParam{}
}

// Validate checks current values in model
func (p *CreateIconParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Icon"].Commands["create"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Icon"].Commands["create"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--image", p.Image)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Icon"].Commands["create"].Params["image"].ValidateFunc
		errs := validator("--image", p.Image)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *CreateIconParam) getResourceDef() *schema.Resource {
	return define.Resources["Icon"]
}

func (p *CreateIconParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["create"]
}

func (p *CreateIconParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *CreateIconParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *CreateIconParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *CreateIconParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *CreateIconParam) SetName(v string) {
	p.Name = v
}

func (p *CreateIconParam) GetName() string {
	return p.Name
}
func (p *CreateIconParam) SetTags(v []string) {
	p.Tags = v
}

func (p *CreateIconParam) GetTags() []string {
	return p.Tags
}
func (p *CreateIconParam) SetImage(v string) {
	p.Image = v
}

func (p *CreateIconParam) GetImage() string {
	return p.Image
}

// ReadIconParam is input parameters for the sacloud API
type ReadIconParam struct {
	Id int64
}

// NewReadIconParam return new ReadIconParam
func NewReadIconParam() *ReadIconParam {
	return &ReadIconParam{}
}

// Validate checks current values in model
func (p *ReadIconParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Icon"].Commands["read"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ReadIconParam) getResourceDef() *schema.Resource {
	return define.Resources["Icon"]
}

func (p *ReadIconParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["read"]
}

func (p *ReadIconParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ReadIconParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ReadIconParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ReadIconParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ReadIconParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadIconParam) GetId() int64 {
	return p.Id
}

// UpdateIconParam is input parameters for the sacloud API
type UpdateIconParam struct {
	Tags []string
	Id   int64
	Name string
}

// NewUpdateIconParam return new UpdateIconParam
func NewUpdateIconParam() *UpdateIconParam {
	return &UpdateIconParam{}
}

// Validate checks current values in model
func (p *UpdateIconParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["Icon"].Commands["update"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
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
		validator := define.Resources["Icon"].Commands["update"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Icon"].Commands["update"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *UpdateIconParam) getResourceDef() *schema.Resource {
	return define.Resources["Icon"]
}

func (p *UpdateIconParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["update"]
}

func (p *UpdateIconParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *UpdateIconParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *UpdateIconParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *UpdateIconParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *UpdateIconParam) SetTags(v []string) {
	p.Tags = v
}

func (p *UpdateIconParam) GetTags() []string {
	return p.Tags
}
func (p *UpdateIconParam) SetId(v int64) {
	p.Id = v
}

func (p *UpdateIconParam) GetId() int64 {
	return p.Id
}
func (p *UpdateIconParam) SetName(v string) {
	p.Name = v
}

func (p *UpdateIconParam) GetName() string {
	return p.Name
}

// DeleteIconParam is input parameters for the sacloud API
type DeleteIconParam struct {
	Id int64
}

// NewDeleteIconParam return new DeleteIconParam
func NewDeleteIconParam() *DeleteIconParam {
	return &DeleteIconParam{}
}

// Validate checks current values in model
func (p *DeleteIconParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Icon"].Commands["delete"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *DeleteIconParam) getResourceDef() *schema.Resource {
	return define.Resources["Icon"]
}

func (p *DeleteIconParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["delete"]
}

func (p *DeleteIconParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *DeleteIconParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *DeleteIconParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *DeleteIconParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *DeleteIconParam) SetId(v int64) {
	p.Id = v
}

func (p *DeleteIconParam) GetId() int64 {
	return p.Id
}
