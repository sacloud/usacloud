// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package command

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// ListLicenseParam is input parameters for the sacloud API
type ListLicenseParam struct {
	Sort []string
	Name []string
	Id   []int64
	From int
	Max  int
}

// NewListLicenseParam return new ListLicenseParam
func NewListLicenseParam() *ListLicenseParam {
	return &ListLicenseParam{}
}

// Validate checks current values in model
func (p *ListLicenseParam) Validate() []error {
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
		validator := define.Resources["License"].Commands["list"].Params["id"].ValidateFunc
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

func (p *ListLicenseParam) getResourceDef() *schema.Resource {
	return define.Resources["License"]
}

func (p *ListLicenseParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["list"]
}

func (p *ListLicenseParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ListLicenseParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ListLicenseParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ListLicenseParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ListLicenseParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListLicenseParam) GetSort() []string {
	return p.Sort
}
func (p *ListLicenseParam) SetName(v []string) {
	p.Name = v
}

func (p *ListLicenseParam) GetName() []string {
	return p.Name
}
func (p *ListLicenseParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListLicenseParam) GetId() []int64 {
	return p.Id
}
func (p *ListLicenseParam) SetFrom(v int) {
	p.From = v
}

func (p *ListLicenseParam) GetFrom() int {
	return p.From
}
func (p *ListLicenseParam) SetMax(v int) {
	p.Max = v
}

func (p *ListLicenseParam) GetMax() int {
	return p.Max
}

// CreateLicenseParam is input parameters for the sacloud API
type CreateLicenseParam struct {
	LicenseInfoId int64
	Name          string
}

// NewCreateLicenseParam return new CreateLicenseParam
func NewCreateLicenseParam() *CreateLicenseParam {
	return &CreateLicenseParam{}
}

// Validate checks current values in model
func (p *CreateLicenseParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["License"].Commands["create"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *CreateLicenseParam) getResourceDef() *schema.Resource {
	return define.Resources["License"]
}

func (p *CreateLicenseParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["create"]
}

func (p *CreateLicenseParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *CreateLicenseParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *CreateLicenseParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *CreateLicenseParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *CreateLicenseParam) SetLicenseInfoId(v int64) {
	p.LicenseInfoId = v
}

func (p *CreateLicenseParam) GetLicenseInfoId() int64 {
	return p.LicenseInfoId
}
func (p *CreateLicenseParam) SetName(v string) {
	p.Name = v
}

func (p *CreateLicenseParam) GetName() string {
	return p.Name
}

// ReadLicenseParam is input parameters for the sacloud API
type ReadLicenseParam struct {
	Id int64
}

// NewReadLicenseParam return new ReadLicenseParam
func NewReadLicenseParam() *ReadLicenseParam {
	return &ReadLicenseParam{}
}

// Validate checks current values in model
func (p *ReadLicenseParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["License"].Commands["read"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ReadLicenseParam) getResourceDef() *schema.Resource {
	return define.Resources["License"]
}

func (p *ReadLicenseParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["read"]
}

func (p *ReadLicenseParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ReadLicenseParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ReadLicenseParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ReadLicenseParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ReadLicenseParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadLicenseParam) GetId() int64 {
	return p.Id
}

// UpdateLicenseParam is input parameters for the sacloud API
type UpdateLicenseParam struct {
	Id   int64
	Name string
}

// NewUpdateLicenseParam return new UpdateLicenseParam
func NewUpdateLicenseParam() *UpdateLicenseParam {
	return &UpdateLicenseParam{}
}

// Validate checks current values in model
func (p *UpdateLicenseParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["License"].Commands["update"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["License"].Commands["update"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *UpdateLicenseParam) getResourceDef() *schema.Resource {
	return define.Resources["License"]
}

func (p *UpdateLicenseParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["update"]
}

func (p *UpdateLicenseParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *UpdateLicenseParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *UpdateLicenseParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *UpdateLicenseParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *UpdateLicenseParam) SetId(v int64) {
	p.Id = v
}

func (p *UpdateLicenseParam) GetId() int64 {
	return p.Id
}
func (p *UpdateLicenseParam) SetName(v string) {
	p.Name = v
}

func (p *UpdateLicenseParam) GetName() string {
	return p.Name
}

// DeleteLicenseParam is input parameters for the sacloud API
type DeleteLicenseParam struct {
	Id int64
}

// NewDeleteLicenseParam return new DeleteLicenseParam
func NewDeleteLicenseParam() *DeleteLicenseParam {
	return &DeleteLicenseParam{}
}

// Validate checks current values in model
func (p *DeleteLicenseParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["License"].Commands["delete"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *DeleteLicenseParam) getResourceDef() *schema.Resource {
	return define.Resources["License"]
}

func (p *DeleteLicenseParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["delete"]
}

func (p *DeleteLicenseParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *DeleteLicenseParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *DeleteLicenseParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *DeleteLicenseParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *DeleteLicenseParam) SetId(v int64) {
	p.Id = v
}

func (p *DeleteLicenseParam) GetId() int64 {
	return p.Id
}
