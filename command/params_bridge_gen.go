// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package command

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// ListBridgeParam is input parameters for the sacloud API
type ListBridgeParam struct {
	Name []string
	Id   []int64
	From int
	Max  int
	Sort []string
}

// NewListBridgeParam return new ListBridgeParam
func NewListBridgeParam() *ListBridgeParam {
	return &ListBridgeParam{}
}

// Validate checks current values in model
func (p *ListBridgeParam) Validate() []error {
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
		validator := define.Resources["Bridge"].Commands["list"].Params["id"].ValidateFunc
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

func (p *ListBridgeParam) getResourceDef() *schema.Resource {
	return define.Resources["Bridge"]
}

func (p *ListBridgeParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["list"]
}

func (p *ListBridgeParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ListBridgeParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ListBridgeParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ListBridgeParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ListBridgeParam) SetName(v []string) {
	p.Name = v
}

func (p *ListBridgeParam) GetName() []string {
	return p.Name
}
func (p *ListBridgeParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListBridgeParam) GetId() []int64 {
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
}

// NewCreateBridgeParam return new CreateBridgeParam
func NewCreateBridgeParam() *CreateBridgeParam {
	return &CreateBridgeParam{}
}

// Validate checks current values in model
func (p *CreateBridgeParam) Validate() []error {
	errors := []error{}
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

	return errors
}

func (p *CreateBridgeParam) getResourceDef() *schema.Resource {
	return define.Resources["Bridge"]
}

func (p *CreateBridgeParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["create"]
}

func (p *CreateBridgeParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *CreateBridgeParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *CreateBridgeParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *CreateBridgeParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
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
	Id int64
}

// NewReadBridgeParam return new ReadBridgeParam
func NewReadBridgeParam() *ReadBridgeParam {
	return &ReadBridgeParam{}
}

// Validate checks current values in model
func (p *ReadBridgeParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Bridge"].Commands["read"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ReadBridgeParam) getResourceDef() *schema.Resource {
	return define.Resources["Bridge"]
}

func (p *ReadBridgeParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["read"]
}

func (p *ReadBridgeParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ReadBridgeParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ReadBridgeParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ReadBridgeParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ReadBridgeParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadBridgeParam) GetId() int64 {
	return p.Id
}

// UpdateBridgeParam is input parameters for the sacloud API
type UpdateBridgeParam struct {
	Description string
	Id          int64
	Name        string
}

// NewUpdateBridgeParam return new UpdateBridgeParam
func NewUpdateBridgeParam() *UpdateBridgeParam {
	return &UpdateBridgeParam{}
}

// Validate checks current values in model
func (p *UpdateBridgeParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["Bridge"].Commands["update"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
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
		validator := define.Resources["Bridge"].Commands["update"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Bridge"].Commands["update"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *UpdateBridgeParam) getResourceDef() *schema.Resource {
	return define.Resources["Bridge"]
}

func (p *UpdateBridgeParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["update"]
}

func (p *UpdateBridgeParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *UpdateBridgeParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *UpdateBridgeParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *UpdateBridgeParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *UpdateBridgeParam) SetDescription(v string) {
	p.Description = v
}

func (p *UpdateBridgeParam) GetDescription() string {
	return p.Description
}
func (p *UpdateBridgeParam) SetId(v int64) {
	p.Id = v
}

func (p *UpdateBridgeParam) GetId() int64 {
	return p.Id
}
func (p *UpdateBridgeParam) SetName(v string) {
	p.Name = v
}

func (p *UpdateBridgeParam) GetName() string {
	return p.Name
}

// DeleteBridgeParam is input parameters for the sacloud API
type DeleteBridgeParam struct {
	Id int64
}

// NewDeleteBridgeParam return new DeleteBridgeParam
func NewDeleteBridgeParam() *DeleteBridgeParam {
	return &DeleteBridgeParam{}
}

// Validate checks current values in model
func (p *DeleteBridgeParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Bridge"].Commands["delete"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *DeleteBridgeParam) getResourceDef() *schema.Resource {
	return define.Resources["Bridge"]
}

func (p *DeleteBridgeParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["delete"]
}

func (p *DeleteBridgeParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *DeleteBridgeParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *DeleteBridgeParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *DeleteBridgeParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *DeleteBridgeParam) SetId(v int64) {
	p.Id = v
}

func (p *DeleteBridgeParam) GetId() int64 {
	return p.Id
}
