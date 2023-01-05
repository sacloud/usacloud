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

package query

import (
	"encoding/json"
	"fmt"

	"github.com/fatih/structs"
	"github.com/itchyny/gojq"
	"github.com/jmespath/go-jmespath"
)

func ByJMESPath(v interface{}, query string, printer func(interface{}) error) (err error) {
	defer func() {
		ret := recover()
		if ret != nil {
			err = fmt.Errorf("jmespath: failed to process query: %s", ret)
		}
	}()
	var result interface{}
	result, err = jmespath.Search(query, v)
	if err != nil {
		return err
	}

	return printer(result)
}

func ByGoJQ(input interface{}, query string, printer func(interface{}) error) (err error) {
	q, err := gojq.Parse(query)
	if err != nil {
		return fmt.Errorf("gojq parse failed: %v", err)
	}

	// gojqにinputを渡す前に[]map[string]interface{}へ変換しておく
	// see: https://pkg.go.dev/github.com/itchyny/gojq#readme-usage-as-a-library
	// > the query input should have type []interface{} for an array and map[string]interface{} for a map
	mv, err := convertInputToMap(input)
	if err != nil {
		return fmt.Errorf("failed to convert to map: %v", err)
	}
	iter := q.Run(mv)

	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			return fmt.Errorf("gojq: %s", err.Error())
		}
		if err := printer(v); err != nil {
			return err
		}
	}
	return nil
}

func convertInputToMap(input interface{}) (interface{}, error) {
	var inputs []interface{}

	switch input := input.(type) {
	case map[string]interface{}:
		return input, nil // 既にmap[string]interface{}の場合はそのまま返す
	case []interface{}:
		inputs = input
	default:
		inputs = []interface{}{input}
	}

	for i, v := range inputs {
		if !structs.IsStruct(v) {
			continue
		}
		mv, err := struct2map(v)
		if err != nil {
			return nil, err
		}
		inputs[i] = mv
	}
	return inputs, nil
}

func struct2map(v interface{}) (interface{}, error) {
	out := make(map[string]interface{})
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, err
	}
	return out, nil
}
