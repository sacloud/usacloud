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

package certificateauthority

import (
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud/types"

	"github.com/sacloud/libsacloud/v2/helper/service/certificateauthority"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/cmd/examples"
	"github.com/sacloud/usacloud/pkg/util"
)

var createCommand = &core.Command{
	Name:     "create",
	Category: "basic",
	Order:    20,

	ColumnDefs: defaultServiceColumnDefs,

	ParameterInitializer: func() interface{} {
		return newCreateParameter()
	},
}

type createParameter struct {
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag.NameParameter   `cli:",squash" mapconv:",squash"`
	cflag.DescParameter   `cli:",squash" mapconv:",squash"`
	cflag.TagsParameter   `cli:",squash" mapconv:",squash"`
	cflag.IconIDParameter `cli:",squash" mapconv:",squash"`

	Country          string `validate:"required"`
	Organization     string `validate:"required"`
	OrganizationUnit []string
	CommonName       string `validate:"required"`

	ValidityPeriodHours int       `mapconv:"-" validate:"required"` // Customize()で時間数->NotAfterへの変換が行われる
	NotAfter            time.Time `cli:"-" json:"-"`                // Customize()の中で現在時刻+ValidityPeriodHoursが設定される

	ClientsData string                             `cli:"clients" mapconv:"-" json:"-"`
	Clients     []*certificateauthority.ClientCert `cli:"-"`

	ServersData string                             `cli:"servers" mapconv:"-" json:"-"`
	Servers     []*certificateauthority.ServerCert `cli:"-"`
}

func newCreateParameter() *createParameter {
	return &createParameter{}
}

func init() {
	Resource.AddCommand(createCommand)
}

// Customize パラメータ変換処理
func (p *createParameter) Customize(_ cli.Context) error {
	p.NotAfter = time.Now().Add(time.Duration(p.ValidityPeriodHours) * time.Hour)

	var clients []*certificateauthority.ClientCert
	if p.ClientsData != "" {
		if err := util.MarshalJSONFromPathOrContent(p.ClientsData, &clients); err != nil {
			return err
		}
	}
	p.Clients = append(p.Clients, clients...)

	var servers []*certificateauthority.ServerCert
	if p.ServersData != "" {
		if err := util.MarshalJSONFromPathOrContent(p.ServersData, &servers); err != nil {
			return err
		}
	}
	p.Servers = append(p.Servers, servers...)
	return nil
}

func (p *createParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &createParameter{
		NameParameter:       examples.Name,
		DescParameter:       examples.Description,
		TagsParameter:       examples.Tags,
		IconIDParameter:     examples.IconID,
		Country:             "JP",
		Organization:        "usacloud",
		OrganizationUnit:    []string{"ou1", "ou2"},
		CommonName:          "example.usacloud.jp",
		ValidityPeriodHours: 24 * 365,
		Clients: []*certificateauthority.ClientCert{
			{
				Country:                   "JP",
				Organization:              "usacloud",
				OrganizationUnit:          []string{"ou1", "ou2"},
				CommonName:                "client.usacloud.jp",
				NotAfter:                  time.Now().Add(24 * time.Hour * 365),
				IssuanceMethod:            types.ECertificateAuthorityIssuanceMethod(examples.OptionsString("certificate_authority_issuance_method")),
				EMail:                     "example@example.com",
				CertificateSigningRequest: "-----BEGIN CERTIFICATE REQUEST-----\n...",
				PublicKey:                 "-----BEGIN PUBLIC KEY-----\n...",
				Hold:                      true,
			},
		},
		Servers: []*certificateauthority.ServerCert{
			{
				Country:                   "JP",
				Organization:              "usacloud",
				OrganizationUnit:          []string{"ou1", "ou2"},
				CommonName:                "client.usacloud.jp",
				NotAfter:                  time.Now().Add(24 * time.Hour * 365),
				SANs:                      []string{"www1.usacloud.jp", "www2.usacloud.jp"},
				CertificateSigningRequest: "-----BEGIN CERTIFICATE REQUEST-----\n...",
				PublicKey:                 "-----BEGIN PUBLIC KEY-----\n...",
				Hold:                      true,
			},
		},
	}
}
