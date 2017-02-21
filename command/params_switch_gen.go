// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package command

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// BridgeConnectSwitchParam is input parameters for the sacloud API
type BridgeConnectSwitchParam struct {
	Id       int64
	BridgeId int64
}

// NewBridgeConnectSwitchParam return new BridgeConnectSwitchParam
func NewBridgeConnectSwitchParam() *BridgeConnectSwitchParam {
	return &BridgeConnectSwitchParam{}
}

// Validate checks current values in model
func (p *BridgeConnectSwitchParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Switch"].Commands["bridge-connect"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--bridge-id", p.BridgeId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Switch"].Commands["bridge-connect"].Params["bridge-id"].ValidateFunc
		errs := validator("--bridge-id", p.BridgeId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *BridgeConnectSwitchParam) getResourceDef() *schema.Resource {
	return define.Resources["Switch"]
}

func (p *BridgeConnectSwitchParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["bridge-connect"]
}

func (p *BridgeConnectSwitchParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *BridgeConnectSwitchParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *BridgeConnectSwitchParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *BridgeConnectSwitchParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *BridgeConnectSwitchParam) SetId(v int64) {
	p.Id = v
}

func (p *BridgeConnectSwitchParam) GetId() int64 {
	return p.Id
}
func (p *BridgeConnectSwitchParam) SetBridgeId(v int64) {
	p.BridgeId = v
}

func (p *BridgeConnectSwitchParam) GetBridgeId() int64 {
	return p.BridgeId
}

// BridgeDisconnectSwitchParam is input parameters for the sacloud API
type BridgeDisconnectSwitchParam struct {
	Id int64
}

// NewBridgeDisconnectSwitchParam return new BridgeDisconnectSwitchParam
func NewBridgeDisconnectSwitchParam() *BridgeDisconnectSwitchParam {
	return &BridgeDisconnectSwitchParam{}
}

// Validate checks current values in model
func (p *BridgeDisconnectSwitchParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Switch"].Commands["bridge-disconnect"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *BridgeDisconnectSwitchParam) getResourceDef() *schema.Resource {
	return define.Resources["Switch"]
}

func (p *BridgeDisconnectSwitchParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["bridge-disconnect"]
}

func (p *BridgeDisconnectSwitchParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *BridgeDisconnectSwitchParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *BridgeDisconnectSwitchParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *BridgeDisconnectSwitchParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *BridgeDisconnectSwitchParam) SetId(v int64) {
	p.Id = v
}

func (p *BridgeDisconnectSwitchParam) GetId() int64 {
	return p.Id
}

// ListSwitchParam is input parameters for the sacloud API
type ListSwitchParam struct {
	Name []string
	Id   []int64
	From int
	Max  int
	Sort []string
}

// NewListSwitchParam return new ListSwitchParam
func NewListSwitchParam() *ListSwitchParam {
	return &ListSwitchParam{}
}

// Validate checks current values in model
func (p *ListSwitchParam) Validate() []error {
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
		validator := define.Resources["Switch"].Commands["list"].Params["id"].ValidateFunc
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

func (p *ListSwitchParam) getResourceDef() *schema.Resource {
	return define.Resources["Switch"]
}

func (p *ListSwitchParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["list"]
}

func (p *ListSwitchParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ListSwitchParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ListSwitchParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ListSwitchParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ListSwitchParam) SetName(v []string) {
	p.Name = v
}

func (p *ListSwitchParam) GetName() []string {
	return p.Name
}
func (p *ListSwitchParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListSwitchParam) GetId() []int64 {
	return p.Id
}
func (p *ListSwitchParam) SetFrom(v int) {
	p.From = v
}

func (p *ListSwitchParam) GetFrom() int {
	return p.From
}
func (p *ListSwitchParam) SetMax(v int) {
	p.Max = v
}

func (p *ListSwitchParam) GetMax() int {
	return p.Max
}
func (p *ListSwitchParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListSwitchParam) GetSort() []string {
	return p.Sort
}

// CreateSwitchParam is input parameters for the sacloud API
type CreateSwitchParam struct {
	Name        string
	Description string
	Tags        []string
	IconId      int64
}

// NewCreateSwitchParam return new CreateSwitchParam
func NewCreateSwitchParam() *CreateSwitchParam {
	return &CreateSwitchParam{}
}

// Validate checks current values in model
func (p *CreateSwitchParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Switch"].Commands["create"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Switch"].Commands["create"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Switch"].Commands["create"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Switch"].Commands["create"].Params["icon-id"].ValidateFunc
		errs := validator("--icon-id", p.IconId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *CreateSwitchParam) getResourceDef() *schema.Resource {
	return define.Resources["Switch"]
}

func (p *CreateSwitchParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["create"]
}

func (p *CreateSwitchParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *CreateSwitchParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *CreateSwitchParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *CreateSwitchParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *CreateSwitchParam) SetName(v string) {
	p.Name = v
}

func (p *CreateSwitchParam) GetName() string {
	return p.Name
}
func (p *CreateSwitchParam) SetDescription(v string) {
	p.Description = v
}

func (p *CreateSwitchParam) GetDescription() string {
	return p.Description
}
func (p *CreateSwitchParam) SetTags(v []string) {
	p.Tags = v
}

func (p *CreateSwitchParam) GetTags() []string {
	return p.Tags
}
func (p *CreateSwitchParam) SetIconId(v int64) {
	p.IconId = v
}

func (p *CreateSwitchParam) GetIconId() int64 {
	return p.IconId
}

// ReadSwitchParam is input parameters for the sacloud API
type ReadSwitchParam struct {
	Id int64
}

// NewReadSwitchParam return new ReadSwitchParam
func NewReadSwitchParam() *ReadSwitchParam {
	return &ReadSwitchParam{}
}

// Validate checks current values in model
func (p *ReadSwitchParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Switch"].Commands["read"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ReadSwitchParam) getResourceDef() *schema.Resource {
	return define.Resources["Switch"]
}

func (p *ReadSwitchParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["read"]
}

func (p *ReadSwitchParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ReadSwitchParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ReadSwitchParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ReadSwitchParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ReadSwitchParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadSwitchParam) GetId() int64 {
	return p.Id
}

// UpdateSwitchParam is input parameters for the sacloud API
type UpdateSwitchParam struct {
	Id          int64
	Name        string
	Description string
	Tags        []string
	IconId      int64
}

// NewUpdateSwitchParam return new UpdateSwitchParam
func NewUpdateSwitchParam() *UpdateSwitchParam {
	return &UpdateSwitchParam{}
}

// Validate checks current values in model
func (p *UpdateSwitchParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Switch"].Commands["update"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Switch"].Commands["update"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Switch"].Commands["update"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Switch"].Commands["update"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Switch"].Commands["update"].Params["icon-id"].ValidateFunc
		errs := validator("--icon-id", p.IconId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *UpdateSwitchParam) getResourceDef() *schema.Resource {
	return define.Resources["Switch"]
}

func (p *UpdateSwitchParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["update"]
}

func (p *UpdateSwitchParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *UpdateSwitchParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *UpdateSwitchParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *UpdateSwitchParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *UpdateSwitchParam) SetId(v int64) {
	p.Id = v
}

func (p *UpdateSwitchParam) GetId() int64 {
	return p.Id
}
func (p *UpdateSwitchParam) SetName(v string) {
	p.Name = v
}

func (p *UpdateSwitchParam) GetName() string {
	return p.Name
}
func (p *UpdateSwitchParam) SetDescription(v string) {
	p.Description = v
}

func (p *UpdateSwitchParam) GetDescription() string {
	return p.Description
}
func (p *UpdateSwitchParam) SetTags(v []string) {
	p.Tags = v
}

func (p *UpdateSwitchParam) GetTags() []string {
	return p.Tags
}
func (p *UpdateSwitchParam) SetIconId(v int64) {
	p.IconId = v
}

func (p *UpdateSwitchParam) GetIconId() int64 {
	return p.IconId
}

// DeleteSwitchParam is input parameters for the sacloud API
type DeleteSwitchParam struct {
	Id int64
}

// NewDeleteSwitchParam return new DeleteSwitchParam
func NewDeleteSwitchParam() *DeleteSwitchParam {
	return &DeleteSwitchParam{}
}

// Validate checks current values in model
func (p *DeleteSwitchParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Switch"].Commands["delete"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *DeleteSwitchParam) getResourceDef() *schema.Resource {
	return define.Resources["Switch"]
}

func (p *DeleteSwitchParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["delete"]
}

func (p *DeleteSwitchParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *DeleteSwitchParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *DeleteSwitchParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *DeleteSwitchParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *DeleteSwitchParam) SetId(v int64) {
	p.Id = v
}

func (p *DeleteSwitchParam) GetId() int64 {
	return p.Id
}
