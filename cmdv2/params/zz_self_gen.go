// Copyright 2017-2020 The Usacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package params

import (
	"io"

	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/pkg/utils"
	"github.com/sacloud/usacloud/schema"
)

// InfoSelfParam is input parameters for the sacloud API
type InfoSelfParam struct {
	ParamTemplate     string
	Parameters        string
	ParamTemplateFile string
	ParameterFile     string
	GenerateSkeleton  bool

	input Input
}

// NewInfoSelfParam return new InfoSelfParam
func NewInfoSelfParam() *InfoSelfParam {
	return &InfoSelfParam{}
}

// Initialize init InfoSelfParam
func (p *InfoSelfParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *InfoSelfParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *InfoSelfParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if utils.IsEmpty(p.Parameters) {
		p.Parameters = ""
	}
	if utils.IsEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if utils.IsEmpty(p.ParameterFile) {
		p.ParameterFile = ""
	}
	if utils.IsEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}

}

func (p *InfoSelfParam) validate() error {
	var errors []error

	{
		errs := validateParameterOptions(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	return utils.FlattenErrors(errors)
}

func (p *InfoSelfParam) ResourceDef() *schema.Resource {
	return define.Resources["Self"]
}

func (p *InfoSelfParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["info"]
}

func (p *InfoSelfParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *InfoSelfParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *InfoSelfParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *InfoSelfParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *InfoSelfParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *InfoSelfParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *InfoSelfParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *InfoSelfParam) GetParameters() string {
	return p.Parameters
}
func (p *InfoSelfParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *InfoSelfParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *InfoSelfParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *InfoSelfParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *InfoSelfParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *InfoSelfParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
