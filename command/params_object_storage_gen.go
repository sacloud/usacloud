// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package command

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// ListObjectStorageParam is input parameters for the sacloud API
type ListObjectStorageParam struct {
	AccessKey string
	SecretKey string
	Bucket    string
}

// NewListObjectStorageParam return new ListObjectStorageParam
func NewListObjectStorageParam() *ListObjectStorageParam {
	return &ListObjectStorageParam{}
}

// Validate checks current values in model
func (p *ListObjectStorageParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--access-key", p.AccessKey)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--secret-key", p.SecretKey)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ListObjectStorageParam) getResourceDef() *schema.Resource {
	return define.Resources["ObjectStorage"]
}

func (p *ListObjectStorageParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["list"]
}

func (p *ListObjectStorageParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ListObjectStorageParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ListObjectStorageParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ListObjectStorageParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ListObjectStorageParam) SetAccessKey(v string) {
	p.AccessKey = v
}

func (p *ListObjectStorageParam) GetAccessKey() string {
	return p.AccessKey
}
func (p *ListObjectStorageParam) SetSecretKey(v string) {
	p.SecretKey = v
}

func (p *ListObjectStorageParam) GetSecretKey() string {
	return p.SecretKey
}
func (p *ListObjectStorageParam) SetBucket(v string) {
	p.Bucket = v
}

func (p *ListObjectStorageParam) GetBucket() string {
	return p.Bucket
}

// PutObjectStorageParam is input parameters for the sacloud API
type PutObjectStorageParam struct {
	AccessKey   string
	SecretKey   string
	Bucket      string
	ContentType string
}

// NewPutObjectStorageParam return new PutObjectStorageParam
func NewPutObjectStorageParam() *PutObjectStorageParam {
	return &PutObjectStorageParam{

		ContentType: "application/octet-stream",
	}
}

// Validate checks current values in model
func (p *PutObjectStorageParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--access-key", p.AccessKey)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--secret-key", p.SecretKey)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *PutObjectStorageParam) getResourceDef() *schema.Resource {
	return define.Resources["ObjectStorage"]
}

func (p *PutObjectStorageParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["put"]
}

func (p *PutObjectStorageParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *PutObjectStorageParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *PutObjectStorageParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *PutObjectStorageParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *PutObjectStorageParam) SetAccessKey(v string) {
	p.AccessKey = v
}

func (p *PutObjectStorageParam) GetAccessKey() string {
	return p.AccessKey
}
func (p *PutObjectStorageParam) SetSecretKey(v string) {
	p.SecretKey = v
}

func (p *PutObjectStorageParam) GetSecretKey() string {
	return p.SecretKey
}
func (p *PutObjectStorageParam) SetBucket(v string) {
	p.Bucket = v
}

func (p *PutObjectStorageParam) GetBucket() string {
	return p.Bucket
}
func (p *PutObjectStorageParam) SetContentType(v string) {
	p.ContentType = v
}

func (p *PutObjectStorageParam) GetContentType() string {
	return p.ContentType
}

// GetObjectStorageParam is input parameters for the sacloud API
type GetObjectStorageParam struct {
	AccessKey string
	SecretKey string
	Bucket    string
}

// NewGetObjectStorageParam return new GetObjectStorageParam
func NewGetObjectStorageParam() *GetObjectStorageParam {
	return &GetObjectStorageParam{}
}

// Validate checks current values in model
func (p *GetObjectStorageParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--access-key", p.AccessKey)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--secret-key", p.SecretKey)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *GetObjectStorageParam) getResourceDef() *schema.Resource {
	return define.Resources["ObjectStorage"]
}

func (p *GetObjectStorageParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["get"]
}

func (p *GetObjectStorageParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *GetObjectStorageParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *GetObjectStorageParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *GetObjectStorageParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *GetObjectStorageParam) SetAccessKey(v string) {
	p.AccessKey = v
}

func (p *GetObjectStorageParam) GetAccessKey() string {
	return p.AccessKey
}
func (p *GetObjectStorageParam) SetSecretKey(v string) {
	p.SecretKey = v
}

func (p *GetObjectStorageParam) GetSecretKey() string {
	return p.SecretKey
}
func (p *GetObjectStorageParam) SetBucket(v string) {
	p.Bucket = v
}

func (p *GetObjectStorageParam) GetBucket() string {
	return p.Bucket
}

// DeleteObjectStorageParam is input parameters for the sacloud API
type DeleteObjectStorageParam struct {
	SecretKey string
	Bucket    string
	AccessKey string
}

// NewDeleteObjectStorageParam return new DeleteObjectStorageParam
func NewDeleteObjectStorageParam() *DeleteObjectStorageParam {
	return &DeleteObjectStorageParam{}
}

// Validate checks current values in model
func (p *DeleteObjectStorageParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--secret-key", p.SecretKey)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--access-key", p.AccessKey)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *DeleteObjectStorageParam) getResourceDef() *schema.Resource {
	return define.Resources["ObjectStorage"]
}

func (p *DeleteObjectStorageParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["delete"]
}

func (p *DeleteObjectStorageParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *DeleteObjectStorageParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *DeleteObjectStorageParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *DeleteObjectStorageParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *DeleteObjectStorageParam) SetSecretKey(v string) {
	p.SecretKey = v
}

func (p *DeleteObjectStorageParam) GetSecretKey() string {
	return p.SecretKey
}
func (p *DeleteObjectStorageParam) SetBucket(v string) {
	p.Bucket = v
}

func (p *DeleteObjectStorageParam) GetBucket() string {
	return p.Bucket
}
func (p *DeleteObjectStorageParam) SetAccessKey(v string) {
	p.AccessKey = v
}

func (p *DeleteObjectStorageParam) GetAccessKey() string {
	return p.AccessKey
}
