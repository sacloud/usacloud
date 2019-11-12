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

package cli

import (
	"github.com/sacloud/usacloud/schema"
	"gopkg.in/urfave/cli.v2"
)

// Commands are added by tools/gen-cli-commands
var Commands []*cli.Command

// ResourceCategoryMap are added by tools/gen-cli-commands
//
// ex. ResourceCategoryMap["server"] -> *schema.Category
var ResourceCategoryMap = map[string]*schema.Category{}

func AppendResourceCategoryMap(r string, category *schema.Category) {
	ResourceCategoryMap[r] = category
}

// CommandCategoryMap are added by tools/gen-cli-commands
//
// ex. CommandCategoryMap["server"]["create"] -> *schema.Category
var CommandCategoryMap = map[string]map[string]*schema.Category{}

func AppendCommandCategoryMap(r string, c string, category *schema.Category) {
	if _, ok := CommandCategoryMap[r]; !ok {
		CommandCategoryMap[r] = map[string]*schema.Category{}
	}
	CommandCategoryMap[r][c] = category
}

// FlagCategoryMap are added by tools/gen-cli-commands
//
// ex. FlagCategoryMap["server"]["build"]["list"] -> *schema.Category
var FlagCategoryMap = map[string]map[string]map[string]*schema.Category{}

func AppendFlagCategoryMap(r string, c string, f string, category *schema.Category) {

	if _, ok := FlagCategoryMap[r]; !ok {
		FlagCategoryMap[r] = map[string]map[string]*schema.Category{}
	}

	if _, ok := FlagCategoryMap[r][c]; !ok {
		FlagCategoryMap[r][c] = map[string]*schema.Category{}
	}

	FlagCategoryMap[r][c][f] = category
}
