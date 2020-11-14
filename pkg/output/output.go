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

package output

type Output interface {
	Print(Contents) error
}

type Formatter interface {
	ColumnDefs() []ColumnDef
}

type DefaultFormatter struct{}

func (f *DefaultFormatter) IncludeFields() []string {
	return []string{}
}

func (f *DefaultFormatter) ExcludeFields() []string {
	return []string{}
}

func (f *DefaultFormatter) ColumnDefs() []ColumnDef {
	return []ColumnDef{}
}

func (f *DefaultFormatter) TableType() TableType {
	return TableSimple
}

type Option interface {
	OutputTypeFlagValue() string
	FormatFlagValue() string
	FormatFileFlagValue() string
	QuietFlagValue() bool
	QueryFlagValue() string
	QueryFileFlagValue() string
}

type TableType int //go:generate stringer -type=OutputTableType :: manual
const (
	TableDetail TableType = iota
	TableSimple
)

type tableWriter interface {
	append(map[string]string)
	render()
}

type ColumnDef struct {
	Name         string
	Sources      []string
	Format       string
	ValueMapping []map[string]string
	FormatFunc   func(values map[string]string) string
}

func (d *ColumnDef) GetSources() []string {
	if len(d.Sources) == 0 {
		return []string{d.Name}
	}
	return d.Sources
}

func (d *ColumnDef) GetFormat() string {
	if d.Format == "" {
		return "%s"
	}
	return d.Format
}
