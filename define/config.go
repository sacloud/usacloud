package define

import (
	"github.com/sacloud/usacloud/schema"
)

func ConfigResource() *schema.Resource {

	commands := map[string]*schema.Command{
		"list": {
			Type:             schema.CommandCustom,
			Aliases:          []string{"ls"},
			Params:           emptyParam(),
			UseCustomCommand: true,
			NeedlessConfirm:  true,
			NoOutput:         true,
		},
		"edit": {
			Type:                   schema.CommandCustom,
			Params:                 configEditParam(),
			UseCustomCommand:       true,
			UseCustomArgCompletion: true,
			NeedlessConfirm:        true,
			NoOutput:               true,
		},
		"use": {
			Type:                   schema.CommandCustom,
			Params:                 emptyParam(),
			UseCustomCommand:       true,
			UseCustomArgCompletion: true,
			NeedlessConfirm:        true,
			NoOutput:               true,
		},
		"show": {
			Type:                   schema.CommandCustom,
			Params:                 emptyParam(),
			UseCustomCommand:       true,
			UseCustomArgCompletion: true,
			NeedlessConfirm:        true,
			NoOutput:               true,
		},
		"current": {
			Type:             schema.CommandCustom,
			Params:           emptyParam(),
			UseCustomCommand: true,
			NeedlessConfirm:  true,
			NoOutput:         true,
		},
		"delete": {
			Type:                   schema.CommandCustom,
			Aliases:                []string{"rm"},
			Params:                 emptyParam(),
			UseCustomCommand:       true,
			UseCustomArgCompletion: true,
			NoOutput:               true,
		},
		"migrate": {
			Type:             schema.CommandCustom,
			Params:           emptyParam(),
			UseCustomCommand: true,
			NeedlessConfirm:  true,
			NoOutput:         true,
		},
	}

	return &schema.Resource{
		Commands:            commands,
		Aliases:             []string{"profile"},
		DefaultCommand:      "edit",
		Usage:               "A manage command of APIKey settings",
		ResourceCategory:    CategoryConfig,
		SkipApplyConfigFile: true,
	}
}

var AllowZones = []string{"is1a", "is1b", "tk1a", "tk1v"}
var AllowOutputTypes = []string{"table", "json", "csv", "tsv"}

func configEditParam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"token": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "API Token of SakuraCloud",
			Category:    "config",
			Order:       10,
		},
		"secret": {
			Type:        schema.TypeString,
			HandlerType: schema.HandlerNoop,
			Description: "API Secret of SakuraCloud",
			Category:    "config",
			Order:       20,
		},
		"zone": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "Target zone of SakuraCloud",
			ValidateFunc: validateInStrValues(AllowZones...),
			CompleteFunc: completeInStrValues(AllowZones...),
			Category:     "config",
			Order:        30,
		},
		"default-output-type": {
			Type:         schema.TypeString,
			HandlerType:  schema.HandlerNoop,
			Description:  "Default output format type",
			ValidateFunc: validateInStrValues(AllowOutputTypes...),
			CompleteFunc: completeInStrValues(AllowOutputTypes...),
			Category:     "config",
			Order:        40,
		},
	}
}
