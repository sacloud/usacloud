// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package command

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// ListInternetParam is input parameters for the sacloud API
type ListInternetParam struct {
	Max  int
	Sort []string
	Name []string
	Id   []int64
	From int
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
func (p *ListInternetParam) SetFrom(v int) {
	p.From = v
}

func (p *ListInternetParam) GetFrom() int {
	return p.From
}
