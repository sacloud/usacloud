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

package tools

import (
	"fmt"

	"github.com/sacloud/usacloud/pkg/naming"
)

func (r *Resource) ServiceSourceFileName() string {
	return fmt.Sprintf("%s_services_gen.go", naming.ToSnakeCase(r.Name))
}

func (r *Resource) ChildResourceServiceSourceFileName(child *Resource) string {
	return fmt.Sprintf("%s_%s_services_gen.go", naming.ToSnakeCase(r.Name), naming.ToSnakeCase(child.Name))
}

func (r *Resource) ServiceRepositoryName() string {
	switch r.PlatformName {
	case "phy":
		return "github.com/sacloud/phy-service-go"
	case "objectstorage":
		return "github.com/sacloud/object-storage-service-go"
	case "webaccel":
		return "github.com/sacloud/webaccel-service-go"
	case "iaas", "":
		return "github.com/sacloud/iaas-service-go"
	}
	panic(fmt.Sprintf("unsupported platform name: %s", r.PlatformName))
}
