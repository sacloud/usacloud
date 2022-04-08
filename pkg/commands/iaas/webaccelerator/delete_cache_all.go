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
	cflag2 "github.com/sacloud/usacloud/pkg/cflag"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/core"
)

var deleteCacheAllCommand = &core.Command{
	Name:     "delete-cache-all",
	Aliases:  []string{"cache-delete-all"},
	Category: "cache",
	Order:    20,

	ParameterInitializer: func() interface{} {
		return newDeleteCacheAllParameter()
	},
	ListAllFunc: listAllFunc,
	Func:        deleteCacheAllFunc,
}

type deleteCacheAllParameter struct {
	cflag2.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag2.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag2.CommonParameter  `cli:",squash" mapconv:"-"`

	Domain string `validate:"required"`
}

func newDeleteCacheAllParameter() *deleteCacheAllParameter {
	return &deleteCacheAllParameter{}
}

func init() {
	Resource.AddCommand(deleteCacheAllCommand)
}

func deleteCacheAllFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	p, ok := parameter.(*deleteCacheAllParameter)
	if !ok {
		return nil, fmt.Errorf("got invalid parameter type: %#v", parameter)
	}
	webAccelOp := iaas.NewWebAccelOp(ctx.Client().(iaas.APICaller))

	return nil, webAccelOp.DeleteAllCache(ctx, &iaas.WebAccelDeleteAllCacheRequest{Domain: p.Domain})
}
