package define

import (
	"fmt"
	"github.com/sacloud/usacloud/schema"
)

var (
	paramID   = getParamResourceID("resource ID")
	paramName = &schema.Schema{
		Type:         schema.TypeString,
		HandlerType:  schema.HandlerPathThrough,
		Description:  "set resource display name",
		ValidateFunc: validateStrLen(1, 64),
	}
	paramRequiredName = &schema.Schema{
		Type:         schema.TypeString,
		HandlerType:  schema.HandlerPathThrough,
		Description:  "set resource display name",
		Required:     true,
		ValidateFunc: validateStrLen(1, 64),
	}
	paramDescription = &schema.Schema{
		Type:         schema.TypeString,
		HandlerType:  schema.HandlerPathThrough,
		Description:  "set resource description",
		Aliases:      []string{"desc"},
		ValidateFunc: validateStrLen(0, 254),
	}
	paramTags = &schema.Schema{
		Type:         schema.TypeStringList,
		HandlerType:  schema.HandlerPathThrough,
		Description:  "set resource tags",
		ValidateFunc: validateStringSlice(validateStrLen(1, 32)),
	}
)

func getParamResourceID(resourceName string) *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeInt64,
		HandlerType:  schema.HandlerPathThrough,
		Description:  fmt.Sprintf("set %s", resourceName),
		Required:     true,
		ValidateFunc: validateSakuraID(),
	}
}

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
}

var CommonListParam map[string]*schema.Schema = map[string]*schema.Schema{

	"name": {
		Type:          schema.TypeStringList,
		HandlerType:   schema.HandlerOrParams,
		Description:   "set filter by name(s)",
		ConflictsWith: []string{"id"},
	},
	"id": {
		Type:            schema.TypeIntList,
		HandlerType:     schema.HandlerAndParams,
		DestinationProp: "ID",
		Description:     "set filter by id(s)",
		ConflictsWith:   []string{"name"},
		ValidateFunc:    validateIntSlice(validateSakuraID()),
	},
	"from": {
		Type:            schema.TypeInt,
		HandlerType:     schema.HandlerPathThrough,
		DestinationProp: "SetOffset",
		Description:     "set offset",
	},
	"max": {
		Type:            schema.TypeInt,
		HandlerType:     schema.HandlerPathThrough,
		DestinationProp: "SetLimit",
		Description:     "set limit",
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
	},
}
