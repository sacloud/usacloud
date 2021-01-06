// Copyright 2017-2021 The Usacloud Authors
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
	"log"
	"reflect"
)

type ServiceFuncMeta struct {
	HasRequestValue    bool
	HasReturnValue     bool
	IsReturnValueSlice bool
}

func (c *Command) ServiceFuncReturnValueType() *ServiceFuncMeta {
	svcType := c.Resource.ServiceType
	method, ok := svcType.MethodByName(c.ServiceFuncName())
	if !ok {
		log.Fatalf(
			"reading meta info of service func failed: service-type: %s func-name: %s error: not found",
			svcType.Name(), c.ServiceFuncName())
	}

	hasRequest := false
	numIn := method.Type.NumIn()
	if numIn > 2 { // [0]はレシーバー、[1]はctx
		hasRequest = true
	}

	numOut := method.Type.NumOut()
	if numOut < 2 { // 戻り値なし(ないはず) or errorのみの場合
		return &ServiceFuncMeta{HasRequestValue: hasRequest, HasReturnValue: false, IsReturnValueSlice: false}
	}

	firstReturnValue := method.Type.Out(0)
	if firstReturnValue.Kind() == reflect.Slice {
		return &ServiceFuncMeta{HasRequestValue: hasRequest, HasReturnValue: true, IsReturnValueSlice: true}
	}
	return &ServiceFuncMeta{HasRequestValue: hasRequest, HasReturnValue: true, IsReturnValueSlice: false}
}
