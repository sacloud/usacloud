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

package webaccelerator

import (
	"fmt"

	"github.com/sacloud/iaas-api-go"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
)

var readCertificateCommand = &core.Command{
	Name:       "read-certificate",
	Aliases:    []string{"certificate-read", "cert-read"},
	Category:   "certificate",
	Order:      10,
	NoProgress: true,

	ColumnDefs: certificateColumnDefs,

	SelectorType: core.SelectorTypeRequireSingle,

	ParameterInitializer: func() interface{} {
		return newReadCertificateParameter()
	},
	ListAllFunc: listAllFunc,
	Func:        readCertificateFunc,
}

type readCertificateParameter struct {
	cflag.IDParameter     `cli:",squash" mapconv:",squash"`
	cflag.CommonParameter `cli:",squash" mapconv:"-"`
	cflag.OutputParameter `cli:",squash" mapconv:"-"`
}

func newReadCertificateParameter() *readCertificateParameter {
	return &readCertificateParameter{}
}

func init() {
	Resource.AddCommand(readCertificateCommand)
}

func readCertificateFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	p, ok := parameter.(*readCertificateParameter)
	if !ok {
		return nil, fmt.Errorf("got invalid parameter type: %#v", parameter)
	}
	webAccelOp := iaas.NewWebAccelOp(ctx.Client().(iaas.APICaller))
	result, err := webAccelOp.ReadCertificate(ctx, p.ID)
	if err != nil {
		return nil, err
	}
	if result == nil || result.Current == nil {
		return nil, nil
	}
	return []interface{}{result}, nil
}
