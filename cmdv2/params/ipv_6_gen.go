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

// ListIpv6Param is input parameters for the sacloud API
type ListIpv6Param struct {
	From       int
	Max        int
	Sort       []string
	Name       []string
	Ipv6netId  sacloud.ID
	InternetId sacloud.ID
	Id         []sacloud.ID

	input Input
}

// NewListIpv6Param return new ListIpv6Param
func NewListIpv6Param() *ListIpv6Param {
	return &ListIpv6Param{}
}

// Initialize init ListIpv6Param
func (p *ListIpv6Param) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ListIpv6Param) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *ListIpv6Param) fillValueToSkeleton() {
	if utils.IsEmpty(p.From) {
		p.From = 0
	}
	if utils.IsEmpty(p.Max) {
		p.Max = 0
	}
	if utils.IsEmpty(p.Sort) {
		p.Sort = []string{""}
	}
	if utils.IsEmpty(p.Name) {
		p.Name = []string{""}
	}
	if utils.IsEmpty(p.Ipv6netId) {
		p.Ipv6netId = sacloud.ID(0)
	}
	if utils.IsEmpty(p.InternetId) {
		p.InternetId = sacloud.ID(0)
	}
	if utils.IsEmpty(p.Id) {
		p.Id = []sacloud.ID{}
	}

}

func (p *ListIpv6Param) validate() error {
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
		validator := define.Resources["IPv6"].Commands["list"].Params["ipv6net-id"].ValidateFunc
		errs := validator("--ipv-6net-id", p.Ipv6netId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["IPv6"].Commands["list"].Params["internet-id"].ValidateFunc
		errs := validator("--internet-id", p.InternetId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["IPv6"].Commands["list"].Params["id"].ValidateFunc
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

	return utils.FlattenErrors(errors)
}

func (p *ListIpv6Param) ResourceDef() *schema.Resource {
	return define.Resources["IPv6"]
}

func (p *ListIpv6Param) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["list"]
}

func (p *ListIpv6Param) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ListIpv6Param) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ListIpv6Param) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ListIpv6Param) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ListIpv6Param) SetFrom(v int) {
	p.From = v
}

func (p *ListIpv6Param) GetFrom() int {
	return p.From
}
func (p *ListIpv6Param) SetMax(v int) {
	p.Max = v
}

func (p *ListIpv6Param) GetMax() int {
	return p.Max
}
func (p *ListIpv6Param) SetSort(v []string) {
	p.Sort = v
}

func (p *ListIpv6Param) GetSort() []string {
	return p.Sort
}
func (p *ListIpv6Param) SetName(v []string) {
	p.Name = v
}

func (p *ListIpv6Param) GetName() []string {
	return p.Name
}
func (p *ListIpv6Param) SetIpv6netId(v sacloud.ID) {
	p.Ipv6netId = v
}

func (p *ListIpv6Param) GetIpv6netId() sacloud.ID {
	return p.Ipv6netId
}
func (p *ListIpv6Param) SetInternetId(v sacloud.ID) {
	p.InternetId = v
}

func (p *ListIpv6Param) GetInternetId() sacloud.ID {
	return p.InternetId
}
func (p *ListIpv6Param) SetId(v []sacloud.ID) {
	p.Id = v
}

func (p *ListIpv6Param) GetId() []sacloud.ID {
	return p.Id
}

// PtrAddIpv6Param is input parameters for the sacloud API
type PtrAddIpv6Param struct {
	Hostname string

	input Input
}

// NewPtrAddIpv6Param return new PtrAddIpv6Param
func NewPtrAddIpv6Param() *PtrAddIpv6Param {
	return &PtrAddIpv6Param{}
}

// Initialize init PtrAddIpv6Param
func (p *PtrAddIpv6Param) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *PtrAddIpv6Param) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *PtrAddIpv6Param) fillValueToSkeleton() {
	if utils.IsEmpty(p.Hostname) {
		p.Hostname = ""
	}

}

func (p *PtrAddIpv6Param) validate() error {
	var errors []error

	{
		validator := validateRequired
		errs := validator("--hostname", p.Hostname)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return utils.FlattenErrors(errors)
}

func (p *PtrAddIpv6Param) ResourceDef() *schema.Resource {
	return define.Resources["IPv6"]
}

func (p *PtrAddIpv6Param) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["ptr-add"]
}

func (p *PtrAddIpv6Param) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *PtrAddIpv6Param) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *PtrAddIpv6Param) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *PtrAddIpv6Param) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *PtrAddIpv6Param) SetHostname(v string) {
	p.Hostname = v
}

func (p *PtrAddIpv6Param) GetHostname() string {
	return p.Hostname
}

// PtrReadIpv6Param is input parameters for the sacloud API
type PtrReadIpv6Param struct {
	input Input
}

// NewPtrReadIpv6Param return new PtrReadIpv6Param
func NewPtrReadIpv6Param() *PtrReadIpv6Param {
	return &PtrReadIpv6Param{}
}

// Initialize init PtrReadIpv6Param
func (p *PtrReadIpv6Param) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *PtrReadIpv6Param) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *PtrReadIpv6Param) fillValueToSkeleton() {

}

func (p *PtrReadIpv6Param) validate() error {
	var errors []error

	return utils.FlattenErrors(errors)
}

func (p *PtrReadIpv6Param) ResourceDef() *schema.Resource {
	return define.Resources["IPv6"]
}

func (p *PtrReadIpv6Param) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["ptr-read"]
}

func (p *PtrReadIpv6Param) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *PtrReadIpv6Param) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *PtrReadIpv6Param) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *PtrReadIpv6Param) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

// PtrUpdateIpv6Param is input parameters for the sacloud API
type PtrUpdateIpv6Param struct {
	Hostname string

	input Input
}

// NewPtrUpdateIpv6Param return new PtrUpdateIpv6Param
func NewPtrUpdateIpv6Param() *PtrUpdateIpv6Param {
	return &PtrUpdateIpv6Param{}
}

// Initialize init PtrUpdateIpv6Param
func (p *PtrUpdateIpv6Param) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *PtrUpdateIpv6Param) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *PtrUpdateIpv6Param) fillValueToSkeleton() {
	if utils.IsEmpty(p.Hostname) {
		p.Hostname = ""
	}

}

func (p *PtrUpdateIpv6Param) validate() error {
	var errors []error

	{
		validator := validateRequired
		errs := validator("--hostname", p.Hostname)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return utils.FlattenErrors(errors)
}

func (p *PtrUpdateIpv6Param) ResourceDef() *schema.Resource {
	return define.Resources["IPv6"]
}

func (p *PtrUpdateIpv6Param) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["ptr-update"]
}

func (p *PtrUpdateIpv6Param) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *PtrUpdateIpv6Param) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *PtrUpdateIpv6Param) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *PtrUpdateIpv6Param) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *PtrUpdateIpv6Param) SetHostname(v string) {
	p.Hostname = v
}

func (p *PtrUpdateIpv6Param) GetHostname() string {
	return p.Hostname
}

// PtrDeleteIpv6Param is input parameters for the sacloud API
type PtrDeleteIpv6Param struct {
	input Input
}

// NewPtrDeleteIpv6Param return new PtrDeleteIpv6Param
func NewPtrDeleteIpv6Param() *PtrDeleteIpv6Param {
	return &PtrDeleteIpv6Param{}
}

// Initialize init PtrDeleteIpv6Param
func (p *PtrDeleteIpv6Param) Initialize(in Input) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *PtrDeleteIpv6Param) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

func (p *PtrDeleteIpv6Param) fillValueToSkeleton() {

}

func (p *PtrDeleteIpv6Param) validate() error {
	var errors []error

	return utils.FlattenErrors(errors)
}

func (p *PtrDeleteIpv6Param) ResourceDef() *schema.Resource {
	return define.Resources["IPv6"]
}

func (p *PtrDeleteIpv6Param) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["ptr-delete"]
}

func (p *PtrDeleteIpv6Param) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *PtrDeleteIpv6Param) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *PtrDeleteIpv6Param) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *PtrDeleteIpv6Param) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}
