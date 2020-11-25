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

package localrouter

import (
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/util"
)

var createCommand = &core.Command{
	Name:     "create",
	Category: "basic",
	Order:    20,

	ColumnDefs: defaultColumnDefs,

	ParameterInitializer: func() interface{} {
		return newCreateParameter()
	},
}

type createParameter struct {
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	Name        string   `validate:"required"`
	Description string   `validate:"description"`
	Tags        []string `validate:"tags"`
	IconID      types.ID

	Switch struct {
		Code     string
		Category string
		ZoneID   string
	}

	Interface struct {
		VirtualIPAddress string   `validate:"omitempty,ipv4"`
		IPAddress        []string `validate:"omitempty,min=2,max=2,dive,ipv4"`
		NetworkMaskLen   int      `validate:"omitempty,min=8,max=28"`
		VRID             int      `validate:"omitempty"`
	} `cli:",squash"`

	PeersData string                     `cli:"peers" mapconv:"-"`
	Peers     []*sacloud.LocalRouterPeer `cli:"-"`

	StaticRoutesData string                            `cli:"static-routes" mapconv:"-"`
	StaticRoutes     []*sacloud.LocalRouterStaticRoute `cli:"-"`
}

func newCreateParameter() *createParameter {
	return &createParameter{}
}

func init() {
	Resource.AddCommand(createCommand)
}

// Customize パラメータ変換処理
func (p *createParameter) Customize() error {
	if p.PeersData != "" {
		var peers []*sacloud.LocalRouterPeer
		if err := util.MarshalJSONFromPathOrContent(p.PeersData, &peers); err != nil {
			return err
		}
		p.Peers = append(p.Peers, peers...)
	}
	if p.StaticRoutesData != "" {
		var staticRoutes []*sacloud.LocalRouterStaticRoute
		if err := util.MarshalJSONFromPathOrContent(p.StaticRoutesData, &staticRoutes); err != nil {
			return err
		}
		p.StaticRoutes = append(p.StaticRoutes, staticRoutes...)
	}

	return nil
}
