// Copyright 2016-2021 The Libsacloud Authors
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

package service

import (
	"fmt"
	"time"

	"github.com/sacloud/libsacloud/v2/pkg/mapconv"
	"github.com/sacloud/libsacloud/v2/pkg/size"
	"github.com/sacloud/libsacloud/v2/sacloud"
)

func HandleNotFoundError(err error, ignoreNotFoundError bool) error {
	if ignoreNotFoundError && sacloud.IsNotFoundError(err) {
		return nil // ignore: 404 not found
	}
	return err
}

func MonitorCondition(start, end time.Time) (*sacloud.MonitorCondition, error) {
	e := end
	if e.IsZero() {
		e = time.Now()
	}

	s := start
	if s.IsZero() {
		s = e.Add(-1 * time.Hour)
	}
	if !(s.Unix() <= e.Unix()) {
		return nil, fmt.Errorf("start(%s) or end(%s) is invalid", start.String(), end.String())
	}
	return &sacloud.MonitorCondition{Start: s, End: e}, nil
}

func RequestConvertTo(source interface{}, dest interface{}) error {
	decoder := &mapconv.Decoder{
		Config: &mapconv.DecoderConfig{
			TagName: "request",
			FilterFuncs: map[string]mapconv.FilterFunc{
				"gb_to_mb": gbToMb,
			},
		},
	}
	return decoder.ConvertTo(source, dest)
}

func gbToMb(v interface{}) (interface{}, error) {
	s, ok := v.(int)
	if !ok {
		return nil, fmt.Errorf("invalid size value: %v", v)
	}
	return size.GiBToMiB(s), nil
}
