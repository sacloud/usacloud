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

package category

import (
	"math"
)

var ParameterCategories = []*Category{
	{
		Key:         "common",
		DisplayName: "Common options",
		Order:       50,
	},
	{
		Key:         "plan",
		DisplayName: "Plan options",
		Order:       60,
	},
	{
		Key:         "diskedit",
		DisplayName: "Edit disk options",
		Order:       200,
	},
	{
		Key:         "network",
		DisplayName: "Network options",
		Order:       200,
	},
	{
		Key:         "health",
		DisplayName: "Health check options",
		Order:       200,
	},
	{
		Key:         "filter",
		DisplayName: "Filter options",
		Order:       math.MaxInt32 - 100,
	},
	{
		Key:         "limit-offset",
		DisplayName: "Limit/Offset options",
		Order:       math.MaxInt32 - 90,
	},
	{
		Key:         "sort",
		DisplayName: "Sort options",
		Order:       math.MaxInt32 - 80,
	},
	{
		Key:         "monitor",
		DisplayName: "Monitor options",
		Order:       math.MaxInt32 - 70,
	},
	{
		Key:         "delete",
		DisplayName: "Delete options",
		Order:       math.MaxInt32 - 60,
	},
	{
		Key:         "zone",
		DisplayName: "Zone options",
		Order:       math.MaxInt32 - 50,
	},
	{
		Key:         "error",
		DisplayName: "Error handling options",
		Order:       math.MaxInt32 - 40,
	},
	{
		Key:         "wait",
		DisplayName: "Wait options",
		Order:       math.MaxInt32 - 30,
	},
	{
		Key:         "input",
		DisplayName: "Input options",
		Order:       math.MaxInt32 - 20,
	},
	{
		Key:         "output",
		DisplayName: "Output options",
		Order:       math.MaxInt32 - 10,
	},
	{
		Key:         "example",
		DisplayName: "Parameter example",
		Order:       math.MaxInt32 - 5,
	},
	{
		Key:         "default",
		DisplayName: "Other options",
		Order:       math.MaxInt32,
	},
}
