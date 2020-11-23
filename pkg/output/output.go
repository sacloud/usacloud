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

type Option interface {
	OutputTypeFlagValue() string
	FormatFlagValue() string
	QuietFlagValue() bool
	QueryFlagValue() string
}
