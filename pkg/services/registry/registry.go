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

package registry

import (
	"github.com/sacloud/usacloud/pkg/cli"
)

type ServiceFunc func(ctx cli.Context, parameter interface{}) ([]interface{}, error)
type ListAllFunc func(ctx cli.Context, parameter interface{}) ([]interface{}, error)

var platformServiceFuncRegistry = map[string]map[string]ServiceFunc{}
var platformListAllFuncRegistry = map[string]map[string]ListAllFunc{}

func SetDefaultServiceFunc(platformName, resourceName, commandName string, fn ServiceFunc) {
	serviceFuncRegistry(platformName)[resourceName+commandName] = fn
}

func SetDefaultListAllFunc(platformName, resourceName, commandName string, fn ListAllFunc) {
	listAllFuncRegistry(platformName)[resourceName+commandName] = fn
}

func DefaultServiceFunc(platformName, resourceName, commandName string) (ServiceFunc, bool) {
	fn, ok := serviceFuncRegistry(platformName)[resourceName+commandName]
	return fn, ok
}

func DefaultListAllFunc(platformName, resourceName, commandName string) (ListAllFunc, bool) {
	fn, ok := listAllFuncRegistry(platformName)[resourceName+commandName]
	return fn, ok
}

func serviceFuncRegistry(platformName string) map[string]ServiceFunc {
	name := platform(platformName)
	v, ok := platformServiceFuncRegistry[name]
	if !ok {
		v = map[string]ServiceFunc{}
		platformServiceFuncRegistry[name] = v
	}
	return v
}

func listAllFuncRegistry(platformName string) map[string]ListAllFunc {
	name := platform(platformName)
	v, ok := platformListAllFuncRegistry[name]
	if !ok {
		v = map[string]ListAllFunc{}
		platformListAllFuncRegistry[name] = v
	}
	return v
}

func platform(name string) string {
	if name == "" {
		return "iaas"
	}
	return name
}
