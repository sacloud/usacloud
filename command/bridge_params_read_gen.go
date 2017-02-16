// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package command

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

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
