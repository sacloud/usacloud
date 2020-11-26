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

package mobilegateway

import (
	"github.com/sacloud/usacloud/pkg/cmd/ccol"
	"github.com/sacloud/usacloud/pkg/output"
)

var defaultColumnDefs = []output.ColumnDef{
	ccol.Zone,
	ccol.ID,
	ccol.Name,
	ccol.Tags,
	ccol.Description,
	{
		Name:     "IPAddresses",
		Template: `{{ range .Interfaces }}{{ .IPAddress }}/{{ .NetworkMaskLen }} {{ end }}`,
	},
	{Name: "Internet", Template: "{{.InternetConnectionEnabled}}"},
	{Name: "InterDeviceCommunication", Template: "{{.InterDeviceCommunicationEnabled}}"},
	ccol.InstanceStatus,
}

/*
	ID                              types.ID
	Name                            string `validate:"required"`
	Description                     string `validate:"min=0,max=512"`
	Tags                            types.Tags
	Availability                    types.EAvailability
	Class                           string
	IconID                          types.ID `mapconv:"Icon.ID"`
	CreatedAt                       time.Time
	InstanceHostName                string                           `mapconv:"Instance.Host.Name"`
	InstanceHostInfoURL             string                           `mapconv:"Instance.Host.InfoURL"`
	InstanceStatus                  types.EServerInstanceStatus      `mapconv:"Instance.Status"`
	InstanceStatusChangedAt         time.Time                        `mapconv:"Instance.StatusChangedAt"`
	Interfaces                      []*MobileGatewayInterface        `json:",omitempty" mapconv:"[]Interfaces,recursive,omitempty"`
	ZoneID                          types.ID                         `mapconv:"Remark.Zone.ID"`
	InterfaceSettings               []*MobileGatewayInterfaceSetting `mapconv:"Settings.MobileGateway.[]Interfaces,recursive"`
	StaticRoutes                    []*MobileGatewayStaticRoute      `mapconv:"Settings.MobileGateway.[]StaticRoutes,recursive"`
	InternetConnectionEnabled       types.StringFlag                 `mapconv:"Settings.MobileGateway.InternetConnection.Enabled"`
	InterDeviceCommunicationEnabled types.StringFlag                 `mapconv:"Settings.MobileGateway.InterDeviceCommunication.Enabled"`
	SettingsHash                    string                           `json:",omitempty" mapconv:",omitempty"`
*/
