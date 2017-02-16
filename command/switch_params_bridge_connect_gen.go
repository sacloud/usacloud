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
