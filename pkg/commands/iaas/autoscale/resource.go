// Copyright 2017-2023 The sacloud/usacloud Authors
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

package autoscale

import (
	"reflect"

	"github.com/sacloud/iaas-service-go/autoscale"
	"github.com/sacloud/usacloud/pkg/commands/iaas/category"
	"github.com/sacloud/usacloud/pkg/core"
)

var Resource = &core.Resource{
	PlatformName:     "iaas",
	Name:             "auto-scale",
	Aliases:          []string{"autoscale"},
	ServiceType:      reflect.TypeOf(&autoscale.Service{}),
	Category:         category.ResourceCategoryComputing,
	IsGlobalResource: true,
}
