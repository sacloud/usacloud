// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package command

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// DeleteSSHKeyParam is input parameters for the sacloud API
type DeleteSSHKeyParam struct {
	Id int64
}

// NewDeleteSSHKeyParam return new DeleteSSHKeyParam
func NewDeleteSSHKeyParam() *DeleteSSHKeyParam {
	return &DeleteSSHKeyParam{}
}

// Validate checks current values in model
func (p *DeleteSSHKeyParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["SSHKey"].Commands["delete"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *DeleteSSHKeyParam) getResourceDef() *schema.Resource {
	return define.Resources["SSHKey"]
}

func (p *DeleteSSHKeyParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["delete"]
}

func (p *DeleteSSHKeyParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *DeleteSSHKeyParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *DeleteSSHKeyParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *DeleteSSHKeyParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *DeleteSSHKeyParam) SetId(v int64) {
	p.Id = v
}

func (p *DeleteSSHKeyParam) GetId() int64 {
	return p.Id
}

// GenerateSSHKeyParam is input parameters for the sacloud API
type GenerateSSHKeyParam struct {
	Name             string
	Description      string
	PassPhrase       string
	PrivateKeyOutput string
}

// NewGenerateSSHKeyParam return new GenerateSSHKeyParam
func NewGenerateSSHKeyParam() *GenerateSSHKeyParam {
	return &GenerateSSHKeyParam{}
}

// Validate checks current values in model
func (p *GenerateSSHKeyParam) Validate() []error {
	errors := []error{}
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
	{
		validator := define.Resources["SSHKey"].Commands["generate"].Params["pass-phrase"].ValidateFunc
		errs := validator("--pass-phrase", p.PassPhrase)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *GenerateSSHKeyParam) getResourceDef() *schema.Resource {
	return define.Resources["SSHKey"]
}

func (p *GenerateSSHKeyParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["generate"]
}

func (p *GenerateSSHKeyParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *GenerateSSHKeyParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *GenerateSSHKeyParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *GenerateSSHKeyParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
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

// ListSSHKeyParam is input parameters for the sacloud API
type ListSSHKeyParam struct {
	Name []string
	Id   []int64
	From int
	Max  int
	Sort []string
}

// NewListSSHKeyParam return new ListSSHKeyParam
func NewListSSHKeyParam() *ListSSHKeyParam {
	return &ListSSHKeyParam{}
}

// Validate checks current values in model
func (p *ListSSHKeyParam) Validate() []error {
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
		validator := define.Resources["SSHKey"].Commands["list"].Params["id"].ValidateFunc
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

func (p *ListSSHKeyParam) getResourceDef() *schema.Resource {
	return define.Resources["SSHKey"]
}

func (p *ListSSHKeyParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["list"]
}

func (p *ListSSHKeyParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ListSSHKeyParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ListSSHKeyParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ListSSHKeyParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ListSSHKeyParam) SetName(v []string) {
	p.Name = v
}

func (p *ListSSHKeyParam) GetName() []string {
	return p.Name
}
func (p *ListSSHKeyParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListSSHKeyParam) GetId() []int64 {
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
func (p *ListSSHKeyParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListSSHKeyParam) GetSort() []string {
	return p.Sort
}

// CreateSSHKeyParam is input parameters for the sacloud API
type CreateSSHKeyParam struct {
	Description      string
	PublicKeyContent string
	PublicKey        string
	Name             string
}

// NewCreateSSHKeyParam return new CreateSSHKeyParam
func NewCreateSSHKeyParam() *CreateSSHKeyParam {
	return &CreateSSHKeyParam{}
}

// Validate checks current values in model
func (p *CreateSSHKeyParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["SSHKey"].Commands["create"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateConflicts("--public-key-content", p.PublicKeyContent, map[string]interface{}{

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

	return errors
}

func (p *CreateSSHKeyParam) getResourceDef() *schema.Resource {
	return define.Resources["SSHKey"]
}

func (p *CreateSSHKeyParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["create"]
}

func (p *CreateSSHKeyParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *CreateSSHKeyParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *CreateSSHKeyParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *CreateSSHKeyParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *CreateSSHKeyParam) SetDescription(v string) {
	p.Description = v
}

func (p *CreateSSHKeyParam) GetDescription() string {
	return p.Description
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

// ReadSSHKeyParam is input parameters for the sacloud API
type ReadSSHKeyParam struct {
	Id int64
}

// NewReadSSHKeyParam return new ReadSSHKeyParam
func NewReadSSHKeyParam() *ReadSSHKeyParam {
	return &ReadSSHKeyParam{}
}

// Validate checks current values in model
func (p *ReadSSHKeyParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["SSHKey"].Commands["read"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ReadSSHKeyParam) getResourceDef() *schema.Resource {
	return define.Resources["SSHKey"]
}

func (p *ReadSSHKeyParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["read"]
}

func (p *ReadSSHKeyParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ReadSSHKeyParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ReadSSHKeyParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ReadSSHKeyParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ReadSSHKeyParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadSSHKeyParam) GetId() int64 {
	return p.Id
}

// UpdateSSHKeyParam is input parameters for the sacloud API
type UpdateSSHKeyParam struct {
	Id          int64
	Name        string
	Description string
}

// NewUpdateSSHKeyParam return new UpdateSSHKeyParam
func NewUpdateSSHKeyParam() *UpdateSSHKeyParam {
	return &UpdateSSHKeyParam{}
}

// Validate checks current values in model
func (p *UpdateSSHKeyParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["SSHKey"].Commands["update"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
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

	return errors
}

func (p *UpdateSSHKeyParam) getResourceDef() *schema.Resource {
	return define.Resources["SSHKey"]
}

func (p *UpdateSSHKeyParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["update"]
}

func (p *UpdateSSHKeyParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *UpdateSSHKeyParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *UpdateSSHKeyParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *UpdateSSHKeyParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *UpdateSSHKeyParam) SetId(v int64) {
	p.Id = v
}

func (p *UpdateSSHKeyParam) GetId() int64 {
	return p.Id
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
