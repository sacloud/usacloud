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
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/commands/iaas/category"
	"github.com/sacloud/usacloud/pkg/core"
	"github.com/sacloud/webaccel-api-go"
)

var Resource = &core.Resource{
	PlatformName:     "webaccel",
	Name:             "web-accelerator",
	Usage:            "SubCommands for WebAccelerator",
	Aliases:          []string{"web-accel", "webaccel"},
	IsGlobalResource: true,
	Category:         category.ResourceCategoryWebAccel,
}

func listAllFunc(ctx cli.Context, parameter interface{}) ([]interface{}, error) {
	webAccelOp := webaccel.NewOp(ctx.Client().(*webaccel.Client))
	searched, err := webAccelOp.List(ctx)
	if err != nil {
		return nil, err
	}
	var results []interface{}
	for _, v := range searched.Sites {
		results = append(results, v)
	}
	return results, nil
}
