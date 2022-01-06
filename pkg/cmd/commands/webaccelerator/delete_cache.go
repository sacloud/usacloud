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

	"github.com/sacloud/usacloud/pkg/output"

	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/cflag"
	"github.com/sacloud/usacloud/pkg/cmd/core"
)

var deleteCacheCommand = &core.Command{
	Name:     "delete-cache",
	Aliases:  []string{"cache-delete"},
	Category: "cache",
	Order:    10,

	ColumnDefs: []output.ColumnDef{
		{Name: "URL"},
		{Name: "Status"},
		{Name: "Result"},
	},

	ParameterInitializer: func() interface{} {
		return newDeleteCacheParameter()
	},
	Func: deleteCacheFunc,
}

type deleteCacheParameter struct {
	cflag.IDParameter      `cli:",squash" mapconv:",squash"`
	cflag.ConfirmParameter `cli:",squash" mapconv:"-"`
	cflag.CommonParameter  `cli:",squash" mapconv:"-"`
	cflag.OutputParameter  `cli:",squash" mapconv:"-"`

	URLs []string `cli:"url" validate:"required"`
}

func newDeleteCacheParameter() *deleteCacheParameter {
	return &deleteCacheParameter{}
}

func init() {
	Resource.AddCommand(deleteCacheCommand)
}

func deleteCacheFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	p, ok := parameter.(*deleteCacheParameter)
	if !ok {
		return nil, fmt.Errorf("got invalid parameter type: %#v", parameter)
	}
	webAccelOp := sacloud.NewWebAccelOp(ctx.Client())
	deleteResults, err := webAccelOp.DeleteCache(ctx, &sacloud.WebAccelDeleteCacheRequest{URL: p.URLs})
	if err != nil {
		return nil, err
	}
	var results []interface{}
	for _, r := range deleteResults {
		results = append(results, r)
	}
	return results, nil
}
