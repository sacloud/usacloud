// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package command

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// ListDiskParam is input parameters for the sacloud API
type ListDiskParam struct {
	Max   int
	Sort  []string
	Name  []string
	Scope string
	Id    []int64
	From  int
}

// NewListDiskParam return new ListDiskParam
func NewListDiskParam() *ListDiskParam {
	return &ListDiskParam{}
}

// Validate checks current values in model
func (p *ListDiskParam) Validate() []error {
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
		validator := define.Resources["Disk"].Commands["list"].Params["scope"].ValidateFunc
		errs := validator("--scope", p.Scope)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Disk"].Commands["list"].Params["id"].ValidateFunc
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

func (p *ListDiskParam) getResourceDef() *schema.Resource {
	return define.Resources["Disk"]
}

func (p *ListDiskParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["list"]
}

func (p *ListDiskParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ListDiskParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ListDiskParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ListDiskParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ListDiskParam) SetMax(v int) {
	p.Max = v
}

func (p *ListDiskParam) GetMax() int {
	return p.Max
}
func (p *ListDiskParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListDiskParam) GetSort() []string {
	return p.Sort
}
func (p *ListDiskParam) SetName(v []string) {
	p.Name = v
}

func (p *ListDiskParam) GetName() []string {
	return p.Name
}
func (p *ListDiskParam) SetScope(v string) {
	p.Scope = v
}

func (p *ListDiskParam) GetScope() string {
	return p.Scope
}
func (p *ListDiskParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListDiskParam) GetId() []int64 {
	return p.Id
}
func (p *ListDiskParam) SetFrom(v int) {
	p.From = v
}

func (p *ListDiskParam) GetFrom() int {
	return p.From
}

// CreateDiskParam is input parameters for the sacloud API
type CreateDiskParam struct {
	SourceArchiveId int64
	DistantFrom     []int64
	Async           bool
	Name            string
	Tags            []string
	IconId          int64
	Plan            string
	Size            int
	Description     string
	Connection      string
	SourceDiskId    int64
}

// NewCreateDiskParam return new CreateDiskParam
func NewCreateDiskParam() *CreateDiskParam {
	return &CreateDiskParam{

		Plan: "ssd",

		Size: 20,

		Connection: "virtio",
	}
}

// Validate checks current values in model
func (p *CreateDiskParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["Disk"].Commands["create"].Params["source-archive-id"].ValidateFunc
		errs := validator("--source-archive-id", p.SourceArchiveId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateConflicts("--source-archive-id", p.SourceArchiveId, map[string]interface{}{

			"--source-disk-id": p.SourceDiskId,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Disk"].Commands["create"].Params["distant-from"].ValidateFunc
		errs := validator("--distant-from", p.DistantFrom)
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
		validator := define.Resources["Disk"].Commands["create"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Disk"].Commands["create"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Disk"].Commands["create"].Params["icon-id"].ValidateFunc
		errs := validator("--icon-id", p.IconId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--plan", p.Plan)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Disk"].Commands["create"].Params["plan"].ValidateFunc
		errs := validator("--plan", p.Plan)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--size", p.Size)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Disk"].Commands["create"].Params["size"].ValidateFunc
		errs := validator("--size", p.Size)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Disk"].Commands["create"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--connection", p.Connection)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Disk"].Commands["create"].Params["connection"].ValidateFunc
		errs := validator("--connection", p.Connection)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Disk"].Commands["create"].Params["source-disk-id"].ValidateFunc
		errs := validator("--source-disk-id", p.SourceDiskId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateConflicts("--source-disk-id", p.SourceDiskId, map[string]interface{}{

			"--source-archive-id": p.SourceArchiveId,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *CreateDiskParam) getResourceDef() *schema.Resource {
	return define.Resources["Disk"]
}

func (p *CreateDiskParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["create"]
}

func (p *CreateDiskParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *CreateDiskParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *CreateDiskParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *CreateDiskParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *CreateDiskParam) SetSourceArchiveId(v int64) {
	p.SourceArchiveId = v
}

func (p *CreateDiskParam) GetSourceArchiveId() int64 {
	return p.SourceArchiveId
}
func (p *CreateDiskParam) SetDistantFrom(v []int64) {
	p.DistantFrom = v
}

func (p *CreateDiskParam) GetDistantFrom() []int64 {
	return p.DistantFrom
}
func (p *CreateDiskParam) SetAsync(v bool) {
	p.Async = v
}

func (p *CreateDiskParam) GetAsync() bool {
	return p.Async
}
func (p *CreateDiskParam) SetName(v string) {
	p.Name = v
}

func (p *CreateDiskParam) GetName() string {
	return p.Name
}
func (p *CreateDiskParam) SetTags(v []string) {
	p.Tags = v
}

func (p *CreateDiskParam) GetTags() []string {
	return p.Tags
}
func (p *CreateDiskParam) SetIconId(v int64) {
	p.IconId = v
}

func (p *CreateDiskParam) GetIconId() int64 {
	return p.IconId
}
func (p *CreateDiskParam) SetPlan(v string) {
	p.Plan = v
}

func (p *CreateDiskParam) GetPlan() string {
	return p.Plan
}
func (p *CreateDiskParam) SetSize(v int) {
	p.Size = v
}

func (p *CreateDiskParam) GetSize() int {
	return p.Size
}
func (p *CreateDiskParam) SetDescription(v string) {
	p.Description = v
}

func (p *CreateDiskParam) GetDescription() string {
	return p.Description
}
func (p *CreateDiskParam) SetConnection(v string) {
	p.Connection = v
}

func (p *CreateDiskParam) GetConnection() string {
	return p.Connection
}
func (p *CreateDiskParam) SetSourceDiskId(v int64) {
	p.SourceDiskId = v
}

func (p *CreateDiskParam) GetSourceDiskId() int64 {
	return p.SourceDiskId
}

// ReadDiskParam is input parameters for the sacloud API
type ReadDiskParam struct {
	Id int64
}

// NewReadDiskParam return new ReadDiskParam
func NewReadDiskParam() *ReadDiskParam {
	return &ReadDiskParam{}
}

// Validate checks current values in model
func (p *ReadDiskParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Disk"].Commands["read"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ReadDiskParam) getResourceDef() *schema.Resource {
	return define.Resources["Disk"]
}

func (p *ReadDiskParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["read"]
}

func (p *ReadDiskParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ReadDiskParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ReadDiskParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ReadDiskParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ReadDiskParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadDiskParam) GetId() int64 {
	return p.Id
}

// UpdateDiskParam is input parameters for the sacloud API
type UpdateDiskParam struct {
	Tags        []string
	IconId      int64
	Connection  string
	Id          int64
	Name        string
	Description string
}

// NewUpdateDiskParam return new UpdateDiskParam
func NewUpdateDiskParam() *UpdateDiskParam {
	return &UpdateDiskParam{}
}

// Validate checks current values in model
func (p *UpdateDiskParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["Disk"].Commands["update"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Disk"].Commands["update"].Params["icon-id"].ValidateFunc
		errs := validator("--icon-id", p.IconId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Disk"].Commands["update"].Params["connection"].ValidateFunc
		errs := validator("--connection", p.Connection)
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
		validator := define.Resources["Disk"].Commands["update"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Disk"].Commands["update"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Disk"].Commands["update"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *UpdateDiskParam) getResourceDef() *schema.Resource {
	return define.Resources["Disk"]
}

func (p *UpdateDiskParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["update"]
}

func (p *UpdateDiskParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *UpdateDiskParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *UpdateDiskParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *UpdateDiskParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *UpdateDiskParam) SetTags(v []string) {
	p.Tags = v
}

func (p *UpdateDiskParam) GetTags() []string {
	return p.Tags
}
func (p *UpdateDiskParam) SetIconId(v int64) {
	p.IconId = v
}

func (p *UpdateDiskParam) GetIconId() int64 {
	return p.IconId
}
func (p *UpdateDiskParam) SetConnection(v string) {
	p.Connection = v
}

func (p *UpdateDiskParam) GetConnection() string {
	return p.Connection
}
func (p *UpdateDiskParam) SetId(v int64) {
	p.Id = v
}

func (p *UpdateDiskParam) GetId() int64 {
	return p.Id
}
func (p *UpdateDiskParam) SetName(v string) {
	p.Name = v
}

func (p *UpdateDiskParam) GetName() string {
	return p.Name
}
func (p *UpdateDiskParam) SetDescription(v string) {
	p.Description = v
}

func (p *UpdateDiskParam) GetDescription() string {
	return p.Description
}

// WaitForCopyDiskParam is input parameters for the sacloud API
type WaitForCopyDiskParam struct {
	Id int64
}

// NewWaitForCopyDiskParam return new WaitForCopyDiskParam
func NewWaitForCopyDiskParam() *WaitForCopyDiskParam {
	return &WaitForCopyDiskParam{}
}

// Validate checks current values in model
func (p *WaitForCopyDiskParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Disk"].Commands["wait-for-copy"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *WaitForCopyDiskParam) getResourceDef() *schema.Resource {
	return define.Resources["Disk"]
}

func (p *WaitForCopyDiskParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["wait-for-copy"]
}

func (p *WaitForCopyDiskParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *WaitForCopyDiskParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *WaitForCopyDiskParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *WaitForCopyDiskParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *WaitForCopyDiskParam) SetId(v int64) {
	p.Id = v
}

func (p *WaitForCopyDiskParam) GetId() int64 {
	return p.Id
}

// ReinstallFromArchiveDiskParam is input parameters for the sacloud API
type ReinstallFromArchiveDiskParam struct {
	SourceArchiveId int64
	DistantFrom     []int64
	Async           bool
	Id              int64
}

// NewReinstallFromArchiveDiskParam return new ReinstallFromArchiveDiskParam
func NewReinstallFromArchiveDiskParam() *ReinstallFromArchiveDiskParam {
	return &ReinstallFromArchiveDiskParam{}
}

// Validate checks current values in model
func (p *ReinstallFromArchiveDiskParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--source-archive-id", p.SourceArchiveId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Disk"].Commands["reinstall-from-archive"].Params["source-archive-id"].ValidateFunc
		errs := validator("--source-archive-id", p.SourceArchiveId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Disk"].Commands["reinstall-from-archive"].Params["distant-from"].ValidateFunc
		errs := validator("--distant-from", p.DistantFrom)
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
		validator := define.Resources["Disk"].Commands["reinstall-from-archive"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ReinstallFromArchiveDiskParam) getResourceDef() *schema.Resource {
	return define.Resources["Disk"]
}

func (p *ReinstallFromArchiveDiskParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["reinstall-from-archive"]
}

func (p *ReinstallFromArchiveDiskParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ReinstallFromArchiveDiskParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ReinstallFromArchiveDiskParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ReinstallFromArchiveDiskParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ReinstallFromArchiveDiskParam) SetSourceArchiveId(v int64) {
	p.SourceArchiveId = v
}

func (p *ReinstallFromArchiveDiskParam) GetSourceArchiveId() int64 {
	return p.SourceArchiveId
}
func (p *ReinstallFromArchiveDiskParam) SetDistantFrom(v []int64) {
	p.DistantFrom = v
}

func (p *ReinstallFromArchiveDiskParam) GetDistantFrom() []int64 {
	return p.DistantFrom
}
func (p *ReinstallFromArchiveDiskParam) SetAsync(v bool) {
	p.Async = v
}

func (p *ReinstallFromArchiveDiskParam) GetAsync() bool {
	return p.Async
}
func (p *ReinstallFromArchiveDiskParam) SetId(v int64) {
	p.Id = v
}

func (p *ReinstallFromArchiveDiskParam) GetId() int64 {
	return p.Id
}

// ServerDisconnectDiskParam is input parameters for the sacloud API
type ServerDisconnectDiskParam struct {
	Id int64
}

// NewServerDisconnectDiskParam return new ServerDisconnectDiskParam
func NewServerDisconnectDiskParam() *ServerDisconnectDiskParam {
	return &ServerDisconnectDiskParam{}
}

// Validate checks current values in model
func (p *ServerDisconnectDiskParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Disk"].Commands["server-disconnect"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ServerDisconnectDiskParam) getResourceDef() *schema.Resource {
	return define.Resources["Disk"]
}

func (p *ServerDisconnectDiskParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["server-disconnect"]
}

func (p *ServerDisconnectDiskParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ServerDisconnectDiskParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ServerDisconnectDiskParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ServerDisconnectDiskParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ServerDisconnectDiskParam) SetId(v int64) {
	p.Id = v
}

func (p *ServerDisconnectDiskParam) GetId() int64 {
	return p.Id
}

// DeleteDiskParam is input parameters for the sacloud API
type DeleteDiskParam struct {
	Id int64
}

// NewDeleteDiskParam return new DeleteDiskParam
func NewDeleteDiskParam() *DeleteDiskParam {
	return &DeleteDiskParam{}
}

// Validate checks current values in model
func (p *DeleteDiskParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Disk"].Commands["delete"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *DeleteDiskParam) getResourceDef() *schema.Resource {
	return define.Resources["Disk"]
}

func (p *DeleteDiskParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["delete"]
}

func (p *DeleteDiskParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *DeleteDiskParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *DeleteDiskParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *DeleteDiskParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *DeleteDiskParam) SetId(v int64) {
	p.Id = v
}

func (p *DeleteDiskParam) GetId() int64 {
	return p.Id
}

// EditDiskParam is input parameters for the sacloud API
type EditDiskParam struct {
	Id                  int64
	Hostname            string
	DisablePasswordAuth bool
	StartupScriptIds    []int64
	DefaultRoute        string
	Password            string
	SshKeyIds           []int64
	Ipaddress           string
	NwMasklen           int
}

// NewEditDiskParam return new EditDiskParam
func NewEditDiskParam() *EditDiskParam {
	return &EditDiskParam{

		NwMasklen: 24,
	}
}

// Validate checks current values in model
func (p *EditDiskParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Disk"].Commands["edit"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Disk"].Commands["edit"].Params["startup-script-ids"].ValidateFunc
		errs := validator("--startup-script-ids", p.StartupScriptIds)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Disk"].Commands["edit"].Params["ssh-key-ids"].ValidateFunc
		errs := validator("--ssh-key-ids", p.SshKeyIds)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Disk"].Commands["edit"].Params["nw-masklen"].ValidateFunc
		errs := validator("--nw-masklen", p.NwMasklen)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *EditDiskParam) getResourceDef() *schema.Resource {
	return define.Resources["Disk"]
}

func (p *EditDiskParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["edit"]
}

func (p *EditDiskParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *EditDiskParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *EditDiskParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *EditDiskParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *EditDiskParam) SetId(v int64) {
	p.Id = v
}

func (p *EditDiskParam) GetId() int64 {
	return p.Id
}
func (p *EditDiskParam) SetHostname(v string) {
	p.Hostname = v
}

func (p *EditDiskParam) GetHostname() string {
	return p.Hostname
}
func (p *EditDiskParam) SetDisablePasswordAuth(v bool) {
	p.DisablePasswordAuth = v
}

func (p *EditDiskParam) GetDisablePasswordAuth() bool {
	return p.DisablePasswordAuth
}
func (p *EditDiskParam) SetStartupScriptIds(v []int64) {
	p.StartupScriptIds = v
}

func (p *EditDiskParam) GetStartupScriptIds() []int64 {
	return p.StartupScriptIds
}
func (p *EditDiskParam) SetDefaultRoute(v string) {
	p.DefaultRoute = v
}

func (p *EditDiskParam) GetDefaultRoute() string {
	return p.DefaultRoute
}
func (p *EditDiskParam) SetPassword(v string) {
	p.Password = v
}

func (p *EditDiskParam) GetPassword() string {
	return p.Password
}
func (p *EditDiskParam) SetSshKeyIds(v []int64) {
	p.SshKeyIds = v
}

func (p *EditDiskParam) GetSshKeyIds() []int64 {
	return p.SshKeyIds
}
func (p *EditDiskParam) SetIpaddress(v string) {
	p.Ipaddress = v
}

func (p *EditDiskParam) GetIpaddress() string {
	return p.Ipaddress
}
func (p *EditDiskParam) SetNwMasklen(v int) {
	p.NwMasklen = v
}

func (p *EditDiskParam) GetNwMasklen() int {
	return p.NwMasklen
}

// ReinstallFromDiskDiskParam is input parameters for the sacloud API
type ReinstallFromDiskDiskParam struct {
	Id           int64
	SourceDiskId int64
	DistantFrom  []int64
	Async        bool
}

// NewReinstallFromDiskDiskParam return new ReinstallFromDiskDiskParam
func NewReinstallFromDiskDiskParam() *ReinstallFromDiskDiskParam {
	return &ReinstallFromDiskDiskParam{}
}

// Validate checks current values in model
func (p *ReinstallFromDiskDiskParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Disk"].Commands["reinstall-from-disk"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--source-disk-id", p.SourceDiskId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Disk"].Commands["reinstall-from-disk"].Params["source-disk-id"].ValidateFunc
		errs := validator("--source-disk-id", p.SourceDiskId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Disk"].Commands["reinstall-from-disk"].Params["distant-from"].ValidateFunc
		errs := validator("--distant-from", p.DistantFrom)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ReinstallFromDiskDiskParam) getResourceDef() *schema.Resource {
	return define.Resources["Disk"]
}

func (p *ReinstallFromDiskDiskParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["reinstall-from-disk"]
}

func (p *ReinstallFromDiskDiskParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ReinstallFromDiskDiskParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ReinstallFromDiskDiskParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ReinstallFromDiskDiskParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ReinstallFromDiskDiskParam) SetId(v int64) {
	p.Id = v
}

func (p *ReinstallFromDiskDiskParam) GetId() int64 {
	return p.Id
}
func (p *ReinstallFromDiskDiskParam) SetSourceDiskId(v int64) {
	p.SourceDiskId = v
}

func (p *ReinstallFromDiskDiskParam) GetSourceDiskId() int64 {
	return p.SourceDiskId
}
func (p *ReinstallFromDiskDiskParam) SetDistantFrom(v []int64) {
	p.DistantFrom = v
}

func (p *ReinstallFromDiskDiskParam) GetDistantFrom() []int64 {
	return p.DistantFrom
}
func (p *ReinstallFromDiskDiskParam) SetAsync(v bool) {
	p.Async = v
}

func (p *ReinstallFromDiskDiskParam) GetAsync() bool {
	return p.Async
}

// ReinstallToBlankDiskParam is input parameters for the sacloud API
type ReinstallToBlankDiskParam struct {
	Id          int64
	DistantFrom []int64
	Async       bool
}

// NewReinstallToBlankDiskParam return new ReinstallToBlankDiskParam
func NewReinstallToBlankDiskParam() *ReinstallToBlankDiskParam {
	return &ReinstallToBlankDiskParam{}
}

// Validate checks current values in model
func (p *ReinstallToBlankDiskParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Disk"].Commands["reinstall-to-blank"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Disk"].Commands["reinstall-to-blank"].Params["distant-from"].ValidateFunc
		errs := validator("--distant-from", p.DistantFrom)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ReinstallToBlankDiskParam) getResourceDef() *schema.Resource {
	return define.Resources["Disk"]
}

func (p *ReinstallToBlankDiskParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["reinstall-to-blank"]
}

func (p *ReinstallToBlankDiskParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ReinstallToBlankDiskParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ReinstallToBlankDiskParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ReinstallToBlankDiskParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ReinstallToBlankDiskParam) SetId(v int64) {
	p.Id = v
}

func (p *ReinstallToBlankDiskParam) GetId() int64 {
	return p.Id
}
func (p *ReinstallToBlankDiskParam) SetDistantFrom(v []int64) {
	p.DistantFrom = v
}

func (p *ReinstallToBlankDiskParam) GetDistantFrom() []int64 {
	return p.DistantFrom
}
func (p *ReinstallToBlankDiskParam) SetAsync(v bool) {
	p.Async = v
}

func (p *ReinstallToBlankDiskParam) GetAsync() bool {
	return p.Async
}

// ServerConnectDiskParam is input parameters for the sacloud API
type ServerConnectDiskParam struct {
	Id       int64
	ServerId int64
}

// NewServerConnectDiskParam return new ServerConnectDiskParam
func NewServerConnectDiskParam() *ServerConnectDiskParam {
	return &ServerConnectDiskParam{}
}

// Validate checks current values in model
func (p *ServerConnectDiskParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Disk"].Commands["server-connect"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--server-id", p.ServerId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Disk"].Commands["server-connect"].Params["server-id"].ValidateFunc
		errs := validator("--server-id", p.ServerId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ServerConnectDiskParam) getResourceDef() *schema.Resource {
	return define.Resources["Disk"]
}

func (p *ServerConnectDiskParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["server-connect"]
}

func (p *ServerConnectDiskParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ServerConnectDiskParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ServerConnectDiskParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ServerConnectDiskParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ServerConnectDiskParam) SetId(v int64) {
	p.Id = v
}

func (p *ServerConnectDiskParam) GetId() int64 {
	return p.Id
}
func (p *ServerConnectDiskParam) SetServerId(v int64) {
	p.ServerId = v
}

func (p *ServerConnectDiskParam) GetServerId() int64 {
	return p.ServerId
}
