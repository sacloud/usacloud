// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package params

import (
	"io"

	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/pkg/utils"
	"github.com/sacloud/usacloud/schema"
)

// CurrentConfigParam is input parameters for the sacloud API
type CurrentConfigParam struct {
	ParamTemplate     string `json:"param-template"`
	Parameters        string `json:"parameters"`
	ParamTemplateFile string `json:"param-template-file"`
	ParameterFile     string `json:"parameter-file"`
	GenerateSkeleton  bool   `json:"generate-skeleton"`

	input Input
}

// NewCurrentConfigParam return new CurrentConfigParam
func NewCurrentConfigParam(in Input) (*CurrentConfigParam, error) {
	v := &CurrentConfigParam{
		input: in,
	}
	if err := v.validate(); err != nil {
		return nil, err
	}
	if err := loadParameters(v); err != nil {
		return nil, err
	}
	return v, nil
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *CurrentConfigParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *CurrentConfigParam) fillValueToSkeleton() {
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

func (p *CurrentConfigParam) validate() error {
	var errors []error

	return utils.FlattenErrors(errors)
}

func (p *CurrentConfigParam) ResourceDef() *schema.Resource {
	return define.Resources["Config"]
}

func (p *CurrentConfigParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["current"]
}

func (p *CurrentConfigParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *CurrentConfigParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *CurrentConfigParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *CurrentConfigParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *CurrentConfigParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *CurrentConfigParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *CurrentConfigParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *CurrentConfigParam) GetParameters() string {
	return p.Parameters
}
func (p *CurrentConfigParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *CurrentConfigParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *CurrentConfigParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *CurrentConfigParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *CurrentConfigParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *CurrentConfigParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}

// DeleteConfigParam is input parameters for the sacloud API
type DeleteConfigParam struct {
	Assumeyes         bool   `json:"assumeyes"`
	ParamTemplate     string `json:"param-template"`
	Parameters        string `json:"parameters"`
	ParamTemplateFile string `json:"param-template-file"`
	ParameterFile     string `json:"parameter-file"`
	GenerateSkeleton  bool   `json:"generate-skeleton"`

	input Input
}

// NewDeleteConfigParam return new DeleteConfigParam
func NewDeleteConfigParam(in Input) (*DeleteConfigParam, error) {
	v := &DeleteConfigParam{
		input: in,
	}
	if err := v.validate(); err != nil {
		return nil, err
	}
	if err := loadParameters(v); err != nil {
		return nil, err
	}
	return v, nil
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *DeleteConfigParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *DeleteConfigParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.Assumeyes) {
		p.Assumeyes = false
	}
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

func (p *DeleteConfigParam) validate() error {
	var errors []error

	return utils.FlattenErrors(errors)
}

func (p *DeleteConfigParam) ResourceDef() *schema.Resource {
	return define.Resources["Config"]
}

func (p *DeleteConfigParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["delete"]
}

func (p *DeleteConfigParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *DeleteConfigParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *DeleteConfigParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *DeleteConfigParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *DeleteConfigParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *DeleteConfigParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *DeleteConfigParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *DeleteConfigParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *DeleteConfigParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *DeleteConfigParam) GetParameters() string {
	return p.Parameters
}
func (p *DeleteConfigParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *DeleteConfigParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *DeleteConfigParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *DeleteConfigParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *DeleteConfigParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *DeleteConfigParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}

// EditConfigParam is input parameters for the sacloud API
type EditConfigParam struct {
	Token             string `json:"token"`
	Secret            string `json:"secret"`
	Zone              string `json:"zone"`
	DefaultOutputType string `json:"default-output-type"`
	ParamTemplate     string `json:"param-template"`
	Parameters        string `json:"parameters"`
	ParamTemplateFile string `json:"param-template-file"`
	ParameterFile     string `json:"parameter-file"`
	GenerateSkeleton  bool   `json:"generate-skeleton"`

	input Input
}

// NewEditConfigParam return new EditConfigParam
func NewEditConfigParam(in Input) (*EditConfigParam, error) {
	v := &EditConfigParam{
		input: in,
	}
	if err := v.validate(); err != nil {
		return nil, err
	}
	if err := loadParameters(v); err != nil {
		return nil, err
	}
	return v, nil
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *EditConfigParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *EditConfigParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.Token) {
		p.Token = ""
	}
	if utils.IsEmpty(p.Secret) {
		p.Secret = ""
	}
	if utils.IsEmpty(p.Zone) {
		p.Zone = ""
	}
	if utils.IsEmpty(p.DefaultOutputType) {
		p.DefaultOutputType = ""
	}
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

func (p *EditConfigParam) validate() error {
	var errors []error
	{
		validator := define.Resources["Config"].Commands["edit"].Params["zone"].ValidateFunc
		errs := validator("--zone", p.Zone)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Config"].Commands["edit"].Params["default-output-type"].ValidateFunc
		errs := validator("--default-output-type", p.DefaultOutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return utils.FlattenErrors(errors)
}

func (p *EditConfigParam) ResourceDef() *schema.Resource {
	return define.Resources["Config"]
}

func (p *EditConfigParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["edit"]
}

func (p *EditConfigParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *EditConfigParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *EditConfigParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *EditConfigParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *EditConfigParam) SetToken(v string) {
	p.Token = v
}

func (p *EditConfigParam) GetToken() string {
	return p.Token
}
func (p *EditConfigParam) SetSecret(v string) {
	p.Secret = v
}

func (p *EditConfigParam) GetSecret() string {
	return p.Secret
}
func (p *EditConfigParam) SetZone(v string) {
	p.Zone = v
}

func (p *EditConfigParam) GetZone() string {
	return p.Zone
}
func (p *EditConfigParam) SetDefaultOutputType(v string) {
	p.DefaultOutputType = v
}

func (p *EditConfigParam) GetDefaultOutputType() string {
	return p.DefaultOutputType
}
func (p *EditConfigParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *EditConfigParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *EditConfigParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *EditConfigParam) GetParameters() string {
	return p.Parameters
}
func (p *EditConfigParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *EditConfigParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *EditConfigParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *EditConfigParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *EditConfigParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *EditConfigParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}

// ListConfigParam is input parameters for the sacloud API
type ListConfigParam struct {
	ParamTemplate     string `json:"param-template"`
	Parameters        string `json:"parameters"`
	ParamTemplateFile string `json:"param-template-file"`
	ParameterFile     string `json:"parameter-file"`
	GenerateSkeleton  bool   `json:"generate-skeleton"`

	input Input
}

// NewListConfigParam return new ListConfigParam
func NewListConfigParam(in Input) (*ListConfigParam, error) {
	v := &ListConfigParam{
		input: in,
	}
	if err := v.validate(); err != nil {
		return nil, err
	}
	if err := loadParameters(v); err != nil {
		return nil, err
	}
	return v, nil
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ListConfigParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *ListConfigParam) fillValueToSkeleton() {
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

func (p *ListConfigParam) validate() error {
	var errors []error

	return utils.FlattenErrors(errors)
}

func (p *ListConfigParam) ResourceDef() *schema.Resource {
	return define.Resources["Config"]
}

func (p *ListConfigParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["list"]
}

func (p *ListConfigParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ListConfigParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ListConfigParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ListConfigParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ListConfigParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ListConfigParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ListConfigParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *ListConfigParam) GetParameters() string {
	return p.Parameters
}
func (p *ListConfigParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ListConfigParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ListConfigParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *ListConfigParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *ListConfigParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ListConfigParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}

// MigrateConfigParam is input parameters for the sacloud API
type MigrateConfigParam struct {
	ParamTemplate     string `json:"param-template"`
	Parameters        string `json:"parameters"`
	ParamTemplateFile string `json:"param-template-file"`
	ParameterFile     string `json:"parameter-file"`
	GenerateSkeleton  bool   `json:"generate-skeleton"`

	input Input
}

// NewMigrateConfigParam return new MigrateConfigParam
func NewMigrateConfigParam(in Input) (*MigrateConfigParam, error) {
	v := &MigrateConfigParam{
		input: in,
	}
	if err := v.validate(); err != nil {
		return nil, err
	}
	if err := loadParameters(v); err != nil {
		return nil, err
	}
	return v, nil
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *MigrateConfigParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *MigrateConfigParam) fillValueToSkeleton() {
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

func (p *MigrateConfigParam) validate() error {
	var errors []error

	return utils.FlattenErrors(errors)
}

func (p *MigrateConfigParam) ResourceDef() *schema.Resource {
	return define.Resources["Config"]
}

func (p *MigrateConfigParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["migrate"]
}

func (p *MigrateConfigParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *MigrateConfigParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *MigrateConfigParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *MigrateConfigParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *MigrateConfigParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *MigrateConfigParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *MigrateConfigParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *MigrateConfigParam) GetParameters() string {
	return p.Parameters
}
func (p *MigrateConfigParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *MigrateConfigParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *MigrateConfigParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *MigrateConfigParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *MigrateConfigParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *MigrateConfigParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}

// ShowConfigParam is input parameters for the sacloud API
type ShowConfigParam struct {
	ParamTemplate     string `json:"param-template"`
	Parameters        string `json:"parameters"`
	ParamTemplateFile string `json:"param-template-file"`
	ParameterFile     string `json:"parameter-file"`
	GenerateSkeleton  bool   `json:"generate-skeleton"`

	input Input
}

// NewShowConfigParam return new ShowConfigParam
func NewShowConfigParam(in Input) (*ShowConfigParam, error) {
	v := &ShowConfigParam{
		input: in,
	}
	if err := v.validate(); err != nil {
		return nil, err
	}
	if err := loadParameters(v); err != nil {
		return nil, err
	}
	return v, nil
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ShowConfigParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *ShowConfigParam) fillValueToSkeleton() {
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

func (p *ShowConfigParam) validate() error {
	var errors []error

	return utils.FlattenErrors(errors)
}

func (p *ShowConfigParam) ResourceDef() *schema.Resource {
	return define.Resources["Config"]
}

func (p *ShowConfigParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["show"]
}

func (p *ShowConfigParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ShowConfigParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ShowConfigParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ShowConfigParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ShowConfigParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ShowConfigParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ShowConfigParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *ShowConfigParam) GetParameters() string {
	return p.Parameters
}
func (p *ShowConfigParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ShowConfigParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ShowConfigParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *ShowConfigParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *ShowConfigParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ShowConfigParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}

// UseConfigParam is input parameters for the sacloud API
type UseConfigParam struct {
	ParamTemplate     string `json:"param-template"`
	Parameters        string `json:"parameters"`
	ParamTemplateFile string `json:"param-template-file"`
	ParameterFile     string `json:"parameter-file"`
	GenerateSkeleton  bool   `json:"generate-skeleton"`

	input Input
}

// NewUseConfigParam return new UseConfigParam
func NewUseConfigParam(in Input) (*UseConfigParam, error) {
	v := &UseConfigParam{
		input: in,
	}
	if err := v.validate(); err != nil {
		return nil, err
	}
	if err := loadParameters(v); err != nil {
		return nil, err
	}
	return v, nil
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *UseConfigParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *UseConfigParam) fillValueToSkeleton() {
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

func (p *UseConfigParam) validate() error {
	var errors []error

	return utils.FlattenErrors(errors)
}

func (p *UseConfigParam) ResourceDef() *schema.Resource {
	return define.Resources["Config"]
}

func (p *UseConfigParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["use"]
}

func (p *UseConfigParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *UseConfigParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *UseConfigParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *UseConfigParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *UseConfigParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *UseConfigParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *UseConfigParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *UseConfigParam) GetParameters() string {
	return p.Parameters
}
func (p *UseConfigParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *UseConfigParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *UseConfigParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *UseConfigParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *UseConfigParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *UseConfigParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
