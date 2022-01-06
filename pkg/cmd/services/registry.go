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

package services

import "github.com/sacloud/usacloud/pkg/cli"

type ServiceFunc func(ctx cli.Context, parameter interface{}) ([]interface{}, error)
type ListAllFunc func(ctx cli.Context, parameter interface{}) ([]interface{}, error)

var defaultServiceFuncRegistry = map[string]ServiceFunc{}
var defaultListAllFuncRegistry = map[string]ListAllFunc{}

func setDefaultServiceFunc(resourceName, commandName string, fn ServiceFunc) {
	defaultServiceFuncRegistry[resourceName+commandName] = fn
}

func setDefaultListAllFunc(resourceName, commandName string, fn ListAllFunc) {
	defaultListAllFuncRegistry[resourceName+commandName] = fn
}

func DefaultServiceFunc(resourceName, commandName string) (ServiceFunc, bool) {
	fn, ok := defaultServiceFuncRegistry[resourceName+commandName]
	return fn, ok
}

func DefaultListAllFunc(resourceName, commandName string) (ListAllFunc, bool) {
	fn, ok := defaultListAllFuncRegistry[resourceName+commandName]
	return fn, ok
}
