// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package command

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// CreateInterfaceParam is input parameters for the sacloud API
type CreateInterfaceParam struct {
	ServerId int64
}

// NewCreateInterfaceParam return new CreateInterfaceParam
func NewCreateInterfaceParam() *CreateInterfaceParam {
	return &CreateInterfaceParam{}
}

// Validate checks current values in model
func (p *CreateInterfaceParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--server-id", p.ServerId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Interface"].Commands["create"].Params["server-id"].ValidateFunc
		errs := validator("--server-id", p.ServerId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *CreateInterfaceParam) getResourceDef() *schema.Resource {
	return define.Resources["Interface"]
}

func (p *CreateInterfaceParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["create"]
}

func (p *CreateInterfaceParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *CreateInterfaceParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *CreateInterfaceParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *CreateInterfaceParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *CreateInterfaceParam) SetServerId(v int64) {
	p.ServerId = v
}

func (p *CreateInterfaceParam) GetServerId() int64 {
	return p.ServerId
}

// DeleteInterfaceParam is input parameters for the sacloud API
type DeleteInterfaceParam struct {
	Id int64
}

// NewDeleteInterfaceParam return new DeleteInterfaceParam
func NewDeleteInterfaceParam() *DeleteInterfaceParam {
	return &DeleteInterfaceParam{}
}

// Validate checks current values in model
func (p *DeleteInterfaceParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Interface"].Commands["delete"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *DeleteInterfaceParam) getResourceDef() *schema.Resource {
	return define.Resources["Interface"]
}

func (p *DeleteInterfaceParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["delete"]
}

func (p *DeleteInterfaceParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *DeleteInterfaceParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *DeleteInterfaceParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *DeleteInterfaceParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *DeleteInterfaceParam) SetId(v int64) {
	p.Id = v
}

func (p *DeleteInterfaceParam) GetId() int64 {
	return p.Id
}

// ListInterfaceParam is input parameters for the sacloud API
type ListInterfaceParam struct {
	From int
	Id   []int64
	Max  int
	Name []string
	Sort []string
}

// NewListInterfaceParam return new ListInterfaceParam
func NewListInterfaceParam() *ListInterfaceParam {
	return &ListInterfaceParam{}
}

// Validate checks current values in model
func (p *ListInterfaceParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["Interface"].Commands["list"].Params["id"].ValidateFunc
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
		errs := validateConflicts("--name", p.Name, map[string]interface{}{

			"--id": p.Id,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ListInterfaceParam) getResourceDef() *schema.Resource {
	return define.Resources["Interface"]
}

func (p *ListInterfaceParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["list"]
}

func (p *ListInterfaceParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ListInterfaceParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ListInterfaceParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ListInterfaceParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ListInterfaceParam) SetFrom(v int) {
	p.From = v
}

func (p *ListInterfaceParam) GetFrom() int {
	return p.From
}
func (p *ListInterfaceParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListInterfaceParam) GetId() []int64 {
	return p.Id
}
func (p *ListInterfaceParam) SetMax(v int) {
	p.Max = v
}

func (p *ListInterfaceParam) GetMax() int {
	return p.Max
}
func (p *ListInterfaceParam) SetName(v []string) {
	p.Name = v
}

func (p *ListInterfaceParam) GetName() []string {
	return p.Name
}
func (p *ListInterfaceParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListInterfaceParam) GetSort() []string {
	return p.Sort
}

// PacketFilterConnectInterfaceParam is input parameters for the sacloud API
type PacketFilterConnectInterfaceParam struct {
	Id             int64
	PacketFilterId int64
}

// NewPacketFilterConnectInterfaceParam return new PacketFilterConnectInterfaceParam
func NewPacketFilterConnectInterfaceParam() *PacketFilterConnectInterfaceParam {
	return &PacketFilterConnectInterfaceParam{}
}

// Validate checks current values in model
func (p *PacketFilterConnectInterfaceParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Interface"].Commands["packet-filter-connect"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--packet-filter-id", p.PacketFilterId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Interface"].Commands["packet-filter-connect"].Params["packet-filter-id"].ValidateFunc
		errs := validator("--packet-filter-id", p.PacketFilterId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *PacketFilterConnectInterfaceParam) getResourceDef() *schema.Resource {
	return define.Resources["Interface"]
}

func (p *PacketFilterConnectInterfaceParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["packet-filter-connect"]
}

func (p *PacketFilterConnectInterfaceParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *PacketFilterConnectInterfaceParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *PacketFilterConnectInterfaceParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *PacketFilterConnectInterfaceParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *PacketFilterConnectInterfaceParam) SetId(v int64) {
	p.Id = v
}

func (p *PacketFilterConnectInterfaceParam) GetId() int64 {
	return p.Id
}
func (p *PacketFilterConnectInterfaceParam) SetPacketFilterId(v int64) {
	p.PacketFilterId = v
}

func (p *PacketFilterConnectInterfaceParam) GetPacketFilterId() int64 {
	return p.PacketFilterId
}

// PacketFilterDisconnectInterfaceParam is input parameters for the sacloud API
type PacketFilterDisconnectInterfaceParam struct {
	Id             int64
	PacketFilterId int64
}

// NewPacketFilterDisconnectInterfaceParam return new PacketFilterDisconnectInterfaceParam
func NewPacketFilterDisconnectInterfaceParam() *PacketFilterDisconnectInterfaceParam {
	return &PacketFilterDisconnectInterfaceParam{}
}

// Validate checks current values in model
func (p *PacketFilterDisconnectInterfaceParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Interface"].Commands["packet-filter-disconnect"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--packet-filter-id", p.PacketFilterId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Interface"].Commands["packet-filter-disconnect"].Params["packet-filter-id"].ValidateFunc
		errs := validator("--packet-filter-id", p.PacketFilterId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *PacketFilterDisconnectInterfaceParam) getResourceDef() *schema.Resource {
	return define.Resources["Interface"]
}

func (p *PacketFilterDisconnectInterfaceParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["packet-filter-disconnect"]
}

func (p *PacketFilterDisconnectInterfaceParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *PacketFilterDisconnectInterfaceParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *PacketFilterDisconnectInterfaceParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *PacketFilterDisconnectInterfaceParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *PacketFilterDisconnectInterfaceParam) SetId(v int64) {
	p.Id = v
}

func (p *PacketFilterDisconnectInterfaceParam) GetId() int64 {
	return p.Id
}
func (p *PacketFilterDisconnectInterfaceParam) SetPacketFilterId(v int64) {
	p.PacketFilterId = v
}

func (p *PacketFilterDisconnectInterfaceParam) GetPacketFilterId() int64 {
	return p.PacketFilterId
}

// ReadInterfaceParam is input parameters for the sacloud API
type ReadInterfaceParam struct {
	Id int64
}

// NewReadInterfaceParam return new ReadInterfaceParam
func NewReadInterfaceParam() *ReadInterfaceParam {
	return &ReadInterfaceParam{}
}

// Validate checks current values in model
func (p *ReadInterfaceParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Interface"].Commands["read"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ReadInterfaceParam) getResourceDef() *schema.Resource {
	return define.Resources["Interface"]
}

func (p *ReadInterfaceParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["read"]
}

func (p *ReadInterfaceParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ReadInterfaceParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ReadInterfaceParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ReadInterfaceParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ReadInterfaceParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadInterfaceParam) GetId() int64 {
	return p.Id
}

// UpdateInterfaceParam is input parameters for the sacloud API
type UpdateInterfaceParam struct {
	Id            int64
	UserIpaddress string
}

// NewUpdateInterfaceParam return new UpdateInterfaceParam
func NewUpdateInterfaceParam() *UpdateInterfaceParam {
	return &UpdateInterfaceParam{}
}

// Validate checks current values in model
func (p *UpdateInterfaceParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Interface"].Commands["update"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Interface"].Commands["update"].Params["user-ipaddress"].ValidateFunc
		errs := validator("--user-ipaddress", p.UserIpaddress)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *UpdateInterfaceParam) getResourceDef() *schema.Resource {
	return define.Resources["Interface"]
}

func (p *UpdateInterfaceParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["update"]
}

func (p *UpdateInterfaceParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *UpdateInterfaceParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *UpdateInterfaceParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *UpdateInterfaceParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *UpdateInterfaceParam) SetId(v int64) {
	p.Id = v
}

func (p *UpdateInterfaceParam) GetId() int64 {
	return p.Id
}
func (p *UpdateInterfaceParam) SetUserIpaddress(v string) {
	p.UserIpaddress = v
}

func (p *UpdateInterfaceParam) GetUserIpaddress() string {
	return p.UserIpaddress
}
