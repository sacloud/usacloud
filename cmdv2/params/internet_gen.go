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

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/pkg/utils"
	"github.com/sacloud/usacloud/pkg/validation"
	"github.com/sacloud/usacloud/schema"
)

// ListInternetParam is input parameters for the sacloud API
type ListInternetParam struct {
	Max  int
	Sort []string
	Name []string
	Id   []sacloud.ID
	From int
	Tags []string

	input Input
}

// NewListInternetParam return new ListInternetParam
func NewListInternetParam() *ListInternetParam {
	return &ListInternetParam{}
}

// Initialize init ListInternetParam
func (p *ListInternetParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ListInternetParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *ListInternetParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.Max) {
		p.Max = 0
	}
	if utils.IsEmpty(p.Sort) {
		p.Sort = []string{""}
	}
	if utils.IsEmpty(p.Name) {
		p.Name = []string{""}
	}
	if utils.IsEmpty(p.Id) {
		p.Id = []sacloud.ID{}
	}
	if utils.IsEmpty(p.From) {
		p.From = 0
	}
	if utils.IsEmpty(p.Tags) {
		p.Tags = []string{""}
	}

}

func (p *ListInternetParam) validate() error {
	var errors []error

	{
		errs := validation.ConflictsWith("--name", p.Name, map[string]interface{}{

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
		errs := validation.ConflictsWith("--id", p.Id, map[string]interface{}{

			"--name": p.Name,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["Internet"].Commands["list"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return utils.FlattenErrors(errors)
}

func (p *ListInternetParam) ResourceDef() *schema.Resource {
	return define.Resources["Internet"]
}

func (p *ListInternetParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["list"]
}

func (p *ListInternetParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ListInternetParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ListInternetParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ListInternetParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
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
func (p *ListInternetParam) SetId(v []sacloud.ID) {
	p.Id = v
}

func (p *ListInternetParam) GetId() []sacloud.ID {
	return p.Id
}
func (p *ListInternetParam) SetFrom(v int) {
	p.From = v
}

func (p *ListInternetParam) GetFrom() int {
	return p.From
}
func (p *ListInternetParam) SetTags(v []string) {
	p.Tags = v
}

func (p *ListInternetParam) GetTags() []string {
	return p.Tags
}

// CreateInternetParam is input parameters for the sacloud API
type CreateInternetParam struct {
	Description string
	Tags        []string
	IconId      sacloud.ID
	NwMasklen   int
	BandWidth   int
	Name        string

	input Input
}

// NewCreateInternetParam return new CreateInternetParam
func NewCreateInternetParam() *CreateInternetParam {
	return &CreateInternetParam{
		NwMasklen: 28, BandWidth: 100}
}

// Initialize init CreateInternetParam
func (p *CreateInternetParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *CreateInternetParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *CreateInternetParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.Description) {
		p.Description = ""
	}
	if utils.IsEmpty(p.Tags) {
		p.Tags = []string{""}
	}
	if utils.IsEmpty(p.IconId) {
		p.IconId = sacloud.ID(0)
	}
	if utils.IsEmpty(p.NwMasklen) {
		p.NwMasklen = 0
	}
	if utils.IsEmpty(p.BandWidth) {
		p.BandWidth = 0
	}
	if utils.IsEmpty(p.Name) {
		p.Name = ""
	}

}

func (p *CreateInternetParam) validate() error {
	var errors []error

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

	{
		validator := validateRequired
		errs := validator("--band-width", p.BandWidth)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["create"].Params["band-width"].ValidateFunc
		errs := validator("--band-width", p.BandWidth)
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
		validator := define.Resources["Internet"].Commands["create"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return utils.FlattenErrors(errors)
}

func (p *CreateInternetParam) ResourceDef() *schema.Resource {
	return define.Resources["Internet"]
}

func (p *CreateInternetParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["create"]
}

func (p *CreateInternetParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *CreateInternetParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *CreateInternetParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *CreateInternetParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
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
func (p *CreateInternetParam) SetIconId(v sacloud.ID) {
	p.IconId = v
}

func (p *CreateInternetParam) GetIconId() sacloud.ID {
	return p.IconId
}
func (p *CreateInternetParam) SetNwMasklen(v int) {
	p.NwMasklen = v
}

func (p *CreateInternetParam) GetNwMasklen() int {
	return p.NwMasklen
}
func (p *CreateInternetParam) SetBandWidth(v int) {
	p.BandWidth = v
}

func (p *CreateInternetParam) GetBandWidth() int {
	return p.BandWidth
}
func (p *CreateInternetParam) SetName(v string) {
	p.Name = v
}

func (p *CreateInternetParam) GetName() string {
	return p.Name
}

// ReadInternetParam is input parameters for the sacloud API
type ReadInternetParam struct {
	input Input
}

// NewReadInternetParam return new ReadInternetParam
func NewReadInternetParam() *ReadInternetParam {
	return &ReadInternetParam{}
}

// Initialize init ReadInternetParam
func (p *ReadInternetParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ReadInternetParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *ReadInternetParam) fillValueToSkeleton() {

}

func (p *ReadInternetParam) validate() error {
	var errors []error

	return utils.FlattenErrors(errors)
}

func (p *ReadInternetParam) ResourceDef() *schema.Resource {
	return define.Resources["Internet"]
}

func (p *ReadInternetParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["read"]
}

func (p *ReadInternetParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ReadInternetParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ReadInternetParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ReadInternetParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

// UpdateInternetParam is input parameters for the sacloud API
type UpdateInternetParam struct {
	IconId      sacloud.ID
	BandWidth   int
	Name        string
	Description string
	Tags        []string

	input Input
}

// NewUpdateInternetParam return new UpdateInternetParam
func NewUpdateInternetParam() *UpdateInternetParam {
	return &UpdateInternetParam{}
}

// Initialize init UpdateInternetParam
func (p *UpdateInternetParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *UpdateInternetParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *UpdateInternetParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.IconId) {
		p.IconId = sacloud.ID(0)
	}
	if utils.IsEmpty(p.BandWidth) {
		p.BandWidth = 0
	}
	if utils.IsEmpty(p.Name) {
		p.Name = ""
	}
	if utils.IsEmpty(p.Description) {
		p.Description = ""
	}
	if utils.IsEmpty(p.Tags) {
		p.Tags = []string{""}
	}

}

func (p *UpdateInternetParam) validate() error {
	var errors []error

	{
		validator := define.Resources["Internet"].Commands["update"].Params["icon-id"].ValidateFunc
		errs := validator("--icon-id", p.IconId)
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

	return utils.FlattenErrors(errors)
}

func (p *UpdateInternetParam) ResourceDef() *schema.Resource {
	return define.Resources["Internet"]
}

func (p *UpdateInternetParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["update"]
}

func (p *UpdateInternetParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *UpdateInternetParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *UpdateInternetParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *UpdateInternetParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *UpdateInternetParam) SetIconId(v sacloud.ID) {
	p.IconId = v
}

func (p *UpdateInternetParam) GetIconId() sacloud.ID {
	return p.IconId
}
func (p *UpdateInternetParam) SetBandWidth(v int) {
	p.BandWidth = v
}

func (p *UpdateInternetParam) GetBandWidth() int {
	return p.BandWidth
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

// DeleteInternetParam is input parameters for the sacloud API
type DeleteInternetParam struct {
	input Input
}

// NewDeleteInternetParam return new DeleteInternetParam
func NewDeleteInternetParam() *DeleteInternetParam {
	return &DeleteInternetParam{}
}

// Initialize init DeleteInternetParam
func (p *DeleteInternetParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *DeleteInternetParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *DeleteInternetParam) fillValueToSkeleton() {

}

func (p *DeleteInternetParam) validate() error {
	var errors []error

	return utils.FlattenErrors(errors)
}

func (p *DeleteInternetParam) ResourceDef() *schema.Resource {
	return define.Resources["Internet"]
}

func (p *DeleteInternetParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["delete"]
}

func (p *DeleteInternetParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *DeleteInternetParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *DeleteInternetParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *DeleteInternetParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

// UpdateBandwidthInternetParam is input parameters for the sacloud API
type UpdateBandwidthInternetParam struct {
	BandWidth int

	input Input
}

// NewUpdateBandwidthInternetParam return new UpdateBandwidthInternetParam
func NewUpdateBandwidthInternetParam() *UpdateBandwidthInternetParam {
	return &UpdateBandwidthInternetParam{
		BandWidth: 100}
}

// Initialize init UpdateBandwidthInternetParam
func (p *UpdateBandwidthInternetParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *UpdateBandwidthInternetParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *UpdateBandwidthInternetParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.BandWidth) {
		p.BandWidth = 0
	}

}

func (p *UpdateBandwidthInternetParam) validate() error {
	var errors []error

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

	return utils.FlattenErrors(errors)
}

func (p *UpdateBandwidthInternetParam) ResourceDef() *schema.Resource {
	return define.Resources["Internet"]
}

func (p *UpdateBandwidthInternetParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["update-bandwidth"]
}

func (p *UpdateBandwidthInternetParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *UpdateBandwidthInternetParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *UpdateBandwidthInternetParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *UpdateBandwidthInternetParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *UpdateBandwidthInternetParam) SetBandWidth(v int) {
	p.BandWidth = v
}

func (p *UpdateBandwidthInternetParam) GetBandWidth() int {
	return p.BandWidth
}

// SubnetInfoInternetParam is input parameters for the sacloud API
type SubnetInfoInternetParam struct {
	input Input
}

// NewSubnetInfoInternetParam return new SubnetInfoInternetParam
func NewSubnetInfoInternetParam() *SubnetInfoInternetParam {
	return &SubnetInfoInternetParam{}
}

// Initialize init SubnetInfoInternetParam
func (p *SubnetInfoInternetParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *SubnetInfoInternetParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *SubnetInfoInternetParam) fillValueToSkeleton() {

}

func (p *SubnetInfoInternetParam) validate() error {
	var errors []error

	return utils.FlattenErrors(errors)
}

func (p *SubnetInfoInternetParam) ResourceDef() *schema.Resource {
	return define.Resources["Internet"]
}

func (p *SubnetInfoInternetParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["subnet-info"]
}

func (p *SubnetInfoInternetParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *SubnetInfoInternetParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *SubnetInfoInternetParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *SubnetInfoInternetParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

// SubnetAddInternetParam is input parameters for the sacloud API
type SubnetAddInternetParam struct {
	NwMasklen int
	NextHop   string

	input Input
}

// NewSubnetAddInternetParam return new SubnetAddInternetParam
func NewSubnetAddInternetParam() *SubnetAddInternetParam {
	return &SubnetAddInternetParam{
		NwMasklen: 28}
}

// Initialize init SubnetAddInternetParam
func (p *SubnetAddInternetParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *SubnetAddInternetParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *SubnetAddInternetParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.NwMasklen) {
		p.NwMasklen = 0
	}
	if utils.IsEmpty(p.NextHop) {
		p.NextHop = ""
	}

}

func (p *SubnetAddInternetParam) validate() error {
	var errors []error

	{
		validator := validateRequired
		errs := validator("--nw-masklen", p.NwMasklen)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["subnet-add"].Params["nw-masklen"].ValidateFunc
		errs := validator("--nw-masklen", p.NwMasklen)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := validateRequired
		errs := validator("--next-hop", p.NextHop)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["subnet-add"].Params["next-hop"].ValidateFunc
		errs := validator("--next-hop", p.NextHop)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return utils.FlattenErrors(errors)
}

func (p *SubnetAddInternetParam) ResourceDef() *schema.Resource {
	return define.Resources["Internet"]
}

func (p *SubnetAddInternetParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["subnet-add"]
}

func (p *SubnetAddInternetParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *SubnetAddInternetParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *SubnetAddInternetParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *SubnetAddInternetParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *SubnetAddInternetParam) SetNwMasklen(v int) {
	p.NwMasklen = v
}

func (p *SubnetAddInternetParam) GetNwMasklen() int {
	return p.NwMasklen
}
func (p *SubnetAddInternetParam) SetNextHop(v string) {
	p.NextHop = v
}

func (p *SubnetAddInternetParam) GetNextHop() string {
	return p.NextHop
}

// SubnetDeleteInternetParam is input parameters for the sacloud API
type SubnetDeleteInternetParam struct {
	SubnetId sacloud.ID

	input Input
}

// NewSubnetDeleteInternetParam return new SubnetDeleteInternetParam
func NewSubnetDeleteInternetParam() *SubnetDeleteInternetParam {
	return &SubnetDeleteInternetParam{}
}

// Initialize init SubnetDeleteInternetParam
func (p *SubnetDeleteInternetParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *SubnetDeleteInternetParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *SubnetDeleteInternetParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.SubnetId) {
		p.SubnetId = sacloud.ID(0)
	}

}

func (p *SubnetDeleteInternetParam) validate() error {
	var errors []error

	{
		validator := define.Resources["Internet"].Commands["subnet-delete"].Params["subnet-id"].ValidateFunc
		errs := validator("--subnet-id", p.SubnetId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return utils.FlattenErrors(errors)
}

func (p *SubnetDeleteInternetParam) ResourceDef() *schema.Resource {
	return define.Resources["Internet"]
}

func (p *SubnetDeleteInternetParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["subnet-delete"]
}

func (p *SubnetDeleteInternetParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *SubnetDeleteInternetParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *SubnetDeleteInternetParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *SubnetDeleteInternetParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *SubnetDeleteInternetParam) SetSubnetId(v sacloud.ID) {
	p.SubnetId = v
}

func (p *SubnetDeleteInternetParam) GetSubnetId() sacloud.ID {
	return p.SubnetId
}

// SubnetUpdateInternetParam is input parameters for the sacloud API
type SubnetUpdateInternetParam struct {
	SubnetId sacloud.ID
	NextHop  string

	input Input
}

// NewSubnetUpdateInternetParam return new SubnetUpdateInternetParam
func NewSubnetUpdateInternetParam() *SubnetUpdateInternetParam {
	return &SubnetUpdateInternetParam{}
}

// Initialize init SubnetUpdateInternetParam
func (p *SubnetUpdateInternetParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *SubnetUpdateInternetParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *SubnetUpdateInternetParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.SubnetId) {
		p.SubnetId = sacloud.ID(0)
	}
	if utils.IsEmpty(p.NextHop) {
		p.NextHop = ""
	}

}

func (p *SubnetUpdateInternetParam) validate() error {
	var errors []error

	{
		validator := define.Resources["Internet"].Commands["subnet-update"].Params["subnet-id"].ValidateFunc
		errs := validator("--subnet-id", p.SubnetId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := validateRequired
		errs := validator("--next-hop", p.NextHop)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Internet"].Commands["subnet-update"].Params["next-hop"].ValidateFunc
		errs := validator("--next-hop", p.NextHop)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return utils.FlattenErrors(errors)
}

func (p *SubnetUpdateInternetParam) ResourceDef() *schema.Resource {
	return define.Resources["Internet"]
}

func (p *SubnetUpdateInternetParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["subnet-update"]
}

func (p *SubnetUpdateInternetParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *SubnetUpdateInternetParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *SubnetUpdateInternetParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *SubnetUpdateInternetParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *SubnetUpdateInternetParam) SetSubnetId(v sacloud.ID) {
	p.SubnetId = v
}

func (p *SubnetUpdateInternetParam) GetSubnetId() sacloud.ID {
	return p.SubnetId
}
func (p *SubnetUpdateInternetParam) SetNextHop(v string) {
	p.NextHop = v
}

func (p *SubnetUpdateInternetParam) GetNextHop() string {
	return p.NextHop
}

// Ipv6InfoInternetParam is input parameters for the sacloud API
type Ipv6InfoInternetParam struct {
	input Input
}

// NewIpv6InfoInternetParam return new Ipv6InfoInternetParam
func NewIpv6InfoInternetParam() *Ipv6InfoInternetParam {
	return &Ipv6InfoInternetParam{}
}

// Initialize init Ipv6InfoInternetParam
func (p *Ipv6InfoInternetParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *Ipv6InfoInternetParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *Ipv6InfoInternetParam) fillValueToSkeleton() {

}

func (p *Ipv6InfoInternetParam) validate() error {
	var errors []error

	return utils.FlattenErrors(errors)
}

func (p *Ipv6InfoInternetParam) ResourceDef() *schema.Resource {
	return define.Resources["Internet"]
}

func (p *Ipv6InfoInternetParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["ipv6-info"]
}

func (p *Ipv6InfoInternetParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *Ipv6InfoInternetParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *Ipv6InfoInternetParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *Ipv6InfoInternetParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

// Ipv6EnableInternetParam is input parameters for the sacloud API
type Ipv6EnableInternetParam struct {
	input Input
}

// NewIpv6EnableInternetParam return new Ipv6EnableInternetParam
func NewIpv6EnableInternetParam() *Ipv6EnableInternetParam {
	return &Ipv6EnableInternetParam{}
}

// Initialize init Ipv6EnableInternetParam
func (p *Ipv6EnableInternetParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *Ipv6EnableInternetParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *Ipv6EnableInternetParam) fillValueToSkeleton() {

}

func (p *Ipv6EnableInternetParam) validate() error {
	var errors []error

	return utils.FlattenErrors(errors)
}

func (p *Ipv6EnableInternetParam) ResourceDef() *schema.Resource {
	return define.Resources["Internet"]
}

func (p *Ipv6EnableInternetParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["ipv6-enable"]
}

func (p *Ipv6EnableInternetParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *Ipv6EnableInternetParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *Ipv6EnableInternetParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *Ipv6EnableInternetParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

// Ipv6DisableInternetParam is input parameters for the sacloud API
type Ipv6DisableInternetParam struct {
	input Input
}

// NewIpv6DisableInternetParam return new Ipv6DisableInternetParam
func NewIpv6DisableInternetParam() *Ipv6DisableInternetParam {
	return &Ipv6DisableInternetParam{}
}

// Initialize init Ipv6DisableInternetParam
func (p *Ipv6DisableInternetParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *Ipv6DisableInternetParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *Ipv6DisableInternetParam) fillValueToSkeleton() {

}

func (p *Ipv6DisableInternetParam) validate() error {
	var errors []error

	return utils.FlattenErrors(errors)
}

func (p *Ipv6DisableInternetParam) ResourceDef() *schema.Resource {
	return define.Resources["Internet"]
}

func (p *Ipv6DisableInternetParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["ipv6-disable"]
}

func (p *Ipv6DisableInternetParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *Ipv6DisableInternetParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *Ipv6DisableInternetParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *Ipv6DisableInternetParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

// MonitorInternetParam is input parameters for the sacloud API
type MonitorInternetParam struct {
	End       string
	KeyFormat string
	Start     string

	input Input
}

// NewMonitorInternetParam return new MonitorInternetParam
func NewMonitorInternetParam() *MonitorInternetParam {
	return &MonitorInternetParam{
		KeyFormat: "sakuracloud.internet.{{.ID}}.nic"}
}

// Initialize init MonitorInternetParam
func (p *MonitorInternetParam) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *MonitorInternetParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *MonitorInternetParam) fillValueToSkeleton() {
	if utils.IsEmpty(p.End) {
		p.End = ""
	}
	if utils.IsEmpty(p.KeyFormat) {
		p.KeyFormat = ""
	}
	if utils.IsEmpty(p.Start) {
		p.Start = ""
	}

}

func (p *MonitorInternetParam) validate() error {
	var errors []error

	{
		validator := define.Resources["Internet"].Commands["monitor"].Params["end"].ValidateFunc
		errs := validator("--end", p.End)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := validateRequired
		errs := validator("--key-format", p.KeyFormat)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["Internet"].Commands["monitor"].Params["start"].ValidateFunc
		errs := validator("--start", p.Start)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return utils.FlattenErrors(errors)
}

func (p *MonitorInternetParam) ResourceDef() *schema.Resource {
	return define.Resources["Internet"]
}

func (p *MonitorInternetParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["monitor"]
}

func (p *MonitorInternetParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *MonitorInternetParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *MonitorInternetParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *MonitorInternetParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *MonitorInternetParam) SetEnd(v string) {
	p.End = v
}

func (p *MonitorInternetParam) GetEnd() string {
	return p.End
}
func (p *MonitorInternetParam) SetKeyFormat(v string) {
	p.KeyFormat = v
}

func (p *MonitorInternetParam) GetKeyFormat() string {
	return p.KeyFormat
}
func (p *MonitorInternetParam) SetStart(v string) {
	p.Start = v
}

func (p *MonitorInternetParam) GetStart() string {
	return p.Start
}
