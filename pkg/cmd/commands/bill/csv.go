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

package bill

import (
	"github.com/sacloud/libsacloud/v2/helper/service/bill"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
)

var csvCommand = &core.Command{
	Name:       "csv",
	Category:   "basic",
	Order:      30,
	NoProgress: true,

	ColumnDefs: csvColumnDefs,

	SelectorType: core.SelectorTypeRequireSingle, // TODO libsacloud service側はbill csvでID省略可能(最新の請求情報が対象となる)がusacloudは現状非対応

	ParameterInitializer: func() interface{} {
		return newCSVParameter()
	},
	// ListAllが通常と異なるシグニチャのため個別対応する
	ListAllFunc: func(ctx cli.Context, _ interface{}) ([]interface{}, error) {
		svc := bill.New(ctx.Client())
		res, err := svc.ListWithContext(ctx, &bill.ListRequest{})
		if err != nil {
			return nil, err
		}

		var results []interface{}
		for _, v := range res {
			results = append(results, v)
		}
		return results, nil
	},
}

type csvParameter struct {
	cflag.IDParameter     `cli:",squash" mapconv:",squash"`
	cflag.InputParameter  `cli:",squash" mapconv:"-"`
	cflag.OutputParameter `cli:",squash" mapconv:"-"`
}

func newCSVParameter() *csvParameter {
	return &csvParameter{}
}

func init() {
	Resource.AddCommand(csvCommand)
}
