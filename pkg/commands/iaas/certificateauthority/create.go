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

package certificateauthority

import (
	"time"

	"github.com/sacloud/iaas-api-go/types"
	"github.com/sacloud/iaas-service-go/certificateauthority/builder"
	cflag2 "github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/examples"
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
	cflag2.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag2.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag2.OutputParameter  `cli:",squash" mapconv:"-"`

	cflag2.NameParameter   `cli:",squash" mapconv:",squash"`
	cflag2.DescParameter   `cli:",squash" mapconv:",squash"`
	cflag2.TagsParameter   `cli:",squash" mapconv:",squash"`
	cflag2.IconIDParameter `cli:",squash" mapconv:",squash"`

	Country          string `validate:"required"`
	Organization     string `validate:"required"`
	OrganizationUnit []string
	CommonName       string `validate:"required"`

	ValidityPeriodHours int       `mapconv:"-" validate:"required"` // Customize()で時間数->NotAfterへの変換が行われる
	NotAfter            time.Time `cli:"-" json:"-"`                // Customize()の中で現在時刻+ValidityPeriodHoursが設定される

	ClientsData string                `cli:"clients" mapconv:"-" json:"-"`
	Clients     []*builder.ClientCert `cli:"-"`

	ServersData string                `cli:"servers" mapconv:"-" json:"-"`
	Servers     []*builder.ServerCert `cli:"-"`
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

	var clients []*builder.ClientCert
	if p.ClientsData != "" {
		if err := util.MarshalJSONFromPathOrContent(p.ClientsData, &clients); err != nil {
			return err
		}
	}
	p.Clients = append(p.Clients, clients...)

	var servers []*builder.ServerCert
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
		Clients: []*builder.ClientCert{
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
		Servers: []*builder.ServerCert{
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
