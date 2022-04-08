// Copyright 2017-2022 The Usacloud Authors
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

type CommonParameterValueHolder interface {
	ParametersFlagValue() string
	GenerateSkeletonFlagValue() bool
}

type ExampleParameterValueHolder interface {
	ExampleFlagValue() bool
}

// CommonParameter 全コマンド共通フィールド
type CommonParameter struct {
	Parameters       string `cli:",category=input,desc=Input parameters in JSON format" json:"-"`
	GenerateSkeleton bool   `cli:",category=input,aliases=skeleton,desc=Output skeleton of parameters with JSON format" json:"-"`
	Example          bool   `cli:",category=example,desc=Output example parameters with JSON format" json:"-"`
}

func (p *CommonParameter) ParametersFlagValue() string {
	return p.Parameters
}

func (p *CommonParameter) GenerateSkeletonFlagValue() bool {
	return p.GenerateSkeleton
}

func (p *CommonParameter) ExampleFlagValue() bool {
	return p.Example
}
