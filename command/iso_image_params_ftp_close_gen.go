// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package command

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// FtpCloseISOImageParam is input parameters for the sacloud API
type FtpCloseISOImageParam struct {
	Id int64
}

// NewFtpCloseISOImageParam return new FtpCloseISOImageParam
func NewFtpCloseISOImageParam() *FtpCloseISOImageParam {
	return &FtpCloseISOImageParam{}
}

// Validate checks current values in model
func (p *FtpCloseISOImageParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ISOImage"].Commands["ftp-close"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *FtpCloseISOImageParam) getResourceDef() *schema.Resource {
	return define.Resources["ISOImage"]
}

func (p *FtpCloseISOImageParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["ftp-close"]
}

func (p *FtpCloseISOImageParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *FtpCloseISOImageParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *FtpCloseISOImageParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *FtpCloseISOImageParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *FtpCloseISOImageParam) SetId(v int64) {
	p.Id = v
}

func (p *FtpCloseISOImageParam) GetId() int64 {
	return p.Id
}
