package define

import (
	"fmt"
	"github.com/sacloud/usacloud/schema"
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
		Type:         schema.TypeInt64,
		HandlerType:  schema.HandlerPathThrough,
		Description:  fmt.Sprintf("set %s", resourceName),
		Required:     true,
		ValidateFunc: validateSakuraShortID(digit),
	}
}

var paramIconResourceID = &schema.Schema{
	Type:            schema.TypeInt64,
	HandlerType:     schema.HandlerPathThrough,
	DestinationProp: "SetIconByID",
	Description:     "set Icon ID",
	ValidateFunc:    validateSakuraID(),
	CompleteFunc:    completeIconID(),
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
		Type:            schema.TypeIntList,
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

var scopeCondStrings = []string{"user", "shared"}
var paramScopeCond = map[string]*schema.Schema{
	"scope": {
		Type:         schema.TypeString,
		HandlerType:  schema.HandlerFilterBy,
		Description:  "set filter by scope('user' or 'shared')",
		ValidateFunc: validateInStrValues(scopeCondStrings...),
		CompleteFunc: completeInStrValues(scopeCondStrings...),
		Category:     "filter",
		Order:        3,
	},
}
