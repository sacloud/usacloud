// Copyright 2017-2025 The sacloud/usacloud Authors
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

package webaccelerator

import (
	"fmt"

	"github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/usacloud/pkg/util"
	"github.com/sacloud/webaccel-api-go"
)

var createCertificateCommand = &core.Command{
	Name:     "create-certificate",
	Aliases:  []string{"certificate-create", "cert-create"},
	Category: "certificate",
	Order:    20,

	ColumnDefs: certificateColumnDefs,

	SelectorType: core.SelectorTypeRequireSingle,

	ParameterInitializer: func() interface{} {
		return newCreateCertificateParameter()
	},
	ListAllFunc: listAllFunc,
	Func:        createCertificateFunc,
}

type createCertificateParameter struct {
	cflag.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`

	CertificateChain string `validate:"required"`
	Key              string `validate:"required"`
}

func newCreateCertificateParameter() *createCertificateParameter {
	return &createCertificateParameter{}
}

func init() {
	Resource.AddCommand(createCertificateCommand)
}

func createCertificateFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	p, ok := parameter.(*createCertificateParameter)
	if !ok {
		return nil, fmt.Errorf("got invalid parameter type: %#v", parameter)
	}

	certs, err := util.StringFromPathOrContent(p.CertificateChain)
	if err != nil {
		return nil, err
	}
	key, err := util.StringFromPathOrContent(p.Key)
	if err != nil {
		return nil, err
	}

	webAccelOp := webaccel.NewOp(ctx.Client().(*webaccel.Client))
	result, err := webAccelOp.CreateCertificate(ctx, p.ID, &webaccel.CreateOrUpdateCertificateRequest{
		CertificateChain: certs,
		Key:              key,
	})
	if err != nil {
		return nil, err
	}
	if result == nil || result.Current == nil {
		return nil, nil
	}
	return []interface{}{result}, nil
}

func (p *createCertificateParameter) ExampleParameters(ctx cli.Context) interface{} {
	return &createCertificateParameter{
		CertificateChain: "/path/to/your/certificate/chain | -----BEGIN CERTIFICATE-----\n...",
		Key:              "/path/to/your/private-key | -----BEGIN RSA PRIVATE KEY-----\n...",
	}
}
