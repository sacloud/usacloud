// Copyright 2017-2021 The Usacloud Authors
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

package cflag

type OutputParameter struct {
	OutputType  string `cli:",short=o,aliases=out,category=output,desc=Output format options: [table/json/yaml]" validate:"omitempty,output_type" json:"-"`
	Quiet       bool   `cli:",short=q,category=output,desc=Output IDs only" json:"-"`
	Format      string `cli:",aliases=fmt,category=output,desc=Output format in Go templates" json:"-"`
	Query       string `cli:",category=output,desc=Query for JSON output" json:"-"`
	QueryDriver string `cli:",category=output,desc=Name of the driver that handles queries to JSON output options: [jmespath/jq]" json:"-" validate:"omitempty,oneof=jmespath jq"`
}

func (p *OutputParameter) OutputTypeFlagValue() string {
	return p.OutputType
}

func (p *OutputParameter) QuietFlagValue() bool {
	return p.Quiet
}

func (p *OutputParameter) FormatFlagValue() string {
	return p.Format
}

func (p *OutputParameter) QueryFlagValue() string {
	return p.Query
}

func (p *OutputParameter) QueryDriverFlagValue() string {
	return p.QueryDriver
}
