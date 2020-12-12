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

package iface

import (
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
)

var listCommand = &core.Command{
	Name:               "list",
	Aliases:            []string{"ls", "find", "select"},
	Category:           "basic",
	Order:              10,
	ServiceFuncAltName: "Find",
	NoProgress:         true,

	ColumnDefs: defaultColumnDefs,

	ParameterInitializer: func() interface{} {
		return newListParameter()
	},
}

type listParameter struct {
	cflag.ZoneParameter        `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter      `cli:",squash" mapconv:"-"`
	cflag.LimitOffsetParameter `cli:",squash" mapconv:",squash"`
	cflag.OutputParameter      `cli:",squash" mapconv:"-"`

	MACAddresses      []string `cli:",category=filter" validate:"omitempty,dive,mac"`
	ServerIDs         []string `cli:"server-ids,category=filter"`
	ServerNames       []string `cli:"server-names,category=filter"`
	PacketFilterIDs   []string `cli:"packet-filer-ids,category=filter"`
	PacketFilterNames []string `cli:"packet-filter-names,category=filter"`
}

func newListParameter() *listParameter {
	return &listParameter{}
}

func init() {
	Resource.AddCommand(listCommand)
}
