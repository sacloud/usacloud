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
