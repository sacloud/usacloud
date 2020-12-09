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

package webaccelerator

import (
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/cmd/core"
)

var Resource = &core.Resource{
	Name:             "web-accelerator",
	Aliases:          []string{"web-accel", "webaccel"},
	IsGlobalResource: true,
	Category:         core.ResourceCategoryOther,
}

func listAllFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	webAccelOp := sacloud.NewWebAccelOp(ctx.Client())
	searched, err := webAccelOp.List(ctx)
	if err != nil {
		return nil, err
	}
	var results []interface{}
	for _, v := range searched.WebAccels {
		results = append(results, v)
	}
	return results, nil
}
