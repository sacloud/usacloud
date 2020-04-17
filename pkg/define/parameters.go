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

package define

import (
	"fmt"

	"github.com/sacloud/libsacloud/sacloud"

	"github.com/sacloud/usacloud/pkg/schema"
)

var (
	paramName = &schema.Schema{
		Type:         schema.TypeString,
		HandlerType:  schema.HandlerPathThrough,
		Description:  "set resource display name",
		ValidateFunc: validateStrLen(1, 64),
		Category:     "common",
		Order:        500,
	}
	paramRequiredName = &schema.Schema{
		Type:         schema.TypeString,
		HandlerType:  schema.HandlerPathThrough,
		Description:  "set resource display name",
		Required:     true,
		ValidateFunc: validateStrLen(1, 64),
		Category:     "common",
		Order:        510,
	}
	paramDescription = &schema.Schema{
		Type:         schema.TypeString,
		HandlerType:  schema.HandlerPathThrough,
		Description:  "set resource description",
		Aliases:      []string{"desc"},
		ValidateFunc: validateStrLen(0, 254),
		Category:     "common",
		Order:        520,
	}
	paramTags = &schema.Schema{
		Type:         schema.TypeStringList,
		HandlerType:  schema.HandlerPathThrough,
		Description:  "set resource tags",
		ValidateFunc: validateStringSlice(validateStrLen(1, 32)),
		Category:     "common",
		Order:        530,
	}
)

func getParamResourceShortID(resourceName string, digit int) *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeId,
		HandlerType:  schema.HandlerPathThrough,
		Description:  fmt.Sprintf("set %s", resourceName),
		Required:     true,
		ValidateFunc: validateSakuraShortID(digit),
	}
}

var paramIconResourceID = &schema.Schema{
	Type:            schema.TypeId,
	HandlerType:     schema.HandlerPathThrough,
	DestinationProp: "SetIconByID",
	Description:     "set Icon ID",
	ValidateFunc:    validateSakuraID(),
	Category:        "common",
	Order:           540,
}

var CommonListParam = map[string]*schema.Schema{

	"name": {
		Type:          schema.TypeStringList,
		HandlerType:   schema.HandlerOrParams,
		Description:   "set filter by name(s)",
		ConflictsWith: []string{"id"},
		Category:      "filter",
		Order:         1,
	},
	"id": {
		Type:            schema.TypeIdList,
		HandlerType:     schema.HandlerAndParams,
		DestinationProp: "ID",
		Description:     "set filter by id(s)",
		ConflictsWith:   []string{"name"},
		ValidateFunc:    validateIntSlice(validateSakuraID()),
		Category:        "filter",
		Order:           2,
	},
	"from": {
		Type:            schema.TypeInt,
		HandlerType:     schema.HandlerPathThrough,
		Aliases:         []string{"offset"},
		DestinationProp: "SetOffset",
		Description:     "set offset",
		Category:        "limit-offset",
		Order:           1,
	},
	"max": {
		Type:            schema.TypeInt,
		HandlerType:     schema.HandlerPathThrough,
		Aliases:         []string{"limit"},
		DestinationProp: "SetLimit",
		Description:     "set limit",
		Category:        "limit-offset",
		Order:           2,
	},
	//"exclude": {
	//	Type:        schema.TypeStringList,
	//	HandlerType: schema.HandlerPathThroughEach,
	//	Description: "set exclude field(s)",
	//},
	//"include": {
	//	Type:        schema.TypeStringList,
	//	HandlerType: schema.HandlerPathThroughEach,
	//	Description: "set include field(s)",
	//},
	"sort": {
		Type:        schema.TypeStringList,
		HandlerType: schema.HandlerSort,
		Description: "set field(s) for sort",
		Category:    "sort",
		Order:       1,
	},
}

func emptyParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{}
}

var scopeCondStrings = []string{"user", "shared"}
var paramScopeCond = map[string]*schema.Schema{
	"scope": {
		Type:         schema.TypeString,
		HandlerType:  schema.HandlerFilterBy,
		Description:  "set filter by scope('user' or 'shared')",
		ValidateFunc: validateInStrValues(scopeCondStrings...),
		Category:     "filter",
		Order:        3,
	},
}

var paramTagsCond = map[string]*schema.Schema{
	"tags": {
		Type:         schema.TypeStringList,
		Aliases:      []string{"selector"},
		HandlerType:  schema.HandlerFilterFunc,
		FilterFunc:   filterListByTags,
		Description:  "set filter by tags(AND)",
		Category:     "filter",
		ValidateFunc: validateStringSlice(validateStrLen(1, 32)),
		Order:        4,
	},
}

func filterListByTags(_ []interface{}, item interface{}, param interface{}) bool {

	type tagHandler interface {
		HasTag(target string) bool
	}

	tagHolder, ok := item.(tagHandler)
	if !ok {
		return false
	}

	paramTags := param.([]string)

	// 完全一致 + AND条件
	res := true
	for _, p := range paramTags {
		if !tagHolder.HasTag(p) {
			res = false
			break
		}
	}
	return res
}

var paramSourceArchiveIDCond = map[string]*schema.Schema{
	"source-archive-id": {
		Type:         schema.TypeId,
		HandlerType:  schema.HandlerFilterFunc,
		FilterFunc:   filterBySourceArchiveID,
		Description:  "set filter by source-archive-id",
		Category:     "filter",
		ValidateFunc: validateSakuraID(),
		Order:        5,
	},
}

var paramSourceDiskCond = map[string]*schema.Schema{
	"source-disk-id": {
		Type:         schema.TypeId,
		HandlerType:  schema.HandlerFilterFunc,
		FilterFunc:   filterBySourceDiskID,
		Description:  "set filter by source-disk-id",
		Category:     "filter",
		ValidateFunc: validateSakuraID(),
		Order:        6,
	},
}

func filterBySourceArchiveID(_ []interface{}, item interface{}, param interface{}) bool {

	type archiveIDHandler interface {
		GetSourceArchiveID() sacloud.ID
	}

	archiveIDHolder, ok := item.(archiveIDHandler)
	if !ok {
		return false
	}

	id := param.(sacloud.ID)
	if id == 0 {
		return true
	}

	return archiveIDHolder.GetSourceArchiveID() == id
}

func filterBySourceDiskID(_ []interface{}, item interface{}, param interface{}) bool {

	type diskIDHandler interface {
		GetSourceDiskID() sacloud.ID
	}

	diskIDHolder, ok := item.(diskIDHandler)
	if !ok {
		return false
	}

	id := param.(sacloud.ID)
	if id == 0 {
		return true
	}

	return diskIDHolder.GetSourceDiskID() == id
}

var paramClassCond = map[string]*schema.Schema{
	"class": {
		Type:        schema.TypeStringList,
		HandlerType: schema.HandlerFilterBy,
		Description: "set filter by class(es)",
		Category:    "filter",
		Order:       7,
	},
}
