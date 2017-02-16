// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package command

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// CreateInternetParam is input parameters for the sacloud API
type CreateInternetParam struct {
	Icon        int64
	NwMasklen   int
	Name        string
	Description string
	Tags        []string
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
		validator := define.Resources["Internet"].Commands["create"].Params["icon"].ValidateFunc
		errs := validator("--icon", p.Icon)
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

func (p *CreateInternetParam) SetIcon(v int64) {
	p.Icon = v
}

func (p *CreateInternetParam) GetIcon() int64 {
	return p.Icon
}
func (p *CreateInternetParam) SetNwMasklen(v int) {
	p.NwMasklen = v
}

func (p *CreateInternetParam) GetNwMasklen() int {
	return p.NwMasklen
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
