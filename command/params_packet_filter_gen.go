// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package command

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// RuleListPacketFilterParam is input parameters for the sacloud API
type RuleListPacketFilterParam struct {
	Id int64
}

// NewRuleListPacketFilterParam return new RuleListPacketFilterParam
func NewRuleListPacketFilterParam() *RuleListPacketFilterParam {
	return &RuleListPacketFilterParam{}
}

// Validate checks current values in model
func (p *RuleListPacketFilterParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["rule-list"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *RuleListPacketFilterParam) getResourceDef() *schema.Resource {
	return define.Resources["PacketFilter"]
}

func (p *RuleListPacketFilterParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["rule-list"]
}

func (p *RuleListPacketFilterParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *RuleListPacketFilterParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *RuleListPacketFilterParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *RuleListPacketFilterParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *RuleListPacketFilterParam) SetId(v int64) {
	p.Id = v
}

func (p *RuleListPacketFilterParam) GetId() int64 {
	return p.Id
}

// RuleDeletePacketFilterParam is input parameters for the sacloud API
type RuleDeletePacketFilterParam struct {
	Id    int64
	Index int
}

// NewRuleDeletePacketFilterParam return new RuleDeletePacketFilterParam
func NewRuleDeletePacketFilterParam() *RuleDeletePacketFilterParam {
	return &RuleDeletePacketFilterParam{}
}

// Validate checks current values in model
func (p *RuleDeletePacketFilterParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["rule-delete"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--index", p.Index)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *RuleDeletePacketFilterParam) getResourceDef() *schema.Resource {
	return define.Resources["PacketFilter"]
}

func (p *RuleDeletePacketFilterParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["rule-delete"]
}

func (p *RuleDeletePacketFilterParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *RuleDeletePacketFilterParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *RuleDeletePacketFilterParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *RuleDeletePacketFilterParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *RuleDeletePacketFilterParam) SetId(v int64) {
	p.Id = v
}

func (p *RuleDeletePacketFilterParam) GetId() int64 {
	return p.Id
}
func (p *RuleDeletePacketFilterParam) SetIndex(v int) {
	p.Index = v
}

func (p *RuleDeletePacketFilterParam) GetIndex() int {
	return p.Index
}

// CreatePacketFilterParam is input parameters for the sacloud API
type CreatePacketFilterParam struct {
	Name        string
	Description string
}

// NewCreatePacketFilterParam return new CreatePacketFilterParam
func NewCreatePacketFilterParam() *CreatePacketFilterParam {
	return &CreatePacketFilterParam{}
}

// Validate checks current values in model
func (p *CreatePacketFilterParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["create"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["create"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *CreatePacketFilterParam) getResourceDef() *schema.Resource {
	return define.Resources["PacketFilter"]
}

func (p *CreatePacketFilterParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["create"]
}

func (p *CreatePacketFilterParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *CreatePacketFilterParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *CreatePacketFilterParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *CreatePacketFilterParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *CreatePacketFilterParam) SetName(v string) {
	p.Name = v
}

func (p *CreatePacketFilterParam) GetName() string {
	return p.Name
}
func (p *CreatePacketFilterParam) SetDescription(v string) {
	p.Description = v
}

func (p *CreatePacketFilterParam) GetDescription() string {
	return p.Description
}

// ReadPacketFilterParam is input parameters for the sacloud API
type ReadPacketFilterParam struct {
	Id int64
}

// NewReadPacketFilterParam return new ReadPacketFilterParam
func NewReadPacketFilterParam() *ReadPacketFilterParam {
	return &ReadPacketFilterParam{}
}

// Validate checks current values in model
func (p *ReadPacketFilterParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["read"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ReadPacketFilterParam) getResourceDef() *schema.Resource {
	return define.Resources["PacketFilter"]
}

func (p *ReadPacketFilterParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["read"]
}

func (p *ReadPacketFilterParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ReadPacketFilterParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ReadPacketFilterParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ReadPacketFilterParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ReadPacketFilterParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadPacketFilterParam) GetId() int64 {
	return p.Id
}

// DeletePacketFilterParam is input parameters for the sacloud API
type DeletePacketFilterParam struct {
	Id int64
}

// NewDeletePacketFilterParam return new DeletePacketFilterParam
func NewDeletePacketFilterParam() *DeletePacketFilterParam {
	return &DeletePacketFilterParam{}
}

// Validate checks current values in model
func (p *DeletePacketFilterParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["delete"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *DeletePacketFilterParam) getResourceDef() *schema.Resource {
	return define.Resources["PacketFilter"]
}

func (p *DeletePacketFilterParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["delete"]
}

func (p *DeletePacketFilterParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *DeletePacketFilterParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *DeletePacketFilterParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *DeletePacketFilterParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *DeletePacketFilterParam) SetId(v int64) {
	p.Id = v
}

func (p *DeletePacketFilterParam) GetId() int64 {
	return p.Id
}

// RuleAddPacketFilterParam is input parameters for the sacloud API
type RuleAddPacketFilterParam struct {
	Description     string
	Index           int
	SourceNetwork   string
	SourcePort      string
	DestinationPort string
	Action          string
	Id              int64
	Protocol        string
}

// NewRuleAddPacketFilterParam return new RuleAddPacketFilterParam
func NewRuleAddPacketFilterParam() *RuleAddPacketFilterParam {
	return &RuleAddPacketFilterParam{

		Index: 1,
	}
}

// Validate checks current values in model
func (p *RuleAddPacketFilterParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["PacketFilter"].Commands["rule-add"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["rule-add"].Params["source-network"].ValidateFunc
		errs := validator("--source-network", p.SourceNetwork)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["rule-add"].Params["source-port"].ValidateFunc
		errs := validator("--source-port", p.SourcePort)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["rule-add"].Params["destination-port"].ValidateFunc
		errs := validator("--destination-port", p.DestinationPort)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["rule-add"].Params["action"].ValidateFunc
		errs := validator("--action", p.Action)
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
		validator := define.Resources["PacketFilter"].Commands["rule-add"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["rule-add"].Params["protocol"].ValidateFunc
		errs := validator("--protocol", p.Protocol)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *RuleAddPacketFilterParam) getResourceDef() *schema.Resource {
	return define.Resources["PacketFilter"]
}

func (p *RuleAddPacketFilterParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["rule-add"]
}

func (p *RuleAddPacketFilterParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *RuleAddPacketFilterParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *RuleAddPacketFilterParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *RuleAddPacketFilterParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *RuleAddPacketFilterParam) SetDescription(v string) {
	p.Description = v
}

func (p *RuleAddPacketFilterParam) GetDescription() string {
	return p.Description
}
func (p *RuleAddPacketFilterParam) SetIndex(v int) {
	p.Index = v
}

func (p *RuleAddPacketFilterParam) GetIndex() int {
	return p.Index
}
func (p *RuleAddPacketFilterParam) SetSourceNetwork(v string) {
	p.SourceNetwork = v
}

func (p *RuleAddPacketFilterParam) GetSourceNetwork() string {
	return p.SourceNetwork
}
func (p *RuleAddPacketFilterParam) SetSourcePort(v string) {
	p.SourcePort = v
}

func (p *RuleAddPacketFilterParam) GetSourcePort() string {
	return p.SourcePort
}
func (p *RuleAddPacketFilterParam) SetDestinationPort(v string) {
	p.DestinationPort = v
}

func (p *RuleAddPacketFilterParam) GetDestinationPort() string {
	return p.DestinationPort
}
func (p *RuleAddPacketFilterParam) SetAction(v string) {
	p.Action = v
}

func (p *RuleAddPacketFilterParam) GetAction() string {
	return p.Action
}
func (p *RuleAddPacketFilterParam) SetId(v int64) {
	p.Id = v
}

func (p *RuleAddPacketFilterParam) GetId() int64 {
	return p.Id
}
func (p *RuleAddPacketFilterParam) SetProtocol(v string) {
	p.Protocol = v
}

func (p *RuleAddPacketFilterParam) GetProtocol() string {
	return p.Protocol
}

// RuleUpdatePacketFilterParam is input parameters for the sacloud API
type RuleUpdatePacketFilterParam struct {
	Index           int
	Protocol        string
	SourcePort      string
	DestinationPort string
	Action          string
	Id              int64
	SourceNetwork   string
	Description     string
}

// NewRuleUpdatePacketFilterParam return new RuleUpdatePacketFilterParam
func NewRuleUpdatePacketFilterParam() *RuleUpdatePacketFilterParam {
	return &RuleUpdatePacketFilterParam{}
}

// Validate checks current values in model
func (p *RuleUpdatePacketFilterParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--index", p.Index)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["rule-update"].Params["protocol"].ValidateFunc
		errs := validator("--protocol", p.Protocol)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["rule-update"].Params["source-port"].ValidateFunc
		errs := validator("--source-port", p.SourcePort)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["rule-update"].Params["destination-port"].ValidateFunc
		errs := validator("--destination-port", p.DestinationPort)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["rule-update"].Params["action"].ValidateFunc
		errs := validator("--action", p.Action)
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
		validator := define.Resources["PacketFilter"].Commands["rule-update"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["rule-update"].Params["source-network"].ValidateFunc
		errs := validator("--source-network", p.SourceNetwork)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["rule-update"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *RuleUpdatePacketFilterParam) getResourceDef() *schema.Resource {
	return define.Resources["PacketFilter"]
}

func (p *RuleUpdatePacketFilterParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["rule-update"]
}

func (p *RuleUpdatePacketFilterParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *RuleUpdatePacketFilterParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *RuleUpdatePacketFilterParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *RuleUpdatePacketFilterParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *RuleUpdatePacketFilterParam) SetIndex(v int) {
	p.Index = v
}

func (p *RuleUpdatePacketFilterParam) GetIndex() int {
	return p.Index
}
func (p *RuleUpdatePacketFilterParam) SetProtocol(v string) {
	p.Protocol = v
}

func (p *RuleUpdatePacketFilterParam) GetProtocol() string {
	return p.Protocol
}
func (p *RuleUpdatePacketFilterParam) SetSourcePort(v string) {
	p.SourcePort = v
}

func (p *RuleUpdatePacketFilterParam) GetSourcePort() string {
	return p.SourcePort
}
func (p *RuleUpdatePacketFilterParam) SetDestinationPort(v string) {
	p.DestinationPort = v
}

func (p *RuleUpdatePacketFilterParam) GetDestinationPort() string {
	return p.DestinationPort
}
func (p *RuleUpdatePacketFilterParam) SetAction(v string) {
	p.Action = v
}

func (p *RuleUpdatePacketFilterParam) GetAction() string {
	return p.Action
}
func (p *RuleUpdatePacketFilterParam) SetId(v int64) {
	p.Id = v
}

func (p *RuleUpdatePacketFilterParam) GetId() int64 {
	return p.Id
}
func (p *RuleUpdatePacketFilterParam) SetSourceNetwork(v string) {
	p.SourceNetwork = v
}

func (p *RuleUpdatePacketFilterParam) GetSourceNetwork() string {
	return p.SourceNetwork
}
func (p *RuleUpdatePacketFilterParam) SetDescription(v string) {
	p.Description = v
}

func (p *RuleUpdatePacketFilterParam) GetDescription() string {
	return p.Description
}

// InterfaceConnectPacketFilterParam is input parameters for the sacloud API
type InterfaceConnectPacketFilterParam struct {
	Id          int64
	InterfaceId int64
}

// NewInterfaceConnectPacketFilterParam return new InterfaceConnectPacketFilterParam
func NewInterfaceConnectPacketFilterParam() *InterfaceConnectPacketFilterParam {
	return &InterfaceConnectPacketFilterParam{}
}

// Validate checks current values in model
func (p *InterfaceConnectPacketFilterParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["interface-connect"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--interface-id", p.InterfaceId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["interface-connect"].Params["interface-id"].ValidateFunc
		errs := validator("--interface-id", p.InterfaceId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *InterfaceConnectPacketFilterParam) getResourceDef() *schema.Resource {
	return define.Resources["PacketFilter"]
}

func (p *InterfaceConnectPacketFilterParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["interface-connect"]
}

func (p *InterfaceConnectPacketFilterParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *InterfaceConnectPacketFilterParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *InterfaceConnectPacketFilterParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *InterfaceConnectPacketFilterParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *InterfaceConnectPacketFilterParam) SetId(v int64) {
	p.Id = v
}

func (p *InterfaceConnectPacketFilterParam) GetId() int64 {
	return p.Id
}
func (p *InterfaceConnectPacketFilterParam) SetInterfaceId(v int64) {
	p.InterfaceId = v
}

func (p *InterfaceConnectPacketFilterParam) GetInterfaceId() int64 {
	return p.InterfaceId
}

// InterfaceDisconnectPacketFilterParam is input parameters for the sacloud API
type InterfaceDisconnectPacketFilterParam struct {
	Id          int64
	InterfaceId int64
}

// NewInterfaceDisconnectPacketFilterParam return new InterfaceDisconnectPacketFilterParam
func NewInterfaceDisconnectPacketFilterParam() *InterfaceDisconnectPacketFilterParam {
	return &InterfaceDisconnectPacketFilterParam{}
}

// Validate checks current values in model
func (p *InterfaceDisconnectPacketFilterParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["interface-disconnect"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--interface-id", p.InterfaceId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["interface-disconnect"].Params["interface-id"].ValidateFunc
		errs := validator("--interface-id", p.InterfaceId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *InterfaceDisconnectPacketFilterParam) getResourceDef() *schema.Resource {
	return define.Resources["PacketFilter"]
}

func (p *InterfaceDisconnectPacketFilterParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["interface-disconnect"]
}

func (p *InterfaceDisconnectPacketFilterParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *InterfaceDisconnectPacketFilterParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *InterfaceDisconnectPacketFilterParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *InterfaceDisconnectPacketFilterParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *InterfaceDisconnectPacketFilterParam) SetId(v int64) {
	p.Id = v
}

func (p *InterfaceDisconnectPacketFilterParam) GetId() int64 {
	return p.Id
}
func (p *InterfaceDisconnectPacketFilterParam) SetInterfaceId(v int64) {
	p.InterfaceId = v
}

func (p *InterfaceDisconnectPacketFilterParam) GetInterfaceId() int64 {
	return p.InterfaceId
}

// ListPacketFilterParam is input parameters for the sacloud API
type ListPacketFilterParam struct {
	Name []string
	Id   []int64
	From int
	Max  int
	Sort []string
}

// NewListPacketFilterParam return new ListPacketFilterParam
func NewListPacketFilterParam() *ListPacketFilterParam {
	return &ListPacketFilterParam{}
}

// Validate checks current values in model
func (p *ListPacketFilterParam) Validate() []error {
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
		validator := define.Resources["PacketFilter"].Commands["list"].Params["id"].ValidateFunc
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

func (p *ListPacketFilterParam) getResourceDef() *schema.Resource {
	return define.Resources["PacketFilter"]
}

func (p *ListPacketFilterParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["list"]
}

func (p *ListPacketFilterParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ListPacketFilterParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ListPacketFilterParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ListPacketFilterParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ListPacketFilterParam) SetName(v []string) {
	p.Name = v
}

func (p *ListPacketFilterParam) GetName() []string {
	return p.Name
}
func (p *ListPacketFilterParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListPacketFilterParam) GetId() []int64 {
	return p.Id
}
func (p *ListPacketFilterParam) SetFrom(v int) {
	p.From = v
}

func (p *ListPacketFilterParam) GetFrom() int {
	return p.From
}
func (p *ListPacketFilterParam) SetMax(v int) {
	p.Max = v
}

func (p *ListPacketFilterParam) GetMax() int {
	return p.Max
}
func (p *ListPacketFilterParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListPacketFilterParam) GetSort() []string {
	return p.Sort
}

// UpdatePacketFilterParam is input parameters for the sacloud API
type UpdatePacketFilterParam struct {
	Id          int64
	Name        string
	Description string
}

// NewUpdatePacketFilterParam return new UpdatePacketFilterParam
func NewUpdatePacketFilterParam() *UpdatePacketFilterParam {
	return &UpdatePacketFilterParam{}
}

// Validate checks current values in model
func (p *UpdatePacketFilterParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["update"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["update"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["PacketFilter"].Commands["update"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *UpdatePacketFilterParam) getResourceDef() *schema.Resource {
	return define.Resources["PacketFilter"]
}

func (p *UpdatePacketFilterParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["update"]
}

func (p *UpdatePacketFilterParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *UpdatePacketFilterParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *UpdatePacketFilterParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *UpdatePacketFilterParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *UpdatePacketFilterParam) SetId(v int64) {
	p.Id = v
}

func (p *UpdatePacketFilterParam) GetId() int64 {
	return p.Id
}
func (p *UpdatePacketFilterParam) SetName(v string) {
	p.Name = v
}

func (p *UpdatePacketFilterParam) GetName() string {
	return p.Name
}
func (p *UpdatePacketFilterParam) SetDescription(v string) {
	p.Description = v
}

func (p *UpdatePacketFilterParam) GetDescription() string {
	return p.Description
}
