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

package cflag

type ConfirmParameter struct {
	AssumeYes bool `cli:"assumeyes,short=y,category=input,desc=Assume that the answer to any question which would be asked is yes" json:"-"`
}

func (p *ConfirmParameter) AssumeYesFlagValue() bool {
	return p.AssumeYes
}

type ConfirmParameterValueHandler interface {
	AssumeYesFlagValue() bool
}
