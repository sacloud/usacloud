// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package command

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// ListZoneParam is input parameters for the sacloud API
type ListZoneParam struct {
	Max  int
	Sort []string
	Name []string
	Id   []int64
	From int
}

// NewListZoneParam return new ListZoneParam
func NewListZoneParam() *ListZoneParam {
	return &ListZoneParam{}
}

// Validate checks current values in model
func (p *ListZoneParam) Validate() []error {
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
		validator := define.Resources["Zone"].Commands["list"].Params["id"].ValidateFunc
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

func (p *ListZoneParam) getResourceDef() *schema.Resource {
	return define.Resources["Zone"]
}

func (p *ListZoneParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["list"]
}

func (p *ListZoneParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ListZoneParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ListZoneParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ListZoneParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ListZoneParam) SetMax(v int) {
	p.Max = v
}

func (p *ListZoneParam) GetMax() int {
	return p.Max
}
func (p *ListZoneParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListZoneParam) GetSort() []string {
	return p.Sort
}
func (p *ListZoneParam) SetName(v []string) {
	p.Name = v
}

func (p *ListZoneParam) GetName() []string {
	return p.Name
}
func (p *ListZoneParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListZoneParam) GetId() []int64 {
	return p.Id
}
func (p *ListZoneParam) SetFrom(v int) {
	p.From = v
}

func (p *ListZoneParam) GetFrom() int {
	return p.From
}

// ReadZoneParam is input parameters for the sacloud API
type ReadZoneParam struct {
	Id int64
}

// NewReadZoneParam return new ReadZoneParam
func NewReadZoneParam() *ReadZoneParam {
	return &ReadZoneParam{}
}

// Validate checks current values in model
func (p *ReadZoneParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Zone"].Commands["read"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ReadZoneParam) getResourceDef() *schema.Resource {
	return define.Resources["Zone"]
}

func (p *ReadZoneParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["read"]
}

func (p *ReadZoneParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ReadZoneParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ReadZoneParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ReadZoneParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ReadZoneParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadZoneParam) GetId() int64 {
	return p.Id
}
