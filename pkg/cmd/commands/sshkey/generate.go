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

package sshkey

import (
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/output"
)

var generateCommand = &core.Command{
	Name:     "generate",
	Category: "basic",
	Order:    25,

	ColumnDefs: append(defaultColumnDefs, output.ColumnDef{Name: "PrivateKey"}),

	ParameterInitializer: func() interface{} {
		return newGenerateParameter()
	},
}

type generateParameter struct {
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag.NameParameter `cli:",squash" mapconv:",squash"`
	cflag.DescParameter `cli:",squash" mapconv:",squash"`
	PassPhrase          string
}

func newGenerateParameter() *generateParameter {
	return &generateParameter{}
}

func init() {
	Resource.AddCommand(generateCommand)
}
