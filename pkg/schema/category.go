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

package schema

import "math"

type Category struct {
	Key         string
	DisplayName string
	Order       int
}

var DefaultResourceCategory = &Category{
	Key:         "default",
	DisplayName: "",
	Order:       math.MaxInt32,
}

var DefaultCommandCategory = &Category{
	Key:         "default",
	DisplayName: "",
	Order:       math.MaxInt32,
}

var FilterParamCategory = &Category{
	Key:         "filter",
	DisplayName: "Filter options",
	Order:       math.MaxInt32 - 60,
}
var LimitOffsetParamCategory = &Category{
	Key:         "limit-offset",
	DisplayName: "Limit/Offset options",
	Order:       math.MaxInt32 - 50,
}
var SortParamCategory = &Category{
	Key:         "sort",
	DisplayName: "Sort options",
	Order:       math.MaxInt32 - 40,
}

var CommonParamCategory = &Category{
	Key:         "common",
	DisplayName: "Common options",
	Order:       math.MaxInt32 - 30,
}

var InputParamCategory = &Category{
	Key:         "Input",
	DisplayName: "Input options",
	Order:       math.MaxInt32 - 20,
}

var OutputParamCategory = &Category{
	Key:         "output",
	DisplayName: "Output options",
	Order:       math.MaxInt32 - 10,
}

var DefaultParamCategory = &Category{
	Key:         "default",
	DisplayName: "Other options",
	Order:       math.MaxInt32,
}
