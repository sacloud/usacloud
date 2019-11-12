// Copyright 2017-2019 The Usacloud Authors
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

package funcs

type MonitorValues []MonitorValue
type MonitorValue map[string]string

func (v MonitorValues) Len() int {
	return len(v)
}

func (v MonitorValues) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v MonitorValues) Less(i, j int) bool {
	_, hasIndex := v[i]["Index"]
	if hasIndex {
		if v[i]["UnixTime"] == v[j]["UnixTime"] {
			return v[i]["Index"] < v[j]["Index"]
		}
	}
	return v[i]["UnixTime"] < v[j]["UnixTime"]
}
