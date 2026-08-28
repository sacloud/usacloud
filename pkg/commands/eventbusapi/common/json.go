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

package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/sacloud/usacloud/pkg/util"
)

func DecodeJSONPathOrContent(pathOrContent string, destination interface{}) error {
	data, err := util.BytesFromPathOrContent(pathOrContent)
	if err != nil {
		return err
	}
	return DecodeJSON(bytes.NewReader(data), destination)
}

func DecodeJSON(reader io.Reader, destination interface{}) error {
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(destination); err != nil {
		return fmt.Errorf("decode JSON: %w", err)
	}
	if err := decoder.Decode(new(interface{})); err != io.EOF {
		if err == nil {
			return fmt.Errorf("decode JSON: multiple values are not allowed")
		}
		return fmt.Errorf("decode JSON: %w", err)
	}
	return nil
}
