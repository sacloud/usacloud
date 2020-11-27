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

package packetfilter

import (
	"reflect"

	"github.com/sacloud/libsacloud/v2/helper/service/packetfilter"
	"github.com/sacloud/usacloud/pkg/cmd/core"
)

var Resource = &core.Resource{
	Name:        "packet-filter",
	Aliases:     []string{"packetfilter"},
	ServiceType: reflect.TypeOf(&packetfilter.Service{}),
	Category:    core.ResourceCategoryStorage,
	CommandCategories: []core.Category{
		{
			Key:         "basic",
			DisplayName: "Basic Commands",
			Order:       10,
		},
		{
			Key:         "operation",
			DisplayName: "Operation Commands",
			Order:       20,
		},
		{
			Key:         "other",
			DisplayName: "Other Commands",
			Order:       1000,
		},
	},
}
