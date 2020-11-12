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

package disk

import (
	"github.com/sacloud/usacloud/pkg/cmd/base"
)

var listCommand = &base.Command{
	Name:               "list",
	Aliases:            []string{"ls", "find", "select"},
	Category:           "basics",
	Order:              10,
	ServiceFuncAltName: "Find",
	NoConfirm:          true,
	NoProgress:         true,

	ParameterInitializer: func() interface{} {
		return newListParameter()
	},
}

type listParameter struct {
	base.ZoneParameter `cli:",squash" mapconv:",squash"`

	Names              []string `cli:",category=filter"`
	Tags               []string `cli:",category=filter"`
	base.FindParameter `cli:",squash" mapconv:",squash"`

	base.OutputParameter `cli:",squash" mapconv:"-"`
}

func newListParameter() *listParameter {
	return &listParameter{
		// TODO デフォルト値はここで設定する
	}
}

func init() {
	Resource.AddCommand(listCommand)
}
