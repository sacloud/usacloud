// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package command

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// DeleteInternetParam is input parameters for the sacloud API
type DeleteInternetParam struct {
	Id int64
}

// NewDeleteInternetParam return new DeleteInternetParam
func NewDeleteInternetParam() *DeleteInternetParam {
	return &DeleteInternetParam{}
}

// Validate checks current values in model
func (p *DeleteInternetParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["delete"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *DeleteInternetParam) getResourceDef() *schema.Resource {
	return define.Resources["Internet"]
}

func (p *DeleteInternetParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["delete"]
}

func (p *DeleteInternetParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *DeleteInternetParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *DeleteInternetParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *DeleteInternetParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *DeleteInternetParam) SetId(v int64) {
	p.Id = v
}

func (p *DeleteInternetParam) GetId() int64 {
	return p.Id
}

// UpdateBandwidthInternetParam is input parameters for the sacloud API
type UpdateBandwidthInternetParam struct {
	Id        int64
	BandWidth int
}

// NewUpdateBandwidthInternetParam return new UpdateBandwidthInternetParam
func NewUpdateBandwidthInternetParam() *UpdateBandwidthInternetParam {
	return &UpdateBandwidthInternetParam{

		BandWidth: 100,
	}
}

// Validate checks current values in model
func (p *UpdateBandwidthInternetParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["update-bandwidth"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--band-width", p.BandWidth)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["update-bandwidth"].Params["band-width"].ValidateFunc
		errs := validator("--band-width", p.BandWidth)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *UpdateBandwidthInternetParam) getResourceDef() *schema.Resource {
	return define.Resources["Internet"]
}

func (p *UpdateBandwidthInternetParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["update-bandwidth"]
}

func (p *UpdateBandwidthInternetParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *UpdateBandwidthInternetParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *UpdateBandwidthInternetParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *UpdateBandwidthInternetParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *UpdateBandwidthInternetParam) SetId(v int64) {
	p.Id = v
}

func (p *UpdateBandwidthInternetParam) GetId() int64 {
	return p.Id
}
func (p *UpdateBandwidthInternetParam) SetBandWidth(v int) {
	p.BandWidth = v
}

func (p *UpdateBandwidthInternetParam) GetBandWidth() int {
	return p.BandWidth
}

// ListInternetParam is input parameters for the sacloud API
type ListInternetParam struct {
	From int
	Max  int
	Sort []string
	Name []string
	Id   []int64
}

// NewListInternetParam return new ListInternetParam
func NewListInternetParam() *ListInternetParam {
	return &ListInternetParam{}
}

// Validate checks current values in model
func (p *ListInternetParam) Validate() []error {
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
		validator := define.Resources["Internet"].Commands["list"].Params["id"].ValidateFunc
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

func (p *ListInternetParam) getResourceDef() *schema.Resource {
	return define.Resources["Internet"]
}

func (p *ListInternetParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["list"]
}

func (p *ListInternetParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ListInternetParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ListInternetParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ListInternetParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ListInternetParam) SetFrom(v int) {
	p.From = v
}

func (p *ListInternetParam) GetFrom() int {
	return p.From
}
func (p *ListInternetParam) SetMax(v int) {
	p.Max = v
}

func (p *ListInternetParam) GetMax() int {
	return p.Max
}
func (p *ListInternetParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListInternetParam) GetSort() []string {
	return p.Sort
}
func (p *ListInternetParam) SetName(v []string) {
	p.Name = v
}

func (p *ListInternetParam) GetName() []string {
	return p.Name
}
func (p *ListInternetParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListInternetParam) GetId() []int64 {
	return p.Id
}

// CreateInternetParam is input parameters for the sacloud API
type CreateInternetParam struct {
	Name        string
	Description string
	Tags        []string
	IconId      int64
	NwMasklen   int
}

// NewCreateInternetParam return new CreateInternetParam
func NewCreateInternetParam() *CreateInternetParam {
	return &CreateInternetParam{

		NwMasklen: 28,
	}
}

// Validate checks current values in model
func (p *CreateInternetParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["create"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["create"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["create"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["create"].Params["icon-id"].ValidateFunc
		errs := validator("--icon-id", p.IconId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--nw-masklen", p.NwMasklen)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["create"].Params["nw-masklen"].ValidateFunc
		errs := validator("--nw-masklen", p.NwMasklen)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *CreateInternetParam) getResourceDef() *schema.Resource {
	return define.Resources["Internet"]
}

func (p *CreateInternetParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["create"]
}

func (p *CreateInternetParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *CreateInternetParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *CreateInternetParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *CreateInternetParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *CreateInternetParam) SetName(v string) {
	p.Name = v
}

func (p *CreateInternetParam) GetName() string {
	return p.Name
}
func (p *CreateInternetParam) SetDescription(v string) {
	p.Description = v
}

func (p *CreateInternetParam) GetDescription() string {
	return p.Description
}
func (p *CreateInternetParam) SetTags(v []string) {
	p.Tags = v
}

func (p *CreateInternetParam) GetTags() []string {
	return p.Tags
}
func (p *CreateInternetParam) SetIconId(v int64) {
	p.IconId = v
}

func (p *CreateInternetParam) GetIconId() int64 {
	return p.IconId
}
func (p *CreateInternetParam) SetNwMasklen(v int) {
	p.NwMasklen = v
}

func (p *CreateInternetParam) GetNwMasklen() int {
	return p.NwMasklen
}

// ReadInternetParam is input parameters for the sacloud API
type ReadInternetParam struct {
	Id int64
}

// NewReadInternetParam return new ReadInternetParam
func NewReadInternetParam() *ReadInternetParam {
	return &ReadInternetParam{}
}

// Validate checks current values in model
func (p *ReadInternetParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["read"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ReadInternetParam) getResourceDef() *schema.Resource {
	return define.Resources["Internet"]
}

func (p *ReadInternetParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["read"]
}

func (p *ReadInternetParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ReadInternetParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ReadInternetParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ReadInternetParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ReadInternetParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadInternetParam) GetId() int64 {
	return p.Id
}

// UpdateInternetParam is input parameters for the sacloud API
type UpdateInternetParam struct {
	BandWidth   int
	Id          int64
	Name        string
	Description string
	Tags        []string
	IconId      int64
}

// NewUpdateInternetParam return new UpdateInternetParam
func NewUpdateInternetParam() *UpdateInternetParam {
	return &UpdateInternetParam{

		BandWidth: 100,
	}
}

// Validate checks current values in model
func (p *UpdateInternetParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--band-width", p.BandWidth)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["update"].Params["band-width"].ValidateFunc
		errs := validator("--band-width", p.BandWidth)
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
		validator := define.Resources["Internet"].Commands["update"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["update"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["update"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["update"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["update"].Params["icon-id"].ValidateFunc
		errs := validator("--icon-id", p.IconId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *UpdateInternetParam) getResourceDef() *schema.Resource {
	return define.Resources["Internet"]
}

func (p *UpdateInternetParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["update"]
}

func (p *UpdateInternetParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *UpdateInternetParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *UpdateInternetParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *UpdateInternetParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *UpdateInternetParam) SetBandWidth(v int) {
	p.BandWidth = v
}

func (p *UpdateInternetParam) GetBandWidth() int {
	return p.BandWidth
}
func (p *UpdateInternetParam) SetId(v int64) {
	p.Id = v
}

func (p *UpdateInternetParam) GetId() int64 {
	return p.Id
}
func (p *UpdateInternetParam) SetName(v string) {
	p.Name = v
}

func (p *UpdateInternetParam) GetName() string {
	return p.Name
}
func (p *UpdateInternetParam) SetDescription(v string) {
	p.Description = v
}

func (p *UpdateInternetParam) GetDescription() string {
	return p.Description
}
func (p *UpdateInternetParam) SetTags(v []string) {
	p.Tags = v
}

func (p *UpdateInternetParam) GetTags() []string {
	return p.Tags
}
func (p *UpdateInternetParam) SetIconId(v int64) {
	p.IconId = v
}

func (p *UpdateInternetParam) GetIconId() int64 {
	return p.IconId
}
