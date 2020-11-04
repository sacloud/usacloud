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

package base

type OutputParameter struct {
	OutputType string `cli:",short=o,aliases=out,category=output,desc=Output format: one of the following [table/json/yaml]"`
	Quiet      bool   `cli:",short=q,category=output,desc=Output IDs only"`
	Format     string `cli:",aliases=fmt,category=output,desc=Output format in Go templates"`
	FormatFile string `cli:",category=output,desc=Output format in Go templates(from file)"`
	Query      string `cli:",category=output,desc=JMESPath query"`
	QueryFile  string `cli:",category=output,desc=JMESPath query(from file)"`
}
